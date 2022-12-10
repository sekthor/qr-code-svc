package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sekthor/qr-code-svc/pkg/controller"
)

func main() {
	router := gin.Default()

	router.GET("/qr/b64/:string", controller.GetQRCodeFromB64)
	router.GET("/qr/str/:string", controller.GetQRCodeFromString)

	router.Run()
}
