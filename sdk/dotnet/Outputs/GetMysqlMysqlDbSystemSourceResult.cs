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
    public sealed class GetMysqlMysqlDbSystemSourceResult
    {
        /// <summary>
        /// The OCID of the backup to be used as the source for the new DB System.
        /// </summary>
        public readonly string BackupId;
        /// <summary>
        /// The specific source identifier.
        /// </summary>
        public readonly string SourceType;

        [OutputConstructor]
        private GetMysqlMysqlDbSystemSourceResult(
            string backupId,

            string sourceType)
        {
            BackupId = backupId;
            SourceType = sourceType;
        }
    }
}