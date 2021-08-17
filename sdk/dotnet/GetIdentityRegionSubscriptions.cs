// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetIdentityRegionSubscriptions
    {
        /// <summary>
        /// This data source provides the list of Region Subscriptions in Oracle Cloud Infrastructure Identity service.
        /// 
        /// Lists the region subscriptions for the specified tenancy.
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
        ///         var testRegionSubscriptions = Output.Create(Oci.GetIdentityRegionSubscriptions.InvokeAsync(new Oci.GetIdentityRegionSubscriptionsArgs
        ///         {
        ///             TenancyId = oci_identity_tenancy.Test_tenancy.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIdentityRegionSubscriptionsResult> InvokeAsync(GetIdentityRegionSubscriptionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIdentityRegionSubscriptionsResult>("oci:index/getIdentityRegionSubscriptions:GetIdentityRegionSubscriptions", args ?? new GetIdentityRegionSubscriptionsArgs(), options.WithVersion());
    }


    public sealed class GetIdentityRegionSubscriptionsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetIdentityRegionSubscriptionsFilterArgs>? _filters;
        public List<Inputs.GetIdentityRegionSubscriptionsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetIdentityRegionSubscriptionsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The OCID of the tenancy.
        /// </summary>
        [Input("tenancyId", required: true)]
        public string TenancyId { get; set; } = null!;

        public GetIdentityRegionSubscriptionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIdentityRegionSubscriptionsResult
    {
        public readonly ImmutableArray<Outputs.GetIdentityRegionSubscriptionsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of region_subscriptions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIdentityRegionSubscriptionsRegionSubscriptionResult> RegionSubscriptions;
        public readonly string TenancyId;

        [OutputConstructor]
        private GetIdentityRegionSubscriptionsResult(
            ImmutableArray<Outputs.GetIdentityRegionSubscriptionsFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetIdentityRegionSubscriptionsRegionSubscriptionResult> regionSubscriptions,

            string tenancyId)
        {
            Filters = filters;
            Id = id;
            RegionSubscriptions = regionSubscriptions;
            TenancyId = tenancyId;
        }
    }
}