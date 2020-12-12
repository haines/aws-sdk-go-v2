// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsRestjson1_serializeOpGetClip struct {
}

func (*awsRestjson1_serializeOpGetClip) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetClip) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetClipInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/getClip")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetClipInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetClipInput(v *GetClipInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetClipInput(v *GetClipInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClipFragmentSelector != nil {
		ok := object.Key("ClipFragmentSelector")
		if err := awsRestjson1_serializeDocumentClipFragmentSelector(v.ClipFragmentSelector, ok); err != nil {
			return err
		}
	}

	if v.StreamARN != nil {
		ok := object.Key("StreamARN")
		ok.String(*v.StreamARN)
	}

	if v.StreamName != nil {
		ok := object.Key("StreamName")
		ok.String(*v.StreamName)
	}

	return nil
}

type awsRestjson1_serializeOpGetDASHStreamingSessionURL struct {
}

func (*awsRestjson1_serializeOpGetDASHStreamingSessionURL) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetDASHStreamingSessionURL) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetDASHStreamingSessionURLInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/getDASHStreamingSessionURL")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetDASHStreamingSessionURLInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetDASHStreamingSessionURLInput(v *GetDASHStreamingSessionURLInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetDASHStreamingSessionURLInput(v *GetDASHStreamingSessionURLInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DASHFragmentSelector != nil {
		ok := object.Key("DASHFragmentSelector")
		if err := awsRestjson1_serializeDocumentDASHFragmentSelector(v.DASHFragmentSelector, ok); err != nil {
			return err
		}
	}

	if len(v.DisplayFragmentNumber) > 0 {
		ok := object.Key("DisplayFragmentNumber")
		ok.String(string(v.DisplayFragmentNumber))
	}

	if len(v.DisplayFragmentTimestamp) > 0 {
		ok := object.Key("DisplayFragmentTimestamp")
		ok.String(string(v.DisplayFragmentTimestamp))
	}

	if v.Expires != nil {
		ok := object.Key("Expires")
		ok.Integer(*v.Expires)
	}

	if v.MaxManifestFragmentResults != nil {
		ok := object.Key("MaxManifestFragmentResults")
		ok.Long(*v.MaxManifestFragmentResults)
	}

	if len(v.PlaybackMode) > 0 {
		ok := object.Key("PlaybackMode")
		ok.String(string(v.PlaybackMode))
	}

	if v.StreamARN != nil {
		ok := object.Key("StreamARN")
		ok.String(*v.StreamARN)
	}

	if v.StreamName != nil {
		ok := object.Key("StreamName")
		ok.String(*v.StreamName)
	}

	return nil
}

type awsRestjson1_serializeOpGetHLSStreamingSessionURL struct {
}

func (*awsRestjson1_serializeOpGetHLSStreamingSessionURL) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetHLSStreamingSessionURL) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetHLSStreamingSessionURLInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/getHLSStreamingSessionURL")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetHLSStreamingSessionURLInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetHLSStreamingSessionURLInput(v *GetHLSStreamingSessionURLInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetHLSStreamingSessionURLInput(v *GetHLSStreamingSessionURLInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.ContainerFormat) > 0 {
		ok := object.Key("ContainerFormat")
		ok.String(string(v.ContainerFormat))
	}

	if len(v.DiscontinuityMode) > 0 {
		ok := object.Key("DiscontinuityMode")
		ok.String(string(v.DiscontinuityMode))
	}

	if len(v.DisplayFragmentTimestamp) > 0 {
		ok := object.Key("DisplayFragmentTimestamp")
		ok.String(string(v.DisplayFragmentTimestamp))
	}

	if v.Expires != nil {
		ok := object.Key("Expires")
		ok.Integer(*v.Expires)
	}

	if v.HLSFragmentSelector != nil {
		ok := object.Key("HLSFragmentSelector")
		if err := awsRestjson1_serializeDocumentHLSFragmentSelector(v.HLSFragmentSelector, ok); err != nil {
			return err
		}
	}

	if v.MaxMediaPlaylistFragmentResults != nil {
		ok := object.Key("MaxMediaPlaylistFragmentResults")
		ok.Long(*v.MaxMediaPlaylistFragmentResults)
	}

	if len(v.PlaybackMode) > 0 {
		ok := object.Key("PlaybackMode")
		ok.String(string(v.PlaybackMode))
	}

	if v.StreamARN != nil {
		ok := object.Key("StreamARN")
		ok.String(*v.StreamARN)
	}

	if v.StreamName != nil {
		ok := object.Key("StreamName")
		ok.String(*v.StreamName)
	}

	return nil
}

type awsRestjson1_serializeOpGetMediaForFragmentList struct {
}

func (*awsRestjson1_serializeOpGetMediaForFragmentList) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetMediaForFragmentList) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetMediaForFragmentListInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/getMediaForFragmentList")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetMediaForFragmentListInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetMediaForFragmentListInput(v *GetMediaForFragmentListInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetMediaForFragmentListInput(v *GetMediaForFragmentListInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Fragments != nil {
		ok := object.Key("Fragments")
		if err := awsRestjson1_serializeDocumentFragmentNumberList(v.Fragments, ok); err != nil {
			return err
		}
	}

	if v.StreamName != nil {
		ok := object.Key("StreamName")
		ok.String(*v.StreamName)
	}

	return nil
}

type awsRestjson1_serializeOpListFragments struct {
}

func (*awsRestjson1_serializeOpListFragments) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListFragments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListFragmentsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/listFragments")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentListFragmentsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListFragmentsInput(v *ListFragmentsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentListFragmentsInput(v *ListFragmentsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.FragmentSelector != nil {
		ok := object.Key("FragmentSelector")
		if err := awsRestjson1_serializeDocumentFragmentSelector(v.FragmentSelector, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Long(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.StreamName != nil {
		ok := object.Key("StreamName")
		ok.String(*v.StreamName)
	}

	return nil
}

func awsRestjson1_serializeDocumentClipFragmentSelector(v *types.ClipFragmentSelector, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.FragmentSelectorType) > 0 {
		ok := object.Key("FragmentSelectorType")
		ok.String(string(v.FragmentSelectorType))
	}

	if v.TimestampRange != nil {
		ok := object.Key("TimestampRange")
		if err := awsRestjson1_serializeDocumentClipTimestampRange(v.TimestampRange, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentClipTimestampRange(v *types.ClipTimestampRange, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EndTimestamp != nil {
		ok := object.Key("EndTimestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.EndTimestamp))
	}

	if v.StartTimestamp != nil {
		ok := object.Key("StartTimestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.StartTimestamp))
	}

	return nil
}

func awsRestjson1_serializeDocumentDASHFragmentSelector(v *types.DASHFragmentSelector, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.FragmentSelectorType) > 0 {
		ok := object.Key("FragmentSelectorType")
		ok.String(string(v.FragmentSelectorType))
	}

	if v.TimestampRange != nil {
		ok := object.Key("TimestampRange")
		if err := awsRestjson1_serializeDocumentDASHTimestampRange(v.TimestampRange, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentDASHTimestampRange(v *types.DASHTimestampRange, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EndTimestamp != nil {
		ok := object.Key("EndTimestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.EndTimestamp))
	}

	if v.StartTimestamp != nil {
		ok := object.Key("StartTimestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.StartTimestamp))
	}

	return nil
}

func awsRestjson1_serializeDocumentFragmentNumberList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentFragmentSelector(v *types.FragmentSelector, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.FragmentSelectorType) > 0 {
		ok := object.Key("FragmentSelectorType")
		ok.String(string(v.FragmentSelectorType))
	}

	if v.TimestampRange != nil {
		ok := object.Key("TimestampRange")
		if err := awsRestjson1_serializeDocumentTimestampRange(v.TimestampRange, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentHLSFragmentSelector(v *types.HLSFragmentSelector, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.FragmentSelectorType) > 0 {
		ok := object.Key("FragmentSelectorType")
		ok.String(string(v.FragmentSelectorType))
	}

	if v.TimestampRange != nil {
		ok := object.Key("TimestampRange")
		if err := awsRestjson1_serializeDocumentHLSTimestampRange(v.TimestampRange, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentHLSTimestampRange(v *types.HLSTimestampRange, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EndTimestamp != nil {
		ok := object.Key("EndTimestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.EndTimestamp))
	}

	if v.StartTimestamp != nil {
		ok := object.Key("StartTimestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.StartTimestamp))
	}

	return nil
}

func awsRestjson1_serializeDocumentTimestampRange(v *types.TimestampRange, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EndTimestamp != nil {
		ok := object.Key("EndTimestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.EndTimestamp))
	}

	if v.StartTimestamp != nil {
		ok := object.Key("StartTimestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.StartTimestamp))
	}

	return nil
}