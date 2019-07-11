// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The DELETE request to delete a usage plan of a given plan Id.
type DeleteUsagePlanInput struct {
	_ struct{} `type:"structure"`

	// [Required] The Id of the to-be-deleted usage plan.
	//
	// UsagePlanId is a required field
	UsagePlanId *string `location:"uri" locationName:"usageplanId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUsagePlanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUsagePlanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUsagePlanInput"}

	if s.UsagePlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UsagePlanId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteUsagePlanInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.UsagePlanId != nil {
		v := *s.UsagePlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "usageplanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteUsagePlanOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUsagePlanOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteUsagePlanOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteUsagePlan = "DeleteUsagePlan"

// DeleteUsagePlanRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Deletes a usage plan of a given plan Id.
//
//    // Example sending a request using DeleteUsagePlanRequest.
//    req := client.DeleteUsagePlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteUsagePlanRequest(input *DeleteUsagePlanInput) DeleteUsagePlanRequest {
	op := &aws.Operation{
		Name:       opDeleteUsagePlan,
		HTTPMethod: "DELETE",
		HTTPPath:   "/usageplans/{usageplanId}",
	}

	if input == nil {
		input = &DeleteUsagePlanInput{}
	}

	req := c.newRequest(op, input, &DeleteUsagePlanOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteUsagePlanRequest{Request: req, Input: input, Copy: c.DeleteUsagePlanRequest}
}

// DeleteUsagePlanRequest is the request type for the
// DeleteUsagePlan API operation.
type DeleteUsagePlanRequest struct {
	*aws.Request
	Input *DeleteUsagePlanInput
	Copy  func(*DeleteUsagePlanInput) DeleteUsagePlanRequest
}

// Send marshals and sends the DeleteUsagePlan API request.
func (r DeleteUsagePlanRequest) Send(ctx context.Context) (*DeleteUsagePlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUsagePlanResponse{
		DeleteUsagePlanOutput: r.Request.Data.(*DeleteUsagePlanOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUsagePlanResponse is the response type for the
// DeleteUsagePlan API operation.
type DeleteUsagePlanResponse struct {
	*DeleteUsagePlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUsagePlan request.
func (r *DeleteUsagePlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}