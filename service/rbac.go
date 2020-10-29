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
var maxExpire int64 = 5

func GetCachedRbacContent() map[string]interface{} {
	if rbacDict == nil {
		rbacDict = make(map[string]interface{})
	}
	_, ok := rbacDict["contentKey"]
	if !ok {
		rbacDict["contentKey"] = getRbacContent()
		rbacDict["expire"] = time.Now().UnixNano() / 1000000000
	} else {
		expire := rbacDict["expire"].(int64)
		if (time.Now().UnixNano()/1000000000 - expire) > maxExpire {
			rbacDict["contentKey"] = getRbacContent()
			rbacDict["expire"] = time.Now().UnixNano() / 1000000000
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
	rbacContent := GetCachedRbacContent()["contentKey"].(map[string]interface{})
	users := rbacContent["users"].([]interface{})
	roles := rbacContent["roles"].(map[interface{}]interface{})
	userList := list.New()
	var curRole string
	for _, user := range users {
		if user.(map[interface{}]interface{})["userName"].(string) == userName {
			curRole = user.(map[interface{}]interface{})["role"].(string)
			break
		}
	}
	for roleName, roleInfo := range roles {
		roleArray := strings.Split(roleName.(string), "-")
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
func GetPathRequirePerm(curPath string) *list.List {
	permList := list.New()
	fileSystemPermMappingObj := GetCachedRbacContent()["contentKey"].(map[string]interface{})["fileSystemPermMapping"]
	fileSystemPermMapping := fileSystemPermMappingObj.([]interface{})
	for _, v := range fileSystemPermMapping {
		fmt.Println(v)
		pathActMap := v.(map[interface{}]interface{})
		path := pathActMap["path"].(string)
		if curPath == path {
			acts := pathActMap["act"].([]interface{})
			for _, a := range acts {
				permList.PushBack(a)
			}
			break
		}
	}
	return permList
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
    获取某个用户在某个路径下拥有的权限列表
	@userName 用戶名
	@inputPath 輸入的文件系統路徑
	@return ["read","write"]。
*/
func GetUserPathAccess(userName string, inputPath string) *list.List {
	roleArray := GetUserRoles(userName)
	actList := list.New()
	for i := roleArray.Front(); i != nil; i = i.Next() {
		for _, v := range i.Value.([]interface{}) {
			pathActV := v.(map[interface{}]interface{})
			path := pathActV["path"].(string)
			act := pathActV["act"].([]interface{})
			ifMatch, _ := regexp.Match(path, []byte(inputPath))
			if ifMatch {
				for _, v1 := range act {
					actList.PushBack(v1)
				}
			}
		}
	}
	return actList
}

/**
檢查當前用戶名對於某個路徑是否有操作的權限
@userName 用戶名
@inputPath 輸入的文件系統路徑
@inputAct 權限
@return true表示驗證通過，false表示驗證失敗。
*/
func CheckUserAct(userName string, inputPath string, inputAct string) bool {
	return CheckUserMulAct(userName, inputPath, []string{inputAct})
}
func CheckUserMulAct(userName string, inputPath string, inputAct []string) bool {
	//当前用户拥有的权限
	actsList := GetUserPathAccess(userName, inputPath)
	haveCount := 0
	//循环询问当前用户是否拥有权限
	for _, v := range inputAct {
		for a := actsList.Front(); a != nil; a = a.Next() {
			if v == a.Value || a.Value == "*" {
				haveCount++
				break
			}
		}
	}
	if haveCount == len(inputAct) {
		return true
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
func CheckUserActRead(userName string, inputPath string) bool {
	return CheckUserAct(userName, inputPath, ActRead)
}

/**
檢查當前用戶名對於某個路徑是否有写的權限
@userName 用戶名
@inputPath 輸入的文件系統路徑
@return true表示驗證通過，false表示驗證失敗。
*/
func CheckUserActWrite(userName string, inputPath string) bool {
	return CheckUserAct(userName, inputPath, ActWrite)
}
