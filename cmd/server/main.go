package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"johny-tuna/internal/config"
	"johny-tuna/internal/handler"
	"johny-tuna/internal/repository"
	"johny-tuna/internal/service"
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
	h.Route(r)

	err := r.Run(fmt.Sprintf(":%s", config.Env.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
}
