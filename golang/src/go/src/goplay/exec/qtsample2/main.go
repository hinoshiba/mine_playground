package main

import (
    "github.com/therecipe/qt/widgets"
    "github.com/therecipe/qt/gui"
    "os"
)

func main() {

    app := widgets.NewQApplication(len(os.Args), os.Args)
    menu := new(widgets.QMenuBar)

	addAction := widgets.NewQAction2("&Add album...", menu)
	deleteAction := widgets.NewQAction2("&Delete album...", menu)
	quitAction := widgets.NewQAction2("&Quit", menu)
	//aboutAction := widgets.NewQAction2("&About", menu)
	//aboutQtAction := widgets.NewQAction2("About &Qt", menu)

	addAction.SetShortcut(gui.QKeySequence_FromString("Ctrl+A", 0))
	deleteAction.SetShortcut(gui.QKeySequence_FromString("Ctrl+D", 0))
	quitAction.SetShortcuts2(gui.QKeySequence__Quit)

/*
	fileMenu := menu.AddMenu2("&File")
	fileMenu.AddActions([]*widgets.QAction{addAction, deleteAction})
	fileMenu.AddSeparator()
	fileMenu.AddActions([]*widgets.QAction{quitAction})

	helpMenu := menu.AddMenu2("&Help")
	helpMenu.AddActions([]*widgets.QAction{aboutAction, aboutQtAction})
	*/

    window := widgets.NewQMainWindow(nil, 0)
    window.SetWindowTitle("Hello World Example")
    window.SetMinimumSize2(200, 200)
	window.SetMenuBar(menu)
    layout := widgets.NewQVBoxLayout()

    mainWidget := widgets.NewQWidget(nil, 0)
    mainWidget.SetLayout(layout)

    label := widgets.NewQLabel2("Hello World", nil, 0)
    layout.AddWidget(label, 0, 0)

    textEdit := widgets.NewQPlainTextEdit2("", nil)
    layout.AddWidget(textEdit, 0, 0)

    button := widgets.NewQPushButton2("Message", nil)
    layout.AddWidget(button, 0, 0)
    button.ConnectClicked(func(checked bool) {
        widgets.QMessageBox_Information(
            nil,
            "",
            textEdit.ToPlainText(),
            widgets.QMessageBox__Ok,
            widgets.QMessageBox__Ok)
    })

    window.SetCentralWidget(mainWidget)
    window.Show()

    app.Exec()
}
