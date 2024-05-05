# Braille-grid Serial Service
This is the http server for serial comms between the microcontroller and the extension.
Uses bubbletea for the terminal UI redirects http comms to serial comms.
Only baud rate is configurable oob, stop bits and parity are set for esp32.


## Setup

```bash
go run main.go
```

