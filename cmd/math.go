package cmd

import "math"
import "errors"

func computeOptimalLocation(startingLocations []Location, endingLocations []Location) Location {
	var locationVariances []LocationVariance

	for _, endingLocation := range endingLocations {
		var variance = getVariance(endingLocation, startingLocations)
		var locationVariance = LocationVariance{
			Location: endingLocation,
			Variance: variance,
		}
		locationVariances = append(locationVariances, locationVariance)
	}

	return getLowestVarianceLocation(locationVariances)
}

func getVariance(endingLocation Location, startingLocations []Location) float64 {
	var distances []float64

	for _, startingLocation := range startingLocations {
		distances = append(distances, getDistance(startingLocation, endingLocation))
	}

	return calculateVariances(distances)
}

func getDistance(startingLocation Location, endingLocation Location) float64 {
	errors.New("not currently implemented")
	return 0.0
}

func calculateVariances(distances []float64) float64 {
	var squaredDifference = calculateSquaredDifference(distances)
	return squaredDifference / float64(len(distances))
}

func calculateSquaredDifference(distances []float64) float64 {
	var mean = calculateMean(distances)
	var totalSquaredDifference = 0.0

	for _, distance := range distances {
		var difference = distance - mean
		totalSquaredDifference += math.Pow(difference, 2)
	}

	return totalSquaredDifference
}

func calculateMean(distances []float64) float64 {
	var totalDistance float64 = 0

	for _, distance := range distances {
		totalDistance += distance
	}

	return totalDistance / float64(len(distances))
}

func getLowestVarianceLocation(locationVariances []LocationVariance) Location {
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
