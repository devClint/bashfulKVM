
#!/bin/bash

# Check if Rocky Linux ISO exists, if not, download it
echo "Debug: Checking for Rocky Linux ISO..."
if [ ! -f "/tmp/rockylinux-minimal.iso" ]; then
    echo "Debug: Downloading Rocky Linux ISO..."
    wget -O /tmp/rockylinux-minimal.iso https://download.rockylinux.org/pub/rocky/9/isos/x86_64/Rocky-9.2-x86_64-minimal.iso
fi

# Check if the download was successful
if [ $? -eq 0 ]; then
    echo "Debug: Rocky Linux ISO downloaded successfully."
else
    echo "Debug: Failed to download Rocky Linux ISO."
    exit 1
fi

# Create a virtual machine
echo "Debug: Creating the virtual machine..."
virt-install     --name lab_vm1     --ram 512     --disk path=/tmp/lab_vm1.img,size=1     --vcpus 1     --os-type linux     --os-variant rhel8.0     --network bridge=virbr0     --graphics none     --console pty,target_type=serial     --location '/tmp/rockylinux-minimal.iso'     --extra-args 'console=ttyS0,115200n8 serial'

# Check if VM was created successfully
if [ $? -eq 0 ]; then
    echo "Debug: Virtual machine created successfully."
else
    echo "Debug: Failed to create the virtual machine."
    exit 1
fi

# Some other lab setup steps can go here

echo "Debug: Lab setup complete."
