package cmd

import (
	"errors"
	"log"
	"math"
)

func ComputeOptimalLocation(startingLocations []string, endingLocations []string) (string, error) {
	var locationVariances []LocationVariance

	for _, endingLocation := range endingLocations {
		variance, err := GetVariance(endingLocation, startingLocations)

		if err != nil {
			return "", errors.New(err.Error())
		}

		var locationVariance = LocationVariance{
			Location: endingLocation,
			Variance: variance,
		}
		locationVariances = append(locationVariances, locationVariance)
	}

	return GetLowestVarianceLocation(locationVariances), nil
}

func GetVariance(endingLocation string, startingLocations []string) (float64, error) {
	var distances []float64

	for _, startingLocation := range startingLocations {
		distance, err := GetDistance(startingLocation, endingLocation)

		if err != nil {
			return 0.0, errors.New(err.Error())
		}
		distances = append(distances, distance)
	}
	log.Println(distances)

	return CalculateVariances(distances), nil
}

func CalculateTotalDistance(distances []float64) float64 {
	totalDistance := 0.0

	for _, distance := range distances {
		totalDistance += distance
	}

	return totalDistance
}

func CalculateVariances(distances []float64) float64 {
	var squaredDifference = CalculateSquaredDifference(distances)
	return squaredDifference / float64(len(distances))
}

func CalculateSquaredDifference(distances []float64) float64 {
	var mean = CalculateMean(distances)
	var totalSquaredDifference = 0.0

	for _, distance := range distances {
		var difference = distance - mean
		totalSquaredDifference += math.Pow(difference, 2)
	}

	return totalSquaredDifference
}

func CalculateMean(distances []float64) float64 {
	var totalDistance float64 = 0

	for _, distance := range distances {
		totalDistance += distance
	}

	return totalDistance / float64(len(distances))
}

func GetLowestVarianceLocation(locationVariances []LocationVariance) string {
	var lowestVariance = math.MaxFloat64
	var bestLocation = locationVariances[0].Location

	for _, locationVariance := range locationVariances {
		if locationVariance.Variance < 0.0 {
			continue
		}
		if locationVariance.Variance < lowestVariance {
			lowestVariance = locationVariance.Variance
			bestLocation = locationVariance.Location
		}
	}

	return bestLocation
}
