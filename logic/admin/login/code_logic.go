package login

import (
	"fmt"
	"sync"
	"time"

	"go-admin-template/config"
	"go-admin-template/pkg/redis"
	"go-admin-template/svc"
	"go-admin-template/types"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/wenlng/go-captcha-assets/resources/imagesv2"
	"github.com/wenlng/go-captcha-assets/resources/tiles"
	"github.com/wenlng/go-captcha/v2/slide"
)

var (
	slideOnce sync.Once
	slideCapt slide.Captcha
)

// Code 获取验证码
func Code(ctx *svc.ServiceContext) (resp types.CodeResponse, err error) {
	slideOnce.Do(func() {
		imgs, _ := imagesv2.GetImages()
		graphs, _ := tiles.GetTiles()

		var newGraphs = make([]*slide.GraphImage, 0, len(graphs))
		for i := 0; i < len(graphs); i++ {
			graph := graphs[i]
			newGraphs = append(newGraphs, &slide.GraphImage{
				OverlayImage: graph.OverlayImage,
				MaskImage:    graph.MaskImage,
				ShadowImage:  graph.ShadowImage,
			})
		}

		builder := slide.NewBuilder()
		builder.SetResources(
			slide.WithGraphImages(newGraphs),
			slide.WithBackgrounds(imgs),
		)
		slideCapt = builder.Make()
	})

	captData, err := slideCapt.Generate()
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		return
	}

	data := captData.GetData()
	key := uuid.NewString()
	redis.Client.Set(ctx, config.ServerConf.Name+":CAPTCHAKEY:"+key, fmt.Sprintf("%d,%d", data.X, data.Y), 5*time.Minute)

	resp.CaptchaKey = key
	resp.ImageBase64, _ = captData.GetMasterImage().ToBase64()
	resp.TileBase64, _ = captData.GetTileImage().ToBase64()
	resp.TileWidth = data.Width
	resp.TileHeight = data.Height
	resp.TileX = data.DX
	resp.TileY = data.DY
	return
}
