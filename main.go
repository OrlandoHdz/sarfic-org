package main

import (
	"fmt"

	"github.com/OrlandoHdz/sarfic-org/encrypt"
)

func main() {
	se, err := encrypt.Sencrypt("gsgsgsgs", "fsyuywjhd")

	if err != nil {
		fmt.Println("Ocurrio un error:")
		fmt.Println(err)
		return
	}
	fmt.Println(se)

}
