package BasicUsages

import (
	"context"
	"os"
	"path/filepath"

	MinioClient "go-learning/src/Utils/Minio"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/minio/minio-go/v7"
)

func UploadFile(c *fiber.Ctx) error {
	_, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	file, err := c.FormFile("file")

	buffer, err := file.Open()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	defer buffer.Close()

	// objectName := file.Filename // for filename
	// fileBuffer := buffer // ??
	mime := file.Header["Content-Type"][0]
	if mime != "image/jpeg" && mime != "image/png" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "file must be image",
		})
	}

	fileSize := file.Size / (1024 * 1024)
	log.Debug(fileSize)

	// getwd for get current directory
	currentDir, _ := os.Getwd()
	log.Debug(currentDir)

	// filepath join for join path
	destination := filepath.Join(currentDir, "./uploads")
	log.Debug(destination)
	_, err = os.Stat(destination) // Declare fileInfo and err to receive the values
	if os.IsNotExist(err) {
		// Create directory with 0755 permission
		os.MkdirAll(destination, 0755)
	} else {
		log.Debug("Directory already exists")
	}

	// save file
	err = c.SaveFile(file, destination+"/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":    true,
			"messages": "success-upload-file",
		},
	})
}

func UploadFileAws(c *fiber.Ctx) error {
	ctx := context.Background()
	minioClient, err := MinioClient.ConnectMinio()
	if err != nil {
		return err
	}

	// Make a new bucket called mymusic.
	bucketName := "eddc1edcfb2"
	location := "us-east-1"
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	if !exists {
		err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			return err
		}
	}

	currentDir, _ := os.Getwd()
	objectName := "testing.jpg"
	filePath := filepath.Join(currentDir, "./uploads/"+objectName)
	// Upload the zip file
	contentType := "image/jpg"

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":    true,
			"messages": "success-upload-file",
			"data":     info,
		},
	})
}
