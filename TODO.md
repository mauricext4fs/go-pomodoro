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
- x BUG: Notification still make the app crash on Intel
- x !!! Refactor playNotificationSound: 
    Not going to do it... beep lib does not make it easy at all to do this right.
- x Resample the notification sound (maybe that will lower the amount of debug output)
- x BUG: Name of binary in build dmg is wrong
- x Find how to enable notification in dmg (notarytool)
- x Migrate build.sh to use notarytool instead of altool
- x BUG: dmg Sound crash app (stupidly... the .wav must still be packed in the .app dir)
- x BUG: Notification icon is incorrect
- x BUG: Sometime the app just close by itself without crash report
			 This seems to happened often after the end of a Pomodoro (on notification). 
			 Seems to have gotten especially worse since adding the systray icon. 

## Version 7 / 8

- Add counter for completed Pomodoro
- Add "About" item in menu
- Switch to Data Binding for the countdown: https://docs.fyne.io/explore/binding
- Investigate what is need to make it work for iOS
- Add Clock Animation
- Try harder to get some of the error thrown by AudioQueueObject.cpp away 
- Try to make it work for Linux
- Try to make it work for Windows
- Add History with custom label (work, study, etc...)
- BUG: 0 Second missing

