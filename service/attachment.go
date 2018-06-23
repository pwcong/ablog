package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pwcong/ablog/model"
)

type AttachmentService struct {
	Base *BaseService
}

func GetExtension(filename string) string {

	if len(filename) < 1 {
		return "unknown"
	}

	res := strings.Split(filename, ".")

	if len(res) < 2 {
		return "unknown"
	}

	return res[1]
}

func (ctx *AttachmentService) GetAttachmentBySymbol(symbol string) (model.Attachment, error) {

	db := ctx.Base.DB

	var attachment model.Attachment

	if notFound := db.Where("symbol = ?", symbol).First(&attachment).RecordNotFound(); notFound {
		return model.Attachment{}, errors.New("attachment is not existed")
	}

	return attachment, nil

}

func (ctx *AttachmentService) SaveAttachment(file *multipart.FileHeader) (model.Attachment, error) {
	src, err := file.Open()
	if err != nil {
		return model.Attachment{}, err
	}

	defer src.Close()

	buffer, err := ioutil.ReadAll(src)
	if err != nil {
		return model.Attachment{}, err
	}

	h := md5.New()
	_, err = h.Write(buffer)
	if err != nil {
		return model.Attachment{}, err
	}

	symbol := hex.EncodeToString(h.Sum(nil))

	var attachment model.Attachment

	db := ctx.Base.DB
	now := time.Now()

	notFound := db.Where("symbol = ?", symbol).First(&attachment).RecordNotFound()
	if notFound {
		attachment = model.Attachment{
			Symbol:   symbol,
			Filename: file.Filename,
			Year:     now.Format("2006"),
			Month:    now.Format("01"),
			Date:     now.Format("02"),
			ExtName:  GetExtension(file.Filename),
		}

		root, err := os.Getwd()
		if err != nil {
			root = filepath.Dir(os.Args[0])
		}

		dir := filepath.Join(root, "public/"+attachment.Year+"/"+attachment.Month+"/"+attachment.Date)

		err = os.MkdirAll(dir, 0666)
		if err != nil {
			return model.Attachment{}, err
		}

		err = ioutil.WriteFile(filepath.Join(dir, attachment.Symbol+"."+attachment.ExtName), buffer, 0666)
		if err != nil {
			return model.Attachment{}, err
		}

		db.Create(&attachment)
	}

	return attachment, nil
}
