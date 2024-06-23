// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: internal/infrastructure/grpc/scheduleGrpc/schedule_service.proto

package scheduleGrpc

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

type CreateScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From           string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Type           string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	CronExpression string `protobuf:"bytes,3,opt,name=cron_expression,json=cronExpression,proto3" json:"cron_expression,omitempty"`
	Deadline       string `protobuf:"bytes,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
	ReplyOn        string `protobuf:"bytes,5,opt,name=reply_on,json=replyOn,proto3" json:"reply_on,omitempty"`
	XApiKey        string `protobuf:"bytes,6,opt,name=x_api_key,json=xApiKey,proto3" json:"x_api_key,omitempty"`
	Data           string `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateScheduleRequest) Reset() {
	*x = CreateScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScheduleRequest) ProtoMessage() {}

func (x *CreateScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScheduleRequest.ProtoReflect.Descriptor instead.
func (*CreateScheduleRequest) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateScheduleRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CreateScheduleRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateScheduleRequest) GetCronExpression() string {
	if x != nil {
		return x.CronExpression
	}
	return ""
}

func (x *CreateScheduleRequest) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *CreateScheduleRequest) GetReplyOn() string {
	if x != nil {
		return x.ReplyOn
	}
	return ""
}

func (x *CreateScheduleRequest) GetXApiKey() string {
	if x != nil {
		return x.XApiKey
	}
	return ""
}

func (x *CreateScheduleRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type CreateScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
}

func (x *CreateScheduleResponse) Reset() {
	*x = CreateScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScheduleResponse) ProtoMessage() {}

func (x *CreateScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScheduleResponse.ProtoReflect.Descriptor instead.
func (*CreateScheduleResponse) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateScheduleResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

var File_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto protoreflect.FileDescriptor

var file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x72, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x5f, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x4f, 0x6e, 0x12, 0x1a, 0x0a, 0x09, 0x78, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x78, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x36, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x56, 0x0a, 0x0f,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x43, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x47, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDescOnce sync.Once
	file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDescData = file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDesc
)

func file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDescGZIP() []byte {
	file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDescOnce.Do(func() {
		file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDescData)
	})
	return file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDescData
}

var file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_goTypes = []interface{}{
	(*CreateScheduleRequest)(nil),  // 0: CreateScheduleRequest
	(*CreateScheduleResponse)(nil), // 1: CreateScheduleResponse
}
var file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_depIdxs = []int32{
	0, // 0: ScheduleService.CreateSchedule:input_type -> CreateScheduleRequest
	1, // 1: ScheduleService.CreateSchedule:output_type -> CreateScheduleResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_init() }
func file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_init() {
	if File_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScheduleRequest); i {
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
		file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScheduleResponse); i {
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
			RawDescriptor: file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_goTypes,
		DependencyIndexes: file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_depIdxs,
		MessageInfos:      file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_msgTypes,
	}.Build()
	File_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto = out.File
	file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_rawDesc = nil
	file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_goTypes = nil
	file_internal_infrastructure_grpc_scheduleGrpc_schedule_service_proto_depIdxs = nil
}
