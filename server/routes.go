package server
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"shards-backend/controllers"
	"shards-backend/utils"
)

// NewRouter <function>
// is used to create a GIN engine instance where all controller and routes will be placed
func NewRouter() *gin.Engine {
	var envVars = utils.GetEnvVars()

	if envVars.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// middlewares
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.Default())

	// static files serving
	router.Static("/images", "./images")

	// endpoints
	v1 := router.Group("v1")
	{
		shards := v1.Group("shards")
		{
			nc := controllers.ShardsController{}

			shards.GET("/", nc.Get)
			shards.GET("/sources", nc.GetSources)
			shards.GET("/types", nc.GetTypes)
		}
	}

	return router
}
