package main

import (
	"fmt"
	"github.com/HJain01/compute-optimal-location/cmd"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	NoOriginError      = fmt.Errorf("you need to supply at least 1 origin")
	NoDestinationError = fmt.Errorf("you need to supply at least 1 destination")
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginError := range c.Errors {
			switch ginError.Err {
			case NoOriginError:
				c.JSON(http.StatusBadRequest, NoOriginError.Error())
			case NoDestinationError:
				c.JSON(http.StatusBadRequest, NoDestinationError.Error())
			default:
				c.JSON(http.StatusInternalServerError, ginError.Err.Error())
			}
		}
	}
}

func main() {
	r := gin.Default()

	r.Use(ErrorHandler())

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "All OK",
		})
	})

	r.GET("/getOptimalLocation", func(c *gin.Context) {
		origins, validOrigin := c.GetQueryArray("origins")
		destinations, validDestination := c.GetQueryArray("destinations")

		if !validOrigin {
			c.Error(NoOriginError)
			return
		}

		if !validDestination {
			c.Error(NoDestinationError)
			return
		}

		optimalLocation, err := cmd.ComputeOptimalLocation(origins, destinations)

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
