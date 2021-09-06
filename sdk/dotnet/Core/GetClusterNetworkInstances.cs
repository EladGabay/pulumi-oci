// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Core
{
    public static class GetClusterNetworkInstances
    {
        /// <summary>
        /// This data source provides the list of Cluster Network Instances in Oracle Cloud Infrastructure Core service.
        /// 
        /// Lists the instances in the specified cluster network.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Oci = Pulumi.Oci;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var testClusterNetworkInstances = Output.Create(Oci.Core.GetClusterNetworkInstances.InvokeAsync(new Oci.Core.GetClusterNetworkInstancesArgs
        ///         {
        ///             ClusterNetworkId = oci_core_cluster_network.Test_cluster_network.Id,
        ///             CompartmentId = @var.Compartment_id,
        ///             DisplayName = @var.Cluster_network_instance_display_name,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterNetworkInstancesResult> InvokeAsync(GetClusterNetworkInstancesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterNetworkInstancesResult>("oci:core/getClusterNetworkInstances:getClusterNetworkInstances", args ?? new GetClusterNetworkInstancesArgs(), options.WithVersion());
    }


    public sealed class GetClusterNetworkInstancesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster network.
        /// </summary>
        [Input("clusterNetworkId", required: true)]
        public string ClusterNetworkId { get; set; } = null!;

        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// A filter to return only resources that match the given display name exactly.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("filters")]
        private List<Inputs.GetClusterNetworkInstancesFilterArgs>? _filters;
        public List<Inputs.GetClusterNetworkInstancesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetClusterNetworkInstancesFilterArgs>());
            set => _filters = value;
        }

        public GetClusterNetworkInstancesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterNetworkInstancesResult
    {
        public readonly string ClusterNetworkId;
        /// <summary>
        /// The OCID of the compartment that contains the instance.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// The user-friendly name. Does not have to be unique.
        /// </summary>
        public readonly string? DisplayName;
        public readonly ImmutableArray<Outputs.GetClusterNetworkInstancesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of instances.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterNetworkInstancesInstanceResult> Instances;

        [OutputConstructor]
        private GetClusterNetworkInstancesResult(
            string clusterNetworkId,

            string compartmentId,

            string? displayName,

            ImmutableArray<Outputs.GetClusterNetworkInstancesFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetClusterNetworkInstancesInstanceResult> instances)
        {
            ClusterNetworkId = clusterNetworkId;
            CompartmentId = compartmentId;
            DisplayName = displayName;
            Filters = filters;
            Id = id;
            Instances = instances;
        }
    }
}