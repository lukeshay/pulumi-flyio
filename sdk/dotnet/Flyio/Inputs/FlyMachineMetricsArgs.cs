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

    public sealed class FlyMachineMetricsArgs : global::Pulumi.ResourceArgs
    {
        [Input("https")]
        public Input<bool>? Https { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        public FlyMachineMetricsArgs()
        {
        }
        public static new FlyMachineMetricsArgs Empty => new FlyMachineMetricsArgs();
    }
}
