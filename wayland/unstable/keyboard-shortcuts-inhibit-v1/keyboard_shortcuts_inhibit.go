// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.31/unstable/keyboard-shortcuts-inhibit/keyboard-shortcuts-inhibit-unstable-v1.xml?ref_type=tags
//
// keyboard_shortcuts_inhibit_unstable_v1 Protocol Copyright:
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

package keyboard_shortcuts_inhibit

import "github.com/rajveermalviya/go-wayland/wayland/client"

// KeyboardShortcutsInhibitManager : context object for keyboard grab_manager
//
// A global interface used for inhibiting the compositor keyboard shortcuts.
type KeyboardShortcutsInhibitManager struct {
	client.BaseProxy
}

// NewKeyboardShortcutsInhibitManager : context object for keyboard grab_manager
//
// A global interface used for inhibiting the compositor keyboard shortcuts.
func NewKeyboardShortcutsInhibitManager(ctx *client.Context) *KeyboardShortcutsInhibitManager {
	zwpKeyboardShortcutsInhibitManagerV1 := &KeyboardShortcutsInhibitManager{}
	ctx.Register(zwpKeyboardShortcutsInhibitManagerV1)
	return zwpKeyboardShortcutsInhibitManagerV1
}

// Destroy : destroy the keyboard shortcuts inhibitor object
//
// Destroy the keyboard shortcuts inhibitor manager.
func (i *KeyboardShortcutsInhibitManager) Destroy() error {
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

// InhibitShortcuts : create a new keyboard shortcuts inhibitor object
//
// Create a new keyboard shortcuts inhibitor object associated with
// the given surface for the given seat.
//
// If shortcuts are already inhibited for the specified seat and surface,
// a protocol error "already_inhibited" is raised by the compositor.
//
//	surface: the surface that inhibits the keyboard shortcuts behavior
//	seat: the wl_seat for which keyboard shortcuts should be disabled
func (i *KeyboardShortcutsInhibitManager) InhibitShortcuts(surface *client.Surface, seat *client.Seat) (*KeyboardShortcutsInhibitor, error) {
	id := NewKeyboardShortcutsInhibitor(i.Context())
	const opcode = 1
	const _reqBufLen = 8 + 4 + 4 + 4
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
	client.PutUint32(_reqBuf[l:l+4], seat.ID())
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

type KeyboardShortcutsInhibitManagerError uint32

// KeyboardShortcutsInhibitManagerError :
const (
	// KeyboardShortcutsInhibitManagerErrorAlreadyInhibited : the shortcuts are already inhibited for this surface
	KeyboardShortcutsInhibitManagerErrorAlreadyInhibited KeyboardShortcutsInhibitManagerError = 0
)

func (e KeyboardShortcutsInhibitManagerError) Name() string {
	switch e {
	case KeyboardShortcutsInhibitManagerErrorAlreadyInhibited:
		return "already_inhibited"
	default:
		return ""
	}
}

func (e KeyboardShortcutsInhibitManagerError) Value() string {
	switch e {
	case KeyboardShortcutsInhibitManagerErrorAlreadyInhibited:
		return "0"
	default:
		return ""
	}
}

func (e KeyboardShortcutsInhibitManagerError) String() string {
	return e.Name() + "=" + e.Value()
}

// KeyboardShortcutsInhibitor : context object for keyboard shortcuts inhibitor
//
// A keyboard shortcuts inhibitor instructs the compositor to ignore
// its own keyboard shortcuts when the associated surface has keyboard
// focus. As a result, when the surface has keyboard focus on the given
// seat, it will receive all key events originating from the specified
// seat, even those which would normally be caught by the compositor for
// its own shortcuts.
//
// The Wayland compositor is however under no obligation to disable
// all of its shortcuts, and may keep some special key combo for its own
// use, including but not limited to one allowing the user to forcibly
// restore normal keyboard events routing in the case of an unwilling
// client. The compositor may also use the same key combo to reactivate
// an existing shortcut inhibitor that was previously deactivated on
// user request.
//
// When the compositor restores its own keyboard shortcuts, an
// "inactive" event is emitted to notify the client that the keyboard
// shortcuts inhibitor is not effectively active for the surface and
// seat any more, and the client should not expect to receive all
// keyboard events.
//
// When the keyboard shortcuts inhibitor is inactive, the client has
// no way to forcibly reactivate the keyboard shortcuts inhibitor.
//
// The user can chose to re-enable a previously deactivated keyboard
// shortcuts inhibitor using any mechanism the compositor may offer,
// in which case the compositor will send an "active" event to notify
// the client.
//
// If the surface is destroyed, unmapped, or loses the seat's keyboard
// focus, the keyboard shortcuts inhibitor becomes irrelevant and the
// compositor will restore its own keyboard shortcuts but no "inactive"
// event is emitted in this case.
type KeyboardShortcutsInhibitor struct {
	client.BaseProxy
	activeHandler   KeyboardShortcutsInhibitorActiveHandlerFunc
	inactiveHandler KeyboardShortcutsInhibitorInactiveHandlerFunc
}

// NewKeyboardShortcutsInhibitor : context object for keyboard shortcuts inhibitor
//
// A keyboard shortcuts inhibitor instructs the compositor to ignore
// its own keyboard shortcuts when the associated surface has keyboard
// focus. As a result, when the surface has keyboard focus on the given
// seat, it will receive all key events originating from the specified
// seat, even those which would normally be caught by the compositor for
// its own shortcuts.
//
// The Wayland compositor is however under no obligation to disable
// all of its shortcuts, and may keep some special key combo for its own
// use, including but not limited to one allowing the user to forcibly
// restore normal keyboard events routing in the case of an unwilling
// client. The compositor may also use the same key combo to reactivate
// an existing shortcut inhibitor that was previously deactivated on
// user request.
//
// When the compositor restores its own keyboard shortcuts, an
// "inactive" event is emitted to notify the client that the keyboard
// shortcuts inhibitor is not effectively active for the surface and
// seat any more, and the client should not expect to receive all
// keyboard events.
//
// When the keyboard shortcuts inhibitor is inactive, the client has
// no way to forcibly reactivate the keyboard shortcuts inhibitor.
//
// The user can chose to re-enable a previously deactivated keyboard
// shortcuts inhibitor using any mechanism the compositor may offer,
// in which case the compositor will send an "active" event to notify
// the client.
//
// If the surface is destroyed, unmapped, or loses the seat's keyboard
// focus, the keyboard shortcuts inhibitor becomes irrelevant and the
// compositor will restore its own keyboard shortcuts but no "inactive"
// event is emitted in this case.
func NewKeyboardShortcutsInhibitor(ctx *client.Context) *KeyboardShortcutsInhibitor {
	zwpKeyboardShortcutsInhibitorV1 := &KeyboardShortcutsInhibitor{}
	ctx.Register(zwpKeyboardShortcutsInhibitorV1)
	return zwpKeyboardShortcutsInhibitorV1
}

// Destroy : destroy the keyboard shortcuts inhibitor object
//
// Remove the keyboard shortcuts inhibitor from the associated wl_surface.
func (i *KeyboardShortcutsInhibitor) Destroy() error {
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

// KeyboardShortcutsInhibitorActiveEvent : shortcuts are inhibited
//
// This event indicates that the shortcut inhibitor is active.
//
// The compositor sends this event every time compositor shortcuts
// are inhibited on behalf of the surface. When active, the client
// may receive input events normally reserved by the compositor
// (see zwp_keyboard_shortcuts_inhibitor_v1).
//
// This occurs typically when the initial request "inhibit_shortcuts"
// first becomes active or when the user instructs the compositor to
// re-enable and existing shortcuts inhibitor using any mechanism
// offered by the compositor.
type KeyboardShortcutsInhibitorActiveEvent struct{}
type KeyboardShortcutsInhibitorActiveHandlerFunc func(KeyboardShortcutsInhibitorActiveEvent)

// SetActiveHandler : sets handler for KeyboardShortcutsInhibitorActiveEvent
func (i *KeyboardShortcutsInhibitor) SetActiveHandler(f KeyboardShortcutsInhibitorActiveHandlerFunc) {
	i.activeHandler = f
}

// KeyboardShortcutsInhibitorInactiveEvent : shortcuts are restored
//
// This event indicates that the shortcuts inhibitor is inactive,
// normal shortcuts processing is restored by the compositor.
type KeyboardShortcutsInhibitorInactiveEvent struct{}
type KeyboardShortcutsInhibitorInactiveHandlerFunc func(KeyboardShortcutsInhibitorInactiveEvent)

// SetInactiveHandler : sets handler for KeyboardShortcutsInhibitorInactiveEvent
func (i *KeyboardShortcutsInhibitor) SetInactiveHandler(f KeyboardShortcutsInhibitorInactiveHandlerFunc) {
	i.inactiveHandler = f
}

func (i *KeyboardShortcutsInhibitor) Dispatch(opcode uint32, fd int, data []byte) {
	switch opcode {
	case 0:
		if i.activeHandler == nil {
			return
		}
		var e KeyboardShortcutsInhibitorActiveEvent

		i.activeHandler(e)
	case 1:
		if i.inactiveHandler == nil {
			return
		}
		var e KeyboardShortcutsInhibitorInactiveEvent

		i.inactiveHandler(e)
	}
}
