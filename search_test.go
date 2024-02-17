package gotidal

import (
	"context"
	"errors"
	"testing"
)

func TestClient_Search(t *testing.T) {
	t.Parallel()

	type args struct {
		params SearchParams
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			"Missing query",
			args{
				params: SearchParams{CountryCode: "AU"},
			},
			ErrMissingRequiredParameters,
		},
		{
			"Missing country code",
			args{
				params: SearchParams{Query: "Devo"},
			},
			ErrMissingRequiredParameters,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &Client{}
			_, err := c.Search(context.Background(), tt.args.params)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Client.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
