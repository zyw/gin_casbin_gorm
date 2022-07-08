package initialize

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var enforcer *casbin.Enforcer

func init() {
	fmt.Println("开始初始化Enforcer.........")
	sqlDB := MyDB()

	// 将数据库连接同步给插件，插件用来操作数据库
	po, err := gormadapter.NewAdapterByDB(sqlDB)
	if err != nil {
		fmt.Println("NewAdapterByDB error")
		panic(err)
	}
	// 创建Enforcer，这里也可以使用原生字符串方式
	modelStr :=
		`
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
	`
	m, err := model.NewModelFromString(modelStr)
	if err != nil {
		fmt.Println("NewModelFromString error")
		panic(err)
	}
	enforcer, err = casbin.NewEnforcer(m, po)
	if err != nil {
		fmt.Println("NewEnforcer error")
		panic(err)
	}

	// 开启权限认证日志
	enforcer.EnableLog(true)
	// 加载数据库中的策略
	err = enforcer.LoadPolicy()
	if err != nil {
		fmt.Println("loadPolicy error")
		panic(err)
	}
}

func Enforcer() *casbin.Enforcer {
	return enforcer
}
