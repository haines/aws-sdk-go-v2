// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Amplify App represents different branches of a repository for building, deploying,
// and hosting.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/App
type App struct {
	_ struct{} `type:"structure"`

	// ARN for the Amplify App.
	//
	// AppArn is a required field
	AppArn *string `locationName:"appArn" type:"string" required:"true"`

	// Unique Id for the Amplify App.
	//
	// AppId is a required field
	AppId *string `locationName:"appId" min:"1" type:"string" required:"true"`

	// Basic Authorization credentials for branches for the Amplify App.
	BasicAuthCredentials *string `locationName:"basicAuthCredentials" type:"string"`

	// BuildSpec content for Amplify App.
	BuildSpec *string `locationName:"buildSpec" min:"1" type:"string"`

	// Create date / time for the Amplify App.
	//
	// CreateTime is a required field
	CreateTime *time.Time `locationName:"createTime" type:"timestamp" timestampFormat:"unix" required:"true"`

	// Custom redirect / rewrite rules for the Amplify App.
	CustomRules []CustomRule `locationName:"customRules" type:"list"`

	// Default domain for the Amplify App.
	//
	// DefaultDomain is a required field
	DefaultDomain *string `locationName:"defaultDomain" min:"1" type:"string" required:"true"`

	// Description for the Amplify App.
	//
	// Description is a required field
	Description *string `locationName:"description" type:"string" required:"true"`

	// Enables Basic Authorization for branches for the Amplify App.
	//
	// EnableBasicAuth is a required field
	EnableBasicAuth *bool `locationName:"enableBasicAuth" type:"boolean" required:"true"`

	// Enables auto-building of branches for the Amplify App.
	//
	// EnableBranchAutoBuild is a required field
	EnableBranchAutoBuild *bool `locationName:"enableBranchAutoBuild" type:"boolean" required:"true"`

	// Environment Variables for the Amplify App.
	//
	// EnvironmentVariables is a required field
	EnvironmentVariables map[string]string `locationName:"environmentVariables" type:"map" required:"true"`

	// IAM service role ARN for the Amplify App.
	IamServiceRoleArn *string `locationName:"iamServiceRoleArn" min:"1" type:"string"`

	// Name for the Amplify App.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// Platform for the Amplify App.
	//
	// Platform is a required field
	Platform Platform `locationName:"platform" type:"string" required:"true" enum:"true"`

	// Structure with Production Branch information.
	ProductionBranch *ProductionBranch `locationName:"productionBranch" type:"structure"`

	// Repository for the Amplify App.
	//
	// Repository is a required field
	Repository *string `locationName:"repository" type:"string" required:"true"`

	// Tag for Amplify App.
	Tags map[string]string `locationName:"tags" type:"map"`

	// Update date / time for the Amplify App.
	//
	// UpdateTime is a required field
	UpdateTime *time.Time `locationName:"updateTime" type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s App) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s App) MarshalFields(e protocol.FieldEncoder) error {
	if s.AppArn != nil {
		v := *s.AppArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "appArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BasicAuthCredentials != nil {
		v := *s.BasicAuthCredentials

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "basicAuthCredentials", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BuildSpec != nil {
		v := *s.BuildSpec

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "buildSpec", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreateTime != nil {
		v := *s.CreateTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.CustomRules != nil {
		v := s.CustomRules

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "customRules", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DefaultDomain != nil {
		v := *s.DefaultDomain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultDomain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnableBasicAuth != nil {
		v := *s.EnableBasicAuth

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableBasicAuth", protocol.BoolValue(v), metadata)
	}
	if s.EnableBranchAutoBuild != nil {
		v := *s.EnableBranchAutoBuild

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableBranchAutoBuild", protocol.BoolValue(v), metadata)
	}
	if s.EnvironmentVariables != nil {
		v := s.EnvironmentVariables

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "environmentVariables", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.IamServiceRoleArn != nil {
		v := *s.IamServiceRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "iamServiceRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Platform) > 0 {
		v := s.Platform

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "platform", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ProductionBranch != nil {
		v := s.ProductionBranch

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "productionBranch", v, metadata)
	}
	if s.Repository != nil {
		v := *s.Repository

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "repository", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.UpdateTime != nil {
		v := *s.UpdateTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "updateTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	return nil
}

// Branch for an Amplify App, which maps to a 3rd party repository branch.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/Branch
type Branch struct {
	_ struct{} `type:"structure"`

	// Id of the active job for a branch, part of an Amplify App.
	//
	// ActiveJobId is a required field
	ActiveJobId *string `locationName:"activeJobId" type:"string" required:"true"`

	// Basic Authorization credentials for a branch, part of an Amplify App.
	BasicAuthCredentials *string `locationName:"basicAuthCredentials" type:"string"`

	// ARN for a branch, part of an Amplify App.
	//
	// BranchArn is a required field
	BranchArn *string `locationName:"branchArn" type:"string" required:"true"`

	// Name for a branch, part of an Amplify App.
	//
	// BranchName is a required field
	BranchName *string `locationName:"branchName" min:"1" type:"string" required:"true"`

	// BuildSpec content for branch for Amplify App.
	BuildSpec *string `locationName:"buildSpec" min:"1" type:"string"`

	// Creation date and time for a branch, part of an Amplify App.
	//
	// CreateTime is a required field
	CreateTime *time.Time `locationName:"createTime" type:"timestamp" timestampFormat:"unix" required:"true"`

	// Custom domains for a branch, part of an Amplify App.
	//
	// CustomDomains is a required field
	CustomDomains []string `locationName:"customDomains" type:"list" required:"true"`

	// Description for a branch, part of an Amplify App.
	//
	// Description is a required field
	Description *string `locationName:"description" type:"string" required:"true"`

	// Display name for a branch, part of an Amplify App.
	DisplayName *string `locationName:"displayName" type:"string"`

	// Enables auto-building on push for a branch, part of an Amplify App.
	//
	// EnableAutoBuild is a required field
	EnableAutoBuild *bool `locationName:"enableAutoBuild" type:"boolean" required:"true"`

	// Enables Basic Authorization for a branch, part of an Amplify App.
	//
	// EnableBasicAuth is a required field
	EnableBasicAuth *bool `locationName:"enableBasicAuth" type:"boolean" required:"true"`

	// Enables notifications for a branch, part of an Amplify App.
	//
	// EnableNotification is a required field
	EnableNotification *bool `locationName:"enableNotification" type:"boolean" required:"true"`

	// Environment Variables specific to a branch, part of an Amplify App.
	//
	// EnvironmentVariables is a required field
	EnvironmentVariables map[string]string `locationName:"environmentVariables" type:"map" required:"true"`

	// Framework for a branch, part of an Amplify App.
	//
	// Framework is a required field
	Framework *string `locationName:"framework" type:"string" required:"true"`

	// Stage for a branch, part of an Amplify App.
	//
	// Stage is a required field
	Stage Stage `locationName:"stage" type:"string" required:"true" enum:"true"`

	// Tag for branch for Amplify App.
	Tags map[string]string `locationName:"tags" type:"map"`

	// Thumbnail Url for the branch.
	ThumbnailUrl *string `locationName:"thumbnailUrl" min:"1" type:"string"`

	// Total number of Jobs part of an Amplify App.
	//
	// TotalNumberOfJobs is a required field
	TotalNumberOfJobs *string `locationName:"totalNumberOfJobs" type:"string" required:"true"`

	// The content TTL for the website in seconds.
	//
	// Ttl is a required field
	Ttl *string `locationName:"ttl" type:"string" required:"true"`

	// Last updated date and time for a branch, part of an Amplify App.
	//
	// UpdateTime is a required field
	UpdateTime *time.Time `locationName:"updateTime" type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s Branch) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Branch) MarshalFields(e protocol.FieldEncoder) error {
	if s.ActiveJobId != nil {
		v := *s.ActiveJobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "activeJobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BasicAuthCredentials != nil {
		v := *s.BasicAuthCredentials

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "basicAuthCredentials", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BranchArn != nil {
		v := *s.BranchArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "branchArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BranchName != nil {
		v := *s.BranchName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "branchName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BuildSpec != nil {
		v := *s.BuildSpec

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "buildSpec", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreateTime != nil {
		v := *s.CreateTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.CustomDomains != nil {
		v := s.CustomDomains

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "customDomains", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DisplayName != nil {
		v := *s.DisplayName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "displayName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnableAutoBuild != nil {
		v := *s.EnableAutoBuild

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableAutoBuild", protocol.BoolValue(v), metadata)
	}
	if s.EnableBasicAuth != nil {
		v := *s.EnableBasicAuth

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableBasicAuth", protocol.BoolValue(v), metadata)
	}
	if s.EnableNotification != nil {
		v := *s.EnableNotification

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableNotification", protocol.BoolValue(v), metadata)
	}
	if s.EnvironmentVariables != nil {
		v := s.EnvironmentVariables

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "environmentVariables", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Framework != nil {
		v := *s.Framework

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "framework", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Stage) > 0 {
		v := s.Stage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stage", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ThumbnailUrl != nil {
		v := *s.ThumbnailUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thumbnailUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TotalNumberOfJobs != nil {
		v := *s.TotalNumberOfJobs

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "totalNumberOfJobs", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Ttl != nil {
		v := *s.Ttl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ttl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UpdateTime != nil {
		v := *s.UpdateTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "updateTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	return nil
}

// Custom rewrite / redirect rule.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/CustomRule
type CustomRule struct {
	_ struct{} `type:"structure"`

	// The condition for a URL rewrite or redirect rule, e.g. country code.
	Condition *string `locationName:"condition" min:"1" type:"string"`

	// The source pattern for a URL rewrite or redirect rule.
	//
	// Source is a required field
	Source *string `locationName:"source" min:"1" type:"string" required:"true"`

	// The status code for a URL rewrite or redirect rule.
	Status *string `locationName:"status" min:"3" type:"string"`

	// The target pattern for a URL rewrite or redirect rule.
	//
	// Target is a required field
	Target *string `locationName:"target" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CustomRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CustomRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CustomRule"}
	if s.Condition != nil && len(*s.Condition) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Condition", 1))
	}

	if s.Source == nil {
		invalidParams.Add(aws.NewErrParamRequired("Source"))
	}
	if s.Source != nil && len(*s.Source) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Source", 1))
	}
	if s.Status != nil && len(*s.Status) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Status", 3))
	}

	if s.Target == nil {
		invalidParams.Add(aws.NewErrParamRequired("Target"))
	}
	if s.Target != nil && len(*s.Target) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Target", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CustomRule) MarshalFields(e protocol.FieldEncoder) error {
	if s.Condition != nil {
		v := *s.Condition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "condition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Source != nil {
		v := *s.Source

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "source", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Status != nil {
		v := *s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Target != nil {
		v := *s.Target

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "target", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Structure for Domain Association, which associates a custom domain with an
// Amplify App.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DomainAssociation
type DomainAssociation struct {
	_ struct{} `type:"structure"`

	// DNS Record for certificate verification.
	//
	// CertificateVerificationDNSRecord is a required field
	CertificateVerificationDNSRecord *string `locationName:"certificateVerificationDNSRecord" type:"string" required:"true"`

	// ARN for the Domain Association.
	//
	// DomainAssociationArn is a required field
	DomainAssociationArn *string `locationName:"domainAssociationArn" type:"string" required:"true"`

	// Name of the domain.
	//
	// DomainName is a required field
	DomainName *string `locationName:"domainName" type:"string" required:"true"`

	// Status fo the Domain Association.
	//
	// DomainStatus is a required field
	DomainStatus DomainStatus `locationName:"domainStatus" type:"string" required:"true" enum:"true"`

	// Enables automated creation of Subdomains for branches.
	//
	// EnableAutoSubDomain is a required field
	EnableAutoSubDomain *bool `locationName:"enableAutoSubDomain" type:"boolean" required:"true"`

	// Reason for the current status of the Domain Association.
	//
	// StatusReason is a required field
	StatusReason *string `locationName:"statusReason" type:"string" required:"true"`

	// Subdomains for the Domain Association.
	//
	// SubDomains is a required field
	SubDomains []SubDomain `locationName:"subDomains" type:"list" required:"true"`
}

// String returns the string representation
func (s DomainAssociation) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DomainAssociation) MarshalFields(e protocol.FieldEncoder) error {
	if s.CertificateVerificationDNSRecord != nil {
		v := *s.CertificateVerificationDNSRecord

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateVerificationDNSRecord", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainAssociationArn != nil {
		v := *s.DomainAssociationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainAssociationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DomainStatus) > 0 {
		v := s.DomainStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.EnableAutoSubDomain != nil {
		v := *s.EnableAutoSubDomain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableAutoSubDomain", protocol.BoolValue(v), metadata)
	}
	if s.StatusReason != nil {
		v := *s.StatusReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SubDomains != nil {
		v := s.SubDomains

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "subDomains", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Structure for an execution job for an Amplify App.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/Job
type Job struct {
	_ struct{} `type:"structure"`

	// Execution steps for an execution job, for an Amplify App.
	//
	// Steps is a required field
	Steps []Step `locationName:"steps" type:"list" required:"true"`

	// Summary for an execution job for an Amplify App.
	//
	// Summary is a required field
	Summary *JobSummary `locationName:"summary" type:"structure" required:"true"`
}

// String returns the string representation
func (s Job) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Job) MarshalFields(e protocol.FieldEncoder) error {
	if s.Steps != nil {
		v := s.Steps

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "steps", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Summary != nil {
		v := s.Summary

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "summary", v, metadata)
	}
	return nil
}

// Structure for the summary of a Job.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/JobSummary
type JobSummary struct {
	_ struct{} `type:"structure"`

	// Commit Id from 3rd party repository provider for the Job.
	//
	// CommitId is a required field
	CommitId *string `locationName:"commitId" type:"string" required:"true"`

	// Commit message from 3rd party repository provider for the Job.
	//
	// CommitMessage is a required field
	CommitMessage *string `locationName:"commitMessage" type:"string" required:"true"`

	// Commit date / time for the Job.
	//
	// CommitTime is a required field
	CommitTime *time.Time `locationName:"commitTime" type:"timestamp" timestampFormat:"unix" required:"true"`

	// End date / time for the Job.
	EndTime *time.Time `locationName:"endTime" type:"timestamp" timestampFormat:"unix"`

	// Arn for the Job.
	//
	// JobArn is a required field
	JobArn *string `locationName:"jobArn" type:"string" required:"true"`

	// Unique Id for the Job.
	//
	// JobId is a required field
	JobId *string `locationName:"jobId" type:"string" required:"true"`

	// Type for the Job.
	//
	// JobType is a required field
	JobType JobType `locationName:"jobType" type:"string" required:"true" enum:"true"`

	// Start date / time for the Job.
	//
	// StartTime is a required field
	StartTime *time.Time `locationName:"startTime" type:"timestamp" timestampFormat:"unix" required:"true"`

	// Status for the Job.
	//
	// Status is a required field
	Status JobStatus `locationName:"status" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s JobSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s JobSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.CommitId != nil {
		v := *s.CommitId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "commitId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CommitMessage != nil {
		v := *s.CommitMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "commitMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CommitTime != nil {
		v := *s.CommitTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "commitTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.JobArn != nil {
		v := *s.JobArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.JobType) > 0 {
		v := s.JobType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StartTime != nil {
		v := *s.StartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Structure with Production Branch information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/ProductionBranch
type ProductionBranch struct {
	_ struct{} `type:"structure"`

	// Branch Name for Production Branch.
	BranchName *string `locationName:"branchName" min:"1" type:"string"`

	// Last Deploy Time of Production Branch.
	LastDeployTime *time.Time `locationName:"lastDeployTime" type:"timestamp" timestampFormat:"unix"`

	// Status of Production Branch.
	Status *string `locationName:"status" min:"3" type:"string"`

	// Thumbnail Url for Production Branch.
	ThumbnailUrl *string `locationName:"thumbnailUrl" min:"1" type:"string"`
}

// String returns the string representation
func (s ProductionBranch) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ProductionBranch) MarshalFields(e protocol.FieldEncoder) error {
	if s.BranchName != nil {
		v := *s.BranchName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "branchName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastDeployTime != nil {
		v := *s.LastDeployTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastDeployTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Status != nil {
		v := *s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThumbnailUrl != nil {
		v := *s.ThumbnailUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thumbnailUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Structure for an execution step for an execution job, for an Amplify App.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/Step
type Step struct {
	_ struct{} `type:"structure"`

	// Url to teh artifact for the execution step.
	ArtifactsUrl *string `locationName:"artifactsUrl" type:"string"`

	// End date/ time of the execution step.
	//
	// EndTime is a required field
	EndTime *time.Time `locationName:"endTime" type:"timestamp" timestampFormat:"unix" required:"true"`

	// Url to the logs for the execution step.
	LogUrl *string `locationName:"logUrl" type:"string"`

	// List of screenshot Urls for the execution step, if relevant.
	Screenshots map[string]string `locationName:"screenshots" type:"map"`

	// Start date/ time of the execution step.
	//
	// StartTime is a required field
	StartTime *time.Time `locationName:"startTime" type:"timestamp" timestampFormat:"unix" required:"true"`

	// Status of the execution step.
	//
	// Status is a required field
	Status JobStatus `locationName:"status" type:"string" required:"true" enum:"true"`

	// Name of the execution step.
	//
	// StepName is a required field
	StepName *string `locationName:"stepName" type:"string" required:"true"`
}

// String returns the string representation
func (s Step) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Step) MarshalFields(e protocol.FieldEncoder) error {
	if s.ArtifactsUrl != nil {
		v := *s.ArtifactsUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "artifactsUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.LogUrl != nil {
		v := *s.LogUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "logUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Screenshots != nil {
		v := s.Screenshots

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "screenshots", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.StartTime != nil {
		v := *s.StartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StepName != nil {
		v := *s.StepName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stepName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Subdomain for the Domain Association.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/SubDomain
type SubDomain struct {
	_ struct{} `type:"structure"`

	// DNS record for the Subdomain.
	//
	// DnsRecord is a required field
	DnsRecord *string `locationName:"dnsRecord" type:"string" required:"true"`

	// Setting structure for the Subdomain.
	//
	// SubDomainSetting is a required field
	SubDomainSetting *SubDomainSetting `locationName:"subDomainSetting" type:"structure" required:"true"`

	// Verified status of the Subdomain
	//
	// Verified is a required field
	Verified *bool `locationName:"verified" type:"boolean" required:"true"`
}

// String returns the string representation
func (s SubDomain) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SubDomain) MarshalFields(e protocol.FieldEncoder) error {
	if s.DnsRecord != nil {
		v := *s.DnsRecord

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dnsRecord", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SubDomainSetting != nil {
		v := s.SubDomainSetting

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "subDomainSetting", v, metadata)
	}
	if s.Verified != nil {
		v := *s.Verified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "verified", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Setting for the Subdomain.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/SubDomainSetting
type SubDomainSetting struct {
	_ struct{} `type:"structure"`

	// Branch name setting for the Subdomain.
	//
	// BranchName is a required field
	BranchName *string `locationName:"branchName" min:"1" type:"string" required:"true"`

	// Prefix setting for the Subdomain.
	//
	// Prefix is a required field
	Prefix *string `locationName:"prefix" type:"string" required:"true"`
}

// String returns the string representation
func (s SubDomainSetting) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SubDomainSetting) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SubDomainSetting"}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if s.Prefix == nil {
		invalidParams.Add(aws.NewErrParamRequired("Prefix"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SubDomainSetting) MarshalFields(e protocol.FieldEncoder) error {
	if s.BranchName != nil {
		v := *s.BranchName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "branchName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Prefix != nil {
		v := *s.Prefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "prefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}