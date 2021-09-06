// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.ContainerEngine
{
    public static class GetWorkRequestErrors
    {
        /// <summary>
        /// This data source provides the list of Work Request Errors in Oracle Cloud Infrastructure Container Engine service.
        /// 
        /// Get the errors of a work request.
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
        ///         var testWorkRequestErrors = Output.Create(Oci.ContainerEngine.GetWorkRequestErrors.InvokeAsync(new Oci.ContainerEngine.GetWorkRequestErrorsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             WorkRequestId = oci_containerengine_work_request.Test_work_request.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetWorkRequestErrorsResult> InvokeAsync(GetWorkRequestErrorsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkRequestErrorsResult>("oci:containerengine/getWorkRequestErrors:getWorkRequestErrors", args ?? new GetWorkRequestErrorsArgs(), options.WithVersion());
    }


    public sealed class GetWorkRequestErrorsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The OCID of the compartment.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetWorkRequestErrorsFilterArgs>? _filters;
        public List<Inputs.GetWorkRequestErrorsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetWorkRequestErrorsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The OCID of the work request.
        /// </summary>
        [Input("workRequestId", required: true)]
        public string WorkRequestId { get; set; } = null!;

        public GetWorkRequestErrorsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkRequestErrorsResult
    {
        public readonly string CompartmentId;
        public readonly ImmutableArray<Outputs.GetWorkRequestErrorsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of work_request_errors.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetWorkRequestErrorsWorkRequestErrorResult> WorkRequestErrors;
        public readonly string WorkRequestId;

        [OutputConstructor]
        private GetWorkRequestErrorsResult(
            string compartmentId,

            ImmutableArray<Outputs.GetWorkRequestErrorsFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetWorkRequestErrorsWorkRequestErrorResult> workRequestErrors,

            string workRequestId)
        {
            CompartmentId = compartmentId;
            Filters = filters;
            Id = id;
            WorkRequestErrors = workRequestErrors;
            WorkRequestId = workRequestId;
        }
    }
}