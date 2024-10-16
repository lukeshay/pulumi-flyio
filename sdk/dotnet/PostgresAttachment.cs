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
    [FlyioResourceType("flyio:index:PostgresAttachment")]
    public partial class PostgresAttachment : global::Pulumi.CustomResource
    {
        [Output("app")]
        public Output<string> App { get; private set; } = null!;

        [Output("attachmentId")]
        public Output<string> AttachmentId { get; private set; } = null!;

        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        [Output("databaseName")]
        public Output<string?> DatabaseName { get; private set; } = null!;

        [Output("databaseUser")]
        public Output<string?> DatabaseUser { get; private set; } = null!;

        [Output("manualEntry")]
        public Output<bool?> ManualEntry { get; private set; } = null!;

        [Output("postgres")]
        public Output<string> Postgres { get; private set; } = null!;

        [Output("variableName")]
        public Output<string> VariableName { get; private set; } = null!;


        /// <summary>
        /// Create a PostgresAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PostgresAttachment(string name, PostgresAttachmentArgs args, CustomResourceOptions? options = null)
            : base("flyio:index:PostgresAttachment", name, args ?? new PostgresAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PostgresAttachment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("flyio:index:PostgresAttachment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lukeshay",
                AdditionalSecretOutputs =
                {
                    "connectionString",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PostgresAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PostgresAttachment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PostgresAttachment(name, id, options);
        }
    }

    public sealed class PostgresAttachmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("app", required: true)]
        public Input<string> App { get; set; } = null!;

        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("databaseUser")]
        public Input<string>? DatabaseUser { get; set; }

        [Input("manualEntry")]
        public Input<bool>? ManualEntry { get; set; }

        [Input("postgres", required: true)]
        public Input<string> Postgres { get; set; } = null!;

        [Input("variableName")]
        public Input<string>? VariableName { get; set; }

        public PostgresAttachmentArgs()
        {
        }
        public static new PostgresAttachmentArgs Empty => new PostgresAttachmentArgs();
    }
}
