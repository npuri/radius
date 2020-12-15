package rfc4679

import (
	"testing"

	"github.com/holgermetschulat/radius"
	"github.com/holgermetschulat/radius/rfc2866"
	"github.com/stretchr/testify/assert"
)

// _ADSLForum_SetVendor
func TestSetVendor(t *testing.T) {
	a := assert.New(t)
	packet := radius.New(radius.CodeAccessAccept, []byte("secret"))
	attr := radius.Attribute([]byte("asdf"))
	err := rfc2866.AcctStatusType_Set(packet, rfc2866.AcctStatusType_Value_Start)
	a.Nil(err)
	err = rfc2866.AcctInputOctets_Set(packet, rfc2866.AcctInputOctets(1))
	a.Nil(err)
	err = _ADSLForum_SetVendor(packet, 2, attr)
	a.Nil(err)
	err = _ADSLForum_SetVendor(packet, 1, attr)
	a.Nil(err)
	err = _ADSLForum_SetVendor(packet, 1, attr)
	a.Nil(err)
	adslAttribs := _ADSLForum_GetsVendor(packet, 1)
	a.Len(adslAttribs, 1)
	a.Len(packet.Attributes, 4)
}

func TestSetVendorNoStandardAttributes(t *testing.T) {
	a := assert.New(t)
	packet := radius.New(radius.CodeAccessAccept, []byte("secret"))
	attr := radius.Attribute([]byte("asdf"))
	err := _ADSLForum_SetVendor(packet, 1, attr)
	a.Nil(err)
	err = _ADSLForum_SetVendor(packet, 2, attr)
	a.Nil(err)
	err = _ADSLForum_SetVendor(packet, 1, attr)
	a.Nil(err)
	adslAttribs := _ADSLForum_GetsVendor(packet, 1)
	a.Len(adslAttribs, 1)
	a.Len(packet.Attributes, 2)
}

func TestSetVendorNoVendorAttributes(t *testing.T) {
	a := assert.New(t)
	packet := radius.New(radius.CodeAccessAccept, []byte("secret"))
	attr := radius.Attribute([]byte("asdf"))
	err := _ADSLForum_SetVendor(packet, 1, attr)
	a.Nil(err)
	err = _ADSLForum_SetVendor(packet, 1, attr)
	a.Nil(err)
	adslAttribs := _ADSLForum_GetsVendor(packet, 1)
	a.Len(adslAttribs, 1)
	a.Len(packet.Attributes, 1)
}
