# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ComputeImageCapabilitySchemaArgs', 'ComputeImageCapabilitySchema']

@pulumi.input_type
class ComputeImageCapabilitySchemaArgs:
    def __init__(__self__, *,
                 compartment_id: pulumi.Input[str],
                 compute_global_image_capability_schema_version_name: pulumi.Input[str],
                 image_id: pulumi.Input[str],
                 schema_data: pulumi.Input[Mapping[str, Any]],
                 defined_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 freeform_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a ComputeImageCapabilitySchema resource.
        :param pulumi.Input[str] compartment_id: (Updatable) The OCID of the compartment that contains the resource.
        :param pulumi.Input[str] compute_global_image_capability_schema_version_name: The name of the compute global image capability schema version
        :param pulumi.Input[str] image_id: The ocid of the image
        :param pulumi.Input[Mapping[str, Any]] schema_data: (Updatable) The map of each capability name to its ImageCapabilitySchemaDescriptor.
        :param pulumi.Input[Mapping[str, Any]] defined_tags: (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        :param pulumi.Input[str] display_name: (Updatable) A user-friendly name for the compute image capability schema
        :param pulumi.Input[Mapping[str, Any]] freeform_tags: (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        """
        pulumi.set(__self__, "compartment_id", compartment_id)
        pulumi.set(__self__, "compute_global_image_capability_schema_version_name", compute_global_image_capability_schema_version_name)
        pulumi.set(__self__, "image_id", image_id)
        pulumi.set(__self__, "schema_data", schema_data)
        if defined_tags is not None:
            pulumi.set(__self__, "defined_tags", defined_tags)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if freeform_tags is not None:
            pulumi.set(__self__, "freeform_tags", freeform_tags)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> pulumi.Input[str]:
        """
        (Updatable) The OCID of the compartment that contains the resource.
        """
        return pulumi.get(self, "compartment_id")

    @compartment_id.setter
    def compartment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "compartment_id", value)

    @property
    @pulumi.getter(name="computeGlobalImageCapabilitySchemaVersionName")
    def compute_global_image_capability_schema_version_name(self) -> pulumi.Input[str]:
        """
        The name of the compute global image capability schema version
        """
        return pulumi.get(self, "compute_global_image_capability_schema_version_name")

    @compute_global_image_capability_schema_version_name.setter
    def compute_global_image_capability_schema_version_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "compute_global_image_capability_schema_version_name", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Input[str]:
        """
        The ocid of the image
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="schemaData")
    def schema_data(self) -> pulumi.Input[Mapping[str, Any]]:
        """
        (Updatable) The map of each capability name to its ImageCapabilitySchemaDescriptor.
        """
        return pulumi.get(self, "schema_data")

    @schema_data.setter
    def schema_data(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "schema_data", value)

    @property
    @pulumi.getter(name="definedTags")
    def defined_tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        """
        return pulumi.get(self, "defined_tags")

    @defined_tags.setter
    def defined_tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "defined_tags", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        (Updatable) A user-friendly name for the compute image capability schema
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="freeformTags")
    def freeform_tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        """
        return pulumi.get(self, "freeform_tags")

    @freeform_tags.setter
    def freeform_tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "freeform_tags", value)


@pulumi.input_type
class _ComputeImageCapabilitySchemaState:
    def __init__(__self__, *,
                 compartment_id: Optional[pulumi.Input[str]] = None,
                 compute_global_image_capability_schema_id: Optional[pulumi.Input[str]] = None,
                 compute_global_image_capability_schema_version_name: Optional[pulumi.Input[str]] = None,
                 defined_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 freeform_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 schema_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 time_created: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ComputeImageCapabilitySchema resources.
        :param pulumi.Input[str] compartment_id: (Updatable) The OCID of the compartment that contains the resource.
        :param pulumi.Input[str] compute_global_image_capability_schema_id: The ocid of the compute global image capability schema
        :param pulumi.Input[str] compute_global_image_capability_schema_version_name: The name of the compute global image capability schema version
        :param pulumi.Input[Mapping[str, Any]] defined_tags: (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        :param pulumi.Input[str] display_name: (Updatable) A user-friendly name for the compute image capability schema
        :param pulumi.Input[Mapping[str, Any]] freeform_tags: (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        :param pulumi.Input[str] image_id: The ocid of the image
        :param pulumi.Input[Mapping[str, Any]] schema_data: (Updatable) The map of each capability name to its ImageCapabilitySchemaDescriptor.
        :param pulumi.Input[str] time_created: The date and time the compute image capability schema was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
        """
        if compartment_id is not None:
            pulumi.set(__self__, "compartment_id", compartment_id)
        if compute_global_image_capability_schema_id is not None:
            pulumi.set(__self__, "compute_global_image_capability_schema_id", compute_global_image_capability_schema_id)
        if compute_global_image_capability_schema_version_name is not None:
            pulumi.set(__self__, "compute_global_image_capability_schema_version_name", compute_global_image_capability_schema_version_name)
        if defined_tags is not None:
            pulumi.set(__self__, "defined_tags", defined_tags)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if freeform_tags is not None:
            pulumi.set(__self__, "freeform_tags", freeform_tags)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if schema_data is not None:
            pulumi.set(__self__, "schema_data", schema_data)
        if time_created is not None:
            pulumi.set(__self__, "time_created", time_created)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> Optional[pulumi.Input[str]]:
        """
        (Updatable) The OCID of the compartment that contains the resource.
        """
        return pulumi.get(self, "compartment_id")

    @compartment_id.setter
    def compartment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compartment_id", value)

    @property
    @pulumi.getter(name="computeGlobalImageCapabilitySchemaId")
    def compute_global_image_capability_schema_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ocid of the compute global image capability schema
        """
        return pulumi.get(self, "compute_global_image_capability_schema_id")

    @compute_global_image_capability_schema_id.setter
    def compute_global_image_capability_schema_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute_global_image_capability_schema_id", value)

    @property
    @pulumi.getter(name="computeGlobalImageCapabilitySchemaVersionName")
    def compute_global_image_capability_schema_version_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the compute global image capability schema version
        """
        return pulumi.get(self, "compute_global_image_capability_schema_version_name")

    @compute_global_image_capability_schema_version_name.setter
    def compute_global_image_capability_schema_version_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute_global_image_capability_schema_version_name", value)

    @property
    @pulumi.getter(name="definedTags")
    def defined_tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        """
        return pulumi.get(self, "defined_tags")

    @defined_tags.setter
    def defined_tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "defined_tags", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        (Updatable) A user-friendly name for the compute image capability schema
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="freeformTags")
    def freeform_tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        """
        return pulumi.get(self, "freeform_tags")

    @freeform_tags.setter
    def freeform_tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "freeform_tags", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ocid of the image
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="schemaData")
    def schema_data(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        (Updatable) The map of each capability name to its ImageCapabilitySchemaDescriptor.
        """
        return pulumi.get(self, "schema_data")

    @schema_data.setter
    def schema_data(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "schema_data", value)

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time the compute image capability schema was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
        """
        return pulumi.get(self, "time_created")

    @time_created.setter
    def time_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_created", value)


class ComputeImageCapabilitySchema(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compartment_id: Optional[pulumi.Input[str]] = None,
                 compute_global_image_capability_schema_version_name: Optional[pulumi.Input[str]] = None,
                 defined_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 freeform_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 schema_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        This resource provides the Compute Image Capability Schema resource in Oracle Cloud Infrastructure Core service.

        Creates compute image capability schema.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_compute_image_capability_schema = oci.core.ComputeImageCapabilitySchema("testComputeImageCapabilitySchema",
            compartment_id=var["compartment_id"],
            compute_global_image_capability_schema_version_name=var["compute_image_capability_schema_compute_global_image_capability_schema_version_name"],
            image_id=oci_core_image["test_image"]["id"],
            schema_data=[{
                "descriptorType": var["compute_image_capability_schema_schema_data_descriptor_type"],
                "source": var["compute_image_capability_schema_schema_data_source"],
                "defaultValue": var["compute_image_capability_schema_schema_data_default_value"],
                "values": var["compute_image_capability_schema_schema_data_values"],
            }],
            defined_tags={
                "Operations.CostCenter": "42",
            },
            display_name=var["compute_image_capability_schema_display_name"],
            freeform_tags={
                "Department": "Finance",
            })
        ```

        ## Import

        ComputeImageCapabilitySchemas can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:core/computeImageCapabilitySchema:ComputeImageCapabilitySchema test_compute_image_capability_schema "id"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compartment_id: (Updatable) The OCID of the compartment that contains the resource.
        :param pulumi.Input[str] compute_global_image_capability_schema_version_name: The name of the compute global image capability schema version
        :param pulumi.Input[Mapping[str, Any]] defined_tags: (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        :param pulumi.Input[str] display_name: (Updatable) A user-friendly name for the compute image capability schema
        :param pulumi.Input[Mapping[str, Any]] freeform_tags: (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        :param pulumi.Input[str] image_id: The ocid of the image
        :param pulumi.Input[Mapping[str, Any]] schema_data: (Updatable) The map of each capability name to its ImageCapabilitySchemaDescriptor.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ComputeImageCapabilitySchemaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource provides the Compute Image Capability Schema resource in Oracle Cloud Infrastructure Core service.

        Creates compute image capability schema.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_compute_image_capability_schema = oci.core.ComputeImageCapabilitySchema("testComputeImageCapabilitySchema",
            compartment_id=var["compartment_id"],
            compute_global_image_capability_schema_version_name=var["compute_image_capability_schema_compute_global_image_capability_schema_version_name"],
            image_id=oci_core_image["test_image"]["id"],
            schema_data=[{
                "descriptorType": var["compute_image_capability_schema_schema_data_descriptor_type"],
                "source": var["compute_image_capability_schema_schema_data_source"],
                "defaultValue": var["compute_image_capability_schema_schema_data_default_value"],
                "values": var["compute_image_capability_schema_schema_data_values"],
            }],
            defined_tags={
                "Operations.CostCenter": "42",
            },
            display_name=var["compute_image_capability_schema_display_name"],
            freeform_tags={
                "Department": "Finance",
            })
        ```

        ## Import

        ComputeImageCapabilitySchemas can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:core/computeImageCapabilitySchema:ComputeImageCapabilitySchema test_compute_image_capability_schema "id"
        ```

        :param str resource_name: The name of the resource.
        :param ComputeImageCapabilitySchemaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ComputeImageCapabilitySchemaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compartment_id: Optional[pulumi.Input[str]] = None,
                 compute_global_image_capability_schema_version_name: Optional[pulumi.Input[str]] = None,
                 defined_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 freeform_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 schema_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ComputeImageCapabilitySchemaArgs.__new__(ComputeImageCapabilitySchemaArgs)

            if compartment_id is None and not opts.urn:
                raise TypeError("Missing required property 'compartment_id'")
            __props__.__dict__["compartment_id"] = compartment_id
            if compute_global_image_capability_schema_version_name is None and not opts.urn:
                raise TypeError("Missing required property 'compute_global_image_capability_schema_version_name'")
            __props__.__dict__["compute_global_image_capability_schema_version_name"] = compute_global_image_capability_schema_version_name
            __props__.__dict__["defined_tags"] = defined_tags
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["freeform_tags"] = freeform_tags
            if image_id is None and not opts.urn:
                raise TypeError("Missing required property 'image_id'")
            __props__.__dict__["image_id"] = image_id
            if schema_data is None and not opts.urn:
                raise TypeError("Missing required property 'schema_data'")
            __props__.__dict__["schema_data"] = schema_data
            __props__.__dict__["compute_global_image_capability_schema_id"] = None
            __props__.__dict__["time_created"] = None
        super(ComputeImageCapabilitySchema, __self__).__init__(
            'oci:core/computeImageCapabilitySchema:ComputeImageCapabilitySchema',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compartment_id: Optional[pulumi.Input[str]] = None,
            compute_global_image_capability_schema_id: Optional[pulumi.Input[str]] = None,
            compute_global_image_capability_schema_version_name: Optional[pulumi.Input[str]] = None,
            defined_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            freeform_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            schema_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            time_created: Optional[pulumi.Input[str]] = None) -> 'ComputeImageCapabilitySchema':
        """
        Get an existing ComputeImageCapabilitySchema resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compartment_id: (Updatable) The OCID of the compartment that contains the resource.
        :param pulumi.Input[str] compute_global_image_capability_schema_id: The ocid of the compute global image capability schema
        :param pulumi.Input[str] compute_global_image_capability_schema_version_name: The name of the compute global image capability schema version
        :param pulumi.Input[Mapping[str, Any]] defined_tags: (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        :param pulumi.Input[str] display_name: (Updatable) A user-friendly name for the compute image capability schema
        :param pulumi.Input[Mapping[str, Any]] freeform_tags: (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        :param pulumi.Input[str] image_id: The ocid of the image
        :param pulumi.Input[Mapping[str, Any]] schema_data: (Updatable) The map of each capability name to its ImageCapabilitySchemaDescriptor.
        :param pulumi.Input[str] time_created: The date and time the compute image capability schema was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ComputeImageCapabilitySchemaState.__new__(_ComputeImageCapabilitySchemaState)

        __props__.__dict__["compartment_id"] = compartment_id
        __props__.__dict__["compute_global_image_capability_schema_id"] = compute_global_image_capability_schema_id
        __props__.__dict__["compute_global_image_capability_schema_version_name"] = compute_global_image_capability_schema_version_name
        __props__.__dict__["defined_tags"] = defined_tags
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["freeform_tags"] = freeform_tags
        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["schema_data"] = schema_data
        __props__.__dict__["time_created"] = time_created
        return ComputeImageCapabilitySchema(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> pulumi.Output[str]:
        """
        (Updatable) The OCID of the compartment that contains the resource.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="computeGlobalImageCapabilitySchemaId")
    def compute_global_image_capability_schema_id(self) -> pulumi.Output[str]:
        """
        The ocid of the compute global image capability schema
        """
        return pulumi.get(self, "compute_global_image_capability_schema_id")

    @property
    @pulumi.getter(name="computeGlobalImageCapabilitySchemaVersionName")
    def compute_global_image_capability_schema_version_name(self) -> pulumi.Output[str]:
        """
        The name of the compute global image capability schema version
        """
        return pulumi.get(self, "compute_global_image_capability_schema_version_name")

    @property
    @pulumi.getter(name="definedTags")
    def defined_tags(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        """
        return pulumi.get(self, "defined_tags")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        (Updatable) A user-friendly name for the compute image capability schema
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="freeformTags")
    def freeform_tags(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        """
        return pulumi.get(self, "freeform_tags")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[str]:
        """
        The ocid of the image
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="schemaData")
    def schema_data(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        (Updatable) The map of each capability name to its ImageCapabilitySchemaDescriptor.
        """
        return pulumi.get(self, "schema_data")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> pulumi.Output[str]:
        """
        The date and time the compute image capability schema was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
        """
        return pulumi.get(self, "time_created")
