// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: rpc_update_task.proto

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

type UpdateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	TaskStatus   string `protobuf:"bytes,4,opt,name=task_status,json=taskStatus,proto3" json:"task_status,omitempty"`
	TaskPriority string `protobuf:"bytes,5,opt,name=task_priority,json=taskPriority,proto3" json:"task_priority,omitempty"`
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_update_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_update_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_rpc_update_task_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateTaskRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateTaskRequest) GetTaskStatus() string {
	if x != nil {
		return x.TaskStatus
	}
	return ""
}

func (x *UpdateTaskRequest) GetTaskPriority() string {
	if x != nil {
		return x.TaskPriority
	}
	return ""
}

type UpdateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *UpdateTaskResponse) Reset() {
	*x = UpdateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_update_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskResponse) ProtoMessage() {}

func (x *UpdateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_update_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskResponse.ProtoReflect.Descriptor instead.
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return file_rpc_update_task_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

var File_rpc_update_task_proto protoreflect.FileDescriptor

var file_rpc_update_task_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x32, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x42,
	0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65,
	0x6d, 0x6f, 0x6c, 0x61, 0x32, 0x33, 0x34, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_update_task_proto_rawDescOnce sync.Once
	file_rpc_update_task_proto_rawDescData = file_rpc_update_task_proto_rawDesc
)

func file_rpc_update_task_proto_rawDescGZIP() []byte {
	file_rpc_update_task_proto_rawDescOnce.Do(func() {
		file_rpc_update_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_update_task_proto_rawDescData)
	})
	return file_rpc_update_task_proto_rawDescData
}

var file_rpc_update_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_update_task_proto_goTypes = []interface{}{
	(*UpdateTaskRequest)(nil),  // 0: pb.UpdateTaskRequest
	(*UpdateTaskResponse)(nil), // 1: pb.UpdateTaskResponse
	(*Task)(nil),               // 2: pb.Task
}
var file_rpc_update_task_proto_depIdxs = []int32{
	2, // 0: pb.UpdateTaskResponse.task:type_name -> pb.Task
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_update_task_proto_init() }
func file_rpc_update_task_proto_init() {
	if File_rpc_update_task_proto != nil {
		return
	}
	file_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_update_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskRequest); i {
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
		file_rpc_update_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskResponse); i {
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
			RawDescriptor: file_rpc_update_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_update_task_proto_goTypes,
		DependencyIndexes: file_rpc_update_task_proto_depIdxs,
		MessageInfos:      file_rpc_update_task_proto_msgTypes,
	}.Build()
	File_rpc_update_task_proto = out.File
	file_rpc_update_task_proto_rawDesc = nil
	file_rpc_update_task_proto_goTypes = nil
	file_rpc_update_task_proto_depIdxs = nil
}
