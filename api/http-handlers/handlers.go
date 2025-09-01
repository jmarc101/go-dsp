package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmarc101/go-dsp/internal/auction"
)

func AuctionHandler(c *gin.Context) {
	var req auction.BidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't decode request"})
		return
	}

	resp, err := auction.Run(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
