package common

import (
	"path/filepath"

	"go-admin-template/svc"

	"github.com/pkg/errors"
	"github.com/samber/lo"
)

// UploadImage 上传图片
func UploadImage(ctx *svc.ServiceContext) (resp string, err error) {
	file, err := ctx.GinContext.FormFile("file")
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		return
	}

	ext := filepath.Ext(file.Filename)
	if !lo.Contains([]string{".png", ".jpg", ".jpeg"}, ext) {
		err = errors.New("不支持的文件格式")
		return
	}

	open, err := file.Open()
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("系统错误")
		return
	}
	defer open.Close()

	// TODO 上传图片至 OSS

	//resp = fmt.Sprintf("https://foo.bar.com/%s", filePath)
	resp = "http://tc.masterjoy.top/typory/image-20240828121602348.png"
	return
}
