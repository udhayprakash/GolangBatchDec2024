# AWS

## AWS CLI

### 1. Install AWS CLI

    $ curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
    $ unzip awscliv2.zip

    $ sudo ./aws/install


### 2. configuration

    In AWs, create IAM User, attach a group and policy, and then under 
        User -> <specific user> -> Security credentials -> Create Access key
    $ code ~/.aws/credentials

        [default]
        aws_access_key_id = YOUR_ACCESS_KEY_ID
        aws_secret_access_key = YOUR_SECRET_ACCESS_KEY


    $ code ~/.aws/config

        [default]
        region = eu-west-1
        output = json

        [profile terraform]
        role_arn = arn:aws:iam::ACCOUNT_ID:role/TerraformAccessRole
        source_profile = default

    $ aws configure get region
    

### Verification

    $ aws --version
    aws-cli/2.23.2 Python/3.12.6 Linux/6.5.0-1025-azure exe/x86_64.ubuntu.20

    $ aws sts get-caller-identity

To View Current Configuration

    $ aws configure list


To get specific configuration

    $ aws configure get aws_access_key_id
    $ aws configure get region
    $ aws configure get output

To set Specific Configuration Values

    $ aws configure set aws_access_key_id AKIAIOSFODNN7EXAMPLE
    $ aws configure set aws_secret_access_key wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
    $ aws configure set region us-west-2
    $ aws configure set output table

To Configure Multiple Profiles

    $ aws configure --profile myprofile

To remove or clear specific configuration values:

    $ aws configure set aws_access_key_id ""
    $ aws configure set region ""


### Using Environment variables

These credentials can be passed as environment variables, instead of storing in config files

    export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
    export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
    export AWS_DEFAULT_REGION=us-east-1


## AWS CLI for s3 bucket

We need to ensure, we have policies attached to the user, for performing action

aws setup

    create a policy, and then 
    create a group with that policy, and 
    attach that group to the user

Bucket Naming Rules:

    Bucket names must be globally unique across all AWS accounts.
    Bucket names must follow DNS naming conventions (e.g., no uppercase letters, no underscores).

aws s3 vs aws s3api

    aws s3 

        - High-level commands for common S3 operations (e.g., upload, download, sync).
        - Easier to use for common tasks; But limited to common operations.	
        - Ex: aws s3 cp, aws s3 ls, aws s3 sync
    
    aws s3api

        - Low-level commands that map directly to the S3 API (e.g., create bucket, set ACLs).
        - More complex, requires knowledge of S3 API parameters.
        - Provides access to all S3 API operations.
        - Ex: aws s3api create-bucket, aws s3api put-object, aws s3api list-objects


To List all existing Buckets:

    $ aws s3  ls
    $ aws s3api list-buckets

To check Availability of a Bucket

    $ aws s3api head-bucket --bucket my-first-bucket

To create new Bucket:

    LocationConstraint is needed for all regions, except us-east-1

    $ aws s3api create-bucket --bucket gotrainingdecjan2025 --region eu-west-1     --create-bucket-configuration LocationConstraint=eu-west-1

        {
            "Location": "http://gotrainingdecjan2025.s3.amazonaws.com/"
        }

To upload a Folder to an S3 Bucket

    $ mkdir myfolder

    $ aws s3 cp myfolder/  s3://gotrainingdecjan2025/ --recursive
    (Fails if the folder is empty)

    $ aws s3 cp myfolder/  s3://gotrainingdecjan2025/ --recursive
    (Doesnt through error, but nothing will happen)

    $ cd myFolder
    $ code myfile.txt


    $ aws s3 cp myfolder/  s3://gotrainingdecjan2025/
    upload failed: myfolder/ to s3://gotrainingdecjan2025/ Parameter validation failed:
    Invalid length for parameter Key, value: 0, valid min length: 1

    $ aws s3 cp myfolder/  s3://gotrainingdecjan2025/ --recursive
    upload: myfolder/myfile.txt to s3://gotrainingdecjan2025/myfile.txt

To download the file from s3 bucket

    $ aws s3 cp s3://gotrainingdecjan2025/  secondFolder/  --recursive
    download: s3://gotrainingdecjan2025/myfile.txt to secondFolder/myfile.txt
    (creates the folder `secondFolder`, if not exists)

To download a Folder from an S3 Bucket

    In console, create a folder `myFirsts3Folder/` and add a file

    $ aws s3 cp s3://gotrainingdecjan2025/myfolder/ . --recursive

To list Files in an S3 Bucket

    $ aws s3 ls s3://gotrainingdecjan2025/

To list Files in a Specific Folder in an S3 Bucket

    $ aws s3 ls s3://gotrainingdecjan2025/myfolder/

To delete a File from an S3 Bucket

    $ aws s3 rm s3://gotrainingdecjan2025/myfile.txt

To delete a Folder from an S3 Bucket

    $ aws s3 rm s3://gotrainingdecjan2025/myfolder/ --recursive

To delete an S3 Bucket

    $ aws s3api delete-bucket --bucket gotrainingdecjan2025
    (bucket must be empty before deletion)

Empty an S3 Bucket

    $ aws s3 rm s3://gotrainingdecjan2025/ --recursive

Change Storage Class of an Object

    Supported storage classes: 
        STANDARD, 
        STANDARD_IA,
        ONEZONE_IA, 
        GLACIER, 
        DEEP_ARCHIVE, 
        INTELLIGENT_TIERING.


    $ aws s3 cp s3://gotrainingdecjan2025/myfile.txt s3://gotrainingdecjan2025/myfile.txt --storage-class GLACIER

Set Storage Class for Uploaded Objects

    $ aws s3 cp myfile.txt s3://gotrainingdecjan2025/ --storage-class STANDARD_IA

Enable Versioning on an S3 Bucket

    $ aws s3api put-bucket-versioning --bucket gotrainingdecjan2025 --versioning-configuration Status=Enabled

Disable Versioning on an S3 Bucket

    $ aws s3api put-bucket-versioning --bucket gotrainingdecjan2025 --versioning-configuration Status=Suspended

List Object Versions

    $ aws s3api list-object-versions --bucket gotrainingdecjan2025 --prefix myfile.txt

Delete a Specific Version of an Object

    $ aws s3api delete-object --bucket gotrainingdecjan2025 --key myfile.txt --version-id <version-id>

Set Bucket Policy

    $ aws s3api put-bucket-policy --bucket gotrainingdecjan2025 --policy file://s3_policies/policy.json

Get Bucket Policy

    $ aws s3api get-bucket-policy --bucket gotrainingdecjan2025

Delete Bucket Policy

    $ aws s3api delete-bucket-policy --bucket gotrainingdecjan2025

    https://awspolicygen.s3.amazonaws.com/policygen.html

Enable Static Website Hosting

    $ aws s3 website s3://gotrainingdecjan2025/ --index-document index.html --error-document error.html


Sync Local Folder with S3 Bucket

    $ aws s3 sync myfolder/ s3://gotrainingdecjan2025/

Sync S3 Bucket with Local Folder

    $ aws s3 sync s3://gotrainingdecjan2025/ myfolder/

Set Object Metadata

    $ aws s3 cp s3://gotrainingdecjan2025/myfile.txt s3://gotrainingdecjan2025/myfile.txt --metadata-directive REPLACE --content-type "text/html"


Generate a Pre-Signed URL

    $ aws s3 presign s3://gotrainingdecjan2025/myfile.txt --expires-in 3600


Enable Server Access Logging

    $ aws s3api put-bucket-logging --bucket gotrainingdecjan2025 --bucket-logging-status file://logging.json

Enable Cross-Region Replication (CRR)

    Create a replication configuration JSON file. Then, apply the configuration:

    $ aws s3api put-bucket-replication --bucket gotrainingdecjan2025 --replication-configuration file://replication.json


Enable Lifecycle Rules

    $ aws s3api put-bucket-lifecycle-configuration --bucket gotrainingdecjan2025 --lifecycle-configuration file://lifecycle.json


Check Object Encryption Status

    $ aws s3api head-object --bucket gotrainingdecjan2025 --key myfile.txt
    (Look for the ServerSideEncryption field in the output)

Enable Default Encryption for a Bucket

    $ aws s3api put-bucket-encryption --bucket gotrainingdecjan2025 --server-side-encryption-configuration file://encryption.json

Disable Default Encryption for a Bucket

    $ aws s3api delete-bucket-encryption --bucket gotrainingdecjan2025


Restore an Object from Glacier

    $ aws s3api restore-object --bucket gotrainingdecjan2025 --key myfile.txt --restore-request '{"Days":7,"GlacierJobParameters":{"Tier":"Bulk"}}'

Check Bucket Location

    $ aws s3api get-bucket-location --bucket gotrainingdecjan2025


Set CORS Configuration

    $ aws s3api put-bucket-cors --bucket gotrainingdecjan2025 --cors-configuration file://cors.json


Get CORS Configuration

    $ aws s3api get-bucket-cors --bucket gotrainingdecjan2025


Delete CORS Configuration

    $ aws s3api delete-bucket-cors --bucket gotrainingdecjan2025
