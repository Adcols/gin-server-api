package main

import (
	"path/filepath"
	"runtime"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// 获取项目根目录的绝对路径
func getProjectRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	// 当前文件在 scripts/gen/generate.go，需要回退两级到项目根目录
	return filepath.Join(filepath.Dir(filename), "../..")
}

// 生成代码的主程序
func main() {
	// 获取项目根目录
	projectRoot := getProjectRoot()

	// 连接数据库
	dsn := "root:Hbwz666..@tcp(127.0.0.1:3306)/gin_server_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	// 生成实例
	g := gen.NewGenerator(gen.Config{
		// 查询代码输出目录
		OutPath: filepath.Join(projectRoot, "internal/query"),
		// 实体结构体输出目录
		ModelPkgPath:      filepath.Join(projectRoot, "internal/model/entity"),
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	// 复用 GORM 的数据库连接
	g.UseDB(db)

	// 从 model 目录下的所有 model 生成查询代码
	g.ApplyBasic(
		// 这里指定需要生成查询代码的 model
		g.GenerateModel("user"),
	)

	// 生成代码
	g.Execute()
}
