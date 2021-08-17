// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Outputs
{

    [OutputType]
    public sealed class ContainerengineNodePoolNode
    {
        /// <summary>
        /// (Updatable) The availability domain in which to place nodes. Example: `Uocm:PHX-AD-1`
        /// </summary>
        public readonly string? AvailabilityDomain;
        /// <summary>
        /// An error that may be associated with the node.
        /// </summary>
        public readonly Outputs.ContainerengineNodePoolNodeError? Error;
        /// <summary>
        /// The fault domain of this node.
        /// </summary>
        public readonly string? FaultDomain;
        /// <summary>
        /// The OCID of the compute instance backing this node.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// (Updatable) The version of Kubernetes to install on the nodes in the node pool.
        /// </summary>
        public readonly string? KubernetesVersion;
        /// <summary>
        /// Details about the state of the node.
        /// </summary>
        public readonly string? LifecycleDetails;
        /// <summary>
        /// (Updatable) The name of the node pool. Avoid entering confidential information.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The OCID of the node pool to which this node belongs.
        /// </summary>
        public readonly string? NodePoolId;
        /// <summary>
        /// The private IP address of this node.
        /// </summary>
        public readonly string? PrivateIp;
        /// <summary>
        /// The public IP address of this node.
        /// </summary>
        public readonly string? PublicIp;
        /// <summary>
        /// The state of the node.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// (Updatable) The OCID of the subnet in which to place nodes.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private ContainerengineNodePoolNode(
            string? availabilityDomain,

            Outputs.ContainerengineNodePoolNodeError? error,

            string? faultDomain,

            string? id,

            string? kubernetesVersion,

            string? lifecycleDetails,

            string? name,

            string? nodePoolId,

            string? privateIp,

            string? publicIp,

            string? state,

            string? subnetId)
        {
            AvailabilityDomain = availabilityDomain;
            Error = error;
            FaultDomain = faultDomain;
            Id = id;
            KubernetesVersion = kubernetesVersion;
            LifecycleDetails = lifecycleDetails;
            Name = name;
            NodePoolId = nodePoolId;
            PrivateIp = privateIp;
            PublicIp = publicIp;
            State = state;
            SubnetId = subnetId;
        }
    }
}