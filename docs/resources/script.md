---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "installer_script Resource - terraform-provider-installer"
subcategory: ""
description: |-
  installer_script manages an application using a custom script.
  Adding an installer_script resource means that Terraform will install application in the path by running the install_script when creating the resource.
---

# installer_script (Resource)

`installer_script` manages an application using a custom script.

Adding an `installer_script` resource means that Terraform will install application in the `path` by running the `install_script` when creating the resource.

## Example Usage

```terraform
resource "installer_script" "this" {
  path           = "/tmp/installer-myapp-test"
  install_script = <<-EOF
  /bin/bash

  touch /tmp/installer-myapp-test
  chmod +x /tmp/installer-myapp-test
  exit 0
  EOF

  uninstall_script = <<-EOF
  /bin/bash

  rm -f /tmp/installer-myapp-test
  exit 0
  EOF
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `install_script` (String) is the script that will be called by Terraform when executing `terraform plan/apply`.
- `path` (String) is the location of the application installed by the install script. If the application does not exist at path, the resource is considered not exist by Terraform
- `uninstall_script` (String) is the script that will be called by Terraform when executing `terraform destroy`.

### Read-Only

- `id` (String) Internal ID of the resource.


