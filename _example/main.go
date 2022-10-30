package main

import (
	"fmt"
	"github.com/logocomune/callprefix"
)

func main() {

	info, ok := callprefix.Info("iu5pmp")
	if !ok {
		fmt.Println("No info found on callsing")
	}

	fmt.Printf("%+v\n", info)

}
