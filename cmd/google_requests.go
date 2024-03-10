package cmd

import (
	"context"
	"errors"
	"fmt"
	"googlemaps.github.io/maps"
	"log"
	"os"
)

func GetRoute(startingLocation string, endingLocation string) ([]maps.Route, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal error: %s\n", err))
	}

	request := &maps.DirectionsRequest{
		Origin:      startingLocation,
		Destination: endingLocation,
	}
	routes, _, err := client.Directions(context.Background(), request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal error: %s\n", err))
	}

	return routes, nil
}

func GetTravelTime(startingLocation string, endingLocation string, time chan float64, potentialError chan error) {
	routes, err := GetRoute(startingLocation, endingLocation)
	if err != nil {
		time <- 0.0
		potentialError <- err
		log.Fatalf("fatal error: %s", err)
	}

	var routeTimes []float64
	for _, legs := range routes[0].Legs {
		routeTimes = append(routeTimes, float64(legs.Duration))
	}

	totalTime := CalculateTotalTime(routeTimes)
	time <- totalTime
	potentialError <- nil
}
