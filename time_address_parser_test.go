package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseDirectoryName(t *testing.T) {
	type args struct {
		directoryName string
	}
	tests := []struct {
		name string
		args args
		want TimeAddress
	}{
		{"", args{"2020-01-01 13-03 Street House"}, TimeAddress{time.Date(2020, time.January, 1, 13, 3, 0, 0, time.UTC), "Street House"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseDirectoryName(tt.args.directoryName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDirectoryName() = %v, want %v", got, tt.want)
			}
		})
	}
}
