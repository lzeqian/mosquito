package tools

import (
	"regexp"
	"strings"
)

func  TrimLeft(dirPth string) (string) {
	str:=strings.TrimLeft(dirPth,PathSeparator)
	return strings.TrimLeft(str,"\\");
}
func  FormatPath(dirPth string) (string) {
	re3, _ := regexp.Compile("[\\|/]+]");
	return re3.ReplaceAllString(dirPth,PathSeparator)
}
/**
  获取第一层目录名称。
  如：/ab/a/b.txt 结果:ab
 */
func  GetRootName(dirPth string) (string) {
	formatDirPth:=FormatPath(dirPth)
	trimDirPath:=TrimLeft(formatDirPth)
	return strings.Split(trimDirPath,PathSeparator)[0];
}