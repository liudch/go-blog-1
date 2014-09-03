package models

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/astaxie/beego"
)

func Md5(origStr string) string {
	h := md5.New()
	h.Write([]byte("Bing.L"))                            //固定盐
	h.Write([]byte(origStr))                             //原始数据
	h.Write([]byte(beego.AppConfig.String("SecretKey"))) //用户配置的安全KEY

	return hex.EncodeToString(h.Sum(nil))
}
