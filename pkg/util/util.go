package util

import (
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
)

// CopyValue 数据复制
func CopyValue(toValue, fromValue interface{}) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					if src == nil {
						return "", nil
					}
					s, ok := src.(time.Time)
					if !ok {
						return nil, errors.New("CopyValue 时间转换错误")
					}
					return s.Format(time.DateTime), nil
				},
			},
		},
	})
}

// Title 蛇形转驼峰
func Title(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// Camel 蛇形转驼峰,首字母小写
func Camel(s string) string {
	if len(s) == 0 {
		return s
	}
	s = Title(s)
	return strings.ToLower(s[:1]) + s[1:]
}

type empty struct{}

var (
	pkgNameOnce sync.Once
	pkgName     string
)

func GetPkgName() string {
	pkgNameOnce.Do(func() {
		pkgNames := reflect.TypeOf(empty{}).PkgPath()
		split := strings.Split(pkgNames, "/")
		pkgName = split[0]
	})
	return pkgName
}

const charset = "0123456789"

func RandomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
