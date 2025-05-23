package problem10

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Base case",
			args: args{
				input: 10,
			},
			want: 17,
		},
		{
			name: "Test case",
			args: args{
				input: 2_000_000,
			},
			want: 142913828922,
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
