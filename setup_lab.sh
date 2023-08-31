
#!/bin/bash

# Variables
ROCKY_ISO_URL="https://download.rockylinux.org/pub/rocky/8/isos/x86_64/Rocky-8.4-x86_64-minimal.iso"
ISO_PATH="/var/lib/libvirt/boot/Rocky-8.4-x86_64-minimal.iso"

# Check if Rocky Linux ISO is present
if [ ! -f "$ISO_PATH" ]; then
    echo "Rocky Linux ISO not found. Downloading..."
    wget -O "$ISO_PATH" "$ROCKY_ISO_URL"
fi

# Create VMs (modify the configurations as needed)
sudo virt-install --name server1 --ram 512 --vcpus 1 --disk size=1 --os-type linux --os-variant rhel8 --cdrom "$ISO_PATH" --network default
sudo virt-install --name server2 --ram 512 --vcpus 1 --disk size=1 --os-type linux --os-variant rhel8 --cdrom "$ISO_PATH" --network default

echo "VMs are being created..."
