// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.DatabaseMigration.Outputs
{

    [OutputType]
    public sealed class GetMigrationsMigrationCollectionItemDatapumpSettingsResult
    {
        /// <summary>
        /// Optional parameters for Datapump Export and Import. Refer to https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/ODMS_DATAPUMP.html#GUID-62324358-2F26-4A94-B69F-1075D53FA96D__BABDECJE
        /// </summary>
        public readonly Outputs.GetMigrationsMigrationCollectionItemDatapumpSettingsDataPumpParametersResult DataPumpParameters;
        /// <summary>
        /// Directory object details, used to define either import or export directory objects in Data Pump Settings.
        /// </summary>
        public readonly Outputs.GetMigrationsMigrationCollectionItemDatapumpSettingsExportDirectoryObjectResult ExportDirectoryObject;
        /// <summary>
        /// Directory object details, used to define either import or export directory objects in Data Pump Settings.
        /// </summary>
        public readonly Outputs.GetMigrationsMigrationCollectionItemDatapumpSettingsImportDirectoryObjectResult ImportDirectoryObject;
        /// <summary>
        /// DataPump job mode. Refer to docs.oracle.com/en/database/oracle/oracle-database/19/arpls/ODMS_DATAPUMP.html#GUID-92C2CB46-8BC9-414D-B62E-79CD788C1E62__BABBDEHD
        /// </summary>
        public readonly string JobMode;
        /// <summary>
        /// Defines remapping to be applied to objects as they are processed. Refer to https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/ODMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMigrationsMigrationCollectionItemDatapumpSettingsMetadataRemapResult> MetadataRemaps;

        [OutputConstructor]
        private GetMigrationsMigrationCollectionItemDatapumpSettingsResult(
            Outputs.GetMigrationsMigrationCollectionItemDatapumpSettingsDataPumpParametersResult dataPumpParameters,

            Outputs.GetMigrationsMigrationCollectionItemDatapumpSettingsExportDirectoryObjectResult exportDirectoryObject,

            Outputs.GetMigrationsMigrationCollectionItemDatapumpSettingsImportDirectoryObjectResult importDirectoryObject,

            string jobMode,

            ImmutableArray<Outputs.GetMigrationsMigrationCollectionItemDatapumpSettingsMetadataRemapResult> metadataRemaps)
        {
            DataPumpParameters = dataPumpParameters;
            ExportDirectoryObject = exportDirectoryObject;
            ImportDirectoryObject = importDirectoryObject;
            JobMode = jobMode;
            MetadataRemaps = metadataRemaps;
        }
    }
}