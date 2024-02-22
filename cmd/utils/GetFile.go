package utils

import (
	"github.com/ncruces/zenity"
)

const defaultPathFile = ``

func GetFile() string {
	path, _ := zenity.SelectFile(zenity.Filename(defaultPath))

	return path
}
