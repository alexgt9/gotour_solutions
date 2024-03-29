package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    pic := make( [][]uint8, dx )
    for	x:= 0; x < dx ; x++{
        pic[x] = make( []uint8, dy )
        for y:= 0; y < dy ; y++{
         	pic[x][y] = uint8( (x+y)/2 )
        }
    }
    return pic
}

func main() {
    pic.Show(Pic)
}
