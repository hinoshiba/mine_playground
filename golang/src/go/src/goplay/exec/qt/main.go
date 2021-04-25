package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	core.QCoreApplication_SetApplicationName("sample")
	core.QCoreApplication_SetOrganizationName("sample")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(400, 300)
	window.SetWindowTitle("Hello World")
	window.Show()

	widgets.QApplication_Exec()
}
