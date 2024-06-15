// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: proto/device.proto

package proto

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

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceType string `protobuf:"bytes,1,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	DevicePath string `protobuf:"bytes,2,opt,name=device_path,json=devicePath,proto3" json:"device_path,omitempty"`
	VendorId   string `protobuf:"bytes,3,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	ProductId  string `protobuf:"bytes,4,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_proto_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *Device) GetDevicePath() string {
	if x != nil {
		return x.DevicePath
	}
	return ""
}

func (x *Device) GetVendorId() string {
	if x != nil {
		return x.VendorId
	}
	return ""
}

func (x *Device) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type DeviceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *DeviceList) Reset() {
	*x = DeviceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceList) ProtoMessage() {}

func (x *DeviceList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceList.ProtoReflect.Descriptor instead.
func (*DeviceList) Descriptor() ([]byte, []int) {
	return file_proto_device_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceList) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type GetDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDevicesRequest) Reset() {
	*x = GetDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDevicesRequest) ProtoMessage() {}

func (x *GetDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDevicesRequest.ProtoReflect.Descriptor instead.
func (*GetDevicesRequest) Descriptor() ([]byte, []int) {
	return file_proto_device_proto_rawDescGZIP(), []int{2}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_device_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_device_proto protoreflect.FileDescriptor

var file_proto_device_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x06,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x4a, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_device_proto_rawDescOnce sync.Once
	file_proto_device_proto_rawDescData = file_proto_device_proto_rawDesc
)

func file_proto_device_proto_rawDescGZIP() []byte {
	file_proto_device_proto_rawDescOnce.Do(func() {
		file_proto_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_device_proto_rawDescData)
	})
	return file_proto_device_proto_rawDescData
}

var file_proto_device_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_device_proto_goTypes = []interface{}{
	(*Device)(nil),            // 0: proto.Device
	(*DeviceList)(nil),        // 1: proto.DeviceList
	(*GetDevicesRequest)(nil), // 2: proto.GetDevicesRequest
	(*Error)(nil),             // 3: proto.Error
}
var file_proto_device_proto_depIdxs = []int32{
	0, // 0: proto.DeviceList.devices:type_name -> proto.Device
	2, // 1: proto.DeviceService.GetDevices:input_type -> proto.GetDevicesRequest
	1, // 2: proto.DeviceService.GetDevices:output_type -> proto.DeviceList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_device_proto_init() }
func file_proto_device_proto_init() {
	if File_proto_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_proto_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceList); i {
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
		file_proto_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDevicesRequest); i {
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
		file_proto_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_proto_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_device_proto_goTypes,
		DependencyIndexes: file_proto_device_proto_depIdxs,
		MessageInfos:      file_proto_device_proto_msgTypes,
	}.Build()
	File_proto_device_proto = out.File
	file_proto_device_proto_rawDesc = nil
	file_proto_device_proto_goTypes = nil
	file_proto_device_proto_depIdxs = nil
}
