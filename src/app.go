package src

import (
	"fmt"
	"miiky976/comp/compose"
	"miiky976/comp/compose/basic"
	"miiky976/comp/compose/modifier"
)

func App() compose.Composable {
	return basic.Column(modifier.NewModifier(),
		basic.Button(
			modifier.NewModifier(),
			basic.Label(modifier.NewModifier(), "label0"),
			func() {
				fmt.Println("Button 1 Clicked")
			},
		),
		basic.Button(
			modifier.NewModifier().FitMaxHeight(),
			basic.Label(modifier.NewModifier(), "label1"),
			func() {
				fmt.Println("Button 2 clicked")
			},
		),
		basic.Row(
			modifier.NewModifier(),
			basic.Button(
				modifier.NewModifier(),
				basic.Label(modifier.NewModifier(), "label2"),
				func() {},
			),
			basic.Button(
				modifier.NewModifier(),
				basic.Label(modifier.NewModifier(), "label3"),
				func() {},
			),
			basic.Button(
				modifier.NewModifier(),
				basic.Label(modifier.NewModifier(), "label4"),
				func() {},
			),
		),
	)
}
