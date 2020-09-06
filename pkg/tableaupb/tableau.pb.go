// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: tableau.proto

package tableaupb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Cardinality of field
type FieldCard int32

const (
	FieldCard_FIELD_CARD_REQUIRED FieldCard = 0 // appears exactly one time
	FieldCard_FIELD_CARD_OPTIONAL FieldCard = 1 // appears zero or one times
	FieldCard_FIELD_CARD_REPEATED FieldCard = 2 // appears zero or more times
)

// Enum value maps for FieldCard.
var (
	FieldCard_name = map[int32]string{
		0: "FIELD_CARD_REQUIRED",
		1: "FIELD_CARD_OPTIONAL",
		2: "FIELD_CARD_REPEATED",
	}
	FieldCard_value = map[string]int32{
		"FIELD_CARD_REQUIRED": 0,
		"FIELD_CARD_OPTIONAL": 1,
		"FIELD_CARD_REPEATED": 2,
	}
)

func (x FieldCard) Enum() *FieldCard {
	p := new(FieldCard)
	*p = x
	return p
}

func (x FieldCard) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldCard) Descriptor() protoreflect.EnumDescriptor {
	return file_tableau_proto_enumTypes[0].Descriptor()
}

func (FieldCard) Type() protoreflect.EnumType {
	return &file_tableau_proto_enumTypes[0]
}

func (x FieldCard) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldCard.Descriptor instead.
func (FieldCard) EnumDescriptor() ([]byte, []int) {
	return file_tableau_proto_rawDescGZIP(), []int{0}
}

type MetaTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workbook  string       `protobuf:"bytes,1,opt,name=workbook,proto3" json:"workbook,omitempty"`
	Worksheet string       `protobuf:"bytes,2,opt,name=worksheet,proto3" json:"worksheet,omitempty"`
	Metarow   int32        `protobuf:"varint,3,opt,name=metarow,proto3" json:"metarow,omitempty"` // [default = 1]; // exact row number of meta info of this tableau
	Descrow   int32        `protobuf:"varint,4,opt,name=descrow,proto3" json:"descrow,omitempty"` // [default = 1]; // exact row number of description of this tableau
	Datarow   int32        `protobuf:"varint,5,opt,name=datarow,proto3" json:"datarow,omitempty"` // [default = 2]; // start row number of data
	Fields    []*MetaField `protobuf:"bytes,10,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *MetaTable) Reset() {
	*x = MetaTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaTable) ProtoMessage() {}

func (x *MetaTable) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaTable.ProtoReflect.Descriptor instead.
func (*MetaTable) Descriptor() ([]byte, []int) {
	return file_tableau_proto_rawDescGZIP(), []int{0}
}

func (x *MetaTable) GetWorkbook() string {
	if x != nil {
		return x.Workbook
	}
	return ""
}

func (x *MetaTable) GetWorksheet() string {
	if x != nil {
		return x.Worksheet
	}
	return ""
}

func (x *MetaTable) GetMetarow() int32 {
	if x != nil {
		return x.Metarow
	}
	return 0
}

func (x *MetaTable) GetDescrow() int32 {
	if x != nil {
		return x.Descrow
	}
	return 0
}

func (x *MetaTable) GetDatarow() int32 {
	if x != nil {
		return x.Datarow
	}
	return 0
}

func (x *MetaTable) GetFields() []*MetaField {
	if x != nil {
		return x.Fields
	}
	return nil
}

type MetaField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Card     FieldCard `protobuf:"varint,1,opt,name=card,proto3,enum=tableau.FieldCard" json:"card,omitempty"`
	Type     FieldType `protobuf:"varint,2,opt,name=type,proto3,enum=FieldType" json:"type,omitempty"`
	Name     string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`          // varaible name
	Caption  string    `protobuf:"bytes,4,opt,name=caption,proto3" json:"caption,omitempty"`    // column name or name prefix
	Required bool      `protobuf:"varint,5,opt,name=required,proto3" json:"required,omitempty"` // [default = true];
	Sep      string    `protobuf:"bytes,6,opt,name=sep,proto3" json:"sep,omitempty"`            // [default = ',']; separator
	Subsep   string    `protobuf:"bytes,7,opt,name=subsep,proto3" json:"subsep,omitempty"`      // [default = ':']; sub separator
	// valid if type is FIELD_TYPE_MESSSAGE
	Fields []*MetaField `protobuf:"bytes,10,rep,name=fields,proto3" json:"fields,omitempty"`
	////////////////////////////////////////////////////////////////////////
	// Simple validators below
	////////////////////////////////////////////////////////////////////////
	// Different meanings:
	// repeated: size range of list or map
	// integer: value range
	// string: count of utf-8 code point
	Min int32 `protobuf:"varint,20,opt,name=min,proto3" json:"min,omitempty"` // min value
	Max int32 `protobuf:"varint,21,opt,name=max,proto3" json:"max,omitempty"` // max value
}

func (x *MetaField) Reset() {
	*x = MetaField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaField) ProtoMessage() {}

func (x *MetaField) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaField.ProtoReflect.Descriptor instead.
func (*MetaField) Descriptor() ([]byte, []int) {
	return file_tableau_proto_rawDescGZIP(), []int{1}
}

func (x *MetaField) GetCard() FieldCard {
	if x != nil {
		return x.Card
	}
	return FieldCard_FIELD_CARD_REQUIRED
}

func (x *MetaField) GetType() FieldType {
	if x != nil {
		return x.Type
	}
	return FieldType_FIELD_TYPE_DEFAULT
}

func (x *MetaField) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetaField) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *MetaField) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *MetaField) GetSep() string {
	if x != nil {
		return x.Sep
	}
	return ""
}

func (x *MetaField) GetSubsep() string {
	if x != nil {
		return x.Subsep
	}
	return ""
}

func (x *MetaField) GetFields() []*MetaField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *MetaField) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *MetaField) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

var File_tableau_proto protoreflect.FileDescriptor

var file_tableau_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x1a, 0x15, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61,
	0x75, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbf, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x68, 0x65, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x61, 0x72,
	0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x61, 0x72, 0x6f,
	0x77, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x61, 0x74, 0x61, 0x72, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x61,
	0x74, 0x61, 0x72, 0x6f, 0x77, 0x12, 0x2a, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x22, 0x97, 0x02, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x26, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x72,
	0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x65, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x73, 0x65, 0x70, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x73, 0x65, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x2a, 0x56, 0x0a, 0x09, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x5f,
	0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x57, 0x65, 0x6e, 0x63, 0x68, 0x79, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tableau_proto_rawDescOnce sync.Once
	file_tableau_proto_rawDescData = file_tableau_proto_rawDesc
)

func file_tableau_proto_rawDescGZIP() []byte {
	file_tableau_proto_rawDescOnce.Do(func() {
		file_tableau_proto_rawDescData = protoimpl.X.CompressGZIP(file_tableau_proto_rawDescData)
	})
	return file_tableau_proto_rawDescData
}

var file_tableau_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tableau_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tableau_proto_goTypes = []interface{}{
	(FieldCard)(0),    // 0: tableau.FieldCard
	(*MetaTable)(nil), // 1: tableau.MetaTable
	(*MetaField)(nil), // 2: tableau.MetaField
	(FieldType)(0),    // 3: FieldType
}
var file_tableau_proto_depIdxs = []int32{
	2, // 0: tableau.MetaTable.fields:type_name -> tableau.MetaField
	0, // 1: tableau.MetaField.card:type_name -> tableau.FieldCard
	3, // 2: tableau.MetaField.type:type_name -> FieldType
	2, // 3: tableau.MetaField.fields:type_name -> tableau.MetaField
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tableau_proto_init() }
func file_tableau_proto_init() {
	if File_tableau_proto != nil {
		return
	}
	file_tableau_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tableau_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaTable); i {
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
		file_tableau_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaField); i {
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
			RawDescriptor: file_tableau_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tableau_proto_goTypes,
		DependencyIndexes: file_tableau_proto_depIdxs,
		EnumInfos:         file_tableau_proto_enumTypes,
		MessageInfos:      file_tableau_proto_msgTypes,
	}.Build()
	File_tableau_proto = out.File
	file_tableau_proto_rawDesc = nil
	file_tableau_proto_goTypes = nil
	file_tableau_proto_depIdxs = nil
}
