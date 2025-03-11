// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AppArgs {
    /**
     * @return Whether to enable subdomains for the application.
     * 
     */
    private @Nullable Boolean enableSubdomains;
    /**
     * @return The name of the Fly.io application.
     * 
     */
    private String name;
    /**
     * @return The network the application belongs to.
     * 
     */
    private @Nullable String network;
    /**
     * @return The organization the application belongs to.
     * 
     */
    private String org;

    private AppArgs() {}
    /**
     * @return Whether to enable subdomains for the application.
     * 
     */
    public Optional<Boolean> enableSubdomains() {
        return Optional.ofNullable(this.enableSubdomains);
    }
    /**
     * @return The name of the Fly.io application.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The network the application belongs to.
     * 
     */
    public Optional<String> network() {
        return Optional.ofNullable(this.network);
    }
    /**
     * @return The organization the application belongs to.
     * 
     */
    public String org() {
        return this.org;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AppArgs defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enableSubdomains;
        private String name;
        private @Nullable String network;
        private String org;
        public Builder() {}
        public Builder(AppArgs defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enableSubdomains = defaults.enableSubdomains;
    	      this.name = defaults.name;
    	      this.network = defaults.network;
    	      this.org = defaults.org;
        }

        @CustomType.Setter
        public Builder enableSubdomains(@Nullable Boolean enableSubdomains) {

            this.enableSubdomains = enableSubdomains;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("AppArgs", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder network(@Nullable String network) {

            this.network = network;
            return this;
        }
        @CustomType.Setter
        public Builder org(String org) {
            if (org == null) {
              throw new MissingRequiredPropertyException("AppArgs", "org");
            }
            this.org = org;
            return this;
        }
        public AppArgs build() {
            final var _resultValue = new AppArgs();
            _resultValue.enableSubdomains = enableSubdomains;
            _resultValue.name = name;
            _resultValue.network = network;
            _resultValue.org = org;
            return _resultValue;
        }
    }
}
