// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Flyio.Flyio.Inputs
{

    public sealed class FlyProxyProtoOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("version")]
        public Input<string>? Version { get; set; }

        public FlyProxyProtoOptionsArgs()
        {
        }
        public static new FlyProxyProtoOptionsArgs Empty => new FlyProxyProtoOptionsArgs();
    }
}
