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

    public sealed class FlyEnvFromArgs : global::Pulumi.ResourceArgs
    {
        [Input("envVar")]
        public Input<string>? EnvVar { get; set; }

        [Input("fieldRef")]
        public Input<string>? FieldRef { get; set; }

        public FlyEnvFromArgs()
        {
        }
        public static new FlyEnvFromArgs Empty => new FlyEnvFromArgs();
    }
}
