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

// Disables server-side encryption for a specified stream.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// Stopping encryption is an asynchronous operation. Upon receiving the request,
// Kinesis Data Streams returns immediately and sets the status of the stream to
// UPDATING . After the update is complete, Kinesis Data Streams sets the status of
// the stream back to ACTIVE . Stopping encryption normally takes a few seconds to
// complete, but it can take minutes. You can continue to read and write data to
// your stream while its status is UPDATING . Once the status of the stream is
// ACTIVE , records written to the stream are no longer encrypted by Kinesis Data
// Streams.
//
// API Limits: You can successfully disable server-side encryption 25 times in a
// rolling 24-hour period.
//
// Note: It can take up to 5 seconds after the stream is in an ACTIVE status
// before all records written to the stream are no longer subject to encryption.
// After you disabled encryption, you can verify that encryption is not applied by
// inspecting the API response from PutRecord or PutRecords .
func (c *Client) StopStreamEncryption(ctx context.Context, params *StopStreamEncryptionInput, optFns ...func(*Options)) (*StopStreamEncryptionOutput, error) {
	if params == nil {
		params = &StopStreamEncryptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopStreamEncryption", params, optFns, c.addOperationStopStreamEncryptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopStreamEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopStreamEncryptionInput struct {

	// The encryption type. The only valid value is KMS .
	//
	// This member is required.
	EncryptionType types.EncryptionType

	// The GUID for the customer-managed Amazon Web Services KMS key to use for
	// encryption. This value can be a globally unique identifier, a fully specified
	// Amazon Resource Name (ARN) to either an alias or a key, or an alias name
	// prefixed by "alias/".You can also use a master key owned by Kinesis Data Streams
	// by specifying the alias aws/kinesis .
	//
	//   - Key ARN example:
	//   arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//   - Alias ARN example: arn:aws:kms:us-east-1:123456789012:alias/MyAliasName
	//
	//   - Globally unique key ID example: 12345678-1234-1234-1234-123456789012
	//
	//   - Alias name example: alias/MyAliasName
	//
	//   - Master key owned by Kinesis Data Streams: alias/aws/kinesis
	//
	// This member is required.
	KeyId *string

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream on which to stop encrypting records.
	StreamName *string

	noSmithyDocumentSerde
}

func (in *StopStreamEncryptionInput) bindEndpointParams(p *EndpointParameters) {
	p.StreamARN = in.StreamARN
	p.OperationType = ptr.String("control")
}

type StopStreamEncryptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopStreamEncryptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopStreamEncryption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopStreamEncryption{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StopStreamEncryption"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpStopStreamEncryptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopStreamEncryption(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopStreamEncryption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StopStreamEncryption",
	}
}
