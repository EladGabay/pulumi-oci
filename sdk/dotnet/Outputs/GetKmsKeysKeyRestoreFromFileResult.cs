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
    public sealed class GetKmsKeysKeyRestoreFromFileResult
    {
        public readonly string ContentLength;
        public readonly string ContentMd5;
        public readonly string RestoreKeyFromFileDetails;

        [OutputConstructor]
        private GetKmsKeysKeyRestoreFromFileResult(
            string contentLength,

            string contentMd5,

            string restoreKeyFromFileDetails)
        {
            ContentLength = contentLength;
            ContentMd5 = contentMd5;
            RestoreKeyFromFileDetails = restoreKeyFromFileDetails;
        }
    }
}