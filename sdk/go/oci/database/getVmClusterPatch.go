// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Vm Cluster Patch resource in Oracle Cloud Infrastructure Database service.
//
// Gets information about a specified patch package.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/database"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := database.GetVmClusterPatch(ctx, &database.GetVmClusterPatchArgs{
// 			PatchId:     oci_database_patch.Test_patch.Id,
// 			VmClusterId: oci_database_vm_cluster.Test_vm_cluster.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetVmClusterPatch(ctx *pulumi.Context, args *GetVmClusterPatchArgs, opts ...pulumi.InvokeOption) (*GetVmClusterPatchResult, error) {
	var rv GetVmClusterPatchResult
	err := ctx.Invoke("oci:database/getVmClusterPatch:getVmClusterPatch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVmClusterPatch.
type GetVmClusterPatchArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch.
	PatchId string `pulumi:"patchId"`
	// The VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	VmClusterId string `pulumi:"vmClusterId"`
}

// A collection of values returned by getVmClusterPatch.
type GetVmClusterPatchResult struct {
	// Actions that can possibly be performed using this patch.
	AvailableActions []string `pulumi:"availableActions"`
	// The text describing this patch package.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Action that is currently being performed or was completed last.
	LastAction string `pulumi:"lastAction"`
	// A descriptive text associated with the lifecycleState. Typically can contain additional displayable text.
	LifecycleDetails string `pulumi:"lifecycleDetails"`
	PatchId          string `pulumi:"patchId"`
	// The current state of the patch as a result of lastAction.
	State string `pulumi:"state"`
	// The date and time that the patch was released.
	TimeReleased string `pulumi:"timeReleased"`
	// The version of this patch package.
	Version     string `pulumi:"version"`
	VmClusterId string `pulumi:"vmClusterId"`
}