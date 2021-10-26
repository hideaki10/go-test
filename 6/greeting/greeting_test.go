package greeting_test

import (
	"bytes"
	"test/6/greeting"
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestGreeting_Do(t *testing.T) {
	defer greeting.ExportSetLang(language.Japanese)()

	g := greeting.Greeting{
		Clock: greeting.ClockFunc(
			func() time.Time {
				return time.Date(2021, 10, 25, 07, 30, 58, 0, time.Local)
			}),
	}

	var buf bytes.Buffer

	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error: ", err)
	}

	if expected, actual := "おはよう", buf.String(); expected != actual {
		t.Errorf("greeting message wont %s but got %s", expected, actual)
	}
}
