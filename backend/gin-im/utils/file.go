package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

type File struct {
}

// Upload 上传文件
func (f *File) Upload(ctx *gin.Context) (filePath string, err error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return "", err
	}

	ext := path.Ext(file.Filename)
	if !f.checkExt(ext) {
		return "", errors.New("文件类型不符合要求")
	}

	filePath = fmt.Sprintf("./upload/%s", file.Filename)
	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		return "", err
	}
	filePath = strings.TrimLeft(filePath, ".")
	return
}

// checkExt 校验文件类型
func (f *File) checkExt(ext string) bool {
	fileExtList := [...]string{
		".png", ".jpeg", ".gif", ".jpg",
	}
	for _, s := range fileExtList {
		if ext == s {
			return true
		}
	}
	return false
}
