package swagger

import (
	"bytes"
	_ "embed"
	"text/template"

	"admin_template/config"

	"github.com/gin-gonic/gin"
)

//go:embed swagger.html
var swaggerHtml string

// Setup 生成 swagger 格式的文档
func Setup(c *gin.Engine) {
	if config.Env == "local" || config.Env == "dev" {
		c.GET("/swagger", func(ctx *gin.Context) {
			ctx.Header("Content-Type", "text/html; charset=utf-8")
			tpl := template.Must(template.New("doc").Parse(swaggerHtml))
			buf := bytes.NewBuffer(nil)
			tpl.Execute(buf, map[string]string{
				"Title":   "后台管理系统",
				"SpecURL": "/swaggerfs/swagger.json",
			})
			ctx.Writer.Write(buf.Bytes())
		})

		c.Static("/swaggerfs", "asset/swagger")
	}
}
