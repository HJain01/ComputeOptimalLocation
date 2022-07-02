package main

import (
	"fmt"
	"github.com/HJain01/compute-optimal-location/cmd"
	. "github.com/JosephWoodward/gin-errorhandling/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	NotFoundError      = fmt.Errorf("resource could not be found")
	NoDestinationError = fmt.Errorf("you need to supply at least 1 origin")
	NoOriginError      = fmt.Errorf("you need to supply at least 1 destination")
)

func main() {
	r := gin.Default()

	r.Use(
		ErrorHandler(
			Map(NoDestinationError, NoOriginError).ToResponse(func(c *gin.Context, err error) {
				c.Status(http.StatusBadRequest)
				c.Writer.Write([]byte(err.Error()))
			}),
			Map(NotFoundError).ToResponse(func(c *gin.Context, err error) {
				c.Status(http.StatusNotFound)
				c.Writer.Write([]byte(err.Error()))
			}),
		))

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "All OK",
		})
	})

	r.GET("/getOptimalLocation", func(c *gin.Context) {
		startingLocations, emptyStart := c.GetQueryArray("startingLocations")
		endingLocations, emptyEnd := c.GetQueryArray("endingLocations")

		if emptyEnd {
			c.Error(NoDestinationError)
			return
		}

		if emptyStart {
			c.Error(NoOriginError)
			return
		}

		optimalLocation, err := cmd.ComputeOptimalLocation(startingLocations, endingLocations)

		if err != nil {
			c.Error(err)
		}

		c.JSON(200, gin.H{
			"optimalLocation": optimalLocation,
		})

	})

	err := r.Run(":4010")

	if err != nil {
		return
	}
}
