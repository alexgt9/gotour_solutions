package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {
    z := 0.4+0i

    for i:=0; i<2000000; i++ {
    	z = z - ( ( cmplx.Pow( z, 3 ) - x ) / 3 * cmplx.Pow( z, 2 ) )
    }

    return z
}

func main() {
    fmt.Println( cmplx.Pow( Cbrt(2), 3 ), Cbrt(2) )
}
