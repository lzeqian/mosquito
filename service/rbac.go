package service

import (
	"container/list"
	"fmt"
	"gopkg.in/yaml.v2"
	"gpm/tools"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)
import "github.com/chenhg5/collection"

/**
  读取rbac配置文件"conf/rbac.yaml"
*/
func getRbacContent() map[string]interface{} {
	yamlFile, _ := ioutil.ReadFile("conf/rbac.yml")
	resultMap := make(map[string]interface{})
	err := yaml.Unmarshal(yamlFile, resultMap)
	if err != nil {
		fmt.Println(err)
	}
	//resultMap, err := yaml.ReadYmlReader("conf/rbac.yml")
	//if err != nil {
	//	fmt.Println(err)
	//}
	return resultMap
}

/**
  rabcDict字典用于保存当前内容和读取时间。
*/
var rbacDict map[string]interface{}

/**
  当前文本缓存过期时间，单位s。
*/
var maxExpire int64 = 60

func GetCachedRbacContent() map[string]interface{} {
	if rbacDict == nil {
		rbacDict = make(map[string]interface{})
	}
	_, ok := rbacDict["contentKey"]
	if !ok {
		rbacDict["contentKey"] = getRbacContent()
		rbacDict["expire"] = time.Now().UnixNano() / 1000
	} else {
		expire := rbacDict["expire"].(int64)
		if (time.Now().UnixNano()/1000 - expire) > maxExpire {
			rbacDict["contentKey"] = getRbacContent()
			rbacDict["expire"] = time.Now().UnixNano() / 1000
		}
	}
	return rbacDict
}

/**
  通过当前用户名，获取当前用户详情信息。
  @userName 用户名
  @return 用户详情
*/
func GetUser(userName string) map[interface{}]interface{} {
	rbacContent := GetCachedRbacContent()
	contentKey := rbacContent["contentKey"].(map[string]interface{})
	users := contentKey["users"]
	userArray := users.([]interface{})
	for _, user := range userArray {
		userInfo := user.(map[interface{}]interface{})
		if userInfo["userName"] == userName {
			return userInfo
		}
	}
	return nil
}

/**
 通过用户名获取用户对应角色的权限信息
@userName 用户名
@return
 list.List([{
  path: '/.+'
  act:
    - 'read'}])
*/
func GetUserRoles(userName string) *list.List {
	rbacContent := GetCachedRbacContent()
	users := rbacContent["users"].([]map[string]interface{})
	roles := rbacContent["roles"].(map[string]interface{})
	userList := list.New()
	var curRole string
	for _, user := range users {
		if user["userName"] == userName {
			curRole = user["role"].(string)
			break
		}
	}
	for roleName, roleInfo := range roles {
		roleArray := strings.Split(roleName, "-")
		for _, role := range roleArray {
			if role == "*" {
				userList.PushBack(roleInfo)
			} else {
				if curRole == role {
					userList.PushBack(roleInfo)
				}
			}
		}
	}
	return userList
}

/**
檢查用戶名和密碼是否匹配
@userName 用戶名
@password 密碼
@return true表示驗證通過，false表示驗證失敗。
*/
func CheckUserPassword(userName string, password string) bool {
	userInfo := GetUser(userName)
	userPassword := userInfo["password"].(string)
	md5Password := tools.Md5(password)
	ifActivate := userInfo["ifActivate"].(int)
	if userPassword == md5Password && ifActivate == 1 {
		return true
	}
	return false
}

/**
檢查當前用戶名對於某個路徑是否有操作的權限
@userName 用戶名
@inputPath 輸入的文件系統路徑
@inputAct 權限
@return true表示驗證通過，false表示驗證失敗。
*/
func CheckUserAccess(userName string, inputPath string, inputAct string) bool {
	roleArray := GetUserRoles(userName)
	for i := roleArray.Front(); i != nil; i = i.Next() {
		for _, v := range i.Value.([]map[string]interface{}) {
			path := v["path"].(string)
			act := v["act"].([]string)
			ifMatch, _ := regexp.Match(path, []byte(inputPath))
			if ifMatch && collection.Collect(act).Contains(inputAct) {
				return true
			}
		}
	}
	return false
}

const (
	ActRead       string = "read"
	ActWrite      string = "write"
	ActCreateDir  string = "createDir"
	ActListDir    string = "listDir"
	ActCreateFile string = "createFile"
	ActDeleteDir  string = "deleteDir"
	ActDeleteFile string = "deleteFile"
)

/**
檢查當前用戶名對於某個路徑是否有读的權限
@userName 用戶名
@inputPath 輸入的文件系統路徑
@return true表示驗證通過，false表示驗證失敗。
*/
func CheckUserRead(userName string, inputPath string) bool {
	return CheckUserAccess(userName, inputPath, ActRead)
}

/**
檢查當前用戶名對於某個路徑是否有写的權限
@userName 用戶名
@inputPath 輸入的文件系統路徑
@return true表示驗證通過，false表示驗證失敗。
*/
func CheckUserWrite(userName string, inputPath string) bool {
	return CheckUserAccess(userName, inputPath, ActWrite)
}
