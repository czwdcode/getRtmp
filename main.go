package main

import (
	"capture/core"
)

func main() {
	device := core.GetAllEth()
	core.FilterInfo(device)

}
