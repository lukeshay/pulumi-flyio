// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Flyio.Flyio.Outputs
{

    [OutputType]
    public sealed class FlyDNSConfig
    {
        public readonly ImmutableArray<Outputs.FlyDnsForwardRule> DnsForwardRules;
        public readonly ImmutableArray<string> Nameservers;
        public readonly ImmutableArray<Outputs.FlyDnsOption> Options;
        public readonly ImmutableArray<string> Searches;
        public readonly bool? SkipRegistration;

        [OutputConstructor]
        private FlyDNSConfig(
            ImmutableArray<Outputs.FlyDnsForwardRule> dnsForwardRules,

            ImmutableArray<string> nameservers,

            ImmutableArray<Outputs.FlyDnsOption> options,

            ImmutableArray<string> searches,

            bool? skipRegistration)
        {
            DnsForwardRules = dnsForwardRules;
            Nameservers = nameservers;
            Options = options;
            Searches = searches;
            SkipRegistration = skipRegistration;
        }
    }
}
