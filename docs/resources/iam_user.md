---
layout: ""
page_title: "dynatrace_iam_user Resource - terraform-provider-dynatrace"
subcategory: "IAM"
description: |-
  The resource `dynatrace_iam_user` covers user configuration via Account Management API for SaaS Accounts
---

# dynatrace_iam_user (Resource)

-> This resource is excluded by default in the export utility since it is part of the account management API. You can, of course, specify that resource explicitly in order to export it. In that case, don't forget to specify the environment variables `DT_CLIENT_ID`, `DT_ACCOUNT_ID` and `DT_CLIENT_SECRET` for authentication.

-> This resource requires the API token scopes **Allow read access for identity resources (users and groups)** (`account-idm-read`) and **Allow write access for identity resources (users and groups)** (`account-idm-write`)

## Dynatrace Documentation

- Dynatrace IAM - https://www.dynatrace.com/support/help/how-to-use-dynatrace/user-management-and-sso/manage-groups-and-permissions

- Settings API - https://www.dynatrace.com/support/help/how-to-use-dynatrace/user-management-and-sso/manage-groups-and-permissions/iam/iam-getting-started

## Prerequisites

Using this resource requires an OAuth client to be configured within your account settings.
The scopes of the OAuth Client need to include `account-idm-read`, `account-idm-write`, `account-env-read`, `account-env-write`, `iam-policies-management`, `iam:policies:write`, `iam:policies:read`, `iam:bindings:write`, `iam:bindings:read` and `iam:effective-permissions:read`.

Finally the provider configuration requires the credentials for that OAuth Client.
The configuration section of your provider needs to look like this.
```terraform
provider "dynatrace" {
  dt_env_url   = "https://########.live.dynatrace.com/"
  dt_api_token = "######.########################.################################################################"  

  iam_client_id = "######.########"
  iam_account_id = "########-####-####-####-############"
  iam_client_secret = "######.########.################################################################"  
}
```

## Resource Example Usage

```terraform
resource "dynatrace_iam_user" "john_doe_gmail_com" {
  email  = "john.doe@gmail.com"
  groups = [ data.dynatrace_iam_group.Restricted.id ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email` (String)

### Optional

- `groups` (Set of String)

### Read-Only

- `id` (String) The ID of this resource.
