package gotidal

import (
	"testing"
)

func Test_lowercaseFirstLetter(t *testing.T) {
	t.Parallel()

	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Empty",
			args{
				"",
			},
			"",
		},
		{
			"All lowercase",
			args{
				"helloworld",
			},
			"helloworld",
		},
		{
			"Camel Case",
			args{
				"HelloWorld",
			},
			"helloWorld",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := lowercaseFirstLetter(tt.args.str); got != tt.want {
				t.Errorf("lowercaseFirstLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_concat(t *testing.T) {
	t.Parallel()

	type args struct {
		strs []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Single string",
			args{
				[]string{"hello"},
			},
			"hello",
		},
		{
			"Double string",
			args{
				[]string{"hello", "world"},
			},
			"helloworld",
		},
		{
			"Triple string",
			args{
				[]string{"hello", "world", "today"},
			},
			"helloworldtoday",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := concat(tt.args.strs...); got != tt.want {
				t.Errorf("concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
