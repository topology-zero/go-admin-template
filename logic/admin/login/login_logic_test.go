package login

import (
	"fmt"
	"testing"

	"go-admin-template/config"
	"go-admin-template/pkg/jwt"
)

func TestLogin(t *testing.T) {
	config.Setup("etc/admin-local.yaml")
	token, _ := jwt.MakeToken(1, 1, "10086")
	fmt.Println(token)
}
