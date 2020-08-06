package qiniu

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"wozaizhao.com/diancan/common"
)

type uploadForm struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

const (
	blockBits = 22 // Indicate that the blockSize is 4M
	blockSize = 1 << blockBits
)

func blockCount(fsize int64) int {

	return int((fsize + (blockSize - 1)) >> blockBits)
}

func calSha1(b []byte, r io.Reader) ([]byte, error) {

	h := sha1.New()
	_, err := io.Copy(h, r)
	if err != nil {
		return nil, err
	}
	return h.Sum(b), nil
}

func getEtag(file *multipart.FileHeader) (etag string, err error) {

	f, err := file.Open()
	if err != nil {
		common.Log("file.Open", err)
		return
	}
	fsize := file.Size
	blockCnt := blockCount(fsize)
	sha1Buf := make([]byte, 0, 21)

	defer f.Close()
	if blockCnt <= 1 { // file size <= 4M
		sha1Buf = append(sha1Buf, 0x16)
		sha1Buf, err = calSha1(sha1Buf, f)
		if err != nil {
			return
		}
	} else { // file size > 4M
		sha1Buf = append(sha1Buf, 0x96)
		sha1BlockBuf := make([]byte, 0, blockCnt*20)
		for i := 0; i < blockCnt; i++ {
			body := io.LimitReader(f, blockSize)
			sha1BlockBuf, err = calSha1(sha1BlockBuf, body)
			if err != nil {
				return
			}
		}
		sha1Buf, _ = calSha1(sha1Buf, bytes.NewReader(sha1BlockBuf))
	}
	etag = base64.URLEncoding.EncodeToString(sha1Buf)
	return
}

// Upload 上传到七牛云
func Upload(c *gin.Context) {
	key := common.GetQiniuKey()
	var form uploadForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	file, _ := c.FormFile("file")
	f, err := file.Open()
	if err != nil {
		common.Log("file.Open", err)
		return
	}
	etag, _ := getEtag(file)
	defer f.Close()
	bucket := "wozaizhao"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(key["ak"], key["sk"])
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	uploadErr := formUploader.Put(context.Background(), &ret, upToken, etag, io.Reader(f), file.Size, &putExtra)
	if uploadErr != nil {
		common.Log("uploadErr", uploadErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": ret.Key})
}
