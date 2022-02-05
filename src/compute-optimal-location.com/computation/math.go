package computation

import "math"
import "errors"
import "compute-optimal-location.com/model"

func computeOptimalLocation(startingLocations []model.Location, endingLocations []model.Location) model.Location {
	var locationVariances []model.LocationVariance

	for _, endingLocation := range endingLocations {
		var variance = getVariance(endingLocation, startingLocations)
		var locationVariance = model.LocationVariance{
			Location: endingLocation,
			Variance: variance,
		}
		locationVariances = append(locationVariances, locationVariance)
	}

	return getLowestVarianceLocation(locationVariances)
}

func getVariance(endingLocation model.Location, startingLocations []model.Location) float64 {
	var distances []float64

	for _, startingLocation := range startingLocations {
		distances = append(distances, getDistance(startingLocation, endingLocation))
	}

	return calculateVariances(distances)
}

func getDistance(startingLocation model.Location, endingLocation model.Location) float64 {
	errors.New("not currently implemented")
	return 0.0
}

func calculateVariances(distances []float64) float64 {
	var mean = CalculateMean(distances)
	var squaredDifference = calculateSquaredDifference(distances, mean)
	return squaredDifference / float64(len(distances))
}

func CalculateMean(distances []float64) float64 {
	var totalDistance float64 = 0

	for _, distance := range distances {
		totalDistance += distance
	}

	return totalDistance / float64(len(distances))
}

func calculateSquaredDifference(distances []float64, mean float64) float64 {
	var totalSquaredDifference = 0.0

	for _, distance := range distances {
		var difference = distance - mean
		totalSquaredDifference += math.Pow(difference, 2)
	}

	return totalSquaredDifference
}

func getLowestVarianceLocation(locationVariances []model.LocationVariance) model.Location {
	var lowestVariance = math.MaxFloat64
	var bestLocation = locationVariances[0].Location

	for _, locationVariance := range locationVariances {
		if locationVariance.Variance < lowestVariance {
			lowestVariance = locationVariance.Variance
			bestLocation = locationVariance.Location
		}
	}

	return bestLocation
}
