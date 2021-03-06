// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateClassifierInput struct {
	_ struct{} `type:"structure"`

	// A CsvClassifier object with updated fields.
	CsvClassifier *UpdateCsvClassifierRequest `type:"structure"`

	// A GrokClassifier object with updated fields.
	GrokClassifier *UpdateGrokClassifierRequest `type:"structure"`

	// A JsonClassifier object with updated fields.
	JsonClassifier *UpdateJsonClassifierRequest `type:"structure"`

	// An XMLClassifier object with updated fields.
	XMLClassifier *UpdateXMLClassifierRequest `type:"structure"`
}

// String returns the string representation
func (s UpdateClassifierInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateClassifierInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateClassifierInput"}
	if s.CsvClassifier != nil {
		if err := s.CsvClassifier.Validate(); err != nil {
			invalidParams.AddNested("CsvClassifier", err.(aws.ErrInvalidParams))
		}
	}
	if s.GrokClassifier != nil {
		if err := s.GrokClassifier.Validate(); err != nil {
			invalidParams.AddNested("GrokClassifier", err.(aws.ErrInvalidParams))
		}
	}
	if s.JsonClassifier != nil {
		if err := s.JsonClassifier.Validate(); err != nil {
			invalidParams.AddNested("JsonClassifier", err.(aws.ErrInvalidParams))
		}
	}
	if s.XMLClassifier != nil {
		if err := s.XMLClassifier.Validate(); err != nil {
			invalidParams.AddNested("XMLClassifier", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateClassifierOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateClassifierOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateClassifier = "UpdateClassifier"

// UpdateClassifierRequest returns a request value for making API operation for
// AWS Glue.
//
// Modifies an existing classifier (a GrokClassifier, an XMLClassifier, a JsonClassifier,
// or a CsvClassifier, depending on which field is present).
//
//    // Example sending a request using UpdateClassifierRequest.
//    req := client.UpdateClassifierRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateClassifier
func (c *Client) UpdateClassifierRequest(input *UpdateClassifierInput) UpdateClassifierRequest {
	op := &aws.Operation{
		Name:       opUpdateClassifier,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateClassifierInput{}
	}

	req := c.newRequest(op, input, &UpdateClassifierOutput{})
	return UpdateClassifierRequest{Request: req, Input: input, Copy: c.UpdateClassifierRequest}
}

// UpdateClassifierRequest is the request type for the
// UpdateClassifier API operation.
type UpdateClassifierRequest struct {
	*aws.Request
	Input *UpdateClassifierInput
	Copy  func(*UpdateClassifierInput) UpdateClassifierRequest
}

// Send marshals and sends the UpdateClassifier API request.
func (r UpdateClassifierRequest) Send(ctx context.Context) (*UpdateClassifierResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateClassifierResponse{
		UpdateClassifierOutput: r.Request.Data.(*UpdateClassifierOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateClassifierResponse is the response type for the
// UpdateClassifier API operation.
type UpdateClassifierResponse struct {
	*UpdateClassifierOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateClassifier request.
func (r *UpdateClassifierResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
