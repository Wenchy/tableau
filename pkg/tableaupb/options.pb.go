// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: options.proto

package tableaupb

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// type of the message field.
type FieldType int32

const (
	// auto deduced protobuf types
	FieldType_FIELD_TYPE_DEFAULT FieldType = 0
	//////////////////////////////
	/////Build-in Types///////////
	//////////////////////////////
	// interger
	FieldType_FIELD_TYPE_INT32  FieldType = 1
	FieldType_FIELD_TYPE_UINT32 FieldType = 2
	FieldType_FIELD_TYPE_INT64  FieldType = 3
	FieldType_FIELD_TYPE_UINT64 FieldType = 4
	// floating-point number
	FieldType_FIELD_TYPE_DOUBLE FieldType = 5
	FieldType_FIELD_TYPE_FLOAT  FieldType = 6
	// bool
	FieldType_FIELD_TYPE_BOOL FieldType = 7
	// string
	FieldType_FIELD_TYPE_STRING FieldType = 8
	////////////////////////
	/////Message Type//////
	////////////////////////
	FieldType_FIELD_TYPE_MESSSAGE FieldType = 10
	////////////////////////
	/////Extended Types/////
	////////////////////////
	// time
	FieldType_FIELD_TYPE_DATE     FieldType = 21 // format: "yyyy-MM-dd"
	FieldType_FIELD_TYPE_TIME     FieldType = 22 // format: "HH:mm:ss"
	FieldType_FIELD_TYPE_DATETIME FieldType = 23 // format: "yyyy-MM-dd HH:mm:ss"
	// list in a cell:
	// - the list **item** must be **built-in** type
	// - format: ',' separated items
	FieldType_FIELD_TYPE_CELL_LIST FieldType = 24
	// map in a cell:
	// - both the **key** and **value** must be **built-in** type
	// - format: key-value pairs is separated by ',', and
	//           key and value is separated by ':'
	FieldType_FIELD_TYPE_CELL_MAP FieldType = 25
	// message in a cell
	FieldType_FIELD_TYPE_CELL_MESSAGE FieldType = 26
)

// Enum value maps for FieldType.
var (
	FieldType_name = map[int32]string{
		0:  "FIELD_TYPE_DEFAULT",
		1:  "FIELD_TYPE_INT32",
		2:  "FIELD_TYPE_UINT32",
		3:  "FIELD_TYPE_INT64",
		4:  "FIELD_TYPE_UINT64",
		5:  "FIELD_TYPE_DOUBLE",
		6:  "FIELD_TYPE_FLOAT",
		7:  "FIELD_TYPE_BOOL",
		8:  "FIELD_TYPE_STRING",
		10: "FIELD_TYPE_MESSSAGE",
		21: "FIELD_TYPE_DATE",
		22: "FIELD_TYPE_TIME",
		23: "FIELD_TYPE_DATETIME",
		24: "FIELD_TYPE_CELL_LIST",
		25: "FIELD_TYPE_CELL_MAP",
		26: "FIELD_TYPE_CELL_MESSAGE",
	}
	FieldType_value = map[string]int32{
		"FIELD_TYPE_DEFAULT":      0,
		"FIELD_TYPE_INT32":        1,
		"FIELD_TYPE_UINT32":       2,
		"FIELD_TYPE_INT64":        3,
		"FIELD_TYPE_UINT64":       4,
		"FIELD_TYPE_DOUBLE":       5,
		"FIELD_TYPE_FLOAT":        6,
		"FIELD_TYPE_BOOL":         7,
		"FIELD_TYPE_STRING":       8,
		"FIELD_TYPE_MESSSAGE":     10,
		"FIELD_TYPE_DATE":         21,
		"FIELD_TYPE_TIME":         22,
		"FIELD_TYPE_DATETIME":     23,
		"FIELD_TYPE_CELL_LIST":    24,
		"FIELD_TYPE_CELL_MAP":     25,
		"FIELD_TYPE_CELL_MESSAGE": 26,
	}
)

func (x FieldType) Enum() *FieldType {
	p := new(FieldType)
	*p = x
	return p
}

func (x FieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_options_proto_enumTypes[0].Descriptor()
}

func (FieldType) Type() protoreflect.EnumType {
	return &file_options_proto_enumTypes[0]
}

func (x FieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldType.Descriptor instead.
func (FieldType) EnumDescriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{0}
}

// layout of composite types, such as list and map.
type CompositeLayout int32

const (
	CompositeLayout_COMPOSITE_LAYOUT_DEFAULT    CompositeLayout = 0 // default direction: vertical for map, horizontal for list
	CompositeLayout_COMPOSITE_LAYOUT_VERTICAL   CompositeLayout = 1 // vertical direction
	CompositeLayout_COMPOSITE_LAYOUT_HORIZONTAL CompositeLayout = 2 // horizontal direction
)

// Enum value maps for CompositeLayout.
var (
	CompositeLayout_name = map[int32]string{
		0: "COMPOSITE_LAYOUT_DEFAULT",
		1: "COMPOSITE_LAYOUT_VERTICAL",
		2: "COMPOSITE_LAYOUT_HORIZONTAL",
	}
	CompositeLayout_value = map[string]int32{
		"COMPOSITE_LAYOUT_DEFAULT":    0,
		"COMPOSITE_LAYOUT_VERTICAL":   1,
		"COMPOSITE_LAYOUT_HORIZONTAL": 2,
	}
)

func (x CompositeLayout) Enum() *CompositeLayout {
	p := new(CompositeLayout)
	*p = x
	return p
}

func (x CompositeLayout) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompositeLayout) Descriptor() protoreflect.EnumDescriptor {
	return file_options_proto_enumTypes[1].Descriptor()
}

func (CompositeLayout) Type() protoreflect.EnumType {
	return &file_options_proto_enumTypes[1]
}

func (x CompositeLayout) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompositeLayout.Descriptor instead.
func (CompositeLayout) EnumDescriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{1}
}

var file_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51001,
		Name:          "workbook",
		Tag:           "bytes,51001,opt,name=workbook",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         52001,
		Name:          "worksheet",
		Tag:           "bytes,52001,opt,name=worksheet",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         52002,
		Name:          "captrow",
		Tag:           "varint,52002,opt,name=captrow",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         52003,
		Name:          "descrow",
		Tag:           "varint,52003,opt,name=descrow",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         52004,
		Name:          "datarow",
		Tag:           "varint,52004,opt,name=datarow",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         52005,
		Name:          "transpose",
		Tag:           "varint,52005,opt,name=transpose",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         53001,
		Name:          "caption",
		Tag:           "bytes,53001,opt,name=caption",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FieldType)(nil),
		Field:         53002,
		Name:          "type",
		Tag:           "varint,53002,opt,name=type,enum=FieldType",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         53003,
		Name:          "key",
		Tag:           "bytes,53003,opt,name=key",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*CompositeLayout)(nil),
		Field:         53004,
		Name:          "layout",
		Tag:           "varint,53004,opt,name=layout,enum=CompositeLayout",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         53005,
		Name:          "sep",
		Tag:           "bytes,53005,opt,name=sep",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         53006,
		Name:          "subsep",
		Tag:           "bytes,53006,opt,name=subsep",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         53101,
		Name:          "min",
		Tag:           "varint,53101,opt,name=min",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         53102,
		Name:          "max",
		Tag:           "varint,53102,opt,name=max",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         53103,
		Name:          "range",
		Tag:           "bytes,53103,opt,name=range",
		Filename:      "options.proto",
	},
}

// Extension fields to descriptor.FileOptions.
var (
	// optional string workbook = 51001;
	E_Workbook = &file_options_proto_extTypes[0]
)

// Extension fields to descriptor.MessageOptions.
var (
	// optional string worksheet = 52001;
	E_Worksheet = &file_options_proto_extTypes[1]
	// optional int32 captrow = 52002;
	E_Captrow = &file_options_proto_extTypes[2] // [default = 1]; // exact row number of caption of the worksheet
	// optional int32 descrow = 52003;
	E_Descrow = &file_options_proto_extTypes[3] // [default = 1]; // exact row number of description of the wooksheet
	// optional int32 datarow = 52004;
	E_Datarow = &file_options_proto_extTypes[4] // [default = 2]; // start row number of data
	// optional bool transpose = 52005;
	E_Transpose = &file_options_proto_extTypes[5] // [default = false]; // interchange the rows and columns
)

// Extension fields to descriptor.FieldOptions.
var (
	// column caption: the word `caption` is self-explanatory to denote the metarow.
	// It means `caption` if field's type is scalar,
	// or `caption prefix` if field's type is composite such as List and Map.
	//
	// optional string caption = 53001;
	E_Caption = &file_options_proto_extTypes[6]
	// tableau field type
	//
	// optional FieldType type = 53002;
	E_Type = &file_options_proto_extTypes[7] // [default = FIELD_TYPE_DEFAULT];
	// optional string key = 53003;
	E_Key = &file_options_proto_extTypes[8]
	// optional CompositeLayout layout = 53004;
	E_Layout = &file_options_proto_extTypes[9] // [default = COMPOSITE_LAYOUT_DEFAULT];
	// optional string sep = 53005;
	E_Sep = &file_options_proto_extTypes[10] // [default = ',']; separator
	// optional string subsep = 53006;
	E_Subsep = &file_options_proto_extTypes[11] // [default = ':']; sub separator
	/////////////////////////////
	// Simple Validators Below //
	/////////////////////////////
	// Different meanings:
	// repeated: size range of array
	// integer: value range
	// string: count of utf-8 code point
	//
	// optional int32 min = 53101;
	E_Min = &file_options_proto_extTypes[12] // min value
	// optional int32 max = 53102;
	E_Max = &file_options_proto_extTypes[13] // max value
	// optional string range = 53103;
	E_Range = &file_options_proto_extTypes[14] // format like set description: [1,10], (1,10], [1,10), [1,~]
)

var File_options_proto protoreflect.FileDescriptor

var file_options_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x82, 0x03, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x12, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45,
	0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x01, 0x12, 0x15, 0x0a,
	0x11, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54,
	0x33, 0x32, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10,
	0x04, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x06, 0x12, 0x13,
	0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f,
	0x4c, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x15, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x16, 0x12, 0x17, 0x0a,
	0x13, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x45,
	0x54, 0x49, 0x4d, 0x45, 0x10, 0x17, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x18,
	0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x45, 0x4c, 0x4c, 0x5f, 0x4d, 0x41, 0x50, 0x10, 0x19, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x49, 0x45,
	0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x10, 0x1a, 0x2a, 0x6f, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x65, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4d,
	0x50, 0x4f, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x4c, 0x41, 0x59, 0x4f, 0x55, 0x54, 0x5f, 0x44, 0x45,
	0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x50, 0x4f,
	0x53, 0x49, 0x54, 0x45, 0x5f, 0x4c, 0x41, 0x59, 0x4f, 0x55, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x54,
	0x49, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x53,
	0x49, 0x54, 0x45, 0x5f, 0x4c, 0x41, 0x59, 0x4f, 0x55, 0x54, 0x5f, 0x48, 0x4f, 0x52, 0x49, 0x5a,
	0x4f, 0x4e, 0x54, 0x41, 0x4c, 0x10, 0x02, 0x3a, 0x3a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x62,
	0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xb9, 0x8e, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x62,
	0x6f, 0x6f, 0x6b, 0x3a, 0x3f, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x65, 0x65, 0x74,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xa1, 0x96, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x68, 0x65, 0x65, 0x74, 0x3a, 0x3b, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x72, 0x6f, 0x77, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xa2, 0x96, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x72, 0x6f,
	0x77, 0x3a, 0x3b, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x12, 0x1f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa3, 0x96,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x3a, 0x3b,
	0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x72, 0x6f, 0x77, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa4, 0x96, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x72, 0x6f, 0x77, 0x3a, 0x3f, 0x0a, 0x09, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa5, 0x96, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x3a, 0x39, 0x0a, 0x07,
	0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x89, 0x9e, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8a,
	0x9e, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x31, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8b,
	0x9e, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x3a, 0x49, 0x0a, 0x06, 0x6c,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8c, 0x9e, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x06,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x3a, 0x31, 0x0a, 0x03, 0x73, 0x65, 0x70, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8d, 0x9e, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x70, 0x3a, 0x37, 0x0a, 0x06, 0x73, 0x75, 0x62,
	0x73, 0x65, 0x70, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x8e, 0x9e, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x73,
	0x65, 0x70, 0x3a, 0x31, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xed, 0x9e, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x3a, 0x31, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xee, 0x9e, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x3a, 0x35, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xef, 0x9e, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x42,
	0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x65,
	0x6e, 0x63, 0x68, 0x79, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x61, 0x75, 0x70, 0x62, 0x3b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_proto_rawDescOnce sync.Once
	file_options_proto_rawDescData = file_options_proto_rawDesc
)

func file_options_proto_rawDescGZIP() []byte {
	file_options_proto_rawDescOnce.Do(func() {
		file_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_proto_rawDescData)
	})
	return file_options_proto_rawDescData
}

var file_options_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_options_proto_goTypes = []interface{}{
	(FieldType)(0),                    // 0: FieldType
	(CompositeLayout)(0),              // 1: CompositeLayout
	(*descriptor.FileOptions)(nil),    // 2: google.protobuf.FileOptions
	(*descriptor.MessageOptions)(nil), // 3: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
}
var file_options_proto_depIdxs = []int32{
	2,  // 0: workbook:extendee -> google.protobuf.FileOptions
	3,  // 1: worksheet:extendee -> google.protobuf.MessageOptions
	3,  // 2: captrow:extendee -> google.protobuf.MessageOptions
	3,  // 3: descrow:extendee -> google.protobuf.MessageOptions
	3,  // 4: datarow:extendee -> google.protobuf.MessageOptions
	3,  // 5: transpose:extendee -> google.protobuf.MessageOptions
	4,  // 6: caption:extendee -> google.protobuf.FieldOptions
	4,  // 7: type:extendee -> google.protobuf.FieldOptions
	4,  // 8: key:extendee -> google.protobuf.FieldOptions
	4,  // 9: layout:extendee -> google.protobuf.FieldOptions
	4,  // 10: sep:extendee -> google.protobuf.FieldOptions
	4,  // 11: subsep:extendee -> google.protobuf.FieldOptions
	4,  // 12: min:extendee -> google.protobuf.FieldOptions
	4,  // 13: max:extendee -> google.protobuf.FieldOptions
	4,  // 14: range:extendee -> google.protobuf.FieldOptions
	0,  // 15: type:type_name -> FieldType
	1,  // 16: layout:type_name -> CompositeLayout
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	15, // [15:17] is the sub-list for extension type_name
	0,  // [0:15] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_options_proto_init() }
func file_options_proto_init() {
	if File_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_options_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 15,
			NumServices:   0,
		},
		GoTypes:           file_options_proto_goTypes,
		DependencyIndexes: file_options_proto_depIdxs,
		EnumInfos:         file_options_proto_enumTypes,
		ExtensionInfos:    file_options_proto_extTypes,
	}.Build()
	File_options_proto = out.File
	file_options_proto_rawDesc = nil
	file_options_proto_goTypes = nil
	file_options_proto_depIdxs = nil
}
