// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns information about replication instances for your account in the current
// region.
func (c *Client) DescribeReplicationInstances(ctx context.Context, params *DescribeReplicationInstancesInput, optFns ...func(*Options)) (*DescribeReplicationInstancesOutput, error) {
	if params == nil {
		params = &DescribeReplicationInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationInstances", params, optFns, c.addOperationDescribeReplicationInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeReplicationInstancesInput struct {

	// Filters applied to replication instances. Valid filter names:
	// replication-instance-arn | replication-instance-id | replication-instance-class
	// | engine-version
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

//
type DescribeReplicationInstancesOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The replication instances described.
	ReplicationInstances []types.ReplicationInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeReplicationInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReplicationInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReplicationInstances{}, middleware.After)
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
	if err = addOpDescribeReplicationInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationInstances(options.Region), middleware.Before); err != nil {
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

// DescribeReplicationInstancesAPIClient is a client that implements the
// DescribeReplicationInstances operation.
type DescribeReplicationInstancesAPIClient interface {
	DescribeReplicationInstances(context.Context, *DescribeReplicationInstancesInput, ...func(*Options)) (*DescribeReplicationInstancesOutput, error)
}

var _ DescribeReplicationInstancesAPIClient = (*Client)(nil)

// DescribeReplicationInstancesPaginatorOptions is the paginator options for
// DescribeReplicationInstances
type DescribeReplicationInstancesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReplicationInstancesPaginator is a paginator for
// DescribeReplicationInstances
type DescribeReplicationInstancesPaginator struct {
	options   DescribeReplicationInstancesPaginatorOptions
	client    DescribeReplicationInstancesAPIClient
	params    *DescribeReplicationInstancesInput
	nextToken *string
	firstPage bool
}

// NewDescribeReplicationInstancesPaginator returns a new
// DescribeReplicationInstancesPaginator
func NewDescribeReplicationInstancesPaginator(client DescribeReplicationInstancesAPIClient, params *DescribeReplicationInstancesInput, optFns ...func(*DescribeReplicationInstancesPaginatorOptions)) *DescribeReplicationInstancesPaginator {
	if params == nil {
		params = &DescribeReplicationInstancesInput{}
	}

	options := DescribeReplicationInstancesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReplicationInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReplicationInstancesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeReplicationInstances page.
func (p *DescribeReplicationInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReplicationInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeReplicationInstances(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ReplicationInstanceAvailableWaiterOptions are waiter options for
// ReplicationInstanceAvailableWaiter
type ReplicationInstanceAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ReplicationInstanceAvailableWaiter will use default minimum delay of 60 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ReplicationInstanceAvailableWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *DescribeReplicationInstancesInput, *DescribeReplicationInstancesOutput, error) (bool, error)
}

// ReplicationInstanceAvailableWaiter defines the waiters for
// ReplicationInstanceAvailable
type ReplicationInstanceAvailableWaiter struct {
	client DescribeReplicationInstancesAPIClient

	options ReplicationInstanceAvailableWaiterOptions
}

// NewReplicationInstanceAvailableWaiter constructs a
// ReplicationInstanceAvailableWaiter.
func NewReplicationInstanceAvailableWaiter(client DescribeReplicationInstancesAPIClient, optFns ...func(*ReplicationInstanceAvailableWaiterOptions)) *ReplicationInstanceAvailableWaiter {
	options := ReplicationInstanceAvailableWaiterOptions{}
	options.MinDelay = 60 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = replicationInstanceAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ReplicationInstanceAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ReplicationInstanceAvailable waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ReplicationInstanceAvailableWaiter) Wait(ctx context.Context, params *DescribeReplicationInstancesInput, maxWaitDur time.Duration, optFns ...func(*ReplicationInstanceAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ReplicationInstanceAvailable waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ReplicationInstanceAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeReplicationInstancesInput, maxWaitDur time.Duration, optFns ...func(*ReplicationInstanceAvailableWaiterOptions)) (*DescribeReplicationInstancesOutput, error) {
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

		out, err := w.client.DescribeReplicationInstances(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ReplicationInstanceAvailable waiter")
}

func replicationInstanceAvailableStateRetryable(ctx context.Context, input *DescribeReplicationInstancesInput, output *DescribeReplicationInstancesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("ReplicationInstances[].ReplicationInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ReplicationInstances[].ReplicationInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleting"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ReplicationInstances[].ReplicationInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-credentials"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ReplicationInstances[].ReplicationInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-network"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ReplicationInstances[].ReplicationInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "inaccessible-encryption-credentials"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

// ReplicationInstanceDeletedWaiterOptions are waiter options for
// ReplicationInstanceDeletedWaiter
type ReplicationInstanceDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ReplicationInstanceDeletedWaiter will use default minimum delay of 15 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ReplicationInstanceDeletedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *DescribeReplicationInstancesInput, *DescribeReplicationInstancesOutput, error) (bool, error)
}

// ReplicationInstanceDeletedWaiter defines the waiters for
// ReplicationInstanceDeleted
type ReplicationInstanceDeletedWaiter struct {
	client DescribeReplicationInstancesAPIClient

	options ReplicationInstanceDeletedWaiterOptions
}

// NewReplicationInstanceDeletedWaiter constructs a
// ReplicationInstanceDeletedWaiter.
func NewReplicationInstanceDeletedWaiter(client DescribeReplicationInstancesAPIClient, optFns ...func(*ReplicationInstanceDeletedWaiterOptions)) *ReplicationInstanceDeletedWaiter {
	options := ReplicationInstanceDeletedWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = replicationInstanceDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ReplicationInstanceDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ReplicationInstanceDeleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ReplicationInstanceDeletedWaiter) Wait(ctx context.Context, params *DescribeReplicationInstancesInput, maxWaitDur time.Duration, optFns ...func(*ReplicationInstanceDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ReplicationInstanceDeleted waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ReplicationInstanceDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeReplicationInstancesInput, maxWaitDur time.Duration, optFns ...func(*ReplicationInstanceDeletedWaiterOptions)) (*DescribeReplicationInstancesOutput, error) {
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

		out, err := w.client.DescribeReplicationInstances(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ReplicationInstanceDeleted waiter")
}

func replicationInstanceDeletedStateRetryable(ctx context.Context, input *DescribeReplicationInstancesInput, output *DescribeReplicationInstancesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("ReplicationInstances[].ReplicationInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err != nil {
		var errorType *types.ResourceNotFoundFault
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeReplicationInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeReplicationInstances",
	}
}
