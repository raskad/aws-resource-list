# aws-resource-list

This project is WIP.  
Due to limited real world testing methods, the maintainers cannot guarantee that this project works as intended in all cases.
Any real world feedback is welcome.  
The goal of this project is to list all resources in an aws account and check which resources are not created by an IaC tool like CloudFormation or Terraform.

## Getting Started

You can run the project, by either downloading the latest binary from the [releases page](https://github.com/raskad/aws-resource-list/releases) or building the project from source.

## Example

```bash
# Refresh resources per aws api
aws-resource-list refresh real
aws-resource-list refresh cfn

# Refresh terraform resources
terraform show -json > tf.json
aws-resource-list refresh tf tf.json

# Show resources that are not created via Cloudformation or Terraform
aws-resource-list compare
```

## Contributing

Outstanding work should be documented in issues.  
Pull requests are welcome.
