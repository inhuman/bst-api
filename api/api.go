package api

import (
	"fmt"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/inhuman/bst-api/interfaces"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type Opts struct {
	BstContainer interfaces.Container
	Logger       *zerolog.Logger
}

type Container struct {
	BstContainer interfaces.Container
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

	r.GET("/search", func(context *gin.Context) {
		con.GetByKey(context)
	})

	r.POST("/insert", func(context *gin.Context) {
		con.Insert(context)
	})

	r.DELETE("/delete", func(context *gin.Context) {
		con.DeleteByKey(context)
	})

	return r.Run(":8080")
}

func (con *Container) GetByKey(c interfaces.GinContext) {

	keyStr := c.Query("val")

	if keyStr != "" {

		key, err := strconv.ParseUint(keyStr, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		val := con.BstContainer.Find(con.BstContainer.GetRoot(), int(key))

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

func (con *Container) DeleteByKey(c interfaces.GinContext) {
	keyStr := c.Query("val")

	if keyStr != "" {

		key, err := strconv.ParseUint(keyStr, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		val := con.BstContainer.Find(con.BstContainer.GetRoot(), int(key))

		if val == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("value for key %d not found", key)},
			)
			return
		}

		con.BstContainer.Delete(con.BstContainer.GetRoot(), int(key))
		c.AbortWithStatus(http.StatusOK)
	}
}

type InsertParams struct {
	Key   int         `json:"key"`
	Value interface{} `json:"value"`
}

func (con *Container) Insert(c interfaces.GinContext) {

	p := InsertParams{}

	err := c.BindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = con.BstContainer.Insert(con.BstContainer.GetRoot(), p.Key, p.Value)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
