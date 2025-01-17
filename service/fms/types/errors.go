// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The operation failed because of a system problem, even though the request was
// valid. Retry your request.
type InternalErrorException struct {
	Message *string
}

func (e *InternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalErrorException) ErrorCode() string             { return "InternalErrorException" }
func (e *InternalErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The parameters of the request were invalid.
type InvalidInputException struct {
	Message *string
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string             { return "InvalidInputException" }
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation failed because there was nothing to do or the operation wasn't
// possible. For example, you might have submitted an AssociateAdminAccount request
// for an account ID that was already set as the Firewall Manager administrator. Or
// you might have tried to access a Region that's disabled by default, and that you
// need to enable for the Firewall Manager administrator account and for
// Organizations before you can access it.
type InvalidOperationException struct {
	Message *string
}

func (e *InvalidOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidOperationException) ErrorCode() string             { return "InvalidOperationException" }
func (e *InvalidOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The value of the Type parameter is invalid.
type InvalidTypeException struct {
	Message *string
}

func (e *InvalidTypeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTypeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTypeException) ErrorCode() string             { return "InvalidTypeException" }
func (e *InvalidTypeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation exceeds a resource limit, for example, the maximum number of
// policy objects that you can create for an Amazon Web Services account. For more
// information, see Firewall Manager Limits
// (https://docs.aws.amazon.com/waf/latest/developerguide/fms-limits.html) in the
// WAF Developer Guide.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource was not found.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
