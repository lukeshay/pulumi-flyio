// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.flyio.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.flyio.flyio.outputs.FlyHTTPOptions;
import com.pulumi.flyio.flyio.outputs.FlyProxyProtoOptions;
import com.pulumi.flyio.flyio.outputs.FlyTLSOptions;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FlyMachinePort {
    private @Nullable Integer endPort;
    private @Nullable Boolean forceHttps;
    private @Nullable List<String> handlers;
    private @Nullable FlyHTTPOptions httpOptions;
    private @Nullable Integer port;
    private @Nullable FlyProxyProtoOptions proxyProtoOptions;
    private @Nullable Integer startPort;
    private @Nullable FlyTLSOptions tlsOptions;

    private FlyMachinePort() {}
    public Optional<Integer> endPort() {
        return Optional.ofNullable(this.endPort);
    }
    public Optional<Boolean> forceHttps() {
        return Optional.ofNullable(this.forceHttps);
    }
    public List<String> handlers() {
        return this.handlers == null ? List.of() : this.handlers;
    }
    public Optional<FlyHTTPOptions> httpOptions() {
        return Optional.ofNullable(this.httpOptions);
    }
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }
    public Optional<FlyProxyProtoOptions> proxyProtoOptions() {
        return Optional.ofNullable(this.proxyProtoOptions);
    }
    public Optional<Integer> startPort() {
        return Optional.ofNullable(this.startPort);
    }
    public Optional<FlyTLSOptions> tlsOptions() {
        return Optional.ofNullable(this.tlsOptions);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FlyMachinePort defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer endPort;
        private @Nullable Boolean forceHttps;
        private @Nullable List<String> handlers;
        private @Nullable FlyHTTPOptions httpOptions;
        private @Nullable Integer port;
        private @Nullable FlyProxyProtoOptions proxyProtoOptions;
        private @Nullable Integer startPort;
        private @Nullable FlyTLSOptions tlsOptions;
        public Builder() {}
        public Builder(FlyMachinePort defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endPort = defaults.endPort;
    	      this.forceHttps = defaults.forceHttps;
    	      this.handlers = defaults.handlers;
    	      this.httpOptions = defaults.httpOptions;
    	      this.port = defaults.port;
    	      this.proxyProtoOptions = defaults.proxyProtoOptions;
    	      this.startPort = defaults.startPort;
    	      this.tlsOptions = defaults.tlsOptions;
        }

        @CustomType.Setter
        public Builder endPort(@Nullable Integer endPort) {

            this.endPort = endPort;
            return this;
        }
        @CustomType.Setter
        public Builder forceHttps(@Nullable Boolean forceHttps) {

            this.forceHttps = forceHttps;
            return this;
        }
        @CustomType.Setter
        public Builder handlers(@Nullable List<String> handlers) {

            this.handlers = handlers;
            return this;
        }
        public Builder handlers(String... handlers) {
            return handlers(List.of(handlers));
        }
        @CustomType.Setter
        public Builder httpOptions(@Nullable FlyHTTPOptions httpOptions) {

            this.httpOptions = httpOptions;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {

            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder proxyProtoOptions(@Nullable FlyProxyProtoOptions proxyProtoOptions) {

            this.proxyProtoOptions = proxyProtoOptions;
            return this;
        }
        @CustomType.Setter
        public Builder startPort(@Nullable Integer startPort) {

            this.startPort = startPort;
            return this;
        }
        @CustomType.Setter
        public Builder tlsOptions(@Nullable FlyTLSOptions tlsOptions) {

            this.tlsOptions = tlsOptions;
            return this;
        }
        public FlyMachinePort build() {
            final var _resultValue = new FlyMachinePort();
            _resultValue.endPort = endPort;
            _resultValue.forceHttps = forceHttps;
            _resultValue.handlers = handlers;
            _resultValue.httpOptions = httpOptions;
            _resultValue.port = port;
            _resultValue.proxyProtoOptions = proxyProtoOptions;
            _resultValue.startPort = startPort;
            _resultValue.tlsOptions = tlsOptions;
            return _resultValue;
        }
    }
}
