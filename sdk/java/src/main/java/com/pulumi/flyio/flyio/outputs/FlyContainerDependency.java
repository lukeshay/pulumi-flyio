// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.flyio.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FlyContainerDependency {
    private @Nullable String condition;
    private @Nullable String name;

    private FlyContainerDependency() {}
    public Optional<String> condition() {
        return Optional.ofNullable(this.condition);
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FlyContainerDependency defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String condition;
        private @Nullable String name;
        public Builder() {}
        public Builder(FlyContainerDependency defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.condition = defaults.condition;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder condition(@Nullable String condition) {

            this.condition = condition;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        public FlyContainerDependency build() {
            final var _resultValue = new FlyContainerDependency();
            _resultValue.condition = condition;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}
