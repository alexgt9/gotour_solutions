package main

import "fmt"

// fibonacci es una función que devuelve
// una función que devuelve un int.
func fibonacci() func() int {
    f0, f1 := 0, 1
    return func() int{
        f2 := f0 + f1
        ret := f0
        f0, f1 = f1, f2
        return ret
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
