// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifacts

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Generic Artifact resource in Oracle Cloud Infrastructure Artifacts service.
//
// Gets information about an artifact with a specified [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/artifacts"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifacts.LookupGenericArtifact(ctx, &artifacts.LookupGenericArtifactArgs{
// 			ArtifactId: oci_artifacts_artifact.Test_artifact.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupGenericArtifact(ctx *pulumi.Context, args *LookupGenericArtifactArgs, opts ...pulumi.InvokeOption) (*LookupGenericArtifactResult, error) {
	var rv LookupGenericArtifactResult
	err := ctx.Invoke("oci:artifacts/getGenericArtifact:getGenericArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGenericArtifact.
type LookupGenericArtifactArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the artifact.  Example: `ocid1.genericartifact.oc1..exampleuniqueID`
	ArtifactId string `pulumi:"artifactId"`
}

// A collection of values returned by getGenericArtifact.
type LookupGenericArtifactResult struct {
	ArtifactId string `pulumi:"artifactId"`
	// A user-defined path to describe the location of an artifact. Slashes do not create a directory structure, but you can use slashes to organize the repository. An artifact path does not include an artifact version.  Example: `project01/my-web-app/artifact-abc`
	ArtifactPath string `pulumi:"artifactPath"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository's compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The artifact name with the format of `<artifact-path>:<artifact-version>`. The artifact name is truncated to a maximum length of 255.  Example: `project01/my-web-app/artifact-abc:1.0.0`
	DisplayName string `pulumi:"displayName"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the artifact.  Example: `ocid1.genericartifact.oc1..exampleuniqueID`
	Id string `pulumi:"id"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository.
	RepositoryId string `pulumi:"repositoryId"`
	// The SHA256 digest for the artifact. When you upload an artifact to the repository, a SHA256 digest is calculated and added to the artifact properties.
	Sha256 string `pulumi:"sha256"`
	// The size of the artifact in bytes.
	SizeInBytes string `pulumi:"sizeInBytes"`
	// The current state of the artifact.
	State string `pulumi:"state"`
	// An RFC 3339 timestamp indicating when the repository was created.
	TimeCreated string `pulumi:"timeCreated"`
	// A user-defined string to describe the artifact version.  Example: `1.1.0` or `1.2-beta-2`
	Version string `pulumi:"version"`
}
