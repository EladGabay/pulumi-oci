// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package integration

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Integration Instances in Oracle Cloud Infrastructure Integration service.
//
// Returns a list of Integration Instances.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/integration"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Integration_instance_display_name
// 		opt1 := _var.Integration_instance_state
// 		_, err := integration.GetIntegrationInstances(ctx, &integration.GetIntegrationInstancesArgs{
// 			CompartmentId: _var.Compartment_id,
// 			DisplayName:   &opt0,
// 			State:         &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIntegrationInstances(ctx *pulumi.Context, args *GetIntegrationInstancesArgs, opts ...pulumi.InvokeOption) (*GetIntegrationInstancesResult, error) {
	var rv GetIntegrationInstancesResult
	err := ctx.Invoke("oci:integration/getIntegrationInstances:getIntegrationInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIntegrationInstances.
type GetIntegrationInstancesArgs struct {
	// The ID of the compartment in which to list resources.
	CompartmentId string `pulumi:"compartmentId"`
	// A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource`
	DisplayName *string                         `pulumi:"displayName"`
	Filters     []GetIntegrationInstancesFilter `pulumi:"filters"`
	// Life cycle state to query on.
	State *string `pulumi:"state"`
}

// A collection of values returned by getIntegrationInstances.
type GetIntegrationInstancesResult struct {
	// Compartment Identifier.
	CompartmentId string `pulumi:"compartmentId"`
	// Integration Instance Identifier, can be renamed.
	DisplayName *string                         `pulumi:"displayName"`
	Filters     []GetIntegrationInstancesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of integration_instances.
	IntegrationInstances []GetIntegrationInstancesIntegrationInstance `pulumi:"integrationInstances"`
	// The current state of the integration instance.
	State *string `pulumi:"state"`
}