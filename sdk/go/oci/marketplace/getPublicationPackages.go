// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package marketplace

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Publication Packages in Oracle Cloud Infrastructure Marketplace service.
//
// Lists the packages in the given Publication
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/marketplace"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Publication_package_package_type
// 		opt1 := _var.Publication_package_package_version
// 		_, err := marketplace.GetPublicationPackages(ctx, &marketplace.GetPublicationPackagesArgs{
// 			PublicationId:  oci_marketplace_publication.Test_publication.Id,
// 			PackageType:    &opt0,
// 			PackageVersion: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetPublicationPackages(ctx *pulumi.Context, args *GetPublicationPackagesArgs, opts ...pulumi.InvokeOption) (*GetPublicationPackagesResult, error) {
	var rv GetPublicationPackagesResult
	err := ctx.Invoke("oci:marketplace/getPublicationPackages:getPublicationPackages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicationPackages.
type GetPublicationPackagesArgs struct {
	Filters []GetPublicationPackagesFilter `pulumi:"filters"`
	// A filter to return only packages that match the given package type exactly.
	PackageType *string `pulumi:"packageType"`
	// The version of the package. Package versions are unique within a listing.
	PackageVersion *string `pulumi:"packageVersion"`
	// The unique identifier for the listing.
	PublicationId string `pulumi:"publicationId"`
}

// A collection of values returned by getPublicationPackages.
type GetPublicationPackagesResult struct {
	Filters []GetPublicationPackagesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The specified package's type.
	PackageType    *string `pulumi:"packageType"`
	PackageVersion *string `pulumi:"packageVersion"`
	PublicationId  string  `pulumi:"publicationId"`
	// The list of publication_packages.
	PublicationPackages []GetPublicationPackagesPublicationPackage `pulumi:"publicationPackages"`
}
