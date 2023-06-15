package test

import (
	"fmt"

	"github.com/INFURA/go-ethlibs/eth"
)

func main() {
	res := eth.ToChecksumAddress("0x12ae66cdc592e10b60f9097a7b0d3c59fce29876")
	fmt.Println(res)
}
