package controllers

import "github.com/oal/beego-pongo2"

type BookmarkController struct {
	BaseController
}

func (this *BookmarkController) Get() {
	//email := this.GetSession("Email")
	//if email == nil {
	//	this.Redirect(beego.UrlFor("IndexController.GET"), 302)
	//}

	pongo2.Render(this.Ctx, "bookmark.html", pongo2.Context{
		"APP":   this.Data,
		"Title": "书签",
	})
}
