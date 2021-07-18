# terraform-provider-lightlytics
Terraform Provider for Lightlytics

## How to build/install

Run the following

`make install`


## Usage

Configure provider credentials and host

```
terraform {
  required_providers {
    lightlytics = {
      version = "0.2"
      source  = "lightlytics.com/api/lightlytics"
    }
  }
}

provider "lightlytics" {
  host = "https://staging.lightops.io"
  username = ""
  password = ""
}
```

Configure AWS account


```
resource "lightlytics_account" "aws" {
  account_type = "AWS"
  aws_account_id = "123234818678"
  display_name = "liran-test"
  aws_regions = ["us-east-1", "us-west-1"]
  stack_region = "us-east-1"
}
```

Find more examples in `/examples` 
