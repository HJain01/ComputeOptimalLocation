package cmd

import (
	"context"
	"googlemaps.github.io/maps"
	"log"
	"os"
)

func GetDistance(startingLocation string, endingLocation string) (float64, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return 0.0, err
	}

	request := &maps.DirectionsRequest{
		Origin:      startingLocation,
		Destination: endingLocation,
	}
	routes, _, err := client.Directions(context.Background(), request)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return 0.0, err
	}

	var routeDistances []float64
	for _, legs := range routes[0].Legs {
		routeDistances = append(routeDistances, float64(legs.Meters))
	}

	totalDistance := CalculateTotalDistance(routeDistances)
	return totalDistance, nil
}
