package main

import (
	"gorm.io/gen"
	"online-judge/biz/dal/mysql"
	"online-judge/biz/model"
)

func init() {
	mysql.InitDB()
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./biz/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(mysql.DBHook) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.User{})

	// Generate the code
	g.Execute()
}
