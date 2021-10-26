package greeting_test

import (
	"bytes"
	"test/7/greeting"
	"testing"
	"time"

	"golang.org/x/text/language"
)

func mockClock(t *testing.T, v string) greeting.Clock {
	t.Helper()
	now, err := time.Parse("2006/01/02 15:04:05", v)
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	return greeting.ClockFunc(func() time.Time {
		return now
	})
}

func TestGreeting_Do(t *testing.T) {
	defer greeting.ExportSetLang(language.Japanese)()

	g := greeting.Greeting{
		Clock: mockClock(t, "2018/08/31 16:00:00"),
	}

	var buf bytes.Buffer

	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error: ", err)
	}

	if expected, actual := "おはよう", buf.String(); expected != actual {
		t.Errorf("greeting message wont %s but got %s", expected, actual)
	}
}
