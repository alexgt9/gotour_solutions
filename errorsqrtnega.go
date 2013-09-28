package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {

		return 0,
			ErrNegativeSqrt(x)
	}

	z, lim := 0.4, 0.0004
	last := 0.0

	for math.Abs(last-z) > lim {
		last = z
		z = z - ((math.Pow(z, 2) - x) / 2 * z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("No se puede calcular la raiz de un numero negativo: %v",
		float64(e))
}
