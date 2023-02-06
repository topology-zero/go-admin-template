package login

import (
	"fmt"
	"testing"

	"admin_template/config"
	"admin_template/pkg/jwt"
)

func TestLogin(t *testing.T) {
	config.Setup("etc/admin-local.yaml")
	token, _ := jwt.MakeToken(1, 1, "10086")
	fmt.Println(token)
}
