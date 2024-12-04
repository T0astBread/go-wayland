// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.31/staging/fractional-scale/fractional-scale-v1.xml?ref_type=tags
//
// fractional_scale_v1 Protocol Copyright:
//
// Copyright © 2022 Kenny Levinsen
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

package fractional_scale

import "github.com/rajveermalviya/go-wayland/wayland/client"

// FractionalScaleManagerInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const FractionalScaleManagerInterfaceName = "wp_fractional_scale_manager_v1"

// FractionalScaleManager : fractional surface scale information
//
// A global interface for requesting surfaces to use fractional scales.
type FractionalScaleManager struct {
	client.BaseProxy
}

// NewFractionalScaleManager : fractional surface scale information
//
// A global interface for requesting surfaces to use fractional scales.
func NewFractionalScaleManager(ctx *client.Context) *FractionalScaleManager {
	wpFractionalScaleManagerV1 := &FractionalScaleManager{}
	ctx.Register(wpFractionalScaleManagerV1)
	return wpFractionalScaleManagerV1
}

// Destroy : unbind the fractional surface scale interface
//
// Informs the server that the client will not be using this protocol
// object anymore. This does not affect any other objects,
// wp_fractional_scale_v1 objects included.
func (i *FractionalScaleManager) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
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

// GetFractionalScale : extend surface interface for scale information
//
// Create an add-on object for the the wl_surface to let the compositor
// request fractional scales. If the given wl_surface already has a
// wp_fractional_scale_v1 object associated, the fractional_scale_exists
// protocol error is raised.
//
//	surface: the surface
func (i *FractionalScaleManager) GetFractionalScale(surface *client.Surface) (*FractionalScale, error) {
	id := NewFractionalScale(i.Context())
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
	client.PutUint32(_reqBuf[l:l+4], surface.ID())
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

type FractionalScaleManagerError uint32

// FractionalScaleManagerError :
const (
	// FractionalScaleManagerErrorFractionalScaleExists : the surface already has a fractional_scale object associated
	FractionalScaleManagerErrorFractionalScaleExists FractionalScaleManagerError = 0
)

func (e FractionalScaleManagerError) Name() string {
	switch e {
	case FractionalScaleManagerErrorFractionalScaleExists:
		return "fractional_scale_exists"
	default:
		return ""
	}
}

func (e FractionalScaleManagerError) Value() string {
	switch e {
	case FractionalScaleManagerErrorFractionalScaleExists:
		return "0"
	default:
		return ""
	}
}

func (e FractionalScaleManagerError) String() string {
	return e.Name() + "=" + e.Value()
}

// FractionalScaleInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const FractionalScaleInterfaceName = "wp_fractional_scale_v1"

// FractionalScale : fractional scale interface to a wl_surface
//
// An additional interface to a wl_surface object which allows the compositor
// to inform the client of the preferred scale.
type FractionalScale struct {
	client.BaseProxy
	preferredScaleHandler FractionalScalePreferredScaleHandlerFunc
}

// NewFractionalScale : fractional scale interface to a wl_surface
//
// An additional interface to a wl_surface object which allows the compositor
// to inform the client of the preferred scale.
func NewFractionalScale(ctx *client.Context) *FractionalScale {
	wpFractionalScaleV1 := &FractionalScale{}
	ctx.Register(wpFractionalScaleV1)
	return wpFractionalScaleV1
}

// Destroy : remove surface scale information for surface
//
// Destroy the fractional scale object. When this object is destroyed,
// preferred_scale events will no longer be sent.
func (i *FractionalScale) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
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

// FractionalScalePreferredScaleEvent : notify of new preferred scale
//
// Notification of a new preferred scale for this surface that the
// compositor suggests that the client should use.
//
// The sent scale is the numerator of a fraction with a denominator of 120.
type FractionalScalePreferredScaleEvent struct {
	Scale uint32
}
type FractionalScalePreferredScaleHandlerFunc func(FractionalScalePreferredScaleEvent)

// SetPreferredScaleHandler : sets handler for FractionalScalePreferredScaleEvent
func (i *FractionalScale) SetPreferredScaleHandler(f FractionalScalePreferredScaleHandlerFunc) {
	i.preferredScaleHandler = f
}

func (i *FractionalScale) Dispatch(opcode uint32, fd int, data []byte) {
	switch opcode {
	case 0:
		if i.preferredScaleHandler == nil {
			return
		}
		var e FractionalScalePreferredScaleEvent
		l := 0
		e.Scale = client.Uint32(data[l : l+4])
		l += 4

		i.preferredScaleHandler(e)
	}
}
