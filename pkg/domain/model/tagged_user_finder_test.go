package model

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FindTaggedUser", func() {
	It("正常系_取得できた場合",
		func() {
			str := "@aaa bbb @ccc"

			actual := FindTaggedUser(str)
			Expect(len(actual)).To(Equal(2))
			Expect(actual[0]).To(Equal("aaa"))
			Expect(actual[1]).To(Equal("ccc"))
		},
	)

	It("正常系_取得できなかった場合",
		func() {
			str := "aaa bbb ccc"

			actual := FindTaggedUser(str)
			Expect(len(actual)).To(Equal(0))
		},
	)
})
