package greeting

import (
	"golang.org/x/text/language"
)

func ExportSetLang(l language.Tag) func() {

	orglang := lang
	lang = l

	return func() {
		lang = orglang
	}
}
