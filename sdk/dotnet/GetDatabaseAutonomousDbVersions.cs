// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetDatabaseAutonomousDbVersions
    {
        /// <summary>
        /// This data source provides the list of Autonomous Db Versions in Oracle Cloud Infrastructure Database service.
        /// 
        /// Gets a list of supported Autonomous Database versions.
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
        ///         var testAutonomousDbVersions = Output.Create(Oci.GetDatabaseAutonomousDbVersions.InvokeAsync(new Oci.GetDatabaseAutonomousDbVersionsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             DbWorkload = @var.Autonomous_db_version_db_workload,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabaseAutonomousDbVersionsResult> InvokeAsync(GetDatabaseAutonomousDbVersionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseAutonomousDbVersionsResult>("oci:index/getDatabaseAutonomousDbVersions:GetDatabaseAutonomousDbVersions", args ?? new GetDatabaseAutonomousDbVersionsArgs(), options.WithVersion());
    }


    public sealed class GetDatabaseAutonomousDbVersionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// A filter to return only autonomous database resources that match the specified workload type.
        /// </summary>
        [Input("dbWorkload")]
        public string? DbWorkload { get; set; }

        [Input("filters")]
        private List<Inputs.GetDatabaseAutonomousDbVersionsFilterArgs>? _filters;
        public List<Inputs.GetDatabaseAutonomousDbVersionsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDatabaseAutonomousDbVersionsFilterArgs>());
            set => _filters = value;
        }

        public GetDatabaseAutonomousDbVersionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseAutonomousDbVersionsResult
    {
        /// <summary>
        /// The list of autonomous_db_versions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseAutonomousDbVersionsAutonomousDbVersionResult> AutonomousDbVersions;
        public readonly string CompartmentId;
        /// <summary>
        /// The Autonomous Database workload type. The following values are valid:
        /// * OLTP - indicates an Autonomous Transaction Processing database
        /// * DW - indicates an Autonomous Data Warehouse database
        /// * AJD - indicates an Autonomous JSON Database
        /// * APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.
        /// </summary>
        public readonly string? DbWorkload;
        public readonly ImmutableArray<Outputs.GetDatabaseAutonomousDbVersionsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDatabaseAutonomousDbVersionsResult(
            ImmutableArray<Outputs.GetDatabaseAutonomousDbVersionsAutonomousDbVersionResult> autonomousDbVersions,

            string compartmentId,

            string? dbWorkload,

            ImmutableArray<Outputs.GetDatabaseAutonomousDbVersionsFilterResult> filters,

            string id)
        {
            AutonomousDbVersions = autonomousDbVersions;
            CompartmentId = compartmentId;
            DbWorkload = dbWorkload;
            Filters = filters;
            Id = id;
        }
    }
}