// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns information about provisioned RDS instances. This API supports
// pagination. This operation can also return information for Amazon Neptune DB
// instances and Amazon DocumentDB instances.
func (c *Client) DescribeDBInstances(ctx context.Context, params *DescribeDBInstancesInput, optFns ...func(*Options)) (*DescribeDBInstancesOutput, error) {
	if params == nil {
		params = &DescribeDBInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBInstances", params, optFns, c.addOperationDescribeDBInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBInstancesInput struct {

	// The user-supplied instance identifier. If this parameter is specified,
	// information from only the specific DB instance is returned. This parameter isn't
	// case-sensitive. Constraints:
	//
	// * If supplied, must match the identifier of an
	// existing DBInstance.
	DBInstanceIdentifier *string

	// A filter that specifies one or more DB instances to describe. Supported
	// filters:
	//
	// * db-cluster-id - Accepts DB cluster identifiers and DB cluster Amazon
	// Resource Names (ARNs). The results list will only include information about the
	// DB instances associated with the DB clusters identified by these ARNs.
	//
	// *
	// db-instance-id - Accepts DB instance identifiers and DB instance Amazon Resource
	// Names (ARNs). The results list will only include information about the DB
	// instances identified by these ARNs.
	//
	// * dbi-resource-id - Accepts DB instance
	// resource identifiers. The results list will only include information about the
	// DB instances identified by these DB instance resource identifiers.
	//
	// * domain -
	// Accepts Active Directory directory IDs. The results list will only include
	// information about the DB instances associated with these domains.
	//
	// * engine -
	// Accepts engine names. The results list will only include information about the
	// DB instances for these engines.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeDBInstances request.
	// If this parameter is specified, the response includes only records beyond the
	// marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

// Contains the result of a successful invocation of the DescribeDBInstances
// action.
type DescribeDBInstancesOutput struct {

	// A list of DBInstance instances.
	DBInstances []types.DBInstance

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeDBInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBInstances{}, middleware.After)
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
	if err = addOpDescribeDBInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBInstances(options.Region), middleware.Before); err != nil {
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

// DescribeDBInstancesAPIClient is a client that implements the DescribeDBInstances
// operation.
type DescribeDBInstancesAPIClient interface {
	DescribeDBInstances(context.Context, *DescribeDBInstancesInput, ...func(*Options)) (*DescribeDBInstancesOutput, error)
}

var _ DescribeDBInstancesAPIClient = (*Client)(nil)

// DescribeDBInstancesPaginatorOptions is the paginator options for
// DescribeDBInstances
type DescribeDBInstancesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBInstancesPaginator is a paginator for DescribeDBInstances
type DescribeDBInstancesPaginator struct {
	options   DescribeDBInstancesPaginatorOptions
	client    DescribeDBInstancesAPIClient
	params    *DescribeDBInstancesInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBInstancesPaginator returns a new DescribeDBInstancesPaginator
func NewDescribeDBInstancesPaginator(client DescribeDBInstancesAPIClient, params *DescribeDBInstancesInput, optFns ...func(*DescribeDBInstancesPaginatorOptions)) *DescribeDBInstancesPaginator {
	if params == nil {
		params = &DescribeDBInstancesInput{}
	}

	options := DescribeDBInstancesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBInstancesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeDBInstances page.
func (p *DescribeDBInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBInstancesOutput, error) {
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

	result, err := p.client.DescribeDBInstances(ctx, &params, optFns...)
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

// DBInstanceAvailableWaiterOptions are waiter options for
// DBInstanceAvailableWaiter
type DBInstanceAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// DBInstanceAvailableWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, DBInstanceAvailableWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeDBInstancesInput, *DescribeDBInstancesOutput, error) (bool, error)
}

// DBInstanceAvailableWaiter defines the waiters for DBInstanceAvailable
type DBInstanceAvailableWaiter struct {
	client DescribeDBInstancesAPIClient

	options DBInstanceAvailableWaiterOptions
}

// NewDBInstanceAvailableWaiter constructs a DBInstanceAvailableWaiter.
func NewDBInstanceAvailableWaiter(client DescribeDBInstancesAPIClient, optFns ...func(*DBInstanceAvailableWaiterOptions)) *DBInstanceAvailableWaiter {
	options := DBInstanceAvailableWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = dBInstanceAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &DBInstanceAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for DBInstanceAvailable waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *DBInstanceAvailableWaiter) Wait(ctx context.Context, params *DescribeDBInstancesInput, maxWaitDur time.Duration, optFns ...func(*DBInstanceAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for DBInstanceAvailable waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *DBInstanceAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeDBInstancesInput, maxWaitDur time.Duration, optFns ...func(*DBInstanceAvailableWaiterOptions)) (*DescribeDBInstancesOutput, error) {
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

		out, err := w.client.DescribeDBInstances(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for DBInstanceAvailable waiter")
}

func dBInstanceAvailableStateRetryable(ctx context.Context, input *DescribeDBInstancesInput, output *DescribeDBInstancesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("DBInstances[].DBInstanceStatus", output)
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
		pathValue, err := jmespath.Search("DBInstances[].DBInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
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
		pathValue, err := jmespath.Search("DBInstances[].DBInstanceStatus", output)
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
		pathValue, err := jmespath.Search("DBInstances[].DBInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "failed"
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
		pathValue, err := jmespath.Search("DBInstances[].DBInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-restore"
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
		pathValue, err := jmespath.Search("DBInstances[].DBInstanceStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-parameters"
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

func newServiceMetadataMiddleware_opDescribeDBInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBInstances",
	}
}
