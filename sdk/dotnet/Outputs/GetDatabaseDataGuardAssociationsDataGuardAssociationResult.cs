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
    public sealed class GetDatabaseDataGuardAssociationsDataGuardAssociationResult
    {
        /// <summary>
        /// The lag time between updates to the primary database and application of the redo data on the standby database, as computed by the reporting database.  Example: `9 seconds`
        /// </summary>
        public readonly string ApplyLag;
        /// <summary>
        /// The rate at which redo logs are synced between the associated databases.  Example: `180 Mb per second`
        /// </summary>
        public readonly string ApplyRate;
        public readonly string AvailabilityDomain;
        public readonly ImmutableArray<string> BackupNetworkNsgIds;
        public readonly bool CreateAsync;
        public readonly string CreationType;
        public readonly string DatabaseAdminPassword;
        /// <summary>
        /// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        public readonly string DatabaseId;
        public readonly string DatabaseSoftwareImageId;
        public readonly string DeleteStandbyDbHomeOnDelete;
        public readonly string DisplayName;
        public readonly string Hostname;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Data Guard association.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Additional information about the current lifecycleState, if available.
        /// </summary>
        public readonly string LifecycleDetails;
        public readonly ImmutableArray<string> NsgIds;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer database's Data Guard association.
        /// </summary>
        public readonly string PeerDataGuardAssociationId;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated peer database.
        /// </summary>
        public readonly string PeerDatabaseId;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home containing the associated peer database.
        /// </summary>
        public readonly string PeerDbHomeId;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system containing the associated peer database.
        /// </summary>
        public readonly string PeerDbSystemId;
        /// <summary>
        /// The role of the peer database in this Data Guard association.
        /// </summary>
        public readonly string PeerRole;
        public readonly string PeerVmClusterId;
        /// <summary>
        /// The protection mode of this Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation.
        /// </summary>
        public readonly string ProtectionMode;
        /// <summary>
        /// The role of the reporting database in this Data Guard association.
        /// </summary>
        public readonly string Role;
        public readonly string Shape;
        /// <summary>
        /// The current state of the Data Guard association.
        /// </summary>
        public readonly string State;
        public readonly string SubnetId;
        /// <summary>
        /// The date and time the Data Guard association was created.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The redo transport type used by this Data Guard association.  For more information, see [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400) in the Oracle Data Guard documentation.
        /// </summary>
        public readonly string TransportType;

        [OutputConstructor]
        private GetDatabaseDataGuardAssociationsDataGuardAssociationResult(
            string applyLag,

            string applyRate,

            string availabilityDomain,

            ImmutableArray<string> backupNetworkNsgIds,

            bool createAsync,

            string creationType,

            string databaseAdminPassword,

            string databaseId,

            string databaseSoftwareImageId,

            string deleteStandbyDbHomeOnDelete,

            string displayName,

            string hostname,

            string id,

            string lifecycleDetails,

            ImmutableArray<string> nsgIds,

            string peerDataGuardAssociationId,

            string peerDatabaseId,

            string peerDbHomeId,

            string peerDbSystemId,

            string peerRole,

            string peerVmClusterId,

            string protectionMode,

            string role,

            string shape,

            string state,

            string subnetId,

            string timeCreated,

            string transportType)
        {
            ApplyLag = applyLag;
            ApplyRate = applyRate;
            AvailabilityDomain = availabilityDomain;
            BackupNetworkNsgIds = backupNetworkNsgIds;
            CreateAsync = createAsync;
            CreationType = creationType;
            DatabaseAdminPassword = databaseAdminPassword;
            DatabaseId = databaseId;
            DatabaseSoftwareImageId = databaseSoftwareImageId;
            DeleteStandbyDbHomeOnDelete = deleteStandbyDbHomeOnDelete;
            DisplayName = displayName;
            Hostname = hostname;
            Id = id;
            LifecycleDetails = lifecycleDetails;
            NsgIds = nsgIds;
            PeerDataGuardAssociationId = peerDataGuardAssociationId;
            PeerDatabaseId = peerDatabaseId;
            PeerDbHomeId = peerDbHomeId;
            PeerDbSystemId = peerDbSystemId;
            PeerRole = peerRole;
            PeerVmClusterId = peerVmClusterId;
            ProtectionMode = protectionMode;
            Role = role;
            Shape = shape;
            State = state;
            SubnetId = subnetId;
            TimeCreated = timeCreated;
            TransportType = transportType;
        }
    }
}