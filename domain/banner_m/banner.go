package banner_m

import "frozen-go-mini/common/domain"

type Banner struct {
	PicUrl string
}

func GetAllBanners(model *domain.Model) []Banner {
	var res []Banner
	res = append(res, Banner{"https://prod-mall-cos-1252929494.cos.ap-guangzhou.myqcloud.com/f08c02795cea481e9d149a2bb93353dc.png"})
	res = append(res, Banner{"https://prod-mall-cos-1252929494.cos.ap-guangzhou.myqcloud.com/1c67f9aab45d4db6bbc4842f4ed368d3.png"})
	res = append(res, Banner{"https://prod-mall-cos-1252929494.cos.ap-guangzhou.myqcloud.com/a77139578ada4a0e96f777d60370e524.jpg"})
	res = append(res, Banner{"https://prod-mall-cos-1252929494.cos.ap-guangzhou.myqcloud.com/883314bc5a22440f964542417c5f28db.jpg"})
	res = append(res, Banner{"https://prod-mall-cos-1252929494.cos.ap-guangzhou.myqcloud.com/952bc6aa8d564d4a9aacaceba25dc62c.png"})
	res = append(res, Banner{"https://prod-mall-cos-1252929494.cos.ap-guangzhou.myqcloud.com/58f841cb3a4545368e96bb0613462f03.jpg"})
	res = append(res, Banner{"https://prod-mall-cos-1252929494.cos.ap-guangzhou.myqcloud.com/8ca4c50e8ed8410183c36f9c40720cb8.jpg"})
	res = append(res, Banner{"https://prod-mall-cos-1252929494.cos.ap-guangzhou.myqcloud.com/e06330ff00164227be5444f6cdf80eac.jpg"})
	return res
}
