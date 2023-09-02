#!/bin/bash

echo "Debug: Lab setup started."
echo "Lab setup is in progress..."

# Check if Rocky Linux ISO exists
if [ ! -f "/tmp/rockylinux-minimal.iso" ]; then
  echo "Debug: Rocky Linux ISO not found. Downloading..."
  wget -O /tmp/rockylinux-minimal.iso https://download.rockylinux.org/pub/rocky/9/isos/x86_64/Rocky-9.2-x86_64-minimal.iso
fi

# Further KVM setup commands will go here

echo "Debug: Lab setup complete."
