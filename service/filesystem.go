package service

import (
	"github.com/astaxie/beego"
	"gpm/models"
	"os"
	"reflect"
)

type PolicyType int32

const (
	APPEND  PolicyType = 0 //追加
	OVERIDE PolicyType = 1 //直接覆盖
)

type FileSystem interface {
	/**
	    读取文件字节内容
		:param 目录名称
	    :param 文件名称
	*/
	ReadByte(string, string) ([]byte, error)
	/**
	    读取文本内容
		:param 目录名称
	    :param 文件名称
	*/
	ReadText(string, string) (string, error)
	/**
	    创建目录
		:param 目录名称
	    :param 文件名称
	*/
	Mkdir(string, string) error
	/**
	    删除目录
		:param 目录名称
	    :param 文件名称
	*/
	RmDir(string, string) error
	/**
	    列表目录下所有子目录和文件
		:param 目录名称
	*/
	ListDir(string) ([]models.Node, error)
	/**
	  判断路径是否为目录
	  :param 文件路径
	*/
	IsDir(string) bool
	/**
	    列表根目录下所有子目录和文件
		:param 目录名称
	*/
	ListRoot() ([]models.Node, error)
	/**
	    删除文件
		:param 目录名称
	    :param 文件名称
	*/
	DeleteFile(string, string) error
	/**
	    创建文件
		:param 目录名称
	    :param 文件名称
	*/
	CreateFile(string, string) error
	/**
	    保存文件内容
		:param 目录名称
	    :param 文件名称
	    :param 文本内容
	    :param 覆盖还是追加
	*/
	SaveTextFile(string, string, string, os.FileMode) error
	/**
	    重命名
	    :param 文件所在目录
		:param 原始文件或者目录名称
	    :param 目标文件或者目录名称
	*/
	Rename(string, string, string) error
	/**
	       检查当前文件系统是否正确连接，如果异常 net.OpError
		**/
	Ping() error
}

type FileSystemFactoryI interface {
	/**
	  实现文件系统给予名称和app.conf中rootmodel值适配
	*/
	Name() string
	/**
	  创建对应的文件系统实例
	*/
	Create(prefix string) (FileSystem, error)
}

/**
---------------------------------------
 抽象工厂，负责通过文件系统工厂类型实例化生成工厂类
---------------------------------------
*/

type FileSystemAbstractFactory struct {
	factoryType map[string]FileSystemFactoryI
	ifInit      bool
}

func (s *FileSystemAbstractFactory) InitFactory() {
	if !s.ifInit {
		s.factoryType = make(map[string]FileSystemFactoryI)
		lfsf := LocalFileSystemFactory{}
		s.factoryType[lfsf.Name()] = &lfsf
		sfsf := SambaFileSystemFactory{}
		s.factoryType[sfsf.Name()] = &sfsf
		ffsf := FtpFileSystemFactory{}
		s.factoryType[ffsf.Name()] = &ffsf
		s.ifInit = true
	}
}
func (s *FileSystemAbstractFactory) AppendFactory(fi FileSystemFactoryI) {
	if !s.ifInit {
		s.factoryType[fi.Name()] = fi
	}
}
func (s *FileSystemAbstractFactory) ConstructFactory() (FileSystem, error) {
	systemfactoryClassStr := beego.AppConfig.String("systemfactory")
	systemfactoryInstance := s.factoryType[systemfactoryClassStr]
	return systemfactoryInstance.Create("")
}
func (s *FileSystemAbstractFactory) ConstructFactoryCustom(prefix string) (FileSystem, error) {
	var prefixStr string
	if prefix == "" {
		prefixStr = ""
	} else {
		prefixStr = prefix + "."
	}
	systemfactoryClassStr := beego.AppConfig.String(prefixStr + "systemfactory")
	systemfactoryInstance := s.factoryType[systemfactoryClassStr]
	systemfactoryInstanceType := reflect.TypeOf(systemfactoryInstance)
	ptrValue := reflect.New(systemfactoryInstanceType.Elem())
	systemfactoryInstance = ptrValue.Interface().(FileSystemFactoryI)
	return systemfactoryInstance.Create(prefixStr)
}
