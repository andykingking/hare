package hare_test

import (
	. "github.com/captainpete/hare"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fact", func() {

	Describe("NewFact()", func() {

		var (
			view   *View
			target *Document
			domain *Document
		)

		BeforeEach(func() {
			view, target, domain = &View{}, &Document{}, &Document{}
			view.SetSaved(true)
			target.SetSaved(true)
			domain.SetSaved(true)
		})

		It("returns nil when view is unsaved", func() {
			view.SetSaved(false)
			Expect(NewFact(view, domain, target)).To(BeNil())
		})

		It("returns nil when target is unsaved", func() {
			target.SetSaved(false)
			Expect(NewFact(view, domain, target)).To(BeNil())
		})

		It("returns nil when domain is unsaved", func() {
			domain.SetSaved(false)
			Expect(NewFact(view, domain, target)).To(BeNil())
		})

		It("returns a *Fact with assigned keys", func() {
			view.K, target.K, domain.K = uint64(1), uint64(2), uint64(3)
			fact := NewFact(view, domain, target)

			Expect(fact.KView).To(Equal(view.K))
			Expect(fact.KTarget).To(Equal(target.K))
			Expect(fact.KDomain).To(Equal(domain.K))
		})

	})

	Describe("BucketName()", func() {
		// TODO
	})

	Describe("Key()", func() {
		// TODO
	})

})
