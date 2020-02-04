# aws-resource-list

This project is WIP.  
Due to limited real world testing methods, the maintainers cannot guarantee that this project works as intended in all cases.
Any real world feedback is welcome.  
The goal of this project is to list all resources in an aws account and check which resources are not created by an IaC tool like CloudFormation or Terraform.

## Getting Started

You can run the project, by either downloading the latest binary from the [releases page](https://github.com/raskad/aws-resource-list/releases) or building the project from source.

## Implemented resources

A list of resources and their implementation status is located at [resources.yaml](resources.yaml)  
Currently the list of resources is populated by all available CloudFormation resources.  
CloudFormation resources that cannot / will not be implemented are noted in [resourcesBlacklistCloudFormation.yaml](resourcesBlacklistCloudFormation.yaml)

## Contributing

Outstanding work should be documented in issues.  
Pull requests are welcome.
