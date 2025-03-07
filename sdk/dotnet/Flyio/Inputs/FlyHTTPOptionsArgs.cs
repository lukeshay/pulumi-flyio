// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PulumiFlyio.Flyio.Flyio.Inputs
{

    public sealed class FlyHTTPOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("compress")]
        public Input<bool>? Compress { get; set; }

        [Input("h2Backend")]
        public Input<bool>? H2Backend { get; set; }

        [Input("headersReadTimeout")]
        public Input<int>? HeadersReadTimeout { get; set; }

        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        [Input("response")]
        public Input<Inputs.FlyHTTPResponseOptionsArgs>? Response { get; set; }

        public FlyHTTPOptionsArgs()
        {
        }
        public static new FlyHTTPOptionsArgs Empty => new FlyHTTPOptionsArgs();
    }
}
