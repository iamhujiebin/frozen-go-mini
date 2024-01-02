package cat_r

import (
	"frozen-go-mini/common/domain"
	"frozen-go-mini/common/mycontext"
	"frozen-go-mini/domain/cat_m"
	"frozen-go-mini/resp"
	"github.com/gin-gonic/gin"
)

// @Tags 分类
// @Summary 获取所有分类
// @Param Authorization header string true "token"
// @Success 200 {object} []cat_m.Category
// @Router /v1.0/category [get]
func CategoryList(c *gin.Context) (*mycontext.MyContext, error) {
	myCtx := mycontext.CreateMyContext(c.Keys)
	model := domain.CreateModelContext(myCtx)
	response := cat_m.GetAllCategories(model)
	resp.ResponseOk(c, response)
	return myCtx, nil
}
