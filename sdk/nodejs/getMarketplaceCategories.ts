// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Categories in Oracle Cloud Infrastructure Marketplace service.
 *
 * Gets the list of all the categories for listings published to Oracle Cloud Infrastructure Marketplace. Categories apply
 * to the software product provided by the listing.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testCategories = oci.GetMarketplaceCategories({
 *     compartmentId: _var.compartment_id,
 * });
 * ```
 */
export function getMarketplaceCategories(args?: GetMarketplaceCategoriesArgs, opts?: pulumi.InvokeOptions): Promise<GetMarketplaceCategoriesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getMarketplaceCategories:GetMarketplaceCategories", {
        "compartmentId": args.compartmentId,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking GetMarketplaceCategories.
 */
export interface GetMarketplaceCategoriesArgs {
    /**
     * The unique identifier for the compartment.
     */
    compartmentId?: string;
    filters?: inputs.GetMarketplaceCategoriesFilter[];
}

/**
 * A collection of values returned by GetMarketplaceCategories.
 */
export interface GetMarketplaceCategoriesResult {
    /**
     * The list of categories.
     */
    readonly categories: outputs.GetMarketplaceCategoriesCategory[];
    readonly compartmentId?: string;
    readonly filters?: outputs.GetMarketplaceCategoriesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}