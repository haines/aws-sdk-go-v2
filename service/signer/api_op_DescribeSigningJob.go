// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns information about a specific code signing job. You specify the job by
// using the jobId value that is returned by the StartSigningJob operation.
func (c *Client) DescribeSigningJob(ctx context.Context, params *DescribeSigningJobInput, optFns ...func(*Options)) (*DescribeSigningJobOutput, error) {
	if params == nil {
		params = &DescribeSigningJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSigningJob", params, optFns, c.addOperationDescribeSigningJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSigningJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSigningJobInput struct {

	// The ID of the signing job on input.
	//
	// This member is required.
	JobId *string
}

type DescribeSigningJobOutput struct {

	// Date and time that the signing job was completed.
	CompletedAt *time.Time

	// Date and time that the signing job was created.
	CreatedAt *time.Time

	// The ID of the signing job on output.
	JobId *string

	// The IAM entity that initiated the signing job.
	JobInvoker *string

	// The AWS account ID of the job owner.
	JobOwner *string

	// A list of any overrides that were applied to the signing operation.
	Overrides *types.SigningPlatformOverrides

	// A human-readable name for the signing platform associated with the signing job.
	PlatformDisplayName *string

	// The microcontroller platform to which your signed code image will be
	// distributed.
	PlatformId *string

	// The name of the profile that initiated the signing operation.
	ProfileName *string

	// The version of the signing profile used to initiate the signing job.
	ProfileVersion *string

	// The IAM principal that requested the signing job.
	RequestedBy *string

	// A revocation record if the signature generated by the signing job has been
	// revoked. Contains a timestamp and the ID of the IAM entity that revoked the
	// signature.
	RevocationRecord *types.SigningJobRevocationRecord

	// Thr expiration timestamp for the signature generated by the signing job.
	SignatureExpiresAt *time.Time

	// Name of the S3 bucket where the signed code image is saved by code signing.
	SignedObject *types.SignedObject

	// The Amazon Resource Name (ARN) of your code signing certificate.
	SigningMaterial *types.SigningMaterial

	// Map of user-assigned key-value pairs used during signing. These values contain
	// any information that you specified for use in your signing job.
	SigningParameters map[string]string

	// The object that contains the name of your S3 bucket or your raw code.
	Source *types.Source

	// Status of the signing job.
	Status types.SigningStatus

	// String value that contains the status reason.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeSigningJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSigningJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSigningJob{}, middleware.After)
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
	if err = addOpDescribeSigningJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSigningJob(options.Region), middleware.Before); err != nil {
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

// DescribeSigningJobAPIClient is a client that implements the DescribeSigningJob
// operation.
type DescribeSigningJobAPIClient interface {
	DescribeSigningJob(context.Context, *DescribeSigningJobInput, ...func(*Options)) (*DescribeSigningJobOutput, error)
}

var _ DescribeSigningJobAPIClient = (*Client)(nil)

// SuccessfulSigningJobWaiterOptions are waiter options for
// SuccessfulSigningJobWaiter
type SuccessfulSigningJobWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// SuccessfulSigningJobWaiter will use default minimum delay of 20 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, SuccessfulSigningJobWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeSigningJobInput, *DescribeSigningJobOutput, error) (bool, error)
}

// SuccessfulSigningJobWaiter defines the waiters for SuccessfulSigningJob
type SuccessfulSigningJobWaiter struct {
	client DescribeSigningJobAPIClient

	options SuccessfulSigningJobWaiterOptions
}

// NewSuccessfulSigningJobWaiter constructs a SuccessfulSigningJobWaiter.
func NewSuccessfulSigningJobWaiter(client DescribeSigningJobAPIClient, optFns ...func(*SuccessfulSigningJobWaiterOptions)) *SuccessfulSigningJobWaiter {
	options := SuccessfulSigningJobWaiterOptions{}
	options.MinDelay = 20 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = successfulSigningJobStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &SuccessfulSigningJobWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for SuccessfulSigningJob waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *SuccessfulSigningJobWaiter) Wait(ctx context.Context, params *DescribeSigningJobInput, maxWaitDur time.Duration, optFns ...func(*SuccessfulSigningJobWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for SuccessfulSigningJob waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *SuccessfulSigningJobWaiter) WaitForOutput(ctx context.Context, params *DescribeSigningJobInput, maxWaitDur time.Duration, optFns ...func(*SuccessfulSigningJobWaiterOptions)) (*DescribeSigningJobOutput, error) {
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

		out, err := w.client.DescribeSigningJob(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for SuccessfulSigningJob waiter")
}

func successfulSigningJobStateRetryable(ctx context.Context, input *DescribeSigningJobInput, output *DescribeSigningJobOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Succeeded"
		value, ok := pathValue.(types.SigningStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.SigningStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.SigningStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.SigningStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeSigningJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "signer",
		OperationName: "DescribeSigningJob",
	}
}
