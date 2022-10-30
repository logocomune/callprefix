package callprefix

import (
	"reflect"
	"testing"
)

func TestInfo(t *testing.T) {
	type args struct {
		callsign string
	}
	tests := []struct {
		name  string
		args  args
		want  Prefix
		want1 bool
	}{
		{
			name:  "No Prefix info",
			args:  args{"000000"},
			want:  Prefix{},
			want1: false,
		},
		{
			name: "Valid callsign",
			args: args{"iu5pmp"},
			want: Prefix{
				CountryName:    "Italy",
				CQZone:         15,
				ITUZone:        28,
				Continent:      "EU",
				Lat:            42.82,
				Lng:            -12.58,
				TimeOffset:     -1,
				Prefix:         "I",
				IsFullCallsign: false,
				Flag:           "🇮🇹",
				CountryCode:    "IT",
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Info(tt.args.callsign)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Info() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Info() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
