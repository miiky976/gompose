package compose

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type Composable struct {
	Object gtk.IWidget
}

type Size struct {
	Width int
	Height int
}

func Init(title string, size Size) *gtk.Window {
	gtk.Init(nil)
	master, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {log.Fatal(err)}

	master.SetTitle(title)
    master.Connect("destroy", func() {
        gtk.MainQuit()
    })
    master.SetDefaultSize(size.Width, size.Height)
	return master
}

func Run(master *gtk.Window, app gtk.IWidget) {
	master.Add(app)
	master.ShowAll()
	gtk.Main()
}
