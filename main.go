package main

import (
	"fmt"
	"log"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main() {
	k, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(k)
}
