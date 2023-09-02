#!/bin/bash

echo "Debug: Checking for Rocky Linux ISO..."

# Check if Rocky Linux ISO exists
if [ ! -f /tmp/rockylinux-minimal.iso ]; then
    echo "Debug: Downloading Rocky Linux ISO..."
    wget -O /tmp/rockylinux-minimal.iso https://download.rockylinux.org/pub/rocky/9/isos/x86_64/Rocky-9.2-x86_64-minimal.iso
    if [ $? -ne 0 ]; then
        echo "Debug: Failed to download Rocky Linux ISO."
        exit 1
    fi
fi

echo "Debug: Rocky Linux ISO downloaded successfully."

echo "Debug: Creating the virtual machine..."

# Create the virtual machine using QEMU (since KVM is not available)
virt-install \
    --name lab_vm1 \
    --os-variant rhel8.0 \
    --ram 512 \
    --disk path=/tmp/lab_vm1.img,size=1 \
    --vcpus 1 \
    --cdrom /tmp/rockylinux-minimal.iso \
    --graphics none \
    --network network=default \
    --noautoconsole \
    --check all=off

if [ $? -ne 0 ]; then
    echo "Debug: Failed to create the virtual machine."
    exit 1
fi

echo "Debug: Virtual machine created successfully."
