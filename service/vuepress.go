package service

import (
	"gpm/database"
	"gpm/tools"
	"strings"
)

func GetVuePressMapping(vuepress database.VuePress) string {
	return GetVuePressMappingV1(vuepress.AppPath, vuepress.Workspace, vuepress.ShareUserName)
}
func GetVuePressMappingV1(appPath string, workspace int, shareUserName string) string {
	mapping := appPath
	if !strings.HasPrefix(appPath, tools.PathSeparator) {
		appPath = tools.PathSeparator + appPath
	}
	if workspace == 1 {
		mapping = tools.PathSeparator + shareUserName + appPath
	} else {
		mapping = tools.PathSeparator + "public" + appPath
	}
	return mapping
}
