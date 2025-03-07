// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class App extends pulumi.CustomResource {
    /**
     * Get an existing App resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): App {
        return new App(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'flyio:index:App';

    /**
     * Returns true if the given object is an instance of App.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is App {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === App.__pulumiType;
    }

    public readonly enableSubdomains!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly input!: pulumi.Output<outputs.AppArgs>;
    public /*out*/ readonly ipv6Address!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly network!: pulumi.Output<string | undefined>;
    public readonly org!: pulumi.Output<string>;
    public /*out*/ readonly sharedIPAddress!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string | undefined>;
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a App resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.org === undefined) && !opts.urn) {
                throw new Error("Missing required property 'org'");
            }
            resourceInputs["enableSubdomains"] = args ? args.enableSubdomains : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["input"] = undefined /*out*/;
            resourceInputs["ipv6Address"] = undefined /*out*/;
            resourceInputs["sharedIPAddress"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["enableSubdomains"] = undefined /*out*/;
            resourceInputs["input"] = undefined /*out*/;
            resourceInputs["ipv6Address"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["org"] = undefined /*out*/;
            resourceInputs["sharedIPAddress"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(App.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a App resource.
 */
export interface AppArgs {
    enableSubdomains?: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
    network?: pulumi.Input<string>;
    org: pulumi.Input<string>;
}
