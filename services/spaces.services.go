package services

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

/*
Handles all interaction with digital ocean spaces using aws sdk v3
*/

type Spaces struct {
	s3Client *s3.S3
}

func NewSpacesService(s3Client *s3.S3) Spaces {
	return Spaces{s3Client: s3Client}
}

func getImageName(fullPath string) string {
	pathSlice := strings.Split(fullPath, "/")

	return pathSlice[len(pathSlice) - 1]
}


func (sp Spaces) getPicturePath(fullPath string) string {
	pathSlice := strings.Split(fullPath, "/")
	pathOnly := ""

	for i:=0; i < len(pathSlice)-1; i++ {
		pathOnly += fmt.Sprintf("%s/", pathSlice[i])
	}

	return pathOnly
}

func (sp Spaces) GetObject(filePath string) (string, error) {

	bucket := os.Getenv("DO_BUCKET")
	folder := os.Getenv("DO_FOLDER")

	//fileName := "advert_images/oSfYTd1UYTglibEbUV3xTJuJyFlWNX1Xvt7Tp0rl.jpg"
	key := filePath
	key = fmt.Sprintf("%s/%s", folder, filePath)

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := sp.s3Client.GetObject(input)
	if err != nil {
		return "", err
	}

	imageName := getImageName(filePath)
	saveLocation, err := os.Create(filepath.Join(imagePath, imageName))
	if err != nil {
		return "", err
	}
	defer saveLocation.Close()

	_, err = io.Copy(saveLocation, result.Body)
	if err != nil {
		return "", err
	}
	return imageName, nil
}

func (sp Spaces) UploadObject(fileName string, originalImagePath string) (string,error) {
	imgFile, err := os.Open(filepath.Join(imagePath, fileName))
	if err != nil {
		return "", err
	}

	bucket := os.Getenv("DO_BUCKET")
	folder := os.Getenv("DO_FOLDER")

	pathOnly := sp.getPicturePath(originalImagePath)
	key :=  pathOnly + fileName
	key = fmt.Sprintf("%s/%s", folder, key)

	//fmt.Println("key: ", key)

	object := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   imgFile,
	}

	_, err = sp.s3Client.PutObject(object)
	if err != nil {
		return "", err
	}

	return pathOnly + fileName, nil

}

func (sp Spaces) ListAllObjectsInAFolder(folderName string) error {

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String("drugstore-prac"),
	}

	objects, err := sp.s3Client.ListObjectsV2(input)
	if err != nil {
		return err
	}

	for _, obj := range objects.Contents {
		fmt.Println(aws.StringValue(obj.Key))
	}

	return nil
}
