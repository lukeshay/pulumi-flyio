// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.flyio.SecretsArgs;
import com.pulumi.flyio.Utilities;
import java.lang.String;
import java.util.List;
import java.util.Map;
import javax.annotation.Nullable;

/**
 * Manages application secrets for a Fly.io application.
 * 
 */
@ResourceType(type="flyio:index:Secrets")
public class Secrets extends com.pulumi.resources.CustomResource {
    /**
     * The Fly.io application the secrets are set for.
     * 
     */
    @Export(name="app", refs={String.class}, tree="[0]")
    private Output<String> app;

    /**
     * @return The Fly.io application the secrets are set for.
     * 
     */
    public Output<String> app() {
        return this.app;
    }
    /**
     * The secret values, as key-value pairs.
     * 
     */
    @Export(name="values", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> values;

    /**
     * @return The secret values, as key-value pairs.
     * 
     */
    public Output<Map<String,String>> values() {
        return this.values;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Secrets(java.lang.String name) {
        this(name, SecretsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Secrets(java.lang.String name, SecretsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Secrets(java.lang.String name, SecretsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("flyio:index:Secrets", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Secrets(java.lang.String name, Output<java.lang.String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("flyio:index:Secrets", name, null, makeResourceOptions(options, id), false);
    }

    private static SecretsArgs makeArgs(SecretsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretsArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "values"
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
    public static Secrets get(java.lang.String name, Output<java.lang.String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Secrets(name, id, options);
    }
}
