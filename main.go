package main

import (
	"os"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("VPN Client")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText(".......")
	widget.Layout().AddWidget(input)

	button := widgets.NewQPushButton2("click", nil)
	button.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	widget.Layout().AddWidget(button)

	// add context menu and trayicon
	menu := widgets.NewQMenu(nil)
	minAction := menu.AddAction("Min")
	minAction.ConnectTriggered(func(bool) {
		window.ShowNormal()
	})
	maxAction := menu.AddAction("Max")
	maxAction.ConnectTriggered(func(bool) {
		window.ShowMaximized()
	})
	hideAction := menu.AddAction("Hide")
	hideAction.ConnectTriggered(func(bool) {
		window.Hide()
	})
	menu.AddSeparator()
	closeAction := menu.AddAction("Close")
	closeAction.ConnectTriggered(func(bool) {
		app.Quit()
	})

	icon := gui.NewQIcon5("./images/icon.png")
	trayIcon := widgets.NewQSystemTrayIcon(nil)
	trayIcon.SetContextMenu(menu)
	trayIcon.SetIcon(icon)
	trayIcon.SetToolTip("VPN client")
	widget.SetWindowIcon(icon)

	trayIcon.Show()
	window.Show()

	app.Exec()
}
