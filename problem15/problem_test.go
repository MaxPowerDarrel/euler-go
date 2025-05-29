package problem15

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Base case",
			args: args{
				input: 2,
			},
			want: 6,
		},
		{
			name: "Test case",
			args: args{
				input: 20,
			},
			want: 137846528820,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.input); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
