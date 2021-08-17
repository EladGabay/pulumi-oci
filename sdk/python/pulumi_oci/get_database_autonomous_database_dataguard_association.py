# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetDatabaseAutonomousDatabaseDataguardAssociationResult',
    'AwaitableGetDatabaseAutonomousDatabaseDataguardAssociationResult',
    'get_database_autonomous_database_dataguard_association',
]

@pulumi.output_type
class GetDatabaseAutonomousDatabaseDataguardAssociationResult:
    """
    A collection of values returned by GetDatabaseAutonomousDatabaseDataguardAssociation.
    """
    def __init__(__self__, apply_lag=None, apply_rate=None, autonomous_database_dataguard_association_id=None, autonomous_database_id=None, id=None, lifecycle_details=None, peer_autonomous_database_id=None, peer_autonomous_database_life_cycle_state=None, peer_role=None, protection_mode=None, role=None, state=None, time_created=None, time_last_role_changed=None):
        if apply_lag and not isinstance(apply_lag, str):
            raise TypeError("Expected argument 'apply_lag' to be a str")
        pulumi.set(__self__, "apply_lag", apply_lag)
        if apply_rate and not isinstance(apply_rate, str):
            raise TypeError("Expected argument 'apply_rate' to be a str")
        pulumi.set(__self__, "apply_rate", apply_rate)
        if autonomous_database_dataguard_association_id and not isinstance(autonomous_database_dataguard_association_id, str):
            raise TypeError("Expected argument 'autonomous_database_dataguard_association_id' to be a str")
        pulumi.set(__self__, "autonomous_database_dataguard_association_id", autonomous_database_dataguard_association_id)
        if autonomous_database_id and not isinstance(autonomous_database_id, str):
            raise TypeError("Expected argument 'autonomous_database_id' to be a str")
        pulumi.set(__self__, "autonomous_database_id", autonomous_database_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lifecycle_details and not isinstance(lifecycle_details, str):
            raise TypeError("Expected argument 'lifecycle_details' to be a str")
        pulumi.set(__self__, "lifecycle_details", lifecycle_details)
        if peer_autonomous_database_id and not isinstance(peer_autonomous_database_id, str):
            raise TypeError("Expected argument 'peer_autonomous_database_id' to be a str")
        pulumi.set(__self__, "peer_autonomous_database_id", peer_autonomous_database_id)
        if peer_autonomous_database_life_cycle_state and not isinstance(peer_autonomous_database_life_cycle_state, str):
            raise TypeError("Expected argument 'peer_autonomous_database_life_cycle_state' to be a str")
        pulumi.set(__self__, "peer_autonomous_database_life_cycle_state", peer_autonomous_database_life_cycle_state)
        if peer_role and not isinstance(peer_role, str):
            raise TypeError("Expected argument 'peer_role' to be a str")
        pulumi.set(__self__, "peer_role", peer_role)
        if protection_mode and not isinstance(protection_mode, str):
            raise TypeError("Expected argument 'protection_mode' to be a str")
        pulumi.set(__self__, "protection_mode", protection_mode)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if time_created and not isinstance(time_created, str):
            raise TypeError("Expected argument 'time_created' to be a str")
        pulumi.set(__self__, "time_created", time_created)
        if time_last_role_changed and not isinstance(time_last_role_changed, str):
            raise TypeError("Expected argument 'time_last_role_changed' to be a str")
        pulumi.set(__self__, "time_last_role_changed", time_last_role_changed)

    @property
    @pulumi.getter(name="applyLag")
    def apply_lag(self) -> str:
        """
        The lag time between updates to the primary database and application of the redo data on the standby database, as computed by the reporting database.  Example: `9 seconds`
        """
        return pulumi.get(self, "apply_lag")

    @property
    @pulumi.getter(name="applyRate")
    def apply_rate(self) -> str:
        """
        The rate at which redo logs are synced between the associated databases.  Example: `180 Mb per second`
        """
        return pulumi.get(self, "apply_rate")

    @property
    @pulumi.getter(name="autonomousDatabaseDataguardAssociationId")
    def autonomous_database_dataguard_association_id(self) -> str:
        return pulumi.get(self, "autonomous_database_dataguard_association_id")

    @property
    @pulumi.getter(name="autonomousDatabaseId")
    def autonomous_database_id(self) -> str:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database that has a relationship with the peer Autonomous Database.
        """
        return pulumi.get(self, "autonomous_database_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lifecycleDetails")
    def lifecycle_details(self) -> str:
        """
        Additional information about the current lifecycleState, if available.
        """
        return pulumi.get(self, "lifecycle_details")

    @property
    @pulumi.getter(name="peerAutonomousDatabaseId")
    def peer_autonomous_database_id(self) -> str:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous Database.
        """
        return pulumi.get(self, "peer_autonomous_database_id")

    @property
    @pulumi.getter(name="peerAutonomousDatabaseLifeCycleState")
    def peer_autonomous_database_life_cycle_state(self) -> str:
        """
        The current state of the Autonomous Dataguard.
        """
        return pulumi.get(self, "peer_autonomous_database_life_cycle_state")

    @property
    @pulumi.getter(name="peerRole")
    def peer_role(self) -> str:
        """
        The role of the Autonomous Dataguard enabled Autonomous Container Database.
        """
        return pulumi.get(self, "peer_role")

    @property
    @pulumi.getter(name="protectionMode")
    def protection_mode(self) -> str:
        """
        The protection mode of this Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation.
        """
        return pulumi.get(self, "protection_mode")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        The role of the Autonomous Dataguard enabled Autonomous Container Database.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the Autonomous Dataguard.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        The date and time the Data Guard association was created.
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter(name="timeLastRoleChanged")
    def time_last_role_changed(self) -> str:
        """
        The date and time when the last role change action happened.
        """
        return pulumi.get(self, "time_last_role_changed")


class AwaitableGetDatabaseAutonomousDatabaseDataguardAssociationResult(GetDatabaseAutonomousDatabaseDataguardAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseAutonomousDatabaseDataguardAssociationResult(
            apply_lag=self.apply_lag,
            apply_rate=self.apply_rate,
            autonomous_database_dataguard_association_id=self.autonomous_database_dataguard_association_id,
            autonomous_database_id=self.autonomous_database_id,
            id=self.id,
            lifecycle_details=self.lifecycle_details,
            peer_autonomous_database_id=self.peer_autonomous_database_id,
            peer_autonomous_database_life_cycle_state=self.peer_autonomous_database_life_cycle_state,
            peer_role=self.peer_role,
            protection_mode=self.protection_mode,
            role=self.role,
            state=self.state,
            time_created=self.time_created,
            time_last_role_changed=self.time_last_role_changed)


def get_database_autonomous_database_dataguard_association(autonomous_database_dataguard_association_id: Optional[str] = None,
                                                           autonomous_database_id: Optional[str] = None,
                                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseAutonomousDatabaseDataguardAssociationResult:
    """
    This data source provides details about a specific Autonomous Database Dataguard Association resource in Oracle Cloud Infrastructure Database service.

    Gets an Autonomous Database dataguard assocation for the specified Autonomous Database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_autonomous_database_dataguard_association = oci.get_database_autonomous_database_dataguard_association(autonomous_database_dataguard_association_id=oci_database_autonomous_database_dataguard_association["test_autonomous_database_dataguard_association"]["id"],
        autonomous_database_id=oci_database_autonomous_database["test_autonomous_database"]["id"])
    ```


    :param str autonomous_database_dataguard_association_id: The Autonomous Database Dataguard Association [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
    :param str autonomous_database_id: The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
    """
    __args__ = dict()
    __args__['autonomousDatabaseDataguardAssociationId'] = autonomous_database_dataguard_association_id
    __args__['autonomousDatabaseId'] = autonomous_database_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:index/getDatabaseAutonomousDatabaseDataguardAssociation:GetDatabaseAutonomousDatabaseDataguardAssociation', __args__, opts=opts, typ=GetDatabaseAutonomousDatabaseDataguardAssociationResult).value

    return AwaitableGetDatabaseAutonomousDatabaseDataguardAssociationResult(
        apply_lag=__ret__.apply_lag,
        apply_rate=__ret__.apply_rate,
        autonomous_database_dataguard_association_id=__ret__.autonomous_database_dataguard_association_id,
        autonomous_database_id=__ret__.autonomous_database_id,
        id=__ret__.id,
        lifecycle_details=__ret__.lifecycle_details,
        peer_autonomous_database_id=__ret__.peer_autonomous_database_id,
        peer_autonomous_database_life_cycle_state=__ret__.peer_autonomous_database_life_cycle_state,
        peer_role=__ret__.peer_role,
        protection_mode=__ret__.protection_mode,
        role=__ret__.role,
        state=__ret__.state,
        time_created=__ret__.time_created,
        time_last_role_changed=__ret__.time_last_role_changed)