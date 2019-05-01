// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create an association between a subnet and routing table.
type RouteTableAssociation struct {
	s *pulumi.ResourceState
}

// NewRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewRouteTableAssociation(ctx *pulumi.Context,
	name string, args *RouteTableAssociationArgs, opts ...pulumi.ResourceOpt) (*RouteTableAssociation, error) {
	if args == nil || args.RouteTableId == nil {
		return nil, errors.New("missing required argument 'RouteTableId'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["routeTableId"] = nil
		inputs["subnetId"] = nil
	} else {
		inputs["routeTableId"] = args.RouteTableId
		inputs["subnetId"] = args.SubnetId
	}
	s, err := ctx.RegisterResource("aws:ec2/routeTableAssociation:RouteTableAssociation", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RouteTableAssociation{s: s}, nil
}

// GetRouteTableAssociation gets an existing RouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RouteTableAssociationState, opts ...pulumi.ResourceOpt) (*RouteTableAssociation, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["routeTableId"] = state.RouteTableId
		inputs["subnetId"] = state.SubnetId
	}
	s, err := ctx.ReadResource("aws:ec2/routeTableAssociation:RouteTableAssociation", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RouteTableAssociation{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RouteTableAssociation) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RouteTableAssociation) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ID of the routing table to associate with.
func (r *RouteTableAssociation) RouteTableId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["routeTableId"])
}

// The subnet ID to create an association.
func (r *RouteTableAssociation) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// Input properties used for looking up and filtering RouteTableAssociation resources.
type RouteTableAssociationState struct {
	// The ID of the routing table to associate with.
	RouteTableId interface{}
	// The subnet ID to create an association.
	SubnetId interface{}
}

// The set of arguments for constructing a RouteTableAssociation resource.
type RouteTableAssociationArgs struct {
	// The ID of the routing table to associate with.
	RouteTableId interface{}
	// The subnet ID to create an association.
	SubnetId interface{}
}
