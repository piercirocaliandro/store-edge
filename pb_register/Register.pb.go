// this file contains the definition of the Register interface, using Protocol Buffer syntax. This interface will be implemented
// by the server node
//
// Compile with: protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb_register/Register.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: pb_register/Register.proto

package pb_register

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

// A storage record is an element in which a value is associated to a specific key
type StorageServerMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddr string `protobuf:"bytes,1,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	Port   int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *StorageServerMsg) Reset() {
	*x = StorageServerMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_register_Register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageServerMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageServerMsg) ProtoMessage() {}

func (x *StorageServerMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pb_register_Register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageServerMsg.ProtoReflect.Descriptor instead.
func (*StorageServerMsg) Descriptor() ([]byte, []int) {
	return file_pb_register_Register_proto_rawDescGZIP(), []int{0}
}

func (x *StorageServerMsg) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *StorageServerMsg) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type ListOfStorageServerMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddr []string `protobuf:"bytes,1,rep,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	Port   []int32  `protobuf:"varint,2,rep,packed,name=port,proto3" json:"port,omitempty"`
}

func (x *ListOfStorageServerMsg) Reset() {
	*x = ListOfStorageServerMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_register_Register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfStorageServerMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfStorageServerMsg) ProtoMessage() {}

func (x *ListOfStorageServerMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pb_register_Register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfStorageServerMsg.ProtoReflect.Descriptor instead.
func (*ListOfStorageServerMsg) Descriptor() ([]byte, []int) {
	return file_pb_register_Register_proto_rawDescGZIP(), []int{1}
}

func (x *ListOfStorageServerMsg) GetIpAddr() []string {
	if x != nil {
		return x.IpAddr
	}
	return nil
}

func (x *ListOfStorageServerMsg) GetPort() []int32 {
	if x != nil {
		return x.Port
	}
	return nil
}

type ClientMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddr string `protobuf:"bytes,1,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	Port   int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ClientMsg) Reset() {
	*x = ClientMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_register_Register_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMsg) ProtoMessage() {}

func (x *ClientMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pb_register_Register_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMsg.ProtoReflect.Descriptor instead.
func (*ClientMsg) Descriptor() ([]byte, []int) {
	return file_pb_register_Register_proto_rawDescGZIP(), []int{2}
}

func (x *ClientMsg) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *ClientMsg) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_pb_register_Register_proto protoreflect.FileDescriptor

var file_pb_register_Register_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x62, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x70,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x45, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x66, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73,
	0x67, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x38,
	0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x32, 0xfc, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12,
	0x13, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x73, 0x67, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x66, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x09, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x70, 0x62, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_register_Register_proto_rawDescOnce sync.Once
	file_pb_register_Register_proto_rawDescData = file_pb_register_Register_proto_rawDesc
)

func file_pb_register_Register_proto_rawDescGZIP() []byte {
	file_pb_register_Register_proto_rawDescOnce.Do(func() {
		file_pb_register_Register_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_register_Register_proto_rawDescData)
	})
	return file_pb_register_Register_proto_rawDescData
}

var file_pb_register_Register_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_register_Register_proto_goTypes = []interface{}{
	(*StorageServerMsg)(nil),       // 0: register.StorageServerMsg
	(*ListOfStorageServerMsg)(nil), // 1: register.ListOfStorageServerMsg
	(*ClientMsg)(nil),              // 2: register.ClientMsg
}
var file_pb_register_Register_proto_depIdxs = []int32{
	2, // 0: register.Register.GetCloseStorageServers:input_type -> register.ClientMsg
	0, // 1: register.Register.GetListOfPeers:input_type -> register.StorageServerMsg
	0, // 2: register.Register.Heartbeat:input_type -> register.StorageServerMsg
	1, // 3: register.Register.GetCloseStorageServers:output_type -> register.ListOfStorageServerMsg
	1, // 4: register.Register.GetListOfPeers:output_type -> register.ListOfStorageServerMsg
	1, // 5: register.Register.Heartbeat:output_type -> register.ListOfStorageServerMsg
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_register_Register_proto_init() }
func file_pb_register_Register_proto_init() {
	if File_pb_register_Register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_register_Register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageServerMsg); i {
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
		file_pb_register_Register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfStorageServerMsg); i {
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
		file_pb_register_Register_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMsg); i {
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
			RawDescriptor: file_pb_register_Register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_register_Register_proto_goTypes,
		DependencyIndexes: file_pb_register_Register_proto_depIdxs,
		MessageInfos:      file_pb_register_Register_proto_msgTypes,
	}.Build()
	File_pb_register_Register_proto = out.File
	file_pb_register_Register_proto_rawDesc = nil
	file_pb_register_Register_proto_goTypes = nil
	file_pb_register_Register_proto_depIdxs = nil
}
