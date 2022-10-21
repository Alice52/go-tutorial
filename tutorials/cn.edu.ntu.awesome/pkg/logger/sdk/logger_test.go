package sdk

import "testing"

func TestHttpGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "google", args: args{url: "http://www.google.com"}},
		{name: "web", args: args{url: "https://www.liwenzhou.com/"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpGet(tt.args.url)
		})
	}
}
