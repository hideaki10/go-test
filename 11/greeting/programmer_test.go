package greeting_test

import (
	"test/11/greeting"
	"testing"
)

func TestIsGopher(t *testing.T) {
	t.Parallel()

	type args struct {
		language string
	}

	test := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Gopherの場合、trueを期待する",
			args: args{
				language: "GoLang",
			},
			want: true,
		},
		{
			name: "PHPerの場合、falseを期待する",
			args: args{
				language: "PHP",
			},
			want: false,
		},
		{
			name: "Dartの場合、falseを期待する",
			args: args{
				language: "DART",
			},
			want: true,
		},
		{
			name: "flutterの場合、falseを期待する",
			args: args{
				language: "Flutter",
			},
			want: false,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := greeting.IsGopher(tt.args.language); got != tt.want {
				t.Errorf("IsGopher() = %v, want %v", got, tt.want)
			}
		})
	}
}
