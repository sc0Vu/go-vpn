package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func createSettingWindow() *widgets.QMainWindow {
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Setting")
	return window
}

func createSettingWidget() *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Please select file")
	input.SetEnabled(false)
	widget.Layout().AddWidget(input)

	fileButton := widgets.NewQPushButton2("Select .ovpn file", nil)
	fileButton.ConnectClicked(func(bool) {
		fileDialog := widgets.NewQFileDialog(nil, core.Qt__Window)
		fileName := fileDialog.GetOpenFileName(nil, "Select .ovpn file", "", "*.ovpn", "", widgets.QFileDialog__ReadOnly)
		input.SetPlaceholderText(fileName)
		fileButton.SetEnabled(false)
		startButton.SetEnabled(true)
	})
	widget.Layout().AddWidget(fileButton)

	startButton := widgets.NewQPushButton2("Start", nil)
	startButton.ConnectClicked(func(bool) {
		// start vpn
	})
	startButton.SetEnabled(false)
	widget.Layout().AddWidget(startButton)
	return widget
}

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	settingWindow := createSettingWindow()
	settingWidget := createSettingWidget()
	settingWindow.SetCentralWidget(settingWidget)

	// add context menu and trayicon
	menu := widgets.NewQMenu(nil)
	aboutAction := menu.AddAction("About")
	aboutAction.ConnectTriggered(func(bool) {
		widgets.QMessageBox_About(nil, "About Go VPN", fmt.Sprintf("GO VPN v0.0.1\n\nBuilded with qt v6.1.0.\n\nNote: It's still in development. Use at your own risk.\n\n"))
	})
	settingAction := menu.AddAction("Setting")
	settingAction.ConnectTriggered(func(bool) {
		settingWindow.Show()
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
	settingWidget.SetWindowIcon(icon)

	trayIcon.Show()

	app.Exec()
}
