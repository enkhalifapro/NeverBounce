package main

import (
	"fmt"
	"github.com/enkhalifapro/neverBounce/src"
)

func main() {
	fmt.Println("in main")
	err, neverBounce := neverBounce.New("secret_nvrbnc_golang")
	if err != nil {
		panic(err)
	}
	err, info := neverBounce.Info()
	fmt.Println(info)
	err, singleCheckInfo := neverBounce.Single.Check("enkhalifapro@gmail.com", true, true, "")
	fmt.Println(err)
	fmt.Println(singleCheckInfo)
}
