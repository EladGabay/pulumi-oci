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
    public sealed class GetVaultSecretsSecretResult
    {
        /// <summary>
        /// The OCID of the compartment.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// A brief description of the secret. Avoid entering confidential information.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// The OCID of the secret.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The OCID of the master encryption key that is used to encrypt the secret.
        /// </summary>
        public readonly string KeyId;
        /// <summary>
        /// Additional information about the current lifecycle state of the secret.
        /// </summary>
        public readonly string LifecycleDetails;
        /// <summary>
        /// The user-friendly name of the secret. Avoid entering confidential information.
        /// </summary>
        public readonly string SecretName;
        /// <summary>
        /// A filter that returns only resources that match the specified lifecycle state. The state value is case-insensitive.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// A property indicating when the secret was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z`
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// An optional property indicating when the current secret version will expire, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z`
        /// </summary>
        public readonly string TimeOfCurrentVersionExpiry;
        /// <summary>
        /// An optional property indicating when to delete the secret, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z`
        /// </summary>
        public readonly string TimeOfDeletion;
        /// <summary>
        /// The OCID of the vault.
        /// </summary>
        public readonly string VaultId;

        [OutputConstructor]
        private GetVaultSecretsSecretResult(
            string compartmentId,

            ImmutableDictionary<string, object> definedTags,

            string description,

            ImmutableDictionary<string, object> freeformTags,

            string id,

            string keyId,

            string lifecycleDetails,

            string secretName,

            string state,

            string timeCreated,

            string timeOfCurrentVersionExpiry,

            string timeOfDeletion,

            string vaultId)
        {
            CompartmentId = compartmentId;
            DefinedTags = definedTags;
            Description = description;
            FreeformTags = freeformTags;
            Id = id;
            KeyId = keyId;
            LifecycleDetails = lifecycleDetails;
            SecretName = secretName;
            State = state;
            TimeCreated = timeCreated;
            TimeOfCurrentVersionExpiry = timeOfCurrentVersionExpiry;
            TimeOfDeletion = timeOfDeletion;
            VaultId = vaultId;
        }
    }
}