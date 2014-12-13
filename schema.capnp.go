package hare

// AUTO GENERATED - DO NOT EDIT

import (
	C "github.com/glycerine/go-capnproto"
	"unsafe"
)

type KeyCapn C.Struct

func NewKeyCapn(s *C.Segment) KeyCapn      { return KeyCapn(s.NewStruct(16, 0)) }
func NewRootKeyCapn(s *C.Segment) KeyCapn  { return KeyCapn(s.NewRootStruct(16, 0)) }
func AutoNewKeyCapn(s *C.Segment) KeyCapn  { return KeyCapn(s.NewStructAR(16, 0)) }
func ReadRootKeyCapn(s *C.Segment) KeyCapn { return KeyCapn(s.Root(0).ToStruct()) }
func (s KeyCapn) Digest() uint64           { return C.Struct(s).Get64(0) }
func (s KeyCapn) SetDigest(v uint64)       { C.Struct(s).Set64(0, v) }
func (s KeyCapn) Index() uint64            { return C.Struct(s).Get64(8) }
func (s KeyCapn) SetIndex(v uint64)        { C.Struct(s).Set64(8, v) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s KeyCapn) MarshalJSON() (bs []byte, err error) { return }

type KeyCapn_List C.PointerList

func NewKeyCapnList(s *C.Segment, sz int) KeyCapn_List {
	return KeyCapn_List(s.NewCompositeList(16, 0, sz))
}
func (s KeyCapn_List) Len() int         { return C.PointerList(s).Len() }
func (s KeyCapn_List) At(i int) KeyCapn { return KeyCapn(C.PointerList(s).At(i).ToStruct()) }
func (s KeyCapn_List) ToArray() []KeyCapn {
	return *(*[]KeyCapn)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s KeyCapn_List) Set(i int, item KeyCapn) { C.PointerList(s).Set(i, C.Object(item)) }
