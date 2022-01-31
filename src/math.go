package src

import "math"
import "errors"

func computeOptimalLocation(startingLocations []location, endingLocations []location) location {
	var locationVariances []locationVariance

	for _, endingLocation := range endingLocations {
		var variance float64 = getVariance(endingLocation, startingLocations)
		var locationVariance = locationVariance{
			Location: endingLocation,
			Variance: variance,
		}
		locationVariances = append(locationVariances, locationVariance)
	}

	return getLowestVarianceLocation(locationVariances)
}

func getVariance(endingLocation location, startingLocations []location) float64 {
	var distances []int

	for _, startingLocation := range startingLocations {
		distances = append(distances, getDistance(startingLocation, endingLocation))
	}

	return calculateVariances(distances)
}

func getDistance(startingLocation location, endingLocation location) int {
	errors.New("not currently implemented")
	return 0
}

func calculateVariances(distances []int) float64 {
	var mean int = calculateMean(distances)
	var squaredDifference float64 = calculateSquaredDifference(distances, mean)
	return squaredDifference / float64(len(distances))
}

func calculateMean(distances []int) int {
	var totalDistance int = 0

	for _, distance := range distances {
		totalDistance += distance
	}

	return totalDistance / len(distances)
}

func calculateSquaredDifference(distances []int, mean int) float64 {
	var totalSquaredDifference float64 = 0.0

	for _, distance := range distances {
		var difference float64 = float64(distance - mean)
		totalSquaredDifference += math.Pow(difference, 2)
	}

	return totalSquaredDifference
}

func getLowestVarianceLocation(locationVariances []locationVariance) location {
	var lowestVariance float64 = math.MaxFloat64
	var bestLocation location = locationVariances[0].Location

	for _, locationVariance := range locationVariances {
		if locationVariance.Variance < lowestVariance {
			lowestVariance = locationVariance.Variance
			bestLocation = locationVariance.Location
		}
	}

	return bestLocation
}
