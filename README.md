AWS Cloud Formation Driver for Docker Machine
---

This is a driver for use with Docker Machine for creating instances using Amazon Cloud Formation. 

After building the driver (instructions coming soon) place the resulting binary in your path.

The driver assumes that you have a cloud formation template already uplaoded to s3.


```
docker-machine  create --driver amazoncf --cloudformation-url https://s3.amazonaws.com/cformation-jellin/template1 --cloudformation-keypairname jeff --cloudformation-keypath /Users/jellin/.ssh/id_rsa --cloudformation-use-private-address  dockerdemo
```
- Code has been tested with docker-machine .52

