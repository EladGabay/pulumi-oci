// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.DatabaseMigration
{
    public static class GetConnections
    {
        /// <summary>
        /// This data source provides the list of Connections in Oracle Cloud Infrastructure Database Migration service.
        /// 
        /// List all Database Connections.
        /// 
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
        ///         var testConnections = Output.Create(Oci.DatabaseMigration.GetConnections.InvokeAsync(new Oci.DatabaseMigration.GetConnectionsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             DisplayName = @var.Connection_display_name,
        ///             State = @var.Connection_state,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetConnectionsResult> InvokeAsync(GetConnectionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectionsResult>("oci:databasemigration/getConnections:getConnections", args ?? new GetConnectionsArgs(), options.WithVersion());
    }


    public sealed class GetConnectionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the compartment in which to list resources.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// A filter to return only resources that match the entire display name given.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("filters")]
        private List<Inputs.GetConnectionsFilterArgs>? _filters;
        public List<Inputs.GetConnectionsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetConnectionsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The current state of the Database Migration Deployment.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetConnectionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectionsResult
    {
        /// <summary>
        /// OCID of the compartment where the secret containing the credentials will be created.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// The list of connection_collection.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetConnectionsConnectionCollectionResult> ConnectionCollections;
        /// <summary>
        /// Database Connection display name identifier.
        /// </summary>
        public readonly string? DisplayName;
        public readonly ImmutableArray<Outputs.GetConnectionsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The current state of the Connection resource.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private GetConnectionsResult(
            string compartmentId,

            ImmutableArray<Outputs.GetConnectionsConnectionCollectionResult> connectionCollections,

            string? displayName,

            ImmutableArray<Outputs.GetConnectionsFilterResult> filters,

            string id,

            string? state)
        {
            CompartmentId = compartmentId;
            ConnectionCollections = connectionCollections;
            DisplayName = displayName;
            Filters = filters;
            Id = id;
            State = state;
        }
    }
}