// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetDataflowInvokeRuns
    {
        /// <summary>
        /// This data source provides the list of Invoke Runs in Oracle Cloud Infrastructure Data Flow service.
        /// 
        /// Lists all runs of an application in the specified compartment.  Only one parameter other than compartmentId may also be included in a query. The query must include compartmentId. If the query does not include compartmentId, or includes compartmentId but two or more other parameters an error is returned.
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
        ///         var testInvokeRuns = Output.Create(Oci.GetDataflowInvokeRuns.InvokeAsync(new Oci.GetDataflowInvokeRunsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             ApplicationId = oci_dataflow_application.Test_application.Id,
        ///             DisplayName = @var.Invoke_run_display_name,
        ///             DisplayNameStartsWith = @var.Invoke_run_display_name_starts_with,
        ///             OwnerPrincipalId = oci_dataflow_owner_principal.Test_owner_principal.Id,
        ///             State = @var.Invoke_run_state,
        ///             TimeCreatedGreaterThan = @var.Invoke_run_time_created_greater_than,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDataflowInvokeRunsResult> InvokeAsync(GetDataflowInvokeRunsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataflowInvokeRunsResult>("oci:index/getDataflowInvokeRuns:GetDataflowInvokeRuns", args ?? new GetDataflowInvokeRunsArgs(), options.WithVersion());
    }


    public sealed class GetDataflowInvokeRunsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the application.
        /// </summary>
        [Input("applicationId")]
        public string? ApplicationId { get; set; }

        /// <summary>
        /// The OCID of the compartment.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// The query parameter for the Spark application name.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// The displayName prefix.
        /// </summary>
        [Input("displayNameStartsWith")]
        public string? DisplayNameStartsWith { get; set; }

        [Input("filters")]
        private List<Inputs.GetDataflowInvokeRunsFilterArgs>? _filters;
        public List<Inputs.GetDataflowInvokeRunsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDataflowInvokeRunsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The OCID of the user who created the resource.
        /// </summary>
        [Input("ownerPrincipalId")]
        public string? OwnerPrincipalId { get; set; }

        /// <summary>
        /// The LifecycleState of the run.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        /// <summary>
        /// The epoch time that the resource was created.
        /// </summary>
        [Input("timeCreatedGreaterThan")]
        public string? TimeCreatedGreaterThan { get; set; }

        public GetDataflowInvokeRunsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataflowInvokeRunsResult
    {
        /// <summary>
        /// The application ID.
        /// </summary>
        public readonly string? ApplicationId;
        /// <summary>
        /// The OCID of a compartment.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// A user-friendly name. This name is not necessarily unique.
        /// </summary>
        public readonly string? DisplayName;
        public readonly string? DisplayNameStartsWith;
        public readonly ImmutableArray<Outputs.GetDataflowInvokeRunsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The OCID of the user who created the resource.
        /// </summary>
        public readonly string? OwnerPrincipalId;
        /// <summary>
        /// The list of runs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDataflowInvokeRunsRunResult> Runs;
        /// <summary>
        /// The current state of this run.
        /// </summary>
        public readonly string? State;
        public readonly string? TimeCreatedGreaterThan;

        [OutputConstructor]
        private GetDataflowInvokeRunsResult(
            string? applicationId,

            string compartmentId,

            string? displayName,

            string? displayNameStartsWith,

            ImmutableArray<Outputs.GetDataflowInvokeRunsFilterResult> filters,

            string id,

            string? ownerPrincipalId,

            ImmutableArray<Outputs.GetDataflowInvokeRunsRunResult> runs,

            string? state,

            string? timeCreatedGreaterThan)
        {
            ApplicationId = applicationId;
            CompartmentId = compartmentId;
            DisplayName = displayName;
            DisplayNameStartsWith = displayNameStartsWith;
            Filters = filters;
            Id = id;
            OwnerPrincipalId = ownerPrincipalId;
            Runs = runs;
            State = state;
            TimeCreatedGreaterThan = timeCreatedGreaterThan;
        }
    }
}