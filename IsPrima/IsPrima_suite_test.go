package IsPrima_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIsPrima(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IsPrima Suite")
}
