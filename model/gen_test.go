package model

import (
	"flag"
	"fmt"
	"testing"

	"admin_template/config"
	"admin_template/model/internal"
	"admin_template/pkg/logger"

	"gorm.io/gen"
)

// 运行该测试用例时,需要设置工作目录为项目目录
func TestGEN(t *testing.T) {
	flag.Parse()

	configFile := fmt.Sprintf("etc/admin-%s.yaml", config.Env)
	config.Setup(configFile)
	logger.Setup()
	Setup()

	// 生成实例
	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath: "./query",

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: false, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type

		// 生成 gorm 标签的字段添加索引 tag
		FieldWithIndexTag: false, // generate with gorm index tag

		// 生成 gorm 标签的字段类型 tag
		FieldWithTypeTag: false, // generate with gorm column type tag
	})

	// 不需要 json tag
	g.WithJSONTagNameStrategy(func(columnName string) (tagContent string) {
		return ""
	})

	// model 命名
	g.WithModelNameStrategy(func(tableName string) (modelName string) {
		return db.Config.NamingStrategy.SchemaName(tableName) + "Model"
	})

	// 不需要将数据库中的 int 转换成 struct 的 int32
	g.WithDataTypeMap(map[string]func(detailType string) (dataType string){
		"int": func(detailType string) (dataType string) { return "int" },
	})

	// 默认 deleted_at 字段是 gorm.Delete 类型, 如果其他字段需要实现软删除,则需要将这个设置
	g.WithOpts(gen.FieldType("delete_time", "gorm.DeletedAt"))

	// 设置目标 db
	g.UseDB(db)

	// 生成全部表的 model
	g.GenerateAllTable()

	// 创建模型,不创建 query 文件
	//g.GenerateModel("admin_user")

	// 创建模型 + query ,不加入自定义方法
	g.ApplyBasic(g.GenerateModel("auth"))
	g.ApplyBasic(g.GenerateModel("casbin_rule"))

	// 创建模型的方法 + 创建 query 文件 + 自定义方法
	g.ApplyInterface(func(internal.AdminUser) {}, g.GenerateModel("admin_user"))
	g.ApplyInterface(func(internal.Role) {}, g.GenerateModel("role"))

	// 最后执行文件生成
	g.Execute()
}
