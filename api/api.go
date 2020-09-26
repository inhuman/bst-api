package api

import (
	"fmt"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/inhuman/bst-api/bst"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type Opts struct {
	BstContainer *bst.Container
	Logger       *zerolog.Logger
}

type Container struct {
	BstContainer *bst.Container
}

func Run(opts Opts) error {

	r := gin.New()

	apiLog := opts.Logger.With().Str("source", "api").Logger()

	con := &Container{
		BstContainer: opts.BstContainer,
	}

	r.Use(logger.SetLogger(logger.Config{
		Logger: &apiLog,
		UTC:    true,
	}))

	r.GET("/search", con.GetByKey)
	r.POST("/insert", con.Insert)
	r.DELETE("/delete", con.DeleteByKey)

	return r.Run(":8080")
}

func (con *Container) GetByKey(c *gin.Context) {

	keyStr := c.Query("val")

	if keyStr != "" {

		key, err := strconv.ParseUint(keyStr, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		val := con.BstContainer.Find(con.BstContainer.Root, int(key))

		if val == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("value for key %d not found", key)},
			)
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"value": val,
		})
	}
}

func (con *Container) DeleteByKey(c *gin.Context) {
	keyStr := c.Query("val")

	if keyStr != "" {

		key, err := strconv.ParseUint(keyStr, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		val := con.BstContainer.Find(con.BstContainer.Root, int(key))

		if val == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("value for key %d not found", key)},
			)
			return
		}

		con.BstContainer.Delete(con.BstContainer.Root, int(key))
		c.AbortWithStatus(http.StatusOK)
	}
}

func (con *Container) Insert(c *gin.Context) {

	// TODO: add post,
}
