package controller

import (
	"media-matrix/http/response"
	"media-matrix/logic/model"

	"github.com/kataras/iris"
)

var Index = new(index)

type index struct {
}

func (l *index) Banner(ctx iris.Context) {

	var res []model.Banner = make([]model.Banner, 0)
	model.Banner{}.Model().
		Order("sort desc").
		Scan(&res)
	ctx.JSON(response.Success(res))

}

func (l *index) Cate(ctx iris.Context) {
	var res []model.Cate = make([]model.Cate, 0)
	model.Cate{}.Model().
		Where("pid=0").
		Order("sort desc").
		Limit(8).
		Scan(&res)
	ctx.JSON(response.Success(res))
}

func (l *index) Recommend(ctx iris.Context) {
	var res []model.Cate = make([]model.Cate, 0)
	model.Cate{}.Model().
		Where("recommend=1").
		Order("sort desc").
		Limit(4).
		Scan(&res)
	ctx.JSON(response.Success(res))
}

func (l *index) Article(ctx iris.Context) {
	var res []model.Article = make([]model.Article, 0)
	model.Article{}.Model().
		Where("recommend=1").
		Order("sort desc").
		Select("id,name,img_url").
		Limit(4).
		Scan(&res)
	ctx.JSON(response.Success(res))
}
