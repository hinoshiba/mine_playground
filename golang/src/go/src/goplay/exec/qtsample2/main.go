package main

import (
    "github.com/therecipe/qt/widgets"
    "os"
)

func main() {

    app := widgets.NewQApplication(len(os.Args), os.Args)

    window := widgets.NewQMainWindow(nil, 0)
    window.SetWindowTitle("Hello World Example")
    window.SetMinimumSize2(200, 200)
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
