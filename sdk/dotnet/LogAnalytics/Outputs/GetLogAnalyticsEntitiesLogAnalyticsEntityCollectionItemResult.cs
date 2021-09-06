// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.LogAnalytics.Outputs
{

    [OutputType]
    public sealed class GetLogAnalyticsEntitiesLogAnalyticsEntityCollectionItemResult
    {
        /// <summary>
        /// The Boolean flag to indicate if logs are collected for an entity for log analytics usage.
        /// </summary>
        public readonly bool AreLogsCollected;
        /// <summary>
        /// A filter to return only log analytics entities whose cloudResourceId matches the cloudResourceId given.
        /// </summary>
        public readonly string CloudResourceId;
        /// <summary>
        /// The ID of the compartment in which to list resources.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// Internal name for the log analytics entity type.
        /// </summary>
        public readonly string EntityTypeInternalName;
        /// <summary>
        /// A filter to return only log analytics entities whose entityTypeName matches the entire log analytics entity type name of one of the entityTypeNames given in the list. The match is case-insensitive.
        /// </summary>
        public readonly string EntityTypeName;
        /// <summary>
        /// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// A filter to return only log analytics entities whose hostname matches the entire hostname given.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The log analytics entity OCID. This ID is a reference used by log analytics features and it represents a resource that is provisioned and managed by the customer on their premises or on the cloud.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// lifecycleDetails has additional information regarding substeps such as management agent plugin deployment.
        /// </summary>
        public readonly string LifecycleDetails;
        /// <summary>
        /// Management agent (management-agents resource kind) compartment OCID
        /// </summary>
        public readonly string ManagementAgentCompartmentId;
        /// <summary>
        /// Management agent (management-agents resource kind) display name
        /// </summary>
        public readonly string ManagementAgentDisplayName;
        /// <summary>
        /// The OCID of the Management Agent.
        /// </summary>
        public readonly string ManagementAgentId;
        /// <summary>
        /// A filter to return only log analytics entities whose name matches the entire name given. The match is case-insensitive.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Logging Analytics namespace used for the request.
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// The name/value pairs for parameter values to be used in file patterns specified in log sources.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Properties;
        /// <summary>
        /// A filter to return only log analytics entities whose sourceId matches the sourceId given.
        /// </summary>
        public readonly string SourceId;
        /// <summary>
        /// A filter to return only those log analytics entities with the specified lifecycle state. The state value is case-insensitive.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The date and time the resource was created, in the format defined by RFC3339.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The date and time the resource was last updated, in the format defined by RFC3339.
        /// </summary>
        public readonly string TimeUpdated;
        /// <summary>
        /// The timezone region of the log analytics entity.
        /// </summary>
        public readonly string TimezoneRegion;

        [OutputConstructor]
        private GetLogAnalyticsEntitiesLogAnalyticsEntityCollectionItemResult(
            bool areLogsCollected,

            string cloudResourceId,

            string compartmentId,

            ImmutableDictionary<string, object> definedTags,

            string entityTypeInternalName,

            string entityTypeName,

            ImmutableDictionary<string, object> freeformTags,

            string hostname,

            string id,

            string lifecycleDetails,

            string managementAgentCompartmentId,

            string managementAgentDisplayName,

            string managementAgentId,

            string name,

            string @namespace,

            ImmutableDictionary<string, object> properties,

            string sourceId,

            string state,

            string timeCreated,

            string timeUpdated,

            string timezoneRegion)
        {
            AreLogsCollected = areLogsCollected;
            CloudResourceId = cloudResourceId;
            CompartmentId = compartmentId;
            DefinedTags = definedTags;
            EntityTypeInternalName = entityTypeInternalName;
            EntityTypeName = entityTypeName;
            FreeformTags = freeformTags;
            Hostname = hostname;
            Id = id;
            LifecycleDetails = lifecycleDetails;
            ManagementAgentCompartmentId = managementAgentCompartmentId;
            ManagementAgentDisplayName = managementAgentDisplayName;
            ManagementAgentId = managementAgentId;
            Name = name;
            Namespace = @namespace;
            Properties = properties;
            SourceId = sourceId;
            State = state;
            TimeCreated = timeCreated;
            TimeUpdated = timeUpdated;
            TimezoneRegion = timezoneRegion;
        }
    }
}