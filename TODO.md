# TODO

## Version 8

- Change "Minutes and " for ":"??
- Add version in title /| menu
- Add logger (would hopefully help with the Audio trash log)
- Switch to Data Binding for the countdown: https://docs.fyne.io/explore/binding
- Notification crash on Linux i3
- Fix: 2024/12/07 16:00:11 Fyne error:  Preferences API requires a unique ID, use app.NewWithID() or the FyneApp.toml ID field
2024/12/07 16:00:11   At: /home/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/app/app.go:60
- Makefile Conditional build with target?
- x Fix make test is broken!
- x Notification UTF char not showing properly on linux (suck on it... not going to fix it!)
- x Update Build Nr and Icon
- x Try to make it work for Linux
- x Fix: UTF-8 char in Title not showing properly on Linux
- x Fix: Sound crash with Linux Install (the sound file is not copied and packaged)
    - x It search for 'notification.wav' in the "working" path
- x Linux sound is very weird (robot sounding) and is shorten the  first time it plays
  - on fresh Archlinux install with i3 and piper everything works fine
    - issue seems to only happened in VM Ware so... let's not waste time with this.



## Version 9

- Add custom label (work, study, etc...)
- BUG: 0 Second missing
- Add Clock Animation
- Try harder to get some of the error thrown by AudioQueueObject.cpp away 

