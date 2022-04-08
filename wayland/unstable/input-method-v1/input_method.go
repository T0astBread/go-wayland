// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://raw.githubusercontent.com/wayland-project/wayland-protocols/1.25/unstable/input-method/input-method-unstable-v1.xml
//
// input_method_unstable_v1 Protocol Copyright:
//
// Copyright © 2012, 2013 Intel Corporation
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

package input_method

import "github.com/rajveermalviya/go-wayland/wayland/client"

// InputMethodContext : input method context
//
// Corresponds to a text input on the input method side. An input method context
// is created on text input activation on the input method side. It allows
// receiving information about the text input from the application via events.
// Input method contexts do not keep state after deactivation and should be
// destroyed after deactivation is handled.
//
// Text is generally UTF-8 encoded, indices and lengths are in bytes.
//
// Serials are used to synchronize the state between the text input and
// an input method. New serials are sent by the text input in the
// commit_state request and are used by the input method to indicate
// the known text input state in events like preedit_string, commit_string,
// and keysym. The text input can then ignore events from the input method
// which are based on an outdated state (for example after a reset).
//
// Warning! The protocol described in this file is experimental and
// backward incompatible changes may be made. Backward compatible changes
// may be added together with the corresponding interface version bump.
// Backward incompatible changes are done by bumping the version number in
// the protocol and interface names and resetting the interface version.
// Once the protocol is to be declared stable, the 'z' prefix and the
// version number in the protocol and interface names are removed and the
// interface version number is reset.
type InputMethodContext struct {
	client.BaseProxy
	surroundingTextHandlers   []InputMethodContextSurroundingTextHandlerFunc
	resetHandlers             []InputMethodContextResetHandlerFunc
	contentTypeHandlers       []InputMethodContextContentTypeHandlerFunc
	invokeActionHandlers      []InputMethodContextInvokeActionHandlerFunc
	commitStateHandlers       []InputMethodContextCommitStateHandlerFunc
	preferredLanguageHandlers []InputMethodContextPreferredLanguageHandlerFunc
}

// NewInputMethodContext : input method context
//
// Corresponds to a text input on the input method side. An input method context
// is created on text input activation on the input method side. It allows
// receiving information about the text input from the application via events.
// Input method contexts do not keep state after deactivation and should be
// destroyed after deactivation is handled.
//
// Text is generally UTF-8 encoded, indices and lengths are in bytes.
//
// Serials are used to synchronize the state between the text input and
// an input method. New serials are sent by the text input in the
// commit_state request and are used by the input method to indicate
// the known text input state in events like preedit_string, commit_string,
// and keysym. The text input can then ignore events from the input method
// which are based on an outdated state (for example after a reset).
//
// Warning! The protocol described in this file is experimental and
// backward incompatible changes may be made. Backward compatible changes
// may be added together with the corresponding interface version bump.
// Backward incompatible changes are done by bumping the version number in
// the protocol and interface names and resetting the interface version.
// Once the protocol is to be declared stable, the 'z' prefix and the
// version number in the protocol and interface names are removed and the
// interface version number is reset.
func NewInputMethodContext(ctx *client.Context) *InputMethodContext {
	zwpInputMethodContextV1 := &InputMethodContext{}
	ctx.Register(zwpInputMethodContextV1)
	return zwpInputMethodContextV1
}

// Destroy :
//
func (i *InputMethodContext) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const rLen = 8
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// CommitString : commit string
//
// Send the commit string text for insertion to the application.
//
// The text to commit could be either just a single character after a key
// press or the result of some composing (pre-edit). It could be also an
// empty text when some text should be removed (see
// delete_surrounding_text) or when the input cursor should be moved (see
// cursor_position).
//
// Any previously set composing text will be removed.
//
//  serial: serial of the latest known text input state
func (i *InputMethodContext) CommitString(serial uint32, text string) error {
	const opcode = 1
	textLen := client.PaddedLen(len(text) + 1)
	rLen := 8 + 4 + (4 + textLen)
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(serial))
	l += 4
	client.PutString(r[l:l+(4+textLen)], text, textLen)
	l += (4 + textLen)
	err := i.Context().WriteMsg(r, nil)
	return err
}

// PreeditString : pre-edit string
//
// Send the pre-edit string text to the application text input.
//
// The commit text can be used to replace the pre-edit text on reset (for
// example on unfocus).
//
// Previously sent preedit_style and preedit_cursor requests are also
// processed by the text_input.
//
//  serial: serial of the latest known text input state
func (i *InputMethodContext) PreeditString(serial uint32, text, commit string) error {
	const opcode = 2
	textLen := client.PaddedLen(len(text) + 1)
	commitLen := client.PaddedLen(len(commit) + 1)
	rLen := 8 + 4 + (4 + textLen) + (4 + commitLen)
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(serial))
	l += 4
	client.PutString(r[l:l+(4+textLen)], text, textLen)
	l += (4 + textLen)
	client.PutString(r[l:l+(4+commitLen)], commit, commitLen)
	l += (4 + commitLen)
	err := i.Context().WriteMsg(r, nil)
	return err
}

// PreeditStyling : pre-edit styling
//
// Set the styling information on composing text. The style is applied for
// length in bytes from index relative to the beginning of
// the composing text (as byte offset). Multiple styles can
// be applied to a composing text.
//
// This request should be sent before sending a preedit_string request.
//
func (i *InputMethodContext) PreeditStyling(index, length, style uint32) error {
	const opcode = 3
	const rLen = 8 + 4 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(index))
	l += 4
	client.PutUint32(r[l:l+4], uint32(length))
	l += 4
	client.PutUint32(r[l:l+4], uint32(style))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// PreeditCursor : pre-edit cursor
//
// Set the cursor position inside the composing text (as byte offset)
// relative to the start of the composing text.
//
// When index is negative no cursor should be displayed.
//
// This request should be sent before sending a preedit_string request.
//
func (i *InputMethodContext) PreeditCursor(index int32) error {
	const opcode = 4
	const rLen = 8 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(index))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// DeleteSurroundingText : delete text
//
// Remove the surrounding text.
//
// This request will be handled on the text_input side directly following
// a commit_string request.
//
func (i *InputMethodContext) DeleteSurroundingText(index int32, length uint32) error {
	const opcode = 5
	const rLen = 8 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(index))
	l += 4
	client.PutUint32(r[l:l+4], uint32(length))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// CursorPosition : set cursor to a new position
//
// Set the cursor and anchor to a new position. Index is the new cursor
// position in bytes (when >= 0 this is relative to the end of the inserted text,
// otherwise it is relative to the beginning of the inserted text). Anchor is
// the new anchor position in bytes (when >= 0 this is relative to the end of the
// inserted text, otherwise it is relative to the beginning of the inserted
// text). When there should be no selected text, anchor should be the same
// as index.
//
// This request will be handled on the text_input side directly following
// a commit_string request.
//
func (i *InputMethodContext) CursorPosition(index, anchor int32) error {
	const opcode = 6
	const rLen = 8 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(index))
	l += 4
	client.PutUint32(r[l:l+4], uint32(anchor))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// ModifiersMap :
//
func (i *InputMethodContext) ModifiersMap(_map []byte) error {
	const opcode = 7
	_mapLen := len(_map)
	rLen := 8 + _mapLen
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutArray(r[l:l+(4+_mapLen)], _map)
	l += _mapLen
	err := i.Context().WriteMsg(r, nil)
	return err
}

// Keysym : keysym
//
// Notify when a key event was sent. Key events should not be used for
// normal text input operations, which should be done with commit_string,
// delete_surrounding_text, etc. The key event follows the wl_keyboard key
// event convention. Sym is an XKB keysym, state is a wl_keyboard key_state.
//
//  serial: serial of the latest known text input state
func (i *InputMethodContext) Keysym(serial, time, sym, state, modifiers uint32) error {
	const opcode = 8
	const rLen = 8 + 4 + 4 + 4 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(serial))
	l += 4
	client.PutUint32(r[l:l+4], uint32(time))
	l += 4
	client.PutUint32(r[l:l+4], uint32(sym))
	l += 4
	client.PutUint32(r[l:l+4], uint32(state))
	l += 4
	client.PutUint32(r[l:l+4], uint32(modifiers))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// GrabKeyboard : grab hardware keyboard
//
// Allow an input method to receive hardware keyboard input and process
// key events to generate text events (with pre-edit) over the wire. This
// allows input methods which compose multiple key events for inputting
// text like it is done for CJK languages.
//
func (i *InputMethodContext) GrabKeyboard() (*client.Keyboard, error) {
	keyboard := client.NewKeyboard(i.Context())
	const opcode = 9
	const rLen = 8 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], keyboard.ID())
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return keyboard, err
}

// Key : forward key event
//
// Forward a wl_keyboard::key event to the client that was not processed
// by the input method itself. Should be used when filtering key events
// with grab_keyboard.  The arguments should be the ones from the
// wl_keyboard::key event.
//
// For generating custom key events use the keysym request instead.
//
//  serial: serial from wl_keyboard::key
//  time: time from wl_keyboard::key
//  key: key from wl_keyboard::key
//  state: state from wl_keyboard::key
func (i *InputMethodContext) Key(serial, time, key, state uint32) error {
	const opcode = 10
	const rLen = 8 + 4 + 4 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(serial))
	l += 4
	client.PutUint32(r[l:l+4], uint32(time))
	l += 4
	client.PutUint32(r[l:l+4], uint32(key))
	l += 4
	client.PutUint32(r[l:l+4], uint32(state))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// Modifiers : forward modifiers event
//
// Forward a wl_keyboard::modifiers event to the client that was not
// processed by the input method itself.  Should be used when filtering
// key events with grab_keyboard. The arguments should be the ones
// from the wl_keyboard::modifiers event.
//
//  serial: serial from wl_keyboard::modifiers
//  modsDepressed: mods_depressed from wl_keyboard::modifiers
//  modsLatched: mods_latched from wl_keyboard::modifiers
//  modsLocked: mods_locked from wl_keyboard::modifiers
//  group: group from wl_keyboard::modifiers
func (i *InputMethodContext) Modifiers(serial, modsDepressed, modsLatched, modsLocked, group uint32) error {
	const opcode = 11
	const rLen = 8 + 4 + 4 + 4 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(serial))
	l += 4
	client.PutUint32(r[l:l+4], uint32(modsDepressed))
	l += 4
	client.PutUint32(r[l:l+4], uint32(modsLatched))
	l += 4
	client.PutUint32(r[l:l+4], uint32(modsLocked))
	l += 4
	client.PutUint32(r[l:l+4], uint32(group))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// Language :
//
//  serial: serial of the latest known text input state
func (i *InputMethodContext) Language(serial uint32, language string) error {
	const opcode = 12
	languageLen := client.PaddedLen(len(language) + 1)
	rLen := 8 + 4 + (4 + languageLen)
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(serial))
	l += 4
	client.PutString(r[l:l+(4+languageLen)], language, languageLen)
	l += (4 + languageLen)
	err := i.Context().WriteMsg(r, nil)
	return err
}

// TextDirection :
//
//  serial: serial of the latest known text input state
func (i *InputMethodContext) TextDirection(serial, direction uint32) error {
	const opcode = 13
	const rLen = 8 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], uint32(serial))
	l += 4
	client.PutUint32(r[l:l+4], uint32(direction))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// InputMethodContextSurroundingTextEvent : surrounding text event
//
// The plain surrounding text around the input position. Cursor is the
// position in bytes within the surrounding text relative to the beginning
// of the text. Anchor is the position in bytes of the selection anchor
// within the surrounding text relative to the beginning of the text. If
// there is no selected text then anchor is the same as cursor.
type InputMethodContextSurroundingTextEvent struct {
	Text   string
	Cursor uint32
	Anchor uint32
}
type InputMethodContextSurroundingTextHandlerFunc func(InputMethodContextSurroundingTextEvent)

// AddSurroundingTextHandler : adds handler for InputMethodContextSurroundingTextEvent
func (i *InputMethodContext) AddSurroundingTextHandler(f InputMethodContextSurroundingTextHandlerFunc) {
	if f == nil {
		return
	}

	i.surroundingTextHandlers = append(i.surroundingTextHandlers, f)
}

// InputMethodContextResetEvent :
type InputMethodContextResetEvent struct{}
type InputMethodContextResetHandlerFunc func(InputMethodContextResetEvent)

// AddResetHandler : adds handler for InputMethodContextResetEvent
func (i *InputMethodContext) AddResetHandler(f InputMethodContextResetHandlerFunc) {
	if f == nil {
		return
	}

	i.resetHandlers = append(i.resetHandlers, f)
}

// InputMethodContextContentTypeEvent :
type InputMethodContextContentTypeEvent struct {
	Hint    uint32
	Purpose uint32
}
type InputMethodContextContentTypeHandlerFunc func(InputMethodContextContentTypeEvent)

// AddContentTypeHandler : adds handler for InputMethodContextContentTypeEvent
func (i *InputMethodContext) AddContentTypeHandler(f InputMethodContextContentTypeHandlerFunc) {
	if f == nil {
		return
	}

	i.contentTypeHandlers = append(i.contentTypeHandlers, f)
}

// InputMethodContextInvokeActionEvent :
type InputMethodContextInvokeActionEvent struct {
	Button uint32
	Index  uint32
}
type InputMethodContextInvokeActionHandlerFunc func(InputMethodContextInvokeActionEvent)

// AddInvokeActionHandler : adds handler for InputMethodContextInvokeActionEvent
func (i *InputMethodContext) AddInvokeActionHandler(f InputMethodContextInvokeActionHandlerFunc) {
	if f == nil {
		return
	}

	i.invokeActionHandlers = append(i.invokeActionHandlers, f)
}

// InputMethodContextCommitStateEvent :
type InputMethodContextCommitStateEvent struct {
	Serial uint32
}
type InputMethodContextCommitStateHandlerFunc func(InputMethodContextCommitStateEvent)

// AddCommitStateHandler : adds handler for InputMethodContextCommitStateEvent
func (i *InputMethodContext) AddCommitStateHandler(f InputMethodContextCommitStateHandlerFunc) {
	if f == nil {
		return
	}

	i.commitStateHandlers = append(i.commitStateHandlers, f)
}

// InputMethodContextPreferredLanguageEvent :
type InputMethodContextPreferredLanguageEvent struct {
	Language string
}
type InputMethodContextPreferredLanguageHandlerFunc func(InputMethodContextPreferredLanguageEvent)

// AddPreferredLanguageHandler : adds handler for InputMethodContextPreferredLanguageEvent
func (i *InputMethodContext) AddPreferredLanguageHandler(f InputMethodContextPreferredLanguageHandlerFunc) {
	if f == nil {
		return
	}

	i.preferredLanguageHandlers = append(i.preferredLanguageHandlers, f)
}

func (i *InputMethodContext) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.surroundingTextHandlers) == 0 {
			return
		}
		var e InputMethodContextSurroundingTextEvent
		l := 0
		textLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Text = client.String(data[l : l+textLen])
		l += textLen
		e.Cursor = client.Uint32(data[l : l+4])
		l += 4
		e.Anchor = client.Uint32(data[l : l+4])
		l += 4
		for _, f := range i.surroundingTextHandlers {
			f(e)
		}
	case 1:
		if len(i.resetHandlers) == 0 {
			return
		}
		var e InputMethodContextResetEvent
		for _, f := range i.resetHandlers {
			f(e)
		}
	case 2:
		if len(i.contentTypeHandlers) == 0 {
			return
		}
		var e InputMethodContextContentTypeEvent
		l := 0
		e.Hint = client.Uint32(data[l : l+4])
		l += 4
		e.Purpose = client.Uint32(data[l : l+4])
		l += 4
		for _, f := range i.contentTypeHandlers {
			f(e)
		}
	case 3:
		if len(i.invokeActionHandlers) == 0 {
			return
		}
		var e InputMethodContextInvokeActionEvent
		l := 0
		e.Button = client.Uint32(data[l : l+4])
		l += 4
		e.Index = client.Uint32(data[l : l+4])
		l += 4
		for _, f := range i.invokeActionHandlers {
			f(e)
		}
	case 4:
		if len(i.commitStateHandlers) == 0 {
			return
		}
		var e InputMethodContextCommitStateEvent
		l := 0
		e.Serial = client.Uint32(data[l : l+4])
		l += 4
		for _, f := range i.commitStateHandlers {
			f(e)
		}
	case 5:
		if len(i.preferredLanguageHandlers) == 0 {
			return
		}
		var e InputMethodContextPreferredLanguageEvent
		l := 0
		languageLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Language = client.String(data[l : l+languageLen])
		l += languageLen
		for _, f := range i.preferredLanguageHandlers {
			f(e)
		}
	}
}

// InputMethod : input method
//
// An input method object is responsible for composing text in response to
// input from hardware or virtual keyboards. There is one input method
// object per seat. On activate there is a new input method context object
// created which allows the input method to communicate with the text input.
type InputMethod struct {
	client.BaseProxy
	activateHandlers   []InputMethodActivateHandlerFunc
	deactivateHandlers []InputMethodDeactivateHandlerFunc
}

// NewInputMethod : input method
//
// An input method object is responsible for composing text in response to
// input from hardware or virtual keyboards. There is one input method
// object per seat. On activate there is a new input method context object
// created which allows the input method to communicate with the text input.
func NewInputMethod(ctx *client.Context) *InputMethod {
	zwpInputMethodV1 := &InputMethod{}
	ctx.Register(zwpInputMethodV1)
	return zwpInputMethodV1
}

func (i *InputMethod) Destroy() error {
	i.Context().Unregister(i)
	return nil
}

// InputMethodActivateEvent : activate event
//
// A text input was activated. Creates an input method context object
// which allows communication with the text input.
type InputMethodActivateEvent struct {
	Id *InputMethodContext
}
type InputMethodActivateHandlerFunc func(InputMethodActivateEvent)

// AddActivateHandler : adds handler for InputMethodActivateEvent
func (i *InputMethod) AddActivateHandler(f InputMethodActivateHandlerFunc) {
	if f == nil {
		return
	}

	i.activateHandlers = append(i.activateHandlers, f)
}

// InputMethodDeactivateEvent : deactivate event
//
// The text input corresponding to the context argument was deactivated.
// The input method context should be destroyed after deactivation is
// handled.
type InputMethodDeactivateEvent struct {
	Context *InputMethodContext
}
type InputMethodDeactivateHandlerFunc func(InputMethodDeactivateEvent)

// AddDeactivateHandler : adds handler for InputMethodDeactivateEvent
func (i *InputMethod) AddDeactivateHandler(f InputMethodDeactivateHandlerFunc) {
	if f == nil {
		return
	}

	i.deactivateHandlers = append(i.deactivateHandlers, f)
}

func (i *InputMethod) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.activateHandlers) == 0 {
			return
		}
		var e InputMethodActivateEvent
		l := 0
		e.Id = i.Context().GetProxy(client.Uint32(data[l : l+4])).(*InputMethodContext)
		l += 4
		for _, f := range i.activateHandlers {
			f(e)
		}
	case 1:
		if len(i.deactivateHandlers) == 0 {
			return
		}
		var e InputMethodDeactivateEvent
		l := 0
		e.Context = i.Context().GetProxy(client.Uint32(data[l : l+4])).(*InputMethodContext)
		l += 4
		for _, f := range i.deactivateHandlers {
			f(e)
		}
	}
}

// InputPanel : interface for implementing keyboards
//
// Only one client can bind this interface at a time.
type InputPanel struct {
	client.BaseProxy
}

// NewInputPanel : interface for implementing keyboards
//
// Only one client can bind this interface at a time.
func NewInputPanel(ctx *client.Context) *InputPanel {
	zwpInputPanelV1 := &InputPanel{}
	ctx.Register(zwpInputPanelV1)
	return zwpInputPanelV1
}

// GetInputPanelSurface :
//
func (i *InputPanel) GetInputPanelSurface(surface *client.Surface) (*InputPanelSurface, error) {
	id := NewInputPanelSurface(i.Context())
	const opcode = 0
	const rLen = 8 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], id.ID())
	l += 4
	client.PutUint32(r[l:l+4], surface.ID())
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return id, err
}

func (i *InputPanel) Destroy() error {
	i.Context().Unregister(i)
	return nil
}

// InputPanelSurface :
type InputPanelSurface struct {
	client.BaseProxy
}

// NewInputPanelSurface :
func NewInputPanelSurface(ctx *client.Context) *InputPanelSurface {
	zwpInputPanelSurfaceV1 := &InputPanelSurface{}
	ctx.Register(zwpInputPanelSurfaceV1)
	return zwpInputPanelSurfaceV1
}

// SetToplevel : set the surface type as a keyboard
//
// Set the input_panel_surface type to keyboard.
//
// A keyboard surface is only shown when a text input is active.
//
func (i *InputPanelSurface) SetToplevel(output *client.Output, position uint32) error {
	const opcode = 0
	const rLen = 8 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], output.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(position))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// SetOverlayPanel : set the surface type as an overlay panel
//
// Set the input_panel_surface to be an overlay panel.
//
// This is shown near the input cursor above the application window when
// a text input is active.
//
func (i *InputPanelSurface) SetOverlayPanel() error {
	const opcode = 1
	const rLen = 8
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

func (i *InputPanelSurface) Destroy() error {
	i.Context().Unregister(i)
	return nil
}

type InputPanelSurfacePosition uint32

// InputPanelSurfacePosition :
const (
	InputPanelSurfacePositionCenterBottom InputPanelSurfacePosition = 0
)

func (e InputPanelSurfacePosition) Name() string {
	switch e {
	case InputPanelSurfacePositionCenterBottom:
		return "center_bottom"
	default:
		return ""
	}
}

func (e InputPanelSurfacePosition) Value() string {
	switch e {
	case InputPanelSurfacePositionCenterBottom:
		return "0"
	default:
		return ""
	}
}

func (e InputPanelSurfacePosition) String() string {
	return e.Name() + "=" + e.Value()
}
