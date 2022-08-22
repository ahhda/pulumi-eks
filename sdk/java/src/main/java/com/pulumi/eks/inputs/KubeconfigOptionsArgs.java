// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.eks.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * Represents the AWS credentials to scope a given kubeconfig when using a non-default credential chain.
 * 
 * The options can be used independently, or additively.
 * 
 * A scoped kubeconfig is necessary for certain auth scenarios. For example:
 *   1. Assume a role on the default account caller,
 *   2. Use an AWS creds profile instead of the default account caller,
 *   3. Use an AWS creds creds profile instead of the default account caller,
 *      and then assume a given role on the profile. This scenario is also
 *      possible by only using a profile, iff the profile includes a role to
 *      assume in its settings.
 * 
 * See for more details:
 * - https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html
 * - https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html
 * - https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html
 * 
 */
public final class KubeconfigOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubeconfigOptionsArgs Empty = new KubeconfigOptionsArgs();

    /**
     * AWS credential profile name to always use instead of the default AWS credential provider chain.
     * 
     * The profile is passed to kubeconfig as an authentication environment setting.
     * 
     */
    @Import(name="profileName")
    private @Nullable Output<String> profileName;

    /**
     * @return AWS credential profile name to always use instead of the default AWS credential provider chain.
     * 
     * The profile is passed to kubeconfig as an authentication environment setting.
     * 
     */
    public Optional<Output<String>> profileName() {
        return Optional.ofNullable(this.profileName);
    }

    /**
     * Role ARN to assume instead of the default AWS credential provider chain.
     * 
     * The role is passed to kubeconfig as an authentication exec argument.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return Role ARN to assume instead of the default AWS credential provider chain.
     * 
     * The role is passed to kubeconfig as an authentication exec argument.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    private KubeconfigOptionsArgs() {}

    private KubeconfigOptionsArgs(KubeconfigOptionsArgs $) {
        this.profileName = $.profileName;
        this.roleArn = $.roleArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeconfigOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeconfigOptionsArgs $;

        public Builder() {
            $ = new KubeconfigOptionsArgs();
        }

        public Builder(KubeconfigOptionsArgs defaults) {
            $ = new KubeconfigOptionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param profileName AWS credential profile name to always use instead of the default AWS credential provider chain.
         * 
         * The profile is passed to kubeconfig as an authentication environment setting.
         * 
         * @return builder
         * 
         */
        public Builder profileName(@Nullable Output<String> profileName) {
            $.profileName = profileName;
            return this;
        }

        /**
         * @param profileName AWS credential profile name to always use instead of the default AWS credential provider chain.
         * 
         * The profile is passed to kubeconfig as an authentication environment setting.
         * 
         * @return builder
         * 
         */
        public Builder profileName(String profileName) {
            return profileName(Output.of(profileName));
        }

        /**
         * @param roleArn Role ARN to assume instead of the default AWS credential provider chain.
         * 
         * The role is passed to kubeconfig as an authentication exec argument.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn Role ARN to assume instead of the default AWS credential provider chain.
         * 
         * The role is passed to kubeconfig as an authentication exec argument.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        public KubeconfigOptionsArgs build() {
            return $;
        }
    }

}