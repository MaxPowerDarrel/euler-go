package problem1

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "Base case",
			input: 10,
			want:  23,
		},
		{
			name:  "Solved case",
			input: 1000,
			want:  233168,
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
