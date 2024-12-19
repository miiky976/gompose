package basic

import (
	"miiky976/comp/compose"
	"miiky976/comp/compose/modifier"

	"github.com/gotk3/gotk3/gtk"
)

func Label(modifier modifier.Modifier, text string) compose.Composable {
	lbl, _ := gtk.LabelNew(text)
	lbl.SetHExpand(modifier.FillMaxWidth)
	lbl.SetVExpand(modifier.FillMaxHeight)
	return compose.Composable{
		Object: lbl,
	}
}
