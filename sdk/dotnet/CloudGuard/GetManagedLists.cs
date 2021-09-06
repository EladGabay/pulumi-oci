// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.CloudGuard
{
    public static class GetManagedLists
    {
        /// <summary>
        /// This data source provides the list of Managed Lists in Oracle Cloud Infrastructure Cloud Guard service.
        /// 
        /// Returns a list of ListManagedLists.
        /// The ListManagedLists operation returns only the managed lists in `compartmentId` passed.
        /// The list does not include any subcompartments of the compartmentId passed.
        /// 
        /// The parameter `accessLevel` specifies whether to return ManagedLists in only
        /// those compartments for which the requestor has INSPECT permissions on at least one resource directly
        /// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
        /// Principal doesn't have access to even one of the child compartments. This is valid only when
        /// `compartmentIdInSubtree` is set to `true`.
        /// 
        /// The parameter `compartmentIdInSubtree` applies when you perform ListManagedLists on the
        /// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
        /// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
        /// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
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
        ///         var testManagedLists = Output.Create(Oci.CloudGuard.GetManagedLists.InvokeAsync(new Oci.CloudGuard.GetManagedListsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             AccessLevel = @var.Managed_list_access_level,
        ///             CompartmentIdInSubtree = @var.Managed_list_compartment_id_in_subtree,
        ///             DisplayName = @var.Managed_list_display_name,
        ///             ListType = @var.Managed_list_list_type,
        ///             ResourceMetadataOnly = @var.Managed_list_resource_metadata_only,
        ///             State = @var.Managed_list_state,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetManagedListsResult> InvokeAsync(GetManagedListsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagedListsResult>("oci:cloudguard/getManagedLists:getManagedLists", args ?? new GetManagedListsArgs(), options.WithVersion());
    }


    public sealed class GetManagedListsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`. Setting this to `ACCESSIBLE` returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to `RESTRICTED` permissions are checked and no partial results are displayed.
        /// </summary>
        [Input("accessLevel")]
        public string? AccessLevel { get; set; }

        /// <summary>
        /// The ID of the compartment in which to list resources.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.
        /// </summary>
        [Input("compartmentIdInSubtree")]
        public bool? CompartmentIdInSubtree { get; set; }

        /// <summary>
        /// A filter to return only resources that match the entire display name given.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("filters")]
        private List<Inputs.GetManagedListsFilterArgs>? _filters;
        public List<Inputs.GetManagedListsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetManagedListsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The type of the ManagedList.
        /// </summary>
        [Input("listType")]
        public string? ListType { get; set; }

        /// <summary>
        /// Default is false. When set to true, the list of all Oracle Managed Resources Metadata supported by Cloud Guard are returned.
        /// </summary>
        [Input("resourceMetadataOnly")]
        public bool? ResourceMetadataOnly { get; set; }

        /// <summary>
        /// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetManagedListsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetManagedListsResult
    {
        public readonly string? AccessLevel;
        /// <summary>
        /// Compartment Identifier where the resource is created
        /// </summary>
        public readonly string CompartmentId;
        public readonly bool? CompartmentIdInSubtree;
        /// <summary>
        /// ManagedList display name
        /// </summary>
        public readonly string? DisplayName;
        public readonly ImmutableArray<Outputs.GetManagedListsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// type of the list
        /// </summary>
        public readonly string? ListType;
        /// <summary>
        /// The list of managed_list_collection.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetManagedListsManagedListCollectionResult> ManagedListCollections;
        public readonly bool? ResourceMetadataOnly;
        /// <summary>
        /// The current state of the resource.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private GetManagedListsResult(
            string? accessLevel,

            string compartmentId,

            bool? compartmentIdInSubtree,

            string? displayName,

            ImmutableArray<Outputs.GetManagedListsFilterResult> filters,

            string id,

            string? listType,

            ImmutableArray<Outputs.GetManagedListsManagedListCollectionResult> managedListCollections,

            bool? resourceMetadataOnly,

            string? state)
        {
            AccessLevel = accessLevel;
            CompartmentId = compartmentId;
            CompartmentIdInSubtree = compartmentIdInSubtree;
            DisplayName = displayName;
            Filters = filters;
            Id = id;
            ListType = listType;
            ManagedListCollections = managedListCollections;
            ResourceMetadataOnly = resourceMetadataOnly;
            State = state;
        }
    }
}