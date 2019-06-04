package main

import (
	"flag"
	"fmt"
	"govel/models"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"govel/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	cfgTmp := flag.String("c", "config.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	help := flag.Bool("h", false, "help")
	flag.Parse()
	cfg := *cfgTmp
	if *version {
		fmt.Println("0.1")
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("/")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./api/config")
	cfg = strings.Replace(cfg, ".json", "", 1)
	viper.SetConfigName(cfg)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "UserToken", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	})
	routes := gin.Default()
	routes.Use(corsConfig)
	models.InitBS(viper.GetString("source_file"))
	go controllers.StartGin(viper.GetString("http.port"), routes)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		os.Exit(0)
	}()
	select {}

}
