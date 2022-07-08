package gorm

import (
	"errors"
	"fmt"
	"gin_casbin_gorm/initialize"
	"gin_casbin_gorm/models"

	"gorm.io/gorm"
)

func GormDemo() {
	db := initialize.MyDB()

	// insert 插入
	/*
		dt := entity.DictType{
			Id:         123457,
			Name:       "类型",
			Status:     1,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}

		if err := db.Create(&dt).Error; err != nil {
			fmt.Println("插入数据错误！" + err.Error())
			panic("插入数据错误：" + err.Error())
		}
	*/

	// select查询
	// 定义一个保存数据的结构体
	dt := models.DictType{}

	// 查询第一条数据
	// SELECT * FROM `sys_dict_type` WHERE (name='类型') LIMIT 1
	result := db.Where("name = ?", "类型").First(&dt)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("找不到记录")
		panic("找不到数据，" + result.Error.Error())
	}

	fmt.Println(dt)

	fmt.Println("成功")
}
