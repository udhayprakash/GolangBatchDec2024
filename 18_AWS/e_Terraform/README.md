# Terraform 

	infrastructure as code (IaC) tool, developed by Hashicorp
	defining infrastructure using declarative configuration files
		- HashiCorp Configuration Language (HCL) or JSON
		
	Provider Ecosystem( AWS, Azure, Google Cloud Platform, etc)
	
	State Management:
		state file acts as a source of truth for current infrastructure
		This file tracks the resources and their dependencies, enabling Terraform to determine what changes need to be applied when configurations are modified
	 

### Core Workflow

Three stages:

1) Write	: Define the infrastructure resources in configuration files.
2) Plan		: Generate an execution plan that outlines what actions Terraform will take to achieve the desired state.
3) Apply	: Execute the proposed changes after user approval, ensuring that resources are created, updated, or destroyed in the correct order.



### Installation

	In Linux, 
	
		Step 1: Update Your System
				
				sudo apt-get update

		Step 2: Install Required Dependencies

				sudo apt-get install -y gnupg software-properties-common curl
		
		Step 3: Add HashiCorp GPG Key

				curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg

		Step 4: Add the HashiCorp Repository
				
				echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list

		Step 5: Install Terraform
				
				sudo apt-get update && sudo apt-get install terraform

		Step 6: Verify the Installation
				
				terraform -version


### Terraform for AWS cloud 

	Step 1: Configure AWS Provider, in main.tf
	
		terraform {
		  required_providers {
			aws = {
			  source  = "hashicorp/aws"
			  version = "~> 5.56"
			}
		  }
		}

		provider "aws" {
		  region = "us-west-2" # Change to your desired region
		}
	
	Step 2: Initialize Your Configuration
		
		terraform init


	Step 3: Create Infrastructure Resources

		resource "aws_instance" "example" {
		  ami           = "ami-0c55b159cbfafe1f0" # Replace with a valid AMI ID for your region
		  instance_type = "t2.micro"
		}

	Step 4: Plan and Apply Changes

		To preview,
		
			terraform plan

		If everything looks good, apply the configuration:
	
			terraform apply
			(Confirm the action by typing "yes")
			
			
			
### Recommended File Structure

	my-terraform-project/
	├── main.tf          # Core infrastructure configuration
	├── variables.tf     # Input variable definitions
	├── outputs.tf       # Outputs from your infrastructure
	├── provider.tf      # Provider configurations
	└── terraform.tfvars  # Default values for variables

	main.tf (starting point of terraform configuration)
		provider "aws" {
		  region = "us-west-2"
		}

		resource "aws_instance" "example" {
		  ami           = "ami-0c55b159cbfafe1f0"
		  instance_type = "t2.micro"
		}

	variables.tf
		This file declares input variables that your configuration will use. It helps in making your configurations dynamic and reusable
		
		variable "region" {
		  description = "The AWS region to deploy to"
		  default     = "us-west-2"
		}

		variable "instance_type" {
		  description = "The type of instance to create"
		  default     = "t2.micro"
		}




 $ terraform init