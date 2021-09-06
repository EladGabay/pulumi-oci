// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Backend Set Health resource in Oracle Cloud Infrastructure Load Balancer service.
//
// Gets the health status for the specified backend set.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/loadbalancer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := loadbalancer.GetBackendSetHealth(ctx, &loadbalancer.GetBackendSetHealthArgs{
// 			BackendSetName: oci_load_balancer_backend_set.Test_backend_set.Name,
// 			LoadBalancerId: oci_load_balancer_load_balancer.Test_load_balancer.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetBackendSetHealth(ctx *pulumi.Context, args *GetBackendSetHealthArgs, opts ...pulumi.InvokeOption) (*GetBackendSetHealthResult, error) {
	var rv GetBackendSetHealthResult
	err := ctx.Invoke("oci:loadbalancer/getBackendSetHealth:getBackendSetHealth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackendSetHealth.
type GetBackendSetHealthArgs struct {
	// The name of the backend set to retrieve the health status for.  Example: `exampleBackendSet`
	BackendSetName string `pulumi:"backendSetName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend set health status to be retrieved.
	LoadBalancerId string `pulumi:"loadBalancerId"`
}

// A collection of values returned by getBackendSetHealth.
type GetBackendSetHealthResult struct {
	BackendSetName string `pulumi:"backendSetName"`
	// A list of backend servers that are currently in the `CRITICAL` health state. The list identifies each backend server by IP address and port.  Example: `10.0.0.4:8080`
	CriticalStateBackendNames []string `pulumi:"criticalStateBackendNames"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// Overall health status of the backend set.
	// *  **OK:** All backend servers in the backend set return a status of `OK`.
	// *  **WARNING:** Half or more of the backend set's backend servers return a status of `OK` and at least one backend server returns a status of `WARNING`, `CRITICAL`, or `UNKNOWN`.
	// *  **CRITICAL:** Fewer than half of the backend set's backend servers return a status of `OK`.
	// *  **UNKNOWN:** More than half of the backend set's backend servers return a status of `UNKNOWN`, the system was unable to retrieve metrics, or the backend set does not have a listener attached.
	Status string `pulumi:"status"`
	// The total number of backend servers in this backend set.  Example: `7`
	TotalBackendCount int `pulumi:"totalBackendCount"`
	// A list of backend servers that are currently in the `UNKNOWN` health state. The list identifies each backend server by IP address and port.  Example: `10.0.0.5:8080`
	UnknownStateBackendNames []string `pulumi:"unknownStateBackendNames"`
	// A list of backend servers that are currently in the `WARNING` health state. The list identifies each backend server by IP address and port.  Example: `10.0.0.3:8080`
	WarningStateBackendNames []string `pulumi:"warningStateBackendNames"`
}