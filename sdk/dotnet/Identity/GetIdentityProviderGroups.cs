// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Identity
{
    public static class GetIdentityProviderGroups
    {
        /// <summary>
        /// This data source provides the list of Identity Provider Groups in Oracle Cloud Infrastructure Identity service.
        /// 
        /// Lists the identity provider groups.
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
        ///         var testIdentityProviderGroups = Output.Create(Oci.Identity.GetIdentityProviderGroups.InvokeAsync(new Oci.Identity.GetIdentityProviderGroupsArgs
        ///         {
        ///             IdentityProviderId = oci_identity_identity_provider.Test_identity_provider.Id,
        ///             Name = @var.Identity_provider_group_name,
        ///             State = @var.Identity_provider_group_state,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIdentityProviderGroupsResult> InvokeAsync(GetIdentityProviderGroupsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIdentityProviderGroupsResult>("oci:identity/getIdentityProviderGroups:getIdentityProviderGroups", args ?? new GetIdentityProviderGroupsArgs(), options.WithVersion());
    }


    public sealed class GetIdentityProviderGroupsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetIdentityProviderGroupsFilterArgs>? _filters;
        public List<Inputs.GetIdentityProviderGroupsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetIdentityProviderGroupsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The OCID of the identity provider.
        /// </summary>
        [Input("identityProviderId", required: true)]
        public string IdentityProviderId { get; set; } = null!;

        /// <summary>
        /// A filter to only return resources that match the given name exactly.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetIdentityProviderGroupsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIdentityProviderGroupsResult
    {
        public readonly ImmutableArray<Outputs.GetIdentityProviderGroupsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of identity_provider_groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIdentityProviderGroupsIdentityProviderGroupResult> IdentityProviderGroups;
        /// <summary>
        /// The OCID of the `IdentityProvider` this group belongs to.
        /// </summary>
        public readonly string IdentityProviderId;
        /// <summary>
        /// Display name of the group
        /// </summary>
        public readonly string? Name;
        public readonly string? State;

        [OutputConstructor]
        private GetIdentityProviderGroupsResult(
            ImmutableArray<Outputs.GetIdentityProviderGroupsFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetIdentityProviderGroupsIdentityProviderGroupResult> identityProviderGroups,

            string identityProviderId,

            string? name,

            string? state)
        {
            Filters = filters;
            Id = id;
            IdentityProviderGroups = identityProviderGroups;
            IdentityProviderId = identityProviderId;
            Name = name;
            State = state;
        }
    }
}