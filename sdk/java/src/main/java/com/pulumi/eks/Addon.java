// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.eks;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.eks.AddonArgs;
import com.pulumi.eks.Utilities;
import javax.annotation.Nullable;

/**
 * Addon manages an EKS add-on.
 * For more information about supported add-ons, see: https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html
 * 
 */
@ResourceType(type="eks:index:Addon")
public class Addon extends com.pulumi.resources.ComponentResource {
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Addon(String name) {
        this(name, AddonArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Addon(String name, AddonArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Addon(String name, AddonArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("eks:index:Addon", name, args == null ? AddonArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}