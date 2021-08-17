// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Log Group resource in Oracle Cloud Infrastructure Logging service.
//
// Get the specified log group's information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oci.GetLoggingLogGroup(ctx, &GetLoggingLogGroupArgs{
// 			LogGroupId: oci_logging_log_group.Test_log_group.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupLoggingLogGroup(ctx *pulumi.Context, args *LookupLoggingLogGroupArgs, opts ...pulumi.InvokeOption) (*LookupLoggingLogGroupResult, error) {
	var rv LookupLoggingLogGroupResult
	err := ctx.Invoke("oci:index/getLoggingLogGroup:GetLoggingLogGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetLoggingLogGroup.
type LookupLoggingLogGroupArgs struct {
	// OCID of a log group to work with.
	LogGroupId string `pulumi:"logGroupId"`
}

// A collection of values returned by GetLoggingLogGroup.
type LookupLoggingLogGroupResult struct {
	// The OCID of the compartment that the resource belongs to.
	CompartmentId string `pulumi:"compartmentId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// Description for this resource.
	Description string `pulumi:"description"`
	// The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information.
	DisplayName string `pulumi:"displayName"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The OCID of the resource.
	Id         string `pulumi:"id"`
	LogGroupId string `pulumi:"logGroupId"`
	// The log group object state.
	State string `pulumi:"state"`
	// Time the resource was created.
	TimeCreated string `pulumi:"timeCreated"`
	// Time the resource was last modified.
	TimeLastModified string `pulumi:"timeLastModified"`
}