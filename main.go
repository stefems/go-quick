package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	directory := flag.String("d", "public", "usage for dir?")
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Fatal(http.ListenAndServe(":"+port, nil))

	// router := gin.New()
	// router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/public", "public")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "public/index.html", nil)
	// })

	// router.Run(":" + port)
}
