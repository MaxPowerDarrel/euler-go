package problem3

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "Base case",
			input: 13195,
			want:  29,
		},
		{
			name:  "Test case",
			input: 600851475143,
			want:  6857,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.input); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
