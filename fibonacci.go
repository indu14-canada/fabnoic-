package main

import (
    "fmt"
)

// fibonacci generates a Fibonacci series up to n terms
func fibonacci(n int) []int {
    if n <= 0 {
        return []int{}
    }

    fibSeries := make([]int, n)
    fibSeries[0] = 0

    if n > 1 {
        fibSeries[1] = 1
        for i := 2; i < n; i++ {
            fibSeries[i] = fibSeries[i-1] + fibSeries[i-2]
        }
    }

    return fibSeries
}

func main() {
    // Number of terms in the Fibonacci series
    n := 10

    fibSeries := fibonacci(n)

    fmt.Printf("Fibonacci series up to %d terms: %v\n", n, fibSeries)
}
