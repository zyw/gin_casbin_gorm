/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"gin_casbin_gorm/app/gcg"
	"gin_casbin_gorm/initialize"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动gin集成casbin和gorm",
	Long:  `一个demo项目用于测试gin继承casbin和gorm.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// 数据库连接以及角色规则的初始化
func connect() {

	enforcer := initialize.Enforcer()
	// 创建一个角色，并赋予权限
	// admin 这个角色可以访问GET方式访问 /api/v2/ping
	res, err := enforcer.AddPolicy("admin", "/api/v2/ping", "GET")
	if err != nil {
		fmt.Println("AddPolicy error")
		panic(err)
	}
	if !res {
		fmt.Println("policy is exist")
	} else {
		fmt.Println("policy is not exist, adding")
	}
	// 将 test 用户加入一个角色中
	addRes, err := enforcer.AddRoleForUser("test", "root")
	if err != nil {
		fmt.Println("AddRoleForUser(\"test\", \"root\") error")
		panic(err)
	}
	if !addRes {
		fmt.Println("test root is exist")
	}
	addRes, err = enforcer.AddRoleForUser("tom", "admin")
	if err != nil {
		fmt.Println("AddRoleForUser(\"tom\", \"admin\") error")
		panic(err)
	}
	if !addRes {
		fmt.Println("tom admin is exist")
	}
}

func run() {
	connect()
	gcg.InitRouter()
}
