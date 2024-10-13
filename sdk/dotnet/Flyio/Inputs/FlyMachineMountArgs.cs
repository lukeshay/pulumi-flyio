// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Flyio.Flyio.Inputs
{

    public sealed class FlyMachineMountArgs : global::Pulumi.ResourceArgs
    {
        [Input("addSizeGb")]
        public Input<int>? AddSizeGb { get; set; }

        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("extendThresholdPercent")]
        public Input<int>? ExtendThresholdPercent { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("sizeGb")]
        public Input<int>? SizeGb { get; set; }

        [Input("sizeGbLimit")]
        public Input<int>? SizeGbLimit { get; set; }

        [Input("volume")]
        public Input<string>? Volume { get; set; }

        public FlyMachineMountArgs()
        {
        }
        public static new FlyMachineMountArgs Empty => new FlyMachineMountArgs();
    }
}
