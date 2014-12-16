package hare_test

import (
	. "github.com/captainpete/hare"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("View", func() {

	var (
		view *View
	)

	BeforeEach(func() {
		view = &View{}
	})

	Describe("BucketName()", func() {
		It("should be 'views' by default", func() {
			Expect(view.BucketName()).To(Equal([]byte("views")))
		})
	})

})
