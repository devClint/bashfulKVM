
# bashfulKVM

A simple lab setup automation tool built with Go and Bash. 

## Features

- Web GUI for a one-click lab setup.
- Uses KVM to create two minimal Rocky Linux VMs.
- Shows the progress of the lab setup on the web GUI.

## Requirements

- Fedora or any other Linux distribution with support for KVM.
- Go programming language.

## How it Works

1. The Go web server serves the web GUI.
2. On clicking the "Start Lab" button, a Bash script (`setup_lab_with_debugging.sh`) is executed.
3. The Bash script downloads the Rocky Linux ISO if it doesn't exist, and sets up two minimal VMs.

## Steps to Run

1. Download and unzip the project.
2. Navigate to the project directory in the terminal.
3. Run the Go web server using `go run go_web_server_with_debugging.go`.
4. Open a web browser and go to `http://localhost:4099`.
5. Click the "Start Lab" button to begin the lab setup.

## Debugging

Check the terminal where the Go web server is running for debugging information.

