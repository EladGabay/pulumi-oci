// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Object Lifecycle Policy resource in Oracle Cloud Infrastructure Object Storage service.
//
// Gets the object lifecycle policy for the bucket.
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
// 		_, err := oci.GetObjectstorageObjectLifecyclePolicy(ctx, &GetObjectstorageObjectLifecyclePolicyArgs{
// 			Bucket:    _var.Object_lifecycle_policy_bucket,
// 			Namespace: _var.Object_lifecycle_policy_namespace,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupObjectstorageObjectLifecyclePolicy(ctx *pulumi.Context, args *LookupObjectstorageObjectLifecyclePolicyArgs, opts ...pulumi.InvokeOption) (*LookupObjectstorageObjectLifecyclePolicyResult, error) {
	var rv LookupObjectstorageObjectLifecyclePolicyResult
	err := ctx.Invoke("oci:index/getObjectstorageObjectLifecyclePolicy:GetObjectstorageObjectLifecyclePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetObjectstorageObjectLifecyclePolicy.
type LookupObjectstorageObjectLifecyclePolicyArgs struct {
	// The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1`
	Bucket string `pulumi:"bucket"`
	// The Object Storage namespace used for the request.
	Namespace string `pulumi:"namespace"`
}

// A collection of values returned by GetObjectstorageObjectLifecyclePolicy.
type LookupObjectstorageObjectLifecyclePolicyResult struct {
	Bucket    string `pulumi:"bucket"`
	Id        string `pulumi:"id"`
	Namespace string `pulumi:"namespace"`
	// The live lifecycle policy on the bucket.
	Rules []GetObjectstorageObjectLifecyclePolicyRule `pulumi:"rules"`
	// The date and time the object lifecycle policy was created, as described in [RFC 3339](https://tools.ietf.org/html/rfc3339).
	TimeCreated string `pulumi:"timeCreated"`
}