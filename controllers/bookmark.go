package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/oal/beego-pongo2"

	"go-blog/models/db"
)

type BookmarkController struct {
	AuthController
}

func (this *BookmarkController) Get() {
	pongo2.Render(this.Ctx, "bookmark.html", pongo2.Context{
		"APP":   this.Data,
		"Title": "书签",
	})
}

//添加和修改书签
func (this *BookmarkController) Post() {
	o := orm.NewOrm()
	bookmark := &db.Bookmarks{}

	if id, err := this.GetInt("id"); err == nil {
		if id != 0 {
			bookmark.Id = int(id)
			o.Read(bookmark)
		}

		bookmark.Special = this.GetString("special") != ""
		bookmark.Name = this.GetString("name")

		var err error
		if id != 0 {
			_, err = o.Update(bookmark)
		} else {
			_, err = o.Insert(bookmark)
		}

		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]string{"err": "更新数据库错误"}
		} else {
			this.Data["json"] = map[string]string{"ok": "添加成功"}
		}
	} else {
		this.Data["json"] = map[string]string{"err": "未知错误"}
	}

	this.ServeJson()
}
