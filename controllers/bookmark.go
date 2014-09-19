package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/oal/beego-pongo2"

	"go-blog/models/db"
)

type BookmarkController struct {
	AuthController
}

func (this *BookmarkController) Get() {
	o := orm.NewOrm()
	qs := o.QueryTable(&db.Bookmarks{})
	var bookmarks []db.Bookmarks
	if _, err := qs.OrderBy("-Special", "-Count").All(&bookmarks); err != nil {
		beego.Error(err)
		return
	}

	pongo2.Render(this.Ctx, "bookmark.html", pongo2.Context{
		"Title":     "书签",
		"APP":       this.Data,
		"Bookmarks": bookmarks,
	})
}

//删除书签
func (this *BookmarkController) DelBookmark() {
	if id, err := this.GetInt("id"); err == nil {
		o := orm.NewOrm()
		bookmark := &db.Bookmarks{Id: int(id)}

		if _, err := o.Delete(bookmark); err == nil {
			this.Data["json"] = map[string]string{"ok": "删除成功"}
		}
	} else {
		this.Abort("400")
	}

	this.ServeJson()
}

//专辑切换
func (this *BookmarkController) Switch() {
	if id, err := this.GetInt("id"); err == nil { //读取参数
		o := orm.NewOrm()
		bookmark := &db.Bookmarks{Id: int(id)}

		if o.Read(bookmark) != nil {
			this.Data["json"] = map[string]string{"err": "读取数据库失败"}
		} else {
			b, _ := this.GetBool("special") //读取参数
			bookmark.Special = b

			if _, err := o.Update(bookmark); err == nil {
				this.Data["json"] = map[string]string{"ok": "更新成功"}
			}

		}
	} else {
		this.Abort("400")
	}

	this.ServeJson()
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

		bookmark.Name = strings.TrimSpace(this.GetString("name"))

		var err error
		if id != 0 {
			_, err = o.Update(bookmark)
		} else {
			_, err = o.Insert(bookmark)
		}

		if err != nil {
			this.Data["json"] = map[string]string{"err": err.Error()}
		} else {
			this.Data["json"] = map[string]string{"ok": "更新成功"}
		}
	} else {
		this.Data["json"] = map[string]string{"err": "未知错误"}
	}

	this.ServeJson()
}
