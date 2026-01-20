---
title: "Quick testing with Terraform without a backend"
date: 2026-01-20
draft: false
tags: ["terraform"]
categories: []
---

When working with complex Terraform expressions, such as string manipulation
using regular expressions, it is often useful to **test them in isolation**
without needing to connect to a remote backend or authenticate with a cloud
provider.

## Isolated testing using a plan

To test an expression, you can create a temporary configuration file, for
example `/tmp/tf-test/main.tf`, with your variables and outputs:

```terraform
variable "test_inputs" {
  default = ["CompanyEvent", "HTTPSEvent", "URLParser", "APIKeyEvent", "ShipmentEvent"]
}

# This pattern might not produce the desired snake_case output
output "current_pattern" {
  value = {
    for input in var.test_inputs :
    input => lower(join("_", regexall("[A-Z][a-z]*", input)))
  }
}

# This pattern uses a lookahead-like approach to match Go's convention
output "fixed_pattern" {
  value = {
    for input in var.test_inputs :
    input => lower(replace(input, "/([a-z])([A-Z])/", "$${1}_$${2}"))
  }
}
```

You can then initialize [Terraform](https://developer.hashicorp.com/terraform)
(or [OpenTofu](https://opentofu.org/) while explicitly skipping the backend
configuration. This allows you to run a plan and inspect the results
immediately.

> [!NOTE] Tested version
>
> The examples below were verified using **OpenTofu v1.11.3**.

```bash
terraform init -backend=false
terraform plan -out=tfplan
```

Results:

```sh
$ terraform plan -out=tfplan

Changes to Outputs:
  + current_pattern = {
      + APIKeyEvent   = "a_p_i_key_event"
      + CompanyEvent  = "company_event"
      + HTTPSEvent    = "h_t_t_p_s_event"
      + ShipmentEvent = "shipment_event"
      + URLParser     = "u_r_l_parser"
    }
  + fixed_pattern   = {
      + APIKeyEvent   = "apikey_event"
      + CompanyEvent  = "company_event"
      + HTTPSEvent    = "httpsevent"
      + ShipmentEvent = "shipment_event"
      + URLParser     = "urlparser"
    }
```

You can also inspect the whole plan with something like:

```sh
$ terraform show -json tfplan | jq '.planned_values.outputs'

{
  "current_pattern": {
    "sensitive": false,
    "type": [
      "object",
      {
        "APIKeyEvent": "string",
        "CompanyEvent": "string",
        "HTTPSEvent": "string",
        "ShipmentEvent": "string",
        "URLParser": "string"
      }
    ],
    "value": {
      "APIKeyEvent": "a_p_i_key_event",
      "CompanyEvent": "company_event",
      "HTTPSEvent": "h_t_t_p_s_event",
      "ShipmentEvent": "shipment_event",
      "URLParser": "u_r_l_parser"
    }
  },
  "fixed_pattern": {
    "sensitive": false,
    "type": [
      "object",
      {
        "APIKeyEvent": "string",
        "CompanyEvent": "string",
        "HTTPSEvent": "string",
        "ShipmentEvent": "string",
        "URLParser": "string"
      }
    ],
    "value": {
      "APIKeyEvent": "apikey_event",
      "CompanyEvent": "company_event",
      "HTTPSEvent": "httpsevent",
      "ShipmentEvent": "shipment_event",
      "URLParser": "urlparser"
    }
  }
}


```

### Why use `-backend=false`?

The `-backend=false` flag tells Terraform to skip backend initialization. This
provides several benefits for local testing:

1.  **No remote state storage:** You don't need to configure S3, GCS, or
    Terraform Cloud.
2.  **No authentication:** You don't need to have cloud credentials configured
    or active.
3.  **Speed:** It skips backend validation and state retrieval, making the
    feedback loop much faster.

## Quick testing with Terraform Console

For even faster iteration on one-off expressions, you can use
`terraform console`. This does not require `terraform init` for evaluating basic
expressions:

```bash
terraform console <<< 'lower(replace("CompanyEvent", "/([a-z])([A-Z])/", "$${1}_$${2}"))'
# Output: "company_event"
```
