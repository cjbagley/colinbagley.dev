package server_test

import (
	"github.com/cjbagley/colinbagley.dev/internal/server"
	"testing"
)

func TestGetTextDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test date with 'st'", args: args{date: "2024-01-01"}, want: "1st January 2024"},
		{name: "Additional test date with 'st'", args: args{date: "1998-12-21"}, want: "21st December 1998"},
		{name: "Test date with 'th'", args: args{date: "2024-06-30"}, want: "30th June 2024"},
		{name: "Additional test date with 'th'", args: args{date: "2007-11-12"}, want: "12th November 2007"},
		{name: "Test date with 'nd'", args: args{date: "2024-06-22"}, want: "22nd June 2024"},
		{name: "Additional test date with 'nd'", args: args{date: "2024-07-02"}, want: "2nd July 2024"},
		{name: "Test date with 'rd'", args: args{date: "2024-06-23"}, want: "23rd June 2024"},
		{name: "Additional test date with 'rd'", args: args{date: "2024-08-03"}, want: "3rd August 2024"},
		{name: "Incorrect date gives empty string", args: args{date: "29/02/9"}, want: ""},
		{name: "Empty string input gives empty string", args: args{date: ""}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := server.GetTextDate(tt.args.date); got != tt.want {
				t.Errorf("getTextDate() = %v, want %v", got, tt.want)
			}
		})
	}

}
