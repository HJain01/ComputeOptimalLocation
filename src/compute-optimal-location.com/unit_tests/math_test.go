package unit_tests

import (
	. "compute-optimal-location.com/computation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mean Function", func() {
	DescribeTable("Mean Function Properly calculates mean",
		func(inputs []float64, expectedMean float64) {
			mean := CalculateMean(inputs)

			Expect(mean).To(Equal(expectedMean))
		},
		Entry("test", [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}, 3.0),
	)
})
