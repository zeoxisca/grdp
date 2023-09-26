package cliprdr

import (
	"github.com/tomatome/win"
)

const (
	CFSTR_FILECONTENTS    = "FileContents"
	CFSTR_FILEDESCRIPTORW = "FileGroupDescriptorW"
	CF_HDROP              = 15
	DVASPECT_CONTENT      = 0x1
)

type Control struct {
	hwnd uintptr
}

func (c *Control) withOpenClipboard(f func()) {
	if OpenClipboard(c.hwnd) {
		f()
		CloseClipboard()
	}
}
func ClipWatcher(c *CliprdrClient) {
}
func OpenClipboard(hwnd uintptr) bool {
	return false
}
func CloseClipboard() bool {
	return false
}
func CountClipboardFormats() int32 {
	return 0
}
func IsClipboardFormatAvailable(id uint32) bool {
	return false
}
func EnumClipboardFormats(formatId uint32) uint32 {
	return uint32(0)
}
func GetClipboardFormatName(id uint32) string {
	return ""
}
func EmptyClipboard() bool {
	return false
}
func RegisterClipboardFormat(format string) uint32 {
	return 0
}
func IsClipboardOwner(h win.HWND) bool {
	return false
}

func HmemAlloc(data []byte) uintptr {
	return uintptr(0)
}
func SetClipboardData(formatId uint32, hmem uintptr) bool {
	return true
}
func GetClipboardData(formatId uint32) string {
	return ""
}

func GetFormatList(hwnd uintptr) []CliprdrFormat {
	list := make([]CliprdrFormat, 0, 10)
	if OpenClipboard(hwnd) {
		n := CountClipboardFormats()
		if IsClipboardFormatAvailable(CF_HDROP) {
			formatId := RegisterClipboardFormat(CFSTR_FILEDESCRIPTORW)
			var f CliprdrFormat
			f.FormatId = formatId
			f.FormatName = CFSTR_FILEDESCRIPTORW
			list = append(list, f)
			formatId = RegisterClipboardFormat(CFSTR_FILECONTENTS)
			var f1 CliprdrFormat
			f1.FormatId = formatId
			f1.FormatName = CFSTR_FILECONTENTS
			list = append(list, f1)
		} else {
			var id uint32
			for i := 0; i < int(n); i++ {
				id = EnumClipboardFormats(id)
				name := GetClipboardFormatName(id)
				var f CliprdrFormat
				f.FormatId = id
				f.FormatName = name
				list = append(list, f)
			}
		}
		CloseClipboard()
	}
	return list
}

func GlobalSize(hMem uintptr) win.SIZE_T {
	return 0
}
func GlobalLock(hMem uintptr) uintptr {
	return uintptr(0)
}
func GlobalUnlock(hMem uintptr) {
}

func (c *Control) SendCliprdrMessage() {
}
func GetFileInfo(sys interface{}) (uint32, []byte, uint32, uint32) {

	return 0, []byte(""), 0, 0
}

func GetFileNames() []string {
	return []string{}
}

const (
	FILE_ATTRIBUTE_DIRECTORY = 0x00000010
)

type DROPFILES struct {
	pFiles uintptr
	pt     uintptr
	fNC    bool
	fWide  bool
}
