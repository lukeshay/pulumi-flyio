// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Flyio.Flyio.Inputs
{

    public sealed class FlyFileArgs : global::Pulumi.ResourceArgs
    {
        [Input("guestPath")]
        public Input<string>? GuestPath { get; set; }

        [Input("rawValue")]
        public Input<string>? RawValue { get; set; }

        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        public FlyFileArgs()
        {
        }
        public static new FlyFileArgs Empty => new FlyFileArgs();
    }
}
