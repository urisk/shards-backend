package controllers
import (
	"github.com/gin-gonic/gin"
)

// NewsController <controller>
// is used for describing controller actions for news.
type ShardsController struct{}

// Get <function>
// is used to handle get action of news controller which will return <count> number of news.
// url: /v1/news?count=80 , by default <count> = 50
func (nc ShardsController) Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"method":  "v1/shards",
		"message": "Hello from Get function!",
	})
}

// GetSources <function>
// is used to handle get action of shards controller which will return all shard sources.
// url: /v1/shards/sources
func (nc ShardsController) GetSources(c *gin.Context) {
	c.JSON(200, gin.H{
		"method":  "v1/shards/sources",
		"message": "Hello from GetSources function!",
	})
}

// GetTypes <function>
// is used to handle get action of shards controller which will return all shard types.
// url: /v1/news/types
func (nc ShardsController) GetTypes(c *gin.Context) {
	c.JSON(200, gin.H{
		"method":  "v1/shards/types",
		"message": "Hello from GetTypes function!",
	})
}

