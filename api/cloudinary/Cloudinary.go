package cloudinary

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cloudinaryClient *cloudinary.Cloudinary

func Init(){
	cld, _ := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"), 
		os.Getenv("CLOUDINARY_API_KEY"), 
		os.Getenv("CLOUDINARY_API_SECRET"),
	)

	cloudinaryClient = cld
}

func UploadImage(file *multipart.FileHeader) (*uploader.UploadResult, error) {
	res, err := cloudinaryClient.Upload.Upload(
		context.Background(), 
		file, 
		uploader.UploadParams{},
	)

	if err != nil{
		return nil, err
	}

	return res, nil
}