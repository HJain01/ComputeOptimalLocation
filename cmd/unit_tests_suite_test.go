package cmd

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUnitTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UnitTests Suite")
}
