package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	upload := route.Group("/api/upload")
	{
		upload.POST("", uploadFile)
	}
}

func uploadFile(c *gin.Context) {
	fmt.Println("File Upload Endpoint Hit")
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		c.String(http.StatusBadRequest, "Error parsing form: %v", err)
		return
	}
	file, handler, err := c.Request.FormFile("myFile")
	if err != nil {
		c.String(http.StatusBadRequest, "Error retrieving the file: %v", err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempDir := os.TempDir()
	tempImagesDir := filepath.Join(tempDir, "temp-images")
	err = os.MkdirAll(tempImagesDir, 0755)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating temp directory: %v", err)
		return
	}
	tempFile, err := os.CreateTemp(tempImagesDir, "upload-*.png")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating temporary file: %v", err)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error copying file contents: %v", err)
		return
	}

	c.String(http.StatusOK, "Successfully Uploaded File\n, tempImagesDir: %v", tempImagesDir)
}
