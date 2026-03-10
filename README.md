# CoffeeBreak

CoffeeBreak is a small Windows utility written in Go that keeps the machine active by sending a synthetic `Shift` key press when user inactivity reaches a threshold.

## What It Does

- Reads idle time from Windows APIs (`GetLastInputInfo`, `GetTickCount64`)
- Checks idle status every 5 seconds
- Sends `Shift` key down/up with `SendInput` when idle time is >= 290 seconds
- Supports debug mode with `--d`

## Requirements

- Windows (uses `user32.dll` and `kernel32.dll`)
- Go `1.25.6` or newer

## Run

```bash
go run .
```

Debug mode:

```bash
go run . --d
```

## Build

```bash
go build -o CoffeeBreak.exe .
```

Run the executable:

```bash
.\CoffeeBreak.exe
```

With debug:

```bash
.\CoffeeBreak.exe --d
```

## Behavior Details

- Loop interval: `5s`
- Idle threshold for key press: `290s`
- Idle status label switches to `idle` at `>= 300s` in debug output

## Project Files

- `main.go`: app entry point and Windows idle/input logic
- `go.mod`: module and Go version
- `asset/caffe.ico`: icon resource

## License

See [License](./License).

## Author

Stefano Pastore  
p.stefano92@gmail.com
