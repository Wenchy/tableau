package converter

import (
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/Wenchy/tableau/internal/log"
	"github.com/Wenchy/tableau/pkg/tableaupb"
	"github.com/iancoleman/strcase"
	"github.com/tealeg/xlsx/v3"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Format int

// file format
const (
	JSON      Format = 0
	Protobin         = 1
	Prototext        = 2
	// Xlsx             = 3
)

type metasheet struct {
	worksheet string // worksheet name
	captrow   int32  // exact row number of caption at worksheet
	descrow   int32  // exact row number of description at wooksheet
	datarow   int32  // start row number of data
	transpose bool   // interchange the rows and columns
}

type Tableaux struct {
	ProtoPackageName          string // protobuf package name.
	InputPath                 string // root dir of workbooks.
	OutputPath                string // output path of generated files.
	OutputFilenameAsSnakeCase bool   // output filename as snake case, default is camel case same as the protobuf message name.
	OutputFormat              Format // output format: json, protobin, or prototext. Default is json.
	OutputPretty              bool   // output pretty format, with mulitline and indent.
	LocationName              string // Location represents the collection of time offsets in use in a geographical area. Default is "Asia/Shanghai".
	// EmitUnpopulated specifies whether to emit unpopulated fields. It does not
	// emit unpopulated oneof fields or unpopulated extension fields.
	// The JSON value emitted for unpopulated fields are as follows:
	//  ╔═══════╤════════════════════════════╗
	//  ║ JSON  │ Protobuf field             ║
	//  ╠═══════╪════════════════════════════╣
	//  ║ false │ proto3 boolean fields      ║
	//  ║ 0     │ proto3 numeric fields      ║
	//  ║ ""    │ proto3 string/bytes fields ║
	//  ║ null  │ proto2 scalar fields       ║
	//  ║ null  │ message fields             ║
	//  ║ []    │ list fields                ║
	//  ║ {}    │ map fields                 ║
	//  ╚═══════╧════════════════════════════╝
	EmitUnpopulated bool
	metasheet       metasheet // meta info of worksheet
}

var specialMessageMap = map[string]int{
	"google.protobuf.Timestamp": 1,
	"google.protobuf.Duration":  1,
}

func (tbx *Tableaux) Convert() {
	// parseActivity()
	// parseItem()
	// numFiles := protoregistry.GlobalFiles.NumFiles()
	// log.Logger.Debug("numFiles", numFiles)
	// protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
	// 	log.Logger.Debugf("filepath: %s", fd.Path())
	// 	return true
	// })
	// log.Logger.Debug("====================")

	// create oupput dir
	err := os.MkdirAll(tbx.OutputPath, 0700)
	if err != nil {
		panic(err)
	}

	protoPackage := protoreflect.FullName(tbx.ProtoPackageName)
	protoregistry.GlobalFiles.RangeFilesByPackage(protoPackage, func(fd protoreflect.FileDescriptor) bool {
		log.Logger.Debugf("filepath: %s", fd.Path())
		opts := fd.Options().(*descriptorpb.FileOptions)
		workbook := proto.GetExtension(opts, tableaupb.E_Workbook).(string)
		if workbook == "" {
			return true
		}

		log.Logger.Debugf("proto: %s, workbook %s", fd.Path(), workbook)
		msgs := fd.Messages()
		for i := 0; i < msgs.Len(); i++ {
			md := msgs.Get(i)
			// log.Logger.Debugf("%s", md.FullName())
			opts := md.Options().(*descriptorpb.MessageOptions)
			worksheet := proto.GetExtension(opts, tableaupb.E_Worksheet).(string)
			if worksheet != "" {
				log.Logger.Infof("generate: %s, message: %s@%s, worksheet: %s@%s", md.Name(), fd.Path(), md.Name(), workbook, worksheet)
			}
			newMsg := dynamicpb.NewMessage(md)
			tbx.Export(newMsg)
		}
		return true
	})
}

// Export the protomsg message.
func (tbx *Tableaux) Export(protomsg proto.Message) {
	md := protomsg.ProtoReflect().Descriptor()
	msg := protomsg.ProtoReflect()
	_, workbook := TestParseFileOptions(md.ParentFile())
	log.Logger.Debug("==================")
	msgName, worksheet, captrow, descrow, datarow, transpose := TestParseMessageOptions(md)
	tbx.metasheet.worksheet = worksheet
	tbx.metasheet.captrow = captrow
	tbx.metasheet.descrow = descrow
	tbx.metasheet.datarow = datarow
	tbx.metasheet.transpose = transpose

	log.Logger.Debug("==================")
	sheet := ReadSheet(tbx.InputPath+workbook, worksheet)
	if transpose {
		// col caprow: caption row
		// col [datarow, MaxRow]: data
		for ncol := 0; ncol < sheet.MaxCol; ncol++ {
			if ncol >= int(datarow)-1 {
				// row, err := sheet.Row(nrow)
				// if err != nil {
				// 	panic(err)
				// }
				kv := make(map[string]string)
				for i := 0; i < sheet.MaxRow; i++ {
					captionCell, err := sheet.Cell(i, int(captrow)-1)
					if err != nil {
						panic(err)
					}
					key := captionCell.Value
					dataCell, err := sheet.Cell(i, ncol)
					if err != nil {
						panic(err)
					}
					value := dataCell.Value
					re := regexp.MustCompile(`\r?\n?`)
					key = re.ReplaceAllString(key, "")
					kv[key] = value
				}
				tbx.TestParseFieldOptions(msg, kv, 0, "")
			}
		}

	} else {
		// row captrow: caption row
		// row [datarow, MaxRow]: data row
		for nrow := 0; nrow < sheet.MaxRow; nrow++ {
			if nrow >= int(datarow)-1 {
				// row, err := sheet.Row(nrow)
				// if err != nil {
				// 	panic(err)
				// }
				kv := make(map[string]string)
				for i := 0; i < sheet.MaxCol; i++ {
					captionCell, err := sheet.Cell(int(captrow)-1, i)
					if err != nil {
						panic(err)
					}
					key := captionCell.Value
					dataCell, err := sheet.Cell(nrow, i)
					if err != nil {
						panic(err)
					}
					value := dataCell.Value
					re := regexp.MustCompile(`\r?\n?`)
					key = re.ReplaceAllString(key, "")
					kv[key] = value
				}
				tbx.TestParseFieldOptions(msg, kv, 0, "")
			}
		}
	}
	log.Logger.Debug("==================")
	filename := msgName
	if tbx.OutputFilenameAsSnakeCase {
		filename = strcase.ToSnake(msgName)
	}
	filePath := tbx.OutputPath + filename
	switch tbx.OutputFormat {
	case JSON:
		exportJSON(protomsg, filePath, tbx.OutputPretty, tbx.EmitUnpopulated)
	case Protobin:
		exportProtobin(protomsg, filePath)
	case Prototext:
		exportPrototext(protomsg, filePath, tbx.OutputPretty)
	default:
		log.Logger.Debugf("unknown format, default to JSON")
		exportJSON(protomsg, filePath, tbx.OutputPretty, tbx.EmitUnpopulated)
	}
}

func exportJSON(protomsg proto.Message, filePath string, pretty bool, emitUnpopulated bool) {
	var out []byte
	var err error
	if pretty {
		opts := protojson.MarshalOptions{
			Multiline:       true,
			Indent:          "    ",
			EmitUnpopulated: emitUnpopulated,
		}
		out, err = opts.Marshal(protomsg)
		if err != nil {
			log.Logger.Panicf("Failed to marshal protomsg as JSON: %v", err)
		}
	} else {
		out, err = protojson.Marshal(protomsg)
		if err != nil {
			panic(err)
		}
	}
	err = ioutil.WriteFile(filePath+".json", out, 0644)
	if err != nil {
		panic(err)
	}
	// out.WriteTo(os.Stdout)
}

func exportProtobin(protomsg proto.Message, filePath string) {
	out, err := proto.Marshal(protomsg)
	if err != nil {
		log.Logger.Panicf("Failed to encode protomsg: %v", err)
	}
	if err := ioutil.WriteFile(filePath+".protobin", out, 0644); err != nil {
		log.Logger.Panicf("Failed to write file: %v", err)
	}
	// out.WriteTo(os.Stdout)
}

func exportPrototext(protomsg proto.Message, filePath string, pretty bool) {
	var out []byte
	var err error
	if pretty {
		opts := prototext.MarshalOptions{
			Multiline: true,
			Indent:    "    ",
		}
		out, err = opts.Marshal(protomsg)
		if err != nil {
			log.Logger.Panicf("failed to encode protomsg: %v", err)
		}
	} else {
		out, err = prototext.Marshal(protomsg)
		if err != nil {
			panic(err)
		}
	}
	if err := ioutil.WriteFile(filePath+".prototext", out, 0644); err != nil {
		log.Logger.Panicf("failed to write file: %v", err)
	}
	// out.WriteTo(os.Stdout)
}

func getTabStr(depth int) string {
	tab := ""
	for i := 0; i < depth; i++ {
		tab += "\t"
	}
	return tab
}

// ReadSheet read a sheet from specified workbook.
func ReadSheet(workbook string, worksheet string) *xlsx.Sheet {
	// open an existing file
	wb, err := xlsx.OpenFile(workbook)
	if err != nil {
		panic(err)
	}
	sh, ok := wb.Sheet[worksheet]
	if !ok {
		log.Logger.Panicf("Sheet %s does not exist in %s", worksheet, workbook)
	}
	exportSheet(sh)
	log.Logger.Debug("----")
	return sh
}

func exportSheet(sheet *xlsx.Sheet) {
	log.Logger.Debugf("MaxCol: %d, MaxRow: %d", sheet.MaxCol, sheet.MaxRow)
	// row 0: captrow
	// row 1 - MaxRow: datarow
	for nrow := 0; nrow < sheet.MaxRow; nrow++ {
		for ncol := 0; ncol < sheet.MaxCol; ncol++ {
			// get the Cell in D1, which is row 0, column 3
			cell, err := sheet.Cell(nrow, ncol)
			if err != nil {
				panic(err)
			}
			log.Logger.Debugf("%s ", cell.Value)
		}
	}
}

// TestParseFileOptions is aimed to parse the options of a protobuf definition file.
func TestParseFileOptions(fd protoreflect.FileDescriptor) (string, string) {
	opts := fd.Options().(*descriptorpb.FileOptions)
	protofile := string(fd.FullName())
	workbook := proto.GetExtension(opts, tableaupb.E_Workbook).(string)
	log.Logger.Debugf("file:%s.proto, workbook:%s", protofile, workbook)
	return protofile, workbook
}

// TestParseMessageOptions is aimed to parse the options of a protobuf message.
func TestParseMessageOptions(md protoreflect.MessageDescriptor) (string, string, int32, int32, int32, bool) {
	opts := md.Options().(*descriptorpb.MessageOptions)
	msgName := string(md.Name())
	worksheet := proto.GetExtension(opts, tableaupb.E_Worksheet).(string)
	captrow := proto.GetExtension(opts, tableaupb.E_Captrow).(int32)
	if captrow == 0 {
		captrow = 1 // default
	}
	descrow := proto.GetExtension(opts, tableaupb.E_Descrow).(int32)
	if descrow == 0 {
		descrow = 1 // default
	}
	datarow := proto.GetExtension(opts, tableaupb.E_Datarow).(int32)
	if datarow == 0 {
		datarow = 2 // default
	}
	transpose := proto.GetExtension(opts, tableaupb.E_Transpose).(bool)
	log.Logger.Debugf("message:%s, worksheet:%s, captrow:%d, descrow:%d, datarow:%d, transpose:%v", msgName, worksheet, captrow, descrow, datarow, transpose)
	return msgName, worksheet, captrow, descrow, datarow, transpose
}

// TestParseFieldOptions is aimed to parse the options of all the fields of a protobuf message.
func (tbx *Tableaux) TestParseFieldOptions(msg protoreflect.Message, row map[string]string, depth int, prefix string) {
	md := msg.Descriptor()
	opts := md.Options().(*descriptorpb.MessageOptions)
	worksheet := proto.GetExtension(opts, tableaupb.E_Worksheet).(string)
	pkg := md.ParentFile().Package()
	log.Logger.Debugf("%s// %s, '%s', %v, %v, %v", getTabStr(depth), md.FullName(), worksheet, md.IsMapEntry(), prefix, pkg)
	for i := 0; i < md.Fields().Len(); i++ {
		fd := md.Fields().Get(i)
		if string(pkg) != tbx.ProtoPackageName && pkg != "google.protobuf" {
			log.Logger.Debugf("%s// no need to proces package: %v", getTabStr(depth), pkg)
			return
		}
		msgName := ""
		if fd.Kind() == protoreflect.MessageKind {
			msgName = string(fd.Message().FullName())
			// log.Logger.Debug(fd.Cardinality().String(), fd.Kind().String(), fd.FullName(), fd.Number())
			// ParseFieldOptions(fd.Message(), depth+1)
		}

		// if fd.IsList() {
		// 	log.Logger.Debug("repeated", fd.Kind().String(), fd.FullName().Name())
		// 	// Redact(fd.Options().ProtoReflect().Interface())
		// }
		opts := fd.Options().(*descriptorpb.FieldOptions)
		caption := proto.GetExtension(opts, tableaupb.E_Caption).(string)
		etype := proto.GetExtension(opts, tableaupb.E_Type).(tableaupb.FieldType)
		key := proto.GetExtension(opts, tableaupb.E_Key).(string)
		layout := proto.GetExtension(opts, tableaupb.E_Layout).(tableaupb.CompositeLayout)
		sep := proto.GetExtension(opts, tableaupb.E_Sep).(string)
		if sep == "" {
			sep = ","
		}
		subsep := proto.GetExtension(opts, tableaupb.E_Subsep).(string)
		if subsep == "" {
			subsep = ":"
		}
		log.Logger.Debugf("%s%s(%v) %s(%s) %s = %d [(caption) = \"%s\", (type) = %s, (key) = \"%s\", (layout) = \"%s\", (sep) = \"%s\"];",
			getTabStr(depth), fd.Cardinality().String(), fd.IsMap(), fd.Kind().String(), msgName, fd.FullName().Name(), fd.Number(), prefix+caption, etype.String(), layout.String(), key, sep)
		// log.Logger.Debug(fd.ContainingMessage().FullName())

		// if fd.Cardinality() == protoreflect.Repeated && fd.Kind() == protoreflect.MessageKind {
		// 	msg := fd.Message().New()
		// }

		// NOTE(wenchy): `proto.Equal` treats a nil message as not equal to an empty one.
		// doc: [Equal](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#Equal)
		// issue: [APIv2: protoreflect: consider Message nilness test](https://github.com/golang/protobuf/issues/966)
		// ```
		// nilMessage = (*MyMessage)(nil)
		// emptyMessage = new(MyMessage)
		//
		// Equal(nil, nil)                   // true
		// Equal(nil, nilMessage)            // false
		// Equal(nil, emptyMessage)          // false
		// Equal(nilMessage, nilMessage)     // true
		// Equal(nilMessage, emptyMessage)   // ??? false
		// Equal(emptyMessage, emptyMessage) // true
		// ```
		//
		// Case: `subMsg := msg.Mutable(fd).Message()`
		// `Message.Mutable` will allocate new "empty message", and is not equal to "nil"
		//
		// Solution:
		// 1. spawn two values: `emptyValue` and `newValue`
		// 2. set `newValue` back to field if `newValue` is not equal to `emptyValue`
		emptyValue := msg.NewField(fd)
		newValue := msg.NewField(fd)
		if fd.IsMap() {
			// Mutable returns a mutable reference to a composite type.
			if msg.Has(fd) {
				newValue = msg.Mutable(fd)
			}
			reflectMap := newValue.Map()
			// reflectMap := msg.Mutable(fd).Map()
			keyFd := fd.MapKey()
			valueFd := fd.MapValue()
			// newKey := protoreflect.ValueOf(int32(1)).MapKey()
			// newKey := tbx.getFieldValue(keyFd, "1111001").MapKey()
			if etype == tableaupb.FieldType_FIELD_TYPE_CELL_MAP {
				if valueFd.Kind() == protoreflect.MessageKind {
					log.Logger.Panicf("in-cell map do not support value as message type")
				}
				cellValue, ok := row[prefix+caption]
				if !ok {
					log.Logger.Panicf("column caption not found: %s", prefix+caption)
				}
				if cellValue != "" {
					// If s does not contain sep and sep is not empty, Split returns a
					// slice of length 1 whose only element is s.
					splits := strings.Split(cellValue, sep)
					for _, pair := range splits {
						kv := strings.Split(pair, subsep)
						if len(kv) != 2 {
							log.Logger.Panicf("illegal key-value pair: %v, %v", prefix+caption, pair)
						}
						key := tbx.getFieldValue(keyFd, kv[0]).MapKey()
						val := reflectMap.NewValue()
						val = tbx.getFieldValue(valueFd, kv[1])
						reflectMap.Set(key, val)
					}
				}
			} else {
				emptyMapValue := reflectMap.NewValue()
				if valueFd.Kind() == protoreflect.MessageKind {
					if layout == tableaupb.CompositeLayout_COMPOSITE_LAYOUT_HORIZONTAL {
						size := getPrefixSize(row, prefix+caption)
						// log.Logger.Debug("prefix size: ", size)
						for i := 1; i <= size; i++ {
							newMapKey := keyFd.Default().MapKey()
							cellValue, ok := row[prefix+caption+strconv.Itoa(i)+key]
							if !ok {
								log.Logger.Panicf("key not found: %s", prefix+caption+key)
							}
							newMapKey = tbx.getFieldValue(keyFd, cellValue).MapKey()
							var newMapValue protoreflect.Value
							if reflectMap.Has(newMapKey) {
								newMapValue = reflectMap.Mutable(newMapKey)
							} else {
								newMapValue = reflectMap.NewValue()
							}
							tbx.TestParseFieldOptions(newMapValue.Message(), row, depth+1, prefix+caption+strconv.Itoa(i))
							if !MessageValueEqual(emptyMapValue, newMapValue) {
								reflectMap.Set(newMapKey, newMapValue)
							}
						}
					} else {
						newMapKey := keyFd.Default().MapKey()
						cellValue, ok := row[prefix+caption+key]
						if !ok {
							log.Logger.Panicf("key not found: %s", prefix+caption+key)
						}
						newMapKey = tbx.getFieldValue(keyFd, cellValue).MapKey()
						var newMapValue protoreflect.Value
						if reflectMap.Has(newMapKey) {
							newMapValue = reflectMap.Mutable(newMapKey)
						} else {
							newMapValue = reflectMap.NewValue()
						}
						tbx.TestParseFieldOptions(newMapValue.Message(), row, depth+1, prefix+caption)
						if !MessageValueEqual(emptyMapValue, newMapValue) {
							reflectMap.Set(newMapKey, newMapValue)
						}
					}
				} else {
					// value is scalar type
					key := "Key"     // deafult key caption
					value := "Value" // deafult value caption
					newMapKey := keyFd.Default().MapKey()
					// key cell
					cellValue, ok := row[prefix+caption+key]
					if !ok {
						log.Logger.Panicf("key not found: %s", prefix+caption+key)
					}
					newMapKey = tbx.getFieldValue(keyFd, cellValue).MapKey()
					var newMapValue protoreflect.Value
					if reflectMap.Has(newMapKey) {
						newMapValue = reflectMap.Mutable(newMapKey)
					} else {
						newMapValue = reflectMap.NewValue()
					}
					// value cell
					cellValue, ok = row[prefix+caption+value]
					if !ok {
						log.Logger.Panicf("value not found: %s", prefix+caption+value)
					}
					newMapValue = tbx.getFieldValue(fd, cellValue)
					if !reflectMap.Has(newMapKey) {
						reflectMap.Set(newMapKey, newMapValue)
					}
				}
			}
			if !msg.Has(fd) && reflectMap.Len() != 0 {
				msg.Set(fd, newValue)
			}
		} else if fd.IsList() {
			// Mutable returns a mutable reference to a composite type.
			if msg.Has(fd) {
				newValue = msg.Mutable(fd)
			}
			reflectList := newValue.List()
			if fd.Kind() == protoreflect.MessageKind {
				emptyListValue := reflectList.NewElement()
				if layout == tableaupb.CompositeLayout_COMPOSITE_LAYOUT_VERTICAL {
					newListValue := reflectList.NewElement()
					tbx.TestParseFieldOptions(newListValue.Message(), row, depth+1, prefix+caption)
					if !MessageValueEqual(emptyListValue, newListValue) {
						reflectList.Append(newListValue)
					}
				} else {
					size := getPrefixSize(row, prefix+caption)
					// log.Logger.Debug("prefix size: ", size)
					for i := 1; i <= size; i++ {
						newListValue := reflectList.NewElement()
						tbx.TestParseFieldOptions(newListValue.Message(), row, depth+1, prefix+caption+strconv.Itoa(i))
						if !MessageValueEqual(emptyListValue, newListValue) {
							reflectList.Append(newListValue)
						}
					}
				}
			} else {
				if etype == tableaupb.FieldType_FIELD_TYPE_CELL_LIST {
					cellValue, ok := row[prefix+caption]
					if !ok {
						log.Logger.Panicf("caption not found: %s", prefix+caption)
					}
					if cellValue != "" {
						// If s does not contain sep and sep is not empty, Split returns a
						// slice of length 1 whose only element is s.
						splits := strings.Split(cellValue, sep)
						for _, v := range splits {
							value := tbx.getFieldValue(fd, v)
							reflectList.Append(value)
						}
					}
				} else {
					log.Logger.Panicf("unknown list type: %v", etype)
				}
			}
			if !msg.Has(fd) && reflectList.Len() != 0 {
				msg.Set(fd, newValue)
			}
		} else {
			if fd.Kind() == protoreflect.MessageKind {
				if etype == tableaupb.FieldType_FIELD_TYPE_CELL_MESSAGE {
					cellValue, ok := row[prefix+caption]
					if !ok {
						log.Logger.Panicf("not found column caption: %v", prefix+caption)
					}
					if cellValue != "" {
						// If s does not contain sep and sep is not empty, Split returns a
						// slice of length 1 whose only element is s.
						splits := strings.Split(cellValue, sep)
						subMd := newValue.Message().Descriptor()
						for i := 0; i < subMd.Fields().Len() && i < len(splits); i++ {
							fd := subMd.Fields().Get(i)
							// log.Logger.Debugf("fd.FullName().Name(): ", fd.FullName().Name())
							value := tbx.getFieldValue(fd, splits[i])
							newValue.Message().Set(fd, value)
						}
					}
				} else {
					subMsgName := string(fd.Message().FullName())
					_, found := specialMessageMap[subMsgName]
					if found {
						cellValue, ok := row[prefix+caption]
						if !ok {
							log.Logger.Panicf("not found column caption: %v", prefix+caption)
						}
						newValue = tbx.getFieldValue(fd, cellValue)
					} else {
						pkgName := newValue.Message().Descriptor().ParentFile().Package()
						if string(pkgName) != tbx.ProtoPackageName {
							log.Logger.Panicf("unknown message %v in package %v", subMsgName, pkgName)
						}
						tbx.TestParseFieldOptions(newValue.Message(), row, depth+1, prefix+caption)
					}
				}
				if !MessageValueEqual(emptyValue, newValue) {
					msg.Set(fd, newValue)
				}
			} else {
				cellValue, ok := row[prefix+caption]
				if !ok {
					log.Logger.Panicf("not found column caption: %v", prefix+caption)
				}
				newValue = tbx.getFieldValue(fd, cellValue)
				msg.Set(fd, newValue)
			}
		}
	}
}

func MessageValueEqual(v1, v2 protoreflect.Value) bool {
	if proto.Equal(v1.Message().Interface(), v2.Message().Interface()) {
		log.Logger.Debug("empty message exists")
		return true
	}
	return false
}

func getPrefixSize(row map[string]string, prefix string) int {
	// log.Logger.Debug("caption prefix: ", prefix)
	size := 0
	for caption := range row {
		if strings.HasPrefix(caption, prefix) {
			num := 0
			// log.Logger.Debug("caption: ", caption)
			colSuffix := caption[len(prefix):]
			// log.Logger.Debug("caption: suffix ", colSuffix)
			for _, r := range colSuffix {
				if unicode.IsDigit(r) {
					num = num*10 + int(r-'0')
				} else {
					break
				}
			}
			size = int(math.Max(float64(size), float64(num)))
		}
	}
	return size
}

func (tbx *Tableaux) getFieldValue(fd protoreflect.FieldDescriptor, cellValue string) protoreflect.Value {
	switch fd.Kind() {
	case protoreflect.Int32Kind:
		var val int64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseInt(cellValue, 10, 32)
			if err != nil {
				log.Logger.Debug("cellValue:", cellValue)
				panic(err)
			}
		}
		return protoreflect.ValueOf(int32(val))
	case protoreflect.Sint32Kind:
		var val int64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseInt(cellValue, 10, 32)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(int32(val))
	case protoreflect.Sfixed32Kind:
		var val int64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseInt(cellValue, 10, 32)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(int32(val))
	case protoreflect.Uint32Kind:
		var val uint64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseUint(cellValue, 10, 32)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(uint32(val))
	case protoreflect.Fixed32Kind:
		var val uint64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseUint(cellValue, 10, 32)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(uint32(val))
	case protoreflect.Int64Kind:
		var val int64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseInt(cellValue, 10, 64)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(int64(val))
	case protoreflect.Sint64Kind:
		var val int64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseInt(cellValue, 10, 64)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(int64(val))
	case protoreflect.Sfixed64Kind:
		var val int64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseInt(cellValue, 10, 64)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(int64(val))
	case protoreflect.Uint64Kind:
		var val uint64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseUint(cellValue, 10, 64)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(uint64(val))
	case protoreflect.Fixed64Kind:
		var val uint64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseUint(cellValue, 10, 64)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(uint64(val))
	case protoreflect.StringKind:
		return protoreflect.ValueOf(cellValue)
	case protoreflect.BytesKind:
		return protoreflect.ValueOf([]byte(cellValue))
	case protoreflect.BoolKind:
		var val bool // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseBool(cellValue)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(val)
	case protoreflect.FloatKind:
		var val float64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseFloat(cellValue, 32)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(float32(val))
	case protoreflect.DoubleKind:
		var val float64 // default
		var err error
		if cellValue != "" {
			val, err = strconv.ParseFloat(cellValue, 64)
			if err != nil {
				panic(err)
			}
		}
		return protoreflect.ValueOf(float64(val))
	case protoreflect.MessageKind:
		msgName := fd.Message().FullName()
		switch msgName {
		case "google.protobuf.Timestamp":
			ts := &timestamppb.Timestamp{}
			if cellValue != "" {
				// location name: "Asia/Shanghai" or "Asia/Chongqing".
				// NOTE(wenchy): There is no "Asia/Beijing" location name. Whoa!!! Big surprize?
				t, err := parseTimeWithLocation(tbx.LocationName, cellValue)
				if err != nil {
					log.Logger.Panicf("illegal timestamp string format: %v, err: %v", cellValue, err)
				}
				// log.Logger.Debugf("timeStr: %v, unix timestamp: %v", cellValue, t.Unix())
				ts = timestamppb.New(t)
				// make use of t as a *timestamppb.Timestamp
			}
			if err := ts.CheckValid(); err != nil {
				log.Logger.Panicf("invalid timestamp: %v", err)
			}
			return protoreflect.ValueOf(ts.ProtoReflect())
		case "google.protobuf.Duration":
			dur := &durationpb.Duration{} // default
			if cellValue != "" {
				d, err := time.ParseDuration(cellValue)
				if err != nil {
					log.Logger.Panicf("ParseDuration failed, illegal format: %v", cellValue)
				}
				dur = durationpb.New(d)
				// make use of d as a *durationpb.Duration
			}
			if err := dur.CheckValid(); err != nil {
				log.Logger.Panicf("duration CheckValid failed: %v", err)
			}
			return protoreflect.ValueOf(dur.ProtoReflect())
		default:
			log.Logger.Panicf("not supported message type: %s", msgName)
		}
	default:
		log.Logger.Panicf("not supported scalar type: %s", fd.Kind().String())
		// case protoreflect.EnumKind:
		// 	log.Logger.Panicf("not supported key type: %s", fd.Kind().String())
		// case protoreflect.GroupKind:
		// 	log.Logger.Panicf("not supported key type: %s", fd.Kind().String())
		// 	return protoreflect.Value{}
	}
	log.Logger.Panicf("should not go here")
	return protoreflect.Value{}
}

func parseTimeWithLocation(locationName string, timeStr string) (time.Time, error) {
	// see https://golang.org/pkg/time/#LoadLocation
	if location, err := time.LoadLocation(locationName); err != nil {
		log.Logger.Panicf("LoadLocation failed: %s", err)
		return time.Time{}, err
	} else {
		timeLayout := "2006-01-02 15:04:05"
		t, err := time.ParseInLocation(timeLayout, timeStr, location)
		if err != nil {
			log.Logger.Panicf("ParseInLocation failed:%v ,timeStr: %v, locationName: %v", err, timeStr, locationName)
		}
		return t, nil
	}
}
