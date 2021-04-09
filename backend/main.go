package main

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := router()
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fcgi.Serve(l, r)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"health": "OK"})
}

func router() *gin.Engine {
	router := gin.Default()

	// ローカル用
	router.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// Cookieが必要ならtrue
		AllowCredentials: false,
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"http://localhost:8000",
		},
		MaxAge: 1 * time.Hour,
	}))

	path := "/api/v1"
	api := router.Group(path)
	{
		api.GET("/health", healthCheck)
		// d := api.Group("/data")
		// {
		// 	d.POST("/add", controller.DateAdd)
		// 	d.GET("/list", controller.DateSearch)
		// }
	}

	return router
}
