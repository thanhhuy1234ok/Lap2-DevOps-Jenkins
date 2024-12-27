package main

import "github.com/gin-gonic/gin"

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var Products []Product

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/products", func(c *gin.Context) {
		if len(Products) == 0 {
			c.JSON(200, gin.H{"message": "Không có sản phẩm"})
		} else {
			c.JSON(200, Products)
		}
	})

	r.Run(":3000")
}
