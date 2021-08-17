// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Outputs
{

    [OutputType]
    public sealed class DatabaseMigrationMigrationGoldenGateDetailsSettings
    {
        /// <summary>
        /// (Updatable) ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds.
        /// </summary>
        public readonly int? AcceptableLag;
        /// <summary>
        /// (Updatable) Parameters for Extract processes.
        /// </summary>
        public readonly Outputs.DatabaseMigrationMigrationGoldenGateDetailsSettingsExtract? Extract;
        /// <summary>
        /// (Updatable) Parameters for Replicat processes.
        /// </summary>
        public readonly Outputs.DatabaseMigrationMigrationGoldenGateDetailsSettingsReplicat? Replicat;

        [OutputConstructor]
        private DatabaseMigrationMigrationGoldenGateDetailsSettings(
            int? acceptableLag,

            Outputs.DatabaseMigrationMigrationGoldenGateDetailsSettingsExtract? extract,

            Outputs.DatabaseMigrationMigrationGoldenGateDetailsSettingsReplicat? replicat)
        {
            AcceptableLag = acceptableLag;
            Extract = extract;
            Replicat = replicat;
        }
    }
}