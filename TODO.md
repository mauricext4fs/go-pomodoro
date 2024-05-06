# TODO

## Version 6

- x Start Apple dev process to get App store build
- x Export Apple dev keychain to Macbook Air
- x v6 icon 
- x Update Makefile for v6
- x Try to get the countdown font bigger
- Refactor playNotificationSound
- Add "About" item in menu

## Version 7 / 8

- Add counter for completed Pomodoro
- Migrate build.sh to use notarytool instead of altool
- Investigate what is need to make it work for iOS
- Add Clock Animation
- Try to make it work for Windows
- Add History with custom label (work, study, etc...)
- Fix Theme (Font size and color, x Alignment, etc...)
- BUG: Name of binary in build dmg is wrong
- BUG: dmg Sound crash app
- BUG: dmg Notification crash app
- BUG: Notification still make the app crash on Intel
- BUG: Check and replace the log.Fatal so it's not crashing Fyne
- BUG: Notification icon is incorrect
- BUG: 0 Second missing
- BUG: Sometime the app just close by itself without crash report
			 This seems to happened often after the end of a Pomodoro (on notification). 
			 Seems to have gotten especially worse since adding the systray icon. 

