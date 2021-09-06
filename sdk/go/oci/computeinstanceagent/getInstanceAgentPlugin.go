// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package computeinstanceagent

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Instance Agent Plugin resource in Oracle Cloud Infrastructure Compute Instance Agent service.
//
// The API to get information for a plugin.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/computeinstanceagent"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := computeinstanceagent.GetInstanceAgentPlugin(ctx, &computeinstanceagent.GetInstanceAgentPluginArgs{
// 			InstanceagentId: oci_computeinstanceagent_instanceagent.Test_instanceagent.Id,
// 			PluginName:      _var.Instance_agent_plugin_plugin_name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetInstanceAgentPlugin(ctx *pulumi.Context, args *GetInstanceAgentPluginArgs, opts ...pulumi.InvokeOption) (*GetInstanceAgentPluginResult, error) {
	var rv GetInstanceAgentPluginResult
	err := ctx.Invoke("oci:computeinstanceagent/getInstanceAgentPlugin:getInstanceAgentPlugin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceAgentPlugin.
type GetInstanceAgentPluginArgs struct {
	CompartmentId string `pulumi:"compartmentId"`
	// The OCID of the instance.
	InstanceagentId string `pulumi:"instanceagentId"`
	// The name of the plugin.
	PluginName string `pulumi:"pluginName"`
}

// A collection of values returned by getInstanceAgentPlugin.
type GetInstanceAgentPluginResult struct {
	CompartmentId string `pulumi:"compartmentId"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	InstanceagentId string `pulumi:"instanceagentId"`
	// The optional message from the agent plugin
	Message string `pulumi:"message"`
	// The plugin name
	Name       string `pulumi:"name"`
	PluginName string `pulumi:"pluginName"`
	// The plugin status Specified the plugin state on the instance * `RUNNING` - The plugin is in running state * `STOPPED` - The plugin is in stopped state * `NOT_SUPPORTED` - The plugin is not supported on this platform * `INVALID` - The plugin state is not recognizable by the service
	Status string `pulumi:"status"`
	// The last update time of the plugin in UTC
	TimeLastUpdatedUtc string `pulumi:"timeLastUpdatedUtc"`
}