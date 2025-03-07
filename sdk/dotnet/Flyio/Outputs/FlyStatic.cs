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
    public sealed class FlyStatic
    {
        public readonly string GuestPath;
        public readonly string? IndexDocument;
        public readonly string? TigrisBucket;
        public readonly string UrlPrefix;

        [OutputConstructor]
        private FlyStatic(
            string guestPath,

            string? indexDocument,

            string? tigrisBucket,

            string urlPrefix)
        {
            GuestPath = guestPath;
            IndexDocument = indexDocument;
            TigrisBucket = tigrisBucket;
            UrlPrefix = urlPrefix;
        }
    }
}
