package fibonacci

import (
    "reflect"
    "testing"
)

func TestFibonacci(t *testing.T) {
    tests := []struct {
        name string
        n    int
        want []int
    }{
        {"0 terms", 0, []int{}},
        {"1 term", 1, []int{0}},
        {"5 terms", 5, []int{0, 1, 1, 2, 3}},
        {"10 terms", 10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Fibonacci(tt.n); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Fibonacci(%d) = %v, want %v", tt.n, got, tt.want)
            }
        })
    }
}
