package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sunlit-byte/blog/pkg/config"
	"github.com/sunlit-byte/blog/pkg/database"
	log "github.com/sunlit-byte/blog/pkg/logger"
)

func main() {
	defer flag.Parse()
	// 读取配置文件路径
	envFile := flag.String("env", ".env.development", "Specify the environment file (e.g. .env.development or .env.production)")

	// load env
	err := config.LoadConfig(envFile)
	if err != nil {
		fmt.Printf("load env file failed. err: %v", err)
		os.Exit(-1)
	}

	// init log
	err = log.InitLogger()
	if err != nil {
		fmt.Printf("init logger failed. err: %v", err)
		os.Exit(-2)
	}

	// init db
	err = database.InitDB()
	if err != nil {
		fmt.Printf("init db failed. err: %v", err)
		os.Exit(-3)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		log.Infof("this is handler")
		c.JSON(http.StatusOK, gin.H{
			"message": "welcom to my log",
		})
	})
	r.Run(":8080")
}
