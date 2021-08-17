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
    public sealed class GetDatabaseMigrationMigrationGoldenGateDetailsSettingsResult
    {
        /// <summary>
        /// ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds.
        /// </summary>
        public readonly int AcceptableLag;
        /// <summary>
        /// Parameters for Extract processes.
        /// </summary>
        public readonly Outputs.GetDatabaseMigrationMigrationGoldenGateDetailsSettingsExtractResult Extract;
        /// <summary>
        /// Parameters for Replicat processes.
        /// </summary>
        public readonly Outputs.GetDatabaseMigrationMigrationGoldenGateDetailsSettingsReplicatResult Replicat;

        [OutputConstructor]
        private GetDatabaseMigrationMigrationGoldenGateDetailsSettingsResult(
            int acceptableLag,

            Outputs.GetDatabaseMigrationMigrationGoldenGateDetailsSettingsExtractResult extract,

            Outputs.GetDatabaseMigrationMigrationGoldenGateDetailsSettingsReplicatResult replicat)
        {
            AcceptableLag = acceptableLag;
            Extract = extract;
            Replicat = replicat;
        }
    }
}