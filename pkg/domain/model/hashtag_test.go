package model

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const FindHashtagsTestcase1 = `
本文
#ハッシュタグ1 #ハッシュタグ2	#ハッシュタグ3
追伸
#ハッシュタグ4　#ハッシュタグ5#ハッシュタグ6
`

var _ = Describe("FindHashtags", func() {
	It("テキストの中にあるハッシュタグを#を取り除いて返す", func() {
		actual := FindHashtags(FindHashtagsTestcase1)
		Expect(actual).To(HaveLen(6))
		Expect(actual[0]).To(Equal("ハッシュタグ1"))
		Expect(actual[1]).To(Equal("ハッシュタグ2"))
		Expect(actual[2]).To(Equal("ハッシュタグ3"))
		Expect(actual[3]).To(Equal("ハッシュタグ4"))
		Expect(actual[4]).To(Equal("ハッシュタグ5"))
		Expect(actual[5]).To(Equal("ハッシュタグ6"))
	})

	It("重複は排除する", func() {
		actual := FindHashtags(`#ハッシュタグ1 #ハッシュタグ1`)

		Expect(actual).To(HaveLen(1))
		Expect(actual).To(Equal([]string{"ハッシュタグ1"}))
	})
})
