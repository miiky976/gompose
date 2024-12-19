package basic

import (
	"log"
	"miiky976/comp/compose"
	"miiky976/comp/compose/modifier"

	"github.com/gotk3/gotk3/gtk"
)

func Button(modifier modifier.Modifier, content compose.Composable, onclick func()) compose.Composable {
	// do some stuff and operations
	btn, err := gtk.ButtonNew()
	if err != nil {log.Fatal(err)}
	btn.Connect("clicked", onclick) 
	btn.SetHExpand(modifier.FillMaxWidth)
	btn.SetVExpand(modifier.FillMaxHeight)
	btn.Add(content.Object)
	return compose.Composable{
		Object: btn,
	}
}
