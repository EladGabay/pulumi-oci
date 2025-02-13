// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataintegration

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Workspaces in Oracle Cloud Infrastructure Data Integration service.
//
// Retrieves a list of Data Integration workspaces.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/dataintegration"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Workspace_name
// 		opt1 := _var.Workspace_state
// 		_, err := dataintegration.GetWorkspaces(ctx, &dataintegration.GetWorkspacesArgs{
// 			CompartmentId: _var.Compartment_id,
// 			Name:          &opt0,
// 			State:         &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetWorkspaces(ctx *pulumi.Context, args *GetWorkspacesArgs, opts ...pulumi.InvokeOption) (*GetWorkspacesResult, error) {
	var rv GetWorkspacesResult
	err := ctx.Invoke("oci:dataintegration/getWorkspaces:getWorkspaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWorkspaces.
type GetWorkspacesArgs struct {
	// The OCID of the compartment containing the resources you want to list.
	CompartmentId string                `pulumi:"compartmentId"`
	Filters       []GetWorkspacesFilter `pulumi:"filters"`
	// Used to filter by the name of the object.
	Name *string `pulumi:"name"`
	// The lifecycle state of a resource. When specified, the operation only returns resources that match the given lifecycle state. When not specified, all lifecycle states are processed as a match.
	State *string `pulumi:"state"`
}

// A collection of values returned by getWorkspaces.
type GetWorkspacesResult struct {
	// The OCID of the compartment that contains the workspace.
	CompartmentId string                `pulumi:"compartmentId"`
	Filters       []GetWorkspacesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
	// Lifecycle states for workspaces in Data Integration Service CREATING - The resource is being created and may not be usable until the entire metadata is defined UPDATING - The resource is being updated and may not be usable until all changes are commited DELETING - The resource is being deleted and might require deep cleanup of children. ACTIVE   - The resource is valid and available for access INACTIVE - The resource might be incomplete in its definition or might have been made unavailable for administrative reasons DELETED  - The resource has been deleted and isn't available FAILED   - The resource is in a failed state due to validation or other errors STARTING - The resource is being started and may not be usable until becomes ACTIVE again STOPPING - The resource is in the process of Stopping and may not be usable until it Stops or fails STOPPED  - The resource is in Stopped state due to stop operation.
	State *string `pulumi:"state"`
	// The list of workspaces.
	Workspaces []GetWorkspacesWorkspace `pulumi:"workspaces"`
}
