# TODO

## Version 6

- x Start Apple dev process to get App store build
- x Export Apple dev keychain to Macbook Air
- x v6 icon 
- x Update Makefile for v6
- x Try to get the countdown font bigger
- x Investigate and add fyne asset bundling logic so the DMG will not crash
- x Remove dmg from Version 5 as it crash everywhere
- x Replace color of Countdown text to the same as fyne.important
- x BUG: dmg Notification crash app
- BUG: Notification still make the app crash on Intel
- x !!! Refactor playNotificationSound: 
    Not going to do it... beep lib does not make it easy at all to do this right.
- BUG: Name of binary in build dmg is wrong

## Version 7 / 8

- Find how to enable notification in dmg (notarytool)
- Migrate build.sh to use notarytool instead of altool
- Add counter for completed Pomodoro
- Add "About" item in menu
- Switch to Data Binding for the countdown: https://docs.fyne.io/explore/binding
- Investigate what is need to make it work for iOS
- Add Clock Animation
- Try harder to get some of the error thrown by AudioQueueObject.cpp away 
- Try to make it work for Linux
- Try to make it work for Windows
- Add History with custom label (work, study, etc...)
- BUG: dmg Sound crash app
- BUG: Check and replace the log.Fatal so it's not crashing Fyne
- BUG: Notification icon is incorrect
- BUG: 0 Second missing
- BUG: Sometime the app just close by itself without crash report
			 This seems to happened often after the end of a Pomodoro (on notification). 
			 Seems to have gotten especially worse since adding the systray icon. 

