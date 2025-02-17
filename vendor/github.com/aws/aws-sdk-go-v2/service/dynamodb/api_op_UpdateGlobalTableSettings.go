// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates settings for a global table.
//
// This documentation is for version 2017.11.29 (Legacy) of global tables, which
// should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)]when possible,
// because it provides greater flexibility, higher efficiency, and consumes less
// write capacity than 2017.11.29 (Legacy).
//
// To determine which version you're using, see [Determining the global table version you are using]. To update existing global tables
// from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables].
//
// [Global Tables version 2019.11.21 (Current)]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Upgrading global tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html
// [Determining the global table version you are using]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html
func (c *Client) UpdateGlobalTableSettings(ctx context.Context, params *UpdateGlobalTableSettingsInput, optFns ...func(*Options)) (*UpdateGlobalTableSettingsOutput, error) {
	if params == nil {
		params = &UpdateGlobalTableSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGlobalTableSettings", params, optFns, c.addOperationUpdateGlobalTableSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGlobalTableSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGlobalTableSettingsInput struct {

	// The name of the global table
	//
	// This member is required.
	GlobalTableName *string

	// The billing mode of the global table. If GlobalTableBillingMode is not
	// specified, the global table defaults to PROVISIONED capacity billing mode.
	//
	//   - PROVISIONED - We recommend using PROVISIONED for predictable workloads.
	//   PROVISIONED sets the billing mode to [Provisioned capacity mode].
	//
	//   - PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable
	//   workloads. PAY_PER_REQUEST sets the billing mode to [On-demand capacity mode].
	//
	// [Provisioned capacity mode]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/provisioned-capacity-mode.html
	// [On-demand capacity mode]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/on-demand-capacity-mode.html
	GlobalTableBillingMode types.BillingMode

	// Represents the settings of a global secondary index for a global table that
	// will be modified.
	GlobalTableGlobalSecondaryIndexSettingsUpdate []types.GlobalTableGlobalSecondaryIndexSettingsUpdate

	// Auto scaling settings for managing provisioned write capacity for the global
	// table.
	GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate *types.AutoScalingSettingsUpdate

	// The maximum number of writes consumed per second before DynamoDB returns a
	// ThrottlingException.
	GlobalTableProvisionedWriteCapacityUnits *int64

	// Represents the settings for a global table in a Region that will be modified.
	ReplicaSettingsUpdate []types.ReplicaSettingsUpdate

	noSmithyDocumentSerde
}

type UpdateGlobalTableSettingsOutput struct {

	// The name of the global table.
	GlobalTableName *string

	// The Region-specific settings for the global table.
	ReplicaSettings []types.ReplicaSettingsDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGlobalTableSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateGlobalTableSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateGlobalTableSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateGlobalTableSettings"); err != nil {
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
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addOpUpdateGlobalTableSettingsDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addUserAgentAccountIDEndpointMode(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateGlobalTableSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGlobalTableSettings(options.Region), middleware.Before); err != nil {
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
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func addOpUpdateGlobalTableSettingsDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpUpdateGlobalTableSettingsDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
		Region:                       o.Region,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpUpdateGlobalTableSettingsDiscoverEndpoint(ctx context.Context, region string, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*UpdateGlobalTableSettingsInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)
	identifierMap["sdk#Region"] = region

	key := fmt.Sprintf("DynamoDB.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	go c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, region, key, opt)
	return internalEndpointDiscovery.WeightedAddress{}, nil
}

func newServiceMetadataMiddleware_opUpdateGlobalTableSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateGlobalTableSettings",
	}
}
