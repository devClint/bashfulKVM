#!/bin/bash

# Debug: Checking for Rocky Linux ISO
echo "Debug: Checking for Rocky Linux ISO..."
if [ ! -f /tmp/rockylinux-minimal.iso ]; then
    # Debug: Downloading Rocky Linux ISO
    echo "Debug: Downloading Rocky Linux ISO..."
    wget -O /tmp/rockylinux-minimal.iso https://download.rockylinux.org/pub/rocky/9/isos/x86_64/Rocky-9.2-x86_64-minimal.iso
fi

# Debug: Rocky Linux ISO downloaded successfully
echo "Debug: Rocky Linux ISO downloaded successfully."

# Debug: Creating the virtual machine
echo "Debug: Creating the virtual machine..."
virt-install --name lab_vm1 --ram 1024 --disk size=10 --vcpus 1 --os-type linux --os-variant generic --network bridge=virbr0 --graphics none --console pty,target_type=serial --location '/tmp/rockylinux-minimal.iso' --extra-args 'console=ttyS0,115200n8 serial'

# Debug: Virtual machine created successfully
echo "Debug: Virtual machine created successfully."
