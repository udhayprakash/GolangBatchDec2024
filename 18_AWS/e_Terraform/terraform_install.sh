#!/bin/bash

# Install Terraform
TERRAFORM_VERSION="1.10.4"
TERRAFORM_URL="https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip"

# Download Terraform
echo "Downloading Terraform..."
wget -q $TERRAFORM_URL -O terraform.zip

# Install Terraform
echo "Installing Terraform..."
unzip terraform.zip
sudo mv terraform /usr/local/bin/

# Clean up
rm terraform.zip

# Verify installation
echo "Terraform version:"
terraform --version

echo "Terraform installation complete!"