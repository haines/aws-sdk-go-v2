// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Lists information about a specific audit report created by calling the
// CreateCertificateAuthorityAuditReport
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html)
// action. Audit information is created every time the certificate authority (CA)
// private key is used. The private key is used when you call the IssueCertificate
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_IssueCertificate.html)
// action or the RevokeCertificate
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_RevokeCertificate.html)
// action.
func (c *Client) DescribeCertificateAuthorityAuditReport(ctx context.Context, params *DescribeCertificateAuthorityAuditReportInput, optFns ...func(*Options)) (*DescribeCertificateAuthorityAuditReportOutput, error) {
	if params == nil {
		params = &DescribeCertificateAuthorityAuditReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCertificateAuthorityAuditReport", params, optFns, c.addOperationDescribeCertificateAuthorityAuditReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCertificateAuthorityAuditReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCertificateAuthorityAuditReportInput struct {

	// The report ID returned by calling the CreateCertificateAuthorityAuditReport
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html)
	// action.
	//
	// This member is required.
	AuditReportId *string

	// The Amazon Resource Name (ARN) of the private CA. This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string
}

type DescribeCertificateAuthorityAuditReportOutput struct {

	// Specifies whether report creation is in progress, has succeeded, or has failed.
	AuditReportStatus types.AuditReportStatus

	// The date and time at which the report was created.
	CreatedAt *time.Time

	// Name of the S3 bucket that contains the report.
	S3BucketName *string

	// S3 key that uniquely identifies the report file in your S3 bucket.
	S3Key *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeCertificateAuthorityAuditReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCertificateAuthorityAuditReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCertificateAuthorityAuditReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeCertificateAuthorityAuditReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCertificateAuthorityAuditReport(options.Region), middleware.Before); err != nil {
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
	return nil
}

// DescribeCertificateAuthorityAuditReportAPIClient is a client that implements the
// DescribeCertificateAuthorityAuditReport operation.
type DescribeCertificateAuthorityAuditReportAPIClient interface {
	DescribeCertificateAuthorityAuditReport(context.Context, *DescribeCertificateAuthorityAuditReportInput, ...func(*Options)) (*DescribeCertificateAuthorityAuditReportOutput, error)
}

var _ DescribeCertificateAuthorityAuditReportAPIClient = (*Client)(nil)

// AuditReportCreatedWaiterOptions are waiter options for AuditReportCreatedWaiter
type AuditReportCreatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// AuditReportCreatedWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, AuditReportCreatedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeCertificateAuthorityAuditReportInput, *DescribeCertificateAuthorityAuditReportOutput, error) (bool, error)
}

// AuditReportCreatedWaiter defines the waiters for AuditReportCreated
type AuditReportCreatedWaiter struct {
	client DescribeCertificateAuthorityAuditReportAPIClient

	options AuditReportCreatedWaiterOptions
}

// NewAuditReportCreatedWaiter constructs a AuditReportCreatedWaiter.
func NewAuditReportCreatedWaiter(client DescribeCertificateAuthorityAuditReportAPIClient, optFns ...func(*AuditReportCreatedWaiterOptions)) *AuditReportCreatedWaiter {
	options := AuditReportCreatedWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = auditReportCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AuditReportCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AuditReportCreated waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *AuditReportCreatedWaiter) Wait(ctx context.Context, params *DescribeCertificateAuthorityAuditReportInput, maxWaitDur time.Duration, optFns ...func(*AuditReportCreatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for AuditReportCreated waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *AuditReportCreatedWaiter) WaitForOutput(ctx context.Context, params *DescribeCertificateAuthorityAuditReportInput, maxWaitDur time.Duration, optFns ...func(*AuditReportCreatedWaiterOptions)) (*DescribeCertificateAuthorityAuditReportOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeCertificateAuthorityAuditReport(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for AuditReportCreated waiter")
}

func auditReportCreatedStateRetryable(ctx context.Context, input *DescribeCertificateAuthorityAuditReportInput, output *DescribeCertificateAuthorityAuditReportOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("AuditReportStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "SUCCESS"
		value, ok := pathValue.(types.AuditReportStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.AuditReportStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("AuditReportStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.AuditReportStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.AuditReportStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeCertificateAuthorityAuditReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "DescribeCertificateAuthorityAuditReport",
	}
}
