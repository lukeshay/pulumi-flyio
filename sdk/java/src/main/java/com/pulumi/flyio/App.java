// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.flyio.Utilities;
import com.pulumi.flyio.outputs.AppArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="flyio:index:App")
public class App extends com.pulumi.resources.CustomResource {
    @Export(name="enableSubdomains", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableSubdomains;

    public Output<Optional<Boolean>> enableSubdomains() {
        return Codegen.optional(this.enableSubdomains);
    }
    @Export(name="input", refs={AppArgs.class}, tree="[0]")
    private Output<AppArgs> input;

    public Output<AppArgs> input() {
        return this.input;
    }
    @Export(name="ipv6Address", refs={String.class}, tree="[0]")
    private Output<String> ipv6Address;

    public Output<String> ipv6Address() {
        return this.ipv6Address;
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> network;

    public Output<Optional<String>> network() {
        return Codegen.optional(this.network);
    }
    @Export(name="org", refs={String.class}, tree="[0]")
    private Output<String> org;

    public Output<String> org() {
        return this.org;
    }
    @Export(name="sharedIPAddress", refs={String.class}, tree="[0]")
    private Output<String> sharedIPAddress;

    public Output<String> sharedIPAddress() {
        return this.sharedIPAddress;
    }
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> status;

    public Output<Optional<String>> status() {
        return Codegen.optional(this.status);
    }
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public App(java.lang.String name) {
        this(name, com.pulumi.flyio.AppArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public App(java.lang.String name, com.pulumi.flyio.AppArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public App(java.lang.String name, com.pulumi.flyio.AppArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("flyio:index:App", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private App(java.lang.String name, Output<java.lang.String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("flyio:index:App", name, null, makeResourceOptions(options, id), false);
    }

    private static com.pulumi.flyio.AppArgs makeArgs(com.pulumi.flyio.AppArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? com.pulumi.flyio.AppArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static App get(java.lang.String name, Output<java.lang.String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new App(name, id, options);
    }
}
