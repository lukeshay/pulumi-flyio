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

    public sealed class FlyVolumeConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tempDir")]
        public Input<Inputs.FlyTempDirVolumeArgs>? TempDir { get; set; }

        public FlyVolumeConfigArgs()
        {
        }
        public static new FlyVolumeConfigArgs Empty => new FlyVolumeConfigArgs();
    }
}
