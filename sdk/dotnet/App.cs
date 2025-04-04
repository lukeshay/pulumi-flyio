// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PulumiFlyio.Flyio
{
    /// <summary>
    /// A Fly.io application.
    /// </summary>
    [FlyioResourceType("flyio:index:App")]
    public partial class App : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether subdomains are enabled for the application.
        /// </summary>
        [Output("enableSubdomains")]
        public Output<bool?> EnableSubdomains { get; private set; } = null!;

        /// <summary>
        /// The input arguments used to create the application.
        /// </summary>
        [Output("input")]
        public Output<Outputs.AppArgs> Input { get; private set; } = null!;

        /// <summary>
        /// The name of the Fly.io application.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network the application belongs to.
        /// </summary>
        [Output("network")]
        public Output<string?> Network { get; private set; } = null!;

        /// <summary>
        /// The organization the application belongs to.
        /// </summary>
        [Output("org")]
        public Output<string> Org { get; private set; } = null!;

        /// <summary>
        /// The current status of the application.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;


        /// <summary>
        /// Create a App resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public App(string name, AppArgs args, CustomResourceOptions? options = null)
            : base("flyio:index:App", name, args ?? new AppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private App(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("flyio:index:App", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing App resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static App Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new App(name, id, options);
        }
    }

    public sealed class AppArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable subdomains for the application.
        /// </summary>
        [Input("enableSubdomains")]
        public Input<bool>? EnableSubdomains { get; set; }

        /// <summary>
        /// The name of the Fly.io application.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The network the application belongs to.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The organization the application belongs to.
        /// </summary>
        [Input("org", required: true)]
        public Input<string> Org { get; set; } = null!;

        public AppArgs()
        {
        }
        public static new AppArgs Empty => new AppArgs();
    }
}
