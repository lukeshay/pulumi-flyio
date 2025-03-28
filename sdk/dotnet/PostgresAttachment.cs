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
    /// A Fly.io Postgres attachment connects a Postgres database to an application.
    /// </summary>
    [FlyioResourceType("flyio:index:PostgresAttachment")]
    public partial class PostgresAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The application to attach the Postgres database to.
        /// </summary>
        [Output("app")]
        public Output<string> App { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the Postgres attachment.
        /// </summary>
        [Output("attachmentId")]
        public Output<string> AttachmentId { get; private set; } = null!;

        /// <summary>
        /// The PostgreSQL connection string.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// The name of the database to use.
        /// </summary>
        [Output("databaseName")]
        public Output<string?> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// The database user to connect as.
        /// </summary>
        [Output("databaseUser")]
        public Output<string?> DatabaseUser { get; private set; } = null!;

        /// <summary>
        /// Whether to manually enter the connection details.
        /// </summary>
        [Output("manualEntry")]
        public Output<bool?> ManualEntry { get; private set; } = null!;

        /// <summary>
        /// The Postgres cluster to attach.
        /// </summary>
        [Output("postgres")]
        public Output<string> Postgres { get; private set; } = null!;

        /// <summary>
        /// The environment variable name that contains the connection string.
        /// </summary>
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
        /// <summary>
        /// The application to attach the Postgres database to.
        /// </summary>
        [Input("app", required: true)]
        public Input<string> App { get; set; } = null!;

        /// <summary>
        /// The name of the database to use.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The database user to connect as.
        /// </summary>
        [Input("databaseUser")]
        public Input<string>? DatabaseUser { get; set; }

        /// <summary>
        /// Whether to manually enter the connection details.
        /// </summary>
        [Input("manualEntry")]
        public Input<bool>? ManualEntry { get; set; }

        /// <summary>
        /// The Postgres cluster to attach.
        /// </summary>
        [Input("postgres", required: true)]
        public Input<string> Postgres { get; set; } = null!;

        /// <summary>
        /// The environment variable name to store the connection string.
        /// </summary>
        [Input("variableName")]
        public Input<string>? VariableName { get; set; }

        public PostgresAttachmentArgs()
        {
        }
        public static new PostgresAttachmentArgs Empty => new PostgresAttachmentArgs();
    }
}
