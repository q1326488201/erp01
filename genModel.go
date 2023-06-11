package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/erp?charset=utf8mb4&parseTime=True&loc=Local"))
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dao/query",
		ModelPkgPath: "./dao/models",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	g.UseDB(db) // reuse your gorm db
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
