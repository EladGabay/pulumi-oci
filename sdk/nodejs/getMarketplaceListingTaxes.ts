// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Listing Taxes in Oracle Cloud Infrastructure Marketplace service.
 *
 * Returns list of all tax implications that current tenant may be liable to once they launch the listing.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testListingTaxes = oci.GetMarketplaceListingTaxes({
 *     listingId: oci_marketplace_listing.test_listing.id,
 *     compartmentId: _var.compartment_id,
 * });
 * ```
 */
export function getMarketplaceListingTaxes(args: GetMarketplaceListingTaxesArgs, opts?: pulumi.InvokeOptions): Promise<GetMarketplaceListingTaxesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getMarketplaceListingTaxes:GetMarketplaceListingTaxes", {
        "compartmentId": args.compartmentId,
        "filters": args.filters,
        "listingId": args.listingId,
    }, opts);
}

/**
 * A collection of arguments for invoking GetMarketplaceListingTaxes.
 */
export interface GetMarketplaceListingTaxesArgs {
    /**
     * The unique identifier for the compartment.
     */
    compartmentId?: string;
    filters?: inputs.GetMarketplaceListingTaxesFilter[];
    /**
     * The unique identifier for the listing.
     */
    listingId: string;
}

/**
 * A collection of values returned by GetMarketplaceListingTaxes.
 */
export interface GetMarketplaceListingTaxesResult {
    readonly compartmentId?: string;
    readonly filters?: outputs.GetMarketplaceListingTaxesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly listingId: string;
    /**
     * The list of taxes.
     */
    readonly taxes: outputs.GetMarketplaceListingTaxesTax[];
}