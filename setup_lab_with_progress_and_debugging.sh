
#!/bin/bash

echo "Debug: Setup lab script started."
echo "Debug: Current working directory: $(pwd)"

# Check if Rocky Linux ISO is already downloaded
if [ ! -f "/tmp/rockylinux-minimal.iso" ]; then
    echo "Debug: Downloading Rocky Linux ISO..."
    wget -O /tmp/rockylinux-minimal.iso https://download.rockylinux.org/pub/rocky/9/isos/x86_64/Rocky-9.2-x86_64-minimal.iso
    if [ $? -ne 0 ]; then
        echo "Debug: Failed to download ISO."
        exit 1
    fi
else
    echo "Debug: ISO already downloaded."
fi

echo "Debug: Lab setup complete."
