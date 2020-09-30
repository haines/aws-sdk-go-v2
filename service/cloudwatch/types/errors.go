// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// More than one process tried to modify a resource at the same time.
type ConcurrentModificationException struct {
	Message *string
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ConcurrentModificationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ConcurrentModificationException) HasMessage() bool {
	return e.Message != nil
}

// Some part of the dashboard data is invalid.
type DashboardInvalidInputError struct {
	Message *string

	DashboardValidationMessages []*DashboardValidationMessage
}

func (e *DashboardInvalidInputError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DashboardInvalidInputError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DashboardInvalidInputError) ErrorCode() string             { return "DashboardInvalidInputError" }
func (e *DashboardInvalidInputError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DashboardInvalidInputError) GetDashboardValidationMessages() []*DashboardValidationMessage {
	return e.DashboardValidationMessages
}
func (e *DashboardInvalidInputError) HasDashboardValidationMessages() bool {
	return e.DashboardValidationMessages != nil
}
func (e *DashboardInvalidInputError) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DashboardInvalidInputError) HasMessage() bool {
	return e.Message != nil
}

// The specified dashboard does not exist.
type DashboardNotFoundError struct {
	Message *string
}

func (e *DashboardNotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DashboardNotFoundError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DashboardNotFoundError) ErrorCode() string             { return "DashboardNotFoundError" }
func (e *DashboardNotFoundError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DashboardNotFoundError) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DashboardNotFoundError) HasMessage() bool {
	return e.Message != nil
}

// Request processing has failed due to some unknown error, exception, or failure.
type InternalServiceFault struct {
	Message *string
}

func (e *InternalServiceFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceFault) ErrorCode() string             { return "InternalServiceFault" }
func (e *InternalServiceFault) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalServiceFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalServiceFault) HasMessage() bool {
	return e.Message != nil
}

// Data was not syntactically valid JSON.
type InvalidFormatFault struct {
	Message *string
}

func (e *InvalidFormatFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidFormatFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidFormatFault) ErrorCode() string             { return "InvalidFormatFault" }
func (e *InvalidFormatFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidFormatFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidFormatFault) HasMessage() bool {
	return e.Message != nil
}

// The next token specified is invalid.
type InvalidNextToken struct {
	Message *string
}

func (e *InvalidNextToken) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextToken) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextToken) ErrorCode() string             { return "InvalidNextToken" }
func (e *InvalidNextToken) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidNextToken) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidNextToken) HasMessage() bool {
	return e.Message != nil
}

// Parameters were used together that cannot be used together.
type InvalidParameterCombinationException struct {
	Message *string
}

func (e *InvalidParameterCombinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterCombinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterCombinationException) ErrorCode() string {
	return "InvalidParameterCombinationException"
}
func (e *InvalidParameterCombinationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *InvalidParameterCombinationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterCombinationException) HasMessage() bool {
	return e.Message != nil
}

// The value of an input parameter is bad or out-of-range.
type InvalidParameterValueException struct {
	Message *string
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string             { return "InvalidParameterValueException" }
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidParameterValueException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterValueException) HasMessage() bool {
	return e.Message != nil
}

// The operation exceeded one or more limits.
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
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// The quota for alarms for this customer has already been reached.
type LimitExceededFault struct {
	Message *string
}

func (e *LimitExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededFault) ErrorCode() string             { return "LimitExceededFault" }
func (e *LimitExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *LimitExceededFault) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededFault) HasMessage() bool {
	return e.Message != nil
}

// An input parameter that is required is missing.
type MissingRequiredParameterException struct {
	Message *string
}

func (e *MissingRequiredParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingRequiredParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingRequiredParameterException) ErrorCode() string {
	return "MissingRequiredParameterException"
}
func (e *MissingRequiredParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *MissingRequiredParameterException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *MissingRequiredParameterException) HasMessage() bool {
	return e.Message != nil
}

// The named resource does not exist.
type ResourceNotFound struct {
	Message *string
}

func (e *ResourceNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFound) ErrorCode() string             { return "ResourceNotFound" }
func (e *ResourceNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceNotFound) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFound) HasMessage() bool {
	return e.Message != nil
}

// The named resource does not exist.
type ResourceNotFoundException struct {
	Message *string

	ResourceId   *string
	ResourceType *string
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
func (e *ResourceNotFoundException) GetResourceId() string {
	return ptr.ToString(e.ResourceId)
}
func (e *ResourceNotFoundException) HasResourceId() bool {
	return e.ResourceId != nil
}
func (e *ResourceNotFoundException) GetResourceType() string {
	return ptr.ToString(e.ResourceType)
}
func (e *ResourceNotFoundException) HasResourceType() bool {
	return e.ResourceType != nil
}
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}