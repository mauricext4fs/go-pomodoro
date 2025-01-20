# TODO

## Version 8

- x Fix make test is broken!
- Update Build Nr and Icon
- Add logic for linux Build
  - Add target logic in Makefile
  - Conditional build with target?
- Add version in title /| menu
- Change "Minutes and " for ":"??
- Add logger (would hopefully help with the Audio trash log)
- Switch to Data Binding for the countdown: https://docs.fyne.io/explore/binding
- x Try to make it work for Linux
- x Fix: UTF-8 char in Title not showing properly on Linux
- Notification UTF char not showing properly on linux
- Fix: Sound crash with Linux Install (the sound file is not copied and packaged)
    - It search for 'notification.wav' in the "working" path
- Linux notification sound is very weird (robot sounding) and is shorten the  first time it plays
- Fix: 2024/12/07 16:00:11 Fyne error:  Preferences API requires a unique ID, use app.NewWithID() or the FyneApp.toml ID field
2024/12/07 16:00:11   At: /home/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/app/app.go:60


## Version 9

- Add custom label (work, study, etc...)
- BUG: 0 Second missing
- Add Clock Animation
- Try harder to get some of the error thrown by AudioQueueObject.cpp away 

