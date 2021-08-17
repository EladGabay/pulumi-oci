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
    public sealed class GetNosqlTablesTableCollectionSchemaResult
    {
        /// <summary>
        /// The columns of a table.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNosqlTablesTableCollectionSchemaColumnResult> Columns;
        /// <summary>
        /// A list of column names that make up a key.
        /// </summary>
        public readonly ImmutableArray<string> PrimaryKeys;
        /// <summary>
        /// A list of column names that make up a key.
        /// </summary>
        public readonly ImmutableArray<string> ShardKeys;
        /// <summary>
        /// The default Time-to-Live for the table, in days.
        /// </summary>
        public readonly int Ttl;

        [OutputConstructor]
        private GetNosqlTablesTableCollectionSchemaResult(
            ImmutableArray<Outputs.GetNosqlTablesTableCollectionSchemaColumnResult> columns,

            ImmutableArray<string> primaryKeys,

            ImmutableArray<string> shardKeys,

            int ttl)
        {
            Columns = columns;
            PrimaryKeys = primaryKeys;
            ShardKeys = shardKeys;
            Ttl = ttl;
        }
    }
}