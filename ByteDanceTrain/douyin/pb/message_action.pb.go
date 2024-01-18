// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: message_action.proto

package pb

import (
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

type DouyinMessageActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`                              // 用户鉴权token
	ToUserId   *int64  `protobuf:"varint,2,req,name=to_user_id,json=toUserId" json:"to_user_id,omitempty"`     // 对方用户id
	ActionType *int32  `protobuf:"varint,3,req,name=action_type,json=actionType" json:"action_type,omitempty"` // 1-发送消息
	Content    *string `protobuf:"bytes,4,req,name=content" json:"content,omitempty"`                          // 消息内容
}

func (x *DouyinMessageActionRequest) Reset() {
	*x = DouyinMessageActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinMessageActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinMessageActionRequest) ProtoMessage() {}

func (x *DouyinMessageActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinMessageActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinMessageActionRequest) Descriptor() ([]byte, []int) {
	return file_message_action_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinMessageActionRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *DouyinMessageActionRequest) GetToUserId() int64 {
	if x != nil && x.ToUserId != nil {
		return *x.ToUserId
	}
	return 0
}

func (x *DouyinMessageActionRequest) GetActionType() int32 {
	if x != nil && x.ActionType != nil {
		return *x.ActionType
	}
	return 0
}

func (x *DouyinMessageActionRequest) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

type DouyinMessageActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
}

func (x *DouyinMessageActionResponse) Reset() {
	*x = DouyinMessageActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinMessageActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinMessageActionResponse) ProtoMessage() {}

func (x *DouyinMessageActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinMessageActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinMessageActionResponse) Descriptor() ([]byte, []int) {
	return file_message_action_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinMessageActionResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *DouyinMessageActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

var File_message_action_proto protoreflect.FileDescriptor

var file_message_action_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x70,
	0x62, 0x22, 0x8e, 0x01, 0x0a, 0x1d, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x08, 0x74,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x60, 0x0a, 0x1e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
}

var (
	file_message_action_proto_rawDescOnce sync.Once
	file_message_action_proto_rawDescData = file_message_action_proto_rawDesc
)

func file_message_action_proto_rawDescGZIP() []byte {
	file_message_action_proto_rawDescOnce.Do(func() {
		file_message_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_action_proto_rawDescData)
	})
	return file_message_action_proto_rawDescData
}

var file_message_action_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_action_proto_goTypes = []interface{}{
	(*DouyinMessageActionRequest)(nil),  // 0: douyin.pb.douyin_message_action_request
	(*DouyinMessageActionResponse)(nil), // 1: douyin.pb.douyin_message_action_response
}
var file_message_action_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_action_proto_init() }
func file_message_action_proto_init() {
	if File_message_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinMessageActionRequest); i {
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
		file_message_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinMessageActionResponse); i {
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
			RawDescriptor: file_message_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_action_proto_goTypes,
		DependencyIndexes: file_message_action_proto_depIdxs,
		MessageInfos:      file_message_action_proto_msgTypes,
	}.Build()
	File_message_action_proto = out.File
	file_message_action_proto_rawDesc = nil
	file_message_action_proto_goTypes = nil
	file_message_action_proto_depIdxs = nil
}
