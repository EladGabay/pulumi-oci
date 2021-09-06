// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the External Non Container Database Operations Insights Management resource in Oracle Cloud Infrastructure Database service.
//
// Enable Operations Insights for the external non-container database.
// When deleting this resource block , we call disable if it was in enabled state .
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
// 		_, err := database.NewExternalNonContainerDatabaseOperationsInsightsManagement(ctx, "testExternalNonContainerDatabaseOperationsInsightsManagement", &database.ExternalNonContainerDatabaseOperationsInsightsManagementArgs{
// 			ExternalDatabaseConnectorId:    pulumi.Any(oci_database_external_database_connector.Test_external_database_connector.Id),
// 			ExternalNonContainerDatabaseId: pulumi.Any(oci_database_external_non_container_database.Test_external_non_container_database.Id),
// 			EnableOperationsInsights:       pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Import is not supported for this resource.
type ExternalNonContainerDatabaseOperationsInsightsManagement struct {
	pulumi.CustomResourceState

	// (Updatable) Enabling OPSI on External non-container Databases . Requires boolean value "true" or "false".
	EnableOperationsInsights pulumi.BoolOutput `pulumi:"enableOperationsInsights"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails).
	ExternalDatabaseConnectorId pulumi.StringOutput `pulumi:"externalDatabaseConnectorId"`
	// The external non-container database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExternalNonContainerDatabaseId pulumi.StringOutput `pulumi:"externalNonContainerDatabaseId"`
}

// NewExternalNonContainerDatabaseOperationsInsightsManagement registers a new resource with the given unique name, arguments, and options.
func NewExternalNonContainerDatabaseOperationsInsightsManagement(ctx *pulumi.Context,
	name string, args *ExternalNonContainerDatabaseOperationsInsightsManagementArgs, opts ...pulumi.ResourceOption) (*ExternalNonContainerDatabaseOperationsInsightsManagement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnableOperationsInsights == nil {
		return nil, errors.New("invalid value for required argument 'EnableOperationsInsights'")
	}
	if args.ExternalDatabaseConnectorId == nil {
		return nil, errors.New("invalid value for required argument 'ExternalDatabaseConnectorId'")
	}
	if args.ExternalNonContainerDatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'ExternalNonContainerDatabaseId'")
	}
	var resource ExternalNonContainerDatabaseOperationsInsightsManagement
	err := ctx.RegisterResource("oci:database/externalNonContainerDatabaseOperationsInsightsManagement:ExternalNonContainerDatabaseOperationsInsightsManagement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalNonContainerDatabaseOperationsInsightsManagement gets an existing ExternalNonContainerDatabaseOperationsInsightsManagement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalNonContainerDatabaseOperationsInsightsManagement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalNonContainerDatabaseOperationsInsightsManagementState, opts ...pulumi.ResourceOption) (*ExternalNonContainerDatabaseOperationsInsightsManagement, error) {
	var resource ExternalNonContainerDatabaseOperationsInsightsManagement
	err := ctx.ReadResource("oci:database/externalNonContainerDatabaseOperationsInsightsManagement:ExternalNonContainerDatabaseOperationsInsightsManagement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalNonContainerDatabaseOperationsInsightsManagement resources.
type externalNonContainerDatabaseOperationsInsightsManagementState struct {
	// (Updatable) Enabling OPSI on External non-container Databases . Requires boolean value "true" or "false".
	EnableOperationsInsights *bool `pulumi:"enableOperationsInsights"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails).
	ExternalDatabaseConnectorId *string `pulumi:"externalDatabaseConnectorId"`
	// The external non-container database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExternalNonContainerDatabaseId *string `pulumi:"externalNonContainerDatabaseId"`
}

type ExternalNonContainerDatabaseOperationsInsightsManagementState struct {
	// (Updatable) Enabling OPSI on External non-container Databases . Requires boolean value "true" or "false".
	EnableOperationsInsights pulumi.BoolPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails).
	ExternalDatabaseConnectorId pulumi.StringPtrInput
	// The external non-container database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExternalNonContainerDatabaseId pulumi.StringPtrInput
}

func (ExternalNonContainerDatabaseOperationsInsightsManagementState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalNonContainerDatabaseOperationsInsightsManagementState)(nil)).Elem()
}

type externalNonContainerDatabaseOperationsInsightsManagementArgs struct {
	// (Updatable) Enabling OPSI on External non-container Databases . Requires boolean value "true" or "false".
	EnableOperationsInsights bool `pulumi:"enableOperationsInsights"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails).
	ExternalDatabaseConnectorId string `pulumi:"externalDatabaseConnectorId"`
	// The external non-container database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExternalNonContainerDatabaseId string `pulumi:"externalNonContainerDatabaseId"`
}

// The set of arguments for constructing a ExternalNonContainerDatabaseOperationsInsightsManagement resource.
type ExternalNonContainerDatabaseOperationsInsightsManagementArgs struct {
	// (Updatable) Enabling OPSI on External non-container Databases . Requires boolean value "true" or "false".
	EnableOperationsInsights pulumi.BoolInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails).
	ExternalDatabaseConnectorId pulumi.StringInput
	// The external non-container database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExternalNonContainerDatabaseId pulumi.StringInput
}

func (ExternalNonContainerDatabaseOperationsInsightsManagementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalNonContainerDatabaseOperationsInsightsManagementArgs)(nil)).Elem()
}

type ExternalNonContainerDatabaseOperationsInsightsManagementInput interface {
	pulumi.Input

	ToExternalNonContainerDatabaseOperationsInsightsManagementOutput() ExternalNonContainerDatabaseOperationsInsightsManagementOutput
	ToExternalNonContainerDatabaseOperationsInsightsManagementOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementOutput
}

func (*ExternalNonContainerDatabaseOperationsInsightsManagement) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalNonContainerDatabaseOperationsInsightsManagement)(nil))
}

func (i *ExternalNonContainerDatabaseOperationsInsightsManagement) ToExternalNonContainerDatabaseOperationsInsightsManagementOutput() ExternalNonContainerDatabaseOperationsInsightsManagementOutput {
	return i.ToExternalNonContainerDatabaseOperationsInsightsManagementOutputWithContext(context.Background())
}

func (i *ExternalNonContainerDatabaseOperationsInsightsManagement) ToExternalNonContainerDatabaseOperationsInsightsManagementOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalNonContainerDatabaseOperationsInsightsManagementOutput)
}

func (i *ExternalNonContainerDatabaseOperationsInsightsManagement) ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput() ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput {
	return i.ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutputWithContext(context.Background())
}

func (i *ExternalNonContainerDatabaseOperationsInsightsManagement) ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput)
}

type ExternalNonContainerDatabaseOperationsInsightsManagementPtrInput interface {
	pulumi.Input

	ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput() ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput
	ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput
}

type externalNonContainerDatabaseOperationsInsightsManagementPtrType ExternalNonContainerDatabaseOperationsInsightsManagementArgs

func (*externalNonContainerDatabaseOperationsInsightsManagementPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalNonContainerDatabaseOperationsInsightsManagement)(nil))
}

func (i *externalNonContainerDatabaseOperationsInsightsManagementPtrType) ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput() ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput {
	return i.ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutputWithContext(context.Background())
}

func (i *externalNonContainerDatabaseOperationsInsightsManagementPtrType) ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput)
}

// ExternalNonContainerDatabaseOperationsInsightsManagementArrayInput is an input type that accepts ExternalNonContainerDatabaseOperationsInsightsManagementArray and ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput values.
// You can construct a concrete instance of `ExternalNonContainerDatabaseOperationsInsightsManagementArrayInput` via:
//
//          ExternalNonContainerDatabaseOperationsInsightsManagementArray{ ExternalNonContainerDatabaseOperationsInsightsManagementArgs{...} }
type ExternalNonContainerDatabaseOperationsInsightsManagementArrayInput interface {
	pulumi.Input

	ToExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput() ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput
	ToExternalNonContainerDatabaseOperationsInsightsManagementArrayOutputWithContext(context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput
}

type ExternalNonContainerDatabaseOperationsInsightsManagementArray []ExternalNonContainerDatabaseOperationsInsightsManagementInput

func (ExternalNonContainerDatabaseOperationsInsightsManagementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalNonContainerDatabaseOperationsInsightsManagement)(nil)).Elem()
}

func (i ExternalNonContainerDatabaseOperationsInsightsManagementArray) ToExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput() ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput {
	return i.ToExternalNonContainerDatabaseOperationsInsightsManagementArrayOutputWithContext(context.Background())
}

func (i ExternalNonContainerDatabaseOperationsInsightsManagementArray) ToExternalNonContainerDatabaseOperationsInsightsManagementArrayOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput)
}

// ExternalNonContainerDatabaseOperationsInsightsManagementMapInput is an input type that accepts ExternalNonContainerDatabaseOperationsInsightsManagementMap and ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput values.
// You can construct a concrete instance of `ExternalNonContainerDatabaseOperationsInsightsManagementMapInput` via:
//
//          ExternalNonContainerDatabaseOperationsInsightsManagementMap{ "key": ExternalNonContainerDatabaseOperationsInsightsManagementArgs{...} }
type ExternalNonContainerDatabaseOperationsInsightsManagementMapInput interface {
	pulumi.Input

	ToExternalNonContainerDatabaseOperationsInsightsManagementMapOutput() ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput
	ToExternalNonContainerDatabaseOperationsInsightsManagementMapOutputWithContext(context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput
}

type ExternalNonContainerDatabaseOperationsInsightsManagementMap map[string]ExternalNonContainerDatabaseOperationsInsightsManagementInput

func (ExternalNonContainerDatabaseOperationsInsightsManagementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalNonContainerDatabaseOperationsInsightsManagement)(nil)).Elem()
}

func (i ExternalNonContainerDatabaseOperationsInsightsManagementMap) ToExternalNonContainerDatabaseOperationsInsightsManagementMapOutput() ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput {
	return i.ToExternalNonContainerDatabaseOperationsInsightsManagementMapOutputWithContext(context.Background())
}

func (i ExternalNonContainerDatabaseOperationsInsightsManagementMap) ToExternalNonContainerDatabaseOperationsInsightsManagementMapOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput)
}

type ExternalNonContainerDatabaseOperationsInsightsManagementOutput struct {
	*pulumi.OutputState
}

func (ExternalNonContainerDatabaseOperationsInsightsManagementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalNonContainerDatabaseOperationsInsightsManagement)(nil))
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementOutput) ToExternalNonContainerDatabaseOperationsInsightsManagementOutput() ExternalNonContainerDatabaseOperationsInsightsManagementOutput {
	return o
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementOutput) ToExternalNonContainerDatabaseOperationsInsightsManagementOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementOutput {
	return o
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementOutput) ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput() ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput {
	return o.ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutputWithContext(context.Background())
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementOutput) ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput {
	return o.ApplyT(func(v ExternalNonContainerDatabaseOperationsInsightsManagement) *ExternalNonContainerDatabaseOperationsInsightsManagement {
		return &v
	}).(ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput)
}

type ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput struct {
	*pulumi.OutputState
}

func (ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalNonContainerDatabaseOperationsInsightsManagement)(nil))
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput) ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput() ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput {
	return o
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput) ToExternalNonContainerDatabaseOperationsInsightsManagementPtrOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput {
	return o
}

type ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput struct{ *pulumi.OutputState }

func (ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExternalNonContainerDatabaseOperationsInsightsManagement)(nil))
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput) ToExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput() ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput {
	return o
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput) ToExternalNonContainerDatabaseOperationsInsightsManagementArrayOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput {
	return o
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput) Index(i pulumi.IntInput) ExternalNonContainerDatabaseOperationsInsightsManagementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExternalNonContainerDatabaseOperationsInsightsManagement {
		return vs[0].([]ExternalNonContainerDatabaseOperationsInsightsManagement)[vs[1].(int)]
	}).(ExternalNonContainerDatabaseOperationsInsightsManagementOutput)
}

type ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput struct{ *pulumi.OutputState }

func (ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ExternalNonContainerDatabaseOperationsInsightsManagement)(nil))
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput) ToExternalNonContainerDatabaseOperationsInsightsManagementMapOutput() ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput {
	return o
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput) ToExternalNonContainerDatabaseOperationsInsightsManagementMapOutputWithContext(ctx context.Context) ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput {
	return o
}

func (o ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput) MapIndex(k pulumi.StringInput) ExternalNonContainerDatabaseOperationsInsightsManagementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ExternalNonContainerDatabaseOperationsInsightsManagement {
		return vs[0].(map[string]ExternalNonContainerDatabaseOperationsInsightsManagement)[vs[1].(string)]
	}).(ExternalNonContainerDatabaseOperationsInsightsManagementOutput)
}

func init() {
	pulumi.RegisterOutputType(ExternalNonContainerDatabaseOperationsInsightsManagementOutput{})
	pulumi.RegisterOutputType(ExternalNonContainerDatabaseOperationsInsightsManagementPtrOutput{})
	pulumi.RegisterOutputType(ExternalNonContainerDatabaseOperationsInsightsManagementArrayOutput{})
	pulumi.RegisterOutputType(ExternalNonContainerDatabaseOperationsInsightsManagementMapOutput{})
}