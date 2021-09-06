// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Marketplace
{
    public static class GetAcceptedAgreements
    {
        /// <summary>
        /// This data source provides the list of Accepted Agreements in Oracle Cloud Infrastructure Marketplace service.
        /// 
        /// Lists the terms of use agreements that have been accepted in the specified compartment.
        /// You can filter results by specifying query parameters.
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
        ///         var testAcceptedAgreements = Output.Create(Oci.Marketplace.GetAcceptedAgreements.InvokeAsync(new Oci.Marketplace.GetAcceptedAgreementsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             AcceptedAgreementId = oci_marketplace_accepted_agreement.Test_accepted_agreement.Id,
        ///             DisplayName = @var.Accepted_agreement_display_name,
        ///             ListingId = oci_marketplace_listing.Test_listing.Id,
        ///             PackageVersion = @var.Accepted_agreement_package_version,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAcceptedAgreementsResult> InvokeAsync(GetAcceptedAgreementsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAcceptedAgreementsResult>("oci:marketplace/getAcceptedAgreements:getAcceptedAgreements", args ?? new GetAcceptedAgreementsArgs(), options.WithVersion());
    }


    public sealed class GetAcceptedAgreementsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier for the accepted terms of use agreement.
        /// </summary>
        [Input("acceptedAgreementId")]
        public string? AcceptedAgreementId { get; set; }

        /// <summary>
        /// The unique identifier for the compartment.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// The display name of the resource.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("filters")]
        private List<Inputs.GetAcceptedAgreementsFilterArgs>? _filters;
        public List<Inputs.GetAcceptedAgreementsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetAcceptedAgreementsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The unique identifier for the listing.
        /// </summary>
        [Input("listingId")]
        public string? ListingId { get; set; }

        /// <summary>
        /// The version of the package. Package versions are unique within a listing.
        /// </summary>
        [Input("packageVersion")]
        public string? PackageVersion { get; set; }

        public GetAcceptedAgreementsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAcceptedAgreementsResult
    {
        public readonly string? AcceptedAgreementId;
        /// <summary>
        /// The list of accepted_agreements.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAcceptedAgreementsAcceptedAgreementResult> AcceptedAgreements;
        /// <summary>
        /// The unique identifier for the compartment where the agreement was accepted.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// A display name for the accepted agreement.
        /// </summary>
        public readonly string? DisplayName;
        public readonly ImmutableArray<Outputs.GetAcceptedAgreementsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The unique identifier for the listing associated with the agreement.
        /// </summary>
        public readonly string? ListingId;
        /// <summary>
        /// The package version associated with the agreement.
        /// </summary>
        public readonly string? PackageVersion;

        [OutputConstructor]
        private GetAcceptedAgreementsResult(
            string? acceptedAgreementId,

            ImmutableArray<Outputs.GetAcceptedAgreementsAcceptedAgreementResult> acceptedAgreements,

            string compartmentId,

            string? displayName,

            ImmutableArray<Outputs.GetAcceptedAgreementsFilterResult> filters,

            string id,

            string? listingId,

            string? packageVersion)
        {
            AcceptedAgreementId = acceptedAgreementId;
            AcceptedAgreements = acceptedAgreements;
            CompartmentId = compartmentId;
            DisplayName = displayName;
            Filters = filters;
            Id = id;
            ListingId = listingId;
            PackageVersion = packageVersion;
        }
    }
}