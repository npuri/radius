// Code generated by radius-dict-gen. DO NOT EDIT.

package rfc2866

import (
	"strconv"

	"github.com/holgermetschulat/radius"
)

const (
	AcctStatusType_Type     radius.Type = 40
	AcctDelayTime_Type      radius.Type = 41
	AcctInputOctets_Type    radius.Type = 42
	AcctOutputOctets_Type   radius.Type = 43
	AcctSessionID_Type      radius.Type = 44
	AcctAuthentic_Type      radius.Type = 45
	AcctSessionTime_Type    radius.Type = 46
	AcctInputPackets_Type   radius.Type = 47
	AcctOutputPackets_Type  radius.Type = 48
	AcctTerminateCause_Type radius.Type = 49
	AcctMultiSessionID_Type radius.Type = 50
	AcctLinkCount_Type      radius.Type = 51
)

type AcctStatusType uint32

const (
	AcctStatusType_Value_Start         AcctStatusType = 1
	AcctStatusType_Value_Stop          AcctStatusType = 2
	AcctStatusType_Value_InterimUpdate AcctStatusType = 3
	AcctStatusType_Value_AccountingOn  AcctStatusType = 7
	AcctStatusType_Value_AccountingOff AcctStatusType = 8
	AcctStatusType_Value_Failed        AcctStatusType = 15
)

var AcctStatusType_Strings = map[AcctStatusType]string{
	AcctStatusType_Value_Start:         "Start",
	AcctStatusType_Value_Stop:          "Stop",
	AcctStatusType_Value_InterimUpdate: "Interim-Update",
	AcctStatusType_Value_AccountingOn:  "Accounting-On",
	AcctStatusType_Value_AccountingOff: "Accounting-Off",
	AcctStatusType_Value_Failed:        "Failed",
}

func (a AcctStatusType) String() string {
	if str, ok := AcctStatusType_Strings[a]; ok {
		return str
	}
	return "AcctStatusType(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctStatusType_Add(p *radius.Packet, value AcctStatusType) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctStatusType_Type, a)
	return
}

func AcctStatusType_Get(p *radius.Packet) (value AcctStatusType) {
	value, _ = AcctStatusType_Lookup(p)
	return
}

func AcctStatusType_Gets(p *radius.Packet) (values []AcctStatusType, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != AcctStatusType_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctStatusType(i))
	}
	return
}

func AcctStatusType_Lookup(p *radius.Packet) (value AcctStatusType, err error) {
	a, ok := p.Lookup(AcctStatusType_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctStatusType(i)
	return
}

func AcctStatusType_Set(p *radius.Packet, value AcctStatusType) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctStatusType_Type, a)
	return
}

func AcctStatusType_Del(p *radius.Packet) {
	p.Attributes.Del(AcctStatusType_Type)
}

type AcctDelayTime uint32

var AcctDelayTime_Strings = map[AcctDelayTime]string{}

func (a AcctDelayTime) String() string {
	if str, ok := AcctDelayTime_Strings[a]; ok {
		return str
	}
	return "AcctDelayTime(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctDelayTime_Add(p *radius.Packet, value AcctDelayTime) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctDelayTime_Type, a)
	return
}

func AcctDelayTime_Get(p *radius.Packet) (value AcctDelayTime) {
	value, _ = AcctDelayTime_Lookup(p)
	return
}

func AcctDelayTime_Gets(p *radius.Packet) (values []AcctDelayTime, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != AcctDelayTime_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctDelayTime(i))
	}
	return
}

func AcctDelayTime_Lookup(p *radius.Packet) (value AcctDelayTime, err error) {
	a, ok := p.Lookup(AcctDelayTime_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctDelayTime(i)
	return
}

func AcctDelayTime_Set(p *radius.Packet, value AcctDelayTime) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctDelayTime_Type, a)
	return
}

func AcctDelayTime_Del(p *radius.Packet) {
	p.Attributes.Del(AcctDelayTime_Type)
}

type AcctInputOctets uint32

var AcctInputOctets_Strings = map[AcctInputOctets]string{}

func (a AcctInputOctets) String() string {
	if str, ok := AcctInputOctets_Strings[a]; ok {
		return str
	}
	return "AcctInputOctets(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctInputOctets_Add(p *radius.Packet, value AcctInputOctets) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctInputOctets_Type, a)
	return
}

func AcctInputOctets_Get(p *radius.Packet) (value AcctInputOctets) {
	value, _ = AcctInputOctets_Lookup(p)
	return
}

func AcctInputOctets_Gets(p *radius.Packet) (values []AcctInputOctets, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != AcctInputOctets_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctInputOctets(i))
	}
	return
}

func AcctInputOctets_Lookup(p *radius.Packet) (value AcctInputOctets, err error) {
	a, ok := p.Lookup(AcctInputOctets_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctInputOctets(i)
	return
}

func AcctInputOctets_Set(p *radius.Packet, value AcctInputOctets) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctInputOctets_Type, a)
	return
}

func AcctInputOctets_Del(p *radius.Packet) {
	p.Attributes.Del(AcctInputOctets_Type)
}

type AcctOutputOctets uint32

var AcctOutputOctets_Strings = map[AcctOutputOctets]string{}

func (a AcctOutputOctets) String() string {
	if str, ok := AcctOutputOctets_Strings[a]; ok {
		return str
	}
	return "AcctOutputOctets(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctOutputOctets_Add(p *radius.Packet, value AcctOutputOctets) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctOutputOctets_Type, a)
	return
}

func AcctOutputOctets_Get(p *radius.Packet) (value AcctOutputOctets) {
	value, _ = AcctOutputOctets_Lookup(p)
	return
}

func AcctOutputOctets_Gets(p *radius.Packet) (values []AcctOutputOctets, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != AcctOutputOctets_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctOutputOctets(i))
	}
	return
}

func AcctOutputOctets_Lookup(p *radius.Packet) (value AcctOutputOctets, err error) {
	a, ok := p.Lookup(AcctOutputOctets_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctOutputOctets(i)
	return
}

func AcctOutputOctets_Set(p *radius.Packet, value AcctOutputOctets) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctOutputOctets_Type, a)
	return
}

func AcctOutputOctets_Del(p *radius.Packet) {
	p.Attributes.Del(AcctOutputOctets_Type)
}

func AcctSessionID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(AcctSessionID_Type, a)
	return
}

func AcctSessionID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(AcctSessionID_Type, a)
	return
}

func AcctSessionID_Get(p *radius.Packet) (value []byte) {
	value, _ = AcctSessionID_Lookup(p)
	return
}

func AcctSessionID_GetString(p *radius.Packet) (value string) {
	value, _ = AcctSessionID_LookupString(p)
	return
}

func AcctSessionID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, avp := range p.Attributes {
		if avp.Type != AcctSessionID_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AcctSessionID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, avp := range p.Attributes {
		if avp.Type != AcctSessionID_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AcctSessionID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(AcctSessionID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func AcctSessionID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(AcctSessionID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func AcctSessionID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(AcctSessionID_Type, a)
	return
}

func AcctSessionID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(AcctSessionID_Type, a)
	return
}

func AcctSessionID_Del(p *radius.Packet) {
	p.Attributes.Del(AcctSessionID_Type)
}

type AcctAuthentic uint32

const (
	AcctAuthentic_Value_RADIUS   AcctAuthentic = 1
	AcctAuthentic_Value_Local    AcctAuthentic = 2
	AcctAuthentic_Value_Remote   AcctAuthentic = 3
	AcctAuthentic_Value_Diameter AcctAuthentic = 4
)

var AcctAuthentic_Strings = map[AcctAuthentic]string{
	AcctAuthentic_Value_RADIUS:   "RADIUS",
	AcctAuthentic_Value_Local:    "Local",
	AcctAuthentic_Value_Remote:   "Remote",
	AcctAuthentic_Value_Diameter: "Diameter",
}

func (a AcctAuthentic) String() string {
	if str, ok := AcctAuthentic_Strings[a]; ok {
		return str
	}
	return "AcctAuthentic(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctAuthentic_Add(p *radius.Packet, value AcctAuthentic) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctAuthentic_Type, a)
	return
}

func AcctAuthentic_Get(p *radius.Packet) (value AcctAuthentic) {
	value, _ = AcctAuthentic_Lookup(p)
	return
}

func AcctAuthentic_Gets(p *radius.Packet) (values []AcctAuthentic, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != AcctAuthentic_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctAuthentic(i))
	}
	return
}

func AcctAuthentic_Lookup(p *radius.Packet) (value AcctAuthentic, err error) {
	a, ok := p.Lookup(AcctAuthentic_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctAuthentic(i)
	return
}

func AcctAuthentic_Set(p *radius.Packet, value AcctAuthentic) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctAuthentic_Type, a)
	return
}

func AcctAuthentic_Del(p *radius.Packet) {
	p.Attributes.Del(AcctAuthentic_Type)
}

type AcctSessionTime uint32

var AcctSessionTime_Strings = map[AcctSessionTime]string{}

func (a AcctSessionTime) String() string {
	if str, ok := AcctSessionTime_Strings[a]; ok {
		return str
	}
	return "AcctSessionTime(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctSessionTime_Add(p *radius.Packet, value AcctSessionTime) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctSessionTime_Type, a)
	return
}

func AcctSessionTime_Get(p *radius.Packet) (value AcctSessionTime) {
	value, _ = AcctSessionTime_Lookup(p)
	return
}

func AcctSessionTime_Gets(p *radius.Packet) (values []AcctSessionTime, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != AcctSessionTime_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctSessionTime(i))
	}
	return
}

func AcctSessionTime_Lookup(p *radius.Packet) (value AcctSessionTime, err error) {
	a, ok := p.Lookup(AcctSessionTime_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctSessionTime(i)
	return
}

func AcctSessionTime_Set(p *radius.Packet, value AcctSessionTime) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctSessionTime_Type, a)
	return
}

func AcctSessionTime_Del(p *radius.Packet) {
	p.Attributes.Del(AcctSessionTime_Type)
}

type AcctInputPackets uint32

var AcctInputPackets_Strings = map[AcctInputPackets]string{}

func (a AcctInputPackets) String() string {
	if str, ok := AcctInputPackets_Strings[a]; ok {
		return str
	}
	return "AcctInputPackets(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctInputPackets_Add(p *radius.Packet, value AcctInputPackets) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctInputPackets_Type, a)
	return
}

func AcctInputPackets_Get(p *radius.Packet) (value AcctInputPackets) {
	value, _ = AcctInputPackets_Lookup(p)
	return
}

func AcctInputPackets_Gets(p *radius.Packet) (values []AcctInputPackets, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != AcctInputPackets_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctInputPackets(i))
	}
	return
}

func AcctInputPackets_Lookup(p *radius.Packet) (value AcctInputPackets, err error) {
	a, ok := p.Lookup(AcctInputPackets_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctInputPackets(i)
	return
}

func AcctInputPackets_Set(p *radius.Packet, value AcctInputPackets) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctInputPackets_Type, a)
	return
}

func AcctInputPackets_Del(p *radius.Packet) {
	p.Attributes.Del(AcctInputPackets_Type)
}

type AcctOutputPackets uint32

var AcctOutputPackets_Strings = map[AcctOutputPackets]string{}

func (a AcctOutputPackets) String() string {
	if str, ok := AcctOutputPackets_Strings[a]; ok {
		return str
	}
	return "AcctOutputPackets(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctOutputPackets_Add(p *radius.Packet, value AcctOutputPackets) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctOutputPackets_Type, a)
	return
}

func AcctOutputPackets_Get(p *radius.Packet) (value AcctOutputPackets) {
	value, _ = AcctOutputPackets_Lookup(p)
	return
}

func AcctOutputPackets_Gets(p *radius.Packet) (values []AcctOutputPackets, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != AcctOutputPackets_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctOutputPackets(i))
	}
	return
}

func AcctOutputPackets_Lookup(p *radius.Packet) (value AcctOutputPackets, err error) {
	a, ok := p.Lookup(AcctOutputPackets_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctOutputPackets(i)
	return
}

func AcctOutputPackets_Set(p *radius.Packet, value AcctOutputPackets) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctOutputPackets_Type, a)
	return
}

func AcctOutputPackets_Del(p *radius.Packet) {
	p.Attributes.Del(AcctOutputPackets_Type)
}

type AcctTerminateCause uint32

const (
	AcctTerminateCause_Value_UserRequest        AcctTerminateCause = 1
	AcctTerminateCause_Value_LostCarrier        AcctTerminateCause = 2
	AcctTerminateCause_Value_LostService        AcctTerminateCause = 3
	AcctTerminateCause_Value_IdleTimeout        AcctTerminateCause = 4
	AcctTerminateCause_Value_SessionTimeout     AcctTerminateCause = 5
	AcctTerminateCause_Value_AdminReset         AcctTerminateCause = 6
	AcctTerminateCause_Value_AdminReboot        AcctTerminateCause = 7
	AcctTerminateCause_Value_PortError          AcctTerminateCause = 8
	AcctTerminateCause_Value_NASError           AcctTerminateCause = 9
	AcctTerminateCause_Value_NASRequest         AcctTerminateCause = 10
	AcctTerminateCause_Value_NASReboot          AcctTerminateCause = 11
	AcctTerminateCause_Value_PortUnneeded       AcctTerminateCause = 12
	AcctTerminateCause_Value_PortPreempted      AcctTerminateCause = 13
	AcctTerminateCause_Value_PortSuspended      AcctTerminateCause = 14
	AcctTerminateCause_Value_ServiceUnavailable AcctTerminateCause = 15
	AcctTerminateCause_Value_Callback           AcctTerminateCause = 16
	AcctTerminateCause_Value_UserError          AcctTerminateCause = 17
	AcctTerminateCause_Value_HostRequest        AcctTerminateCause = 18
)

var AcctTerminateCause_Strings = map[AcctTerminateCause]string{
	AcctTerminateCause_Value_UserRequest:        "User-Request",
	AcctTerminateCause_Value_LostCarrier:        "Lost-Carrier",
	AcctTerminateCause_Value_LostService:        "Lost-Service",
	AcctTerminateCause_Value_IdleTimeout:        "Idle-Timeout",
	AcctTerminateCause_Value_SessionTimeout:     "Session-Timeout",
	AcctTerminateCause_Value_AdminReset:         "Admin-Reset",
	AcctTerminateCause_Value_AdminReboot:        "Admin-Reboot",
	AcctTerminateCause_Value_PortError:          "Port-Error",
	AcctTerminateCause_Value_NASError:           "NAS-Error",
	AcctTerminateCause_Value_NASRequest:         "NAS-Request",
	AcctTerminateCause_Value_NASReboot:          "NAS-Reboot",
	AcctTerminateCause_Value_PortUnneeded:       "Port-Unneeded",
	AcctTerminateCause_Value_PortPreempted:      "Port-Preempted",
	AcctTerminateCause_Value_PortSuspended:      "Port-Suspended",
	AcctTerminateCause_Value_ServiceUnavailable: "Service-Unavailable",
	AcctTerminateCause_Value_Callback:           "Callback",
	AcctTerminateCause_Value_UserError:          "User-Error",
	AcctTerminateCause_Value_HostRequest:        "Host-Request",
}

func (a AcctTerminateCause) String() string {
	if str, ok := AcctTerminateCause_Strings[a]; ok {
		return str
	}
	return "AcctTerminateCause(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctTerminateCause_Add(p *radius.Packet, value AcctTerminateCause) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctTerminateCause_Type, a)
	return
}

func AcctTerminateCause_Get(p *radius.Packet) (value AcctTerminateCause) {
	value, _ = AcctTerminateCause_Lookup(p)
	return
}

func AcctTerminateCause_Gets(p *radius.Packet) (values []AcctTerminateCause, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != AcctTerminateCause_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctTerminateCause(i))
	}
	return
}

func AcctTerminateCause_Lookup(p *radius.Packet) (value AcctTerminateCause, err error) {
	a, ok := p.Lookup(AcctTerminateCause_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctTerminateCause(i)
	return
}

func AcctTerminateCause_Set(p *radius.Packet, value AcctTerminateCause) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctTerminateCause_Type, a)
	return
}

func AcctTerminateCause_Del(p *radius.Packet) {
	p.Attributes.Del(AcctTerminateCause_Type)
}

func AcctMultiSessionID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(AcctMultiSessionID_Type, a)
	return
}

func AcctMultiSessionID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(AcctMultiSessionID_Type, a)
	return
}

func AcctMultiSessionID_Get(p *radius.Packet) (value []byte) {
	value, _ = AcctMultiSessionID_Lookup(p)
	return
}

func AcctMultiSessionID_GetString(p *radius.Packet) (value string) {
	value, _ = AcctMultiSessionID_LookupString(p)
	return
}

func AcctMultiSessionID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, avp := range p.Attributes {
		if avp.Type != AcctMultiSessionID_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AcctMultiSessionID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, avp := range p.Attributes {
		if avp.Type != AcctMultiSessionID_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AcctMultiSessionID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(AcctMultiSessionID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func AcctMultiSessionID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(AcctMultiSessionID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func AcctMultiSessionID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(AcctMultiSessionID_Type, a)
	return
}

func AcctMultiSessionID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(AcctMultiSessionID_Type, a)
	return
}

func AcctMultiSessionID_Del(p *radius.Packet) {
	p.Attributes.Del(AcctMultiSessionID_Type)
}

type AcctLinkCount uint32

var AcctLinkCount_Strings = map[AcctLinkCount]string{}

func (a AcctLinkCount) String() string {
	if str, ok := AcctLinkCount_Strings[a]; ok {
		return str
	}
	return "AcctLinkCount(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctLinkCount_Add(p *radius.Packet, value AcctLinkCount) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctLinkCount_Type, a)
	return
}

func AcctLinkCount_Get(p *radius.Packet) (value AcctLinkCount) {
	value, _ = AcctLinkCount_Lookup(p)
	return
}

func AcctLinkCount_Gets(p *radius.Packet) (values []AcctLinkCount, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != AcctLinkCount_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctLinkCount(i))
	}
	return
}

func AcctLinkCount_Lookup(p *radius.Packet) (value AcctLinkCount, err error) {
	a, ok := p.Lookup(AcctLinkCount_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctLinkCount(i)
	return
}

func AcctLinkCount_Set(p *radius.Packet, value AcctLinkCount) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctLinkCount_Type, a)
	return
}

func AcctLinkCount_Del(p *radius.Packet) {
	p.Attributes.Del(AcctLinkCount_Type)
}
