// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PulumiFlyio.Flyio.Outputs
{

    [OutputType]
    public sealed class IPArgs
    {
        /// <summary>
        /// The type of IP address (v4 or v6).
        /// </summary>
        public readonly string AddrType;
        /// <summary>
        /// The name of the Fly.io application to allocate the IP address for.
        /// </summary>
        public readonly string App;
        /// <summary>
        /// The network to allocate the IP address in.
        /// </summary>
        public readonly string? Network;
        /// <summary>
        /// The region to allocate the IP address in.
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private IPArgs(
            string addrType,

            string app,

            string? network,

            string region)
        {
            AddrType = addrType;
            App = app;
            Network = network;
            Region = region;
        }
    }
}
