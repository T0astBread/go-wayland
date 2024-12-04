// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.31/stable/viewporter/viewporter.xml?ref_type=tags
//
// viewporter Protocol Copyright:
//
// Copyright © 2013-2016 Collabora, Ltd.
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

package viewporter

import "github.com/rajveermalviya/go-wayland/wayland/client"

// ViewporterInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const ViewporterInterfaceName = "wp_viewporter"

// Viewporter : surface cropping and scaling
//
// The global interface exposing surface cropping and scaling
// capabilities is used to instantiate an interface extension for a
// wl_surface object. This extended interface will then allow
// cropping and scaling the surface contents, effectively
// disconnecting the direct relationship between the buffer and the
// surface size.
type Viewporter struct {
	client.BaseProxy
}

// NewViewporter : surface cropping and scaling
//
// The global interface exposing surface cropping and scaling
// capabilities is used to instantiate an interface extension for a
// wl_surface object. This extended interface will then allow
// cropping and scaling the surface contents, effectively
// disconnecting the direct relationship between the buffer and the
// surface size.
func NewViewporter(ctx *client.Context) *Viewporter {
	wpViewporter := &Viewporter{}
	ctx.Register(wpViewporter)
	return wpViewporter
}

// Destroy : unbind from the cropping and scaling interface
//
// Informs the server that the client will not be using this
// protocol object anymore. This does not affect any other objects,
// wp_viewport objects included.
func (i *Viewporter) Destroy() error {
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

// GetViewport : extend surface interface for crop and scale
//
// Instantiate an interface extension for the given wl_surface to
// crop and scale its content. If the given wl_surface already has
// a wp_viewport object associated, the viewport_exists
// protocol error is raised.
//
//	surface: the surface
func (i *Viewporter) GetViewport(surface *client.Surface) (*Viewport, error) {
	id := NewViewport(i.Context())
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

type ViewporterError uint32

// ViewporterError :
const (
	// ViewporterErrorViewportExists : the surface already has a viewport object associated
	ViewporterErrorViewportExists ViewporterError = 0
)

func (e ViewporterError) Name() string {
	switch e {
	case ViewporterErrorViewportExists:
		return "viewport_exists"
	default:
		return ""
	}
}

func (e ViewporterError) Value() string {
	switch e {
	case ViewporterErrorViewportExists:
		return "0"
	default:
		return ""
	}
}

func (e ViewporterError) String() string {
	return e.Name() + "=" + e.Value()
}

// ViewportInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const ViewportInterfaceName = "wp_viewport"

// Viewport : crop and scale interface to a wl_surface
//
// An additional interface to a wl_surface object, which allows the
// client to specify the cropping and scaling of the surface
// contents.
//
// This interface works with two concepts: the source rectangle (src_x,
// src_y, src_width, src_height), and the destination size (dst_width,
// dst_height). The contents of the source rectangle are scaled to the
// destination size, and content outside the source rectangle is ignored.
// This state is double-buffered, and is applied on the next
// wl_surface.commit.
//
// The two parts of crop and scale state are independent: the source
// rectangle, and the destination size. Initially both are unset, that
// is, no scaling is applied. The whole of the current wl_buffer is
// used as the source, and the surface size is as defined in
// wl_surface.attach.
//
// If the destination size is set, it causes the surface size to become
// dst_width, dst_height. The source (rectangle) is scaled to exactly
// this size. This overrides whatever the attached wl_buffer size is,
// unless the wl_buffer is NULL. If the wl_buffer is NULL, the surface
// has no content and therefore no size. Otherwise, the size is always
// at least 1x1 in surface local coordinates.
//
// If the source rectangle is set, it defines what area of the wl_buffer is
// taken as the source. If the source rectangle is set and the destination
// size is not set, then src_width and src_height must be integers, and the
// surface size becomes the source rectangle size. This results in cropping
// without scaling. If src_width or src_height are not integers and
// destination size is not set, the bad_size protocol error is raised when
// the surface state is applied.
//
// The coordinate transformations from buffer pixel coordinates up to
// the surface-local coordinates happen in the following order:
// 1. buffer_transform (wl_surface.set_buffer_transform)
// 2. buffer_scale (wl_surface.set_buffer_scale)
// 3. crop and scale (wp_viewport.set*)
// This means, that the source rectangle coordinates of crop and scale
// are given in the coordinates after the buffer transform and scale,
// i.e. in the coordinates that would be the surface-local coordinates
// if the crop and scale was not applied.
//
// If src_x or src_y are negative, the bad_value protocol error is raised.
// Otherwise, if the source rectangle is partially or completely outside of
// the non-NULL wl_buffer, then the out_of_buffer protocol error is raised
// when the surface state is applied. A NULL wl_buffer does not raise the
// out_of_buffer error.
//
// If the wl_surface associated with the wp_viewport is destroyed,
// all wp_viewport requests except 'destroy' raise the protocol error
// no_surface.
//
// If the wp_viewport object is destroyed, the crop and scale
// state is removed from the wl_surface. The change will be applied
// on the next wl_surface.commit.
type Viewport struct {
	client.BaseProxy
}

// NewViewport : crop and scale interface to a wl_surface
//
// An additional interface to a wl_surface object, which allows the
// client to specify the cropping and scaling of the surface
// contents.
//
// This interface works with two concepts: the source rectangle (src_x,
// src_y, src_width, src_height), and the destination size (dst_width,
// dst_height). The contents of the source rectangle are scaled to the
// destination size, and content outside the source rectangle is ignored.
// This state is double-buffered, and is applied on the next
// wl_surface.commit.
//
// The two parts of crop and scale state are independent: the source
// rectangle, and the destination size. Initially both are unset, that
// is, no scaling is applied. The whole of the current wl_buffer is
// used as the source, and the surface size is as defined in
// wl_surface.attach.
//
// If the destination size is set, it causes the surface size to become
// dst_width, dst_height. The source (rectangle) is scaled to exactly
// this size. This overrides whatever the attached wl_buffer size is,
// unless the wl_buffer is NULL. If the wl_buffer is NULL, the surface
// has no content and therefore no size. Otherwise, the size is always
// at least 1x1 in surface local coordinates.
//
// If the source rectangle is set, it defines what area of the wl_buffer is
// taken as the source. If the source rectangle is set and the destination
// size is not set, then src_width and src_height must be integers, and the
// surface size becomes the source rectangle size. This results in cropping
// without scaling. If src_width or src_height are not integers and
// destination size is not set, the bad_size protocol error is raised when
// the surface state is applied.
//
// The coordinate transformations from buffer pixel coordinates up to
// the surface-local coordinates happen in the following order:
// 1. buffer_transform (wl_surface.set_buffer_transform)
// 2. buffer_scale (wl_surface.set_buffer_scale)
// 3. crop and scale (wp_viewport.set*)
// This means, that the source rectangle coordinates of crop and scale
// are given in the coordinates after the buffer transform and scale,
// i.e. in the coordinates that would be the surface-local coordinates
// if the crop and scale was not applied.
//
// If src_x or src_y are negative, the bad_value protocol error is raised.
// Otherwise, if the source rectangle is partially or completely outside of
// the non-NULL wl_buffer, then the out_of_buffer protocol error is raised
// when the surface state is applied. A NULL wl_buffer does not raise the
// out_of_buffer error.
//
// If the wl_surface associated with the wp_viewport is destroyed,
// all wp_viewport requests except 'destroy' raise the protocol error
// no_surface.
//
// If the wp_viewport object is destroyed, the crop and scale
// state is removed from the wl_surface. The change will be applied
// on the next wl_surface.commit.
func NewViewport(ctx *client.Context) *Viewport {
	wpViewport := &Viewport{}
	ctx.Register(wpViewport)
	return wpViewport
}

// Destroy : remove scaling and cropping from the surface
//
// The associated wl_surface's crop and scale state is removed.
// The change is applied on the next wl_surface.commit.
func (i *Viewport) Destroy() error {
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

// SetSource : set the source rectangle for cropping
//
// Set the source rectangle of the associated wl_surface. See
// wp_viewport for the description, and relation to the wl_buffer
// size.
//
// If all of x, y, width and height are -1.0, the source rectangle is
// unset instead. Any other set of values where width or height are zero
// or negative, or x or y are negative, raise the bad_value protocol
// error.
//
// The crop and scale state is double-buffered state, and will be
// applied on the next wl_surface.commit.
//
//	x: source rectangle x
//	y: source rectangle y
//	width: source rectangle width
//	height: source rectangle height
func (i *Viewport) SetSource(x, y, width, height float64) error {
	const opcode = 1
	const _reqBufLen = 8 + 4 + 4 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutFixed(_reqBuf[l:l+4], x)
	l += 4
	client.PutFixed(_reqBuf[l:l+4], y)
	l += 4
	client.PutFixed(_reqBuf[l:l+4], width)
	l += 4
	client.PutFixed(_reqBuf[l:l+4], height)
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// SetDestination : set the surface size for scaling
//
// Set the destination size of the associated wl_surface. See
// wp_viewport for the description, and relation to the wl_buffer
// size.
//
// If width is -1 and height is -1, the destination size is unset
// instead. Any other pair of values for width and height that
// contains zero or negative values raises the bad_value protocol
// error.
//
// The crop and scale state is double-buffered state, and will be
// applied on the next wl_surface.commit.
//
//	width: surface width
//	height: surface height
func (i *Viewport) SetDestination(width, height int32) error {
	const opcode = 2
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(width))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(height))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

type ViewportError uint32

// ViewportError :
const (
	// ViewportErrorBadValue : negative or zero values in width or height
	ViewportErrorBadValue ViewportError = 0
	// ViewportErrorBadSize : destination size is not integer
	ViewportErrorBadSize ViewportError = 1
	// ViewportErrorOutOfBuffer : source rectangle extends outside of the content area
	ViewportErrorOutOfBuffer ViewportError = 2
	// ViewportErrorNoSurface : the wl_surface was destroyed
	ViewportErrorNoSurface ViewportError = 3
)

func (e ViewportError) Name() string {
	switch e {
	case ViewportErrorBadValue:
		return "bad_value"
	case ViewportErrorBadSize:
		return "bad_size"
	case ViewportErrorOutOfBuffer:
		return "out_of_buffer"
	case ViewportErrorNoSurface:
		return "no_surface"
	default:
		return ""
	}
}

func (e ViewportError) Value() string {
	switch e {
	case ViewportErrorBadValue:
		return "0"
	case ViewportErrorBadSize:
		return "1"
	case ViewportErrorOutOfBuffer:
		return "2"
	case ViewportErrorNoSurface:
		return "3"
	default:
		return ""
	}
}

func (e ViewportError) String() string {
	return e.Name() + "=" + e.Value()
}
