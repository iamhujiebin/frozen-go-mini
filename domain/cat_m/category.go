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
	res = append(res, Category{Label: "时令上新", Title: "时令上新", Icon: "bad-laugh"})
	res = append(res, Category{Label: "热饮推荐", Title: "热饮推荐", Icon: "sinister-smile", Dot: true})
	res = append(res, Category{Label: "轻负担推荐", Title: "轻负担推荐", Icon: "smile"})
	res = append(res, Category{Label: "清乳茶", Title: "清乳茶", Icon: "joyful", DotCount: 8})
	res = append(res, Category{Label: "清爽真果茶", Title: "清爽真果茶", Icon: "joyful"})
	res = append(res, Category{Label: "传统真果茶", Title: "传统真果茶", Icon: "awkward"})
	res = append(res, Category{Label: "茗茶/咖啡/蛋糕", Title: "茗茶/咖啡/蛋糕", Icon: "sleep"})
	res = append(res, Category{Label: "吃点啥", Title: "吃点啥", Icon: "uncomfortable"})
	res = append(res, Category{Label: "小料&提示", Title: "小料&提示", Icon: "joyful"})
	res = append(res, Category{Label: "经典果茶", Title: "经典果茶", Icon: "crack"})
	return res
}
