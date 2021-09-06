// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Database.Outputs
{

    [OutputType]
    public sealed class GetAutonomousContainerDatabaseDataguardAssociationsAutonomousContainerDatabaseDataguardAssociationResult
    {
        /// <summary>
        /// The lag time between updates to the primary Autonomous Container Database and application of the redo data on the standby Autonomous Container Database, as computed by the reporting database.  Example: `9 seconds`
        /// </summary>
        public readonly string ApplyLag;
        /// <summary>
        /// The rate at which redo logs are synchronized between the associated Autonomous Container Databases.  Example: `180 Mb per second`
        /// </summary>
        public readonly string ApplyRate;
        /// <summary>
        /// The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        public readonly string AutonomousContainerDatabaseId;
        /// <summary>
        /// The OCID of the Autonomous Data Guard created for a given Autonomous Container Database.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Additional information about the current lifecycleState, if available.
        /// </summary>
        public readonly string LifecycleDetails;
        /// <summary>
        /// The OCID of the peer Autonomous Container Database-Autonomous Data Guard association.
        /// </summary>
        public readonly string PeerAutonomousContainerDatabaseDataguardAssociationId;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous Container Database.
        /// </summary>
        public readonly string PeerAutonomousContainerDatabaseId;
        /// <summary>
        /// The current state of Autonomous Data Guard.
        /// </summary>
        public readonly string PeerLifecycleState;
        /// <summary>
        /// The Data Guard role of the Autonomous Container Database, if Autonomous Data Guard is enabled.
        /// </summary>
        public readonly string PeerRole;
        /// <summary>
        /// The protection mode of this Autonomous Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation.
        /// </summary>
        public readonly string ProtectionMode;
        /// <summary>
        /// The Data Guard role of the Autonomous Container Database, if Autonomous Data Guard is enabled.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// The current state of Autonomous Data Guard.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The date and time the Autonomous DataGuard association was created.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The date and time when the last role change action happened.
        /// </summary>
        public readonly string TimeLastRoleChanged;
        /// <summary>
        /// The date and time of the last update to the apply lag, apply rate, and transport lag values.
        /// </summary>
        public readonly string TimeLastSynced;
        /// <summary>
        /// The approximate number of seconds of redo data not yet available on the standby Autonomous Container Database, as computed by the reporting database.  Example: `7 seconds`
        /// </summary>
        public readonly string TransportLag;

        [OutputConstructor]
        private GetAutonomousContainerDatabaseDataguardAssociationsAutonomousContainerDatabaseDataguardAssociationResult(
            string applyLag,

            string applyRate,

            string autonomousContainerDatabaseId,

            string id,

            string lifecycleDetails,

            string peerAutonomousContainerDatabaseDataguardAssociationId,

            string peerAutonomousContainerDatabaseId,

            string peerLifecycleState,

            string peerRole,

            string protectionMode,

            string role,

            string state,

            string timeCreated,

            string timeLastRoleChanged,

            string timeLastSynced,

            string transportLag)
        {
            ApplyLag = applyLag;
            ApplyRate = applyRate;
            AutonomousContainerDatabaseId = autonomousContainerDatabaseId;
            Id = id;
            LifecycleDetails = lifecycleDetails;
            PeerAutonomousContainerDatabaseDataguardAssociationId = peerAutonomousContainerDatabaseDataguardAssociationId;
            PeerAutonomousContainerDatabaseId = peerAutonomousContainerDatabaseId;
            PeerLifecycleState = peerLifecycleState;
            PeerRole = peerRole;
            ProtectionMode = protectionMode;
            Role = role;
            State = state;
            TimeCreated = timeCreated;
            TimeLastRoleChanged = timeLastRoleChanged;
            TimeLastSynced = timeLastSynced;
            TransportLag = transportLag;
        }
    }
}