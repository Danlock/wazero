package binary

import (
	"github.com/tetratelabs/wazero/wasm"
)

// EncodeModule implements wasm.EncodeModule for the WebAssembly 1.0 (MVP) Binary Format.
// Note: If saving to a file, the conventional extension is wasm
// See https://www.w3.org/TR/wasm-core-1/#binary-format%E2%91%A0
func EncodeModule(m *wasm.Module) (bytes []byte) {
	bytes = append(magic, version...)
	for name, data := range m.CustomSections {
		bytes = append(bytes, encodeCustomSection(name, data)...)
	}
	if len(m.TypeSection) > 0 {
		bytes = append(bytes, encodeTypeSection(m.TypeSection)...)
	}
	if len(m.ImportSection) > 0 {
		panic("TODO: ImportSection")
	}
	if len(m.FunctionSection) > 0 {
		panic("TODO: FunctionSection")
	}
	if len(m.TableSection) > 0 {
		panic("TODO: TableSection")
	}
	if len(m.MemorySection) > 0 {
		panic("TODO: MemorySection")
	}
	if len(m.GlobalSection) > 0 {
		panic("TODO: GlobalSection")
	}
	if len(m.ExportSection) > 0 {
		panic("TODO: ExportSection")
	}
	if m.StartSection != nil {
		panic("TODO: StartSection")
	}
	if len(m.ElementSection) > 0 {
		panic("TODO: ElementSection")
	}
	if len(m.CodeSection) > 0 {
		panic("TODO: CodeSection")
	}
	if len(m.DataSection) > 0 {
		panic("TODO: DataSection")
	}
	// >> The name section should appear only once in a module, and only after the data section.
	// See https://www.w3.org/TR/wasm-core-1/#binary-namesec
	if m.NameSection != nil {
		bytes = append(bytes, encodeCustomSection("name", encodeNameSectionData(m.NameSection))...)
	}
	return
}