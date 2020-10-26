package service

import (
	"github.com/astaxie/beego/context"
	"github.com/casbin/casbin"
)
var enforcer *casbin.Enforcer
func GetEnforcer()  *casbin.Enforcer{
	if(enforcer==nil) {
		enforcer= casbin.NewEnforcer("conf/casbin.conf", "conf/casbin_policy.csv")
	}
	return enforcer
}

var FilterUser = func(ctx *context.Context) {

}