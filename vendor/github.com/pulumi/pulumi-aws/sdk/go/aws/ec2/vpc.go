// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a VPC resource.
type Vpc struct {
	s *pulumi.ResourceState
}

// NewVpc registers a new resource with the given unique name, arguments, and options.
func NewVpc(ctx *pulumi.Context,
	name string, args *VpcArgs, opts ...pulumi.ResourceOpt) (*Vpc, error) {
	if args == nil || args.CidrBlock == nil {
		return nil, errors.New("missing required argument 'CidrBlock'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["assignGeneratedIpv6CidrBlock"] = nil
		inputs["cidrBlock"] = nil
		inputs["enableClassiclink"] = nil
		inputs["enableClassiclinkDnsSupport"] = nil
		inputs["enableDnsHostnames"] = nil
		inputs["enableDnsSupport"] = nil
		inputs["instanceTenancy"] = nil
		inputs["tags"] = nil
	} else {
		inputs["assignGeneratedIpv6CidrBlock"] = args.AssignGeneratedIpv6CidrBlock
		inputs["cidrBlock"] = args.CidrBlock
		inputs["enableClassiclink"] = args.EnableClassiclink
		inputs["enableClassiclinkDnsSupport"] = args.EnableClassiclinkDnsSupport
		inputs["enableDnsHostnames"] = args.EnableDnsHostnames
		inputs["enableDnsSupport"] = args.EnableDnsSupport
		inputs["instanceTenancy"] = args.InstanceTenancy
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	inputs["defaultNetworkAclId"] = nil
	inputs["defaultRouteTableId"] = nil
	inputs["defaultSecurityGroupId"] = nil
	inputs["dhcpOptionsId"] = nil
	inputs["ipv6AssociationId"] = nil
	inputs["ipv6CidrBlock"] = nil
	inputs["mainRouteTableId"] = nil
	inputs["ownerId"] = nil
	s, err := ctx.RegisterResource("aws:ec2/vpc:Vpc", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Vpc{s: s}, nil
}

// GetVpc gets an existing Vpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpc(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VpcState, opts ...pulumi.ResourceOpt) (*Vpc, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["assignGeneratedIpv6CidrBlock"] = state.AssignGeneratedIpv6CidrBlock
		inputs["cidrBlock"] = state.CidrBlock
		inputs["defaultNetworkAclId"] = state.DefaultNetworkAclId
		inputs["defaultRouteTableId"] = state.DefaultRouteTableId
		inputs["defaultSecurityGroupId"] = state.DefaultSecurityGroupId
		inputs["dhcpOptionsId"] = state.DhcpOptionsId
		inputs["enableClassiclink"] = state.EnableClassiclink
		inputs["enableClassiclinkDnsSupport"] = state.EnableClassiclinkDnsSupport
		inputs["enableDnsHostnames"] = state.EnableDnsHostnames
		inputs["enableDnsSupport"] = state.EnableDnsSupport
		inputs["instanceTenancy"] = state.InstanceTenancy
		inputs["ipv6AssociationId"] = state.Ipv6AssociationId
		inputs["ipv6CidrBlock"] = state.Ipv6CidrBlock
		inputs["mainRouteTableId"] = state.MainRouteTableId
		inputs["ownerId"] = state.OwnerId
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:ec2/vpc:Vpc", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Vpc{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Vpc) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Vpc) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Amazon Resource Name (ARN) of VPC
func (r *Vpc) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Requests an Amazon-provided IPv6 CIDR
// block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
// the size of the CIDR block. Default is `false`.
func (r *Vpc) AssignGeneratedIpv6CidrBlock() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["assignGeneratedIpv6CidrBlock"])
}

// The CIDR block for the VPC.
func (r *Vpc) CidrBlock() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cidrBlock"])
}

// The ID of the network ACL created by default on VPC creation
func (r *Vpc) DefaultNetworkAclId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultNetworkAclId"])
}

// The ID of the route table created by default on VPC creation
func (r *Vpc) DefaultRouteTableId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultRouteTableId"])
}

// The ID of the security group created by default on VPC creation
func (r *Vpc) DefaultSecurityGroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultSecurityGroupId"])
}

func (r *Vpc) DhcpOptionsId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dhcpOptionsId"])
}

// A boolean flag to enable/disable ClassicLink
// for the VPC. Only valid in regions and accounts that support EC2 Classic.
// See the [ClassicLink documentation][1] for more information. Defaults false.
func (r *Vpc) EnableClassiclink() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableClassiclink"])
}

// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
// Only valid in regions and accounts that support EC2 Classic.
func (r *Vpc) EnableClassiclinkDnsSupport() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableClassiclinkDnsSupport"])
}

// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
func (r *Vpc) EnableDnsHostnames() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableDnsHostnames"])
}

// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
func (r *Vpc) EnableDnsSupport() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableDnsSupport"])
}

// A tenancy option for instances launched into the VPC
func (r *Vpc) InstanceTenancy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceTenancy"])
}

// The association ID for the IPv6 CIDR block.
func (r *Vpc) Ipv6AssociationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipv6AssociationId"])
}

// The IPv6 CIDR block.
func (r *Vpc) Ipv6CidrBlock() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipv6CidrBlock"])
}

// The ID of the main route table associated with
// this VPC. Note that you can change a VPC's main route table by using an
// [`aws_main_route_table_association`](https://www.terraform.io/docs/providers/aws/r/main_route_table_assoc.html).
func (r *Vpc) MainRouteTableId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["mainRouteTableId"])
}

// The ID of the AWS account that owns the VPC.
func (r *Vpc) OwnerId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ownerId"])
}

// A mapping of tags to assign to the resource.
func (r *Vpc) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Vpc resources.
type VpcState struct {
	// Amazon Resource Name (ARN) of VPC
	Arn interface{}
	// Requests an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
	// the size of the CIDR block. Default is `false`.
	AssignGeneratedIpv6CidrBlock interface{}
	// The CIDR block for the VPC.
	CidrBlock interface{}
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId interface{}
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId interface{}
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId interface{}
	DhcpOptionsId interface{}
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation][1] for more information. Defaults false.
	EnableClassiclink interface{}
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	EnableClassiclinkDnsSupport interface{}
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames interface{}
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport interface{}
	// A tenancy option for instances launched into the VPC
	InstanceTenancy interface{}
	// The association ID for the IPv6 CIDR block.
	Ipv6AssociationId interface{}
	// The IPv6 CIDR block.
	Ipv6CidrBlock interface{}
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// [`aws_main_route_table_association`](https://www.terraform.io/docs/providers/aws/r/main_route_table_assoc.html).
	MainRouteTableId interface{}
	// The ID of the AWS account that owns the VPC.
	OwnerId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a Vpc resource.
type VpcArgs struct {
	// Requests an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
	// the size of the CIDR block. Default is `false`.
	AssignGeneratedIpv6CidrBlock interface{}
	// The CIDR block for the VPC.
	CidrBlock interface{}
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation][1] for more information. Defaults false.
	EnableClassiclink interface{}
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	EnableClassiclinkDnsSupport interface{}
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames interface{}
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport interface{}
	// A tenancy option for instances launched into the VPC
	InstanceTenancy interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
