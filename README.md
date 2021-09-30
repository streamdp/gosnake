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

Another one version of the classic snake game written in golang with a library tcell.

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

![crazy gameplay screenshot one](assets/app_screenshot_1.png "Crazy game play in progress :)")
![crazy gameplay screenshot two](assets/app_screenshot_2.png "Crazy game play in progress :)")
![crazy gameplay screenshot three](assets/app_screenshot_3.png "Crazy game play in progress :)")
![and game over screen](assets/app_screenshot_4.png "Crazy game play is over :)")
