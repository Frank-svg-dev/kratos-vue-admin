// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: api/kva-user/auth.proto

package kva_user

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UserLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserLoginReq) Reset() {
	*x = UserLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kva_user_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginReq) ProtoMessage() {}

func (x *UserLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_kva_user_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginReq.ProtoReflect.Descriptor instead.
func (*UserLoginReq) Descriptor() ([]byte, []int) {
	return file_api_kva_user_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserLoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpiresAt string `protobuf:"bytes,2,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
}

func (x *UserLoginRes) Reset() {
	*x = UserLoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kva_user_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRes) ProtoMessage() {}

func (x *UserLoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_kva_user_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRes.ProtoReflect.Descriptor instead.
func (*UserLoginRes) Descriptor() ([]byte, []int) {
	return file_api_kva_user_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserLoginRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserLoginRes) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

var File_api_kva_user_auth_proto protoreflect.FileDescriptor

var file_api_kva_user_auth_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x76, 0x61, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b, 0x76, 0x61, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x28, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x42, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x32, 0x57, 0x0a, 0x04, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x4f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x16, 0x2e, 0x6b, 0x76, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x6b, 0x76, 0x61, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x22,
	0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x06, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a,
	0x01, 0x2a, 0x42, 0x30, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76, 0x61, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x50, 0x01, 0x5a, 0x1e, 0x6b, 0x76, 0x61, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6b, 0x76, 0x61, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x6b, 0x76, 0x61, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_kva_user_auth_proto_rawDescOnce sync.Once
	file_api_kva_user_auth_proto_rawDescData = file_api_kva_user_auth_proto_rawDesc
)

func file_api_kva_user_auth_proto_rawDescGZIP() []byte {
	file_api_kva_user_auth_proto_rawDescOnce.Do(func() {
		file_api_kva_user_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_kva_user_auth_proto_rawDescData)
	})
	return file_api_kva_user_auth_proto_rawDescData
}

var file_api_kva_user_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_kva_user_auth_proto_goTypes = []interface{}{
	(*UserLoginReq)(nil), // 0: kva_user.UserLoginReq
	(*UserLoginRes)(nil), // 1: kva_user.UserLoginRes
}
var file_api_kva_user_auth_proto_depIdxs = []int32{
	0, // 0: kva_user.Auth.CreateAuth:input_type -> kva_user.UserLoginReq
	1, // 1: kva_user.Auth.CreateAuth:output_type -> kva_user.UserLoginRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_kva_user_auth_proto_init() }
func file_api_kva_user_auth_proto_init() {
	if File_api_kva_user_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_kva_user_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginReq); i {
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
		file_api_kva_user_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginRes); i {
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
			RawDescriptor: file_api_kva_user_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_kva_user_auth_proto_goTypes,
		DependencyIndexes: file_api_kva_user_auth_proto_depIdxs,
		MessageInfos:      file_api_kva_user_auth_proto_msgTypes,
	}.Build()
	File_api_kva_user_auth_proto = out.File
	file_api_kva_user_auth_proto_rawDesc = nil
	file_api_kva_user_auth_proto_goTypes = nil
	file_api_kva_user_auth_proto_depIdxs = nil
}
