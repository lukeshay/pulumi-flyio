# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import flyio as _flyio

__all__ = [
    'AppArgs',
    'MachineArgs',
]

@pulumi.output_type
class AppArgs(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "appName":
            suggest = "app_name"
        elif key == "orgSlug":
            suggest = "org_slug"
        elif key == "enableSubdomains":
            suggest = "enable_subdomains"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AppArgs. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AppArgs.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AppArgs.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 app_name: str,
                 org_slug: str,
                 enable_subdomains: Optional[bool] = None,
                 network: Optional[str] = None):
        pulumi.set(__self__, "app_name", app_name)
        pulumi.set(__self__, "org_slug", org_slug)
        if enable_subdomains is not None:
            pulumi.set(__self__, "enable_subdomains", enable_subdomains)
        if network is not None:
            pulumi.set(__self__, "network", network)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> str:
        return pulumi.get(self, "app_name")

    @property
    @pulumi.getter(name="orgSlug")
    def org_slug(self) -> str:
        return pulumi.get(self, "org_slug")

    @property
    @pulumi.getter(name="enableSubdomains")
    def enable_subdomains(self) -> Optional[bool]:
        return pulumi.get(self, "enable_subdomains")

    @property
    @pulumi.getter
    def network(self) -> Optional[str]:
        return pulumi.get(self, "network")


@pulumi.output_type
class MachineArgs(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "appName":
            suggest = "app_name"
        elif key == "leaseTtl":
            suggest = "lease_ttl"
        elif key == "skipLaunch":
            suggest = "skip_launch"
        elif key == "skipServiceRegistration":
            suggest = "skip_service_registration"
        elif key == "updateStrategy":
            suggest = "update_strategy"
        elif key == "waitForChecks":
            suggest = "wait_for_checks"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in MachineArgs. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        MachineArgs.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        MachineArgs.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 app_name: str,
                 config: '_flyio.outputs.FlyMachineConfig',
                 lease_ttl: Optional[int] = None,
                 lsvd: Optional[bool] = None,
                 name: Optional[str] = None,
                 region: Optional[str] = None,
                 skip_launch: Optional[bool] = None,
                 skip_service_registration: Optional[bool] = None,
                 update_strategy: Optional[str] = None,
                 wait_for_checks: Optional[int] = None):
        pulumi.set(__self__, "app_name", app_name)
        pulumi.set(__self__, "config", config)
        if lease_ttl is not None:
            pulumi.set(__self__, "lease_ttl", lease_ttl)
        if lsvd is not None:
            pulumi.set(__self__, "lsvd", lsvd)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if skip_launch is not None:
            pulumi.set(__self__, "skip_launch", skip_launch)
        if skip_service_registration is not None:
            pulumi.set(__self__, "skip_service_registration", skip_service_registration)
        if update_strategy is not None:
            pulumi.set(__self__, "update_strategy", update_strategy)
        if wait_for_checks is not None:
            pulumi.set(__self__, "wait_for_checks", wait_for_checks)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> str:
        return pulumi.get(self, "app_name")

    @property
    @pulumi.getter
    def config(self) -> '_flyio.outputs.FlyMachineConfig':
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="leaseTtl")
    def lease_ttl(self) -> Optional[int]:
        return pulumi.get(self, "lease_ttl")

    @property
    @pulumi.getter
    def lsvd(self) -> Optional[bool]:
        return pulumi.get(self, "lsvd")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="skipLaunch")
    def skip_launch(self) -> Optional[bool]:
        return pulumi.get(self, "skip_launch")

    @property
    @pulumi.getter(name="skipServiceRegistration")
    def skip_service_registration(self) -> Optional[bool]:
        return pulumi.get(self, "skip_service_registration")

    @property
    @pulumi.getter(name="updateStrategy")
    def update_strategy(self) -> Optional[str]:
        return pulumi.get(self, "update_strategy")

    @property
    @pulumi.getter(name="waitForChecks")
    def wait_for_checks(self) -> Optional[int]:
        return pulumi.get(self, "wait_for_checks")


