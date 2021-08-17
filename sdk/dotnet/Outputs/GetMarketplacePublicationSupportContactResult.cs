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
    public sealed class GetMarketplacePublicationSupportContactResult
    {
        /// <summary>
        /// The email of the contact.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// name of the operating system
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The phone number of the contact.
        /// </summary>
        public readonly string Phone;
        /// <summary>
        /// The email subject line to use when contacting support.
        /// </summary>
        public readonly string Subject;

        [OutputConstructor]
        private GetMarketplacePublicationSupportContactResult(
            string email,

            string name,

            string phone,

            string subject)
        {
            Email = email;
            Name = name;
            Phone = phone;
            Subject = subject;
        }
    }
}