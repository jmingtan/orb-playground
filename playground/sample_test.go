package playground_test

import (
	"github.com/jmingtan/orb-playground/playground"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add", func() {
	It("should sum two numbers together", func() {
		Expect(playground.Add(5, 7)).To(Equal(12))
		Expect(playground.Add(2, 5)).To(Equal(7))
		Expect(playground.Add(9, -3)).To(Equal(6))
	})
})
