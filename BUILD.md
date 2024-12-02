# Build

Go version >= 1.23 is required

Make sure go/bin is in $PATH
```sh
go install fyne.io/fyne/v2/cmd/fyne@latest
make build 
```

## Specific steps for Windows before build can be possible

- Install msys2
- pacman -Syu
- pacman -Su
- pacman -S mingw-w64-x86_64-gcc
- pacman -S mingw-w64-x86_64-go
- pacman -S zip
- Add the following to shell profile: 

```sh
PATH=$PATH:/mingw64/bin
export GOROOT=/mingw64/lib/go
export GOPATH=/mingw64
```


