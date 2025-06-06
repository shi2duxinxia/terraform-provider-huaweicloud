---
subcategory: "Image Management Service (IMS)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_images_images"
description: |-
  Use this data source to get available IMS images within HuaweiCloud.
---

# huaweicloud_images_images

Use this data source to get available IMS images within HuaweiCloud.

## Example Usage

```hcl
variable "image_id" {}

data "huaweicloud_images_image" "test" {
  image_id = var.image_id
}

data "huaweicloud_images_images" "ubuntu" {
  name        = "Ubuntu 18.04 server 64bit"
  visibility  = "public"
}

data "huaweicloud_images_images" "centos-1" {
  architecture = "x86"
  os           = "CentOS"
  visibility   = "public"
}

data "huaweicloud_images_images" "centos-2" {
  architecture = "x86"
  name_regex   = "^CentOS 7.4"
  visibility   = "public"
}

data "huaweicloud_images_images" "bms_image" {
  architecture = "x86"
  image_type   = "Ironic"
  visibility   = "public"
}
```

## Argument Reference

* `region` - (Optional, String) Specifies the region in which to obtain the images.
  If omitted, the provider-level region will be used.

* `image_id` - (Optional, String) Specifies the ID of the image.

* `name` - (Optional, String) Specifies the name of the image. Cannot be used simultaneously with `name_regex`.

* `name_regex` - (Optional, String) Specifies the regular expression of the name of the image.
  Cannot be used simultaneously with `name`.

* `image_type` - (Optional, String) Specifies the environment where the image is used.
  The valid values are as follows:
  + **FusionCompute**: Cloud server image, also known as system disk image.
  + **DataImage**: Data disk image.
  + **Ironic**: Bare metal server image.
  + **IsoImage**: ISO image.

* `is_whole_image` - (Optional, Bool) Specifies whether it is a whole image. The valid value is **true** or **false**.
  Defaults to **false**.

* `visibility` - (Optional, String) Specifies the visibility of the image. Must be one of **public**, **private**,
  **market** or **shared**.

* `owner` - (Optional, String) Specifies the owner (UUID) of the image.

* `flavor_id` - (Optional, String) Specifies the ECS flavor ID used to filter out available images.
  You can specify only one flavor ID and only ECS flavor ID is valid, BMS flavor is not supported.

* `sort_key` - (Optional, String) Specifies which field to use for sorting. The valid values are **name**,
  **container_format**, **disk_format**, **status**, **id**, **size**, and **created_at**. Defaults to **name**.

* `sort_direction` - (Optional, String) Specifies whether to sort the query results in ascending or descending order.
  The valid values are as follows:
  + **asc**: Ascending order.
  + **desc**: Descending order.

  Defaults to **asc**.

* `os` - (Optional, String) Specifies the image OS type. The value can be **Windows**, **Ubuntu**, **RedHat**, **SUSE**,
  **CentOS**, **Debian**, **OpenSUSE**, **Oracle Linux**, **Fedora**, **Other**, **CoreOS**, or **EulerOS**.

* `architecture` - (Optional, String) Specifies the image architecture type. The value can be **x86** or **arm**.

* `tag` - (Optional, String) Specifies the image tag in **Key=Value** format.

* `enterprise_project_id` - (Optional, String) Specifies the enterprise project ID of the image.
  For enterprise users, if omitted, will query the images under all enterprise projects.

* `__support_agent_list` - (Optional, String) Specifies whether the image supports host security or host monitoring.
  The valid values are as follows:
  + **hss**: Host security.
  + **ces**: Host monitoring.
  + **hss,ces**: Both support.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `images` - All images that match the filter parameters.
  The [images](#ims_images) structure is documented below.

<a name="ims_images"></a>
The `images` block supports:

* `id` - The ID of the image

* `name` - The name of the image.

* `image_type` - The environment where the image is used.

* `visibility` - The visibility of the image.

* `owner` - The owner (UUID) of the image.

* `os` - The image OS type.

* `architecture` - The image architecture type.

* `enterprise_project_id` - The enterprise project ID of the image.

* `os_version` - The operating system version of the image.

* `file` - The image file download and upload links.

* `schema` - The image view.

* `status` - The status of the image. The valid value is **active**.

* `description` - The description of the image.

* `protected` - Indicates whether the image is protected, protected images cannot be deleted.
  The valid value is **true** or **false**.

* `container_format` - The format of the image's container.

* `min_ram_mb` - The minimum memory required to run an image, in MB unit.

* `max_ram_mb` - The maximum memory supported by the image, in MB unit.

* `min_disk_gb` - The minimum disk space required to run an image, in GB unit.
  + When the operating system is Linux, the value ranges from `10` to `1,024`.
  + When the operating system is Windows, the value ranges from `20` to `1,024`.

* `disk_format` - The image format. The value can be **zvhd2**, **vhd**, **zvhd**, **raw**, **qcow2**, or **iso**.

* `data_origin` - The image source. The format is **server_backup,backup_id**,  **instance,instance_id**,
  **server_backup,vault_id**,  **volume,volume_id**, **file,image_url**, or **image,region,image_id**.

* `backup_id` - The backup ID of the whole image in the CBR vault.

* `size_bytes` - The size of the image file, in bytes unit.

* `active_at` - The time when the image status changes to active, in RFC3339 format.

* `created_at` - The creation time of the image, in RFC3339 format.

* `updated_at` - The last update time of the image, in RFC3339 format.

* `__support_agent_list` - Does the image support host security or host monitoring.
  The valid value is **hss**, **ces**, or **hss,ces**. If it is empty, it means that the image does not support host
  security or host monitoring.
