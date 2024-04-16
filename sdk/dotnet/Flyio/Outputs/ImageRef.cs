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
    public sealed class ImageRef
    {
        public readonly string? Digest;
        public readonly ImmutableDictionary<string, string>? Labels;
        public readonly string? Registry;
        public readonly string? Repository;
        public readonly string? Tag;

        [OutputConstructor]
        private ImageRef(
            string? digest,

            ImmutableDictionary<string, string>? labels,

            string? registry,

            string? repository,

            string? tag)
        {
            Digest = digest;
            Labels = labels;
            Registry = registry;
            Repository = repository;
            Tag = tag;
        }
    }
}
