package user

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestAdd(t *testing.T) {

	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	fmt.Println(string(password))
}
