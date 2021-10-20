package app

import (
	"github.com/gin-gonic/gin"
	"os"
)

var (
	router = gin.Default()
)

func StartApplication() {

	r := MapUrls()
	port := os.Getenv("PORT")
	_ = r.Run(":"+port)
}
