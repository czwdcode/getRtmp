package main

import (
	"getRtmp/core"
)

func main() {
	device := core.GetAllEth()
	core.FilterInfo(device)

}
