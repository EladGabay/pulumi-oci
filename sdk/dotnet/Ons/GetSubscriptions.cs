// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Ons
{
    public static class GetSubscriptions
    {
        /// <summary>
        /// This data source provides the list of Subscriptions in Oracle Cloud Infrastructure Notifications service.
        /// 
        /// Lists the subscriptions in the specified compartment or topic.
        /// 
        /// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
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
        ///         var testSubscriptions = Output.Create(Oci.Ons.GetSubscriptions.InvokeAsync(new Oci.Ons.GetSubscriptionsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             TopicId = oci_ons_notification_topic.Test_notification_topic.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSubscriptionsResult> InvokeAsync(GetSubscriptionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionsResult>("oci:ons/getSubscriptions:getSubscriptions", args ?? new GetSubscriptionsArgs(), options.WithVersion());
    }


    public sealed class GetSubscriptionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetSubscriptionsFilterArgs>? _filters;
        public List<Inputs.GetSubscriptionsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetSubscriptionsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Return all subscriptions that are subscribed to the given topic OCID. Either this query parameter or the compartmentId query parameter must be set.
        /// </summary>
        [Input("topicId")]
        public string? TopicId { get; set; }

        public GetSubscriptionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSubscriptionsResult
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the subscription.
        /// </summary>
        public readonly string CompartmentId;
        public readonly ImmutableArray<Outputs.GetSubscriptionsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of subscriptions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSubscriptionsSubscriptionResult> Subscriptions;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated topic.
        /// </summary>
        public readonly string? TopicId;

        [OutputConstructor]
        private GetSubscriptionsResult(
            string compartmentId,

            ImmutableArray<Outputs.GetSubscriptionsFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetSubscriptionsSubscriptionResult> subscriptions,

            string? topicId)
        {
            CompartmentId = compartmentId;
            Filters = filters;
            Id = id;
            Subscriptions = subscriptions;
            TopicId = topicId;
        }
    }
}