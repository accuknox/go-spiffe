// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: workload.proto

package workload_custom

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The X509SVIDRequest message conveys parameters for requesting an X.509-SVID.
// There are currently no request parameters.
type X509SVIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta map[string]string `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *X509SVIDRequest) Reset() {
	*x = X509SVIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X509SVIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509SVIDRequest) ProtoMessage() {}

func (x *X509SVIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509SVIDRequest.ProtoReflect.Descriptor instead.
func (*X509SVIDRequest) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{0}
}

func (x *X509SVIDRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

// The X509SVIDResponse message carries X.509-SVIDs and related information,
// including a set of global CRLs and a list of bundles the workload may use
// for federating with foreign trust domains.
type X509SVIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. A list of X509SVID messages, each of which includes a single
	// X.509-SVID, its private key, and the bundle for the trust domain.
	Svids []*X509SVID `protobuf:"bytes,1,rep,name=svids,proto3" json:"svids,omitempty"`
	// Optional. ASN.1 DER encoded certificate revocation lists.
	Crl [][]byte `protobuf:"bytes,2,rep,name=crl,proto3" json:"crl,omitempty"`
	// Optional. CA certificate bundles belonging to foreign trust domains that
	// the workload should trust, keyed by the SPIFFE ID of the foreign trust
	// domain. Bundles are ASN.1 DER encoded.
	FederatedBundles map[string][]byte `protobuf:"bytes,3,rep,name=federated_bundles,json=federatedBundles,proto3" json:"federated_bundles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *X509SVIDResponse) Reset() {
	*x = X509SVIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X509SVIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509SVIDResponse) ProtoMessage() {}

func (x *X509SVIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509SVIDResponse.ProtoReflect.Descriptor instead.
func (*X509SVIDResponse) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{1}
}

func (x *X509SVIDResponse) GetSvids() []*X509SVID {
	if x != nil {
		return x.Svids
	}
	return nil
}

func (x *X509SVIDResponse) GetCrl() [][]byte {
	if x != nil {
		return x.Crl
	}
	return nil
}

func (x *X509SVIDResponse) GetFederatedBundles() map[string][]byte {
	if x != nil {
		return x.FederatedBundles
	}
	return nil
}

// The X509SVID message carries a single SVID and all associated information,
// including the X.509 bundle for the trust domain.
type X509SVID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The SPIFFE ID of the SVID in this entry
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// Required. ASN.1 DER encoded certificate chain. MAY include
	// intermediates, the leaf certificate (or SVID itself) MUST come first.
	X509Svid []byte `protobuf:"bytes,2,opt,name=x509_svid,json=x509Svid,proto3" json:"x509_svid,omitempty"`
	// Required. ASN.1 DER encoded PKCS#8 private key. MUST be unencrypted.
	X509SvidKey []byte `protobuf:"bytes,3,opt,name=x509_svid_key,json=x509SvidKey,proto3" json:"x509_svid_key,omitempty"`
	// Required. ASN.1 DER encoded X.509 bundle for the trust domain.
	Bundle []byte `protobuf:"bytes,4,opt,name=bundle,proto3" json:"bundle,omitempty"`
}

func (x *X509SVID) Reset() {
	*x = X509SVID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X509SVID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509SVID) ProtoMessage() {}

func (x *X509SVID) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509SVID.ProtoReflect.Descriptor instead.
func (*X509SVID) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{2}
}

func (x *X509SVID) GetSpiffeId() string {
	if x != nil {
		return x.SpiffeId
	}
	return ""
}

func (x *X509SVID) GetX509Svid() []byte {
	if x != nil {
		return x.X509Svid
	}
	return nil
}

func (x *X509SVID) GetX509SvidKey() []byte {
	if x != nil {
		return x.X509SvidKey
	}
	return nil
}

func (x *X509SVID) GetBundle() []byte {
	if x != nil {
		return x.Bundle
	}
	return nil
}

// The X509BundlesRequest message conveys parameters for requesting X.509
// bundles. There are currently no such parameters.
type X509BundlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta map[string]string `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *X509BundlesRequest) Reset() {
	*x = X509BundlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X509BundlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509BundlesRequest) ProtoMessage() {}

func (x *X509BundlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509BundlesRequest.ProtoReflect.Descriptor instead.
func (*X509BundlesRequest) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{3}
}

func (x *X509BundlesRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

// The X509BundlesResponse message carries a set of global CRLs and a map of
// trust bundles the workload should trust.
type X509BundlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. ASN.1 DER encoded certificate revocation lists.
	Crl [][]byte `protobuf:"bytes,1,rep,name=crl,proto3" json:"crl,omitempty"`
	// Required. CA certificate bundles belonging to trust domains that the
	// workload should trust, keyed by the SPIFFE ID of the trust domain.
	// Bundles are ASN.1 DER encoded.
	Bundles map[string][]byte `protobuf:"bytes,2,rep,name=bundles,proto3" json:"bundles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *X509BundlesResponse) Reset() {
	*x = X509BundlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X509BundlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509BundlesResponse) ProtoMessage() {}

func (x *X509BundlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509BundlesResponse.ProtoReflect.Descriptor instead.
func (*X509BundlesResponse) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{4}
}

func (x *X509BundlesResponse) GetCrl() [][]byte {
	if x != nil {
		return x.Crl
	}
	return nil
}

func (x *X509BundlesResponse) GetBundles() map[string][]byte {
	if x != nil {
		return x.Bundles
	}
	return nil
}

type JWTSVIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The audience(s) the workload intends to authenticate against.
	Audience []string `protobuf:"bytes,1,rep,name=audience,proto3" json:"audience,omitempty"`
	// Optional. The requested SPIFFE ID for the JWT-SVID. If unset, all
	// JWT-SVIDs to which the workload is entitled are requested.
	SpiffeId string `protobuf:"bytes,2,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
}

func (x *JWTSVIDRequest) Reset() {
	*x = JWTSVIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTSVIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTSVIDRequest) ProtoMessage() {}

func (x *JWTSVIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTSVIDRequest.ProtoReflect.Descriptor instead.
func (*JWTSVIDRequest) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{5}
}

func (x *JWTSVIDRequest) GetAudience() []string {
	if x != nil {
		return x.Audience
	}
	return nil
}

func (x *JWTSVIDRequest) GetSpiffeId() string {
	if x != nil {
		return x.SpiffeId
	}
	return ""
}

// The JWTSVIDResponse message conveys JWT-SVIDs.
type JWTSVIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The list of returned JWT-SVIDs.
	Svids []*JWTSVID `protobuf:"bytes,1,rep,name=svids,proto3" json:"svids,omitempty"`
}

func (x *JWTSVIDResponse) Reset() {
	*x = JWTSVIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTSVIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTSVIDResponse) ProtoMessage() {}

func (x *JWTSVIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTSVIDResponse.ProtoReflect.Descriptor instead.
func (*JWTSVIDResponse) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{6}
}

func (x *JWTSVIDResponse) GetSvids() []*JWTSVID {
	if x != nil {
		return x.Svids
	}
	return nil
}

// The JWTSVID message carries the JWT-SVID token and associated metadata.
type JWTSVID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The SPIFFE ID of the JWT-SVID.
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// Required. Encoded JWT using JWS Compact Serialization.
	Svid string `protobuf:"bytes,2,opt,name=svid,proto3" json:"svid,omitempty"`
}

func (x *JWTSVID) Reset() {
	*x = JWTSVID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTSVID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTSVID) ProtoMessage() {}

func (x *JWTSVID) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTSVID.ProtoReflect.Descriptor instead.
func (*JWTSVID) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{7}
}

func (x *JWTSVID) GetSpiffeId() string {
	if x != nil {
		return x.SpiffeId
	}
	return ""
}

func (x *JWTSVID) GetSvid() string {
	if x != nil {
		return x.Svid
	}
	return ""
}

// The JWTBundlesRequest message conveys parameters for requesting JWT bundles.
// There are currently no such parameters.
type JWTBundlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JWTBundlesRequest) Reset() {
	*x = JWTBundlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTBundlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTBundlesRequest) ProtoMessage() {}

func (x *JWTBundlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTBundlesRequest.ProtoReflect.Descriptor instead.
func (*JWTBundlesRequest) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{8}
}

// The JWTBundlesReponse conveys JWT bundles.
type JWTBundlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. JWK encoded JWT bundles, keyed by the SPIFFE ID of the trust
	// domain.
	Bundles map[string][]byte `protobuf:"bytes,1,rep,name=bundles,proto3" json:"bundles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *JWTBundlesResponse) Reset() {
	*x = JWTBundlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTBundlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTBundlesResponse) ProtoMessage() {}

func (x *JWTBundlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTBundlesResponse.ProtoReflect.Descriptor instead.
func (*JWTBundlesResponse) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{9}
}

func (x *JWTBundlesResponse) GetBundles() map[string][]byte {
	if x != nil {
		return x.Bundles
	}
	return nil
}

// The ValidateJWTSVIDRequest message conveys request parameters for
// JWT-SVID validation.
type ValidateJWTSVIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The audience of the validating party. The JWT-SVID must
	// contain an audience claim which contains this value in order to
	// succesfully validate.
	Audience string `protobuf:"bytes,1,opt,name=audience,proto3" json:"audience,omitempty"`
	// Required. The JWT-SVID to validate, encoded using JWS Compact
	// Serialization.
	Svid string `protobuf:"bytes,2,opt,name=svid,proto3" json:"svid,omitempty"`
}

func (x *ValidateJWTSVIDRequest) Reset() {
	*x = ValidateJWTSVIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateJWTSVIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateJWTSVIDRequest) ProtoMessage() {}

func (x *ValidateJWTSVIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateJWTSVIDRequest.ProtoReflect.Descriptor instead.
func (*ValidateJWTSVIDRequest) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{10}
}

func (x *ValidateJWTSVIDRequest) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

func (x *ValidateJWTSVIDRequest) GetSvid() string {
	if x != nil {
		return x.Svid
	}
	return ""
}

// The ValidateJWTSVIDReponse message conveys the JWT-SVID validation results.
type ValidateJWTSVIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The SPIFFE ID of the validated JWT-SVID.
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// Optional. Arbitrary claims contained within the payload of the validated
	// JWT-SVID.
	Claims *structpb.Struct `protobuf:"bytes,2,opt,name=claims,proto3" json:"claims,omitempty"`
}

func (x *ValidateJWTSVIDResponse) Reset() {
	*x = ValidateJWTSVIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workload_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateJWTSVIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateJWTSVIDResponse) ProtoMessage() {}

func (x *ValidateJWTSVIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workload_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateJWTSVIDResponse.ProtoReflect.Descriptor instead.
func (*ValidateJWTSVIDResponse) Descriptor() ([]byte, []int) {
	return file_workload_proto_rawDescGZIP(), []int{11}
}

func (x *ValidateJWTSVIDResponse) GetSpiffeId() string {
	if x != nil {
		return x.SpiffeId
	}
	return ""
}

func (x *ValidateJWTSVIDResponse) GetClaims() *structpb.Struct {
	if x != nil {
		return x.Claims
	}
	return nil
}

var File_workload_proto protoreflect.FileDescriptor

var file_workload_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a,
	0x0a, 0x0f, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe0, 0x01, 0x0a, 0x10, 0x58,
	0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x05, 0x73, 0x76, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x52, 0x05, 0x73, 0x76, 0x69, 0x64, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x63,
	0x72, 0x6c, 0x12, 0x54, 0x0a, 0x11, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x1a, 0x43, 0x0a, 0x15, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80, 0x01,
	0x0a, 0x08, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70,
	0x69, 0x66, 0x66, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x70, 0x69, 0x66, 0x66, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x78, 0x35, 0x30, 0x39, 0x5f,
	0x73, 0x76, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x78, 0x35, 0x30, 0x39,
	0x53, 0x76, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x73, 0x76, 0x69,
	0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x78, 0x35, 0x30,
	0x39, 0x53, 0x76, 0x69, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x22, 0x80, 0x01, 0x0a, 0x12, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xa0, 0x01, 0x0a, 0x13, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x63, 0x72, 0x6c, 0x12, 0x3b, 0x0a,
	0x07, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x49, 0x0a, 0x0e, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x49,
	0x64, 0x22, 0x31, 0x0a, 0x0f, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x73, 0x76, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x52, 0x05, 0x73,
	0x76, 0x69, 0x64, 0x73, 0x22, 0x3a, 0x0a, 0x07, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x76, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x76, 0x69, 0x64,
	0x22, 0x13, 0x0a, 0x11, 0x4a, 0x57, 0x54, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x12, 0x4a, 0x57, 0x54, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x4a, 0x57, 0x54, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x48, 0x0a, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x76,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x76, 0x69, 0x64, 0x22, 0x67,
	0x0a, 0x17, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x69,
	0x66, 0x66, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70,
	0x69, 0x66, 0x66, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x32, 0xc3, 0x02, 0x0a, 0x11, 0x53, 0x70, 0x69, 0x66,
	0x66, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x50, 0x49, 0x12, 0x36, 0x0a,
	0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x12, 0x10,
	0x2e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x10, 0x46, 0x65, 0x74, 0x63, 0x68, 0x58, 0x35,
	0x30, 0x39, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x58, 0x35, 0x30, 0x39,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x58, 0x35, 0x30, 0x39, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x31, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4a,
	0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x12, 0x0f, 0x2e, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x4a, 0x57, 0x54, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x4a,
	0x57, 0x54, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x4a, 0x57, 0x54, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x12, 0x17, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x57,
	0x54, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x48, 0x5a,
	0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x75,
	0x6b, 0x6e, 0x6f, 0x78, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x76,
	0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workload_proto_rawDescOnce sync.Once
	file_workload_proto_rawDescData = file_workload_proto_rawDesc
)

func file_workload_proto_rawDescGZIP() []byte {
	file_workload_proto_rawDescOnce.Do(func() {
		file_workload_proto_rawDescData = protoimpl.X.CompressGZIP(file_workload_proto_rawDescData)
	})
	return file_workload_proto_rawDescData
}

var file_workload_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_workload_proto_goTypes = []interface{}{
	(*X509SVIDRequest)(nil),         // 0: X509SVIDRequest
	(*X509SVIDResponse)(nil),        // 1: X509SVIDResponse
	(*X509SVID)(nil),                // 2: X509SVID
	(*X509BundlesRequest)(nil),      // 3: X509BundlesRequest
	(*X509BundlesResponse)(nil),     // 4: X509BundlesResponse
	(*JWTSVIDRequest)(nil),          // 5: JWTSVIDRequest
	(*JWTSVIDResponse)(nil),         // 6: JWTSVIDResponse
	(*JWTSVID)(nil),                 // 7: JWTSVID
	(*JWTBundlesRequest)(nil),       // 8: JWTBundlesRequest
	(*JWTBundlesResponse)(nil),      // 9: JWTBundlesResponse
	(*ValidateJWTSVIDRequest)(nil),  // 10: ValidateJWTSVIDRequest
	(*ValidateJWTSVIDResponse)(nil), // 11: ValidateJWTSVIDResponse
	nil,                             // 12: X509SVIDRequest.MetaEntry
	nil,                             // 13: X509SVIDResponse.FederatedBundlesEntry
	nil,                             // 14: X509BundlesRequest.MetaEntry
	nil,                             // 15: X509BundlesResponse.BundlesEntry
	nil,                             // 16: JWTBundlesResponse.BundlesEntry
	(*structpb.Struct)(nil),         // 17: google.protobuf.Struct
}
var file_workload_proto_depIdxs = []int32{
	12, // 0: X509SVIDRequest.meta:type_name -> X509SVIDRequest.MetaEntry
	2,  // 1: X509SVIDResponse.svids:type_name -> X509SVID
	13, // 2: X509SVIDResponse.federated_bundles:type_name -> X509SVIDResponse.FederatedBundlesEntry
	14, // 3: X509BundlesRequest.meta:type_name -> X509BundlesRequest.MetaEntry
	15, // 4: X509BundlesResponse.bundles:type_name -> X509BundlesResponse.BundlesEntry
	7,  // 5: JWTSVIDResponse.svids:type_name -> JWTSVID
	16, // 6: JWTBundlesResponse.bundles:type_name -> JWTBundlesResponse.BundlesEntry
	17, // 7: ValidateJWTSVIDResponse.claims:type_name -> google.protobuf.Struct
	0,  // 8: SpiffeWorkloadAPI.FetchX509SVID:input_type -> X509SVIDRequest
	3,  // 9: SpiffeWorkloadAPI.FetchX509Bundles:input_type -> X509BundlesRequest
	5,  // 10: SpiffeWorkloadAPI.FetchJWTSVID:input_type -> JWTSVIDRequest
	8,  // 11: SpiffeWorkloadAPI.FetchJWTBundles:input_type -> JWTBundlesRequest
	10, // 12: SpiffeWorkloadAPI.ValidateJWTSVID:input_type -> ValidateJWTSVIDRequest
	1,  // 13: SpiffeWorkloadAPI.FetchX509SVID:output_type -> X509SVIDResponse
	4,  // 14: SpiffeWorkloadAPI.FetchX509Bundles:output_type -> X509BundlesResponse
	6,  // 15: SpiffeWorkloadAPI.FetchJWTSVID:output_type -> JWTSVIDResponse
	9,  // 16: SpiffeWorkloadAPI.FetchJWTBundles:output_type -> JWTBundlesResponse
	11, // 17: SpiffeWorkloadAPI.ValidateJWTSVID:output_type -> ValidateJWTSVIDResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_workload_proto_init() }
func file_workload_proto_init() {
	if File_workload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X509SVIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X509SVIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X509SVID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X509BundlesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X509BundlesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTSVIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workload_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTSVIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workload_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTSVID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workload_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTBundlesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workload_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTBundlesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workload_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateJWTSVIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workload_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateJWTSVIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_workload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workload_proto_goTypes,
		DependencyIndexes: file_workload_proto_depIdxs,
		MessageInfos:      file_workload_proto_msgTypes,
	}.Build()
	File_workload_proto = out.File
	file_workload_proto_rawDesc = nil
	file_workload_proto_goTypes = nil
	file_workload_proto_depIdxs = nil
}
