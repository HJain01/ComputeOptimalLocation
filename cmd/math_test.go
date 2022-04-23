package cmd

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate Functions", func() {
	Describe("Total Distance", func() {
		DescribeTable("Makes proper calculations no matter the input",
			func(routeDistances []float64, expectedTotalDistance float64) {
				totalDistance := CalculateTotalDistance(routeDistances)

				Expect(totalDistance).To(Equal(expectedTotalDistance))
			},
			Entry("All Positive Numbers", []float64{1.0, 2.0, 3.0, 4.0, 5.0}, 15.0),
			Entry("All Negative Numbers", []float64{-1.0, -2.0, -3.0, -4.0, -5.0}, -15.0),
			Entry("Some Positive and Some Negative Numbers", []float64{1.0, -2.0, 3.0, -4.0, 5.0}, 3.0),
		)
	})
	Describe("Variance", func() {
		DescribeTable("Makes proper calculations no matter the input",
			func(inputs []float64, expectedVariance float64) {
				variance := CalculateVariances(inputs)

				Expect(variance).To(Equal(expectedVariance))
			},
			Entry("Zero Variance in List", []float64{1.0, 1.0, 1.0, 1.0}, 0.0),
			Entry("Zero Variance with Negative Numbers in List", []float64{-1.0, -1.0, -1.0, -1.0}, 0.0),
			Entry("Linear Increase of Numbers in List", []float64{1.0, 2.0, 3.0, 4.0}, 1.25),
			Entry("Linear Increase of Negative Numbers in List", []float64{-1.0, -2.0, -3.0, -4.0}, 1.25),
			Entry("Exponential Increase of Numbers in List", []float64{1.0, 2.0, 4.0, 8.0}, 7.1875),
			Entry("Exponential Increase of Negative Numbers in List", []float64{-1.0, -2.0, -4.0, -8.0}, 7.1875),
		)
	})

	Describe("Squared Difference", func() {
		DescribeTable("Makes proper calculations no matter the input",
			func(distances []float64, expectedSquaredDifference float64) {
				squaredDifference := CalculateSquaredDifference(distances)

				Expect(squaredDifference).To(Equal(expectedSquaredDifference))
			},
			Entry("Zero Difference In Numbers", []float64{1.0, 1.0, 1.0, 1.0, 1.0}, 0.0),
			Entry("Zero Difference with Negative Numbers in List", []float64{-1.0, -1.0, -1.0, -1.0}, 0.0),
			Entry("Linear Increase of Numbers in List", []float64{1.0, 2.0, 3.0, 4.0}, 5.0),
			Entry("Linear Increase of Negative Numbers in List", []float64{-1.0, -2.0, -3.0, -4.0}, 5.0),
			Entry("Exponential Increase of Numbers in List", []float64{1.0, 2.0, 4.0, 8.0}, 28.75),
			Entry("Exponential Increase of Negative Numbers in List", []float64{-1.0, -2.0, -4.0, -8.0}, 28.75),
		)
	})

	Describe("Mean", func() {
		DescribeTable("Makes proper calculations no matter the input",
			func(inputs []float64, expectedMean float64) {
				mean := CalculateMean(inputs)

				Expect(mean).To(Equal(expectedMean))
			},
			Entry("All Positive Numbers", []float64{1.0, 2.0, 3.0, 4.0, 5.0}, 3.0),
			Entry("All Negative Numbers", []float64{-1.0, -2.0, -3.0, -4.0, -5.0}, -3.0),
			Entry("Some Numbers are Negative", []float64{-10.0, 202.0, -3.0, 4.0, 5.0, 6.0}, 34.0),
			Entry("Mean is not an Integer", []float64{0.0, 2.0, 3.0, 4.0, 5.0}, 2.8),
			Entry("Mean is a Repeating Decimal", []float64{10.0, 2202.0, 3.0, 4.0, 5.0, 6.0}, 371+(2.0/3.0)),
		)
	})
})

var _ = Describe("Get Lowest Variance Location Function", func() {
	DescribeTable("Gets the proper lowest variance",
		func(inputs []LocationVariance, expectedLocation string) {
			mostFairLocation := GetLowestVarianceLocation(inputs)

			Expect(mostFairLocation).To(Equal(expectedLocation))
		},
		Entry("All Positive Variances",
			[]LocationVariance{
				{
					Location: "123 GLS Street Good Looking",
					Variance: 0.0,
				},
				{
					Location: "123 Pet Street Lorelei Malcolm",
					Variance: 1.0,
				},
			},
			"123 GLS Street Good Looking",
		),
		Entry("Negative Variances are ignored",
			[]LocationVariance{
				{
					Location: "123 GLS Street Good Looking",
					Variance: -1.0,
				},
				{
					Location: "123 Pet Street Lorelei Malcolm",
					Variance: 1.0,
				},
			},
			"123 Pet Street Lorelei Malcolm",
		),
	)
})
