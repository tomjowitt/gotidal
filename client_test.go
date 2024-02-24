package gotidal

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

type mockHTTPClient struct {
	FilePath   string
	StatusCode int
}

func (c *mockHTTPClient) Do(req *http.Request) (*http.Response, error) { // nolint:revive // req is unused
	data, err := os.ReadFile(c.FilePath)
	if err != nil {
		return nil, fmt.Errorf("could not load payload file: %w", err)
	}

	buffer := bytes.NewBuffer(data)
	readCloser := io.NopCloser(buffer)

	return &http.Response{
		StatusCode: c.StatusCode,
		Body:       readCloser,
	}, nil
}

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
