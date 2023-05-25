package bling

import (
	_ "embed"
	"testing"
)

func Test_BlingMapFilter(t *testing.T) {
	type args struct {
		bling *Bling
	}
	nobling, err := loadBling(none)
	if err != nil {
		t.Fatal(err)
	}
	lowbling, err := loadBling(low)
	if err != nil {
		t.Fatal(err)
	}
	dfltbling, err := loadBling(dflt)
	if err != nil {
		t.Fatal(err)
	}
	highbling, err := loadBling(high)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "none",
			args: args{
				bling: nobling,
			},
			wantErr: true,
		},
		{
			name: "low",
			args: args{
				bling: lowbling,
			},
			wantErr: true,
		},
		{
			name: "default",
			args: args{
				bling: dfltbling,
			},
			wantErr: true,
		},
		{
			name: "high",
			args: args{
				bling: highbling,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := len(tt.args.bling.Packages)
			got := len(tt.args.bling.PackageMap)
			if got != want {
				t.Errorf("%s packages: got %v, want %v", tt.name, got, want)
				return
			}
			progwant := len(tt.args.bling.Programs)
			proggot := len(tt.args.bling.ProgramMap)
			if proggot != progwant {
				t.Errorf("%s programs: got %v, want %v", tt.name, proggot, progwant)
				return
			}
		})

	}
}
