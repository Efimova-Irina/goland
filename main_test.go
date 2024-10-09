package main

import "testing"

func Test_div(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:	"Test 1",   
			args:	args{a: 10, b: 2},
			want:   5,
			wantErr: false,
		},
		{
			name:	"Test 2",   
			args:	args{a: 10, b: 0},
			want:   0,
			wantErr: true,
		},
		{
			name:	"Test 3",   
			args:	args{a: 0, b: 6},
			want:   0,
			wantErr: false,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := div(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("div() = %v, want %v", got, tt.want)
			}
		})
	}
}
