# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DatabasePluggableDatabaseArgs', 'DatabasePluggableDatabase']

@pulumi.input_type
class DatabasePluggableDatabaseArgs:
    def __init__(__self__, *,
                 container_database_id: pulumi.Input[str],
                 pdb_admin_password: pulumi.Input[str],
                 pdb_name: pulumi.Input[str],
                 tde_wallet_password: pulumi.Input[str],
                 defined_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 freeform_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a DatabasePluggableDatabase resource.
        :param pulumi.Input[str] container_database_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB
        :param pulumi.Input[str] pdb_admin_password: A strong password for PDB Admin. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
        :param pulumi.Input[str] pdb_name: The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
        :param pulumi.Input[str] tde_wallet_password: The existing TDE wallet password of the CDB.
        :param pulumi.Input[Mapping[str, Any]] defined_tags: (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        :param pulumi.Input[Mapping[str, Any]] freeform_tags: (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        """
        pulumi.set(__self__, "container_database_id", container_database_id)
        pulumi.set(__self__, "pdb_admin_password", pdb_admin_password)
        pulumi.set(__self__, "pdb_name", pdb_name)
        pulumi.set(__self__, "tde_wallet_password", tde_wallet_password)
        if defined_tags is not None:
            pulumi.set(__self__, "defined_tags", defined_tags)
        if freeform_tags is not None:
            pulumi.set(__self__, "freeform_tags", freeform_tags)

    @property
    @pulumi.getter(name="containerDatabaseId")
    def container_database_id(self) -> pulumi.Input[str]:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB
        """
        return pulumi.get(self, "container_database_id")

    @container_database_id.setter
    def container_database_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_database_id", value)

    @property
    @pulumi.getter(name="pdbAdminPassword")
    def pdb_admin_password(self) -> pulumi.Input[str]:
        """
        A strong password for PDB Admin. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
        """
        return pulumi.get(self, "pdb_admin_password")

    @pdb_admin_password.setter
    def pdb_admin_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "pdb_admin_password", value)

    @property
    @pulumi.getter(name="pdbName")
    def pdb_name(self) -> pulumi.Input[str]:
        """
        The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
        """
        return pulumi.get(self, "pdb_name")

    @pdb_name.setter
    def pdb_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "pdb_name", value)

    @property
    @pulumi.getter(name="tdeWalletPassword")
    def tde_wallet_password(self) -> pulumi.Input[str]:
        """
        The existing TDE wallet password of the CDB.
        """
        return pulumi.get(self, "tde_wallet_password")

    @tde_wallet_password.setter
    def tde_wallet_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "tde_wallet_password", value)

    @property
    @pulumi.getter(name="definedTags")
    def defined_tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        """
        return pulumi.get(self, "defined_tags")

    @defined_tags.setter
    def defined_tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "defined_tags", value)

    @property
    @pulumi.getter(name="freeformTags")
    def freeform_tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        """
        return pulumi.get(self, "freeform_tags")

    @freeform_tags.setter
    def freeform_tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "freeform_tags", value)


@pulumi.input_type
class _DatabasePluggableDatabaseState:
    def __init__(__self__, *,
                 compartment_id: Optional[pulumi.Input[str]] = None,
                 connection_strings: Optional[pulumi.Input['DatabasePluggableDatabaseConnectionStringsArgs']] = None,
                 container_database_id: Optional[pulumi.Input[str]] = None,
                 defined_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 freeform_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 is_restricted: Optional[pulumi.Input[bool]] = None,
                 lifecycle_details: Optional[pulumi.Input[str]] = None,
                 open_mode: Optional[pulumi.Input[str]] = None,
                 pdb_admin_password: Optional[pulumi.Input[str]] = None,
                 pdb_name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tde_wallet_password: Optional[pulumi.Input[str]] = None,
                 time_created: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DatabasePluggableDatabase resources.
        :param pulumi.Input[str] compartment_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        :param pulumi.Input['DatabasePluggableDatabaseConnectionStringsArgs'] connection_strings: Connection strings to connect to an Oracle Pluggable Database.
        :param pulumi.Input[str] container_database_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB
        :param pulumi.Input[Mapping[str, Any]] defined_tags: (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        :param pulumi.Input[Mapping[str, Any]] freeform_tags: (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        :param pulumi.Input[bool] is_restricted: The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode, the user needs both create a session and have restricted session privileges to connect to it.
        :param pulumi.Input[str] lifecycle_details: Detailed message for the lifecycle state.
        :param pulumi.Input[str] open_mode: The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software).
        :param pulumi.Input[str] pdb_admin_password: A strong password for PDB Admin. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
        :param pulumi.Input[str] pdb_name: The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
        :param pulumi.Input[str] state: The current state of the pluggable database.
        :param pulumi.Input[str] tde_wallet_password: The existing TDE wallet password of the CDB.
        :param pulumi.Input[str] time_created: The date and time the pluggable database was created.
        """
        if compartment_id is not None:
            pulumi.set(__self__, "compartment_id", compartment_id)
        if connection_strings is not None:
            pulumi.set(__self__, "connection_strings", connection_strings)
        if container_database_id is not None:
            pulumi.set(__self__, "container_database_id", container_database_id)
        if defined_tags is not None:
            pulumi.set(__self__, "defined_tags", defined_tags)
        if freeform_tags is not None:
            pulumi.set(__self__, "freeform_tags", freeform_tags)
        if is_restricted is not None:
            pulumi.set(__self__, "is_restricted", is_restricted)
        if lifecycle_details is not None:
            pulumi.set(__self__, "lifecycle_details", lifecycle_details)
        if open_mode is not None:
            pulumi.set(__self__, "open_mode", open_mode)
        if pdb_admin_password is not None:
            pulumi.set(__self__, "pdb_admin_password", pdb_admin_password)
        if pdb_name is not None:
            pulumi.set(__self__, "pdb_name", pdb_name)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tde_wallet_password is not None:
            pulumi.set(__self__, "tde_wallet_password", tde_wallet_password)
        if time_created is not None:
            pulumi.set(__self__, "time_created", time_created)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        """
        return pulumi.get(self, "compartment_id")

    @compartment_id.setter
    def compartment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compartment_id", value)

    @property
    @pulumi.getter(name="connectionStrings")
    def connection_strings(self) -> Optional[pulumi.Input['DatabasePluggableDatabaseConnectionStringsArgs']]:
        """
        Connection strings to connect to an Oracle Pluggable Database.
        """
        return pulumi.get(self, "connection_strings")

    @connection_strings.setter
    def connection_strings(self, value: Optional[pulumi.Input['DatabasePluggableDatabaseConnectionStringsArgs']]):
        pulumi.set(self, "connection_strings", value)

    @property
    @pulumi.getter(name="containerDatabaseId")
    def container_database_id(self) -> Optional[pulumi.Input[str]]:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB
        """
        return pulumi.get(self, "container_database_id")

    @container_database_id.setter
    def container_database_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_database_id", value)

    @property
    @pulumi.getter(name="definedTags")
    def defined_tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        """
        return pulumi.get(self, "defined_tags")

    @defined_tags.setter
    def defined_tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "defined_tags", value)

    @property
    @pulumi.getter(name="freeformTags")
    def freeform_tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        """
        return pulumi.get(self, "freeform_tags")

    @freeform_tags.setter
    def freeform_tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "freeform_tags", value)

    @property
    @pulumi.getter(name="isRestricted")
    def is_restricted(self) -> Optional[pulumi.Input[bool]]:
        """
        The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode, the user needs both create a session and have restricted session privileges to connect to it.
        """
        return pulumi.get(self, "is_restricted")

    @is_restricted.setter
    def is_restricted(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_restricted", value)

    @property
    @pulumi.getter(name="lifecycleDetails")
    def lifecycle_details(self) -> Optional[pulumi.Input[str]]:
        """
        Detailed message for the lifecycle state.
        """
        return pulumi.get(self, "lifecycle_details")

    @lifecycle_details.setter
    def lifecycle_details(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lifecycle_details", value)

    @property
    @pulumi.getter(name="openMode")
    def open_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software).
        """
        return pulumi.get(self, "open_mode")

    @open_mode.setter
    def open_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "open_mode", value)

    @property
    @pulumi.getter(name="pdbAdminPassword")
    def pdb_admin_password(self) -> Optional[pulumi.Input[str]]:
        """
        A strong password for PDB Admin. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
        """
        return pulumi.get(self, "pdb_admin_password")

    @pdb_admin_password.setter
    def pdb_admin_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pdb_admin_password", value)

    @property
    @pulumi.getter(name="pdbName")
    def pdb_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
        """
        return pulumi.get(self, "pdb_name")

    @pdb_name.setter
    def pdb_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pdb_name", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The current state of the pluggable database.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="tdeWalletPassword")
    def tde_wallet_password(self) -> Optional[pulumi.Input[str]]:
        """
        The existing TDE wallet password of the CDB.
        """
        return pulumi.get(self, "tde_wallet_password")

    @tde_wallet_password.setter
    def tde_wallet_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tde_wallet_password", value)

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time the pluggable database was created.
        """
        return pulumi.get(self, "time_created")

    @time_created.setter
    def time_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_created", value)


class DatabasePluggableDatabase(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_database_id: Optional[pulumi.Input[str]] = None,
                 defined_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 freeform_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 pdb_admin_password: Optional[pulumi.Input[str]] = None,
                 pdb_name: Optional[pulumi.Input[str]] = None,
                 tde_wallet_password: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource provides the Pluggable Database resource in Oracle Cloud Infrastructure Database service.

        Creates and starts a pluggable database in the specified container database.
        Use the [StartPluggableDatabase](#/en/database/latest/PluggableDatabase/StartPluggableDatabase] and [StopPluggableDatabase](#/en/database/latest/PluggableDatabase/StopPluggableDatabase] APIs to start and stop the pluggable database.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_pluggable_database = oci.DatabasePluggableDatabase("testPluggableDatabase",
            container_database_id=oci_database_database["test_database"]["id"],
            pdb_admin_password=var["pluggable_database_pdb_admin_password"],
            pdb_name=var["pluggable_database_pdb_name"],
            tde_wallet_password=var["pluggable_database_tde_wallet_password"],
            defined_tags=var["pluggable_database_defined_tags"],
            freeform_tags={
                "Department": "Finance",
            })
        ```

        ## Import

        PluggableDatabases can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:index/databasePluggableDatabase:DatabasePluggableDatabase test_pluggable_database "id"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_database_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB
        :param pulumi.Input[Mapping[str, Any]] defined_tags: (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        :param pulumi.Input[Mapping[str, Any]] freeform_tags: (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        :param pulumi.Input[str] pdb_admin_password: A strong password for PDB Admin. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
        :param pulumi.Input[str] pdb_name: The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
        :param pulumi.Input[str] tde_wallet_password: The existing TDE wallet password of the CDB.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabasePluggableDatabaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource provides the Pluggable Database resource in Oracle Cloud Infrastructure Database service.

        Creates and starts a pluggable database in the specified container database.
        Use the [StartPluggableDatabase](#/en/database/latest/PluggableDatabase/StartPluggableDatabase] and [StopPluggableDatabase](#/en/database/latest/PluggableDatabase/StopPluggableDatabase] APIs to start and stop the pluggable database.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_pluggable_database = oci.DatabasePluggableDatabase("testPluggableDatabase",
            container_database_id=oci_database_database["test_database"]["id"],
            pdb_admin_password=var["pluggable_database_pdb_admin_password"],
            pdb_name=var["pluggable_database_pdb_name"],
            tde_wallet_password=var["pluggable_database_tde_wallet_password"],
            defined_tags=var["pluggable_database_defined_tags"],
            freeform_tags={
                "Department": "Finance",
            })
        ```

        ## Import

        PluggableDatabases can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:index/databasePluggableDatabase:DatabasePluggableDatabase test_pluggable_database "id"
        ```

        :param str resource_name: The name of the resource.
        :param DatabasePluggableDatabaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabasePluggableDatabaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_database_id: Optional[pulumi.Input[str]] = None,
                 defined_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 freeform_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 pdb_admin_password: Optional[pulumi.Input[str]] = None,
                 pdb_name: Optional[pulumi.Input[str]] = None,
                 tde_wallet_password: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabasePluggableDatabaseArgs.__new__(DatabasePluggableDatabaseArgs)

            if container_database_id is None and not opts.urn:
                raise TypeError("Missing required property 'container_database_id'")
            __props__.__dict__["container_database_id"] = container_database_id
            __props__.__dict__["defined_tags"] = defined_tags
            __props__.__dict__["freeform_tags"] = freeform_tags
            if pdb_admin_password is None and not opts.urn:
                raise TypeError("Missing required property 'pdb_admin_password'")
            __props__.__dict__["pdb_admin_password"] = pdb_admin_password
            if pdb_name is None and not opts.urn:
                raise TypeError("Missing required property 'pdb_name'")
            __props__.__dict__["pdb_name"] = pdb_name
            if tde_wallet_password is None and not opts.urn:
                raise TypeError("Missing required property 'tde_wallet_password'")
            __props__.__dict__["tde_wallet_password"] = tde_wallet_password
            __props__.__dict__["compartment_id"] = None
            __props__.__dict__["connection_strings"] = None
            __props__.__dict__["is_restricted"] = None
            __props__.__dict__["lifecycle_details"] = None
            __props__.__dict__["open_mode"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["time_created"] = None
        super(DatabasePluggableDatabase, __self__).__init__(
            'oci:index/databasePluggableDatabase:DatabasePluggableDatabase',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compartment_id: Optional[pulumi.Input[str]] = None,
            connection_strings: Optional[pulumi.Input[pulumi.InputType['DatabasePluggableDatabaseConnectionStringsArgs']]] = None,
            container_database_id: Optional[pulumi.Input[str]] = None,
            defined_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            freeform_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            is_restricted: Optional[pulumi.Input[bool]] = None,
            lifecycle_details: Optional[pulumi.Input[str]] = None,
            open_mode: Optional[pulumi.Input[str]] = None,
            pdb_admin_password: Optional[pulumi.Input[str]] = None,
            pdb_name: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            tde_wallet_password: Optional[pulumi.Input[str]] = None,
            time_created: Optional[pulumi.Input[str]] = None) -> 'DatabasePluggableDatabase':
        """
        Get an existing DatabasePluggableDatabase resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compartment_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        :param pulumi.Input[pulumi.InputType['DatabasePluggableDatabaseConnectionStringsArgs']] connection_strings: Connection strings to connect to an Oracle Pluggable Database.
        :param pulumi.Input[str] container_database_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB
        :param pulumi.Input[Mapping[str, Any]] defined_tags: (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        :param pulumi.Input[Mapping[str, Any]] freeform_tags: (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        :param pulumi.Input[bool] is_restricted: The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode, the user needs both create a session and have restricted session privileges to connect to it.
        :param pulumi.Input[str] lifecycle_details: Detailed message for the lifecycle state.
        :param pulumi.Input[str] open_mode: The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software).
        :param pulumi.Input[str] pdb_admin_password: A strong password for PDB Admin. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
        :param pulumi.Input[str] pdb_name: The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
        :param pulumi.Input[str] state: The current state of the pluggable database.
        :param pulumi.Input[str] tde_wallet_password: The existing TDE wallet password of the CDB.
        :param pulumi.Input[str] time_created: The date and time the pluggable database was created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabasePluggableDatabaseState.__new__(_DatabasePluggableDatabaseState)

        __props__.__dict__["compartment_id"] = compartment_id
        __props__.__dict__["connection_strings"] = connection_strings
        __props__.__dict__["container_database_id"] = container_database_id
        __props__.__dict__["defined_tags"] = defined_tags
        __props__.__dict__["freeform_tags"] = freeform_tags
        __props__.__dict__["is_restricted"] = is_restricted
        __props__.__dict__["lifecycle_details"] = lifecycle_details
        __props__.__dict__["open_mode"] = open_mode
        __props__.__dict__["pdb_admin_password"] = pdb_admin_password
        __props__.__dict__["pdb_name"] = pdb_name
        __props__.__dict__["state"] = state
        __props__.__dict__["tde_wallet_password"] = tde_wallet_password
        __props__.__dict__["time_created"] = time_created
        return DatabasePluggableDatabase(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> pulumi.Output[str]:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="connectionStrings")
    def connection_strings(self) -> pulumi.Output['outputs.DatabasePluggableDatabaseConnectionStrings']:
        """
        Connection strings to connect to an Oracle Pluggable Database.
        """
        return pulumi.get(self, "connection_strings")

    @property
    @pulumi.getter(name="containerDatabaseId")
    def container_database_id(self) -> pulumi.Output[str]:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB
        """
        return pulumi.get(self, "container_database_id")

    @property
    @pulumi.getter(name="definedTags")
    def defined_tags(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        """
        return pulumi.get(self, "defined_tags")

    @property
    @pulumi.getter(name="freeformTags")
    def freeform_tags(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        """
        return pulumi.get(self, "freeform_tags")

    @property
    @pulumi.getter(name="isRestricted")
    def is_restricted(self) -> pulumi.Output[bool]:
        """
        The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode, the user needs both create a session and have restricted session privileges to connect to it.
        """
        return pulumi.get(self, "is_restricted")

    @property
    @pulumi.getter(name="lifecycleDetails")
    def lifecycle_details(self) -> pulumi.Output[str]:
        """
        Detailed message for the lifecycle state.
        """
        return pulumi.get(self, "lifecycle_details")

    @property
    @pulumi.getter(name="openMode")
    def open_mode(self) -> pulumi.Output[str]:
        """
        The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software).
        """
        return pulumi.get(self, "open_mode")

    @property
    @pulumi.getter(name="pdbAdminPassword")
    def pdb_admin_password(self) -> pulumi.Output[str]:
        """
        A strong password for PDB Admin. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
        """
        return pulumi.get(self, "pdb_admin_password")

    @property
    @pulumi.getter(name="pdbName")
    def pdb_name(self) -> pulumi.Output[str]:
        """
        The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
        """
        return pulumi.get(self, "pdb_name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of the pluggable database.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="tdeWalletPassword")
    def tde_wallet_password(self) -> pulumi.Output[str]:
        """
        The existing TDE wallet password of the CDB.
        """
        return pulumi.get(self, "tde_wallet_password")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> pulumi.Output[str]:
        """
        The date and time the pluggable database was created.
        """
        return pulumi.get(self, "time_created")
