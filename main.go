package main

//#include "keyboard.h"
import "C"
import (
	"fmt"
	"log"

	"github.com/lxn/walk"
)

var (
	keyNumberPerSecond = 10
	lastAction *walk.Action
)
const (
	BASE_KEY_NUMBER_PER_SECOND float64 = 12.5
)

func main() {
	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Fatal(err)
	}

	icon, err := walk.Resources.Icon("icon.ico")
	if err != nil {
		log.Fatal(err)
	}

	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		log.Fatal(err)
	}

	defer ni.Dispose()

	if err := ni.SetIcon(icon); err != nil {
		log.Fatal(err)
	}

	ni.SetToolTip(fmt.Sprintf("每秒最多输出 N（默认 %.1f） 个相同字符，或退出", BASE_KEY_NUMBER_PER_SECOND))

	actionList := ni.ContextMenu().Actions()

	addInputMaxFrequencyMenuList(actionList)
	actionList.Add(walk.NewSeparatorAction())
	actionList.Add(createMapKeyMenu()) // Ctrl + C/V/F
	actionList.Add(walk.NewSeparatorAction())
	actionList.Add(createExitMenu())

	ni.SetVisible(true)

	if err := ni.ShowInfo("右键图标选择", "右键图标选择击键（默认10）或退出"); err != nil {
		log.Fatal(err)
	}

	go C.Setup()

	// Run the message loop.
	mw.Run()
}

func createExitMenu() *walk.Action {
	exitAction := walk.NewAction()
	exitAction.SetText("E&xit")
	exitAction.Triggered().Attach(func() { 
		walk.App().Exit(0) 
	})
	return exitAction
}

func createMapKeyMenu() *walk.Action {
	action := walk.NewAction()
	action.SetText("映射 Ctrl + C/V/F")

	open := true
	action.Triggered().Attach(func() {
		open = !open
		openMapKey := 0
		if (open) {
			openMapKey = 1
		}
		C.OpenMapComplexKeyboard(C.int(openMapKey))
		action.SetChecked(open)
	})
	action.SetChecked(open)
	return action
}

func addInputMaxFrequencyMenuList(actionList *walk.ActionList) {
	for i := 1; i <= 3; i++ {
		action := walk.NewAction()
		action.SetText(fmt.Sprintf("击键频率 %.1f", float64(i) * BASE_KEY_NUMBER_PER_SECOND))
		action.Triggered().Attach(setKeyTimesFunc(i, action))
		
		actionList.Add(action)
	}
	setActionChecked(actionList.At(0))
}

func setKeyTimesFunc(i int, action *walk.Action) func() {
	return func() {
		setActionChecked(action)
		keyNumberPerSecond = int(float64(i) * BASE_KEY_NUMBER_PER_SECOND)
		C.UpdateKeyPressRate(C.int(keyNumberPerSecond))
	}
}

func setActionChecked(action *walk.Action) {
	if lastAction != nil {
		lastAction.SetChecked(false)
	}
	action.SetChecked(true)
	lastAction = action
}
