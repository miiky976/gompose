package basic

import (
	"log"
	"miiky976/comp/compose"
	"miiky976/comp/compose/modifier"

	"github.com/gotk3/gotk3/gtk"
)

func Column(modifier modifier.Modifier, content ...compose.Composable) compose.Composable {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, modifier.Space)
	if err != nil {log.Fatal(err)}
	for _, i := range content {
		box.Add(i.Object)
	}
	box.SetHExpand(modifier.FillMaxWidth)
	box.SetVExpand(modifier.FillMaxHeight)
	return compose.Composable{
		Object: box,
	}
}


func Row(modifier modifier.Modifier, content ...compose.Composable) compose.Composable {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, modifier.Space)
	if err != nil {log.Fatal(err)}
	for _, i := range content {
		box.Add(i.Object)
	}
	box.SetHExpand(modifier.FillMaxWidth)
	box.SetVExpand(modifier.FillMaxHeight)
	return compose.Composable{
		Object: box,
	}
}
