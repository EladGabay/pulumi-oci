# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetBdsInstanceResult',
    'AwaitableGetBdsInstanceResult',
    'get_bds_instance',
]

@pulumi.output_type
class GetBdsInstanceResult:
    """
    A collection of values returned by getBdsInstance.
    """
    def __init__(__self__, bds_instance_id=None, cloud_sql_details=None, cluster_admin_password=None, cluster_details=None, cluster_public_key=None, cluster_version=None, compartment_id=None, created_by=None, defined_tags=None, display_name=None, freeform_tags=None, id=None, is_cloud_sql_configured=None, is_high_availability=None, is_secure=None, master_node=None, network_config=None, nodes=None, number_of_nodes=None, state=None, time_created=None, time_updated=None, util_node=None, worker_node=None):
        if bds_instance_id and not isinstance(bds_instance_id, str):
            raise TypeError("Expected argument 'bds_instance_id' to be a str")
        pulumi.set(__self__, "bds_instance_id", bds_instance_id)
        if cloud_sql_details and not isinstance(cloud_sql_details, dict):
            raise TypeError("Expected argument 'cloud_sql_details' to be a dict")
        pulumi.set(__self__, "cloud_sql_details", cloud_sql_details)
        if cluster_admin_password and not isinstance(cluster_admin_password, str):
            raise TypeError("Expected argument 'cluster_admin_password' to be a str")
        pulumi.set(__self__, "cluster_admin_password", cluster_admin_password)
        if cluster_details and not isinstance(cluster_details, dict):
            raise TypeError("Expected argument 'cluster_details' to be a dict")
        pulumi.set(__self__, "cluster_details", cluster_details)
        if cluster_public_key and not isinstance(cluster_public_key, str):
            raise TypeError("Expected argument 'cluster_public_key' to be a str")
        pulumi.set(__self__, "cluster_public_key", cluster_public_key)
        if cluster_version and not isinstance(cluster_version, str):
            raise TypeError("Expected argument 'cluster_version' to be a str")
        pulumi.set(__self__, "cluster_version", cluster_version)
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if created_by and not isinstance(created_by, str):
            raise TypeError("Expected argument 'created_by' to be a str")
        pulumi.set(__self__, "created_by", created_by)
        if defined_tags and not isinstance(defined_tags, dict):
            raise TypeError("Expected argument 'defined_tags' to be a dict")
        pulumi.set(__self__, "defined_tags", defined_tags)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if freeform_tags and not isinstance(freeform_tags, dict):
            raise TypeError("Expected argument 'freeform_tags' to be a dict")
        pulumi.set(__self__, "freeform_tags", freeform_tags)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_cloud_sql_configured and not isinstance(is_cloud_sql_configured, bool):
            raise TypeError("Expected argument 'is_cloud_sql_configured' to be a bool")
        pulumi.set(__self__, "is_cloud_sql_configured", is_cloud_sql_configured)
        if is_high_availability and not isinstance(is_high_availability, bool):
            raise TypeError("Expected argument 'is_high_availability' to be a bool")
        pulumi.set(__self__, "is_high_availability", is_high_availability)
        if is_secure and not isinstance(is_secure, bool):
            raise TypeError("Expected argument 'is_secure' to be a bool")
        pulumi.set(__self__, "is_secure", is_secure)
        if master_node and not isinstance(master_node, dict):
            raise TypeError("Expected argument 'master_node' to be a dict")
        pulumi.set(__self__, "master_node", master_node)
        if network_config and not isinstance(network_config, dict):
            raise TypeError("Expected argument 'network_config' to be a dict")
        pulumi.set(__self__, "network_config", network_config)
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        pulumi.set(__self__, "nodes", nodes)
        if number_of_nodes and not isinstance(number_of_nodes, int):
            raise TypeError("Expected argument 'number_of_nodes' to be a int")
        pulumi.set(__self__, "number_of_nodes", number_of_nodes)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if time_created and not isinstance(time_created, str):
            raise TypeError("Expected argument 'time_created' to be a str")
        pulumi.set(__self__, "time_created", time_created)
        if time_updated and not isinstance(time_updated, str):
            raise TypeError("Expected argument 'time_updated' to be a str")
        pulumi.set(__self__, "time_updated", time_updated)
        if util_node and not isinstance(util_node, dict):
            raise TypeError("Expected argument 'util_node' to be a dict")
        pulumi.set(__self__, "util_node", util_node)
        if worker_node and not isinstance(worker_node, dict):
            raise TypeError("Expected argument 'worker_node' to be a dict")
        pulumi.set(__self__, "worker_node", worker_node)

    @property
    @pulumi.getter(name="bdsInstanceId")
    def bds_instance_id(self) -> str:
        return pulumi.get(self, "bds_instance_id")

    @property
    @pulumi.getter(name="cloudSqlDetails")
    def cloud_sql_details(self) -> 'outputs.GetBdsInstanceCloudSqlDetailsResult':
        """
        The information about added Cloud SQL capability
        """
        return pulumi.get(self, "cloud_sql_details")

    @property
    @pulumi.getter(name="clusterAdminPassword")
    def cluster_admin_password(self) -> str:
        return pulumi.get(self, "cluster_admin_password")

    @property
    @pulumi.getter(name="clusterDetails")
    def cluster_details(self) -> 'outputs.GetBdsInstanceClusterDetailsResult':
        """
        Specific info about a Hadoop cluster
        """
        return pulumi.get(self, "cluster_details")

    @property
    @pulumi.getter(name="clusterPublicKey")
    def cluster_public_key(self) -> str:
        return pulumi.get(self, "cluster_public_key")

    @property
    @pulumi.getter(name="clusterVersion")
    def cluster_version(self) -> str:
        """
        Version of the Hadoop distribution.
        """
        return pulumi.get(self, "cluster_version")

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> str:
        """
        The OCID of the compartment.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> str:
        """
        The user who created the cluster.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="definedTags")
    def defined_tags(self) -> Mapping[str, Any]:
        """
        Defined tags for this resource. Each key is predefined and scoped to a namespace. For example, `{"foo-namespace": {"bar-key": "value"}}`
        """
        return pulumi.get(self, "defined_tags")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The name of the node.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="freeformTags")
    def freeform_tags(self) -> Mapping[str, Any]:
        """
        Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. For example, `{"bar-key": "value"}`
        """
        return pulumi.get(self, "freeform_tags")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The OCID of the Big Data Service resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isCloudSqlConfigured")
    def is_cloud_sql_configured(self) -> bool:
        """
        Boolean flag specifying whether or not Cloud SQL should be configured.
        """
        return pulumi.get(self, "is_cloud_sql_configured")

    @property
    @pulumi.getter(name="isHighAvailability")
    def is_high_availability(self) -> bool:
        """
        Boolean flag specifying whether or not the cluster is highly available (HA)
        """
        return pulumi.get(self, "is_high_availability")

    @property
    @pulumi.getter(name="isSecure")
    def is_secure(self) -> bool:
        """
        Boolean flag specifying whether or not the cluster should be set up as secure.
        """
        return pulumi.get(self, "is_secure")

    @property
    @pulumi.getter(name="masterNode")
    def master_node(self) -> 'outputs.GetBdsInstanceMasterNodeResult':
        return pulumi.get(self, "master_node")

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> 'outputs.GetBdsInstanceNetworkConfigResult':
        """
        Additional configuration of the user's network.
        """
        return pulumi.get(self, "network_config")

    @property
    @pulumi.getter
    def nodes(self) -> Sequence['outputs.GetBdsInstanceNodeResult']:
        """
        The list of nodes in the cluster.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter(name="numberOfNodes")
    def number_of_nodes(self) -> int:
        """
        The number of nodes that form the cluster.
        """
        return pulumi.get(self, "number_of_nodes")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the cluster.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        The time the cluster was created, shown as an RFC 3339 formatted datetime string.
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter(name="timeUpdated")
    def time_updated(self) -> str:
        """
        The time the cluster was updated, shown as an RFC 3339 formatted datetime string.
        """
        return pulumi.get(self, "time_updated")

    @property
    @pulumi.getter(name="utilNode")
    def util_node(self) -> 'outputs.GetBdsInstanceUtilNodeResult':
        return pulumi.get(self, "util_node")

    @property
    @pulumi.getter(name="workerNode")
    def worker_node(self) -> 'outputs.GetBdsInstanceWorkerNodeResult':
        return pulumi.get(self, "worker_node")


class AwaitableGetBdsInstanceResult(GetBdsInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBdsInstanceResult(
            bds_instance_id=self.bds_instance_id,
            cloud_sql_details=self.cloud_sql_details,
            cluster_admin_password=self.cluster_admin_password,
            cluster_details=self.cluster_details,
            cluster_public_key=self.cluster_public_key,
            cluster_version=self.cluster_version,
            compartment_id=self.compartment_id,
            created_by=self.created_by,
            defined_tags=self.defined_tags,
            display_name=self.display_name,
            freeform_tags=self.freeform_tags,
            id=self.id,
            is_cloud_sql_configured=self.is_cloud_sql_configured,
            is_high_availability=self.is_high_availability,
            is_secure=self.is_secure,
            master_node=self.master_node,
            network_config=self.network_config,
            nodes=self.nodes,
            number_of_nodes=self.number_of_nodes,
            state=self.state,
            time_created=self.time_created,
            time_updated=self.time_updated,
            util_node=self.util_node,
            worker_node=self.worker_node)


def get_bds_instance(bds_instance_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBdsInstanceResult:
    """
    This data source provides details about a specific Bds Instance resource in Oracle Cloud Infrastructure Big Data Service service.

    Returns information about the Big Data Service cluster identified by the given ID.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_bds_instance = oci.bds.get_bds_instance(bds_instance_id=oci_bds_bds_instance["test_bds_instance"]["id"])
    ```


    :param str bds_instance_id: The OCID of the cluster.
    """
    __args__ = dict()
    __args__['bdsInstanceId'] = bds_instance_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:bds/getBdsInstance:getBdsInstance', __args__, opts=opts, typ=GetBdsInstanceResult).value

    return AwaitableGetBdsInstanceResult(
        bds_instance_id=__ret__.bds_instance_id,
        cloud_sql_details=__ret__.cloud_sql_details,
        cluster_admin_password=__ret__.cluster_admin_password,
        cluster_details=__ret__.cluster_details,
        cluster_public_key=__ret__.cluster_public_key,
        cluster_version=__ret__.cluster_version,
        compartment_id=__ret__.compartment_id,
        created_by=__ret__.created_by,
        defined_tags=__ret__.defined_tags,
        display_name=__ret__.display_name,
        freeform_tags=__ret__.freeform_tags,
        id=__ret__.id,
        is_cloud_sql_configured=__ret__.is_cloud_sql_configured,
        is_high_availability=__ret__.is_high_availability,
        is_secure=__ret__.is_secure,
        master_node=__ret__.master_node,
        network_config=__ret__.network_config,
        nodes=__ret__.nodes,
        number_of_nodes=__ret__.number_of_nodes,
        state=__ret__.state,
        time_created=__ret__.time_created,
        time_updated=__ret__.time_updated,
        util_node=__ret__.util_node,
        worker_node=__ret__.worker_node)