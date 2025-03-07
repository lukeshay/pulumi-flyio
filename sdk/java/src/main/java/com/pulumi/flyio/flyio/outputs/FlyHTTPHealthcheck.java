// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.flyio.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.flyio.flyio.outputs.FlyMachineHTTPHeader;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FlyHTTPHealthcheck {
    private @Nullable List<FlyMachineHTTPHeader> headers;
    private @Nullable String method;
    private @Nullable String path;
    private @Nullable Integer port;
    private @Nullable String scheme;
    private @Nullable String tlsServerName;
    private @Nullable Boolean tlsSkipVerify;

    private FlyHTTPHealthcheck() {}
    public List<FlyMachineHTTPHeader> headers() {
        return this.headers == null ? List.of() : this.headers;
    }
    public Optional<String> method() {
        return Optional.ofNullable(this.method);
    }
    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }
    public Optional<String> scheme() {
        return Optional.ofNullable(this.scheme);
    }
    public Optional<String> tlsServerName() {
        return Optional.ofNullable(this.tlsServerName);
    }
    public Optional<Boolean> tlsSkipVerify() {
        return Optional.ofNullable(this.tlsSkipVerify);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FlyHTTPHealthcheck defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<FlyMachineHTTPHeader> headers;
        private @Nullable String method;
        private @Nullable String path;
        private @Nullable Integer port;
        private @Nullable String scheme;
        private @Nullable String tlsServerName;
        private @Nullable Boolean tlsSkipVerify;
        public Builder() {}
        public Builder(FlyHTTPHealthcheck defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.headers = defaults.headers;
    	      this.method = defaults.method;
    	      this.path = defaults.path;
    	      this.port = defaults.port;
    	      this.scheme = defaults.scheme;
    	      this.tlsServerName = defaults.tlsServerName;
    	      this.tlsSkipVerify = defaults.tlsSkipVerify;
        }

        @CustomType.Setter
        public Builder headers(@Nullable List<FlyMachineHTTPHeader> headers) {

            this.headers = headers;
            return this;
        }
        public Builder headers(FlyMachineHTTPHeader... headers) {
            return headers(List.of(headers));
        }
        @CustomType.Setter
        public Builder method(@Nullable String method) {

            this.method = method;
            return this;
        }
        @CustomType.Setter
        public Builder path(@Nullable String path) {

            this.path = path;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {

            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder scheme(@Nullable String scheme) {

            this.scheme = scheme;
            return this;
        }
        @CustomType.Setter
        public Builder tlsServerName(@Nullable String tlsServerName) {

            this.tlsServerName = tlsServerName;
            return this;
        }
        @CustomType.Setter
        public Builder tlsSkipVerify(@Nullable Boolean tlsSkipVerify) {

            this.tlsSkipVerify = tlsSkipVerify;
            return this;
        }
        public FlyHTTPHealthcheck build() {
            final var _resultValue = new FlyHTTPHealthcheck();
            _resultValue.headers = headers;
            _resultValue.method = method;
            _resultValue.path = path;
            _resultValue.port = port;
            _resultValue.scheme = scheme;
            _resultValue.tlsServerName = tlsServerName;
            _resultValue.tlsSkipVerify = tlsSkipVerify;
            return _resultValue;
        }
    }
}
