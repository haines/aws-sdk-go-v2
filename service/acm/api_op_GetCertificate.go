// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/GetCertificateRequest
type GetCertificateInput struct {
	_ struct{} `type:"structure"`

	// String that contains a certificate ARN in the following format:
	//
	// arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// CertificateArn is a required field
	CertificateArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCertificateInput"}

	if s.CertificateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateArn"))
	}
	if s.CertificateArn != nil && len(*s.CertificateArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/GetCertificateResponse
type GetCertificateOutput struct {
	_ struct{} `type:"structure"`

	// String that contains the ACM certificate represented by the ARN specified
	// at input.
	Certificate *string `min:"1" type:"string"`

	// The certificate chain that contains the root certificate issued by the certificate
	// authority (CA).
	CertificateChain *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCertificate = "GetCertificate"

// GetCertificateRequest returns a request value for making API operation for
// AWS Certificate Manager.
//
// Retrieves a certificate specified by an ARN and its certificate chain . The
// chain is an ordered list of certificates that contains the end entity certificate,
// intermediate certificates of subordinate CAs, and the root certificate in
// that order. The certificate and certificate chain are base64 encoded. If
// you want to decode the certificate to see the individual fields, you can
// use OpenSSL.
//
//    // Example sending a request using GetCertificateRequest.
//    req := client.GetCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/GetCertificate
func (c *Client) GetCertificateRequest(input *GetCertificateInput) GetCertificateRequest {
	op := &aws.Operation{
		Name:       opGetCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCertificateInput{}
	}

	req := c.newRequest(op, input, &GetCertificateOutput{})
	return GetCertificateRequest{Request: req, Input: input, Copy: c.GetCertificateRequest}
}

// GetCertificateRequest is the request type for the
// GetCertificate API operation.
type GetCertificateRequest struct {
	*aws.Request
	Input *GetCertificateInput
	Copy  func(*GetCertificateInput) GetCertificateRequest
}

// Send marshals and sends the GetCertificate API request.
func (r GetCertificateRequest) Send(ctx context.Context) (*GetCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCertificateResponse{
		GetCertificateOutput: r.Request.Data.(*GetCertificateOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCertificateResponse is the response type for the
// GetCertificate API operation.
type GetCertificateResponse struct {
	*GetCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCertificate request.
func (r *GetCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}