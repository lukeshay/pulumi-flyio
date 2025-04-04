# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import flyio as _flyio
from . import outputs

__all__ = ['VolumeArgs', 'Volume']

@pulumi.input_type
class VolumeArgs:
    def __init__(__self__, *,
                 app: pulumi.Input[str],
                 auto_backup_enabled: Optional[pulumi.Input[bool]] = None,
                 compute: Optional[pulumi.Input['_flyio.FlyMachineGuestArgs']] = None,
                 compute_image: Optional[pulumi.Input[str]] = None,
                 encrypted: Optional[pulumi.Input[bool]] = None,
                 fstype: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 require_unique_zone: Optional[pulumi.Input[bool]] = None,
                 size_gb: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 snapshot_retention: Optional[pulumi.Input[int]] = None,
                 source_volume_id: Optional[pulumi.Input[str]] = None,
                 unique_zone_app_wide: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Volume resource.
        :param pulumi.Input[str] app: The Fly.io application to attach the volume to.
        :param pulumi.Input[bool] auto_backup_enabled: Whether to enable automatic backups for the volume.
        """
        pulumi.set(__self__, "app", app)
        if auto_backup_enabled is not None:
            pulumi.set(__self__, "auto_backup_enabled", auto_backup_enabled)
        if compute is not None:
            pulumi.set(__self__, "compute", compute)
        if compute_image is not None:
            pulumi.set(__self__, "compute_image", compute_image)
        if encrypted is not None:
            pulumi.set(__self__, "encrypted", encrypted)
        if fstype is not None:
            pulumi.set(__self__, "fstype", fstype)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if require_unique_zone is not None:
            pulumi.set(__self__, "require_unique_zone", require_unique_zone)
        if size_gb is not None:
            pulumi.set(__self__, "size_gb", size_gb)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if snapshot_retention is not None:
            pulumi.set(__self__, "snapshot_retention", snapshot_retention)
        if source_volume_id is not None:
            pulumi.set(__self__, "source_volume_id", source_volume_id)
        if unique_zone_app_wide is not None:
            pulumi.set(__self__, "unique_zone_app_wide", unique_zone_app_wide)

    @property
    @pulumi.getter
    def app(self) -> pulumi.Input[str]:
        """
        The Fly.io application to attach the volume to.
        """
        return pulumi.get(self, "app")

    @app.setter
    def app(self, value: pulumi.Input[str]):
        pulumi.set(self, "app", value)

    @property
    @pulumi.getter(name="autoBackupEnabled")
    def auto_backup_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable automatic backups for the volume.
        """
        return pulumi.get(self, "auto_backup_enabled")

    @auto_backup_enabled.setter
    def auto_backup_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_backup_enabled", value)

    @property
    @pulumi.getter
    def compute(self) -> Optional[pulumi.Input['_flyio.FlyMachineGuestArgs']]:
        return pulumi.get(self, "compute")

    @compute.setter
    def compute(self, value: Optional[pulumi.Input['_flyio.FlyMachineGuestArgs']]):
        pulumi.set(self, "compute", value)

    @property
    @pulumi.getter(name="computeImage")
    def compute_image(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "compute_image")

    @compute_image.setter
    def compute_image(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute_image", value)

    @property
    @pulumi.getter
    def encrypted(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "encrypted")

    @encrypted.setter
    def encrypted(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "encrypted", value)

    @property
    @pulumi.getter
    def fstype(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fstype")

    @fstype.setter
    def fstype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fstype", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="requireUniqueZone")
    def require_unique_zone(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "require_unique_zone")

    @require_unique_zone.setter
    def require_unique_zone(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_unique_zone", value)

    @property
    @pulumi.getter(name="sizeGb")
    def size_gb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "size_gb")

    @size_gb.setter
    def size_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_gb", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter(name="snapshotRetention")
    def snapshot_retention(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "snapshot_retention")

    @snapshot_retention.setter
    def snapshot_retention(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "snapshot_retention", value)

    @property
    @pulumi.getter(name="sourceVolumeId")
    def source_volume_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_volume_id")

    @source_volume_id.setter
    def source_volume_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_volume_id", value)

    @property
    @pulumi.getter(name="uniqueZoneAppWide")
    def unique_zone_app_wide(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "unique_zone_app_wide")

    @unique_zone_app_wide.setter
    def unique_zone_app_wide(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unique_zone_app_wide", value)


class Volume(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app: Optional[pulumi.Input[str]] = None,
                 auto_backup_enabled: Optional[pulumi.Input[bool]] = None,
                 compute: Optional[pulumi.Input[Union['_flyio.FlyMachineGuestArgs', '_flyio.FlyMachineGuestArgsDict']]] = None,
                 compute_image: Optional[pulumi.Input[str]] = None,
                 encrypted: Optional[pulumi.Input[bool]] = None,
                 fstype: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 require_unique_zone: Optional[pulumi.Input[bool]] = None,
                 size_gb: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 snapshot_retention: Optional[pulumi.Input[int]] = None,
                 source_volume_id: Optional[pulumi.Input[str]] = None,
                 unique_zone_app_wide: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        A Fly.io volume provides persistent storage for your applications.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app: The Fly.io application to attach the volume to.
        :param pulumi.Input[bool] auto_backup_enabled: Whether to enable automatic backups for the volume.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VolumeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Fly.io volume provides persistent storage for your applications.

        :param str resource_name: The name of the resource.
        :param VolumeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VolumeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app: Optional[pulumi.Input[str]] = None,
                 auto_backup_enabled: Optional[pulumi.Input[bool]] = None,
                 compute: Optional[pulumi.Input[Union['_flyio.FlyMachineGuestArgs', '_flyio.FlyMachineGuestArgsDict']]] = None,
                 compute_image: Optional[pulumi.Input[str]] = None,
                 encrypted: Optional[pulumi.Input[bool]] = None,
                 fstype: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 require_unique_zone: Optional[pulumi.Input[bool]] = None,
                 size_gb: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 snapshot_retention: Optional[pulumi.Input[int]] = None,
                 source_volume_id: Optional[pulumi.Input[str]] = None,
                 unique_zone_app_wide: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VolumeArgs.__new__(VolumeArgs)

            if app is None and not opts.urn:
                raise TypeError("Missing required property 'app'")
            __props__.__dict__["app"] = app
            __props__.__dict__["auto_backup_enabled"] = auto_backup_enabled
            __props__.__dict__["compute"] = compute
            __props__.__dict__["compute_image"] = compute_image
            __props__.__dict__["encrypted"] = encrypted
            __props__.__dict__["fstype"] = fstype
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["require_unique_zone"] = require_unique_zone
            __props__.__dict__["size_gb"] = size_gb
            __props__.__dict__["snapshot_id"] = snapshot_id
            __props__.__dict__["snapshot_retention"] = snapshot_retention
            __props__.__dict__["source_volume_id"] = source_volume_id
            __props__.__dict__["unique_zone_app_wide"] = unique_zone_app_wide
            __props__.__dict__["attached_alloc_id"] = None
            __props__.__dict__["attached_machine_id"] = None
            __props__.__dict__["block_size"] = None
            __props__.__dict__["blocks"] = None
            __props__.__dict__["blocks_avail"] = None
            __props__.__dict__["blocks_free"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["fly_id"] = None
            __props__.__dict__["host_status"] = None
            __props__.__dict__["input"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["zone"] = None
        super(Volume, __self__).__init__(
            'flyio:index:Volume',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Volume':
        """
        Get an existing Volume resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VolumeArgs.__new__(VolumeArgs)

        __props__.__dict__["app"] = None
        __props__.__dict__["attached_alloc_id"] = None
        __props__.__dict__["attached_machine_id"] = None
        __props__.__dict__["auto_backup_enabled"] = None
        __props__.__dict__["block_size"] = None
        __props__.__dict__["blocks"] = None
        __props__.__dict__["blocks_avail"] = None
        __props__.__dict__["blocks_free"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["encrypted"] = None
        __props__.__dict__["fly_id"] = None
        __props__.__dict__["fstype"] = None
        __props__.__dict__["host_status"] = None
        __props__.__dict__["input"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["size_gb"] = None
        __props__.__dict__["snapshot_retention"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["zone"] = None
        return Volume(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def app(self) -> pulumi.Output[str]:
        """
        The Fly.io application the volume is attached to.
        """
        return pulumi.get(self, "app")

    @property
    @pulumi.getter(name="attachedAllocId")
    def attached_alloc_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "attached_alloc_id")

    @property
    @pulumi.getter(name="attachedMachineId")
    def attached_machine_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "attached_machine_id")

    @property
    @pulumi.getter(name="autoBackupEnabled")
    def auto_backup_enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "auto_backup_enabled")

    @property
    @pulumi.getter(name="blockSize")
    def block_size(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "block_size")

    @property
    @pulumi.getter
    def blocks(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "blocks")

    @property
    @pulumi.getter(name="blocksAvail")
    def blocks_avail(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "blocks_avail")

    @property
    @pulumi.getter(name="blocksFree")
    def blocks_free(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "blocks_free")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def encrypted(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter(name="flyId")
    def fly_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "fly_id")

    @property
    @pulumi.getter
    def fstype(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "fstype")

    @property
    @pulumi.getter(name="hostStatus")
    def host_status(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "host_status")

    @property
    @pulumi.getter
    def input(self) -> pulumi.Output['outputs.VolumeArgs']:
        """
        The input arguments used to create the volume.
        """
        return pulumi.get(self, "input")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sizeGb")
    def size_gb(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "size_gb")

    @property
    @pulumi.getter(name="snapshotRetention")
    def snapshot_retention(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "snapshot_retention")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "zone")

