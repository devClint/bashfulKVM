# bashfulKVM

## About
This project aims to provide a simplified GUI to set up a lab environment using KVM. 
It's written in Go and Bash, and the GUI is basic HTML. 
The lab uses Rocky Linux for the virtual machines.

## Features
1. Web GUI to start the lab setup process
2. Downloads Rocky Linux minimal ISO if not already downloaded
3. Shows the progress of the lab setup

## Usage
1. Clone this repository.
2. Make sure you have Go and KVM installed.
3. Run `go run go_web_server_with_debugging.go`.
4. Open your web browser and go to `http://localhost:4099`.
5. Click the "Start Lab" button.

## Dependencies
1. Go (for the web server)
2. KVM (for the virtual machines)
3. wget (for downloading ISO)

## Version
- bashfulKVM v04
