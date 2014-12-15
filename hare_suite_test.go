package hare_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHare(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hare Suite")
}
