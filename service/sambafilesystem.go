package service

import (
	"github.com/astaxie/beego"
	"github.com/hirochachacha/go-smb2"
	"gpm/models"
	"gpm/tools"
	"net"
	"os"
	"strings"
)

type SambaFileSystem struct {
	RootPath  string
	fs        *smb2.Share
	shareName string
	session   smb2.Session
}

func (s *SambaFileSystem) ReadByte(parentDir string, fileName string) ([]byte, error) {
	return s.fs.ReadFile(s.getTargetPath(parentDir, fileName))
}
func (s *SambaFileSystem) ReadText(parentDir string, fileName string) (string, error) {
	readByte, err := s.ReadByte(parentDir, fileName)
	return string(readByte), err
}

func (s *SambaFileSystem) Mkdir(parentDir string, fileName string) error {
	return s.fs.MkdirAll(s.getTargetPath(parentDir, fileName), os.ModePerm)
}

func (s *SambaFileSystem) RmDir(parentDir string, fileName string) error {
	return s.fs.RemoveAll(s.getTargetPath(parentDir, fileName))
}

/**
  一般路径为 /测试目录/业务文档
   其中 测试目录为共享名称  通过Mount方法挂载文件系统
   业务文档为目录名称，如果需要处理可通过路径：业务文档/apollo/a.xt访问文件，注意根目录不需要/
*/
func (s *SambaFileSystem) ListRoot() ([]models.Node, error) {
	//未绑定sharename,直接列表sharename
	if s.RootPath == tools.PathSeparator {
		fileArray, _ := s.session.ListSharenames()
		nodeList := make([]models.Node, len(fileArray))
		for index, fi := range fileArray {
			node := models.Node{
				Title:       fi,
				Expand:      false,
				Contextmenu: true,
				IsDir:       true,
				Children:    []models.Node{},
				DirPath:     tools.PathSeparator,
			}
			nodeList[index] = node
		}
		return nodeList, nil
	}
	return s.ListDir("")
}
func (s *SambaFileSystem) IsDir(destPath string) bool {
	formatDestPath := s.getPath(destPath)
	fi, _ := s.fs.Lstat(formatDestPath)
	return fi.IsDir()
}
func (s *SambaFileSystem) getTargetPath(dirPth string, cfileName string) string {
	fomatDirPath := s.getPath(dirPth)
	if fomatDirPath == "" {
		return cfileName
	} else {
		return fomatDirPath + tools.PathSeparator + cfileName
	}
}
func (s *SambaFileSystem) getPath(dirPth string) string {
	formatDirPath := tools.FormatPath(dirPth)
	if s.RootPath == tools.PathSeparator {
		s.shareName = tools.GetRootName(formatDirPath)
		s.fs, _ = s.session.Mount(s.shareName)
	}
	formatDirPath = strings.TrimLeft(formatDirPath, tools.PathSeparator)
	formatDirPath = tools.TrimLeft(strings.Split(formatDirPath, s.shareName)[1])
	return formatDirPath
}

/**
  /测试目录/apollo
  shareName:="测试目录"
  path:="apollo"
  /测试目录
  shareName:="测试目录"
  path:=""
*/
func (s *SambaFileSystem) ListDir(dirPth string) ([]models.Node, error) {
	fileArray, _ := s.fs.ReadDir(s.getPath(dirPth))
	nodeList := make([]models.Node, len(fileArray))
	for index, fi := range fileArray {
		node := models.Node{
			Title:       fi.Name(),
			Expand:      false,
			Contextmenu: true,
			IsDir:       true,
			Children:    []models.Node{},
			DirPath:     dirPth,
		}
		if fi.IsDir() {
			node.IsDir = true
		} else {
			node.IsDir = false
		}
		nodeList[index] = node
	}
	return nodeList, nil
}
func (s *SambaFileSystem) DeleteFile(parentDir string, fileName string) error {
	return s.fs.Remove(s.getTargetPath(parentDir, fileName))
}
func (s *SambaFileSystem) CreateFile(parentDir string, fileName string) error {
	f, err := s.fs.Create(s.getTargetPath(parentDir, fileName))
	defer f.Close()
	return err
}
func (s *SambaFileSystem) SaveTextFile(parentDir string, fileName string, content string, policyType os.FileMode) error {
	return s.fs.WriteFile(s.getTargetPath(parentDir, fileName), []byte(content), policyType)
}
func (s *SambaFileSystem) Rename(srcDir string, src string, dest string) error {
	srcPath := s.getTargetPath(srcDir, src)
	destPath := s.getTargetPath(srcDir, dest)
	return s.fs.Rename(srcPath, destPath)
}
func (s *SambaFileSystem) Ping() error {
	_, err := s.session.ListSharenames()
	return err
}

/**
---------------------------------------
 文件系统工厂类负责读取配置参数生成文件系统实例
---------------------------------------
*/
type SambaFileSystemFactory struct {
}

func (sam *SambaFileSystemFactory) Create(prefix string) (FileSystem, error) {

	sambahost := beego.AppConfig.String(prefix + "sambahost")
	sambaport := beego.AppConfig.String(prefix + "sambaport")
	sambauser := beego.AppConfig.String(prefix + "sambauser")
	sambapassword := beego.AppConfig.String(prefix + "sambapassword")
	rootpath := beego.AppConfig.String(prefix + ".rootpath")

	conn, err := net.Dial("tcp", sambahost+":"+sambaport)
	if err != nil {
		panic(err)
	}
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     sambauser,
			Password: sambapassword,
		},
	}

	session, err := d.Dial(conn)
	if err != nil {
		panic(err)
	}
	formatRootPath := tools.FormatPath(rootpath)
	shareName := tools.GetRootName(formatRootPath)
	shareMount, _ := session.Mount(shareName)
	fileSystem := SambaFileSystem{RootPath: formatRootPath, session: *session, fs: shareMount, shareName: shareName}
	return &fileSystem, nil
}
func (s *SambaFileSystemFactory) Name() string {
	return "service.SambaFileSystemFactory"
}
