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
	var times []float64

	for _, startingLocation := range startingLocations {
		c := make(chan float64)
		errorChan := make(chan error)
		GetTravelTime(startingLocation, endingLocation, c, errorChan)
		time := <-c
		err := <-errorChan

		if err != nil {
			return 0.0, errors.New(err.Error())
		}
		times = append(times, time)
	}
	log.Println(times)

	return CalculateVariances(times), nil
}

func CalculateTotalTime(times []float64) float64 {
	totalTime := 0.0

	for _, time := range times {
		totalTime += time
	}

	return totalTime
}

func CalculateVariances(times []float64) float64 {
	var squaredDifference = CalculateSquaredDifference(times)
	return squaredDifference / float64(len(times))
}

func CalculateSquaredDifference(times []float64) float64 {
	var mean = CalculateMean(times)
	var totalSquaredDifference = 0.0

	for _, time := range times {
		var difference = time - mean
		totalSquaredDifference += math.Pow(difference, 2)
	}

	return totalSquaredDifference
}

func CalculateMean(times []float64) float64 {
	var totalTime = CalculateTotalTime(times)

	return totalTime / float64(len(times))
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
