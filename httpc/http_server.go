package httpc

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RunHttpServerWithGin(ctx context.Context) {
	log.Println("hello")

	 engine :=  gin.Default()
	 routers(engine)
	 engine.Run()
}

func routers(engine *gin.Engine) {
	engine.GET("/hello", func(c *gin.Context) {
			log.Println(c.Request.RequestURI)

			c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
	})
}

func RunNativeHttpServer() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}