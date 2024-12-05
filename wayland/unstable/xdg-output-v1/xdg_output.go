// Generated by go-wayland-scanner
// https://github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.31/unstable/xdg-output/xdg-output-unstable-v1.xml?ref_type=tags
//
// xdg_output_unstable_v1 Protocol Copyright:
//
// Copyright © 2017 Red Hat Inc.
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

package xdg_output

import "github.com/MatthiasKunnen/go-wayland/wayland/client"

// OutputManagerInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const OutputManagerInterfaceName = "zxdg_output_manager_v1"

// OutputManager : manage xdg_output objects
//
// A global factory interface for xdg_output objects.
type OutputManager struct {
	client.BaseProxy
}

// NewOutputManager : manage xdg_output objects
//
// A global factory interface for xdg_output objects.
func NewOutputManager(ctx *client.Context) *OutputManager {
	zxdgOutputManagerV1 := &OutputManager{}
	ctx.Register(zxdgOutputManagerV1)
	return zxdgOutputManagerV1
}

// Destroy : destroy the xdg_output_manager object
//
// Using this request a client can tell the server that it is not
// going to use the xdg_output_manager object anymore.
//
// Any objects already created through this instance are not affected.
func (i *OutputManager) Destroy() error {
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

// GetXdgOutput : create an xdg output from a wl_output
//
// This creates a new xdg_output object for the given wl_output.
func (i *OutputManager) GetXdgOutput(output *client.Output) (*Output, error) {
	id := NewOutput(i.Context())
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
	client.PutUint32(_reqBuf[l:l+4], output.ID())
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

// OutputInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const OutputInterfaceName = "zxdg_output_v1"

// Output : compositor logical output region
//
// An xdg_output describes part of the compositor geometry.
//
// This typically corresponds to a monitor that displays part of the
// compositor space.
//
// For objects version 3 onwards, after all xdg_output properties have been
// sent (when the object is created and when properties are updated), a
// wl_output.done event is sent. This allows changes to the output
// properties to be seen as atomic, even if they happen via multiple events.
type Output struct {
	client.BaseProxy
	logicalPositionHandler OutputLogicalPositionHandlerFunc
	logicalSizeHandler     OutputLogicalSizeHandlerFunc
	doneHandler            OutputDoneHandlerFunc
	nameHandler            OutputNameHandlerFunc
	descriptionHandler     OutputDescriptionHandlerFunc
}

// NewOutput : compositor logical output region
//
// An xdg_output describes part of the compositor geometry.
//
// This typically corresponds to a monitor that displays part of the
// compositor space.
//
// For objects version 3 onwards, after all xdg_output properties have been
// sent (when the object is created and when properties are updated), a
// wl_output.done event is sent. This allows changes to the output
// properties to be seen as atomic, even if they happen via multiple events.
func NewOutput(ctx *client.Context) *Output {
	zxdgOutputV1 := &Output{}
	ctx.Register(zxdgOutputV1)
	return zxdgOutputV1
}

// Destroy : destroy the xdg_output object
//
// Using this request a client can tell the server that it is not
// going to use the xdg_output object anymore.
func (i *Output) Destroy() error {
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

// OutputLogicalPositionEvent : position of the output within the global compositor space
//
// The position event describes the location of the wl_output within
// the global compositor space.
//
// The logical_position event is sent after creating an xdg_output
// (see xdg_output_manager.get_xdg_output) and whenever the location
// of the output changes within the global compositor space.
type OutputLogicalPositionEvent struct {
	X int32
	Y int32
}
type OutputLogicalPositionHandlerFunc func(OutputLogicalPositionEvent)

// SetLogicalPositionHandler : sets handler for OutputLogicalPositionEvent
func (i *Output) SetLogicalPositionHandler(f OutputLogicalPositionHandlerFunc) {
	i.logicalPositionHandler = f
}

// OutputLogicalSizeEvent : size of the output in the global compositor space
//
// The logical_size event describes the size of the output in the
// global compositor space.
//
// For example, a surface without any buffer scale, transformation
// nor rotation set, with the size matching the logical_size will
// have the same size as the corresponding output when displayed.
//
// Most regular Wayland clients should not pay attention to the
// logical size and would rather rely on xdg_shell interfaces.
//
// Some clients such as Xwayland, however, need this to configure
// their surfaces in the global compositor space as the compositor
// may apply a different scale from what is advertised by the output
// scaling property (to achieve fractional scaling, for example).
//
// For example, for a wl_output mode 3840×2160 and a scale factor 2:
//
// - A compositor not scaling the surface buffers will advertise a
// logical size of 3840×2160,
//
// - A compositor automatically scaling the surface buffers will
// advertise a logical size of 1920×1080,
//
// - A compositor using a fractional scale of 1.5 will advertise a
// logical size of 2560×1440.
//
// For example, for a wl_output mode 1920×1080 and a 90 degree rotation,
// the compositor will advertise a logical size of 1080x1920.
//
// The logical_size event is sent after creating an xdg_output
// (see xdg_output_manager.get_xdg_output) and whenever the logical
// size of the output changes, either as a result of a change in the
// applied scale or because of a change in the corresponding output
// mode(see wl_output.mode) or transform (see wl_output.transform).
type OutputLogicalSizeEvent struct {
	Width  int32
	Height int32
}
type OutputLogicalSizeHandlerFunc func(OutputLogicalSizeEvent)

// SetLogicalSizeHandler : sets handler for OutputLogicalSizeEvent
func (i *Output) SetLogicalSizeHandler(f OutputLogicalSizeHandlerFunc) {
	i.logicalSizeHandler = f
}

// OutputDoneEvent : all information about the output have been sent
//
// This event is sent after all other properties of an xdg_output
// have been sent.
//
// This allows changes to the xdg_output properties to be seen as
// atomic, even if they happen via multiple events.
//
// For objects version 3 onwards, this event is deprecated. Compositors
// are not required to send it anymore and must send wl_output.done
// instead.
type OutputDoneEvent struct{}
type OutputDoneHandlerFunc func(OutputDoneEvent)

// SetDoneHandler : sets handler for OutputDoneEvent
func (i *Output) SetDoneHandler(f OutputDoneHandlerFunc) {
	i.doneHandler = f
}

// OutputNameEvent : name of this output
//
// Many compositors will assign names to their outputs, show them to the
// user, allow them to be configured by name, etc. The client may wish to
// know this name as well to offer the user similar behaviors.
//
// The naming convention is compositor defined, but limited to
// alphanumeric characters and dashes (-). Each name is unique among all
// wl_output globals, but if a wl_output global is destroyed the same name
// may be reused later. The names will also remain consistent across
// sessions with the same hardware and software configuration.
//
// Examples of names include 'HDMI-A-1', 'WL-1', 'X11-1', etc. However, do
// not assume that the name is a reflection of an underlying DRM
// connector, X11 connection, etc.
//
// The name event is sent after creating an xdg_output (see
// xdg_output_manager.get_xdg_output). This event is only sent once per
// xdg_output, and the name does not change over the lifetime of the
// wl_output global.
type OutputNameEvent struct {
	Name string
}
type OutputNameHandlerFunc func(OutputNameEvent)

// SetNameHandler : sets handler for OutputNameEvent
func (i *Output) SetNameHandler(f OutputNameHandlerFunc) {
	i.nameHandler = f
}

// OutputDescriptionEvent : human-readable description of this output
//
// Many compositors can produce human-readable descriptions of their
// outputs.  The client may wish to know this description as well, to
// communicate the user for various purposes.
//
// The description is a UTF-8 string with no convention defined for its
// contents. Examples might include 'Foocorp 11" Display' or 'Virtual X11
// output via :1'.
//
// The description event is sent after creating an xdg_output (see
// xdg_output_manager.get_xdg_output) and whenever the description
// changes. The description is optional, and may not be sent at all.
//
// For objects of version 2 and lower, this event is only sent once per
// xdg_output, and the description does not change over the lifetime of
// the wl_output global.
type OutputDescriptionEvent struct {
	Description string
}
type OutputDescriptionHandlerFunc func(OutputDescriptionEvent)

// SetDescriptionHandler : sets handler for OutputDescriptionEvent
func (i *Output) SetDescriptionHandler(f OutputDescriptionHandlerFunc) {
	i.descriptionHandler = f
}

func (i *Output) Dispatch(opcode uint32, fd int, data []byte) {
	switch opcode {
	case 0:
		if i.logicalPositionHandler == nil {
			return
		}
		var e OutputLogicalPositionEvent
		l := 0
		e.X = int32(client.Uint32(data[l : l+4]))
		l += 4
		e.Y = int32(client.Uint32(data[l : l+4]))
		l += 4

		i.logicalPositionHandler(e)
	case 1:
		if i.logicalSizeHandler == nil {
			return
		}
		var e OutputLogicalSizeEvent
		l := 0
		e.Width = int32(client.Uint32(data[l : l+4]))
		l += 4
		e.Height = int32(client.Uint32(data[l : l+4]))
		l += 4

		i.logicalSizeHandler(e)
	case 2:
		if i.doneHandler == nil {
			return
		}
		var e OutputDoneEvent

		i.doneHandler(e)
	case 3:
		if i.nameHandler == nil {
			return
		}
		var e OutputNameEvent
		l := 0
		nameLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Name = client.String(data[l : l+nameLen])
		l += nameLen

		i.nameHandler(e)
	case 4:
		if i.descriptionHandler == nil {
			return
		}
		var e OutputDescriptionEvent
		l := 0
		descriptionLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Description = client.String(data[l : l+descriptionLen])
		l += descriptionLen

		i.descriptionHandler(e)
	}
}
