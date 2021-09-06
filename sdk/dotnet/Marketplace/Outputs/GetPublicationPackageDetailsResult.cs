// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Marketplace.Outputs
{

    [OutputType]
    public sealed class GetPublicationPackageDetailsResult
    {
        public readonly ImmutableArray<Outputs.GetPublicationPackageDetailsEulaResult> Eulas;
        public readonly string ImageId;
        public readonly Outputs.GetPublicationPackageDetailsOperatingSystemResult OperatingSystem;
        /// <summary>
        /// The listing's package type.
        /// </summary>
        public readonly string PackageType;
        public readonly string PackageVersion;

        [OutputConstructor]
        private GetPublicationPackageDetailsResult(
            ImmutableArray<Outputs.GetPublicationPackageDetailsEulaResult> eulas,

            string imageId,

            Outputs.GetPublicationPackageDetailsOperatingSystemResult operatingSystem,

            string packageType,

            string packageVersion)
        {
            Eulas = eulas;
            ImageId = imageId;
            OperatingSystem = operatingSystem;
            PackageType = packageType;
            PackageVersion = packageVersion;
        }
    }
}