package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	AppError "github.com/kaitodecode/nyated-backend/common/error"
	"github.com/kaitodecode/nyated-backend/common/response"
	"github.com/kaitodecode/nyated-backend/config"
	"github.com/kaitodecode/nyated-backend/constants"
	"github.com/kaitodecode/nyated-backend/controllers"
	"github.com/kaitodecode/nyated-backend/database/seeder"
	"github.com/kaitodecode/nyated-backend/domain/models"
	middlewares "github.com/kaitodecode/nyated-backend/middleware"
	"github.com/kaitodecode/nyated-backend/repositories"
	"github.com/kaitodecode/nyated-backend/routes"
	"github.com/kaitodecode/nyated-backend/services"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use: "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		_ = godotenv.Load()
		config.Init()
		AppError.Init()
		db, err := config.InitDatabase()

		if err != nil {
			panic(err)
		}

		loc, err := time.LoadLocation("Asia/Jakarta")

		if err != nil {
			panic(err)
		}

		time.Local = loc

		err = db.AutoMigrate(
			&models.Role{},
			&models.User{},
			&models.Folder{},
			&models.Note{},
		)

		if err != nil {
			panic(err)
		}

		seeder.NewSeederRegistry(db).Run()

		repo := repositories.NewRepositoryRegistry(db)
		service := services.NewServiceRegistry(repo)
		controller := controllers.NewControllerRegistry(service)

		router :=  gin.Default()
		router.Use(middlewares.HandlePanic())
		router.NoRoute(func (c *gin.Context)  {
			c.JSON(http.StatusNotFound, response.Response{
				Status: constants.ERROR ,
				Message: fmt.Sprintf("Path %s", http.StatusText(http.StatusNotFound)),
			})
		})
		router.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, response.Response{
				Status: constants.SUCCESS,
				Message: "Welcome to Nyated App",
			})
		})

		router.Use(func(ctx *gin.Context) {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Writer.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, DELETE")
			ctx.Writer.Header().Set("Access-Control-Allow-Header", "Content-Type, Authorization, x-service-name, x-api-key, x-request-at")
			ctx.Next()
		})

		lmt := tollbooth.NewLimiter(
			float64(config.Config.RateLimiterMaxRequest),
			&limiter.ExpirableOptions{
				DefaultExpirationTTL: time.Duration(config.Config.RateLimiterTimeSecond) * time.Second,
			})

		router.Use(middlewares.RateLimiter(lmt))

		group := router.Group("api/")
		route := routes.NewRouteRegistry(controller, group)

		route.Serve()

		port := fmt.Sprintf(":%s", config.Config.Port)
		router.Run(port)
	},
}

func Run(){
	err := command.Execute()
	if err != nil {
		panic(err)
	}
}