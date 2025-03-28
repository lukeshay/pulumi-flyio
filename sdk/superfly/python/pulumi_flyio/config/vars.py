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
from .. import _utilities

import types

__config__ = pulumi.Config('flyio')


class _ExportableConfig(types.ModuleType):
    @property
    def token(self) -> str:
        """
        API key for the Fly.io API.
        """
        return __config__.get('token') or (_utilities.get_env('FLY_API_TOKEN', 'FLY_TOKEN', 'FLY_API_KEY', 'FLY_KEY') or '')

