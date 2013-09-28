package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func ContadorPalabras(s string) map[string]int {
    counter := make(map[string]int)

    for i, words:= 0,strings.Fields( s ) ; i< len( words ); i++{
        counter[words[i]]++
    }
    return counter
}

func main() {
    wc.Test(ContadorPalabras)
}