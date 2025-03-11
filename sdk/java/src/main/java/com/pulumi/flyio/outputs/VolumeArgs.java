// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.flyio.flyio.outputs.FlyMachineGuest;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VolumeArgs {
    /**
     * @return The Fly.io application to attach the volume to.
     * 
     */
    private String app;
    /**
     * @return Whether to enable automatic backups for the volume.
     * 
     */
    private @Nullable Boolean autoBackupEnabled;
    private @Nullable FlyMachineGuest compute;
    private @Nullable String computeImage;
    private @Nullable Boolean encrypted;
    private @Nullable String fstype;
    private @Nullable String name;
    private @Nullable String region;
    private @Nullable Boolean requireUniqueZone;
    private @Nullable Integer sizeGb;
    private @Nullable String snapshotId;
    private @Nullable Integer snapshotRetention;
    private @Nullable String sourceVolumeId;
    private @Nullable Boolean uniqueZoneAppWide;

    private VolumeArgs() {}
    /**
     * @return The Fly.io application to attach the volume to.
     * 
     */
    public String app() {
        return this.app;
    }
    /**
     * @return Whether to enable automatic backups for the volume.
     * 
     */
    public Optional<Boolean> autoBackupEnabled() {
        return Optional.ofNullable(this.autoBackupEnabled);
    }
    public Optional<FlyMachineGuest> compute() {
        return Optional.ofNullable(this.compute);
    }
    public Optional<String> computeImage() {
        return Optional.ofNullable(this.computeImage);
    }
    public Optional<Boolean> encrypted() {
        return Optional.ofNullable(this.encrypted);
    }
    public Optional<String> fstype() {
        return Optional.ofNullable(this.fstype);
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }
    public Optional<Boolean> requireUniqueZone() {
        return Optional.ofNullable(this.requireUniqueZone);
    }
    public Optional<Integer> sizeGb() {
        return Optional.ofNullable(this.sizeGb);
    }
    public Optional<String> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }
    public Optional<Integer> snapshotRetention() {
        return Optional.ofNullable(this.snapshotRetention);
    }
    public Optional<String> sourceVolumeId() {
        return Optional.ofNullable(this.sourceVolumeId);
    }
    public Optional<Boolean> uniqueZoneAppWide() {
        return Optional.ofNullable(this.uniqueZoneAppWide);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VolumeArgs defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String app;
        private @Nullable Boolean autoBackupEnabled;
        private @Nullable FlyMachineGuest compute;
        private @Nullable String computeImage;
        private @Nullable Boolean encrypted;
        private @Nullable String fstype;
        private @Nullable String name;
        private @Nullable String region;
        private @Nullable Boolean requireUniqueZone;
        private @Nullable Integer sizeGb;
        private @Nullable String snapshotId;
        private @Nullable Integer snapshotRetention;
        private @Nullable String sourceVolumeId;
        private @Nullable Boolean uniqueZoneAppWide;
        public Builder() {}
        public Builder(VolumeArgs defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.app = defaults.app;
    	      this.autoBackupEnabled = defaults.autoBackupEnabled;
    	      this.compute = defaults.compute;
    	      this.computeImage = defaults.computeImage;
    	      this.encrypted = defaults.encrypted;
    	      this.fstype = defaults.fstype;
    	      this.name = defaults.name;
    	      this.region = defaults.region;
    	      this.requireUniqueZone = defaults.requireUniqueZone;
    	      this.sizeGb = defaults.sizeGb;
    	      this.snapshotId = defaults.snapshotId;
    	      this.snapshotRetention = defaults.snapshotRetention;
    	      this.sourceVolumeId = defaults.sourceVolumeId;
    	      this.uniqueZoneAppWide = defaults.uniqueZoneAppWide;
        }

        @CustomType.Setter
        public Builder app(String app) {
            if (app == null) {
              throw new MissingRequiredPropertyException("VolumeArgs", "app");
            }
            this.app = app;
            return this;
        }
        @CustomType.Setter
        public Builder autoBackupEnabled(@Nullable Boolean autoBackupEnabled) {

            this.autoBackupEnabled = autoBackupEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder compute(@Nullable FlyMachineGuest compute) {

            this.compute = compute;
            return this;
        }
        @CustomType.Setter
        public Builder computeImage(@Nullable String computeImage) {

            this.computeImage = computeImage;
            return this;
        }
        @CustomType.Setter
        public Builder encrypted(@Nullable Boolean encrypted) {

            this.encrypted = encrypted;
            return this;
        }
        @CustomType.Setter
        public Builder fstype(@Nullable String fstype) {

            this.fstype = fstype;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {

            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder requireUniqueZone(@Nullable Boolean requireUniqueZone) {

            this.requireUniqueZone = requireUniqueZone;
            return this;
        }
        @CustomType.Setter
        public Builder sizeGb(@Nullable Integer sizeGb) {

            this.sizeGb = sizeGb;
            return this;
        }
        @CustomType.Setter
        public Builder snapshotId(@Nullable String snapshotId) {

            this.snapshotId = snapshotId;
            return this;
        }
        @CustomType.Setter
        public Builder snapshotRetention(@Nullable Integer snapshotRetention) {

            this.snapshotRetention = snapshotRetention;
            return this;
        }
        @CustomType.Setter
        public Builder sourceVolumeId(@Nullable String sourceVolumeId) {

            this.sourceVolumeId = sourceVolumeId;
            return this;
        }
        @CustomType.Setter
        public Builder uniqueZoneAppWide(@Nullable Boolean uniqueZoneAppWide) {

            this.uniqueZoneAppWide = uniqueZoneAppWide;
            return this;
        }
        public VolumeArgs build() {
            final var _resultValue = new VolumeArgs();
            _resultValue.app = app;
            _resultValue.autoBackupEnabled = autoBackupEnabled;
            _resultValue.compute = compute;
            _resultValue.computeImage = computeImage;
            _resultValue.encrypted = encrypted;
            _resultValue.fstype = fstype;
            _resultValue.name = name;
            _resultValue.region = region;
            _resultValue.requireUniqueZone = requireUniqueZone;
            _resultValue.sizeGb = sizeGb;
            _resultValue.snapshotId = snapshotId;
            _resultValue.snapshotRetention = snapshotRetention;
            _resultValue.sourceVolumeId = sourceVolumeId;
            _resultValue.uniqueZoneAppWide = uniqueZoneAppWide;
            return _resultValue;
        }
    }
}
