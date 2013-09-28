package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z, lim := 0.4, 0.0001
    last := 0.0

    for math.Abs( last - z ) > lim {
    	last = z
    	z = z - ( ( math.Pow( z, 2 ) - x ) / 2 * z )
    }

    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println( math.Sqrt(2))
}