// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a deployment group to which application revisions are deployed.
func (c *Client) CreateDeploymentGroup(ctx context.Context, params *CreateDeploymentGroupInput, optFns ...func(*Options)) (*CreateDeploymentGroupOutput, error) {
	if params == nil {
		params = &CreateDeploymentGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeploymentGroup", params, optFns, c.addOperationCreateDeploymentGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeploymentGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateDeploymentGroup operation.
type CreateDeploymentGroupInput struct {

	// The name of an CodeDeploy application associated with the user or Amazon Web
	// Services account.
	//
	// This member is required.
	ApplicationName *string

	// The name of a new deployment group for the specified application.
	//
	// This member is required.
	DeploymentGroupName *string

	// A service role Amazon Resource Name (ARN) that allows CodeDeploy to act on the
	// user's behalf when interacting with Amazon Web Services services.
	//
	// This member is required.
	ServiceRoleArn *string

	// Information to add about Amazon CloudWatch alarms when the deployment group is
	// created.
	AlarmConfiguration *types.AlarmConfiguration

	// Configuration information for an automatic rollback that is added when a
	// deployment group is created.
	AutoRollbackConfiguration *types.AutoRollbackConfiguration

	// A list of associated Amazon EC2 Auto Scaling groups.
	AutoScalingGroups []string

	// Information about blue/green deployment options for a deployment group.
	BlueGreenDeploymentConfiguration *types.BlueGreenDeploymentConfiguration

	// If specified, the deployment configuration name can be either one of the
	// predefined configurations provided with CodeDeploy or a custom deployment
	// configuration that you create by calling the create deployment configuration
	// operation.
	//
	// CodeDeployDefault.OneAtATime is the default deployment configuration. It is
	// used if a configuration isn't specified for the deployment or deployment group.
	//
	// For more information about the predefined deployment configurations in
	// CodeDeploy, see [Working with Deployment Configurations in CodeDeploy]in the CodeDeploy User Guide.
	//
	// [Working with Deployment Configurations in CodeDeploy]: https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations.html
	DeploymentConfigName *string

	// Information about the type of deployment, in-place or blue/green, that you want
	// to run and whether to route deployment traffic behind a load balancer.
	DeploymentStyle *types.DeploymentStyle

	// The Amazon EC2 tags on which to filter. The deployment group includes Amazon
	// EC2 instances with any of the specified tags. Cannot be used in the same call as
	// ec2TagSet.
	Ec2TagFilters []types.EC2TagFilter

	// Information about groups of tags applied to Amazon EC2 instances. The
	// deployment group includes only Amazon EC2 instances identified by all the tag
	// groups. Cannot be used in the same call as ec2TagFilters .
	Ec2TagSet *types.EC2TagSet

	//  The target Amazon ECS services in the deployment group. This applies only to
	// deployment groups that use the Amazon ECS compute platform. A target Amazon ECS
	// service is specified as an Amazon ECS cluster and service name pair using the
	// format : .
	EcsServices []types.ECSService

	// Information about the load balancer used in a deployment.
	LoadBalancerInfo *types.LoadBalancerInfo

	// The on-premises instance tags on which to filter. The deployment group includes
	// on-premises instances with any of the specified tags. Cannot be used in the same
	// call as OnPremisesTagSet .
	OnPremisesInstanceTagFilters []types.TagFilter

	// Information about groups of tags applied to on-premises instances. The
	// deployment group includes only on-premises instances identified by all of the
	// tag groups. Cannot be used in the same call as onPremisesInstanceTagFilters .
	OnPremisesTagSet *types.OnPremisesTagSet

	// Indicates what happens when new Amazon EC2 instances are launched
	// mid-deployment and do not receive the deployed application revision.
	//
	// If this option is set to UPDATE or is unspecified, CodeDeploy initiates one or
	// more 'auto-update outdated instances' deployments to apply the deployed
	// application revision to the new Amazon EC2 instances.
	//
	// If this option is set to IGNORE , CodeDeploy does not initiate a deployment to
	// update the new Amazon EC2 instances. This may result in instances having
	// different revisions.
	OutdatedInstancesStrategy types.OutdatedInstancesStrategy

	//  The metadata that you apply to CodeDeploy deployment groups to help you
	// organize and categorize them. Each tag consists of a key and an optional value,
	// both of which you define.
	Tags []types.Tag

	// This parameter only applies if you are using CodeDeploy with Amazon EC2 Auto
	// Scaling. For more information, see [Integrating CodeDeploy with Amazon EC2 Auto Scaling]in the CodeDeploy User Guide.
	//
	// Set terminationHookEnabled to true to have CodeDeploy install a termination
	// hook into your Auto Scaling group when you create a deployment group. When this
	// hook is installed, CodeDeploy will perform termination deployments.
	//
	// For information about termination deployments, see [Enabling termination deployments during Auto Scaling scale-in events] in the CodeDeploy User
	// Guide.
	//
	// For more information about Auto Scaling scale-in events, see the [Scale in] topic in the
	// Amazon EC2 Auto Scaling User Guide.
	//
	// [Enabling termination deployments during Auto Scaling scale-in events]: https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-auto-scaling.html#integrations-aws-auto-scaling-behaviors-hook-enable
	// [Integrating CodeDeploy with Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-auto-scaling.html
	// [Scale in]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-lifecycle.html#as-lifecycle-scale-in
	TerminationHookEnabled *bool

	// Information about triggers to create when the deployment group is created. For
	// examples, see [Create a Trigger for an CodeDeploy Event]in the CodeDeploy User Guide.
	//
	// [Create a Trigger for an CodeDeploy Event]: https://docs.aws.amazon.com/codedeploy/latest/userguide/how-to-notify-sns.html
	TriggerConfigurations []types.TriggerConfig

	noSmithyDocumentSerde
}

// Represents the output of a CreateDeploymentGroup operation.
type CreateDeploymentGroupOutput struct {

	// A unique deployment group ID.
	DeploymentGroupId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDeploymentGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDeploymentGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDeploymentGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDeploymentGroup"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateDeploymentGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeploymentGroup(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateDeploymentGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDeploymentGroup",
	}
}
