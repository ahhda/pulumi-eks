// Code generated by pulumi-gen-eks DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of the new access entry. Valid values are STANDARD, FARGATE_LINUX, EC2_LINUX, and EC2_WINDOWS.
// Defaults to STANDARD which provides the standard workflow. EC2_LINUX and EC2_WINDOWS types disallow users to input a kubernetesGroup, and prevent associating access policies.
type AccessEntryType string

const (
	// Standard Access Entry Workflow. Allows users to input a username and kubernetesGroup, and to associate access policies.
	AccessEntryTypeStandard = AccessEntryType("STANDARD")
	// For IAM roles used with AWS Fargate profiles.
	AccessEntryTypeFargateLinux = AccessEntryType("FARGATE_LINUX")
	// For IAM roles associated with self-managed Linux node groups. Allows the nodes to join the cluster.
	AccessEntryTypeEC2Linux = AccessEntryType("EC2_LINUX")
	// For IAM roles associated with self-managed Windows node groups. Allows the nodes to join the cluster.
	AccessEntryTypeEC2Windows = AccessEntryType("EC2_WINDOWS")
)

func (AccessEntryType) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessEntryType)(nil)).Elem()
}

func (e AccessEntryType) ToAccessEntryTypeOutput() AccessEntryTypeOutput {
	return pulumi.ToOutput(e).(AccessEntryTypeOutput)
}

func (e AccessEntryType) ToAccessEntryTypeOutputWithContext(ctx context.Context) AccessEntryTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessEntryTypeOutput)
}

func (e AccessEntryType) ToAccessEntryTypePtrOutput() AccessEntryTypePtrOutput {
	return e.ToAccessEntryTypePtrOutputWithContext(context.Background())
}

func (e AccessEntryType) ToAccessEntryTypePtrOutputWithContext(ctx context.Context) AccessEntryTypePtrOutput {
	return AccessEntryType(e).ToAccessEntryTypeOutputWithContext(ctx).ToAccessEntryTypePtrOutputWithContext(ctx)
}

func (e AccessEntryType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessEntryType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessEntryType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessEntryType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessEntryTypeOutput struct{ *pulumi.OutputState }

func (AccessEntryTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessEntryType)(nil)).Elem()
}

func (o AccessEntryTypeOutput) ToAccessEntryTypeOutput() AccessEntryTypeOutput {
	return o
}

func (o AccessEntryTypeOutput) ToAccessEntryTypeOutputWithContext(ctx context.Context) AccessEntryTypeOutput {
	return o
}

func (o AccessEntryTypeOutput) ToAccessEntryTypePtrOutput() AccessEntryTypePtrOutput {
	return o.ToAccessEntryTypePtrOutputWithContext(context.Background())
}

func (o AccessEntryTypeOutput) ToAccessEntryTypePtrOutputWithContext(ctx context.Context) AccessEntryTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessEntryType) *AccessEntryType {
		return &v
	}).(AccessEntryTypePtrOutput)
}

func (o AccessEntryTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessEntryTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessEntryType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessEntryTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessEntryTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessEntryType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessEntryTypePtrOutput struct{ *pulumi.OutputState }

func (AccessEntryTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessEntryType)(nil)).Elem()
}

func (o AccessEntryTypePtrOutput) ToAccessEntryTypePtrOutput() AccessEntryTypePtrOutput {
	return o
}

func (o AccessEntryTypePtrOutput) ToAccessEntryTypePtrOutputWithContext(ctx context.Context) AccessEntryTypePtrOutput {
	return o
}

func (o AccessEntryTypePtrOutput) Elem() AccessEntryTypeOutput {
	return o.ApplyT(func(v *AccessEntryType) AccessEntryType {
		if v != nil {
			return *v
		}
		var ret AccessEntryType
		return ret
	}).(AccessEntryTypeOutput)
}

func (o AccessEntryTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessEntryTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessEntryType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccessEntryTypeInput is an input type that accepts values of the AccessEntryType enum
// A concrete instance of `AccessEntryTypeInput` can be one of the following:
//
//	AccessEntryTypeStandard
//	AccessEntryTypeFargateLinux
//	AccessEntryTypeEC2Linux
//	AccessEntryTypeEC2Windows
type AccessEntryTypeInput interface {
	pulumi.Input

	ToAccessEntryTypeOutput() AccessEntryTypeOutput
	ToAccessEntryTypeOutputWithContext(context.Context) AccessEntryTypeOutput
}

var accessEntryTypePtrType = reflect.TypeOf((**AccessEntryType)(nil)).Elem()

type AccessEntryTypePtrInput interface {
	pulumi.Input

	ToAccessEntryTypePtrOutput() AccessEntryTypePtrOutput
	ToAccessEntryTypePtrOutputWithContext(context.Context) AccessEntryTypePtrOutput
}

type accessEntryTypePtr string

func AccessEntryTypePtr(v string) AccessEntryTypePtrInput {
	return (*accessEntryTypePtr)(&v)
}

func (*accessEntryTypePtr) ElementType() reflect.Type {
	return accessEntryTypePtrType
}

func (in *accessEntryTypePtr) ToAccessEntryTypePtrOutput() AccessEntryTypePtrOutput {
	return pulumi.ToOutput(in).(AccessEntryTypePtrOutput)
}

func (in *accessEntryTypePtr) ToAccessEntryTypePtrOutputWithContext(ctx context.Context) AccessEntryTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessEntryTypePtrOutput)
}

// The authentication mode of the cluster. Valid values are `CONFIG_MAP`, `API` or `API_AND_CONFIG_MAP`.
//
// See for more details:
// https://docs.aws.amazon.com/eks/latest/userguide/grant-k8s-access.html#set-cam
type AuthenticationMode string

const (
	// Only aws-auth ConfigMap will be used for authenticating to the Kubernetes API.
	AuthenticationModeConfigMap = AuthenticationMode("CONFIG_MAP")
	// Only Access Entries will be used for authenticating to the Kubernetes API.
	AuthenticationModeApi = AuthenticationMode("API")
	// Both aws-auth ConfigMap and Access Entries can be used for authenticating to the Kubernetes API.
	AuthenticationModeApiAndConfigMap = AuthenticationMode("API_AND_CONFIG_MAP")
)

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessEntryTypeInput)(nil)).Elem(), AccessEntryType("STANDARD"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessEntryTypePtrInput)(nil)).Elem(), AccessEntryType("STANDARD"))
	pulumi.RegisterOutputType(AccessEntryTypeOutput{})
	pulumi.RegisterOutputType(AccessEntryTypePtrOutput{})
}