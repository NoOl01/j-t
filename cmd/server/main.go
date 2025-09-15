package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"johny-tuna/internal/config"
	"johny-tuna/internal/handler"
	"johny-tuna/internal/repository"
	"johny-tuna/internal/service"
	"johny-tuna/internal/utils"
)

// @title Johny Tuna
// @version 1.0
// @BasePath /api/v1
func main() {
	config.LoadEnv()
	db := repository.Connect()
	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	h := handler.NewHandler(srv)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		AllowHeaders:    []string{"Origin", "Content-type", "Accept", "Authorization"},
	}))

	config.BuildExist = utils.BuildCheck()

	h.Route(r)
	if config.BuildExist {
		r.Static("/assets", "./dist/assets")
		r.NoRoute(func(c *gin.Context) {
			c.File("./dist/index.html")
		})
	}

	err := r.Run(fmt.Sprintf(":%s", config.Env.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
}
