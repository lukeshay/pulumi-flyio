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
    public sealed class CheckStatus
    {
        public readonly string? Name;
        public readonly string? Output;
        public readonly string? Status;
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private CheckStatus(
            string? name,

            string? output,

            string? status,

            string? updatedAt)
        {
            Name = name;
            Output = output;
            Status = status;
            UpdatedAt = updatedAt;
        }
    }
}
