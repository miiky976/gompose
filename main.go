package main

import (
	"fmt"
	comp "miiky976/comp/compose"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)
	col := comp.Column(comp.NewEmptyModifier(), 
		comp.Button(comp.NewEmptyModifier(), "label1", func(){fmt.Println("Button 1 Clicked")}),
		comp.Button(comp.NewEmptyModifier(), "nooo", func() {fmt.Println("Button 2 clicked")}),
		comp.Row(comp.NewEmptyModifier(), 
			comp.Button(comp.NewEmptyModifier(), "nice", func() {}), 
			comp.Button(comp.NewEmptyModifier(), "pachuca", func() {}), 
			comp.Button(comp.NewEmptyModifier(), "beso", func() {}), 
		),
	)
	comp.Init(col)
}
