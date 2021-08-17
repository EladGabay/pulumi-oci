# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DataSafeDataSafeConfigurationArgs', 'DataSafeDataSafeConfiguration']

@pulumi.input_type
class DataSafeDataSafeConfigurationArgs:
    def __init__(__self__, *,
                 is_enabled: pulumi.Input[bool],
                 compartment_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataSafeDataSafeConfiguration resource.
        :param pulumi.Input[bool] is_enabled: (Updatable) Indicates if Data Safe is enabled.
        :param pulumi.Input[str] compartment_id: (Updatable) A filter to return only resources that match the specified compartment OCID.
        """
        pulumi.set(__self__, "is_enabled", is_enabled)
        if compartment_id is not None:
            pulumi.set(__self__, "compartment_id", compartment_id)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Input[bool]:
        """
        (Updatable) Indicates if Data Safe is enabled.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> Optional[pulumi.Input[str]]:
        """
        (Updatable) A filter to return only resources that match the specified compartment OCID.
        """
        return pulumi.get(self, "compartment_id")

    @compartment_id.setter
    def compartment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compartment_id", value)


@pulumi.input_type
class _DataSafeDataSafeConfigurationState:
    def __init__(__self__, *,
                 compartment_id: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 time_enabled: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DataSafeDataSafeConfiguration resources.
        :param pulumi.Input[str] compartment_id: (Updatable) A filter to return only resources that match the specified compartment OCID.
        :param pulumi.Input[bool] is_enabled: (Updatable) Indicates if Data Safe is enabled.
        :param pulumi.Input[str] state: The current state of Data Safe.
        :param pulumi.Input[str] time_enabled: The date and time Data Safe was enabled, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
        :param pulumi.Input[str] url: The URL of the Data Safe service.
        """
        if compartment_id is not None:
            pulumi.set(__self__, "compartment_id", compartment_id)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if time_enabled is not None:
            pulumi.set(__self__, "time_enabled", time_enabled)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> Optional[pulumi.Input[str]]:
        """
        (Updatable) A filter to return only resources that match the specified compartment OCID.
        """
        return pulumi.get(self, "compartment_id")

    @compartment_id.setter
    def compartment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compartment_id", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        (Updatable) Indicates if Data Safe is enabled.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The current state of Data Safe.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="timeEnabled")
    def time_enabled(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time Data Safe was enabled, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
        """
        return pulumi.get(self, "time_enabled")

    @time_enabled.setter
    def time_enabled(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_enabled", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the Data Safe service.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class DataSafeDataSafeConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compartment_id: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        This resource provides the Data Safe Configuration resource in Oracle Cloud Infrastructure Data Safe service.

        Enables Data Safe in the tenancy and region.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_data_safe_configuration = oci.DataSafeDataSafeConfiguration("testDataSafeConfiguration",
            is_enabled=var["data_safe_configuration_is_enabled"],
            compartment_id=var["compartment_id"])
        ```

        ## Import

        Import is not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compartment_id: (Updatable) A filter to return only resources that match the specified compartment OCID.
        :param pulumi.Input[bool] is_enabled: (Updatable) Indicates if Data Safe is enabled.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataSafeDataSafeConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource provides the Data Safe Configuration resource in Oracle Cloud Infrastructure Data Safe service.

        Enables Data Safe in the tenancy and region.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_data_safe_configuration = oci.DataSafeDataSafeConfiguration("testDataSafeConfiguration",
            is_enabled=var["data_safe_configuration_is_enabled"],
            compartment_id=var["compartment_id"])
        ```

        ## Import

        Import is not supported for this resource.

        :param str resource_name: The name of the resource.
        :param DataSafeDataSafeConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataSafeDataSafeConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compartment_id: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
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
            __props__ = DataSafeDataSafeConfigurationArgs.__new__(DataSafeDataSafeConfigurationArgs)

            __props__.__dict__["compartment_id"] = compartment_id
            if is_enabled is None and not opts.urn:
                raise TypeError("Missing required property 'is_enabled'")
            __props__.__dict__["is_enabled"] = is_enabled
            __props__.__dict__["state"] = None
            __props__.__dict__["time_enabled"] = None
            __props__.__dict__["url"] = None
        super(DataSafeDataSafeConfiguration, __self__).__init__(
            'oci:index/dataSafeDataSafeConfiguration:DataSafeDataSafeConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compartment_id: Optional[pulumi.Input[str]] = None,
            is_enabled: Optional[pulumi.Input[bool]] = None,
            state: Optional[pulumi.Input[str]] = None,
            time_enabled: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'DataSafeDataSafeConfiguration':
        """
        Get an existing DataSafeDataSafeConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compartment_id: (Updatable) A filter to return only resources that match the specified compartment OCID.
        :param pulumi.Input[bool] is_enabled: (Updatable) Indicates if Data Safe is enabled.
        :param pulumi.Input[str] state: The current state of Data Safe.
        :param pulumi.Input[str] time_enabled: The date and time Data Safe was enabled, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
        :param pulumi.Input[str] url: The URL of the Data Safe service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataSafeDataSafeConfigurationState.__new__(_DataSafeDataSafeConfigurationState)

        __props__.__dict__["compartment_id"] = compartment_id
        __props__.__dict__["is_enabled"] = is_enabled
        __props__.__dict__["state"] = state
        __props__.__dict__["time_enabled"] = time_enabled
        __props__.__dict__["url"] = url
        return DataSafeDataSafeConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> pulumi.Output[str]:
        """
        (Updatable) A filter to return only resources that match the specified compartment OCID.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Output[bool]:
        """
        (Updatable) Indicates if Data Safe is enabled.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of Data Safe.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="timeEnabled")
    def time_enabled(self) -> pulumi.Output[str]:
        """
        The date and time Data Safe was enabled, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
        """
        return pulumi.get(self, "time_enabled")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the Data Safe service.
        """
        return pulumi.get(self, "url")
