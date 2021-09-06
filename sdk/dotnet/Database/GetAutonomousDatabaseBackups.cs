// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Database
{
    public static class GetAutonomousDatabaseBackups
    {
        /// <summary>
        /// This data source provides the list of Autonomous Database Backups in Oracle Cloud Infrastructure Database service.
        /// 
        /// Gets a list of Autonomous Database backups based on either the `autonomousDatabaseId` or `compartmentId` specified as a query parameter.
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
        ///         var testAutonomousDatabaseBackups = Output.Create(Oci.Database.GetAutonomousDatabaseBackups.InvokeAsync(new Oci.Database.GetAutonomousDatabaseBackupsArgs
        ///         {
        ///             AutonomousDatabaseId = oci_database_autonomous_database.Test_autonomous_database.Id,
        ///             CompartmentId = @var.Compartment_id,
        ///             DisplayName = @var.Autonomous_database_backup_display_name,
        ///             State = @var.Autonomous_database_backup_state,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAutonomousDatabaseBackupsResult> InvokeAsync(GetAutonomousDatabaseBackupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAutonomousDatabaseBackupsResult>("oci:database/getAutonomousDatabaseBackups:getAutonomousDatabaseBackups", args ?? new GetAutonomousDatabaseBackupsArgs(), options.WithVersion());
    }


    public sealed class GetAutonomousDatabaseBackupsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        [Input("autonomousDatabaseId")]
        public string? AutonomousDatabaseId { get; set; }

        /// <summary>
        /// The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        [Input("compartmentId")]
        public string? CompartmentId { get; set; }

        /// <summary>
        /// A filter to return only resources that match the entire display name given. The match is not case sensitive.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("filters")]
        private List<Inputs.GetAutonomousDatabaseBackupsFilterArgs>? _filters;
        public List<Inputs.GetAutonomousDatabaseBackupsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetAutonomousDatabaseBackupsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// A filter to return only resources that match the given lifecycle state exactly.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetAutonomousDatabaseBackupsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAutonomousDatabaseBackupsResult
    {
        /// <summary>
        /// The list of autonomous_database_backups.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAutonomousDatabaseBackupsAutonomousDatabaseBackupResult> AutonomousDatabaseBackups;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
        /// </summary>
        public readonly string? AutonomousDatabaseId;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        /// </summary>
        public readonly string? CompartmentId;
        /// <summary>
        /// The user-friendly name for the backup. The name does not have to be unique.
        /// </summary>
        public readonly string? DisplayName;
        public readonly ImmutableArray<Outputs.GetAutonomousDatabaseBackupsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The current state of the backup.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private GetAutonomousDatabaseBackupsResult(
            ImmutableArray<Outputs.GetAutonomousDatabaseBackupsAutonomousDatabaseBackupResult> autonomousDatabaseBackups,

            string? autonomousDatabaseId,

            string? compartmentId,

            string? displayName,

            ImmutableArray<Outputs.GetAutonomousDatabaseBackupsFilterResult> filters,

            string id,

            string? state)
        {
            AutonomousDatabaseBackups = autonomousDatabaseBackups;
            AutonomousDatabaseId = autonomousDatabaseId;
            CompartmentId = compartmentId;
            DisplayName = displayName;
            Filters = filters;
            Id = id;
            State = state;
        }
    }
}