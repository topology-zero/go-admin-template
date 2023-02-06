package util

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
					return s.Format("2006-01-02 15:04:05"), nil
				},
			},
		},
	})
}

// WarpDbError 将错误打印，且不包含敏感信息返回
func WarpDbError(err error) error {
	if err == nil {
		return nil
	}
	logrus.Errorf("数据库异常：%+v", errors.WithStack(err))
	return errors.New("系统错误")
}
