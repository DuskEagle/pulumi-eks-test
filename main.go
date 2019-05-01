package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/go/aws/eks"
	"github.com/pulumi/pulumi-aws/sdk/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

const EKSVersion = "1.12"

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		region, ok := ctx.GetConfig("aws:region")
		if !ok {
			return fmt.Errorf("could not find configured region in Pulumi")
		}

		vpc, err := createVPC(ctx)
		if err != nil {
			return err
		}
		subnets, err := createSubnet(ctx, vpc, region)
		if err != nil {
			return err
		}
		sg, err := createSecurityGroup(ctx, vpc.ID())
		if err != nil {
			return err
		}
		role, err := createEKSIAM(ctx)
		if err != nil {
			return err
		}
		_, err = createEKSCluster(ctx, role, sg, subnets)
		if err != nil {
			return err
		}

		return nil
	})
}

func createVPC(ctx *pulumi.Context) (*ec2.Vpc, error) {
	vpc, err := ec2.NewVpc(ctx, "test-vpc", &ec2.VpcArgs{
		CidrBlock: "10.10.0.0/16",
		Tags: map[string]string{
			"Name": "test-vpc",
		},
	})
	return vpc, err


}

func createSubnet(ctx *pulumi.Context, vpc *ec2.Vpc, region string) ([]*ec2.Subnet, error) {
	cidrRange := "10.10.%d.0/24"
	var subnets []*ec2.Subnet
	for i, az := range []string{"a", "b"} {
		subnet, err := ec2.NewSubnet(ctx, "region_subnet_" + az, &ec2.SubnetArgs{
			AvailabilityZone: region + az,
			CidrBlock:        fmt.Sprintf(cidrRange,i),
			VpcId:            vpc.ID(),
			Tags: map[string]string{
				"Name": fmt.Sprintf("subnet-" + az),
			},
		})
		if err != nil {
			return nil, err
		}
		subnets = append(subnets, subnet)
	}
	return subnets, nil
}

func createEKSCluster(
	ctx *pulumi.Context,
	role *iam.Role,
	sg *ec2.SecurityGroup,
	subnets []*ec2.Subnet) (*eks.Cluster, error) {
	cluster, err := eks.NewCluster(ctx, "test", &eks.ClusterArgs{
		Name: "test",
		RoleArn: role.Arn(),
		Version: EKSVersion,
		VpcConfig: vpcConfig(sg, subnets),
	})
	return cluster, err
}


func vpcConfig(sg *ec2.SecurityGroup, subnets []*ec2.Subnet) map[string][]*pulumi.IDOutput {
	subnetIDs := make([]*pulumi.IDOutput, len(subnets))
	for i, v := range subnets {
		subnetIDs[i] = v.ID()
	}
	return map[string][]*pulumi.IDOutput{
		"securityGroupIds": []*pulumi.IDOutput{sg.ID()},
		"subnetIds": subnetIDs,
	}
}

func createSecurityGroup(ctx *pulumi.Context, vpcId *pulumi.IDOutput) (*ec2.SecurityGroup, error) {
	securityGroup, err := ec2.NewSecurityGroup(ctx, "cluster", &ec2.SecurityGroupArgs{
		NamePrefix: "test",
		Description: "Security group for EKS control plane",
		VpcId: vpcId,
	})
	if err != nil {
		return nil, err
	}

	_, err = ec2.NewSecurityGroupRule(ctx, "cluster_ingress_internet", &ec2.SecurityGroupRuleArgs{
		Description: "Allow ingress traffic from the Internet to the control plane.",
		Protocol: "-1",
		SecurityGroupId: securityGroup,
		CidrBlocks: []string{"0.0.0.0/0"},
		FromPort: 0,
		ToPort: 0,
		Type: "ingress",
	})
	if err != nil {
		return nil, err
	}

	_, err = ec2.NewSecurityGroupRule(ctx, "cluster_egress_internet", &ec2.SecurityGroupRuleArgs{
		Description: "Allow egress traffic from the Internet to the control plane.",
		Protocol: "-1",
		SecurityGroupId: securityGroup,
		CidrBlocks: []string{"0.0.0.0/0"},
		FromPort: 0,
		ToPort: 0,
		Type: "egress",
	})
	if err != nil {
		return nil, err
	}

	return securityGroup, nil
}

func createEKSIAM(ctx *pulumi.Context) (*iam.Role, error) {
	role, err := iam.NewRole(ctx, "cluster", &iam.RoleArgs{
		NamePrefix: "test",
		AssumeRolePolicy: map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []interface{}{
				map[string]interface{}{
					"Sid": "EKSClusterAssumeRole",
					"Effect": "Allow",
					"Principal": map[string]string{
						"Service": "eks.amazonaws.com",
					},
					"Action": "sts:AssumeRole",
				},
			},
		},
		Path: "/",
		ForceDetachPolicies: true,
	})
	if err != nil {
		return nil, err
	}

	// These policies must be present on the IAM role we create an EKS cluster with,
	// or AWS will refuse to create the cluster.
	for _, policyAttachment := range []string{
		"AmazonEKSClusterPolicy", "AmazonEKSServicePolicy",
	} {
		if err := createIAMRolePolicyAttachment(ctx, role, policyAttachment); err != nil {
			return nil, err
		}
	}

	return role, nil
}

// Attach an Amazon-managed IAM policy to an IAM role.
func createIAMRolePolicyAttachment(ctx *pulumi.Context, role *iam.Role, policyAttachment string) error {
	roleName, _, err := role.Name().Value()
	if err != nil {
		return err
	}
	_, err = iam.NewPolicyAttachment(
		ctx,
		fmt.Sprintf("%s-%s", roleName, policyAttachment),
		&iam.PolicyAttachmentArgs{
			PolicyArn: fmt.Sprintf(
				"arn:aws:iam::aws:policy/%s",
				policyAttachment),
			Roles: []*iam.Role{role},
		},
	)
	return err
}
