package threading

import (
	"go-admin-template/svc"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func GoSafeCtx(ctx *svc.ServiceContext, fn func()) {
	go RunSafeCtx(ctx, fn)
}

func RunSafeCtx(ctx *svc.ServiceContext, fn func()) {
	defer func() {
		if p := recover(); p != nil {
			ctx.Log.Errorf("%+v\n%+v", p, errors.New("go routine error"))
		}
	}()

	fn()
}

func GoSafe(fn func()) {
	go RunSafe(fn)
}

func RunSafe(fn func()) {
	defer func() {
		if p := recover(); p != nil {
			logrus.Errorf("%+v\n%+v", p, errors.New("go routine error"))
		}
	}()

	fn()
}
