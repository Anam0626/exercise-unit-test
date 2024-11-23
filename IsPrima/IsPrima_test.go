package IsPrima_test

import (
	"exercise-unit-test/IsPrima"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("IsPrime", func() {
	Context("when the number is less than or equal to 1", func() {
		It("should return false for 1", func() {
			Expect(IsPrima.IsPrime(1)).To(BeFalse())
		})

		It("should return false for 0", func() {
			Expect(IsPrima.IsPrime(0)).To(BeFalse())
		})

		It("should return false for negative numbers", func() {
			Expect(IsPrima.IsPrime(-5)).To(BeFalse())
		})
	})

	Context("when the number is a prime number", func() {
		It("should return true for 2", func() {
			Expect(IsPrima.IsPrime(2)).To(BeTrue())
		})

		It("should return true for 3", func() {
			Expect(IsPrima.IsPrime(3)).To(BeTrue())
		})

		It("should return true for large prime numbers", func() {
			Expect(IsPrima.IsPrime(17)).To(BeTrue())
			Expect(IsPrima.IsPrime(97)).To(BeTrue())
		})
	})

	Context("when the number is not a prime number", func() {
		It("should return false for even numbers greater than 2", func() {
			Expect(IsPrima.IsPrime(4)).To(BeFalse())
		})

		It("should return false for non-prime odd numbers", func() {
			Expect(IsPrima.IsPrime(9)).To(BeFalse())
		})

		It("should return false for large non-prime numbers", func() {
			Expect(IsPrima.IsPrime(100)).To(BeFalse())
		})
	})
})
