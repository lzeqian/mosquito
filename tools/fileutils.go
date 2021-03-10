package tools

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var PathSeparator string = "/"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ExecCommand(commandName string, params []string) bool {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	err1 := cmd.Start()
	if err1 != nil {
		io.Copy(cmd.Stderr, bytes.NewBufferString(err1.Error()))
		//fmt.Println(err)
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return true
}
func ExportToFormat(srcFileBytes []byte, fileName string, formatType string, tmpDir string) []byte {
	destPath := tmpDir + PathSeparator + fileName
	//拷贝文件内容到服务器中。
	ioutil.WriteFile(destPath, srcFileBytes, 0777)
	defer os.Remove(destPath)
	defer os.Remove(fileName)
	fileBase := strings.Split(filepath.Base(fileName), ".")[0]
	genDestFileName := tmpDir + PathSeparator + fileBase + "." + formatType
	osType := runtime.GOOS
	cmd := ""
	param1 := ""
	libreofficePath := beego.AppConfig.String("libreofficepath")
	if osType == "windows" {
		cmd = "cmd"
		param1 = "/C"
	} else {
		cmd = "/bin/bash"
		param1 = "-c"
	}
	commandParam := "--headless" + " --convert-to" + " " + formatType + " --outdir " + tmpDir + " " + destPath
	//执行后是异步的，无法感知最终生成的文件，只能定时去扫描该文件是否存在
	if ExecCommand(cmd, []string{param1, libreofficePath, commandParam}) {
		//尝试15s后未处理完成直接失败
		i := 1
		//最大尝试次数是30次
		maxCount := 30
		//每次休眠时间为 500ms
		unit := 500
		for {
			if i > maxCount {
				break
			}
			_, err := os.Stat(genDestFileName)
			if os.IsNotExist(err) {
				time.Sleep(time.Duration(unit) * time.Millisecond)
				i = i + 1
				continue
			}
			defer os.Remove(genDestFileName)
			readBytes, _ := ioutil.ReadFile(genDestFileName)
			if len(readBytes) <= 0 {
				time.Sleep(time.Duration(unit) * time.Millisecond)
				i = i + 1
				continue
			}
			return readBytes
		}
	}
	return nil
}
