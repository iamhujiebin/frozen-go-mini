package route

import (
	_ "frozen-go-mini/docs"
	"frozen-go-mini/route/banner_r"
	"frozen-go-mini/route/cat_r"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	var r = gin.Default()
	r.Use(Cors()) // 跨域
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ws
	//r.GET("/ws", JWTApiHandle, ws_r.WsHandler)
	// http
	noLogin := r.Group("")
	noLogin.Use(ExceptionHandle, LoggerHandle)
	v1 := noLogin.Group("/v1.0")
	{
		v1.GET("/banners", wrapper(banner_r.BannerList))
		v1.GET("/category", wrapper(cat_r.CategoryList))
	}
	return r
}
