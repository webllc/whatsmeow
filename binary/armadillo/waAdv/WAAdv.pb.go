// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: waAdv/WAAdv.proto

package waAdv

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ADVEncryptionType int32

const (
	ADVEncryptionType_E2EE   ADVEncryptionType = 0
	ADVEncryptionType_HOSTED ADVEncryptionType = 1
)

// Enum value maps for ADVEncryptionType.
var (
	ADVEncryptionType_name = map[int32]string{
		0: "E2EE",
		1: "HOSTED",
	}
	ADVEncryptionType_value = map[string]int32{
		"E2EE":   0,
		"HOSTED": 1,
	}
)

func (x ADVEncryptionType) Enum() *ADVEncryptionType {
	p := new(ADVEncryptionType)
	*p = x
	return p
}

func (x ADVEncryptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ADVEncryptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_waAdv_WAAdv_proto_enumTypes[0].Descriptor()
}

func (ADVEncryptionType) Type() protoreflect.EnumType {
	return &file_waAdv_WAAdv_proto_enumTypes[0]
}

func (x ADVEncryptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ADVEncryptionType.Descriptor instead.
func (ADVEncryptionType) EnumDescriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{0}
}

type ADVKeyIndexList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawID        uint32            `protobuf:"varint,1,opt,name=rawID,proto3" json:"rawID,omitempty"`
	Timestamp    uint64            `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CurrentIndex uint32            `protobuf:"varint,3,opt,name=currentIndex,proto3" json:"currentIndex,omitempty"`
	ValidIndexes []uint32          `protobuf:"varint,4,rep,packed,name=validIndexes,proto3" json:"validIndexes,omitempty"`
	AccountType  ADVEncryptionType `protobuf:"varint,5,opt,name=accountType,proto3,enum=WAAdv.ADVEncryptionType" json:"accountType,omitempty"`
}

func (x *ADVKeyIndexList) Reset() {
	*x = ADVKeyIndexList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waAdv_WAAdv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADVKeyIndexList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADVKeyIndexList) ProtoMessage() {}

func (x *ADVKeyIndexList) ProtoReflect() protoreflect.Message {
	mi := &file_waAdv_WAAdv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADVKeyIndexList.ProtoReflect.Descriptor instead.
func (*ADVKeyIndexList) Descriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{0}
}

func (x *ADVKeyIndexList) GetRawID() uint32 {
	if x != nil {
		return x.RawID
	}
	return 0
}

func (x *ADVKeyIndexList) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ADVKeyIndexList) GetCurrentIndex() uint32 {
	if x != nil {
		return x.CurrentIndex
	}
	return 0
}

func (x *ADVKeyIndexList) GetValidIndexes() []uint32 {
	if x != nil {
		return x.ValidIndexes
	}
	return nil
}

func (x *ADVKeyIndexList) GetAccountType() ADVEncryptionType {
	if x != nil {
		return x.AccountType
	}
	return ADVEncryptionType_E2EE
}

type ADVSignedKeyIndexList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details             []byte `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	AccountSignature    []byte `protobuf:"bytes,2,opt,name=accountSignature,proto3" json:"accountSignature,omitempty"`
	AccountSignatureKey []byte `protobuf:"bytes,3,opt,name=accountSignatureKey,proto3" json:"accountSignatureKey,omitempty"`
}

func (x *ADVSignedKeyIndexList) Reset() {
	*x = ADVSignedKeyIndexList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waAdv_WAAdv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADVSignedKeyIndexList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADVSignedKeyIndexList) ProtoMessage() {}

func (x *ADVSignedKeyIndexList) ProtoReflect() protoreflect.Message {
	mi := &file_waAdv_WAAdv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADVSignedKeyIndexList.ProtoReflect.Descriptor instead.
func (*ADVSignedKeyIndexList) Descriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{1}
}

func (x *ADVSignedKeyIndexList) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ADVSignedKeyIndexList) GetAccountSignature() []byte {
	if x != nil {
		return x.AccountSignature
	}
	return nil
}

func (x *ADVSignedKeyIndexList) GetAccountSignatureKey() []byte {
	if x != nil {
		return x.AccountSignatureKey
	}
	return nil
}

type ADVDeviceIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawID       uint32            `protobuf:"varint,1,opt,name=rawID,proto3" json:"rawID,omitempty"`
	Timestamp   uint64            `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	KeyIndex    uint32            `protobuf:"varint,3,opt,name=keyIndex,proto3" json:"keyIndex,omitempty"`
	AccountType ADVEncryptionType `protobuf:"varint,4,opt,name=accountType,proto3,enum=WAAdv.ADVEncryptionType" json:"accountType,omitempty"`
	DeviceType  ADVEncryptionType `protobuf:"varint,5,opt,name=deviceType,proto3,enum=WAAdv.ADVEncryptionType" json:"deviceType,omitempty"`
}

func (x *ADVDeviceIdentity) Reset() {
	*x = ADVDeviceIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waAdv_WAAdv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADVDeviceIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADVDeviceIdentity) ProtoMessage() {}

func (x *ADVDeviceIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_waAdv_WAAdv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADVDeviceIdentity.ProtoReflect.Descriptor instead.
func (*ADVDeviceIdentity) Descriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{2}
}

func (x *ADVDeviceIdentity) GetRawID() uint32 {
	if x != nil {
		return x.RawID
	}
	return 0
}

func (x *ADVDeviceIdentity) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ADVDeviceIdentity) GetKeyIndex() uint32 {
	if x != nil {
		return x.KeyIndex
	}
	return 0
}

func (x *ADVDeviceIdentity) GetAccountType() ADVEncryptionType {
	if x != nil {
		return x.AccountType
	}
	return ADVEncryptionType_E2EE
}

func (x *ADVDeviceIdentity) GetDeviceType() ADVEncryptionType {
	if x != nil {
		return x.DeviceType
	}
	return ADVEncryptionType_E2EE
}

type ADVSignedDeviceIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details             []byte `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	AccountSignatureKey []byte `protobuf:"bytes,2,opt,name=accountSignatureKey,proto3" json:"accountSignatureKey,omitempty"`
	AccountSignature    []byte `protobuf:"bytes,3,opt,name=accountSignature,proto3" json:"accountSignature,omitempty"`
	DeviceSignature     []byte `protobuf:"bytes,4,opt,name=deviceSignature,proto3" json:"deviceSignature,omitempty"`
}

func (x *ADVSignedDeviceIdentity) Reset() {
	*x = ADVSignedDeviceIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waAdv_WAAdv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADVSignedDeviceIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADVSignedDeviceIdentity) ProtoMessage() {}

func (x *ADVSignedDeviceIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_waAdv_WAAdv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADVSignedDeviceIdentity.ProtoReflect.Descriptor instead.
func (*ADVSignedDeviceIdentity) Descriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{3}
}

func (x *ADVSignedDeviceIdentity) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ADVSignedDeviceIdentity) GetAccountSignatureKey() []byte {
	if x != nil {
		return x.AccountSignatureKey
	}
	return nil
}

func (x *ADVSignedDeviceIdentity) GetAccountSignature() []byte {
	if x != nil {
		return x.AccountSignature
	}
	return nil
}

func (x *ADVSignedDeviceIdentity) GetDeviceSignature() []byte {
	if x != nil {
		return x.DeviceSignature
	}
	return nil
}

type ADVSignedDeviceIdentityHMAC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details     []byte            `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	HMAC        []byte            `protobuf:"bytes,2,opt,name=HMAC,proto3" json:"HMAC,omitempty"`
	AccountType ADVEncryptionType `protobuf:"varint,3,opt,name=accountType,proto3,enum=WAAdv.ADVEncryptionType" json:"accountType,omitempty"`
}

func (x *ADVSignedDeviceIdentityHMAC) Reset() {
	*x = ADVSignedDeviceIdentityHMAC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waAdv_WAAdv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADVSignedDeviceIdentityHMAC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADVSignedDeviceIdentityHMAC) ProtoMessage() {}

func (x *ADVSignedDeviceIdentityHMAC) ProtoReflect() protoreflect.Message {
	mi := &file_waAdv_WAAdv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADVSignedDeviceIdentityHMAC.ProtoReflect.Descriptor instead.
func (*ADVSignedDeviceIdentityHMAC) Descriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{4}
}

func (x *ADVSignedDeviceIdentityHMAC) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ADVSignedDeviceIdentityHMAC) GetHMAC() []byte {
	if x != nil {
		return x.HMAC
	}
	return nil
}

func (x *ADVSignedDeviceIdentityHMAC) GetAccountType() ADVEncryptionType {
	if x != nil {
		return x.AccountType
	}
	return ADVEncryptionType_E2EE
}

var File_waAdv_WAAdv_proto protoreflect.FileDescriptor

//go:embed WAAdv.pb.raw
var file_waAdv_WAAdv_proto_rawDesc []byte

var (
	file_waAdv_WAAdv_proto_rawDescOnce sync.Once
	file_waAdv_WAAdv_proto_rawDescData = file_waAdv_WAAdv_proto_rawDesc
)

func file_waAdv_WAAdv_proto_rawDescGZIP() []byte {
	file_waAdv_WAAdv_proto_rawDescOnce.Do(func() {
		file_waAdv_WAAdv_proto_rawDescData = protoimpl.X.CompressGZIP(file_waAdv_WAAdv_proto_rawDescData)
	})
	return file_waAdv_WAAdv_proto_rawDescData
}

var file_waAdv_WAAdv_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_waAdv_WAAdv_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_waAdv_WAAdv_proto_goTypes = []interface{}{
	(ADVEncryptionType)(0),              // 0: WAAdv.ADVEncryptionType
	(*ADVKeyIndexList)(nil),             // 1: WAAdv.ADVKeyIndexList
	(*ADVSignedKeyIndexList)(nil),       // 2: WAAdv.ADVSignedKeyIndexList
	(*ADVDeviceIdentity)(nil),           // 3: WAAdv.ADVDeviceIdentity
	(*ADVSignedDeviceIdentity)(nil),     // 4: WAAdv.ADVSignedDeviceIdentity
	(*ADVSignedDeviceIdentityHMAC)(nil), // 5: WAAdv.ADVSignedDeviceIdentityHMAC
}
var file_waAdv_WAAdv_proto_depIdxs = []int32{
	0, // 0: WAAdv.ADVKeyIndexList.accountType:type_name -> WAAdv.ADVEncryptionType
	0, // 1: WAAdv.ADVDeviceIdentity.accountType:type_name -> WAAdv.ADVEncryptionType
	0, // 2: WAAdv.ADVDeviceIdentity.deviceType:type_name -> WAAdv.ADVEncryptionType
	0, // 3: WAAdv.ADVSignedDeviceIdentityHMAC.accountType:type_name -> WAAdv.ADVEncryptionType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_waAdv_WAAdv_proto_init() }
func file_waAdv_WAAdv_proto_init() {
	if File_waAdv_WAAdv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waAdv_WAAdv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADVKeyIndexList); i {
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
		file_waAdv_WAAdv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADVSignedKeyIndexList); i {
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
		file_waAdv_WAAdv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADVDeviceIdentity); i {
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
		file_waAdv_WAAdv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADVSignedDeviceIdentity); i {
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
		file_waAdv_WAAdv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADVSignedDeviceIdentityHMAC); i {
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
			RawDescriptor: file_waAdv_WAAdv_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waAdv_WAAdv_proto_goTypes,
		DependencyIndexes: file_waAdv_WAAdv_proto_depIdxs,
		EnumInfos:         file_waAdv_WAAdv_proto_enumTypes,
		MessageInfos:      file_waAdv_WAAdv_proto_msgTypes,
	}.Build()
	File_waAdv_WAAdv_proto = out.File
	file_waAdv_WAAdv_proto_rawDesc = nil
	file_waAdv_WAAdv_proto_goTypes = nil
	file_waAdv_WAAdv_proto_depIdxs = nil
}
