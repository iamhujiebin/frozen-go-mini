package banner_r

import (
	"frozen-go-mini/common/domain"
	"frozen-go-mini/common/mycontext"
	"frozen-go-mini/domain/banner_m"
	"frozen-go-mini/resp"
	"github.com/gin-gonic/gin"
)

// @Tags 轮播图
// @Summary 获取所有轮播图
// @Param Authorization header string true "token"
// @Success 200 {object} []banner_m.Banner
// @Router /v1.0/banners [get]
func BannerList(c *gin.Context) (*mycontext.MyContext, error) {
	myCtx := mycontext.CreateMyContext(c.Keys)
	model := domain.CreateModelContext(myCtx)
	response := banner_m.GetAllBanners(model)
	resp.ResponseOk(c, response)
	return myCtx, nil
}
