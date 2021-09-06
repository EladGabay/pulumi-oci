// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the External Database Connector resource in Oracle Cloud Infrastructure Database service.
//
// Creates a new external database connector.
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
// 		_, err := database.NewExternalDatabaseConnector(ctx, "testExternalDatabaseConnector", &database.ExternalDatabaseConnectorArgs{
// 			ConnectionCredentials: &database.ExternalDatabaseConnectorConnectionCredentialsArgs{
// 				CredentialName: pulumi.Any(_var.External_database_connector_connection_credentials_credential_name),
// 				CredentialType: pulumi.Any(_var.External_database_connector_connection_credentials_credential_type),
// 				Password:       pulumi.Any(_var.External_database_connector_connection_credentials_password),
// 				Role:           pulumi.Any(_var.External_database_connector_connection_credentials_role),
// 				Username:       pulumi.Any(_var.External_database_connector_connection_credentials_username),
// 			},
// 			ConnectionString: &database.ExternalDatabaseConnectorConnectionStringArgs{
// 				Hostname: pulumi.Any(_var.External_database_connector_connection_string_hostname),
// 				Port:     pulumi.Any(_var.External_database_connector_connection_string_port),
// 				Protocol: pulumi.Any(_var.External_database_connector_connection_string_protocol),
// 				Service:  pulumi.Any(_var.External_database_connector_connection_string_service),
// 			},
// 			ConnectorAgentId:   pulumi.Any(oci_database_connector_agent.Test_connector_agent.Id),
// 			DisplayName:        pulumi.Any(_var.External_database_connector_display_name),
// 			ExternalDatabaseId: pulumi.Any(oci_database_database.Test_database.Id),
// 			ConnectorType:      pulumi.Any(_var.External_database_connector_connector_type),
// 			DefinedTags:        pulumi.Any(_var.External_database_connector_defined_tags),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
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
// ExternalDatabaseConnectors can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:database/externalDatabaseConnector:ExternalDatabaseConnector test_external_database_connector "id"
// ```
type ExternalDatabaseConnector struct {
	pulumi.CustomResourceState

	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector crendentials.
	ConnectionCredentials ExternalDatabaseConnectorConnectionCredentialsOutput `pulumi:"connectionCredentials"`
	// The status of connectivity to the external database.
	ConnectionStatus pulumi.StringOutput `pulumi:"connectionStatus"`
	// (Updatable) The Oracle Database connection string.
	ConnectionString ExternalDatabaseConnectorConnectionStringOutput `pulumi:"connectionString"`
	// The ID of the agent used for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails).
	ConnectorAgentId pulumi.StringOutput `pulumi:"connectorAgentId"`
	// (Updatable) The type of connector used by the external database resource.
	ConnectorType pulumi.StringOutput `pulumi:"connectorType"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) The user-friendly name for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). The name does not have to be unique.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database resource.
	ExternalDatabaseId pulumi.StringOutput `pulumi:"externalDatabaseId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// Additional information about the current lifecycle state.
	LifecycleDetails pulumi.StringOutput `pulumi:"lifecycleDetails"`
	// The current lifecycle state of the external database connector resource.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the `connectionStatus` of this external connector was last updated.
	TimeConnectionStatusLastUpdated pulumi.StringOutput `pulumi:"timeConnectionStatusLastUpdated"`
	// The date and time the external connector was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
}

// NewExternalDatabaseConnector registers a new resource with the given unique name, arguments, and options.
func NewExternalDatabaseConnector(ctx *pulumi.Context,
	name string, args *ExternalDatabaseConnectorArgs, opts ...pulumi.ResourceOption) (*ExternalDatabaseConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionCredentials == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionCredentials'")
	}
	if args.ConnectionString == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionString'")
	}
	if args.ConnectorAgentId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorAgentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ExternalDatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'ExternalDatabaseId'")
	}
	var resource ExternalDatabaseConnector
	err := ctx.RegisterResource("oci:database/externalDatabaseConnector:ExternalDatabaseConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalDatabaseConnector gets an existing ExternalDatabaseConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalDatabaseConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalDatabaseConnectorState, opts ...pulumi.ResourceOption) (*ExternalDatabaseConnector, error) {
	var resource ExternalDatabaseConnector
	err := ctx.ReadResource("oci:database/externalDatabaseConnector:ExternalDatabaseConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalDatabaseConnector resources.
type externalDatabaseConnectorState struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector crendentials.
	ConnectionCredentials *ExternalDatabaseConnectorConnectionCredentials `pulumi:"connectionCredentials"`
	// The status of connectivity to the external database.
	ConnectionStatus *string `pulumi:"connectionStatus"`
	// (Updatable) The Oracle Database connection string.
	ConnectionString *ExternalDatabaseConnectorConnectionString `pulumi:"connectionString"`
	// The ID of the agent used for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails).
	ConnectorAgentId *string `pulumi:"connectorAgentId"`
	// (Updatable) The type of connector used by the external database resource.
	ConnectorType *string `pulumi:"connectorType"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) The user-friendly name for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). The name does not have to be unique.
	DisplayName *string `pulumi:"displayName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database resource.
	ExternalDatabaseId *string `pulumi:"externalDatabaseId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Additional information about the current lifecycle state.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// The current lifecycle state of the external database connector resource.
	State *string `pulumi:"state"`
	// The date and time the `connectionStatus` of this external connector was last updated.
	TimeConnectionStatusLastUpdated *string `pulumi:"timeConnectionStatusLastUpdated"`
	// The date and time the external connector was created.
	TimeCreated *string `pulumi:"timeCreated"`
}

type ExternalDatabaseConnectorState struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector crendentials.
	ConnectionCredentials ExternalDatabaseConnectorConnectionCredentialsPtrInput
	// The status of connectivity to the external database.
	ConnectionStatus pulumi.StringPtrInput
	// (Updatable) The Oracle Database connection string.
	ConnectionString ExternalDatabaseConnectorConnectionStringPtrInput
	// The ID of the agent used for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails).
	ConnectorAgentId pulumi.StringPtrInput
	// (Updatable) The type of connector used by the external database resource.
	ConnectorType pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags pulumi.MapInput
	// (Updatable) The user-friendly name for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). The name does not have to be unique.
	DisplayName pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database resource.
	ExternalDatabaseId pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// Additional information about the current lifecycle state.
	LifecycleDetails pulumi.StringPtrInput
	// The current lifecycle state of the external database connector resource.
	State pulumi.StringPtrInput
	// The date and time the `connectionStatus` of this external connector was last updated.
	TimeConnectionStatusLastUpdated pulumi.StringPtrInput
	// The date and time the external connector was created.
	TimeCreated pulumi.StringPtrInput
}

func (ExternalDatabaseConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalDatabaseConnectorState)(nil)).Elem()
}

type externalDatabaseConnectorArgs struct {
	// (Updatable) Credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector crendentials.
	ConnectionCredentials ExternalDatabaseConnectorConnectionCredentials `pulumi:"connectionCredentials"`
	// (Updatable) The Oracle Database connection string.
	ConnectionString ExternalDatabaseConnectorConnectionString `pulumi:"connectionString"`
	// The ID of the agent used for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails).
	ConnectorAgentId string `pulumi:"connectorAgentId"`
	// (Updatable) The type of connector used by the external database resource.
	ConnectorType *string `pulumi:"connectorType"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) The user-friendly name for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). The name does not have to be unique.
	DisplayName string `pulumi:"displayName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database resource.
	ExternalDatabaseId string `pulumi:"externalDatabaseId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
}

// The set of arguments for constructing a ExternalDatabaseConnector resource.
type ExternalDatabaseConnectorArgs struct {
	// (Updatable) Credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector crendentials.
	ConnectionCredentials ExternalDatabaseConnectorConnectionCredentialsInput
	// (Updatable) The Oracle Database connection string.
	ConnectionString ExternalDatabaseConnectorConnectionStringInput
	// The ID of the agent used for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails).
	ConnectorAgentId pulumi.StringInput
	// (Updatable) The type of connector used by the external database resource.
	ConnectorType pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags pulumi.MapInput
	// (Updatable) The user-friendly name for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). The name does not have to be unique.
	DisplayName pulumi.StringInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database resource.
	ExternalDatabaseId pulumi.StringInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
}

func (ExternalDatabaseConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalDatabaseConnectorArgs)(nil)).Elem()
}

type ExternalDatabaseConnectorInput interface {
	pulumi.Input

	ToExternalDatabaseConnectorOutput() ExternalDatabaseConnectorOutput
	ToExternalDatabaseConnectorOutputWithContext(ctx context.Context) ExternalDatabaseConnectorOutput
}

func (*ExternalDatabaseConnector) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalDatabaseConnector)(nil))
}

func (i *ExternalDatabaseConnector) ToExternalDatabaseConnectorOutput() ExternalDatabaseConnectorOutput {
	return i.ToExternalDatabaseConnectorOutputWithContext(context.Background())
}

func (i *ExternalDatabaseConnector) ToExternalDatabaseConnectorOutputWithContext(ctx context.Context) ExternalDatabaseConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalDatabaseConnectorOutput)
}

func (i *ExternalDatabaseConnector) ToExternalDatabaseConnectorPtrOutput() ExternalDatabaseConnectorPtrOutput {
	return i.ToExternalDatabaseConnectorPtrOutputWithContext(context.Background())
}

func (i *ExternalDatabaseConnector) ToExternalDatabaseConnectorPtrOutputWithContext(ctx context.Context) ExternalDatabaseConnectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalDatabaseConnectorPtrOutput)
}

type ExternalDatabaseConnectorPtrInput interface {
	pulumi.Input

	ToExternalDatabaseConnectorPtrOutput() ExternalDatabaseConnectorPtrOutput
	ToExternalDatabaseConnectorPtrOutputWithContext(ctx context.Context) ExternalDatabaseConnectorPtrOutput
}

type externalDatabaseConnectorPtrType ExternalDatabaseConnectorArgs

func (*externalDatabaseConnectorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalDatabaseConnector)(nil))
}

func (i *externalDatabaseConnectorPtrType) ToExternalDatabaseConnectorPtrOutput() ExternalDatabaseConnectorPtrOutput {
	return i.ToExternalDatabaseConnectorPtrOutputWithContext(context.Background())
}

func (i *externalDatabaseConnectorPtrType) ToExternalDatabaseConnectorPtrOutputWithContext(ctx context.Context) ExternalDatabaseConnectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalDatabaseConnectorPtrOutput)
}

// ExternalDatabaseConnectorArrayInput is an input type that accepts ExternalDatabaseConnectorArray and ExternalDatabaseConnectorArrayOutput values.
// You can construct a concrete instance of `ExternalDatabaseConnectorArrayInput` via:
//
//          ExternalDatabaseConnectorArray{ ExternalDatabaseConnectorArgs{...} }
type ExternalDatabaseConnectorArrayInput interface {
	pulumi.Input

	ToExternalDatabaseConnectorArrayOutput() ExternalDatabaseConnectorArrayOutput
	ToExternalDatabaseConnectorArrayOutputWithContext(context.Context) ExternalDatabaseConnectorArrayOutput
}

type ExternalDatabaseConnectorArray []ExternalDatabaseConnectorInput

func (ExternalDatabaseConnectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalDatabaseConnector)(nil)).Elem()
}

func (i ExternalDatabaseConnectorArray) ToExternalDatabaseConnectorArrayOutput() ExternalDatabaseConnectorArrayOutput {
	return i.ToExternalDatabaseConnectorArrayOutputWithContext(context.Background())
}

func (i ExternalDatabaseConnectorArray) ToExternalDatabaseConnectorArrayOutputWithContext(ctx context.Context) ExternalDatabaseConnectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalDatabaseConnectorArrayOutput)
}

// ExternalDatabaseConnectorMapInput is an input type that accepts ExternalDatabaseConnectorMap and ExternalDatabaseConnectorMapOutput values.
// You can construct a concrete instance of `ExternalDatabaseConnectorMapInput` via:
//
//          ExternalDatabaseConnectorMap{ "key": ExternalDatabaseConnectorArgs{...} }
type ExternalDatabaseConnectorMapInput interface {
	pulumi.Input

	ToExternalDatabaseConnectorMapOutput() ExternalDatabaseConnectorMapOutput
	ToExternalDatabaseConnectorMapOutputWithContext(context.Context) ExternalDatabaseConnectorMapOutput
}

type ExternalDatabaseConnectorMap map[string]ExternalDatabaseConnectorInput

func (ExternalDatabaseConnectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalDatabaseConnector)(nil)).Elem()
}

func (i ExternalDatabaseConnectorMap) ToExternalDatabaseConnectorMapOutput() ExternalDatabaseConnectorMapOutput {
	return i.ToExternalDatabaseConnectorMapOutputWithContext(context.Background())
}

func (i ExternalDatabaseConnectorMap) ToExternalDatabaseConnectorMapOutputWithContext(ctx context.Context) ExternalDatabaseConnectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalDatabaseConnectorMapOutput)
}

type ExternalDatabaseConnectorOutput struct {
	*pulumi.OutputState
}

func (ExternalDatabaseConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalDatabaseConnector)(nil))
}

func (o ExternalDatabaseConnectorOutput) ToExternalDatabaseConnectorOutput() ExternalDatabaseConnectorOutput {
	return o
}

func (o ExternalDatabaseConnectorOutput) ToExternalDatabaseConnectorOutputWithContext(ctx context.Context) ExternalDatabaseConnectorOutput {
	return o
}

func (o ExternalDatabaseConnectorOutput) ToExternalDatabaseConnectorPtrOutput() ExternalDatabaseConnectorPtrOutput {
	return o.ToExternalDatabaseConnectorPtrOutputWithContext(context.Background())
}

func (o ExternalDatabaseConnectorOutput) ToExternalDatabaseConnectorPtrOutputWithContext(ctx context.Context) ExternalDatabaseConnectorPtrOutput {
	return o.ApplyT(func(v ExternalDatabaseConnector) *ExternalDatabaseConnector {
		return &v
	}).(ExternalDatabaseConnectorPtrOutput)
}

type ExternalDatabaseConnectorPtrOutput struct {
	*pulumi.OutputState
}

func (ExternalDatabaseConnectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalDatabaseConnector)(nil))
}

func (o ExternalDatabaseConnectorPtrOutput) ToExternalDatabaseConnectorPtrOutput() ExternalDatabaseConnectorPtrOutput {
	return o
}

func (o ExternalDatabaseConnectorPtrOutput) ToExternalDatabaseConnectorPtrOutputWithContext(ctx context.Context) ExternalDatabaseConnectorPtrOutput {
	return o
}

type ExternalDatabaseConnectorArrayOutput struct{ *pulumi.OutputState }

func (ExternalDatabaseConnectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExternalDatabaseConnector)(nil))
}

func (o ExternalDatabaseConnectorArrayOutput) ToExternalDatabaseConnectorArrayOutput() ExternalDatabaseConnectorArrayOutput {
	return o
}

func (o ExternalDatabaseConnectorArrayOutput) ToExternalDatabaseConnectorArrayOutputWithContext(ctx context.Context) ExternalDatabaseConnectorArrayOutput {
	return o
}

func (o ExternalDatabaseConnectorArrayOutput) Index(i pulumi.IntInput) ExternalDatabaseConnectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExternalDatabaseConnector {
		return vs[0].([]ExternalDatabaseConnector)[vs[1].(int)]
	}).(ExternalDatabaseConnectorOutput)
}

type ExternalDatabaseConnectorMapOutput struct{ *pulumi.OutputState }

func (ExternalDatabaseConnectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ExternalDatabaseConnector)(nil))
}

func (o ExternalDatabaseConnectorMapOutput) ToExternalDatabaseConnectorMapOutput() ExternalDatabaseConnectorMapOutput {
	return o
}

func (o ExternalDatabaseConnectorMapOutput) ToExternalDatabaseConnectorMapOutputWithContext(ctx context.Context) ExternalDatabaseConnectorMapOutput {
	return o
}

func (o ExternalDatabaseConnectorMapOutput) MapIndex(k pulumi.StringInput) ExternalDatabaseConnectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ExternalDatabaseConnector {
		return vs[0].(map[string]ExternalDatabaseConnector)[vs[1].(string)]
	}).(ExternalDatabaseConnectorOutput)
}

func init() {
	pulumi.RegisterOutputType(ExternalDatabaseConnectorOutput{})
	pulumi.RegisterOutputType(ExternalDatabaseConnectorPtrOutput{})
	pulumi.RegisterOutputType(ExternalDatabaseConnectorArrayOutput{})
	pulumi.RegisterOutputType(ExternalDatabaseConnectorMapOutput{})
}