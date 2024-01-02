package cat_m

import "frozen-go-mini/common/domain"

type Category struct {
	Label    string `json:"label"`    // 标签
	Title    string `json:"title"`    // 标题
	Icon     string `json:"icon"`     // icon
	Dot      bool   `json:"dot"`      // 是否红点
	DotCount int    `json:"dotCount"` // 红点数量
}

// 获取所有分类
func GetAllCategories(model *domain.Model) []Category {
	var res []Category
	res = append(res, Category{
		Label: "时令上新",
		Title: "时令上新",
		Icon:  "bad-laugh",
	})
	return res
}
