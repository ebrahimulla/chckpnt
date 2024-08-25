package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
    if len(os.Args) > 4{
        return 
    }
    
    word := os.Args[1]
    old := os.Args[2]
    new := os.Args[3]
	res := ""
    for _, i := range word{
        if string(i) == old{
            res += new
        }else{
            res += string(i)
        }
    }
	for _, c := range res{
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}