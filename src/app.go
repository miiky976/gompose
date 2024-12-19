package src

import (
	"fmt"
	"miiky976/comp/compose"
	"miiky976/comp/compose/basic"
	"miiky976/comp/compose/modifier"
)

func App() compose.Composable{
	return basic.Column(modifier.NewEmptyModifier(), 
		compose.Button(
			modifier.NewEmptyModifier(), 
			"label1", 
			func(){
				fmt.Println("Button 1 Clicked")
			},
		),
		compose.Button(
			modifier.NewEmptyModifier().FitMaxHeight(), 
			"nooo", 
			func() {
				fmt.Println("Button 2 clicked")
			},
		),
		basic.Row(
			modifier.NewEmptyModifier(), 
			compose.Button(
				modifier.NewEmptyModifier(), 
				"nice", 
				func() {},
			), 
			compose.Button(
				modifier.NewEmptyModifier(), 
				"pachuca", 
				func() {},
			), 
			compose.Button(
				modifier.NewEmptyModifier(), 
				"beso", 
				func() {},
			), 
		),
	)
}
