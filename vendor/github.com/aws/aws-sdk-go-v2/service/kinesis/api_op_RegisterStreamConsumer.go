// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a consumer with a Kinesis data stream. When you use this operation,
// the consumer you register can then call SubscribeToShardto receive data from the stream using
// enhanced fan-out, at a rate of up to 2 MiB per second for every shard you
// subscribe to. This rate is unaffected by the total number of consumers that read
// from the same stream.
//
// You can register up to 20 consumers per stream. A given consumer can only be
// registered with one stream at a time.
//
// For an example of how to use this operations, see Enhanced Fan-Out Using the Kinesis Data Streams API.
//
// The use of this operation has a limit of five transactions per second per
// account. Also, only 5 consumers can be created simultaneously. In other words,
// you cannot have more than 5 consumers in a CREATING status at the same time.
// Registering a 6th consumer while there are 5 in a CREATING status results in a
// LimitExceededException .
func (c *Client) RegisterStreamConsumer(ctx context.Context, params *RegisterStreamConsumerInput, optFns ...func(*Options)) (*RegisterStreamConsumerOutput, error) {
	if params == nil {
		params = &RegisterStreamConsumerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterStreamConsumer", params, optFns, c.addOperationRegisterStreamConsumerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterStreamConsumerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterStreamConsumerInput struct {

	// For a given Kinesis data stream, each consumer must have a unique name.
	// However, consumer names don't have to be unique across data streams.
	//
	// This member is required.
	ConsumerName *string

	// The ARN of the Kinesis data stream that you want to register the consumer with.
	// For more info, see [Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces].
	//
	// [Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kinesis-streams
	//
	// This member is required.
	StreamARN *string

	noSmithyDocumentSerde
}

func (in *RegisterStreamConsumerInput) bindEndpointParams(p *EndpointParameters) {
	p.StreamARN = in.StreamARN
	p.OperationType = ptr.String("control")
}

type RegisterStreamConsumerOutput struct {

	// An object that represents the details of the consumer you registered. When you
	// register a consumer, it gets an ARN that is generated by Kinesis Data Streams.
	//
	// This member is required.
	Consumer *types.Consumer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterStreamConsumerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterStreamConsumer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterStreamConsumer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterStreamConsumer"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpRegisterStreamConsumerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterStreamConsumer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterStreamConsumer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterStreamConsumer",
	}
}
