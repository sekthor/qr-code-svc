package controller

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func GetQRCodeFromString(c *gin.Context) {

	str, success := c.Params.Get("string")

	if !success {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "malformed or no parameter provided",
		})
	}

	png, err := qrcode.Encode(str, qrcode.Medium, 256)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not generate qr code",
		})

	}

	c.Data(http.StatusOK, "image/png", png)
}

func GetQRCodeFromB64(c *gin.Context) {

	b64, success := c.Params.Get("string")

	if !success {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "malformed or no parameter provided",
		})
	}

	decode, err := base64.URLEncoding.DecodeString(b64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not decode base64",
		})
	}

	png, err := qrcode.Encode(string(decode), qrcode.Medium, 256)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not generate qr code",
		})

	}

	c.Data(http.StatusOK, "image/png", png)
}
