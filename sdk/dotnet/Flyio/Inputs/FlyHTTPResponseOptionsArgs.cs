// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Flyio.Flyio.Inputs
{

    public sealed class FlyHTTPResponseOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("headers")]
        private InputMap<ImmutableDictionary<string, object>>? _headers;
        public InputMap<ImmutableDictionary<string, object>> Headers
        {
            get => _headers ?? (_headers = new InputMap<ImmutableDictionary<string, object>>());
            set => _headers = value;
        }

        [Input("pristine")]
        public Input<bool>? Pristine { get; set; }

        public FlyHTTPResponseOptionsArgs()
        {
        }
        public static new FlyHTTPResponseOptionsArgs Empty => new FlyHTTPResponseOptionsArgs();
    }
}