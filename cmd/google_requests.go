package cmd

import (
	"context"
	"googlemaps.github.io/maps"
	"log"
	"os"
)

func GetDistance(startingLocation Location, endingLocation Location) (float64, error) {
	client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))

	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return 0.0, err
	}

	origin := startingLocation.Address + startingLocation.City + startingLocation.State
	destination := endingLocation.Address + endingLocation.City + endingLocation.State

	request := &maps.DirectionsRequest{
		Origin:      origin,
		Destination: destination,
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
