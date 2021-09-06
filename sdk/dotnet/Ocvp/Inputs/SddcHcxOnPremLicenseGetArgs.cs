// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Ocvp.Inputs
{

    public sealed class SddcHcxOnPremLicenseGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// HCX on-premise license key value
        /// </summary>
        [Input("activationKey")]
        public Input<string>? ActivationKey { get; set; }

        /// <summary>
        /// status of HCX on-premise license
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Name of the system that consumed the HCX on-premise license
        /// </summary>
        [Input("systemName")]
        public Input<string>? SystemName { get; set; }

        public SddcHcxOnPremLicenseGetArgs()
        {
        }
    }
}