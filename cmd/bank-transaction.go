package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmiepr/bank-transaction/internal/service"
)

func main() {
	r := gin.Default()
	db := service.Connect()
	s := &service.Service{DB: db}

	// Simple group: v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/wallet/balance", s.GetBalance)
		v1.PATCH("/wallet/balance", s.CreateTransaction)
	}

	r.Run(":3000")
}
