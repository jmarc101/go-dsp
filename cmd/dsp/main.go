package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	httphandlers "github.com/jmarc101/go-dsp/api/http-handlers"
)

const PORT = 8082

func main() {
	r := gin.Default()

	r.POST("/auction", httphandlers.AuctionHandler)

	err := r.Run(fmt.Sprintf(":%d", PORT))
	if err != nil {
		panic(err)
	}
}
