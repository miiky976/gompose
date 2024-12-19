package main

import (
	"fmt"
	comp "miiky976/comp/compose"
	"miiky976/comp/compose/modifier"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)
	col := comp.Column(modifier.NewEmptyModifier(), 
		comp.Button(modifier.NewEmptyModifier(), "label1", func(){fmt.Println("Button 1 Clicked")}),
		comp.Button(modifier.NewEmptyModifier(), "nooo", func() {fmt.Println("Button 2 clicked")}),
		comp.Row(modifier.NewEmptyModifier(), 
			comp.Button(modifier.NewEmptyModifier(), "nice", func() {}), 
			comp.Button(modifier.NewEmptyModifier(), "pachuca", func() {}), 
			comp.Button(modifier.NewEmptyModifier(), "beso", func() {}), 
		),
	)
	comp.Init(col)
}
