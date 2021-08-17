// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetCoreFastConnectProviderServices
    {
        /// <summary>
        /// This data source provides the list of Fast Connect Provider Services in Oracle Cloud Infrastructure Core service.
        /// 
        /// Lists the service offerings from supported providers. You need this
        /// information so you can specify your desired provider and service
        /// offering when you create a virtual circuit.
        /// 
        /// For the compartment ID, provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of your tenancy (the root compartment).
        /// 
        /// For more information, see [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).
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
        ///         var testFastConnectProviderServices = Output.Create(Oci.GetCoreFastConnectProviderServices.InvokeAsync(new Oci.GetCoreFastConnectProviderServicesArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCoreFastConnectProviderServicesResult> InvokeAsync(GetCoreFastConnectProviderServicesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCoreFastConnectProviderServicesResult>("oci:index/getCoreFastConnectProviderServices:GetCoreFastConnectProviderServices", args ?? new GetCoreFastConnectProviderServicesArgs(), options.WithVersion());
    }


    public sealed class GetCoreFastConnectProviderServicesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetCoreFastConnectProviderServicesFilterArgs>? _filters;
        public List<Inputs.GetCoreFastConnectProviderServicesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetCoreFastConnectProviderServicesFilterArgs>());
            set => _filters = value;
        }

        public GetCoreFastConnectProviderServicesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCoreFastConnectProviderServicesResult
    {
        public readonly string CompartmentId;
        /// <summary>
        /// The list of fast_connect_provider_services.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCoreFastConnectProviderServicesFastConnectProviderServiceResult> FastConnectProviderServices;
        public readonly ImmutableArray<Outputs.GetCoreFastConnectProviderServicesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetCoreFastConnectProviderServicesResult(
            string compartmentId,

            ImmutableArray<Outputs.GetCoreFastConnectProviderServicesFastConnectProviderServiceResult> fastConnectProviderServices,

            ImmutableArray<Outputs.GetCoreFastConnectProviderServicesFilterResult> filters,

            string id)
        {
            CompartmentId = compartmentId;
            FastConnectProviderServices = fastConnectProviderServices;
            Filters = filters;
            Id = id;
        }
    }
}