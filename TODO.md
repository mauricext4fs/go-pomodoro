# TODO

## Version 7

- x Default Notification = Off
- x Improve Notification msg
- x Add all DB logic for Pomodoro History with possibility for custom label
- x Bug: Update Activity overwrite ActivityType
- x Add Count of finish Pomodoro 
- x Add test for UpdateCountPomodoro
- x Try to make it work for Windows
- x Windows Build
  - Need to "bundle" the .wav as it is not working when not present in the same directory
    - Zip file?
- x Find a way to use FyneApp info in Makefile
  Forget it... it's a nightmare and whatever janky solution it wont be cross-platform (windows, mac, linux)
- x Investigate what is needed to make it work for iOS
  Need to change name... 

## Version 8

- Add History with custom label (work, study, etc...)
- Add logger (would hopefully help with the Audio trash log)
- Switch to Data Binding for the countdown: https://docs.fyne.io/explore/binding
- Add Clock Animation
- Try harder to get some of the error thrown by AudioQueueObject.cpp away 
- Try to make it work for Linux

## Version 9

- BUG: 0 Second missing

