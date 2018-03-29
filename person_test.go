package person_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/yuyalive/ginkgo-sample"
)

var _ = Describe("Person", func() {

	var (
		kaede  Person
		nagisa Person
	)

	BeforeEach(func() {
		kaede = Person{
			Name:   "kaede",
			Gender: "male",
			Age:    19,
			Height: 180,
			Weight: 60,
		}

		nagisa = Person{
			Name:   "nagisa",
			Gender: "female",
			Age:    25,
			Height: 158,
			Weight: 45,
		}
	})

	Describe("Check Gender", func() {
		Context("Case of male", func() {
			It("should return true", func() {
				Expect(kaede.IsMale()).To(Equal(true))
			})
		})

		Context("Case of famale", func() {
			It("should return false", func() {
				Expect(nagisa.IsMale()).To(Equal(false))
			})
		})
	})

	Describe("Check Height", func() {
		Context("With less than 160cm", func() {
			It("should return true", func() {
				Expect(nagisa.IsLess160cm()).To(Equal(true))
			})
		})

		Context("With more than 160cm", func() {
			It("should return false", func() {
				Expect(kaede.IsLess160cm()).To(Equal(false))
			})
		})
	})

	Describe("Check Weight", func() {
		Context("With more than 50kg", func() {
			It("should return true", func() {
				Expect(kaede.IsMore50kg()).To(Equal(true))
			})
		})

		Context("With less than 50kg", func() {
			It("should return true", func() {
				Expect(nagisa.IsMore50kg()).To(Equal(false))
			})
		})
	})

})
