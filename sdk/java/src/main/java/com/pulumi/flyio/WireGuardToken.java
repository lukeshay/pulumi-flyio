// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.flyio.Utilities;
import com.pulumi.flyio.WireGuardTokenArgs;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="flyio:index:WireGuardToken")
public class WireGuardToken extends com.pulumi.resources.CustomResource {
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> name;

    public Output<Optional<String>> name() {
        return Codegen.optional(this.name);
    }
    @Export(name="org", refs={String.class}, tree="[0]")
    private Output<String> org;

    public Output<String> org() {
        return this.org;
    }
    @Export(name="token", refs={String.class}, tree="[0]")
    private Output<String> token;

    public Output<String> token() {
        return this.token;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public WireGuardToken(java.lang.String name) {
        this(name, WireGuardTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public WireGuardToken(java.lang.String name, WireGuardTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public WireGuardToken(java.lang.String name, WireGuardTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("flyio:index:WireGuardToken", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private WireGuardToken(java.lang.String name, Output<java.lang.String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("flyio:index:WireGuardToken", name, null, makeResourceOptions(options, id), false);
    }

    private static WireGuardTokenArgs makeArgs(WireGuardTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? WireGuardTokenArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "token"
            ))
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
    public static WireGuardToken get(java.lang.String name, Output<java.lang.String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new WireGuardToken(name, id, options);
    }
}
