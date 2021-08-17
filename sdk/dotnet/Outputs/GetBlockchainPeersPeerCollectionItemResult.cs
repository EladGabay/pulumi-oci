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
    public sealed class GetBlockchainPeersPeerCollectionItemResult
    {
        /// <summary>
        /// Availability Domain of peer
        /// </summary>
        public readonly string Ad;
        /// <summary>
        /// peer alias
        /// </summary>
        public readonly string Alias;
        /// <summary>
        /// Unique service identifier.
        /// </summary>
        public readonly string BlockchainPlatformId;
        /// <summary>
        /// Host on which the Peer exists
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// OCPU allocation parameter
        /// </summary>
        public readonly Outputs.GetBlockchainPeersPeerCollectionItemOcpuAllocationParamResult OcpuAllocationParam;
        /// <summary>
        /// peer identifier
        /// </summary>
        public readonly string PeerKey;
        /// <summary>
        /// Peer role
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// The current state of the peer.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetBlockchainPeersPeerCollectionItemResult(
            string ad,

            string alias,

            string blockchainPlatformId,

            string host,

            Outputs.GetBlockchainPeersPeerCollectionItemOcpuAllocationParamResult ocpuAllocationParam,

            string peerKey,

            string role,

            string state)
        {
            Ad = ad;
            Alias = alias;
            BlockchainPlatformId = blockchainPlatformId;
            Host = host;
            OcpuAllocationParam = ocpuAllocationParam;
            PeerKey = peerKey;
            Role = role;
            State = state;
        }
    }
}