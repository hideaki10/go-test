package greeting

import (
	"fmt"
	"time"

	"github.com/tenntenn/greeting/v2/text"
)

var lang = text.DefaultLang()

func Do() {
	h := time.Now().Hour()
	switch {
	case h >= 4 && h <= 9:

		fmt.Println(text.GoodMorning(lang))
	case h >= 10 && h <= 16:

		fmt.Println(text.Hello(lang))
	default:

		fmt.Println(text.GoodEvening(lang))
	}
}
