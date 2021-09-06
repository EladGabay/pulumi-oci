# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'BackendSetBackendArgs',
    'BackendSetHealthCheckerArgs',
    'NetworkLoadBalancerIpAddressArgs',
    'NetworkLoadBalancerIpAddressReservedIpArgs',
    'NetworkLoadBalancerReservedIpArgs',
    'GetBackendSetsFilterArgs',
    'GetBackendsFilterArgs',
    'GetListenersFilterArgs',
    'GetNetworkLoadBalancersFilterArgs',
    'GetNetworkLoadBalancersPoliciesFilterArgs',
    'GetNetworkLoadBalancersProtocolsFilterArgs',
]

@pulumi.input_type
class BackendSetBackendArgs:
    def __init__(__self__, *,
                 port: pulumi.Input[int],
                 ip_address: Optional[pulumi.Input[str]] = None,
                 is_backup: Optional[pulumi.Input[bool]] = None,
                 is_drain: Optional[pulumi.Input[bool]] = None,
                 is_offline: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] port: (Updatable) The backend server port against which to run the health check. If the port is not specified, then the network load balancer uses the port information from the `Backend` object. The port must be specified if the backend port is 0.  Example: `8080`
        :param pulumi.Input[str] ip_address: The IP address of the backend server. Example: `10.0.0.3`
        :param pulumi.Input[bool] is_backup: Whether the network load balancer should treat this server as a backup unit. If `true`, then the network load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "isBackup" fail the health check policy.  Example: `false`
        :param pulumi.Input[bool] is_drain: Whether the network load balancer should drain this server. Servers marked "isDrain" receive no  incoming traffic.  Example: `false`
        :param pulumi.Input[bool] is_offline: Whether the network load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false`
        :param pulumi.Input[str] name: A user-friendly name for the backend set that must be unique and cannot be changed.
        :param pulumi.Input[str] target_id: The IP OCID/Instance OCID associated with the backend server. Example: `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>`
        :param pulumi.Input[int] weight: The network load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections as a server weighted '1'. For more information about load balancing policies, see [How Network Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).  Example: `3`
        """
        pulumi.set(__self__, "port", port)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if is_backup is not None:
            pulumi.set(__self__, "is_backup", is_backup)
        if is_drain is not None:
            pulumi.set(__self__, "is_drain", is_drain)
        if is_offline is not None:
            pulumi.set(__self__, "is_offline", is_offline)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if target_id is not None:
            pulumi.set(__self__, "target_id", target_id)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        (Updatable) The backend server port against which to run the health check. If the port is not specified, then the network load balancer uses the port information from the `Backend` object. The port must be specified if the backend port is 0.  Example: `8080`
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of the backend server. Example: `10.0.0.3`
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="isBackup")
    def is_backup(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the network load balancer should treat this server as a backup unit. If `true`, then the network load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "isBackup" fail the health check policy.  Example: `false`
        """
        return pulumi.get(self, "is_backup")

    @is_backup.setter
    def is_backup(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_backup", value)

    @property
    @pulumi.getter(name="isDrain")
    def is_drain(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the network load balancer should drain this server. Servers marked "isDrain" receive no  incoming traffic.  Example: `false`
        """
        return pulumi.get(self, "is_drain")

    @is_drain.setter
    def is_drain(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_drain", value)

    @property
    @pulumi.getter(name="isOffline")
    def is_offline(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the network load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false`
        """
        return pulumi.get(self, "is_offline")

    @is_offline.setter
    def is_offline(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_offline", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A user-friendly name for the backend set that must be unique and cannot be changed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> Optional[pulumi.Input[str]]:
        """
        The IP OCID/Instance OCID associated with the backend server. Example: `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>`
        """
        return pulumi.get(self, "target_id")

    @target_id.setter
    def target_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_id", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        The network load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections as a server weighted '1'. For more information about load balancing policies, see [How Network Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).  Example: `3`
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class BackendSetHealthCheckerArgs:
    def __init__(__self__, *,
                 protocol: pulumi.Input[str],
                 interval_in_millis: Optional[pulumi.Input[int]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 request_data: Optional[pulumi.Input[str]] = None,
                 response_body_regex: Optional[pulumi.Input[str]] = None,
                 response_data: Optional[pulumi.Input[str]] = None,
                 retries: Optional[pulumi.Input[int]] = None,
                 return_code: Optional[pulumi.Input[int]] = None,
                 timeout_in_millis: Optional[pulumi.Input[int]] = None,
                 url_path: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] protocol: (Updatable) The protocol the health check must use; either HTTP or HTTPS, or UDP or TCP.  Example: `HTTP`
        :param pulumi.Input[int] interval_in_millis: (Updatable) The interval between health checks, in milliseconds. The default value is 10000 (10 seconds).  Example: `10000`
        :param pulumi.Input[int] port: (Updatable) The backend server port against which to run the health check. If the port is not specified, then the network load balancer uses the port information from the `Backend` object. The port must be specified if the backend port is 0.  Example: `8080`
        :param pulumi.Input[str] request_data: (Updatable) Base64 encoded pattern to be sent as UDP or TCP health check probe.
        :param pulumi.Input[str] response_body_regex: (Updatable) A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$`
        :param pulumi.Input[str] response_data: (Updatable) Base64 encoded pattern to be validated as UDP or TCP health check probe response.
        :param pulumi.Input[int] retries: (Updatable) The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state. The default value is 3.  Example: `3`
        :param pulumi.Input[int] return_code: (Updatable) The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol, then you can use common HTTP status codes such as "200".  Example: `200`
        :param pulumi.Input[int] timeout_in_millis: (Updatable) The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period. The default value is 3000 (3 seconds).  Example: `3000`
        :param pulumi.Input[str] url_path: (Updatable) The path against which to run the health check.  Example: `/healthcheck`
        """
        pulumi.set(__self__, "protocol", protocol)
        if interval_in_millis is not None:
            pulumi.set(__self__, "interval_in_millis", interval_in_millis)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if request_data is not None:
            pulumi.set(__self__, "request_data", request_data)
        if response_body_regex is not None:
            pulumi.set(__self__, "response_body_regex", response_body_regex)
        if response_data is not None:
            pulumi.set(__self__, "response_data", response_data)
        if retries is not None:
            pulumi.set(__self__, "retries", retries)
        if return_code is not None:
            pulumi.set(__self__, "return_code", return_code)
        if timeout_in_millis is not None:
            pulumi.set(__self__, "timeout_in_millis", timeout_in_millis)
        if url_path is not None:
            pulumi.set(__self__, "url_path", url_path)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        (Updatable) The protocol the health check must use; either HTTP or HTTPS, or UDP or TCP.  Example: `HTTP`
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="intervalInMillis")
    def interval_in_millis(self) -> Optional[pulumi.Input[int]]:
        """
        (Updatable) The interval between health checks, in milliseconds. The default value is 10000 (10 seconds).  Example: `10000`
        """
        return pulumi.get(self, "interval_in_millis")

    @interval_in_millis.setter
    def interval_in_millis(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval_in_millis", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        (Updatable) The backend server port against which to run the health check. If the port is not specified, then the network load balancer uses the port information from the `Backend` object. The port must be specified if the backend port is 0.  Example: `8080`
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="requestData")
    def request_data(self) -> Optional[pulumi.Input[str]]:
        """
        (Updatable) Base64 encoded pattern to be sent as UDP or TCP health check probe.
        """
        return pulumi.get(self, "request_data")

    @request_data.setter
    def request_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_data", value)

    @property
    @pulumi.getter(name="responseBodyRegex")
    def response_body_regex(self) -> Optional[pulumi.Input[str]]:
        """
        (Updatable) A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$`
        """
        return pulumi.get(self, "response_body_regex")

    @response_body_regex.setter
    def response_body_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_body_regex", value)

    @property
    @pulumi.getter(name="responseData")
    def response_data(self) -> Optional[pulumi.Input[str]]:
        """
        (Updatable) Base64 encoded pattern to be validated as UDP or TCP health check probe response.
        """
        return pulumi.get(self, "response_data")

    @response_data.setter
    def response_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_data", value)

    @property
    @pulumi.getter
    def retries(self) -> Optional[pulumi.Input[int]]:
        """
        (Updatable) The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state. The default value is 3.  Example: `3`
        """
        return pulumi.get(self, "retries")

    @retries.setter
    def retries(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retries", value)

    @property
    @pulumi.getter(name="returnCode")
    def return_code(self) -> Optional[pulumi.Input[int]]:
        """
        (Updatable) The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol, then you can use common HTTP status codes such as "200".  Example: `200`
        """
        return pulumi.get(self, "return_code")

    @return_code.setter
    def return_code(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "return_code", value)

    @property
    @pulumi.getter(name="timeoutInMillis")
    def timeout_in_millis(self) -> Optional[pulumi.Input[int]]:
        """
        (Updatable) The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period. The default value is 3000 (3 seconds).  Example: `3000`
        """
        return pulumi.get(self, "timeout_in_millis")

    @timeout_in_millis.setter
    def timeout_in_millis(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_in_millis", value)

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> Optional[pulumi.Input[str]]:
        """
        (Updatable) The path against which to run the health check.  Example: `/healthcheck`
        """
        return pulumi.get(self, "url_path")

    @url_path.setter
    def url_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_path", value)


@pulumi.input_type
class NetworkLoadBalancerIpAddressArgs:
    def __init__(__self__, *,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 reserved_ip: Optional[pulumi.Input['NetworkLoadBalancerIpAddressReservedIpArgs']] = None):
        """
        :param pulumi.Input[str] ip_address: An IP address.  Example: `192.168.0.3`
        :param pulumi.Input[bool] is_public: Whether the IP address is public or private.
        :param pulumi.Input['NetworkLoadBalancerIpAddressReservedIpArgs'] reserved_ip: An object representing a reserved IP address to be attached or that is already attached to a network load balancer.
        """
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if is_public is not None:
            pulumi.set(__self__, "is_public", is_public)
        if reserved_ip is not None:
            pulumi.set(__self__, "reserved_ip", reserved_ip)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        An IP address.  Example: `192.168.0.3`
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the IP address is public or private.
        """
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter(name="reservedIp")
    def reserved_ip(self) -> Optional[pulumi.Input['NetworkLoadBalancerIpAddressReservedIpArgs']]:
        """
        An object representing a reserved IP address to be attached or that is already attached to a network load balancer.
        """
        return pulumi.get(self, "reserved_ip")

    @reserved_ip.setter
    def reserved_ip(self, value: Optional[pulumi.Input['NetworkLoadBalancerIpAddressReservedIpArgs']]):
        pulumi.set(self, "reserved_ip", value)


@pulumi.input_type
class NetworkLoadBalancerIpAddressReservedIpArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: OCID of the reserved public IP address created with the virtual cloud network.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        OCID of the reserved public IP address created with the virtual cloud network.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class NetworkLoadBalancerReservedIpArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: OCID of the reserved public IP address created with the virtual cloud network.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        OCID of the reserved public IP address created with the virtual cloud network.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class GetBackendSetsFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str],
                 regex: Optional[bool] = None):
        """
        :param str name: A user-friendly name for the backend set that must be unique and cannot be changed.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)
        if regex is not None:
            pulumi.set(__self__, "regex", regex)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A user-friendly name for the backend set that must be unique and cannot be changed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def regex(self) -> Optional[bool]:
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: Optional[bool]):
        pulumi.set(self, "regex", value)


@pulumi.input_type
class GetBackendsFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str],
                 regex: Optional[bool] = None):
        """
        :param str name: A read-only field showing the IP address/IP OCID and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080`, or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:443` or `10.0.0.3:0`
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)
        if regex is not None:
            pulumi.set(__self__, "regex", regex)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A read-only field showing the IP address/IP OCID and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080`, or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:443` or `10.0.0.3:0`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def regex(self) -> Optional[bool]:
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: Optional[bool]):
        pulumi.set(self, "regex", value)


@pulumi.input_type
class GetListenersFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str],
                 regex: Optional[bool] = None):
        """
        :param str name: A friendly name for the listener. It must be unique and it cannot be changed.  Example: `example_listener`
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)
        if regex is not None:
            pulumi.set(__self__, "regex", regex)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A friendly name for the listener. It must be unique and it cannot be changed.  Example: `example_listener`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def regex(self) -> Optional[bool]:
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: Optional[bool]):
        pulumi.set(self, "regex", value)


@pulumi.input_type
class GetNetworkLoadBalancersFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str],
                 regex: Optional[bool] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)
        if regex is not None:
            pulumi.set(__self__, "regex", regex)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def regex(self) -> Optional[bool]:
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: Optional[bool]):
        pulumi.set(self, "regex", value)


@pulumi.input_type
class GetNetworkLoadBalancersPoliciesFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str],
                 regex: Optional[bool] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)
        if regex is not None:
            pulumi.set(__self__, "regex", regex)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def regex(self) -> Optional[bool]:
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: Optional[bool]):
        pulumi.set(self, "regex", value)


@pulumi.input_type
class GetNetworkLoadBalancersProtocolsFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str],
                 regex: Optional[bool] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)
        if regex is not None:
            pulumi.set(__self__, "regex", regex)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def regex(self) -> Optional[bool]:
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: Optional[bool]):
        pulumi.set(self, "regex", value)

