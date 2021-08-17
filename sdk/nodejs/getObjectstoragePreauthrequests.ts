// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Preauthenticated Requests in Oracle Cloud Infrastructure Object Storage service.
 *
 * Lists pre-authenticated requests for the bucket.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testPreauthenticatedRequests = oci.GetObjectstoragePreauthrequests({
 *     bucket: _var.preauthenticated_request_bucket,
 *     namespace: _var.preauthenticated_request_namespace,
 *     objectNamePrefix: _var.preauthenticated_request_object_name_prefix,
 * });
 * ```
 */
export function getObjectstoragePreauthrequests(args: GetObjectstoragePreauthrequestsArgs, opts?: pulumi.InvokeOptions): Promise<GetObjectstoragePreauthrequestsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getObjectstoragePreauthrequests:GetObjectstoragePreauthrequests", {
        "bucket": args.bucket,
        "filters": args.filters,
        "namespace": args.namespace,
        "objectNamePrefix": args.objectNamePrefix,
    }, opts);
}

/**
 * A collection of arguments for invoking GetObjectstoragePreauthrequests.
 */
export interface GetObjectstoragePreauthrequestsArgs {
    /**
     * The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1`
     */
    bucket: string;
    filters?: inputs.GetObjectstoragePreauthrequestsFilter[];
    /**
     * The Object Storage namespace used for the request.
     */
    namespace: string;
    /**
     * User-specified object name prefixes can be used to query and return a list of pre-authenticated requests.
     */
    objectNamePrefix?: string;
}

/**
 * A collection of values returned by GetObjectstoragePreauthrequests.
 */
export interface GetObjectstoragePreauthrequestsResult {
    /**
     * The name of the bucket.  Example: `my-new-bucket1`
     */
    readonly bucket: string;
    readonly filters?: outputs.GetObjectstoragePreauthrequestsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Object Storage namespace used for the request.
     */
    readonly namespace: string;
    readonly objectNamePrefix?: string;
    /**
     * The list of preauthenticated_requests.
     */
    readonly preauthenticatedRequests: outputs.GetObjectstoragePreauthrequestsPreauthenticatedRequest[];
}