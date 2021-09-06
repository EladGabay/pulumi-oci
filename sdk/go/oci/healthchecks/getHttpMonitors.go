// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthchecks

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Http Monitors in Oracle Cloud Infrastructure Health Checks service.
//
// Gets a list of HTTP monitors.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/healthchecks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Http_monitor_display_name
// 		opt1 := _var.Http_monitor_home_region
// 		_, err := healthchecks.GetHttpMonitors(ctx, &healthchecks.GetHttpMonitorsArgs{
// 			CompartmentId: _var.Compartment_id,
// 			DisplayName:   &opt0,
// 			HomeRegion:    &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetHttpMonitors(ctx *pulumi.Context, args *GetHttpMonitorsArgs, opts ...pulumi.InvokeOption) (*GetHttpMonitorsResult, error) {
	var rv GetHttpMonitorsResult
	err := ctx.Invoke("oci:healthchecks/getHttpMonitors:getHttpMonitors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHttpMonitors.
type GetHttpMonitorsArgs struct {
	// Filters results by compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// Filters results that exactly match the `displayName` field.
	DisplayName *string                 `pulumi:"displayName"`
	Filters     []GetHttpMonitorsFilter `pulumi:"filters"`
	// Filters results that match the `homeRegion`.
	HomeRegion *string `pulumi:"homeRegion"`
}

// A collection of values returned by getHttpMonitors.
type GetHttpMonitorsResult struct {
	// The OCID of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// A user-friendly and mutable name suitable for display in a user interface.
	DisplayName *string                 `pulumi:"displayName"`
	Filters     []GetHttpMonitorsFilter `pulumi:"filters"`
	// The region where updates must be made and where results must be fetched from.
	HomeRegion *string `pulumi:"homeRegion"`
	// The list of http_monitors.
	HttpMonitors []GetHttpMonitorsHttpMonitor `pulumi:"httpMonitors"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}