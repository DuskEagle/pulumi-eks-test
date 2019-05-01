# About
A simple example of configuring EKS with Pulumi.

# Using
```
pulumi stack init pulumi-eks-test
pulumi config set aws:region us-east-1
pulumi config set aws:profile <AWS_PROFILE>
pulumi up
```

