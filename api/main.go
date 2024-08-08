package main

import (
	"log"

	"github.com/gin-gonic/gin" //import "gin" package

	"github.com/cavdy-play/go_mongo/config"
	"github.com/cavdy-play/go_mongo/routes"
)

// https://github.com/cavdy-play/go_mongo/blob/master/config/db.go

/**
 *Using gin.HandlerChain
 (
	HandlerChain = type HandlerChain []HandlerFunc
 	HandlerFunc = type HandlerFunc func (*Context)
 )
**/

// func A(c *gin.Context) {
// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

// 	if c.Request.Method == "OPTIONS" {
// 		c.AbortWithStatus(204)
// 		return
// 	}
// 	c.Next()
// }

// func B(c *gin.Context) {
// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

// 	if c.Request.Method == "OPTIONS" {
// 		c.AbortWithStatus(204)
// 		return
// 	}
// 	c.Next()
// }

// func CORS() gin.HandlersChain {
// 	return gin.HandlersChain{A, B}
// }

/**
 *Using gin.HandlerFunc
**/

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	// Database
	config.Connect()

	// Init Router
	router := gin.Default()
	// router.Use(CORS()[0])
	router.Use(CORS())

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
