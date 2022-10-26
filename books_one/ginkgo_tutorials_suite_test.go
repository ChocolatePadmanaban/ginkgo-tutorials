package books_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoTutorials(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoTutorials Suite")
}
