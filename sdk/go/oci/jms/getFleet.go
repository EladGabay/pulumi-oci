// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package jms

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Fleet resource in Oracle Cloud Infrastructure Jms service.
//
// Retrieve a Fleet with the specified identifier.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/jms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := jms.LookupFleet(ctx, &jms.LookupFleetArgs{
// 			FleetId: oci_jms_fleet.Test_fleet.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupFleet(ctx *pulumi.Context, args *LookupFleetArgs, opts ...pulumi.InvokeOption) (*LookupFleetResult, error) {
	var rv LookupFleetResult
	err := ctx.Invoke("oci:jms/getFleet:getFleet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFleet.
type LookupFleetArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId string `pulumi:"fleetId"`
}

// A collection of values returned by getFleet.
type LookupFleetResult struct {
	// The approximate count of all unique applications in the Fleet in the past seven days. This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag.
	ApproximateApplicationCount int `pulumi:"approximateApplicationCount"`
	// The approximate count of all unique Java installations in the Fleet in the past seven days. This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag.
	ApproximateInstallationCount int `pulumi:"approximateInstallationCount"`
	// The approximate count of all unique Java Runtimes in the Fleet in the past seven days. This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag.
	ApproximateJreCount int `pulumi:"approximateJreCount"`
	// The approximate count of all unique managed instances in the Fleet in the past seven days. This metric is provided on a best-effort manner, and is not taken into account when computing the resource ETag.
	ApproximateManagedInstanceCount int `pulumi:"approximateManagedInstanceCount"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of the Fleet.
	CompartmentId string `pulumi:"compartmentId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. (See [Understanding Free-form Tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)).
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The Fleet's description.
	Description string `pulumi:"description"`
	// The name of the Fleet.
	DisplayName string `pulumi:"displayName"`
	FleetId     string `pulumi:"fleetId"`
	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`. (See [Managing Tags and Tag Namespaces](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).)
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
	Id string `pulumi:"id"`
	// The lifecycle state of the Fleet.
	State string `pulumi:"state"`
	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// The creation date and time of the Fleet (formatted according to RFC3339).
	TimeCreated string `pulumi:"timeCreated"`
}