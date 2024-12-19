# Compose in Go?

Just a simple proff of concept of how Compose can be implemented in Go.

Made around GTK using gotk3 bindings.

# Screenshots

```go
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
```

![Example of gompose.](https://github.com/miiky976/gompose/blob/master/images/example.png)
