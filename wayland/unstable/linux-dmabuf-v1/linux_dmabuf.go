// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://raw.githubusercontent.com/wayland-project/wayland-protocols/1.23/unstable/linux-dmabuf/linux-dmabuf-unstable-v1.xml
//
// linux_dmabuf_unstable_v1 Protocol Copyright:
//
// Copyright © 2014, 2015 Collabora, Ltd.
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

package linux_dmabuf

import (
	"reflect"

	"github.com/rajveermalviya/go-wayland/wayland/client"
	"golang.org/x/sys/unix"
)

// LinuxDmabuf : factory for creating dmabuf-based wl_buffers
//
// Following the interfaces from:
// https://www.khronos.org/registry/egl/extensions/EXT/EGL_EXT_image_dma_buf_import.txt
// https://www.khronos.org/registry/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import_modifiers.txt
// and the Linux DRM sub-system's AddFb2 ioctl.
//
// This interface offers ways to create generic dmabuf-based
// wl_buffers. Immediately after a client binds to this interface,
// the set of supported formats and format modifiers is sent with
// 'format' and 'modifier' events.
//
// The following are required from clients:
//
// - Clients must ensure that either all data in the dma-buf is
// coherent for all subsequent read access or that coherency is
// correctly handled by the underlying kernel-side dma-buf
// implementation.
//
// - Don't make any more attachments after sending the buffer to the
// compositor. Making more attachments later increases the risk of
// the compositor not being able to use (re-import) an existing
// dmabuf-based wl_buffer.
//
// The underlying graphics stack must ensure the following:
//
// - The dmabuf file descriptors relayed to the server will stay valid
// for the whole lifetime of the wl_buffer. This means the server may
// at any time use those fds to import the dmabuf into any kernel
// sub-system that might accept it.
//
// However, when the underlying graphics stack fails to deliver the
// promise, because of e.g. a device hot-unplug which raises internal
// errors, after the wl_buffer has been successfully created the
// compositor must not raise protocol errors to the client when dmabuf
// import later fails.
//
// To create a wl_buffer from one or more dmabufs, a client creates a
// zwp_linux_dmabuf_params_v1 object with a zwp_linux_dmabuf_v1.create_params
// request. All planes required by the intended format are added with
// the 'add' request. Finally, a 'create' or 'create_immed' request is
// issued, which has the following outcome depending on the import success.
//
// The 'create' request,
// - on success, triggers a 'created' event which provides the final
// wl_buffer to the client.
// - on failure, triggers a 'failed' event to convey that the server
// cannot use the dmabufs received from the client.
//
// For the 'create_immed' request,
// - on success, the server immediately imports the added dmabufs to
// create a wl_buffer. No event is sent from the server in this case.
// - on failure, the server can choose to either:
// - terminate the client by raising a fatal error.
// - mark the wl_buffer as failed, and send a 'failed' event to the
// client. If the client uses a failed wl_buffer as an argument to any
// request, the behaviour is compositor implementation-defined.
//
// Warning! The protocol described in this file is experimental and
// backward incompatible changes may be made. Backward compatible changes
// may be added together with the corresponding interface version bump.
// Backward incompatible changes are done by bumping the version number in
// the protocol and interface names and resetting the interface version.
// Once the protocol is to be declared stable, the 'z' prefix and the
// version number in the protocol and interface names are removed and the
// interface version number is reset.
type LinuxDmabuf struct {
	client.BaseProxy
	formatHandlers   []LinuxDmabufFormatHandlerFunc
	modifierHandlers []LinuxDmabufModifierHandlerFunc
}

// NewLinuxDmabuf : factory for creating dmabuf-based wl_buffers
//
// Following the interfaces from:
// https://www.khronos.org/registry/egl/extensions/EXT/EGL_EXT_image_dma_buf_import.txt
// https://www.khronos.org/registry/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import_modifiers.txt
// and the Linux DRM sub-system's AddFb2 ioctl.
//
// This interface offers ways to create generic dmabuf-based
// wl_buffers. Immediately after a client binds to this interface,
// the set of supported formats and format modifiers is sent with
// 'format' and 'modifier' events.
//
// The following are required from clients:
//
// - Clients must ensure that either all data in the dma-buf is
// coherent for all subsequent read access or that coherency is
// correctly handled by the underlying kernel-side dma-buf
// implementation.
//
// - Don't make any more attachments after sending the buffer to the
// compositor. Making more attachments later increases the risk of
// the compositor not being able to use (re-import) an existing
// dmabuf-based wl_buffer.
//
// The underlying graphics stack must ensure the following:
//
// - The dmabuf file descriptors relayed to the server will stay valid
// for the whole lifetime of the wl_buffer. This means the server may
// at any time use those fds to import the dmabuf into any kernel
// sub-system that might accept it.
//
// However, when the underlying graphics stack fails to deliver the
// promise, because of e.g. a device hot-unplug which raises internal
// errors, after the wl_buffer has been successfully created the
// compositor must not raise protocol errors to the client when dmabuf
// import later fails.
//
// To create a wl_buffer from one or more dmabufs, a client creates a
// zwp_linux_dmabuf_params_v1 object with a zwp_linux_dmabuf_v1.create_params
// request. All planes required by the intended format are added with
// the 'add' request. Finally, a 'create' or 'create_immed' request is
// issued, which has the following outcome depending on the import success.
//
// The 'create' request,
// - on success, triggers a 'created' event which provides the final
// wl_buffer to the client.
// - on failure, triggers a 'failed' event to convey that the server
// cannot use the dmabufs received from the client.
//
// For the 'create_immed' request,
// - on success, the server immediately imports the added dmabufs to
// create a wl_buffer. No event is sent from the server in this case.
// - on failure, the server can choose to either:
// - terminate the client by raising a fatal error.
// - mark the wl_buffer as failed, and send a 'failed' event to the
// client. If the client uses a failed wl_buffer as an argument to any
// request, the behaviour is compositor implementation-defined.
//
// Warning! The protocol described in this file is experimental and
// backward incompatible changes may be made. Backward compatible changes
// may be added together with the corresponding interface version bump.
// Backward incompatible changes are done by bumping the version number in
// the protocol and interface names and resetting the interface version.
// Once the protocol is to be declared stable, the 'z' prefix and the
// version number in the protocol and interface names are removed and the
// interface version number is reset.
func NewLinuxDmabuf(ctx *client.Context) *LinuxDmabuf {
	zwpLinuxDmabufV1 := &LinuxDmabuf{}
	ctx.Register(zwpLinuxDmabufV1)
	return zwpLinuxDmabufV1
}

// Destroy : unbind the factory
//
// Objects created through this interface, especially wl_buffers, will
// remain valid.
//
func (i *LinuxDmabuf) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const rLen = 8
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return err
}

// CreateParams : create a temporary object for buffer parameters
//
// This temporary object is used to collect multiple dmabuf handles into
// a single batch to create a wl_buffer. It can only be used once and
// should be destroyed after a 'created' or 'failed' event has been
// received.
//
func (i *LinuxDmabuf) CreateParams() (*LinuxBufferParams, error) {
	paramsId := NewLinuxBufferParams(i.Context())
	const opcode = 1
	const rLen = 8 + 4
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], paramsId.ID())
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return paramsId, err
}

// LinuxDmabufFormatEvent : supported buffer format
//
// This event advertises one buffer format that the server supports.
// All the supported formats are advertised once when the client
// binds to this interface. A roundtrip after binding guarantees
// that the client has received all supported formats.
//
// For the definition of the format codes, see the
// zwp_linux_buffer_params_v1::create request.
//
// Warning: the 'format' event is likely to be deprecated and replaced
// with the 'modifier' event introduced in zwp_linux_dmabuf_v1
// version 3, described below. Please refrain from using the information
// received from this event.
type LinuxDmabufFormatEvent struct {
	Format uint32
}
type LinuxDmabufFormatHandlerFunc func(LinuxDmabufFormatEvent)

// AddFormatHandler : adds handler for LinuxDmabufFormatEvent
func (i *LinuxDmabuf) AddFormatHandler(f LinuxDmabufFormatHandlerFunc) {
	if f == nil {
		return
	}

	i.formatHandlers = append(i.formatHandlers, f)
}

func (i *LinuxDmabuf) RemoveFormatHandler(f LinuxDmabufFormatHandlerFunc) {
	for j, e := range i.formatHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.formatHandlers = append(i.formatHandlers[:j], i.formatHandlers[j+1:]...)
			return
		}
	}
}

// LinuxDmabufModifierEvent : supported buffer format modifier
//
// This event advertises the formats that the server supports, along with
// the modifiers supported for each format. All the supported modifiers
// for all the supported formats are advertised once when the client
// binds to this interface. A roundtrip after binding guarantees that
// the client has received all supported format-modifier pairs.
//
// For legacy support, DRM_FORMAT_MOD_INVALID (that is, modifier_hi ==
// 0x00ffffff and modifier_lo == 0xffffffff) is allowed in this event.
// It indicates that the server can support the format with an implicit
// modifier. When a plane has DRM_FORMAT_MOD_INVALID as its modifier, it
// is as if no explicit modifier is specified. The effective modifier
// will be derived from the dmabuf.
//
// A compositor that sends valid modifiers and DRM_FORMAT_MOD_INVALID for
// a given format supports both explicit modifiers and implicit modifiers.
//
// For the definition of the format and modifier codes, see the
// zwp_linux_buffer_params_v1::create and zwp_linux_buffer_params_v1::add
// requests.
type LinuxDmabufModifierEvent struct {
	Format     uint32
	ModifierHi uint32
	ModifierLo uint32
}
type LinuxDmabufModifierHandlerFunc func(LinuxDmabufModifierEvent)

// AddModifierHandler : adds handler for LinuxDmabufModifierEvent
func (i *LinuxDmabuf) AddModifierHandler(f LinuxDmabufModifierHandlerFunc) {
	if f == nil {
		return
	}

	i.modifierHandlers = append(i.modifierHandlers, f)
}

func (i *LinuxDmabuf) RemoveModifierHandler(f LinuxDmabufModifierHandlerFunc) {
	for j, e := range i.modifierHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.modifierHandlers = append(i.modifierHandlers[:j], i.modifierHandlers[j+1:]...)
			return
		}
	}
}

func (i *LinuxDmabuf) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.formatHandlers) == 0 {
			return
		}
		var e LinuxDmabufFormatEvent
		l := 0
		e.Format = client.Uint32(data[l : l+4])
		l += 4
		for _, f := range i.formatHandlers {
			f(e)
		}
	case 1:
		if len(i.modifierHandlers) == 0 {
			return
		}
		var e LinuxDmabufModifierEvent
		l := 0
		e.Format = client.Uint32(data[l : l+4])
		l += 4
		e.ModifierHi = client.Uint32(data[l : l+4])
		l += 4
		e.ModifierLo = client.Uint32(data[l : l+4])
		l += 4
		for _, f := range i.modifierHandlers {
			f(e)
		}
	}
}

// LinuxBufferParams : parameters for creating a dmabuf-based wl_buffer
//
// This temporary object is a collection of dmabufs and other
// parameters that together form a single logical buffer. The temporary
// object may eventually create one wl_buffer unless cancelled by
// destroying it before requesting 'create'.
//
// Single-planar formats only require one dmabuf, however
// multi-planar formats may require more than one dmabuf. For all
// formats, an 'add' request must be called once per plane (even if the
// underlying dmabuf fd is identical).
//
// You must use consecutive plane indices ('plane_idx' argument for 'add')
// from zero to the number of planes used by the drm_fourcc format code.
// All planes required by the format must be given exactly once, but can
// be given in any order. Each plane index can be set only once.
type LinuxBufferParams struct {
	client.BaseProxy
	createdHandlers []LinuxBufferParamsCreatedHandlerFunc
	failedHandlers  []LinuxBufferParamsFailedHandlerFunc
}

// NewLinuxBufferParams : parameters for creating a dmabuf-based wl_buffer
//
// This temporary object is a collection of dmabufs and other
// parameters that together form a single logical buffer. The temporary
// object may eventually create one wl_buffer unless cancelled by
// destroying it before requesting 'create'.
//
// Single-planar formats only require one dmabuf, however
// multi-planar formats may require more than one dmabuf. For all
// formats, an 'add' request must be called once per plane (even if the
// underlying dmabuf fd is identical).
//
// You must use consecutive plane indices ('plane_idx' argument for 'add')
// from zero to the number of planes used by the drm_fourcc format code.
// All planes required by the format must be given exactly once, but can
// be given in any order. Each plane index can be set only once.
func NewLinuxBufferParams(ctx *client.Context) *LinuxBufferParams {
	zwpLinuxBufferParamsV1 := &LinuxBufferParams{}
	ctx.Register(zwpLinuxBufferParamsV1)
	return zwpLinuxBufferParamsV1
}

// Destroy : delete this object, used or not
//
// Cleans up the temporary data sent to the server for dmabuf-based
// wl_buffer creation.
//
func (i *LinuxBufferParams) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const rLen = 8
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return err
}

// Add : add a dmabuf to the temporary set
//
// This request adds one dmabuf to the set in this
// zwp_linux_buffer_params_v1.
//
// The 64-bit unsigned value combined from modifier_hi and modifier_lo
// is the dmabuf layout modifier. DRM AddFB2 ioctl calls this the
// fb modifier, which is defined in drm_mode.h of Linux UAPI.
// This is an opaque token. Drivers use this token to express tiling,
// compression, etc. driver-specific modifications to the base format
// defined by the DRM fourcc code.
//
// Warning: It should be an error if the format/modifier pair was not
// advertised with the modifier event. This is not enforced yet because
// some implementations always accept DRM_FORMAT_MOD_INVALID. Also
// version 2 of this protocol does not have the modifier event.
//
// This request raises the PLANE_IDX error if plane_idx is too large.
// The error PLANE_SET is raised if attempting to set a plane that
// was already set.
//
//  fd: dmabuf fd
//  planeIdx: plane index
//  offset: offset in bytes
//  stride: stride in bytes
//  modifierHi: high 32 bits of layout modifier
//  modifierLo: low 32 bits of layout modifier
func (i *LinuxBufferParams) Add(fd uintptr, planeIdx, offset, stride, modifierHi, modifierLo uint32) error {
	const opcode = 1
	const rLen = 8 + 4 + 4 + 4 + 4 + 4
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(planeIdx))
	l += 4
	client.PutUint32(r[l:l+4], uint32(offset))
	l += 4
	client.PutUint32(r[l:l+4], uint32(stride))
	l += 4
	client.PutUint32(r[l:l+4], uint32(modifierHi))
	l += 4
	client.PutUint32(r[l:l+4], uint32(modifierLo))
	l += 4
	oob := unix.UnixRights(int(fd))
	err := i.Context().WriteMsg(r, oob)
	return err
}

// Create : create a wl_buffer from the given dmabufs
//
// This asks for creation of a wl_buffer from the added dmabuf
// buffers. The wl_buffer is not created immediately but returned via
// the 'created' event if the dmabuf sharing succeeds. The sharing
// may fail at runtime for reasons a client cannot predict, in
// which case the 'failed' event is triggered.
//
// The 'format' argument is a DRM_FORMAT code, as defined by the
// libdrm's drm_fourcc.h. The Linux kernel's DRM sub-system is the
// authoritative source on how the format codes should work.
//
// The 'flags' is a bitfield of the flags defined in enum "flags".
// 'y_invert' means the that the image needs to be y-flipped.
//
// Flag 'interlaced' means that the frame in the buffer is not
// progressive as usual, but interlaced. An interlaced buffer as
// supported here must always contain both top and bottom fields.
// The top field always begins on the first pixel row. The temporal
// ordering between the two fields is top field first, unless
// 'bottom_first' is specified. It is undefined whether 'bottom_first'
// is ignored if 'interlaced' is not set.
//
// This protocol does not convey any information about field rate,
// duration, or timing, other than the relative ordering between the
// two fields in one buffer. A compositor may have to estimate the
// intended field rate from the incoming buffer rate. It is undefined
// whether the time of receiving wl_surface.commit with a new buffer
// attached, applying the wl_surface state, wl_surface.frame callback
// trigger, presentation, or any other point in the compositor cycle
// is used to measure the frame or field times. There is no support
// for detecting missed or late frames/fields/buffers either, and
// there is no support whatsoever for cooperating with interlaced
// compositor output.
//
// The composited image quality resulting from the use of interlaced
// buffers is explicitly undefined. A compositor may use elaborate
// hardware features or software to deinterlace and create progressive
// output frames from a sequence of interlaced input buffers, or it
// may produce substandard image quality. However, compositors that
// cannot guarantee reasonable image quality in all cases are recommended
// to just reject all interlaced buffers.
//
// Any argument errors, including non-positive width or height,
// mismatch between the number of planes and the format, bad
// format, bad offset or stride, may be indicated by fatal protocol
// errors: INCOMPLETE, INVALID_FORMAT, INVALID_DIMENSIONS,
// OUT_OF_BOUNDS.
//
// Dmabuf import errors in the server that are not obvious client
// bugs are returned via the 'failed' event as non-fatal. This
// allows attempting dmabuf sharing and falling back in the client
// if it fails.
//
// This request can be sent only once in the object's lifetime, after
// which the only legal request is destroy. This object should be
// destroyed after issuing a 'create' request. Attempting to use this
// object after issuing 'create' raises ALREADY_USED protocol error.
//
// It is not mandatory to issue 'create'. If a client wants to
// cancel the buffer creation, it can just destroy this object.
//
//  width: base plane width in pixels
//  height: base plane height in pixels
//  format: DRM_FORMAT code
//  flags: see enum flags
func (i *LinuxBufferParams) Create(width, height int32, format, flags uint32) error {
	const opcode = 2
	const rLen = 8 + 4 + 4 + 4 + 4
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(width))
	l += 4
	client.PutUint32(r[l:l+4], uint32(height))
	l += 4
	client.PutUint32(r[l:l+4], uint32(format))
	l += 4
	client.PutUint32(r[l:l+4], uint32(flags))
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return err
}

// CreateImmed : immediately create a wl_buffer from the given dmabufs
//
// This asks for immediate creation of a wl_buffer by importing the
// added dmabufs.
//
// In case of import success, no event is sent from the server, and the
// wl_buffer is ready to be used by the client.
//
// Upon import failure, either of the following may happen, as seen fit
// by the implementation:
// - the client is terminated with one of the following fatal protocol
// errors:
// - INCOMPLETE, INVALID_FORMAT, INVALID_DIMENSIONS, OUT_OF_BOUNDS,
// in case of argument errors such as mismatch between the number
// of planes and the format, bad format, non-positive width or
// height, or bad offset or stride.
// - INVALID_WL_BUFFER, in case the cause for failure is unknown or
// plaform specific.
// - the server creates an invalid wl_buffer, marks it as failed and
// sends a 'failed' event to the client. The result of using this
// invalid wl_buffer as an argument in any request by the client is
// defined by the compositor implementation.
//
// This takes the same arguments as a 'create' request, and obeys the
// same restrictions.
//
//  width: base plane width in pixels
//  height: base plane height in pixels
//  format: DRM_FORMAT code
//  flags: see enum flags
func (i *LinuxBufferParams) CreateImmed(width, height int32, format, flags uint32) (*client.Buffer, error) {
	bufferId := client.NewBuffer(i.Context())
	const opcode = 3
	const rLen = 8 + 4 + 4 + 4 + 4 + 4
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], bufferId.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(width))
	l += 4
	client.PutUint32(r[l:l+4], uint32(height))
	l += 4
	client.PutUint32(r[l:l+4], uint32(format))
	l += 4
	client.PutUint32(r[l:l+4], uint32(flags))
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return bufferId, err
}

type LinuxBufferParamsError uint32

// LinuxBufferParamsError :
const (
	// LinuxBufferParamsErrorAlreadyUsed : the dmabuf_batch object has already been used to create a wl_buffer
	LinuxBufferParamsErrorAlreadyUsed LinuxBufferParamsError = 0
	// LinuxBufferParamsErrorPlaneIdx : plane index out of bounds
	LinuxBufferParamsErrorPlaneIdx LinuxBufferParamsError = 1
	// LinuxBufferParamsErrorPlaneSet : the plane index was already set
	LinuxBufferParamsErrorPlaneSet LinuxBufferParamsError = 2
	// LinuxBufferParamsErrorIncomplete : missing or too many planes to create a buffer
	LinuxBufferParamsErrorIncomplete LinuxBufferParamsError = 3
	// LinuxBufferParamsErrorInvalidFormat : format not supported
	LinuxBufferParamsErrorInvalidFormat LinuxBufferParamsError = 4
	// LinuxBufferParamsErrorInvalidDimensions : invalid width or height
	LinuxBufferParamsErrorInvalidDimensions LinuxBufferParamsError = 5
	// LinuxBufferParamsErrorOutOfBounds : offset + stride * height goes out of dmabuf bounds
	LinuxBufferParamsErrorOutOfBounds LinuxBufferParamsError = 6
	// LinuxBufferParamsErrorInvalidWlBuffer : invalid wl_buffer resulted from importing dmabufs via the create_immed request on given buffer_params
	LinuxBufferParamsErrorInvalidWlBuffer LinuxBufferParamsError = 7
)

func (e LinuxBufferParamsError) Name() string {
	switch e {
	case LinuxBufferParamsErrorAlreadyUsed:
		return "already_used"
	case LinuxBufferParamsErrorPlaneIdx:
		return "plane_idx"
	case LinuxBufferParamsErrorPlaneSet:
		return "plane_set"
	case LinuxBufferParamsErrorIncomplete:
		return "incomplete"
	case LinuxBufferParamsErrorInvalidFormat:
		return "invalid_format"
	case LinuxBufferParamsErrorInvalidDimensions:
		return "invalid_dimensions"
	case LinuxBufferParamsErrorOutOfBounds:
		return "out_of_bounds"
	case LinuxBufferParamsErrorInvalidWlBuffer:
		return "invalid_wl_buffer"
	default:
		return ""
	}
}

func (e LinuxBufferParamsError) Value() string {
	switch e {
	case LinuxBufferParamsErrorAlreadyUsed:
		return "0"
	case LinuxBufferParamsErrorPlaneIdx:
		return "1"
	case LinuxBufferParamsErrorPlaneSet:
		return "2"
	case LinuxBufferParamsErrorIncomplete:
		return "3"
	case LinuxBufferParamsErrorInvalidFormat:
		return "4"
	case LinuxBufferParamsErrorInvalidDimensions:
		return "5"
	case LinuxBufferParamsErrorOutOfBounds:
		return "6"
	case LinuxBufferParamsErrorInvalidWlBuffer:
		return "7"
	default:
		return ""
	}
}

func (e LinuxBufferParamsError) String() string {
	return e.Name() + "=" + e.Value()
}

type LinuxBufferParamsFlags uint32

// LinuxBufferParamsFlags :
const (
	// LinuxBufferParamsFlagsYInvert : contents are y-inverted
	LinuxBufferParamsFlagsYInvert LinuxBufferParamsFlags = 1
	// LinuxBufferParamsFlagsInterlaced : content is interlaced
	LinuxBufferParamsFlagsInterlaced LinuxBufferParamsFlags = 2
	// LinuxBufferParamsFlagsBottomFirst : bottom field first
	LinuxBufferParamsFlagsBottomFirst LinuxBufferParamsFlags = 4
)

func (e LinuxBufferParamsFlags) Name() string {
	switch e {
	case LinuxBufferParamsFlagsYInvert:
		return "y_invert"
	case LinuxBufferParamsFlagsInterlaced:
		return "interlaced"
	case LinuxBufferParamsFlagsBottomFirst:
		return "bottom_first"
	default:
		return ""
	}
}

func (e LinuxBufferParamsFlags) Value() string {
	switch e {
	case LinuxBufferParamsFlagsYInvert:
		return "1"
	case LinuxBufferParamsFlagsInterlaced:
		return "2"
	case LinuxBufferParamsFlagsBottomFirst:
		return "4"
	default:
		return ""
	}
}

func (e LinuxBufferParamsFlags) String() string {
	return e.Name() + "=" + e.Value()
}

// LinuxBufferParamsCreatedEvent : buffer creation succeeded
//
// This event indicates that the attempted buffer creation was
// successful. It provides the new wl_buffer referencing the dmabuf(s).
//
// Upon receiving this event, the client should destroy the
// zlinux_dmabuf_params object.
type LinuxBufferParamsCreatedEvent struct {
	Buffer *client.Buffer
}
type LinuxBufferParamsCreatedHandlerFunc func(LinuxBufferParamsCreatedEvent)

// AddCreatedHandler : adds handler for LinuxBufferParamsCreatedEvent
func (i *LinuxBufferParams) AddCreatedHandler(f LinuxBufferParamsCreatedHandlerFunc) {
	if f == nil {
		return
	}

	i.createdHandlers = append(i.createdHandlers, f)
}

func (i *LinuxBufferParams) RemoveCreatedHandler(f LinuxBufferParamsCreatedHandlerFunc) {
	for j, e := range i.createdHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.createdHandlers = append(i.createdHandlers[:j], i.createdHandlers[j+1:]...)
			return
		}
	}
}

// LinuxBufferParamsFailedEvent : buffer creation failed
//
// This event indicates that the attempted buffer creation has
// failed. It usually means that one of the dmabuf constraints
// has not been fulfilled.
//
// Upon receiving this event, the client should destroy the
// zlinux_buffer_params object.
type LinuxBufferParamsFailedEvent struct{}
type LinuxBufferParamsFailedHandlerFunc func(LinuxBufferParamsFailedEvent)

// AddFailedHandler : adds handler for LinuxBufferParamsFailedEvent
func (i *LinuxBufferParams) AddFailedHandler(f LinuxBufferParamsFailedHandlerFunc) {
	if f == nil {
		return
	}

	i.failedHandlers = append(i.failedHandlers, f)
}

func (i *LinuxBufferParams) RemoveFailedHandler(f LinuxBufferParamsFailedHandlerFunc) {
	for j, e := range i.failedHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.failedHandlers = append(i.failedHandlers[:j], i.failedHandlers[j+1:]...)
			return
		}
	}
}

func (i *LinuxBufferParams) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.createdHandlers) == 0 {
			return
		}
		var e LinuxBufferParamsCreatedEvent
		l := 0
		e.Buffer = i.Context().GetProxy(client.Uint32(data[l : l+4])).(*client.Buffer)
		l += 4
		for _, f := range i.createdHandlers {
			f(e)
		}
	case 1:
		if len(i.failedHandlers) == 0 {
			return
		}
		var e LinuxBufferParamsFailedEvent
		for _, f := range i.failedHandlers {
			f(e)
		}
	}
}
