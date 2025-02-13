//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

//
// This is a proxy for MAP application C/D 3GPP 29.002
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: feg/protos/hlr/hlr_proxy.proto

package hlr

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This RPC converts Result-Code from Altran MAP Protocol into gRPC status codes
// ErrorCode reflects Experimental-Result values which are 3GPP failures
// to be processed by HLR Proxy.
type ErrorCode int32

const (
	// Default success code
	ErrorCode_SUCCESS                         ErrorCode = 0
	ErrorCode_UNABLE_TO_DELIVER               ErrorCode = 1
	ErrorCode_AUTHENTICATION_REJECTED         ErrorCode = 2
	ErrorCode_AUTHENTICATION_DATA_UNAVAILABLE ErrorCode = 3
	ErrorCode_UNKNOWN_SUBSCRIBER              ErrorCode = 4
	ErrorCode_NO_PATH_TO_HLR                  ErrorCode = 5
	ErrorCode_NO_HLR_IN_ACTIVE_STATE          ErrorCode = 6
	ErrorCode_NO_RESP_FROM_PEER               ErrorCode = 7
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "SUCCESS",
		1: "UNABLE_TO_DELIVER",
		2: "AUTHENTICATION_REJECTED",
		3: "AUTHENTICATION_DATA_UNAVAILABLE",
		4: "UNKNOWN_SUBSCRIBER",
		5: "NO_PATH_TO_HLR",
		6: "NO_HLR_IN_ACTIVE_STATE",
		7: "NO_RESP_FROM_PEER",
	}
	ErrorCode_value = map[string]int32{
		"SUCCESS":                         0,
		"UNABLE_TO_DELIVER":               1,
		"AUTHENTICATION_REJECTED":         2,
		"AUTHENTICATION_DATA_UNAVAILABLE": 3,
		"UNKNOWN_SUBSCRIBER":              4,
		"NO_PATH_TO_HLR":                  5,
		"NO_HLR_IN_ACTIVE_STATE":          6,
		"NO_RESP_FROM_PEER":               7,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_feg_protos_hlr_hlr_proxy_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_feg_protos_hlr_hlr_proxy_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_feg_protos_hlr_hlr_proxy_proto_rawDescGZIP(), []int{0}
}

type AuthInfoReq_ResyncInfo_Len int32

const (
	AuthInfoReq_ResyncInfo_ZERO_LEN AuthInfoReq_ResyncInfo_Len = 0
	AuthInfoReq_ResyncInfo_RAND_LEN AuthInfoReq_ResyncInfo_Len = 16
	AuthInfoReq_ResyncInfo_AUTH_LEN AuthInfoReq_ResyncInfo_Len = 16
)

// Enum value maps for AuthInfoReq_ResyncInfo_Len.
var (
	AuthInfoReq_ResyncInfo_Len_name = map[int32]string{
		0:  "ZERO_LEN",
		16: "RAND_LEN",
		// Duplicate value: 16: "AUTH_LEN",
	}
	AuthInfoReq_ResyncInfo_Len_value = map[string]int32{
		"ZERO_LEN": 0,
		"RAND_LEN": 16,
		"AUTH_LEN": 16,
	}
)

func (x AuthInfoReq_ResyncInfo_Len) Enum() *AuthInfoReq_ResyncInfo_Len {
	p := new(AuthInfoReq_ResyncInfo_Len)
	*p = x
	return p
}

func (x AuthInfoReq_ResyncInfo_Len) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthInfoReq_ResyncInfo_Len) Descriptor() protoreflect.EnumDescriptor {
	return file_feg_protos_hlr_hlr_proxy_proto_enumTypes[1].Descriptor()
}

func (AuthInfoReq_ResyncInfo_Len) Type() protoreflect.EnumType {
	return &file_feg_protos_hlr_hlr_proxy_proto_enumTypes[1]
}

func (x AuthInfoReq_ResyncInfo_Len) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthInfoReq_ResyncInfo_Len.Descriptor instead.
func (AuthInfoReq_ResyncInfo_Len) EnumDescriptor() ([]byte, []int) {
	return file_feg_protos_hlr_hlr_proxy_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Authentication Information Request (MAP 29.002 section 8.5.2)
type AuthInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Subscriber identifier
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// Number of umts vectors to request in response
	NumRequestedUmtsVectors uint32 `protobuf:"varint,2,opt,name=num_requested_umts_vectors,json=numRequestedUmtsVectors,proto3" json:"num_requested_umts_vectors,omitempty"`
	//ResyncInfo containing RAND and AUTS in the case of a resync attach
	ResyncInfo *AuthInfoReq_ResyncInfo `protobuf:"bytes,3,opt,name=resync_info,json=resyncInfo,proto3" json:"resync_info,omitempty"`
}

func (x *AuthInfoReq) Reset() {
	*x = AuthInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_hlr_hlr_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInfoReq) ProtoMessage() {}

func (x *AuthInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_hlr_hlr_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInfoReq.ProtoReflect.Descriptor instead.
func (*AuthInfoReq) Descriptor() ([]byte, []int) {
	return file_feg_protos_hlr_hlr_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *AuthInfoReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AuthInfoReq) GetNumRequestedUmtsVectors() uint32 {
	if x != nil {
		return x.NumRequestedUmtsVectors
	}
	return 0
}

func (x *AuthInfoReq) GetResyncInfo() *AuthInfoReq_ResyncInfo {
	if x != nil {
		return x.ResyncInfo
	}
	return nil
}

// Authentication Information Answer (MAP 29.002 Section 8.5.2)
type AuthInfoAns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// EPC error code on failure
	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=magma.feg.hlr.ErrorCode" json:"error_code,omitempty"`
	// Authentication vectors matching the requested number
	UmtsVectors []*AuthInfoAns_UMTSVector `protobuf:"bytes,2,rep,name=umts_vectors,json=umtsVectors,proto3" json:"umts_vectors,omitempty"`
}

func (x *AuthInfoAns) Reset() {
	*x = AuthInfoAns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_hlr_hlr_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInfoAns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInfoAns) ProtoMessage() {}

func (x *AuthInfoAns) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_hlr_hlr_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInfoAns.ProtoReflect.Descriptor instead.
func (*AuthInfoAns) Descriptor() ([]byte, []int) {
	return file_feg_protos_hlr_hlr_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *AuthInfoAns) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_SUCCESS
}

func (x *AuthInfoAns) GetUmtsVectors() []*AuthInfoAns_UMTSVector {
	if x != nil {
		return x.UmtsVectors
	}
	return nil
}

type AuthInfoReq_ResyncInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rand []byte `protobuf:"bytes,1,opt,name=rand,proto3" json:"rand,omitempty"`
	Autn []byte `protobuf:"bytes,2,opt,name=autn,proto3" json:"autn,omitempty"`
}

func (x *AuthInfoReq_ResyncInfo) Reset() {
	*x = AuthInfoReq_ResyncInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_hlr_hlr_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInfoReq_ResyncInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInfoReq_ResyncInfo) ProtoMessage() {}

func (x *AuthInfoReq_ResyncInfo) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_hlr_hlr_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInfoReq_ResyncInfo.ProtoReflect.Descriptor instead.
func (*AuthInfoReq_ResyncInfo) Descriptor() ([]byte, []int) {
	return file_feg_protos_hlr_hlr_proxy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AuthInfoReq_ResyncInfo) GetRand() []byte {
	if x != nil {
		return x.Rand
	}
	return nil
}

func (x *AuthInfoReq_ResyncInfo) GetAutn() []byte {
	if x != nil {
		return x.Autn
	}
	return nil
}

// For details about fields read 3GPP 33.401
type AuthInfoAns_UMTSVector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rand []byte `protobuf:"bytes,1,opt,name=rand,proto3" json:"rand,omitempty"`
	Xres []byte `protobuf:"bytes,2,opt,name=xres,proto3" json:"xres,omitempty"`
	Ck   []byte `protobuf:"bytes,3,opt,name=ck,proto3" json:"ck,omitempty"`
	Ik   []byte `protobuf:"bytes,4,opt,name=ik,proto3" json:"ik,omitempty"`
	Autn []byte `protobuf:"bytes,5,opt,name=autn,proto3" json:"autn,omitempty"`
}

func (x *AuthInfoAns_UMTSVector) Reset() {
	*x = AuthInfoAns_UMTSVector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_hlr_hlr_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInfoAns_UMTSVector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInfoAns_UMTSVector) ProtoMessage() {}

func (x *AuthInfoAns_UMTSVector) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_hlr_hlr_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInfoAns_UMTSVector.ProtoReflect.Descriptor instead.
func (*AuthInfoAns_UMTSVector) Descriptor() ([]byte, []int) {
	return file_feg_protos_hlr_hlr_proxy_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AuthInfoAns_UMTSVector) GetRand() []byte {
	if x != nil {
		return x.Rand
	}
	return nil
}

func (x *AuthInfoAns_UMTSVector) GetXres() []byte {
	if x != nil {
		return x.Xres
	}
	return nil
}

func (x *AuthInfoAns_UMTSVector) GetCk() []byte {
	if x != nil {
		return x.Ck
	}
	return nil
}

func (x *AuthInfoAns_UMTSVector) GetIk() []byte {
	if x != nil {
		return x.Ik
	}
	return nil
}

func (x *AuthInfoAns_UMTSVector) GetAutn() []byte {
	if x != nil {
		return x.Autn
	}
	return nil
}

var File_feg_protos_hlr_hlr_proxy_proto protoreflect.FileDescriptor

var file_feg_protos_hlr_hlr_proxy_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x66, 0x65, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x6c, 0x72,
	0x2f, 0x68, 0x6c, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x68, 0x6c, 0x72, 0x22,
	0x9a, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x1a,
	0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x6d,
	0x74, 0x73, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x17, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x55, 0x6d,
	0x74, 0x73, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x46, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x68, 0x6c, 0x72, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x2e, 0x52, 0x65, 0x73, 0x79, 0x6e,
	0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x69, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72,
	0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x61, 0x75, 0x74, 0x6e, 0x22, 0x33, 0x0a, 0x03, 0x4c, 0x65, 0x6e, 0x12, 0x0c,
	0x0a, 0x08, 0x5a, 0x45, 0x52, 0x4f, 0x5f, 0x4c, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x41, 0x4e, 0x44, 0x5f, 0x4c, 0x45, 0x4e, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x55,
	0x54, 0x48, 0x5f, 0x4c, 0x45, 0x4e, 0x10, 0x10, 0x1a, 0x02, 0x10, 0x01, 0x22, 0xfa, 0x01, 0x0a,
	0x0b, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x0a,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x68, 0x6c, 0x72,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x75, 0x6d, 0x74, 0x73, 0x5f, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x68, 0x6c, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x41, 0x6e, 0x73, 0x2e, 0x55, 0x4d, 0x54, 0x53, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x0b, 0x75, 0x6d, 0x74, 0x73, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x1a,
	0x68, 0x0a, 0x0a, 0x55, 0x4d, 0x54, 0x53, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x61, 0x6e,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x78, 0x72, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x75, 0x74, 0x6e, 0x2a, 0xd0, 0x01, 0x0a, 0x09, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x54,
	0x4f, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x41,
	0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45,
	0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x55, 0x54, 0x48,
	0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x16, 0x0a,
	0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49,
	0x42, 0x45, 0x52, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x5f, 0x50, 0x41, 0x54, 0x48,
	0x5f, 0x54, 0x4f, 0x5f, 0x48, 0x4c, 0x52, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x4e, 0x4f, 0x5f,
	0x48, 0x4c, 0x52, 0x5f, 0x49, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x53, 0x50,
	0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x50, 0x45, 0x45, 0x52, 0x10, 0x07, 0x32, 0x50, 0x0a, 0x08,
	0x48, 0x6c, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x44, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67,
	0x2e, 0x68, 0x6c, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x68, 0x6c, 0x72,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x6e, 0x73, 0x22, 0x00, 0x42, 0x1f,
	0x5a, 0x1d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x66, 0x65, 0x67, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x6c, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feg_protos_hlr_hlr_proxy_proto_rawDescOnce sync.Once
	file_feg_protos_hlr_hlr_proxy_proto_rawDescData = file_feg_protos_hlr_hlr_proxy_proto_rawDesc
)

func file_feg_protos_hlr_hlr_proxy_proto_rawDescGZIP() []byte {
	file_feg_protos_hlr_hlr_proxy_proto_rawDescOnce.Do(func() {
		file_feg_protos_hlr_hlr_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_feg_protos_hlr_hlr_proxy_proto_rawDescData)
	})
	return file_feg_protos_hlr_hlr_proxy_proto_rawDescData
}

var file_feg_protos_hlr_hlr_proxy_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_feg_protos_hlr_hlr_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_feg_protos_hlr_hlr_proxy_proto_goTypes = []interface{}{
	(ErrorCode)(0),                  // 0: magma.feg.hlr.ErrorCode
	(AuthInfoReq_ResyncInfo_Len)(0), // 1: magma.feg.hlr.AuthInfoReq.ResyncInfo.Len
	(*AuthInfoReq)(nil),             // 2: magma.feg.hlr.AuthInfoReq
	(*AuthInfoAns)(nil),             // 3: magma.feg.hlr.AuthInfoAns
	(*AuthInfoReq_ResyncInfo)(nil),  // 4: magma.feg.hlr.AuthInfoReq.ResyncInfo
	(*AuthInfoAns_UMTSVector)(nil),  // 5: magma.feg.hlr.AuthInfoAns.UMTSVector
}
var file_feg_protos_hlr_hlr_proxy_proto_depIdxs = []int32{
	4, // 0: magma.feg.hlr.AuthInfoReq.resync_info:type_name -> magma.feg.hlr.AuthInfoReq.ResyncInfo
	0, // 1: magma.feg.hlr.AuthInfoAns.error_code:type_name -> magma.feg.hlr.ErrorCode
	5, // 2: magma.feg.hlr.AuthInfoAns.umts_vectors:type_name -> magma.feg.hlr.AuthInfoAns.UMTSVector
	2, // 3: magma.feg.hlr.HlrProxy.AuthInfo:input_type -> magma.feg.hlr.AuthInfoReq
	3, // 4: magma.feg.hlr.HlrProxy.AuthInfo:output_type -> magma.feg.hlr.AuthInfoAns
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_feg_protos_hlr_hlr_proxy_proto_init() }
func file_feg_protos_hlr_hlr_proxy_proto_init() {
	if File_feg_protos_hlr_hlr_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feg_protos_hlr_hlr_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInfoReq); i {
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
		file_feg_protos_hlr_hlr_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInfoAns); i {
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
		file_feg_protos_hlr_hlr_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInfoReq_ResyncInfo); i {
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
		file_feg_protos_hlr_hlr_proxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInfoAns_UMTSVector); i {
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
			RawDescriptor: file_feg_protos_hlr_hlr_proxy_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feg_protos_hlr_hlr_proxy_proto_goTypes,
		DependencyIndexes: file_feg_protos_hlr_hlr_proxy_proto_depIdxs,
		EnumInfos:         file_feg_protos_hlr_hlr_proxy_proto_enumTypes,
		MessageInfos:      file_feg_protos_hlr_hlr_proxy_proto_msgTypes,
	}.Build()
	File_feg_protos_hlr_hlr_proxy_proto = out.File
	file_feg_protos_hlr_hlr_proxy_proto_rawDesc = nil
	file_feg_protos_hlr_hlr_proxy_proto_goTypes = nil
	file_feg_protos_hlr_hlr_proxy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HlrProxyClient is the client API for HlrProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HlrProxyClient interface {
	AuthInfo(ctx context.Context, in *AuthInfoReq, opts ...grpc.CallOption) (*AuthInfoAns, error)
}

type hlrProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewHlrProxyClient(cc grpc.ClientConnInterface) HlrProxyClient {
	return &hlrProxyClient{cc}
}

func (c *hlrProxyClient) AuthInfo(ctx context.Context, in *AuthInfoReq, opts ...grpc.CallOption) (*AuthInfoAns, error) {
	out := new(AuthInfoAns)
	err := c.cc.Invoke(ctx, "/magma.feg.hlr.HlrProxy/AuthInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HlrProxyServer is the server API for HlrProxy service.
type HlrProxyServer interface {
	AuthInfo(context.Context, *AuthInfoReq) (*AuthInfoAns, error)
}

// UnimplementedHlrProxyServer can be embedded to have forward compatible implementations.
type UnimplementedHlrProxyServer struct {
}

func (*UnimplementedHlrProxyServer) AuthInfo(context.Context, *AuthInfoReq) (*AuthInfoAns, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthInfo not implemented")
}

func RegisterHlrProxyServer(s *grpc.Server, srv HlrProxyServer) {
	s.RegisterService(&_HlrProxy_serviceDesc, srv)
}

func _HlrProxy_AuthInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HlrProxyServer).AuthInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.hlr.HlrProxy/AuthInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HlrProxyServer).AuthInfo(ctx, req.(*AuthInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HlrProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.hlr.HlrProxy",
	HandlerType: (*HlrProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthInfo",
			Handler:    _HlrProxy_AuthInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/hlr/hlr_proxy.proto",
}
