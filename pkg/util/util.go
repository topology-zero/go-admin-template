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

// RFC3339toDateTime 将 RFC3339 格式转换为 DateTime 格式
func RFC3339toDateTime(s string) string {
	if len(s) == 0 {
		return s
	}

	t, err := time.ParseInLocation(time.RFC3339, s, time.Local)
	if err != nil {
		return s
	}
	return t.Format(time.DateTime)
}

// MaskData 数据脱敏
func MaskData[T any](src T) T {
	value := maskValue(reflect.ValueOf(src))
	return value.Interface().(T)
}

func maskValue(v reflect.Value) reflect.Value {
	if !v.IsValid() {
		return v
	}

	switch v.Kind() {
	case reflect.Pointer:
		if v.IsNil() {
			return v
		}
		p := reflect.New(v.Elem().Type())
		p.Elem().Set(maskValue(v.Elem()))
		return p
	case reflect.Struct:
		dst := reflect.New(v.Type()).Elem()

		for i := 0; i < v.NumField(); i++ {

			field := v.Type().Field(i)
			value := v.Field(i)

			// string 根据 tag 脱敏
			if value.Kind() == reflect.String {
				dst.Field(i).SetString(doMask(
					field.Tag.Get("mask"),
					value.String(),
				))
				continue
			}
			// 嵌套 struct / pointer
			switch value.Kind() {
			case reflect.Struct, reflect.Pointer:
				dst.Field(i).Set(maskValue(value))
			default:
				dst.Field(i).Set(value)
			}
		}
		return dst
	default:
		return v
	}
}

func doMask(rule, value string) string {
	switch rule {
	case "phone":
		if len(value) >= 7 {
			return value[:3] + "****" + value[len(value)-4:]
		}
	case "email":
		if i := strings.Index(value, "@"); i > 1 {
			return value[:1] + strings.Repeat("*", i-1) + value[i:]
		}
	case "name":
		r := []rune(value)
		if len(r) > 1 {
			return string(r[:1]) + strings.Repeat("*", len(r)-1)
		}
	case "idcard":
		if len(value) > 8 {
			return value[:4] + strings.Repeat("*", len(value)-8) + value[len(value)-4:]
		}
	case "password":
		return "******"
	}
	return value
}
