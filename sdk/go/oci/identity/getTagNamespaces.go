// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Tag Namespaces in Oracle Cloud Infrastructure Identity service.
//
// Lists the tag namespaces in the specified compartment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/identity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Tag_namespace_include_subcompartments
// 		opt1 := _var.Tag_namespace_state
// 		_, err := identity.GetTagNamespaces(ctx, &identity.GetTagNamespacesArgs{
// 			CompartmentId:          _var.Compartment_id,
// 			IncludeSubcompartments: &opt0,
// 			State:                  &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetTagNamespaces(ctx *pulumi.Context, args *GetTagNamespacesArgs, opts ...pulumi.InvokeOption) (*GetTagNamespacesResult, error) {
	var rv GetTagNamespacesResult
	err := ctx.Invoke("oci:identity/getTagNamespaces:getTagNamespaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTagNamespaces.
type GetTagNamespacesArgs struct {
	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId string                   `pulumi:"compartmentId"`
	Filters       []GetTagNamespacesFilter `pulumi:"filters"`
	// An optional boolean parameter indicating whether to retrieve all tag namespaces in subcompartments. If this parameter is not specified, only the tag namespaces defined in the specified compartment are retrieved.
	IncludeSubcompartments *bool `pulumi:"includeSubcompartments"`
	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	State *string `pulumi:"state"`
}

// A collection of values returned by getTagNamespaces.
type GetTagNamespacesResult struct {
	// The OCID of the compartment that contains the tag namespace.
	CompartmentId string                   `pulumi:"compartmentId"`
	Filters       []GetTagNamespacesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string `pulumi:"id"`
	IncludeSubcompartments *bool  `pulumi:"includeSubcompartments"`
	// The tagnamespace's current state. After creating a tagnamespace, make sure its `lifecycleState` is ACTIVE before using it. After retiring a tagnamespace, make sure its `lifecycleState` is INACTIVE before using it.
	State *string `pulumi:"state"`
	// The list of tag_namespaces.
	TagNamespaces []GetTagNamespacesTagNamespace `pulumi:"tagNamespaces"`
}
