// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package audit

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure Audit service.
//
// Get the configuration
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/audit"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := audit.LookupConfiguration(ctx, &audit.LookupConfigurationArgs{
// 			CompartmentId: _var.Tenancy_ocid,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupConfiguration(ctx *pulumi.Context, args *LookupConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationResult, error) {
	var rv LookupConfigurationResult
	err := ctx.Invoke("oci:audit/getConfiguration:getConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConfiguration.
type LookupConfigurationArgs struct {
	// ID of the root compartment (tenancy)
	CompartmentId string `pulumi:"compartmentId"`
}

// A collection of values returned by getConfiguration.
type LookupConfigurationResult struct {
	CompartmentId string `pulumi:"compartmentId"`
	Id            string `pulumi:"id"`
	// The retention period setting, specified in days. The minimum is 90, the maximum 365.  Example: `90`
	RetentionPeriodDays int `pulumi:"retentionPeriodDays"`
}