// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetDatabaseAutonomousContainerPatches
    {
        /// <summary>
        /// This data source provides the list of Autonomous Container Patches in Oracle Cloud Infrastructure Database service.
        /// 
        /// Lists the patches applicable to the requested container database.
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
        ///         var testAutonomousContainerPatches = Output.Create(Oci.GetDatabaseAutonomousContainerPatches.InvokeAsync(new Oci.GetDatabaseAutonomousContainerPatchesArgs
        ///         {
        ///             AutonomousContainerDatabaseId = oci_database_autonomous_container_database.Test_autonomous_container_database.Id,
        ///             CompartmentId = @var.Compartment_id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabaseAutonomousContainerPatchesResult> InvokeAsync(GetDatabaseAutonomousContainerPatchesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseAutonomousContainerPatchesResult>("oci:index/getDatabaseAutonomousContainerPatches:GetDatabaseAutonomousContainerPatches", args ?? new GetDatabaseAutonomousContainerPatchesArgs(), options.WithVersion());
    }


    public sealed class GetDatabaseAutonomousContainerPatchesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        [Input("autonomousContainerDatabaseId", required: true)]
        public string AutonomousContainerDatabaseId { get; set; } = null!;

        /// <summary>
        /// The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetDatabaseAutonomousContainerPatchesFilterArgs>? _filters;
        public List<Inputs.GetDatabaseAutonomousContainerPatchesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDatabaseAutonomousContainerPatchesFilterArgs>());
            set => _filters = value;
        }

        public GetDatabaseAutonomousContainerPatchesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseAutonomousContainerPatchesResult
    {
        public readonly string AutonomousContainerDatabaseId;
        /// <summary>
        /// The list of autonomous_patches.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseAutonomousContainerPatchesAutonomousPatchResult> AutonomousPatches;
        public readonly string CompartmentId;
        public readonly ImmutableArray<Outputs.GetDatabaseAutonomousContainerPatchesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDatabaseAutonomousContainerPatchesResult(
            string autonomousContainerDatabaseId,

            ImmutableArray<Outputs.GetDatabaseAutonomousContainerPatchesAutonomousPatchResult> autonomousPatches,

            string compartmentId,

            ImmutableArray<Outputs.GetDatabaseAutonomousContainerPatchesFilterResult> filters,

            string id)
        {
            AutonomousContainerDatabaseId = autonomousContainerDatabaseId;
            AutonomousPatches = autonomousPatches;
            CompartmentId = compartmentId;
            Filters = filters;
            Id = id;
        }
    }
}