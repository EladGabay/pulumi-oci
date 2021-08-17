# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetLoadBalancersResult',
    'AwaitableGetLoadBalancersResult',
    'get_load_balancers',
]

@pulumi.output_type
class GetLoadBalancersResult:
    """
    A collection of values returned by GetLoadBalancers.
    """
    def __init__(__self__, compartment_id=None, detail=None, display_name=None, filters=None, id=None, load_balancers=None, state=None):
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if detail and not isinstance(detail, str):
            raise TypeError("Expected argument 'detail' to be a str")
        pulumi.set(__self__, "detail", detail)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if load_balancers and not isinstance(load_balancers, list):
            raise TypeError("Expected argument 'load_balancers' to be a list")
        pulumi.set(__self__, "load_balancers", load_balancers)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> str:
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter
    def detail(self) -> Optional[str]:
        return pulumi.get(self, "detail")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetLoadBalancersFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loadBalancers")
    def load_balancers(self) -> Sequence['outputs.GetLoadBalancersLoadBalancerResult']:
        return pulumi.get(self, "load_balancers")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        return pulumi.get(self, "state")


class AwaitableGetLoadBalancersResult(GetLoadBalancersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadBalancersResult(
            compartment_id=self.compartment_id,
            detail=self.detail,
            display_name=self.display_name,
            filters=self.filters,
            id=self.id,
            load_balancers=self.load_balancers,
            state=self.state)


def get_load_balancers(compartment_id: Optional[str] = None,
                       detail: Optional[str] = None,
                       display_name: Optional[str] = None,
                       filters: Optional[Sequence[pulumi.InputType['GetLoadBalancersFilterArgs']]] = None,
                       state: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadBalancersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['compartmentId'] = compartment_id
    __args__['detail'] = detail
    __args__['displayName'] = display_name
    __args__['filters'] = filters
    __args__['state'] = state
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:index/getLoadBalancers:GetLoadBalancers', __args__, opts=opts, typ=GetLoadBalancersResult).value

    return AwaitableGetLoadBalancersResult(
        compartment_id=__ret__.compartment_id,
        detail=__ret__.detail,
        display_name=__ret__.display_name,
        filters=__ret__.filters,
        id=__ret__.id,
        load_balancers=__ret__.load_balancers,
        state=__ret__.state)