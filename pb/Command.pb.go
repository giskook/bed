// Code generated by protoc-gen-go.
// source: Command.proto
// DO NOT EDIT!

/*
Package Report is a generated protocol buffer package.

It is generated from these files:
	Command.proto

It has these top-level messages:
	BedControl
	ToiletComplete
	Command
	ControlReport
*/
package Report

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Command_CommandType int32

const (
	// up message
	Command_CMT_INVALID           Command_CommandType = 0
	Command_CMT_REPBEDRUN         Command_CommandType = 259
	Command_CMT_REPTOILET         Command_CommandType = 261
	Command_CMT_REPMANUALTOILET   Command_CommandType = 262
	Command_CMT_REPMANUALBEDRUN   Command_CommandType = 260
	Command_CMT_REPTOILETCOMPLETE Command_CommandType = 263
	Command_CMT_REQBEDRUN         Command_CommandType = 33027
	Command_CMT_REQTOILET         Command_CommandType = 33029
	Command_CMT_REQDEDRESET       Command_CommandType = 33030
)

var Command_CommandType_name = map[int32]string{
	0:     "CMT_INVALID",
	259:   "CMT_REPBEDRUN",
	261:   "CMT_REPTOILET",
	262:   "CMT_REPMANUALTOILET",
	260:   "CMT_REPMANUALBEDRUN",
	263:   "CMT_REPTOILETCOMPLETE",
	33027: "CMT_REQBEDRUN",
	33029: "CMT_REQTOILET",
	33030: "CMT_REQDEDRESET",
}
var Command_CommandType_value = map[string]int32{
	"CMT_INVALID":           0,
	"CMT_REPBEDRUN":         259,
	"CMT_REPTOILET":         261,
	"CMT_REPMANUALTOILET":   262,
	"CMT_REPMANUALBEDRUN":   260,
	"CMT_REPTOILETCOMPLETE": 263,
	"CMT_REQBEDRUN":         33027,
	"CMT_REQTOILET":         33029,
	"CMT_REQDEDRESET":       33030,
}

func (x Command_CommandType) String() string {
	return proto.EnumName(Command_CommandType_name, int32(x))
}

type BedControl struct {
	Back    uint32 `protobuf:"varint,1,opt,name=back" json:"back,omitempty"`
	LegCurl uint32 `protobuf:"varint,2,opt,name=legCurl" json:"legCurl,omitempty"`
	Head    uint32 `protobuf:"varint,3,opt,name=head" json:"head,omitempty"`
	Leg     uint32 `protobuf:"varint,4,opt,name=leg" json:"leg,omitempty"`
}

func (m *BedControl) Reset()         { *m = BedControl{} }
func (m *BedControl) String() string { return proto.CompactTextString(m) }
func (*BedControl) ProtoMessage()    {}

type ToiletComplete struct {
	Style  uint32 `protobuf:"varint,1,opt,name=style" json:"style,omitempty"`
	Time   uint32 `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Weight uint32 `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Water  uint32 `protobuf:"varint,4,opt,name=water" json:"water,omitempty"`
	Wind   uint32 `protobuf:"varint,5,opt,name=wind" json:"wind,omitempty"`
}

func (m *ToiletComplete) Reset()         { *m = ToiletComplete{} }
func (m *ToiletComplete) String() string { return proto.CompactTextString(m) }
func (*ToiletComplete) ProtoMessage()    {}

type Command struct {
	Type   Command_CommandType `protobuf:"varint,1,opt,name=type,enum=Report.Command_CommandType" json:"type,omitempty"`
	Bed    *BedControl         `protobuf:"bytes,2,opt,name=bed" json:"bed,omitempty"`
	Toilet *ToiletComplete     `protobuf:"bytes,3,opt,name=toilet" json:"toilet,omitempty"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}

func (m *Command) GetBed() *BedControl {
	if m != nil {
		return m.Bed
	}
	return nil
}

func (m *Command) GetToilet() *ToiletComplete {
	if m != nil {
		return m.Toilet
	}
	return nil
}

type ControlReport struct {
	Tid          uint64   `protobuf:"varint,1,opt,name=tid" json:"tid,omitempty"`
	SerialNumber uint32   `protobuf:"varint,2,opt,name=serial_number" json:"serial_number,omitempty"`
	Command      *Command `protobuf:"bytes,3,opt,name=command" json:"command,omitempty"`
}

func (m *ControlReport) Reset()         { *m = ControlReport{} }
func (m *ControlReport) String() string { return proto.CompactTextString(m) }
func (*ControlReport) ProtoMessage()    {}

func (m *ControlReport) GetCommand() *Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func init() {
	proto.RegisterEnum("Report.Command_CommandType", Command_CommandType_name, Command_CommandType_value)
}