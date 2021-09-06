// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loganalytics

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Namespace resource in Oracle Cloud Infrastructure Log Analytics service.
//
// This API gets the namespace details of a tenancy already onboarded in Logging Analytics Application
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/loganalytics"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := loganalytics.LookupNamespace(ctx, &loganalytics.LookupNamespaceArgs{
// 			Namespace: _var.Namespace_namespace,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("oci:loganalytics/getNamespace:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNamespace.
type LookupNamespaceArgs struct {
	// The Logging Analytics namespace used for the request.
	Namespace string `pulumi:"namespace"`
}

// A collection of values returned by getNamespace.
type LookupNamespaceResult struct {
	// The is the tenancy ID
	CompartmentId string `pulumi:"compartmentId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// This indicates if the tenancy is onboarded to Logging Analytics
	IsOnboarded bool `pulumi:"isOnboarded"`
	// This is the namespace name of a tenancy
	Namespace string `pulumi:"namespace"`
}