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

    public sealed class FlyDnsForwardRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("addr")]
        public Input<string>? Addr { get; set; }

        [Input("basename")]
        public Input<string>? Basename { get; set; }

        public FlyDnsForwardRuleArgs()
        {
        }
        public static new FlyDnsForwardRuleArgs Empty => new FlyDnsForwardRuleArgs();
    }
}
