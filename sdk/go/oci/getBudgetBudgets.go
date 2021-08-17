// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Budgets in Oracle Cloud Infrastructure Budget service.
//
// Gets a list of Budgets in a compartment.
//
// By default, ListBudgets returns budgets of 'COMPARTMENT' target type and the budget records with only ONE target compartment OCID.
//
// To list ALL budgets, set the targetType query parameter to ALL.
// Example:
//   'targetType=ALL'
//
// Additional targetTypes would be available in future releases. Clients should ignore new targetType
// or upgrade to latest version of client SDK to handle new targetType.
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
// 		opt0 := _var.Budget_display_name
// 		opt1 := _var.Budget_state
// 		opt2 := _var.Budget_target_type
// 		_, err := oci.GetBudgetBudgets(ctx, &GetBudgetBudgetsArgs{
// 			CompartmentId: _var.Tenancy_ocid,
// 			DisplayName:   &opt0,
// 			State:         &opt1,
// 			TargetType:    &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetBudgetBudgets(ctx *pulumi.Context, args *GetBudgetBudgetsArgs, opts ...pulumi.InvokeOption) (*GetBudgetBudgetsResult, error) {
	var rv GetBudgetBudgetsResult
	err := ctx.Invoke("oci:index/getBudgetBudgets:GetBudgetBudgets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetBudgetBudgets.
type GetBudgetBudgetsArgs struct {
	// The ID of the compartment in which to list resources.
	CompartmentId string `pulumi:"compartmentId"`
	// A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource`
	DisplayName *string                  `pulumi:"displayName"`
	Filters     []GetBudgetBudgetsFilter `pulumi:"filters"`
	// The current state of the resource to filter by.
	State *string `pulumi:"state"`
	// The type of target to filter by.
	// * ALL - List all budgets
	// * COMPARTMENT - List all budgets with targetType == "COMPARTMENT"
	// * TAG - List all budgets with targetType == "TAG"
	TargetType *string `pulumi:"targetType"`
}

// A collection of values returned by GetBudgetBudgets.
type GetBudgetBudgetsResult struct {
	// The list of budgets.
	Budgets []GetBudgetBudgetsBudget `pulumi:"budgets"`
	// The OCID of the compartment
	CompartmentId string `pulumi:"compartmentId"`
	// The display name of the budget.
	DisplayName *string                  `pulumi:"displayName"`
	Filters     []GetBudgetBudgetsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current state of the budget.
	State *string `pulumi:"state"`
	// The type of target on which the budget is applied.
	TargetType *string `pulumi:"targetType"`
}