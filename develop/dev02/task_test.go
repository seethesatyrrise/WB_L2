package main

import "testing"

func Test_unpackString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"1", args{"a4bc2d5e"}, "aaaabccddddde", false},
		{"2", args{"abcd"}, "abcd", false},
		{"3", args{"A12b2cD3"}, "AAAAAAAAAAAAbbcDDD", false},
		{"4", args{"45"}, "", true},
		{"5", args{"qwe\\4\\5"}, "qwe45", false},
		{"6", args{"qwe\\45"}, "qwe44444", false},
		{"7", args{"qwe\\\\5"}, "qwe\\\\\\\\\\", false},
		{"8", args{"q\\"}, "", true},
		{"9", args{""}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpackString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("unpackString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
