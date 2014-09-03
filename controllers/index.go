package controllers

import "github.com/oal/beego-pongo2"

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	pongo2.Render(this.Ctx, "index.html", pongo2.Context{
		"APP":   this.Data,
		"Title": "抛弃世俗之浮躁，留我钻研之刻苦",
	})
}
