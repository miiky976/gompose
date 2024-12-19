package main

import (
	"miiky976/comp/compose"
	"miiky976/comp/src"

)

func main() {
	master := compose.Init("Compose in Gotk", compose.Size{400, 800})
	app := src.App()
	compose.Run(master, app.Object)
}
