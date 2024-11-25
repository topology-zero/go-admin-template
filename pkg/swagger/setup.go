package swagger

import (
	"go-admin-template/config"

	"github.com/gin-gonic/gin"
)

// Setup 生成 swagger 格式的文档
func Setup(c *gin.Engine) {
	if config.Env == "local" || config.Env == "dev" {
		c.Static("/swagger", "asset/swagger")
	}
}
