// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The specified load balancer does not exist.
type AccessPointNotFoundException struct {
	Message *string
}

func (e *AccessPointNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessPointNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessPointNotFoundException) ErrorCode() string             { return "LoadBalancerNotFound" }
func (e *AccessPointNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified ARN does not refer to a valid SSL certificate in AWS Identity and
// Access Management (IAM) or AWS Certificate Manager (ACM). Note that if you
// recently uploaded the certificate to IAM, this error might indicate that the
// certificate is not fully available yet.
type CertificateNotFoundException struct {
	Message *string
}

func (e *CertificateNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CertificateNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CertificateNotFoundException) ErrorCode() string             { return "CertificateNotFound" }
func (e *CertificateNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A request made by Elastic Load Balancing to another service exceeds the maximum
// request rate permitted for your account.
type DependencyThrottleException struct {
	Message *string
}

func (e *DependencyThrottleException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DependencyThrottleException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DependencyThrottleException) ErrorCode() string             { return "DependencyThrottle" }
func (e *DependencyThrottleException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified load balancer name already exists for this account.
type DuplicateAccessPointNameException struct {
	Message *string
}

func (e *DuplicateAccessPointNameException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateAccessPointNameException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateAccessPointNameException) ErrorCode() string             { return "DuplicateLoadBalancerName" }
func (e *DuplicateAccessPointNameException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A listener already exists for the specified load balancer name and port, but
// with a different instance port, protocol, or SSL certificate.
type DuplicateListenerException struct {
	Message *string
}

func (e *DuplicateListenerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateListenerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateListenerException) ErrorCode() string             { return "DuplicateListener" }
func (e *DuplicateListenerException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A policy with the specified name already exists for this load balancer.
type DuplicatePolicyNameException struct {
	Message *string
}

func (e *DuplicatePolicyNameException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicatePolicyNameException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicatePolicyNameException) ErrorCode() string             { return "DuplicatePolicyName" }
func (e *DuplicatePolicyNameException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A tag key was specified more than once.
type DuplicateTagKeysException struct {
	Message *string
}

func (e *DuplicateTagKeysException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateTagKeysException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateTagKeysException) ErrorCode() string             { return "DuplicateTagKeys" }
func (e *DuplicateTagKeysException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested configuration change is not valid.
type InvalidConfigurationRequestException struct {
	Message *string
}

func (e *InvalidConfigurationRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidConfigurationRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidConfigurationRequestException) ErrorCode() string {
	return "InvalidConfigurationRequest"
}
func (e *InvalidConfigurationRequestException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified endpoint is not valid.
type InvalidEndPointException struct {
	Message *string
}

func (e *InvalidEndPointException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidEndPointException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidEndPointException) ErrorCode() string             { return "InvalidInstance" }
func (e *InvalidEndPointException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified value for the schema is not valid. You can only specify a scheme
// for load balancers in a VPC.
type InvalidSchemeException struct {
	Message *string
}

func (e *InvalidSchemeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSchemeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSchemeException) ErrorCode() string             { return "InvalidScheme" }
func (e *InvalidSchemeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more of the specified security groups do not exist.
type InvalidSecurityGroupException struct {
	Message *string
}

func (e *InvalidSecurityGroupException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSecurityGroupException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSecurityGroupException) ErrorCode() string             { return "InvalidSecurityGroup" }
func (e *InvalidSecurityGroupException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified VPC has no associated Internet gateway.
type InvalidSubnetException struct {
	Message *string
}

func (e *InvalidSubnetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSubnetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSubnetException) ErrorCode() string             { return "InvalidSubnet" }
func (e *InvalidSubnetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The load balancer does not have a listener configured at the specified port.
type ListenerNotFoundException struct {
	Message *string
}

func (e *ListenerNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ListenerNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ListenerNotFoundException) ErrorCode() string             { return "ListenerNotFound" }
func (e *ListenerNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified load balancer attribute does not exist.
type LoadBalancerAttributeNotFoundException struct {
	Message *string
}

func (e *LoadBalancerAttributeNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LoadBalancerAttributeNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LoadBalancerAttributeNotFoundException) ErrorCode() string {
	return "LoadBalancerAttributeNotFound"
}
func (e *LoadBalancerAttributeNotFoundException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// This operation is not allowed.
type OperationNotPermittedException struct {
	Message *string
}

func (e *OperationNotPermittedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationNotPermittedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationNotPermittedException) ErrorCode() string             { return "OperationNotPermitted" }
func (e *OperationNotPermittedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more of the specified policies do not exist.
type PolicyNotFoundException struct {
	Message *string
}

func (e *PolicyNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PolicyNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PolicyNotFoundException) ErrorCode() string             { return "PolicyNotFound" }
func (e *PolicyNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more of the specified policy types do not exist.
type PolicyTypeNotFoundException struct {
	Message *string
}

func (e *PolicyTypeNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PolicyTypeNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PolicyTypeNotFoundException) ErrorCode() string             { return "PolicyTypeNotFound" }
func (e *PolicyTypeNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more of the specified subnets do not exist.
type SubnetNotFoundException struct {
	Message *string
}

func (e *SubnetNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetNotFoundException) ErrorCode() string             { return "SubnetNotFound" }
func (e *SubnetNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The quota for the number of load balancers has been reached.
type TooManyAccessPointsException struct {
	Message *string
}

func (e *TooManyAccessPointsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyAccessPointsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyAccessPointsException) ErrorCode() string             { return "TooManyLoadBalancers" }
func (e *TooManyAccessPointsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The quota for the number of policies for this load balancer has been reached.
type TooManyPoliciesException struct {
	Message *string
}

func (e *TooManyPoliciesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyPoliciesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyPoliciesException) ErrorCode() string             { return "TooManyPolicies" }
func (e *TooManyPoliciesException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The quota for the number of tags that can be assigned to a load balancer has
// been reached.
type TooManyTagsException struct {
	Message *string
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string             { return "TooManyTags" }
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified protocol or signature version is not supported.
type UnsupportedProtocolException struct {
	Message *string
}

func (e *UnsupportedProtocolException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedProtocolException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedProtocolException) ErrorCode() string             { return "UnsupportedProtocol" }
func (e *UnsupportedProtocolException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
