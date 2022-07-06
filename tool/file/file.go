package file

import (
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
)

type ImageBucket struct {
	root string
}

var imgBucket *ImageBucket

func NewImageBucket(r string) *ImageBucket {
	if imgBucket == nil {
		imgBucket = &ImageBucket{
			root: "./media/",
		}
	}
	return imgBucket
}

func (ib *ImageBucket) Save(f multipart.File, fh *multipart.FileHeader) (string, error) {

	name := uuid.NewString() + ".jpg"
	file, err := os.Create(ib.root + name)
	if err != nil {
		log.Println(err)
		return name, err
	}
	io.Copy(file, f)
	err = file.Close()
	return name, err
}
