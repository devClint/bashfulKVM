#!/bin/bash

# This script sets up the virtual lab environment

# Debug: Show where the script is running
echo "Debug: Script is running."

# Check if the Rocky Linux ISO is already downloaded
if [ ! -f "/tmp/rockylinux-minimal.iso" ]; then
    echo "Debug: Rocky Linux ISO not found. Downloading..."
    wget -O /tmp/rockylinux-minimal.iso https://download.rockylinux.org/pub/rocky/9/isos/x86_64/Rocky-9.2-x86_64-minimal.iso
    echo "Debug: Download complete."
else
    echo "Debug: Rocky Linux ISO found."
fi

# Debug: Show that the script has completed
echo "Debug: Script execution complete."

