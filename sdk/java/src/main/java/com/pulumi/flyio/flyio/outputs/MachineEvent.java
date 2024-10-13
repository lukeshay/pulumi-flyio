// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.flyio.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class MachineEvent {
    private String flyId;
    private @Nullable Map<String,Object> request;
    private @Nullable String source;
    private @Nullable String status;
    private @Nullable Integer timestamp;
    private @Nullable String type;

    private MachineEvent() {}
    public String flyId() {
        return this.flyId;
    }
    public Map<String,Object> request() {
        return this.request == null ? Map.of() : this.request;
    }
    public Optional<String> source() {
        return Optional.ofNullable(this.source);
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Optional<Integer> timestamp() {
        return Optional.ofNullable(this.timestamp);
    }
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MachineEvent defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String flyId;
        private @Nullable Map<String,Object> request;
        private @Nullable String source;
        private @Nullable String status;
        private @Nullable Integer timestamp;
        private @Nullable String type;
        public Builder() {}
        public Builder(MachineEvent defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.flyId = defaults.flyId;
    	      this.request = defaults.request;
    	      this.source = defaults.source;
    	      this.status = defaults.status;
    	      this.timestamp = defaults.timestamp;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder flyId(String flyId) {
            if (flyId == null) {
              throw new MissingRequiredPropertyException("MachineEvent", "flyId");
            }
            this.flyId = flyId;
            return this;
        }
        @CustomType.Setter
        public Builder request(@Nullable Map<String,Object> request) {

            this.request = request;
            return this;
        }
        @CustomType.Setter
        public Builder source(@Nullable String source) {

            this.source = source;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {

            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder timestamp(@Nullable Integer timestamp) {

            this.timestamp = timestamp;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {

            this.type = type;
            return this;
        }
        public MachineEvent build() {
            final var _resultValue = new MachineEvent();
            _resultValue.flyId = flyId;
            _resultValue.request = request;
            _resultValue.source = source;
            _resultValue.status = status;
            _resultValue.timestamp = timestamp;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
