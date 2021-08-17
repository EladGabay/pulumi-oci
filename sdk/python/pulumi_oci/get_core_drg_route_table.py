# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetCoreDrgRouteTableResult',
    'AwaitableGetCoreDrgRouteTableResult',
    'get_core_drg_route_table',
]

@pulumi.output_type
class GetCoreDrgRouteTableResult:
    """
    A collection of values returned by GetCoreDrgRouteTable.
    """
    def __init__(__self__, compartment_id=None, defined_tags=None, display_name=None, drg_id=None, drg_route_table_id=None, freeform_tags=None, id=None, import_drg_route_distribution_id=None, is_ecmp_enabled=None, remove_import_trigger=None, state=None, time_created=None):
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if defined_tags and not isinstance(defined_tags, dict):
            raise TypeError("Expected argument 'defined_tags' to be a dict")
        pulumi.set(__self__, "defined_tags", defined_tags)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if drg_id and not isinstance(drg_id, str):
            raise TypeError("Expected argument 'drg_id' to be a str")
        pulumi.set(__self__, "drg_id", drg_id)
        if drg_route_table_id and not isinstance(drg_route_table_id, str):
            raise TypeError("Expected argument 'drg_route_table_id' to be a str")
        pulumi.set(__self__, "drg_route_table_id", drg_route_table_id)
        if freeform_tags and not isinstance(freeform_tags, dict):
            raise TypeError("Expected argument 'freeform_tags' to be a dict")
        pulumi.set(__self__, "freeform_tags", freeform_tags)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if import_drg_route_distribution_id and not isinstance(import_drg_route_distribution_id, str):
            raise TypeError("Expected argument 'import_drg_route_distribution_id' to be a str")
        pulumi.set(__self__, "import_drg_route_distribution_id", import_drg_route_distribution_id)
        if is_ecmp_enabled and not isinstance(is_ecmp_enabled, bool):
            raise TypeError("Expected argument 'is_ecmp_enabled' to be a bool")
        pulumi.set(__self__, "is_ecmp_enabled", is_ecmp_enabled)
        if remove_import_trigger and not isinstance(remove_import_trigger, bool):
            raise TypeError("Expected argument 'remove_import_trigger' to be a bool")
        pulumi.set(__self__, "remove_import_trigger", remove_import_trigger)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if time_created and not isinstance(time_created, str):
            raise TypeError("Expected argument 'time_created' to be a str")
        pulumi.set(__self__, "time_created", time_created)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> str:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the DRG is in. The DRG route table is always in the same compartment as the DRG.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="definedTags")
    def defined_tags(self) -> Mapping[str, Any]:
        """
        Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        """
        return pulumi.get(self, "defined_tags")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="drgId")
    def drg_id(self) -> str:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG the DRG that contains this route table.
        """
        return pulumi.get(self, "drg_id")

    @property
    @pulumi.getter(name="drgRouteTableId")
    def drg_route_table_id(self) -> str:
        return pulumi.get(self, "drg_route_table_id")

    @property
    @pulumi.getter(name="freeformTags")
    def freeform_tags(self) -> Mapping[str, Any]:
        """
        Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        """
        return pulumi.get(self, "freeform_tags")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="importDrgRouteDistributionId")
    def import_drg_route_distribution_id(self) -> str:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution used to specify how incoming route advertisements from referenced attachments are inserted into the DRG route table.
        """
        return pulumi.get(self, "import_drg_route_distribution_id")

    @property
    @pulumi.getter(name="isEcmpEnabled")
    def is_ecmp_enabled(self) -> bool:
        """
        If you want traffic to be routed using ECMP across your virtual circuits or IPSec tunnels to your on-premises network, enable ECMP on the DRG route table to which these attachments import routes.
        """
        return pulumi.get(self, "is_ecmp_enabled")

    @property
    @pulumi.getter(name="removeImportTrigger")
    def remove_import_trigger(self) -> bool:
        return pulumi.get(self, "remove_import_trigger")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The DRG route table's current state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        The date and time the DRG route table was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
        """
        return pulumi.get(self, "time_created")


class AwaitableGetCoreDrgRouteTableResult(GetCoreDrgRouteTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCoreDrgRouteTableResult(
            compartment_id=self.compartment_id,
            defined_tags=self.defined_tags,
            display_name=self.display_name,
            drg_id=self.drg_id,
            drg_route_table_id=self.drg_route_table_id,
            freeform_tags=self.freeform_tags,
            id=self.id,
            import_drg_route_distribution_id=self.import_drg_route_distribution_id,
            is_ecmp_enabled=self.is_ecmp_enabled,
            remove_import_trigger=self.remove_import_trigger,
            state=self.state,
            time_created=self.time_created)


def get_core_drg_route_table(drg_route_table_id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCoreDrgRouteTableResult:
    """
    This data source provides details about a specific Drg Route Table resource in Oracle Cloud Infrastructure Core service.

    Gets the specified DRG route table's information.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_drg_route_table = oci.get_core_drg_route_table(drg_route_table_id=oci_core_drg_route_table["test_drg_route_table"]["id"])
    ```


    :param str drg_route_table_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table.
    """
    __args__ = dict()
    __args__['drgRouteTableId'] = drg_route_table_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:index/getCoreDrgRouteTable:GetCoreDrgRouteTable', __args__, opts=opts, typ=GetCoreDrgRouteTableResult).value

    return AwaitableGetCoreDrgRouteTableResult(
        compartment_id=__ret__.compartment_id,
        defined_tags=__ret__.defined_tags,
        display_name=__ret__.display_name,
        drg_id=__ret__.drg_id,
        drg_route_table_id=__ret__.drg_route_table_id,
        freeform_tags=__ret__.freeform_tags,
        id=__ret__.id,
        import_drg_route_distribution_id=__ret__.import_drg_route_distribution_id,
        is_ecmp_enabled=__ret__.is_ecmp_enabled,
        remove_import_trigger=__ret__.remove_import_trigger,
        state=__ret__.state,
        time_created=__ret__.time_created)