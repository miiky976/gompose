package compose

import (
	"log"
	"miiky976/comp/compose/modifier"

	"github.com/gotk3/gotk3/gtk"
)

type Composable struct {
	object gtk.IWidget
}

func Init(comp Composable) {
	// some stuff to init the thing from the main.go for example gtk.Init() idk
	gtk.Init(nil)
	master, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {log.Fatal(err)}

	master.Add(comp.object)

	master.SetTitle("go sway panel")
    master.Connect("destroy", func() {
        gtk.MainQuit()
    })
    master.SetDefaultSize(800, 30)

    master.ShowAll()
	gtk.Main()
}

func Button(modifier modifier.Modifier, label string, onclick func()) Composable{
	// do some stuff and operations
	btn, err := gtk.ButtonNew()
	if err != nil {log.Fatal(err)}
	btn.SetLabel(label)
	btn.Connect("clicked", onclick) 
	btn.SetHExpand(modifier.FillMaxWidth)
	btn.SetVExpand(modifier.FillMaxHeight)
	return Composable{btn}
}

func Column(modifier modifier.Modifier, content ...Composable) Composable {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, modifier.Space)
	if err != nil {log.Fatal(err)}
	for _, i := range content {
		box.Add(i.object)
	}
	box.SetHExpand(modifier.FillMaxWidth)
	box.SetVExpand(modifier.FillMaxHeight)
	return Composable{box}
}


func Row(modifier modifier.Modifier, content ...Composable) Composable {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, modifier.Space)
	if err != nil {log.Fatal(err)}
	for _, i := range content {
		box.Add(i.object)
	}
	box.SetHExpand(modifier.FillMaxWidth)
	box.SetVExpand(modifier.FillMaxHeight)
	return Composable{box}
}
