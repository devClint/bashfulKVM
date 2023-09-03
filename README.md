# BashfulKVM

## Project Summary

Created by Clint Nitkiewicz, BashfulKVM is designed to simulate a data center lab environment. Its purpose is to allow technicians to emulate, troubleshoot, and prepare for various scenarios before they occur in production. Inspired by Terraform and Ansible, BashfulKVM aims to offer a realistic experience by leveraging the power of Bash scripting and Go programming. It can deploy tens to hundreds of lightweight Linux server VMs using KVM, even on modest hardware. The ultimate goal is to bridge the gap between classroom learning and hands-on data center experience.

## Detailed Project Scope & Tasks

### Backend Development

1. **Initial Setup**
    - [x] Create a Bash script to download the Rocky Linux ISO
    - [x] Add functionality to check for the existence of the ISO before downloading
    - [x] Implement the VM creation using `virt-install` command in Bash

2. **Go Server Implementation**
    - [x] Initialize a Go web server to handle HTTP requests
    - [x] Implement route handlers for different endpoints (`/`, `/start-lab`)
    - [x] Execute Bash script from Go using `os/exec`

3. **Debugging and Verbose Messages**
    - [x] Add debug and verbose messages to the Go server console
    - [x] Pipe Bash `echo` outputs to appear on the Go console

4. **Error Handling and Logging**
    - [ ] Implement error handling in Bash for failed VM creations
    - [ ] Implement error handling in Go for failed script executions
    - [ ] Create log files to store error messages and other logs

5. **Multiple VM Management**
    - [ ] Extend Bash script to handle naming, destroying, and creating new VMs
    - [ ] Implement user prompts in Go for VM management decisions

### Frontend Development

1. **HTML Interface**
    - [x] Create a basic HTML interface with a "Start Lab" button
    - [ ] Improve the HTML interface to include more options and configurations

2. **Styling and User Interface**
    - [ ] Implement CSS for a more visually appealing interface
    - [ ] Design the layout to be responsive and user-friendly

3. **JavaScript and Dynamic Updates**
    - [ ] Use JavaScript to make asynchronous API calls to the Go server
    - [ ] Update the HTML interface dynamically based on the server's response

### Integration and Testing

1. **End-to-End Integration**
    - [x] Link the HTML interface, Go server, and Bash script together
    - [x] Validate that the whole system works as a single unit

2. **Testing**
    - [ ] Perform unit tests on individual components
    - [ ] Conduct integration tests to ensure end-to-end functionality
    - [ ] Test under different scenarios to emulate data center conditions

### Deployment and Documentation

1. **Documentation**
    - [ ] Create README files explaining each component and their interactions
    - [ ] Develop a user guide for the end-users and a developer guide for contributors

2. **Deployment**
    - [ ] Prepare the project for deployment, considering containerization
    - [ ] Test the deployed version in a live environment

### Real-time Data Center Challenges

1. **Hardware**
    - [ ] Emulate NIC card failures
    - [ ] Emulate memory failures
    - [ ] Emulate power supply failures
    - [ ] Emulate fiber cable failures

2. **Network**
    - [ ] Emulate switch failures
    - [ ] Emulate switch port failures
    - [ ] Emulate network latency and packet loss

3. **Operating System**
    - [ ] Emulate kernel panics and system crashes
    - [ ] Emulate corruption of essential OS files
    - [ ] Emulate GRUB bootloader issues
    - [ ] Emulate disk space exhaustion scenarios

### License

GNU General Public License v3.0

