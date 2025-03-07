// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PulumiFlyio.Flyio.Flyio.Outputs
{

    [OutputType]
    public sealed class FlyHTTPResponseOptions
    {
        public readonly ImmutableDictionary<string, ImmutableDictionary<string, object>>? Headers;
        public readonly bool? Pristine;

        [OutputConstructor]
        private FlyHTTPResponseOptions(
            ImmutableDictionary<string, ImmutableDictionary<string, object>>? headers,

            bool? pristine)
        {
            Headers = headers;
            Pristine = pristine;
        }
    }
}
