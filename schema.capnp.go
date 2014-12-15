package hare

// AUTO GENERATED - DO NOT EDIT

import (
	C "github.com/glycerine/go-capnproto"
	"unsafe"
)

type DocumentCapn C.Struct

func NewDocumentCapn(s *C.Segment) DocumentCapn      { return DocumentCapn(s.NewStruct(0, 1)) }
func NewRootDocumentCapn(s *C.Segment) DocumentCapn  { return DocumentCapn(s.NewRootStruct(0, 1)) }
func AutoNewDocumentCapn(s *C.Segment) DocumentCapn  { return DocumentCapn(s.NewStructAR(0, 1)) }
func ReadRootDocumentCapn(s *C.Segment) DocumentCapn { return DocumentCapn(s.Root(0).ToStruct()) }
func (s DocumentCapn) Body() string                  { return C.Struct(s).GetObject(0).ToText() }
func (s DocumentCapn) SetBody(v string)              { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s DocumentCapn) MarshalJSON() (bs []byte, err error) { return }

type DocumentCapn_List C.PointerList

func NewDocumentCapnList(s *C.Segment, sz int) DocumentCapn_List {
	return DocumentCapn_List(s.NewCompositeList(0, 1, sz))
}
func (s DocumentCapn_List) Len() int { return C.PointerList(s).Len() }
func (s DocumentCapn_List) At(i int) DocumentCapn {
	return DocumentCapn(C.PointerList(s).At(i).ToStruct())
}
func (s DocumentCapn_List) ToArray() []DocumentCapn {
	return *(*[]DocumentCapn)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s DocumentCapn_List) Set(i int, item DocumentCapn) { C.PointerList(s).Set(i, C.Object(item)) }

type KeyCapn C.Struct

func NewKeyCapn(s *C.Segment) KeyCapn      { return KeyCapn(s.NewStruct(8, 0)) }
func NewRootKeyCapn(s *C.Segment) KeyCapn  { return KeyCapn(s.NewRootStruct(8, 0)) }
func AutoNewKeyCapn(s *C.Segment) KeyCapn  { return KeyCapn(s.NewStructAR(8, 0)) }
func ReadRootKeyCapn(s *C.Segment) KeyCapn { return KeyCapn(s.Root(0).ToStruct()) }
func (s KeyCapn) Uint64() uint64           { return C.Struct(s).Get64(0) }
func (s KeyCapn) SetUint64(v uint64)       { C.Struct(s).Set64(0, v) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s KeyCapn) MarshalJSON() (bs []byte, err error) { return }

type KeyCapn_List C.PointerList

func NewKeyCapnList(s *C.Segment, sz int) KeyCapn_List { return KeyCapn_List(s.NewUInt64List(sz)) }
func (s KeyCapn_List) Len() int                        { return C.PointerList(s).Len() }
func (s KeyCapn_List) At(i int) KeyCapn                { return KeyCapn(C.PointerList(s).At(i).ToStruct()) }
func (s KeyCapn_List) ToArray() []KeyCapn {
	return *(*[]KeyCapn)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s KeyCapn_List) Set(i int, item KeyCapn) { C.PointerList(s).Set(i, C.Object(item)) }

type SequencerCapn C.Struct

func NewSequencerCapn(s *C.Segment) SequencerCapn      { return SequencerCapn(s.NewStruct(8, 0)) }
func NewRootSequencerCapn(s *C.Segment) SequencerCapn  { return SequencerCapn(s.NewRootStruct(8, 0)) }
func AutoNewSequencerCapn(s *C.Segment) SequencerCapn  { return SequencerCapn(s.NewStructAR(8, 0)) }
func ReadRootSequencerCapn(s *C.Segment) SequencerCapn { return SequencerCapn(s.Root(0).ToStruct()) }
func (s SequencerCapn) Index() int64                   { return int64(C.Struct(s).Get64(0)) }
func (s SequencerCapn) SetIndex(v int64)               { C.Struct(s).Set64(0, uint64(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s SequencerCapn) MarshalJSON() (bs []byte, err error) { return }

type SequencerCapn_List C.PointerList

func NewSequencerCapnList(s *C.Segment, sz int) SequencerCapn_List {
	return SequencerCapn_List(s.NewUInt64List(sz))
}
func (s SequencerCapn_List) Len() int { return C.PointerList(s).Len() }
func (s SequencerCapn_List) At(i int) SequencerCapn {
	return SequencerCapn(C.PointerList(s).At(i).ToStruct())
}
func (s SequencerCapn_List) ToArray() []SequencerCapn {
	return *(*[]SequencerCapn)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s SequencerCapn_List) Set(i int, item SequencerCapn) { C.PointerList(s).Set(i, C.Object(item)) }

type ViewCapn C.Struct

func NewViewCapn(s *C.Segment) ViewCapn      { return ViewCapn(s.NewStruct(0, 1)) }
func NewRootViewCapn(s *C.Segment) ViewCapn  { return ViewCapn(s.NewRootStruct(0, 1)) }
func AutoNewViewCapn(s *C.Segment) ViewCapn  { return ViewCapn(s.NewStructAR(0, 1)) }
func ReadRootViewCapn(s *C.Segment) ViewCapn { return ViewCapn(s.Root(0).ToStruct()) }
func (s ViewCapn) Src() string               { return C.Struct(s).GetObject(0).ToText() }
func (s ViewCapn) SetSrc(v string)           { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s ViewCapn) MarshalJSON() (bs []byte, err error) { return }

type ViewCapn_List C.PointerList

func NewViewCapnList(s *C.Segment, sz int) ViewCapn_List {
	return ViewCapn_List(s.NewCompositeList(0, 1, sz))
}
func (s ViewCapn_List) Len() int          { return C.PointerList(s).Len() }
func (s ViewCapn_List) At(i int) ViewCapn { return ViewCapn(C.PointerList(s).At(i).ToStruct()) }
func (s ViewCapn_List) ToArray() []ViewCapn {
	return *(*[]ViewCapn)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s ViewCapn_List) Set(i int, item ViewCapn) { C.PointerList(s).Set(i, C.Object(item)) }
