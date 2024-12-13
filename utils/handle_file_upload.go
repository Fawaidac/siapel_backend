package utils

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleFileUpload(c *gin.Context, reqName, uploadDir string) (string, error) {
	file, err := c.FormFile(reqName)
	if err != nil {
		return "", err
	}
	sanitizedFilename := filepath.Base(file.Filename)
	uploadPath := filepath.Join(uploadDir, sanitizedFilename)
	if err := os.MkdirAll(filepath.Dir(uploadPath), os.ModePerm); err != nil {
		return "", err
	}
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		return "", err
	}
	return uploadPath, nil
}