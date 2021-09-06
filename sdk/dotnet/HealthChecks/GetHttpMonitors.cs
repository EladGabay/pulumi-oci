// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.HealthChecks
{
    public static class GetHttpMonitors
    {
        /// <summary>
        /// This data source provides the list of Http Monitors in Oracle Cloud Infrastructure Health Checks service.
        /// 
        /// Gets a list of HTTP monitors.
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
        ///         var testHttpMonitors = Output.Create(Oci.HealthChecks.GetHttpMonitors.InvokeAsync(new Oci.HealthChecks.GetHttpMonitorsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             DisplayName = @var.Http_monitor_display_name,
        ///             HomeRegion = @var.Http_monitor_home_region,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetHttpMonitorsResult> InvokeAsync(GetHttpMonitorsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHttpMonitorsResult>("oci:healthchecks/getHttpMonitors:getHttpMonitors", args ?? new GetHttpMonitorsArgs(), options.WithVersion());
    }


    public sealed class GetHttpMonitorsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filters results by compartment.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// Filters results that exactly match the `displayName` field.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("filters")]
        private List<Inputs.GetHttpMonitorsFilterArgs>? _filters;
        public List<Inputs.GetHttpMonitorsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetHttpMonitorsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Filters results that match the `homeRegion`.
        /// </summary>
        [Input("homeRegion")]
        public string? HomeRegion { get; set; }

        public GetHttpMonitorsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHttpMonitorsResult
    {
        /// <summary>
        /// The OCID of the compartment.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// A user-friendly and mutable name suitable for display in a user interface.
        /// </summary>
        public readonly string? DisplayName;
        public readonly ImmutableArray<Outputs.GetHttpMonitorsFilterResult> Filters;
        /// <summary>
        /// The region where updates must be made and where results must be fetched from.
        /// </summary>
        public readonly string? HomeRegion;
        /// <summary>
        /// The list of http_monitors.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHttpMonitorsHttpMonitorResult> HttpMonitors;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetHttpMonitorsResult(
            string compartmentId,

            string? displayName,

            ImmutableArray<Outputs.GetHttpMonitorsFilterResult> filters,

            string? homeRegion,

            ImmutableArray<Outputs.GetHttpMonitorsHttpMonitorResult> httpMonitors,

            string id)
        {
            CompartmentId = compartmentId;
            DisplayName = displayName;
            Filters = filters;
            HomeRegion = homeRegion;
            HttpMonitors = httpMonitors;
            Id = id;
        }
    }
}