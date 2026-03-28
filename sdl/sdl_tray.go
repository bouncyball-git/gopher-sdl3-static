package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

extern void goTrayCallback(void *userdata, SDL_TrayEntry *entry);

static void cgoTrayCallbackTrampoline(void *userdata, SDL_TrayEntry *entry) {
	goTrayCallback(userdata, entry);
}
*/
import "C"

import "unsafe"

// Tray represents a system tray icon.
type Tray struct {
	c *C.SDL_Tray
}

// TrayMenu represents a tray menu.
type TrayMenu struct {
	c *C.SDL_TrayMenu
}

// TrayEntry represents an entry in a tray menu.
type TrayEntry struct {
	c *C.SDL_TrayEntry
}

// TrayEntryFlags represents flags for tray entries.
type TrayEntryFlags uint32

// Tray entry flag values.
const (
	TRAYENTRY_BUTTON   TrayEntryFlags = C.SDL_TRAYENTRY_BUTTON
	TRAYENTRY_CHECKBOX TrayEntryFlags = C.SDL_TRAYENTRY_CHECKBOX
	TRAYENTRY_SUBMENU  TrayEntryFlags = C.SDL_TRAYENTRY_SUBMENU
	TRAYENTRY_DISABLED TrayEntryFlags = C.SDL_TRAYENTRY_DISABLED
	TRAYENTRY_CHECKED  TrayEntryFlags = C.SDL_TRAYENTRY_CHECKED
)

// CreateTray creates a system tray icon.
func CreateTray(icon *Surface, tooltip string) *Tray {
	var ci *C.SDL_Surface
	if icon != nil {
		ci = icon.c
	}
	ct := C.CString(tooltip)
	defer C.free(unsafe.Pointer(ct))
	t := C.SDL_CreateTray(ci, ct)
	if t == nil {
		return nil
	}
	return &Tray{c: t}
}

// Destroy destroys the tray.
func (t *Tray) Destroy() {
	if t.c != nil {
		C.SDL_DestroyTray(t.c)
		t.c = nil
	}
}

// SetIcon sets the tray icon.
func (t *Tray) SetIcon(icon *Surface) {
	var ci *C.SDL_Surface
	if icon != nil {
		ci = icon.c
	}
	C.SDL_SetTrayIcon(t.c, ci)
}

// SetTooltip sets the tray tooltip.
func (t *Tray) SetTooltip(tooltip string) {
	ct := C.CString(tooltip)
	defer C.free(unsafe.Pointer(ct))
	C.SDL_SetTrayTooltip(t.c, ct)
}

// CreateMenu creates a menu for the tray.
func (t *Tray) CreateMenu() *TrayMenu {
	m := C.SDL_CreateTrayMenu(t.c)
	if m == nil {
		return nil
	}
	return &TrayMenu{c: m}
}

// GetMenu returns the tray's menu.
func (t *Tray) GetMenu() *TrayMenu {
	m := C.SDL_GetTrayMenu(t.c)
	if m == nil {
		return nil
	}
	return &TrayMenu{c: m}
}

// InsertEntry inserts an entry at the specified position. Use -1 to append.
func (m *TrayMenu) InsertEntry(pos int, label string, flags TrayEntryFlags) *TrayEntry {
	cl := C.CString(label)
	defer C.free(unsafe.Pointer(cl))
	e := C.SDL_InsertTrayEntryAt(m.c, C.int(pos), cl, C.SDL_TrayEntryFlags(flags))
	if e == nil {
		return nil
	}
	return &TrayEntry{c: e}
}

// Remove removes the entry from its menu.
func (e *TrayEntry) Remove() {
	C.SDL_RemoveTrayEntry(e.c)
}

// SetLabel sets the label of the entry.
func (e *TrayEntry) SetLabel(label string) {
	cl := C.CString(label)
	defer C.free(unsafe.Pointer(cl))
	C.SDL_SetTrayEntryLabel(e.c, cl)
}

// Label returns the label of the entry.
func (e *TrayEntry) Label() string {
	return C.GoString(C.SDL_GetTrayEntryLabel(e.c))
}

// SetChecked sets the checked state.
func (e *TrayEntry) SetChecked(checked bool) {
	C.SDL_SetTrayEntryChecked(e.c, C.bool(checked))
}

// Checked returns whether the entry is checked.
func (e *TrayEntry) Checked() bool {
	return bool(C.SDL_GetTrayEntryChecked(e.c))
}

// SetEnabled sets whether the entry is enabled.
func (e *TrayEntry) SetEnabled(enabled bool) {
	C.SDL_SetTrayEntryEnabled(e.c, C.bool(enabled))
}

// Enabled returns whether the entry is enabled.
func (e *TrayEntry) Enabled() bool {
	return bool(C.SDL_GetTrayEntryEnabled(e.c))
}

// CreateSubmenu creates a submenu for this entry.
func (e *TrayEntry) CreateSubmenu() *TrayMenu {
	m := C.SDL_CreateTraySubmenu(e.c)
	if m == nil {
		return nil
	}
	return &TrayMenu{c: m}
}

// GetSubmenu returns the submenu of this entry.
func (e *TrayEntry) GetSubmenu() *TrayMenu {
	m := C.SDL_GetTraySubmenu(e.c)
	if m == nil {
		return nil
	}
	return &TrayMenu{c: m}
}

// UpdateTrays processes tray events.
func UpdateTrays() {
	C.SDL_UpdateTrays()
}

// CreateTrayWithProperties creates a system tray icon with properties.
func CreateTrayWithProperties(props PropertiesID) *Tray {
	t := C.SDL_CreateTrayWithProperties(C.SDL_PropertiesID(props))
	if t == nil {
		return nil
	}
	return &Tray{c: t}
}

// GetEntries returns all entries in a tray menu.
func (m *TrayMenu) GetEntries() []*TrayEntry {
	var count C.int
	entries := C.SDL_GetTrayEntries(m.c, &count)
	if entries == nil {
		return nil
	}
	n := int(count)
	result := make([]*TrayEntry, n)
	slice := unsafe.Slice((**C.SDL_TrayEntry)(unsafe.Pointer(entries)), n)
	for i, e := range slice {
		result[i] = &TrayEntry{c: e}
	}
	return result
}

// Parent returns the parent menu of a tray entry.
func (e *TrayEntry) Parent() *TrayMenu {
	m := C.SDL_GetTrayEntryParent(e.c)
	if m == nil {
		return nil
	}
	return &TrayMenu{c: m}
}

// ParentEntry returns the parent entry of a tray menu (for submenus).
func (m *TrayMenu) ParentEntry() *TrayEntry {
	e := C.SDL_GetTrayMenuParentEntry(m.c)
	if e == nil {
		return nil
	}
	return &TrayEntry{c: e}
}

// ParentTray returns the parent tray of a tray menu.
func (m *TrayMenu) ParentTray() *Tray {
	t := C.SDL_GetTrayMenuParentTray(m.c)
	if t == nil {
		return nil
	}
	return &Tray{c: t}
}

// Click simulates a click on a tray entry.
func (e *TrayEntry) Click() {
	C.SDL_ClickTrayEntry(e.c)
}

// TrayCallbackFunc is called when a tray entry is activated.
type TrayCallbackFunc func(entry *TrayEntry)

//export goTrayCallback
func goTrayCallback(userdata unsafe.Pointer, entry *C.SDL_TrayEntry) {
	id := uintptr(userdata)
	fn := getCallback(id).(TrayCallbackFunc)
	fn(&TrayEntry{c: entry})
}

// SetCallback sets a callback for the tray entry.
func (e *TrayEntry) SetCallback(callback TrayCallbackFunc) {
	if callback == nil {
		C.SDL_SetTrayEntryCallback(e.c, nil, nil)
		return
	}
	id := registerCallback(callback)
	C.SDL_SetTrayEntryCallback(e.c, C.SDL_TrayCallback(C.cgoTrayCallbackTrampoline), ptrFromID(id))
}
