package app

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thebogie/smacktalkgaming/config"
	"github.com/thebogie/smacktalkgaming/controllers"
	"github.com/thebogie/smacktalkgaming/db"
	"github.com/thebogie/smacktalkgaming/repos"
	"github.com/thebogie/smacktalkgaming/services"
)

var (
	router = gin.Default()
)

// Run is run
func Run() {

	/*
		====== Setup Database Domains ========
	*/
	dbc := db.InitDB()

	/*
		====== Setup repositories =======
	*/
	userRepo := repos.NewUserRepo(dbc, "users")
	contestRepo := repos.NewContestRepo(dbc, "contests")
	gameRepo := repos.NewGameRepo(dbc, "games")
	venueRepo := repos.NewVenueRepo(dbc, "venues")

	/*
		====== Setup services ===========
	*/
	userService := services.NewUserService(userRepo)
	contestService := services.NewContestService(contestRepo)
	gameService := services.NewGameService(gameRepo)
	venueService := services.NewVenueService(venueRepo)

	/*
		====== Setup controllers ========
	*/
	userCtl := controllers.NewUserController(userService, contestService)
	gameCtl := controllers.NewGameController(gameService)
	contestCtl := controllers.NewContestController(userService, contestService, gameService, venueService)

	/*
		====== Setup middlewares ========
	*/
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(cors.Default())

	/*
		====== Setup routes =============
	*/
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	api := router.Group("/api")

	//update or create
	api.POST("/register", userCtl.Register)
	api.POST("/login", userCtl.Login)
	api.POST("/contests", contestCtl.UpdateContest)
	api.POST("/games", gameCtl.UpdateGame)
	api.GET("/users/:userid", userCtl.GetUser)
	api.GET("/users/:userid/stats/:daterange", userCtl.GetUserStats)

	router.Run(config.Config.API.Port)
}
