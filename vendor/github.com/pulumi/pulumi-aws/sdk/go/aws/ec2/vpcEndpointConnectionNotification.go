// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a VPC Endpoint connection notification resource.
// Connection notifications notify subscribers of VPC Endpoint events.
type VpcEndpointConnectionNotification struct {
	s *pulumi.ResourceState
}

// NewVpcEndpointConnectionNotification registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointConnectionNotification(ctx *pulumi.Context,
	name string, args *VpcEndpointConnectionNotificationArgs, opts ...pulumi.ResourceOpt) (*VpcEndpointConnectionNotification, error) {
	if args == nil || args.ConnectionEvents == nil {
		return nil, errors.New("missing required argument 'ConnectionEvents'")
	}
	if args == nil || args.ConnectionNotificationArn == nil {
		return nil, errors.New("missing required argument 'ConnectionNotificationArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["connectionEvents"] = nil
		inputs["connectionNotificationArn"] = nil
		inputs["vpcEndpointId"] = nil
		inputs["vpcEndpointServiceId"] = nil
	} else {
		inputs["connectionEvents"] = args.ConnectionEvents
		inputs["connectionNotificationArn"] = args.ConnectionNotificationArn
		inputs["vpcEndpointId"] = args.VpcEndpointId
		inputs["vpcEndpointServiceId"] = args.VpcEndpointServiceId
	}
	inputs["notificationType"] = nil
	inputs["state"] = nil
	s, err := ctx.RegisterResource("aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcEndpointConnectionNotification{s: s}, nil
}

// GetVpcEndpointConnectionNotification gets an existing VpcEndpointConnectionNotification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointConnectionNotification(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VpcEndpointConnectionNotificationState, opts ...pulumi.ResourceOpt) (*VpcEndpointConnectionNotification, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["connectionEvents"] = state.ConnectionEvents
		inputs["connectionNotificationArn"] = state.ConnectionNotificationArn
		inputs["notificationType"] = state.NotificationType
		inputs["state"] = state.State
		inputs["vpcEndpointId"] = state.VpcEndpointId
		inputs["vpcEndpointServiceId"] = state.VpcEndpointServiceId
	}
	s, err := ctx.ReadResource("aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcEndpointConnectionNotification{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VpcEndpointConnectionNotification) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VpcEndpointConnectionNotification) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
func (r *VpcEndpointConnectionNotification) ConnectionEvents() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["connectionEvents"])
}

// The ARN of the SNS topic for the notifications.
func (r *VpcEndpointConnectionNotification) ConnectionNotificationArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["connectionNotificationArn"])
}

// The type of notification.
func (r *VpcEndpointConnectionNotification) NotificationType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["notificationType"])
}

// The state of the notification.
func (r *VpcEndpointConnectionNotification) State() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["state"])
}

// The ID of the VPC Endpoint to receive notifications for.
func (r *VpcEndpointConnectionNotification) VpcEndpointId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcEndpointId"])
}

// The ID of the VPC Endpoint Service to receive notifications for.
func (r *VpcEndpointConnectionNotification) VpcEndpointServiceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcEndpointServiceId"])
}

// Input properties used for looking up and filtering VpcEndpointConnectionNotification resources.
type VpcEndpointConnectionNotificationState struct {
	// One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
	ConnectionEvents interface{}
	// The ARN of the SNS topic for the notifications.
	ConnectionNotificationArn interface{}
	// The type of notification.
	NotificationType interface{}
	// The state of the notification.
	State interface{}
	// The ID of the VPC Endpoint to receive notifications for.
	VpcEndpointId interface{}
	// The ID of the VPC Endpoint Service to receive notifications for.
	VpcEndpointServiceId interface{}
}

// The set of arguments for constructing a VpcEndpointConnectionNotification resource.
type VpcEndpointConnectionNotificationArgs struct {
	// One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
	ConnectionEvents interface{}
	// The ARN of the SNS topic for the notifications.
	ConnectionNotificationArn interface{}
	// The ID of the VPC Endpoint to receive notifications for.
	VpcEndpointId interface{}
	// The ID of the VPC Endpoint Service to receive notifications for.
	VpcEndpointServiceId interface{}
}
