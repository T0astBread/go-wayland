// Generated by go-wayland-scanner
// https://github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.31/unstable/primary-selection/primary-selection-unstable-v1.xml?ref_type=tags
//
// wp_primary_selection_unstable_v1 Protocol Copyright:
//
// Copyright © 2015, 2016 Red Hat
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice (including the next
// paragraph) shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package primary_selection

import (
	"github.com/MatthiasKunnen/go-wayland/wayland/client"
	"golang.org/x/sys/unix"
)

// PrimarySelectionDeviceManagerInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const PrimarySelectionDeviceManagerInterfaceName = "zwp_primary_selection_device_manager_v1"

// PrimarySelectionDeviceManager : X primary selection emulation
//
// The primary selection device manager is a singleton global object that
// provides access to the primary selection. It allows to create
// wp_primary_selection_source objects, as well as retrieving the per-seat
// wp_primary_selection_device objects.
type PrimarySelectionDeviceManager struct {
	client.BaseProxy
}

// NewPrimarySelectionDeviceManager : X primary selection emulation
//
// The primary selection device manager is a singleton global object that
// provides access to the primary selection. It allows to create
// wp_primary_selection_source objects, as well as retrieving the per-seat
// wp_primary_selection_device objects.
func NewPrimarySelectionDeviceManager(ctx *client.Context) *PrimarySelectionDeviceManager {
	zwpPrimarySelectionDeviceManagerV1 := &PrimarySelectionDeviceManager{}
	ctx.Register(zwpPrimarySelectionDeviceManagerV1)
	return zwpPrimarySelectionDeviceManagerV1
}

// CreateSource : create a new primary selection source
//
// Create a new primary selection source.
func (i *PrimarySelectionDeviceManager) CreateSource() (*PrimarySelectionSource, error) {
	id := NewPrimarySelectionSource(i.Context())
	const opcode = 0
	const _reqBufLen = 8 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

// GetDevice : create a new primary selection device
//
// Create a new data device for a given seat.
func (i *PrimarySelectionDeviceManager) GetDevice(seat *client.Seat) (*PrimarySelectionDevice, error) {
	id := NewPrimarySelectionDevice(i.Context())
	const opcode = 1
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], seat.ID())
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

// Destroy : destroy the primary selection device manager
//
// Destroy the primary selection device manager.
func (i *PrimarySelectionDeviceManager) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 2
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// PrimarySelectionDeviceInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const PrimarySelectionDeviceInterfaceName = "zwp_primary_selection_device_v1"

// PrimarySelectionDevice :
type PrimarySelectionDevice struct {
	client.BaseProxy
	dataOfferHandler PrimarySelectionDeviceDataOfferHandlerFunc
	selectionHandler PrimarySelectionDeviceSelectionHandlerFunc
}

// NewPrimarySelectionDevice :
func NewPrimarySelectionDevice(ctx *client.Context) *PrimarySelectionDevice {
	zwpPrimarySelectionDeviceV1 := &PrimarySelectionDevice{}
	ctx.Register(zwpPrimarySelectionDeviceV1)
	return zwpPrimarySelectionDeviceV1
}

// SetSelection : set the primary selection
//
// Replaces the current selection. The previous owner of the primary
// selection will receive a wp_primary_selection_source.cancelled event.
//
// To unset the selection, set the source to NULL.
//
//	serial: serial of the event that triggered this request
func (i *PrimarySelectionDevice) SetSelection(source *PrimarySelectionSource, serial uint32) error {
	const opcode = 0
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	if source == nil {
		client.PutUint32(_reqBuf[l:l+4], 0)
		l += 4
	} else {
		client.PutUint32(_reqBuf[l:l+4], source.ID())
		l += 4
	}
	client.PutUint32(_reqBuf[l:l+4], uint32(serial))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// Destroy : destroy the primary selection device
//
// Destroy the primary selection device.
func (i *PrimarySelectionDevice) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 1
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// PrimarySelectionDeviceDataOfferEvent : introduce a new wp_primary_selection_offer
//
// Introduces a new wp_primary_selection_offer object that may be used
// to receive the current primary selection. Immediately following this
// event, the new wp_primary_selection_offer object will send
// wp_primary_selection_offer.offer events to describe the offered mime
// types.
type PrimarySelectionDeviceDataOfferEvent struct {
	Offer *PrimarySelectionOffer
}
type PrimarySelectionDeviceDataOfferHandlerFunc func(PrimarySelectionDeviceDataOfferEvent)

// SetDataOfferHandler : sets handler for PrimarySelectionDeviceDataOfferEvent
func (i *PrimarySelectionDevice) SetDataOfferHandler(f PrimarySelectionDeviceDataOfferHandlerFunc) {
	i.dataOfferHandler = f
}

// PrimarySelectionDeviceSelectionEvent : advertise a new primary selection
//
// The wp_primary_selection_device.selection event is sent to notify the
// client of a new primary selection. This event is sent after the
// wp_primary_selection.data_offer event introducing this object, and after
// the offer has announced its mimetypes through
// wp_primary_selection_offer.offer.
//
// The data_offer is valid until a new offer or NULL is received
// or until the client loses keyboard focus. The client must destroy the
// previous selection data_offer, if any, upon receiving this event.
type PrimarySelectionDeviceSelectionEvent struct {
	Id *PrimarySelectionOffer
}
type PrimarySelectionDeviceSelectionHandlerFunc func(PrimarySelectionDeviceSelectionEvent)

// SetSelectionHandler : sets handler for PrimarySelectionDeviceSelectionEvent
func (i *PrimarySelectionDevice) SetSelectionHandler(f PrimarySelectionDeviceSelectionHandlerFunc) {
	i.selectionHandler = f
}

func (i *PrimarySelectionDevice) Dispatch(opcode uint32, fd int, data []byte) {
	switch opcode {
	case 0:
		if i.dataOfferHandler == nil {
			return
		}
		var e PrimarySelectionDeviceDataOfferEvent
		l := 0
		e.Offer = i.Context().GetProxy(client.Uint32(data[l : l+4])).(*PrimarySelectionOffer)
		l += 4

		i.dataOfferHandler(e)
	case 1:
		if i.selectionHandler == nil {
			return
		}
		var e PrimarySelectionDeviceSelectionEvent
		l := 0
		e.Id = i.Context().GetProxy(client.Uint32(data[l : l+4])).(*PrimarySelectionOffer)
		l += 4

		i.selectionHandler(e)
	}
}

// PrimarySelectionOfferInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const PrimarySelectionOfferInterfaceName = "zwp_primary_selection_offer_v1"

// PrimarySelectionOffer : offer to transfer primary selection contents
//
// A wp_primary_selection_offer represents an offer to transfer the contents
// of the primary selection clipboard to the client. Similar to
// wl_data_offer, the offer also describes the mime types that the data can
// be converted to and provides the mechanisms for transferring the data
// directly to the client.
type PrimarySelectionOffer struct {
	client.BaseProxy
	offerHandler PrimarySelectionOfferOfferHandlerFunc
}

// NewPrimarySelectionOffer : offer to transfer primary selection contents
//
// A wp_primary_selection_offer represents an offer to transfer the contents
// of the primary selection clipboard to the client. Similar to
// wl_data_offer, the offer also describes the mime types that the data can
// be converted to and provides the mechanisms for transferring the data
// directly to the client.
func NewPrimarySelectionOffer(ctx *client.Context) *PrimarySelectionOffer {
	zwpPrimarySelectionOfferV1 := &PrimarySelectionOffer{}
	ctx.Register(zwpPrimarySelectionOfferV1)
	return zwpPrimarySelectionOfferV1
}

// Receive : request that the data is transferred
//
// To transfer the contents of the primary selection clipboard, the client
// issues this request and indicates the mime type that it wants to
// receive. The transfer happens through the passed file descriptor
// (typically created with the pipe system call). The source client writes
// the data in the mime type representation requested and then closes the
// file descriptor.
//
// The receiving client reads from the read end of the pipe until EOF and
// closes its end, at which point the transfer is complete.
func (i *PrimarySelectionOffer) Receive(mimeType string, fd int) error {
	const opcode = 0
	mimeTypeLen := client.PaddedLen(len(mimeType) + 1)
	_reqBufLen := 8 + (4 + mimeTypeLen)
	_reqBuf := make([]byte, _reqBufLen)
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutString(_reqBuf[l:l+(4+mimeTypeLen)], mimeType, mimeTypeLen)
	l += (4 + mimeTypeLen)
	oob := unix.UnixRights(int(fd))
	err := i.Context().WriteMsg(_reqBuf, oob)
	return err
}

// Destroy : destroy the primary selection offer
//
// Destroy the primary selection offer.
func (i *PrimarySelectionOffer) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 1
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// PrimarySelectionOfferOfferEvent : advertise offered mime type
//
// Sent immediately after creating announcing the
// wp_primary_selection_offer through
// wp_primary_selection_device.data_offer. One event is sent per offered
// mime type.
type PrimarySelectionOfferOfferEvent struct {
	MimeType string
}
type PrimarySelectionOfferOfferHandlerFunc func(PrimarySelectionOfferOfferEvent)

// SetOfferHandler : sets handler for PrimarySelectionOfferOfferEvent
func (i *PrimarySelectionOffer) SetOfferHandler(f PrimarySelectionOfferOfferHandlerFunc) {
	i.offerHandler = f
}

func (i *PrimarySelectionOffer) Dispatch(opcode uint32, fd int, data []byte) {
	switch opcode {
	case 0:
		if i.offerHandler == nil {
			return
		}
		var e PrimarySelectionOfferOfferEvent
		l := 0
		mimeTypeLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.MimeType = client.String(data[l : l+mimeTypeLen])
		l += mimeTypeLen

		i.offerHandler(e)
	}
}

// PrimarySelectionSourceInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const PrimarySelectionSourceInterfaceName = "zwp_primary_selection_source_v1"

// PrimarySelectionSource : offer to replace the contents of the primary selection
//
// The source side of a wp_primary_selection_offer, it provides a way to
// describe the offered data and respond to requests to transfer the
// requested contents of the primary selection clipboard.
type PrimarySelectionSource struct {
	client.BaseProxy
	sendHandler      PrimarySelectionSourceSendHandlerFunc
	cancelledHandler PrimarySelectionSourceCancelledHandlerFunc
}

// NewPrimarySelectionSource : offer to replace the contents of the primary selection
//
// The source side of a wp_primary_selection_offer, it provides a way to
// describe the offered data and respond to requests to transfer the
// requested contents of the primary selection clipboard.
func NewPrimarySelectionSource(ctx *client.Context) *PrimarySelectionSource {
	zwpPrimarySelectionSourceV1 := &PrimarySelectionSource{}
	ctx.Register(zwpPrimarySelectionSourceV1)
	return zwpPrimarySelectionSourceV1
}

// Offer : add an offered mime type
//
// This request adds a mime type to the set of mime types advertised to
// targets. Can be called several times to offer multiple types.
func (i *PrimarySelectionSource) Offer(mimeType string) error {
	const opcode = 0
	mimeTypeLen := client.PaddedLen(len(mimeType) + 1)
	_reqBufLen := 8 + (4 + mimeTypeLen)
	_reqBuf := make([]byte, _reqBufLen)
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutString(_reqBuf[l:l+(4+mimeTypeLen)], mimeType, mimeTypeLen)
	l += (4 + mimeTypeLen)
	err := i.Context().WriteMsg(_reqBuf, nil)
	return err
}

// Destroy : destroy the primary selection source
//
// Destroy the primary selection source.
func (i *PrimarySelectionSource) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 1
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// PrimarySelectionSourceSendEvent : send the primary selection contents
//
// Request for the current primary selection contents from the client.
// Send the specified mime type over the passed file descriptor, then
// close it.
type PrimarySelectionSourceSendEvent struct {
	MimeType string
	Fd       int
}
type PrimarySelectionSourceSendHandlerFunc func(PrimarySelectionSourceSendEvent)

// SetSendHandler : sets handler for PrimarySelectionSourceSendEvent
func (i *PrimarySelectionSource) SetSendHandler(f PrimarySelectionSourceSendHandlerFunc) {
	i.sendHandler = f
}

// PrimarySelectionSourceCancelledEvent : request for primary selection contents was canceled
//
// This primary selection source is no longer valid. The client should
// clean up and destroy this primary selection source.
type PrimarySelectionSourceCancelledEvent struct{}
type PrimarySelectionSourceCancelledHandlerFunc func(PrimarySelectionSourceCancelledEvent)

// SetCancelledHandler : sets handler for PrimarySelectionSourceCancelledEvent
func (i *PrimarySelectionSource) SetCancelledHandler(f PrimarySelectionSourceCancelledHandlerFunc) {
	i.cancelledHandler = f
}

func (i *PrimarySelectionSource) Dispatch(opcode uint32, fd int, data []byte) {
	switch opcode {
	case 0:
		if i.sendHandler == nil {
			if fd != -1 {
				unix.Close(fd)
			}
			return
		}
		var e PrimarySelectionSourceSendEvent
		l := 0
		mimeTypeLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.MimeType = client.String(data[l : l+mimeTypeLen])
		l += mimeTypeLen
		e.Fd = fd

		i.sendHandler(e)
	case 1:
		if i.cancelledHandler == nil {
			return
		}
		var e PrimarySelectionSourceCancelledEvent

		i.cancelledHandler(e)
	}
}
