package utils

import (
	"github.com/ncruces/zenity"
)

const defaultPath = ``

func GetFolder() string {
	path, _ := zenity.SelectFile(zenity.Filename(defaultPath), zenity.Directory())

	return path
}
