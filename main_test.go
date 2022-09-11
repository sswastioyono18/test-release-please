package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+2=3",
			args: args{
				a: 1,
				b: 2,
			},
			want: 0,
		},
		{
			name: "2+2=3",
			args: args{
				a: 2,
				b: 2,
			},
			want: 2,
		},
		{
			name: "3=6",
			args: args{
				a: 3,
				b: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
