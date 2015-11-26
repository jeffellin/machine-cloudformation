AWS Cloud Formation Driver for Docker Machine
---

##Setup
First you want to make sure you have a valid cloud formration script. You can either use the one here (github.com) or create your own.  The file is JSON an can be edited either using your favorite text editor or the visual editor within the AWS console. 

You may want to test the cloud formation script in the AWS console before running it with docker-machine. 

Once you have your script you will need to upload it to an S3 bucket.  Make sure that the bucket used is in the same region you will be running cloud formation from. Even though you supply the full S3 URL cloud formation cannot find it if it is located in another region.

The cloud formation driver must be in S3 and cannot exist in on your local file system.

###Installing the Driver
Currently the driver is compiled against docker-machine .5.1. You can download the binary from [Github](https://github.com/jeffellin/machine-cloudformation/releases)  Either copy the binary to the same location as docker-machine or  add its location to your PATH variable.

###Credentials
The driver uses the GO SDK for AWS.  This means that the credentials can be in any of the standard locations that AWS tools expect.  If you have the AWS CLI simply running aws configure is sufficient. Otherwise you will want to set the following environment variables.

|Variable|Description|
|---|---|
|AWS_SECRET|AWS Secret|
|AWS_SECRET_KEY|Aws Secret Key|
|AWS_REGION| Aws region, e.g. us-east-1|

The IAM credentials used should have enough privilidges to invoke cloud formation API.

###Available Driver Options
|Option Name|Description|Default Value|required|
|---|---|---|---|
|cloudformation-url|The url of the cloud formation you wish to use.|none|yes|
|cloudformation-keypath|The path to the private key to use |none|yes|
|cloudformation-keypairname|The name of the private key pair on AWS|none|yes|
|cloudformation-useprivateip|Use the private IP to communicate with the instance|false|no|
|cloudformation-parameters|Key value pairs of additional parameters to pass to cloud formation|none|no|
|cloudformation-ssh-user|username for ssh|ubuntu|no|

Paramaters are passed using the extraparams option and must be in the following format. 

KeyName=KeyValue.  

Multiple parameters can be passed using a pipe (|) as a seperator.

e.g. InstanceType=t1.micro|EBSOptimized=false

```
docker-machine create --cloudformaiton-url https://s3.amazonws.amazon.com/somebucket/cloudformation.json --cloudformation-keypairname mykey --cloudformation-keypath /Users/jellin/.ssh/id_rsa
```

Once the machine is started you can work with it as you would a normal docker-machine.

Use the env, ip, stop, start commands as normal.

You can even remove the stack with the docker-machine rm command.  Removing the docker machine will destroy all resources that cloud formation created.

##Building the driver
checkout the driver to $GOPATH/src/github.com/jeffellin/machine-cloudformation

Run the make script within to build the driver binary.
airname jeff --cloudformation-keypath /Users/jellin/.ssh/id_rsa --cloudformation-use-private-address  dockerdemo
```
- Code has been tested with docker-machine .51
---
Todo
- Fix issue when Stack Creation fails,  the process will hang forver
- Test on .51 which this branch is vendorized with.


