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

    public sealed class FlyMachinePortArgs : global::Pulumi.ResourceArgs
    {
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        [Input("forceHttps")]
        public Input<bool>? ForceHttps { get; set; }

        [Input("handlers")]
        private InputList<string>? _handlers;
        public InputList<string> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<string>());
            set => _handlers = value;
        }

        [Input("httpOptions")]
        public Input<Inputs.FlyHTTPOptionsArgs>? HttpOptions { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("proxyProtoOptions")]
        public Input<Inputs.FlyProxyProtoOptionsArgs>? ProxyProtoOptions { get; set; }

        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        [Input("tlsOptions")]
        public Input<Inputs.FlyTLSOptionsArgs>? TlsOptions { get; set; }

        public FlyMachinePortArgs()
        {
        }
        public static new FlyMachinePortArgs Empty => new FlyMachinePortArgs();
    }
}
