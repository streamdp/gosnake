<p align="center">
  <h3 align="center">gosnake</h3>
  <p align="center">
    <a href="https://github.com/streamdp/gosnake/releases/latest">
      <img alt="GitHub release" src="https://img.shields.io/github/v/release/streamdp/gosnake.svg?logo=github&style=flat-square">
    </a>
    <a href="https://goreportcard.com/report/github.com/streamdp/gosnake">
      <img src="https://goreportcard.com/badge/github.com/streamdp/gosnake" alt="Code Status" />
    </a>
  </p>
</p>
<p align="center">
Another one version of the classic snake game written in golang with a library tcell. This is a test task for the position Golang developer.
</p>

# Build the app

```bash
go build -o bin/gosnake main.go
```

or

```bash
task build
```

# Run the app

```bash
./bin/gosnake
```

or

```
task run
```
If you want to run the game from *Windows*, double-click on __*gosnake.exe*__ (you must use version of __*gosnake.exe*__ for the *Windows* architecture from the released binaries). For *macOS*, use the __*"open with terminal"*__ context menu on __*gosnake*__ and confirm launch with the __*"open"*__ button (you have to use version of __*gosnake*__ for the *Darwin* architecture from the released binaries).

# App builtin help

```bash
$ ./gosnake -h
gosnake is a version of the classic snake game written in golang with a library tcell.

Usage of ./gosnake:
  -h display help
  -heigth int
     set heigth of the game desk (default 20)
  -limit int
     set heigth of the game desk (default 10)
  -width int
     set width of the game desk (default 70)
```

# Test the app

```bash
./bin/gosnake                                                      
```

<p align="center" width="100%">
<img width="33%" src="assets/app_screenshot_2.png">
<img width="33%" src="assets/app_screenshot_3.png">
<img width="33%" src="assets/app_screenshot_4.png">
</p>
