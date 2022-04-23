package main

import (
	"fmt"
	cmd "github.com/HJain01/compute-optimal-location/cmd"
	"github.com/gin-gonic/gin"
)

var (
	NotFoundError   = fmt.Errorf("resource could not be found")
	BadRequestError = fmt.Errorf("you have not supplied the right arguments \nsupply starting and ending locations")
)

func main() {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "All OK",
		})
	})

	r.GET("/getOptimalLocation", func(c *gin.Context) {
		startingLocations, emptyStart := c.GetQueryArray("startingLocations")
		endingLocations, emptyEnd := c.GetQueryArray("endingLocations")

		if emptyStart || len(startingLocations) == 0 {
			c.JSON(400, gin.H{
				"Bad Request": "You need to supply at least one starting location",
			})
		}

		if emptyEnd || len(endingLocations) == 0 {
			c.JSON(400, gin.H{
				"Bad Request": "You need to supply at least one ending location",
			})
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
