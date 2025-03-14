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

__all__ = ['SecretsArgs', 'Secrets']

@pulumi.input_type
class SecretsArgs:
    def __init__(__self__, *,
                 app: pulumi.Input[str],
                 values: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        """
        The set of arguments for constructing a Secrets resource.
        :param pulumi.Input[str] app: The Fly.io application to set secrets for.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] values: The secret values to set, as key-value pairs.
        """
        pulumi.set(__self__, "app", app)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def app(self) -> pulumi.Input[str]:
        """
        The Fly.io application to set secrets for.
        """
        return pulumi.get(self, "app")

    @app.setter
    def app(self, value: pulumi.Input[str]):
        pulumi.set(self, "app", value)

    @property
    @pulumi.getter
    def values(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        The secret values to set, as key-value pairs.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "values", value)


class Secrets(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages application secrets for a Fly.io application.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app: The Fly.io application to set secrets for.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] values: The secret values to set, as key-value pairs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages application secrets for a Fly.io application.

        :param str resource_name: The name of the resource.
        :param SecretsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretsArgs.__new__(SecretsArgs)

            if app is None and not opts.urn:
                raise TypeError("Missing required property 'app'")
            __props__.__dict__["app"] = app
            if values is None and not opts.urn:
                raise TypeError("Missing required property 'values'")
            __props__.__dict__["values"] = None if values is None else pulumi.Output.secret(values)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["values"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Secrets, __self__).__init__(
            'flyio:index:Secrets',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Secrets':
        """
        Get an existing Secrets resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SecretsArgs.__new__(SecretsArgs)

        __props__.__dict__["app"] = None
        __props__.__dict__["values"] = None
        return Secrets(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def app(self) -> pulumi.Output[str]:
        """
        The Fly.io application the secrets are set for.
        """
        return pulumi.get(self, "app")

    @property
    @pulumi.getter
    def values(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The secret values, as key-value pairs.
        """
        return pulumi.get(self, "values")

