# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetPublicationResult',
    'AwaitableGetPublicationResult',
    'get_publication',
]

@pulumi.output_type
class GetPublicationResult:
    """
    A collection of values returned by getPublication.
    """
    def __init__(__self__, compartment_id=None, defined_tags=None, freeform_tags=None, icon=None, id=None, is_agreement_acknowledged=None, listing_type=None, long_description=None, name=None, package_details=None, package_type=None, publication_id=None, short_description=None, state=None, support_contacts=None, supported_operating_systems=None, time_created=None):
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if defined_tags and not isinstance(defined_tags, dict):
            raise TypeError("Expected argument 'defined_tags' to be a dict")
        pulumi.set(__self__, "defined_tags", defined_tags)
        if freeform_tags and not isinstance(freeform_tags, dict):
            raise TypeError("Expected argument 'freeform_tags' to be a dict")
        pulumi.set(__self__, "freeform_tags", freeform_tags)
        if icon and not isinstance(icon, dict):
            raise TypeError("Expected argument 'icon' to be a dict")
        pulumi.set(__self__, "icon", icon)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_agreement_acknowledged and not isinstance(is_agreement_acknowledged, bool):
            raise TypeError("Expected argument 'is_agreement_acknowledged' to be a bool")
        pulumi.set(__self__, "is_agreement_acknowledged", is_agreement_acknowledged)
        if listing_type and not isinstance(listing_type, str):
            raise TypeError("Expected argument 'listing_type' to be a str")
        pulumi.set(__self__, "listing_type", listing_type)
        if long_description and not isinstance(long_description, str):
            raise TypeError("Expected argument 'long_description' to be a str")
        pulumi.set(__self__, "long_description", long_description)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if package_details and not isinstance(package_details, dict):
            raise TypeError("Expected argument 'package_details' to be a dict")
        pulumi.set(__self__, "package_details", package_details)
        if package_type and not isinstance(package_type, str):
            raise TypeError("Expected argument 'package_type' to be a str")
        pulumi.set(__self__, "package_type", package_type)
        if publication_id and not isinstance(publication_id, str):
            raise TypeError("Expected argument 'publication_id' to be a str")
        pulumi.set(__self__, "publication_id", publication_id)
        if short_description and not isinstance(short_description, str):
            raise TypeError("Expected argument 'short_description' to be a str")
        pulumi.set(__self__, "short_description", short_description)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if support_contacts and not isinstance(support_contacts, list):
            raise TypeError("Expected argument 'support_contacts' to be a list")
        pulumi.set(__self__, "support_contacts", support_contacts)
        if supported_operating_systems and not isinstance(supported_operating_systems, list):
            raise TypeError("Expected argument 'supported_operating_systems' to be a list")
        pulumi.set(__self__, "supported_operating_systems", supported_operating_systems)
        if time_created and not isinstance(time_created, str):
            raise TypeError("Expected argument 'time_created' to be a str")
        pulumi.set(__self__, "time_created", time_created)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> str:
        """
        The Compartment id where the listings exists
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="definedTags")
    def defined_tags(self) -> Mapping[str, Any]:
        """
        The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
        """
        return pulumi.get(self, "defined_tags")

    @property
    @pulumi.getter(name="freeformTags")
    def freeform_tags(self) -> Mapping[str, Any]:
        """
        The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
        """
        return pulumi.get(self, "freeform_tags")

    @property
    @pulumi.getter
    def icon(self) -> 'outputs.GetPublicationIconResult':
        """
        The model for upload data for images and icons.
        """
        return pulumi.get(self, "icon")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The unique identifier for the listing in Marketplace.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isAgreementAcknowledged")
    def is_agreement_acknowledged(self) -> bool:
        return pulumi.get(self, "is_agreement_acknowledged")

    @property
    @pulumi.getter(name="listingType")
    def listing_type(self) -> str:
        """
        In which catalog the listing should exist.
        """
        return pulumi.get(self, "listing_type")

    @property
    @pulumi.getter(name="longDescription")
    def long_description(self) -> str:
        """
        A long description of the listing.
        """
        return pulumi.get(self, "long_description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        name of the operating system
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="packageDetails")
    def package_details(self) -> 'outputs.GetPublicationPackageDetailsResult':
        return pulumi.get(self, "package_details")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> str:
        """
        The listing's package type.
        """
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="publicationId")
    def publication_id(self) -> str:
        return pulumi.get(self, "publication_id")

    @property
    @pulumi.getter(name="shortDescription")
    def short_description(self) -> str:
        """
        A short description of the listing.
        """
        return pulumi.get(self, "short_description")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the listing in its lifecycle
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="supportContacts")
    def support_contacts(self) -> Sequence['outputs.GetPublicationSupportContactResult']:
        """
        Contact information to use to get support from the publisher for the listing.
        """
        return pulumi.get(self, "support_contacts")

    @property
    @pulumi.getter(name="supportedOperatingSystems")
    def supported_operating_systems(self) -> Sequence['outputs.GetPublicationSupportedOperatingSystemResult']:
        """
        List of operating systems supprted.
        """
        return pulumi.get(self, "supported_operating_systems")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        The date and time this publication was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z`
        """
        return pulumi.get(self, "time_created")


class AwaitableGetPublicationResult(GetPublicationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPublicationResult(
            compartment_id=self.compartment_id,
            defined_tags=self.defined_tags,
            freeform_tags=self.freeform_tags,
            icon=self.icon,
            id=self.id,
            is_agreement_acknowledged=self.is_agreement_acknowledged,
            listing_type=self.listing_type,
            long_description=self.long_description,
            name=self.name,
            package_details=self.package_details,
            package_type=self.package_type,
            publication_id=self.publication_id,
            short_description=self.short_description,
            state=self.state,
            support_contacts=self.support_contacts,
            supported_operating_systems=self.supported_operating_systems,
            time_created=self.time_created)


def get_publication(publication_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPublicationResult:
    """
    This data source provides details about a specific Publication resource in Oracle Cloud Infrastructure Marketplace service.

    Get details of a publication

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_publication = oci.marketplace.get_publication(publication_id=oci_marketplace_publication["test_publication"]["id"])
    ```


    :param str publication_id: The unique identifier for the listing.
    """
    __args__ = dict()
    __args__['publicationId'] = publication_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:marketplace/getPublication:getPublication', __args__, opts=opts, typ=GetPublicationResult).value

    return AwaitableGetPublicationResult(
        compartment_id=__ret__.compartment_id,
        defined_tags=__ret__.defined_tags,
        freeform_tags=__ret__.freeform_tags,
        icon=__ret__.icon,
        id=__ret__.id,
        is_agreement_acknowledged=__ret__.is_agreement_acknowledged,
        listing_type=__ret__.listing_type,
        long_description=__ret__.long_description,
        name=__ret__.name,
        package_details=__ret__.package_details,
        package_type=__ret__.package_type,
        publication_id=__ret__.publication_id,
        short_description=__ret__.short_description,
        state=__ret__.state,
        support_contacts=__ret__.support_contacts,
        supported_operating_systems=__ret__.supported_operating_systems,
        time_created=__ret__.time_created)