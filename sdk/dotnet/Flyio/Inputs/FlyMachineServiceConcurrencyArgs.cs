// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Flyio.Flyio.Inputs
{

    public sealed class FlyMachineServiceConcurrencyArgs : global::Pulumi.ResourceArgs
    {
        [Input("hardLimit")]
        public Input<int>? HardLimit { get; set; }

        [Input("softLimit")]
        public Input<int>? SoftLimit { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public FlyMachineServiceConcurrencyArgs()
        {
        }
        public static new FlyMachineServiceConcurrencyArgs Empty => new FlyMachineServiceConcurrencyArgs();
    }
}