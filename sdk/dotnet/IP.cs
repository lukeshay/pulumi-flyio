// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Flyio
{
    [FlyioResourceType("flyio:index:IP")]
    public partial class IP : global::Pulumi.CustomResource
    {
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        [Output("app")]
        public Output<string> App { get; private set; } = null!;

        [Output("createdAt")]
        public Output<Pulumi.Flyio.Time.Outputs.Time> CreatedAt { get; private set; } = null!;

        [Output("flyId")]
        public Output<string> FlyId { get; private set; } = null!;

        [Output("input")]
        public Output<Outputs.IPArgs> Input { get; private set; } = null!;

        [Output("network")]
        public Output<string?> Network { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IP resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IP(string name, IPArgs args, CustomResourceOptions? options = null)
            : base("flyio:index:IP", name, args ?? new IPArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IP(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("flyio:index:IP", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lukeshay",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IP resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IP Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IP(name, id, options);
        }
    }

    public sealed class IPArgs : global::Pulumi.ResourceArgs
    {
        [Input("addrType", required: true)]
        public Input<string> AddrType { get; set; } = null!;

        [Input("app", required: true)]
        public Input<string> App { get; set; } = null!;

        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public IPArgs()
        {
        }
        public static new IPArgs Empty => new IPArgs();
    }
}
