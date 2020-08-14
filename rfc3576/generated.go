// Code generated by radius-dict-gen. DO NOT EDIT.

package rfc3576

import (
	"strconv"

	"github.com/holgermetschulat/radius"

	. "github.com/holgermetschulat/radius/rfc2865"
)

const (
	ErrorCause_Type radius.Type = 101
)

func init() {
	ServiceType_Strings[ServiceType_Value_AuthorizeOnly] = "Authorize-Only"
}

const (
	ServiceType_Value_AuthorizeOnly ServiceType = 17
)

type ErrorCause uint32

const (
	ErrorCause_Value_ResidualContextRemoved     ErrorCause = 201
	ErrorCause_Value_InvalidEAPPacket           ErrorCause = 202
	ErrorCause_Value_UnsupportedAttribute       ErrorCause = 401
	ErrorCause_Value_MissingAttribute           ErrorCause = 402
	ErrorCause_Value_NASIdentificationMismatch  ErrorCause = 403
	ErrorCause_Value_InvalidRequest             ErrorCause = 404
	ErrorCause_Value_UnsupportedService         ErrorCause = 405
	ErrorCause_Value_UnsupportedExtension       ErrorCause = 406
	ErrorCause_Value_AdministrativelyProhibited ErrorCause = 501
	ErrorCause_Value_ProxyRequestNotRoutable    ErrorCause = 502
	ErrorCause_Value_SessionContextNotFound     ErrorCause = 503
	ErrorCause_Value_SessionContextNotRemovable ErrorCause = 504
	ErrorCause_Value_ProxyProcessingError       ErrorCause = 505
	ErrorCause_Value_ResourcesUnavailable       ErrorCause = 506
	ErrorCause_Value_RequestInitiated           ErrorCause = 507
)

var ErrorCause_Strings = map[ErrorCause]string{
	ErrorCause_Value_ResidualContextRemoved:     "Residual-Context-Removed",
	ErrorCause_Value_InvalidEAPPacket:           "Invalid-EAP-Packet",
	ErrorCause_Value_UnsupportedAttribute:       "Unsupported-Attribute",
	ErrorCause_Value_MissingAttribute:           "Missing-Attribute",
	ErrorCause_Value_NASIdentificationMismatch:  "NAS-Identification-Mismatch",
	ErrorCause_Value_InvalidRequest:             "Invalid-Request",
	ErrorCause_Value_UnsupportedService:         "Unsupported-Service",
	ErrorCause_Value_UnsupportedExtension:       "Unsupported-Extension",
	ErrorCause_Value_AdministrativelyProhibited: "Administratively-Prohibited",
	ErrorCause_Value_ProxyRequestNotRoutable:    "Proxy-Request-Not-Routable",
	ErrorCause_Value_SessionContextNotFound:     "Session-Context-Not-Found",
	ErrorCause_Value_SessionContextNotRemovable: "Session-Context-Not-Removable",
	ErrorCause_Value_ProxyProcessingError:       "Proxy-Processing-Error",
	ErrorCause_Value_ResourcesUnavailable:       "Resources-Unavailable",
	ErrorCause_Value_RequestInitiated:           "Request-Initiated",
}

func (a ErrorCause) String() string {
	if str, ok := ErrorCause_Strings[a]; ok {
		return str
	}
	return "ErrorCause(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func ErrorCause_Add(p *radius.Packet, value ErrorCause) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(ErrorCause_Type, a)
	return
}

func ErrorCause_Get(p *radius.Packet) (value ErrorCause) {
	value, _ = ErrorCause_Lookup(p)
	return
}

func ErrorCause_Gets(p *radius.Packet) (values []ErrorCause, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != ErrorCause_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, ErrorCause(i))
	}
	return
}

func ErrorCause_Lookup(p *radius.Packet) (value ErrorCause, err error) {
	a, ok := p.Lookup(ErrorCause_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = ErrorCause(i)
	return
}

func ErrorCause_Set(p *radius.Packet, value ErrorCause) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(ErrorCause_Type, a)
	return
}

func ErrorCause_Del(p *radius.Packet) {
	p.Attributes.Del(ErrorCause_Type)
}
