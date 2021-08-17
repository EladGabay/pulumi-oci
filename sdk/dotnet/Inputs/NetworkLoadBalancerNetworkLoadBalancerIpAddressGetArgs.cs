// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Inputs
{

    public sealed class NetworkLoadBalancerNetworkLoadBalancerIpAddressGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An IP address.  Example: `192.168.0.3`
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Whether the IP address is public or private.
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        /// <summary>
        /// An object representing a reserved IP address to be attached or that is already attached to a network load balancer.
        /// </summary>
        [Input("reservedIp")]
        public Input<Inputs.NetworkLoadBalancerNetworkLoadBalancerIpAddressReservedIpGetArgs>? ReservedIp { get; set; }

        public NetworkLoadBalancerNetworkLoadBalancerIpAddressGetArgs()
        {
        }
    }
}