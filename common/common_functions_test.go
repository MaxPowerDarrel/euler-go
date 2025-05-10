package common

import "testing"

func TestIsDivisibleByAny(t *testing.T) {
	type args struct {
		n        int
		divisors []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Simple false case",
			args: args{
				n:        12,
				divisors: []int{7, 5},
			},
			want: false,
		},
		{
			name: "Simple true case",
			args: args{
				n:        12,
				divisors: []int{3, 5, 6},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDivisibleByAny(tt.args.n, tt.args.divisors); got != tt.want {
				t.Errorf("IsDivisibleByAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Simple false case",
			args: args{
				n: 6,
			},
			want: false,
		},
		{
			name: "Simple true case",
			args: args{
				n: 13,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.n); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple case",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
