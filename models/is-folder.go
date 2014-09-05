package models

import "os"

// 判断目录是否存在，并且可以不存在可以尝试创建
func IsFolder(pth string, tryMk bool) bool {
	fi, err := os.Stat(pth)

	if err != nil {
		if tryMk {
			return os.MkdirAll(pth, os.ModePerm) == nil
		}

		return false
	}

	return fi.IsDir()
}
