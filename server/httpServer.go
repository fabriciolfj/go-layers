package server

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-layers/controllers"
	"go-layers/repositories"
	"go-layers/services"
	"log"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	runnersController *controllers.RunnersController
	resultsController *controllers.ResultsController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	runnersRepo := repositories.NewRunnersRepository(dbHandler)
	resultsRepo := repositories.NewResultsRepository(dbHandler)
	runnersService := services.NewRunnersService(runnersRepo, resultsRepo)
	resultsService := services.NewResultsService(resultsRepo, runnersRepo)
	runnersController := controllers.NewRunnersController(runnersService)
	resultsController := controllers.NewResultsController(resultsService)

	router := gin.Default()
	router.POST("/runner", runnersController.CreateRunner)
	router.PUT("/runner", runnersController.UpdateRunner)
	router.DELETE("/runner/:id", runnersController.DeleteRunner)
	router.GET("/runner/:id", runnersController.GetRunner)
	router.GET("/runner", runnersController.GetRunnersBatch)

	router.POST("/result", resultsController.CreateResult)
	router.DELETE("/result:id", resultsController.DeelteResult)

	return HttpServer{
		config:            config,
		router:            router,
		runnersController: runnersController,
		resultsController: resultsController,
	}
}

func (hs *HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
