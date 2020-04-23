// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /Users/aloaugus/src/vpp-api-repo/20.05-rc0~540_g77ea42b~b9261/api/plugins/mactime.api.json

/*
Package mactime is a generated VPP binary API for 'mactime' module.

It consists of:
	  6 enums
	  2 aliases
	  2 types
	  7 messages
	  3 services
*/
package mactime

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "mactime"
	// APIVersion is the API version of this module.
	APIVersion = "2.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xd32f01e0
)

// IfStatusFlags represents VPP binary API enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var IfStatusFlags_name = map[uint32]string{
	1: "IF_STATUS_API_FLAG_ADMIN_UP",
	2: "IF_STATUS_API_FLAG_LINK_UP",
}

var IfStatusFlags_value = map[string]uint32{
	"IF_STATUS_API_FLAG_ADMIN_UP": 1,
	"IF_STATUS_API_FLAG_LINK_UP":  2,
}

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IfType represents VPP binary API enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 1
	IF_API_TYPE_SUB      IfType = 2
	IF_API_TYPE_P2P      IfType = 3
	IF_API_TYPE_PIPE     IfType = 4
)

var IfType_name = map[uint32]string{
	1: "IF_API_TYPE_HARDWARE",
	2: "IF_API_TYPE_SUB",
	3: "IF_API_TYPE_P2P",
	4: "IF_API_TYPE_PIPE",
}

var IfType_value = map[string]uint32{
	"IF_API_TYPE_HARDWARE": 1,
	"IF_API_TYPE_SUB":      2,
	"IF_API_TYPE_P2P":      3,
	"IF_API_TYPE_PIPE":     4,
}

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LinkDuplex represents VPP binary API enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var LinkDuplex_name = map[uint32]string{
	0: "LINK_DUPLEX_API_UNKNOWN",
	1: "LINK_DUPLEX_API_HALF",
	2: "LINK_DUPLEX_API_FULL",
}

var LinkDuplex_value = map[string]uint32{
	"LINK_DUPLEX_API_UNKNOWN": 0,
	"LINK_DUPLEX_API_HALF":    1,
	"LINK_DUPLEX_API_FULL":    2,
}

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// MtuProto represents VPP binary API enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 1
	MTU_PROTO_API_IP4  MtuProto = 2
	MTU_PROTO_API_IP6  MtuProto = 3
	MTU_PROTO_API_MPLS MtuProto = 4
	MTU_PROTO_API_N    MtuProto = 5
)

var MtuProto_name = map[uint32]string{
	1: "MTU_PROTO_API_L3",
	2: "MTU_PROTO_API_IP4",
	3: "MTU_PROTO_API_IP6",
	4: "MTU_PROTO_API_MPLS",
	5: "MTU_PROTO_API_N",
}

var MtuProto_value = map[string]uint32{
	"MTU_PROTO_API_L3":   1,
	"MTU_PROTO_API_IP4":  2,
	"MTU_PROTO_API_IP6":  3,
	"MTU_PROTO_API_MPLS": 4,
	"MTU_PROTO_API_N":    5,
}

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// RxMode represents VPP binary API enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var RxMode_name = map[uint32]string{
	0: "RX_MODE_API_UNKNOWN",
	1: "RX_MODE_API_POLLING",
	2: "RX_MODE_API_INTERRUPT",
	3: "RX_MODE_API_ADAPTIVE",
	4: "RX_MODE_API_DEFAULT",
}

var RxMode_value = map[string]uint32{
	"RX_MODE_API_UNKNOWN":   0,
	"RX_MODE_API_POLLING":   1,
	"RX_MODE_API_INTERRUPT": 2,
	"RX_MODE_API_ADAPTIVE":  3,
	"RX_MODE_API_DEFAULT":   4,
}

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// SubIfFlags represents VPP binary API enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var SubIfFlags_name = map[uint32]string{
	1:   "SUB_IF_API_FLAG_NO_TAGS",
	2:   "SUB_IF_API_FLAG_ONE_TAG",
	4:   "SUB_IF_API_FLAG_TWO_TAGS",
	8:   "SUB_IF_API_FLAG_DOT1AD",
	16:  "SUB_IF_API_FLAG_EXACT_MATCH",
	32:  "SUB_IF_API_FLAG_DEFAULT",
	64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
	128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
	254: "SUB_IF_API_FLAG_MASK_VNET",
	256: "SUB_IF_API_FLAG_DOT1AH",
}

var SubIfFlags_value = map[string]uint32{
	"SUB_IF_API_FLAG_NO_TAGS":           1,
	"SUB_IF_API_FLAG_ONE_TAG":           2,
	"SUB_IF_API_FLAG_TWO_TAGS":          4,
	"SUB_IF_API_FLAG_DOT1AD":            8,
	"SUB_IF_API_FLAG_EXACT_MATCH":       16,
	"SUB_IF_API_FLAG_DEFAULT":           32,
	"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
	"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
	"SUB_IF_API_FLAG_MASK_VNET":         254,
	"SUB_IF_API_FLAG_DOT1AH":            256,
}

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// InterfaceIndex represents VPP binary API alias 'interface_index'.
type InterfaceIndex uint32

// MacAddress represents VPP binary API alias 'mac_address'.
type MacAddress [6]uint8

// MactimeTimeRange represents VPP binary API type 'mactime_time_range'.
type MactimeTimeRange struct {
	Start float64
	End   float64
}

func (*MactimeTimeRange) GetTypeName() string { return "mactime_time_range" }

// TimeRange represents VPP binary API type 'time_range'.
type TimeRange struct {
	Start float64
	End   float64
}

func (*TimeRange) GetTypeName() string { return "time_range" }

// MactimeAddDelRange represents VPP binary API message 'mactime_add_del_range'.
type MactimeAddDelRange struct {
	IsAdd      bool
	Drop       bool
	Allow      bool
	AllowQuota uint8
	NoUDP10001 bool
	DataQuota  uint64
	MacAddress MacAddress
	DeviceName string `struc:"[64]byte"`
	Count      uint32 `struc:"sizeof=Ranges"`
	Ranges     []TimeRange
}

func (m *MactimeAddDelRange) Reset()                        { *m = MactimeAddDelRange{} }
func (*MactimeAddDelRange) GetMessageName() string          { return "mactime_add_del_range" }
func (*MactimeAddDelRange) GetCrcString() string            { return "101858ef" }
func (*MactimeAddDelRange) GetMessageType() api.MessageType { return api.RequestMessage }

// MactimeAddDelRangeReply represents VPP binary API message 'mactime_add_del_range_reply'.
type MactimeAddDelRangeReply struct {
	Retval int32
}

func (m *MactimeAddDelRangeReply) Reset()                        { *m = MactimeAddDelRangeReply{} }
func (*MactimeAddDelRangeReply) GetMessageName() string          { return "mactime_add_del_range_reply" }
func (*MactimeAddDelRangeReply) GetCrcString() string            { return "e8d4e804" }
func (*MactimeAddDelRangeReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MactimeDetails represents VPP binary API message 'mactime_details'.
type MactimeDetails struct {
	PoolIndex       uint32
	MacAddress      MacAddress
	DataQuota       uint64
	DataUsedInRange uint64
	Flags           uint32
	DeviceName      string `struc:"[64]byte"`
	Nranges         uint32 `struc:"sizeof=Ranges"`
	Ranges          []MactimeTimeRange
}

func (m *MactimeDetails) Reset()                        { *m = MactimeDetails{} }
func (*MactimeDetails) GetMessageName() string          { return "mactime_details" }
func (*MactimeDetails) GetCrcString() string            { return "44921c06" }
func (*MactimeDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// MactimeDump represents VPP binary API message 'mactime_dump'.
type MactimeDump struct {
	MyTableEpoch uint32
}

func (m *MactimeDump) Reset()                        { *m = MactimeDump{} }
func (*MactimeDump) GetMessageName() string          { return "mactime_dump" }
func (*MactimeDump) GetCrcString() string            { return "8f454e23" }
func (*MactimeDump) GetMessageType() api.MessageType { return api.RequestMessage }

// MactimeDumpReply represents VPP binary API message 'mactime_dump_reply'.
type MactimeDumpReply struct {
	Retval     int32
	TableEpoch uint32
}

func (m *MactimeDumpReply) Reset()                        { *m = MactimeDumpReply{} }
func (*MactimeDumpReply) GetMessageName() string          { return "mactime_dump_reply" }
func (*MactimeDumpReply) GetCrcString() string            { return "49bcc753" }
func (*MactimeDumpReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MactimeEnableDisable represents VPP binary API message 'mactime_enable_disable'.
type MactimeEnableDisable struct {
	EnableDisable bool
	SwIfIndex     InterfaceIndex
}

func (m *MactimeEnableDisable) Reset()                        { *m = MactimeEnableDisable{} }
func (*MactimeEnableDisable) GetMessageName() string          { return "mactime_enable_disable" }
func (*MactimeEnableDisable) GetCrcString() string            { return "3865946c" }
func (*MactimeEnableDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// MactimeEnableDisableReply represents VPP binary API message 'mactime_enable_disable_reply'.
type MactimeEnableDisableReply struct {
	Retval int32
}

func (m *MactimeEnableDisableReply) Reset()                        { *m = MactimeEnableDisableReply{} }
func (*MactimeEnableDisableReply) GetMessageName() string          { return "mactime_enable_disable_reply" }
func (*MactimeEnableDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*MactimeEnableDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*MactimeAddDelRange)(nil), "mactime.MactimeAddDelRange")
	api.RegisterMessage((*MactimeAddDelRangeReply)(nil), "mactime.MactimeAddDelRangeReply")
	api.RegisterMessage((*MactimeDetails)(nil), "mactime.MactimeDetails")
	api.RegisterMessage((*MactimeDump)(nil), "mactime.MactimeDump")
	api.RegisterMessage((*MactimeDumpReply)(nil), "mactime.MactimeDumpReply")
	api.RegisterMessage((*MactimeEnableDisable)(nil), "mactime.MactimeEnableDisable")
	api.RegisterMessage((*MactimeEnableDisableReply)(nil), "mactime.MactimeEnableDisableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*MactimeAddDelRange)(nil),
		(*MactimeAddDelRangeReply)(nil),
		(*MactimeDetails)(nil),
		(*MactimeDump)(nil),
		(*MactimeDumpReply)(nil),
		(*MactimeEnableDisable)(nil),
		(*MactimeEnableDisableReply)(nil),
	}
}

// RPCService represents RPC service API for mactime module.
type RPCService interface {
	DumpMactime(ctx context.Context, in *MactimeDump) (RPCService_DumpMactimeClient, error)
	MactimeAddDelRange(ctx context.Context, in *MactimeAddDelRange) (*MactimeAddDelRangeReply, error)
	MactimeEnableDisable(ctx context.Context, in *MactimeEnableDisable) (*MactimeEnableDisableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpMactime(ctx context.Context, in *MactimeDump) (RPCService_DumpMactimeClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMactimeClient{stream}
	return x, nil
}

type RPCService_DumpMactimeClient interface {
	Recv() (*MactimeDetails, error)
}

type serviceClient_DumpMactimeClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMactimeClient) Recv() (*MactimeDetails, error) {
	m := new(MactimeDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) MactimeAddDelRange(ctx context.Context, in *MactimeAddDelRange) (*MactimeAddDelRangeReply, error) {
	out := new(MactimeAddDelRangeReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MactimeEnableDisable(ctx context.Context, in *MactimeEnableDisable) (*MactimeEnableDisableReply, error) {
	out := new(MactimeEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
