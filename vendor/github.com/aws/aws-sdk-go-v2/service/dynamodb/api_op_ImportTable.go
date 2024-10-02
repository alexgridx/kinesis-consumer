// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports table data from an S3 bucket.
func (c *Client) ImportTable(ctx context.Context, params *ImportTableInput, optFns ...func(*Options)) (*ImportTableOutput, error) {
	if params == nil {
		params = &ImportTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportTable", params, optFns, c.addOperationImportTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportTableInput struct {

	//  The format of the source data. Valid values for ImportFormat are CSV ,
	// DYNAMODB_JSON or ION .
	//
	// This member is required.
	InputFormat types.InputFormat

	//  The S3 bucket that provides the source for the import.
	//
	// This member is required.
	S3BucketSource *types.S3BucketSource

	// Parameters for the table to import the data into.
	//
	// This member is required.
	TableCreationParameters *types.TableCreationParameters

	// Providing a ClientToken makes the call to ImportTableInput idempotent, meaning
	// that multiple identical calls have the same effect as one single call.
	//
	// A client token is valid for 8 hours after the first request that uses it is
	// completed. After 8 hours, any request with the same client token is treated as a
	// new request. Do not resubmit the same request with the same client token for
	// more than 8 hours, or the result might not be idempotent.
	//
	// If you submit a request with the same client token but a change in other
	// parameters within the 8-hour idempotency window, DynamoDB returns an
	// IdempotentParameterMismatch exception.
	ClientToken *string

	//  Type of compression to be used on the input coming from the imported table.
	InputCompressionType types.InputCompressionType

	//  Additional properties that specify how the input is formatted,
	InputFormatOptions *types.InputFormatOptions

	noSmithyDocumentSerde
}

type ImportTableOutput struct {

	//  Represents the properties of the table created for the import, and parameters
	// of the import. The import parameters include import status, how many items were
	// processed, and how many errors were encountered.
	//
	// This member is required.
	ImportTableDescription *types.ImportTableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpImportTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpImportTable{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportTable"); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opImportTableMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpImportTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportTable(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpImportTable struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpImportTable) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpImportTable) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ImportTableInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ImportTableInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opImportTableMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpImportTable{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opImportTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportTable",
	}
}
