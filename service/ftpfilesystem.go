package service

import (
	"bytes"
	"errors"
	"github.com/astaxie/beego"
	"gopkg.in/dutchcoders/goftp.v1"
	"gpm/models"
	"gpm/tools"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type FtpFileSystem struct {
	RootPath string
	fs       *goftp.FTP
}

func (s *FtpFileSystem) ReadByte(parentDir string, fileName string) ([]byte, error) {
	formatDestPath := tools.FormatPath(parentDir)
	var text []byte
	_, e := s.fs.Retr(formatDestPath+tools.PathSeparator+fileName, func(r io.Reader) error {
		a, _ := ioutil.ReadAll(r)
		text = a
		return nil
	})
	return text, e
}
func (s *FtpFileSystem) ReadText(parentDir string, fileName string) (string, error) {
	readByte, err := s.ReadByte(parentDir, fileName)
	return string(readByte), err
}

func (s *FtpFileSystem) Mkdir(parentDir string, fileName string) error {
	formatDestPath := tools.FormatPath(parentDir)
	return s.fs.Mkd(formatDestPath + tools.PathSeparator + fileName)
}

func (s *FtpFileSystem) RmDir(parentDir string, fileName string) error {
	formatDestPath := tools.FormatPath(parentDir)
	return s.fs.Rmd(formatDestPath + tools.PathSeparator + fileName)
}
func (s *FtpFileSystem) ExistFile(parentDir string, fileName string) (bool, error) {
	return false, nil
}

/**
  一般路径为 /测试目录/业务文档
   其中 测试目录为共享名称  通过Mount方法挂载文件系统
   业务文档为目录名称，如果需要处理可通过路径：业务文档/apollo/a.xt访问文件，注意根目录不需要/
*/
func (s *FtpFileSystem) ListRoot() ([]models.Node, error) {
	return s.ListDir("", "")
}
func (s *FtpFileSystem) IsDir(destPath string) (bool, error) {
	formatDestPath := tools.FormatPath(destPath)
	index := strings.LastIndex(formatDestPath, tools.PathSeparator)
	dirName := formatDestPath[0:index]
	lastName := formatDestPath[index+1:]
	listFile, _ := s.fs.List(dirName)
	ifHave := false
	for _, curFile := range listFile {
		if strings.Index(curFile, "<") > 0 {
			reReplace, _ := regexp.Compile("[ |\t]+")
			splitFile := strings.Split(reReplace.ReplaceAllString(curFile, " "), " ")
			typeStr := splitFile[2]
			filename := strings.TrimSpace(splitFile[3])
			if lastName == filename {
				ifHave = true
				if "<DIR>" == typeStr {
					return true, nil
				}
			}
		} else {
			splitFile := strings.Split(curFile, ";")
			typeStr := strings.Split(splitFile[0], "=")[1]
			filename := ""
			isDir := false
			if "dir" == typeStr {
				filename = strings.TrimSpace(splitFile[2])
				isDir = true
			} else {
				filename = strings.TrimSpace(splitFile[3])
			}
			if lastName == filename {
				ifHave = true
				if isDir {
					return true, nil
				}
			}
		}
	}
	if ifHave {
		return false, nil
	} else {
		return false, errors.New("文件不存在")
	}

}

/**
  /测试目录/apollo
  shareName:="测试目录"
  path:="apollo"
  /测试目录
  shareName:="测试目录"
  path:=""
*/
func (s *FtpFileSystem) ListDir(dirPth string, trimPrefix string) ([]models.Node, error) {
	formatDestPath := tools.FormatPath(dirPth)
	listFile, _ := s.fs.List(formatDestPath)
	reReplace, _ := regexp.Compile("[ |\t]+")
	nodeList := make([]models.Node, len(listFile))
	for index, curFile := range listFile {
		//09-29-20 09:50AM       <DIR>          doc
		//08-05-20 03:31PM       <DIR>          gradle
		if strings.Index(curFile, "<") > 0 {
			splitFile := strings.Split(reReplace.ReplaceAllString(curFile, " "), " ")
			typeStr := splitFile[2]
			filename := strings.TrimSpace(splitFile[3])
			node := models.Node{
				Title:       filename,
				FileName:    filename,
				Expand:      false,
				Contextmenu: true,
				IsDir:       true,
				Children:    []models.Node{},
				DirPath:     dirPth,
			}
			if "<DIR>" == typeStr {
				node.IsDir = true
			} else {
				node.IsDir = false
			}
			nodeList[index] = node
		} else {
			//type=dir;modify=20200929015012; doc
			//type=file;modify=20201015025214;size=326785; topology-vue-master.zip
			splitFile := strings.Split(curFile, ";")
			typeStr := strings.Split(splitFile[0], "=")[1]
			isDir := false
			filename := ""
			node := models.Node{
				Expand:      false,
				Contextmenu: true,
				IsDir:       true,
				Children:    []models.Node{},
				DirPath:     dirPth,
			}
			if "dir" == typeStr {
				isDir = true
				filename = strings.TrimSpace(splitFile[2])
			} else {
				filename = strings.TrimSpace(splitFile[3])
			}
			node.Title = filename
			node.FileName = filename
			node.IsDir = isDir
			nodeList[index] = node
		}

	}
	return nodeList, nil
}
func (s *FtpFileSystem) DeleteFile(parentDir string, fileName string) error {
	formatDestPath := tools.FormatPath(parentDir)
	return s.fs.Dele(formatDestPath + tools.PathSeparator + fileName)
}
func (s *FtpFileSystem) CreateFile(parentDir string, fileName string) error {
	formatDestPath := tools.FormatPath(parentDir)
	return s.fs.Stor(formatDestPath+tools.PathSeparator+fileName, bytes.NewReader([]byte("")))
}
func (s *FtpFileSystem) SaveTextFile(parentDir string, fileName string, content string, policyType os.FileMode) error {
	return s.SaveByte(parentDir, fileName, []byte(content), policyType)
}
func (s *FtpFileSystem) SaveByte(parentDir string, fileName string, content []byte, policyType os.FileMode) error {
	formatDestPath := tools.FormatPath(parentDir)
	err := s.DeleteFile(parentDir, fileName)
	for err != nil && !strings.HasPrefix(err.Error(), "550") {
		err = s.DeleteFile(parentDir, fileName)
	}
	err1 := s.fs.Stor(formatDestPath+tools.PathSeparator+fileName, bytes.NewReader(content))
	return err1
}
func (s *FtpFileSystem) Rename(srcDir string, src string, dest string) error {
	formatDestPath := tools.FormatPath(srcDir)
	srcPath := formatDestPath + tools.PathSeparator + src
	destPath := formatDestPath + tools.PathSeparator + dest
	return s.fs.Rename(srcPath, destPath)
}
func (s *FtpFileSystem) Ping() error {
	_, err := s.fs.List("/")
	return err
}

/**
---------------------------------------
 文件系统工厂类负责读取配置参数生成文件系统实例
---------------------------------------
*/
type FtpFileSystemFactory struct {
}

func (s *FtpFileSystemFactory) getShareNames(dirPth string) string {
	formatDirPth := tools.FormatPath(dirPth)
	trimDirPath := tools.TrimLeft(formatDirPth)
	return strings.Split(trimDirPath, tools.PathSeparator)[0]
}
func (sam *FtpFileSystemFactory) Create(prefix string) (FileSystem, error) {
	sambahost := beego.AppConfig.String(prefix + "ftphost")
	sambaport := beego.AppConfig.String(prefix + "ftpport")
	sambauser := beego.AppConfig.String(prefix + "ftpuser")
	sambapassword := beego.AppConfig.String(prefix + "ftppassword")
	rootpath := beego.AppConfig.String(prefix + "rootpath")

	var err error
	var ftp *goftp.FTP

	// For debug messages: goftp.ConnectDbg("ftp.server.com:21")
	if ftp, err = goftp.Connect(sambahost + ":" + sambaport); err != nil {
		panic(err)
	}
	// Username / password authentication
	if err = ftp.Login(sambauser, sambapassword); err != nil {
		panic(err)
	}
	formatRootPath := tools.FormatPath(rootpath)
	fileSystem := FtpFileSystem{RootPath: formatRootPath, fs: ftp}
	return &fileSystem, nil
}
func (s *FtpFileSystemFactory) Name() string {
	return "service.FtpFileSystemFactory"
}
