// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetDataintegrationWorkspace
    {
        /// <summary>
        /// This data source provides details about a specific Workspace resource in Oracle Cloud Infrastructure Data Integration service.
        /// 
        /// Retrieves a Data Integration workspace using the specified identifier.
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
        ///         var testWorkspace = Output.Create(Oci.GetDataintegrationWorkspace.InvokeAsync(new Oci.GetDataintegrationWorkspaceArgs
        ///         {
        ///             WorkspaceId = oci_dataintegration_workspace.Test_workspace.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDataintegrationWorkspaceResult> InvokeAsync(GetDataintegrationWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataintegrationWorkspaceResult>("oci:index/getDataintegrationWorkspace:GetDataintegrationWorkspace", args ?? new GetDataintegrationWorkspaceArgs(), options.WithVersion());
    }


    public sealed class GetDataintegrationWorkspaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The workspace ID.
        /// </summary>
        [Input("workspaceId", required: true)]
        public string WorkspaceId { get; set; } = null!;

        public GetDataintegrationWorkspaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataintegrationWorkspaceResult
    {
        /// <summary>
        /// The OCID of the compartment that contains the workspace.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// A user defined description for the workspace.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A user-friendly display name for the workspace. Does not have to be unique, and can be modified. Avoid entering confidential information.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The IP of the custom DNS.
        /// </summary>
        public readonly string DnsServerIp;
        /// <summary>
        /// The DNS zone of the custom DNS to use to resolve names.
        /// </summary>
        public readonly string DnsServerZone;
        /// <summary>
        /// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// A system-generated and immutable identifier assigned to the workspace upon creation.
        /// </summary>
        public readonly string Id;
        public readonly bool IsForceOperation;
        /// <summary>
        /// Specifies whether the private network connection is enabled or disabled.
        /// </summary>
        public readonly bool IsPrivateNetworkEnabled;
        public readonly int QuiesceTimeout;
        /// <summary>
        /// Lifecycle states for workspaces in Data Integration Service CREATING - The resource is being created and may not be usable until the entire metadata is defined UPDATING - The resource is being updated and may not be usable until all changes are commited DELETING - The resource is being deleted and might require deep cleanup of children. ACTIVE   - The resource is valid and available for access INACTIVE - The resource might be incomplete in its definition or might have been made unavailable for administrative reasons DELETED  - The resource has been deleted and isn't available FAILED   - The resource is in a failed state due to validation or other errors STARTING - The resource is being started and may not be usable until becomes ACTIVE again STOPPING - The resource is in the process of Stopping and may not be usable until it Stops or fails STOPPED  - The resource is in Stopped state due to stop operation.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in failed state.
        /// </summary>
        public readonly string StateMessage;
        /// <summary>
        /// The OCID of the subnet for customer connected databases.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// The date and time the workspace was created, in the timestamp format defined by RFC3339.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The date and time the workspace was updated, in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
        /// </summary>
        public readonly string TimeUpdated;
        /// <summary>
        /// The OCID of the VCN the subnet is in.
        /// </summary>
        public readonly string VcnId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetDataintegrationWorkspaceResult(
            string compartmentId,

            ImmutableDictionary<string, object> definedTags,

            string description,

            string displayName,

            string dnsServerIp,

            string dnsServerZone,

            ImmutableDictionary<string, object> freeformTags,

            string id,

            bool isForceOperation,

            bool isPrivateNetworkEnabled,

            int quiesceTimeout,

            string state,

            string stateMessage,

            string subnetId,

            string timeCreated,

            string timeUpdated,

            string vcnId,

            string workspaceId)
        {
            CompartmentId = compartmentId;
            DefinedTags = definedTags;
            Description = description;
            DisplayName = displayName;
            DnsServerIp = dnsServerIp;
            DnsServerZone = dnsServerZone;
            FreeformTags = freeformTags;
            Id = id;
            IsForceOperation = isForceOperation;
            IsPrivateNetworkEnabled = isPrivateNetworkEnabled;
            QuiesceTimeout = quiesceTimeout;
            State = state;
            StateMessage = stateMessage;
            SubnetId = subnetId;
            TimeCreated = timeCreated;
            TimeUpdated = timeUpdated;
            VcnId = vcnId;
            WorkspaceId = workspaceId;
        }
    }
}