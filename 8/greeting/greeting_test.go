package greeting_test

import (
	"bytes"
	"io"
	"test/8/greeting"
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

type errorWriter struct {
	Err error
}

func (w *errorWriter) Writer(p []byte) (n int, err error) {
	return 0, w.Err
}

func TestGreeting_Do(t *testing.T) {
	defer greeting.ExportSetLang(language.Japanese)()

	cases := map[string]struct {
		writer    io.Writer
		clock     greeting.Clock
		msg       string
		expectErr bool
	}{
		"04時": {new(bytes.Buffer), mockClock(t, "2018/08/31 04:00:00"), "おはよう", false},
		"09時": {new(bytes.Buffer), mockClock(t, "2018/08/31 09:00:00"), "おはよう", false},
		"10時": {new(bytes.Buffer), mockClock(t, "2018/08/31 10:00:00"), "こんにちは", false},
		"16時": {new(bytes.Buffer), mockClock(t, "2018/08/31 16:00:00"), "こんにちは", false},
		"17時": {new(bytes.Buffer), mockClock(t, "2018/08/31 17:00:00"), "こんばんは", false},
		"03時": {new(bytes.Buffer), mockClock(t, "2018/08/31 03:00:00"), "こんばんは", false},
		//"エラー": {new(bytes.Buffer), nil, "", true},
	}

	for n, tt := range cases {
		//tt := tt
		t.Run(n, func(t *testing.T) {
			g := greeting.Greeting{
				Clock: tt.clock,
			}
			switch err := g.Do(tt.writer); true {
			case err == nil && tt.expectErr:
				t.Error("")
			case err != nil && !tt.expectErr:
				t.Error("unexpected error", err)
			}

			if buff, ok := tt.writer.(*bytes.Buffer); ok {
				msg := buff.String()
				if msg != tt.msg {
					t.Errorf("greeting msg wont %s but go %s", tt.msg, msg)

				}
			}
		})
	}

}
