package hare

import (
  capn "github.com/glycerine/go-capnproto"
  "io"
  "fmt"
)




func (s *Key) Save(w io.Writer) {
  	seg := capn.NewBuffer(nil)
  	KeyGoToCapn(seg, s)
  	seg.WriteTo(w)
}
 

 
func (s *Key) Load(r io.Reader) {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
  	}
  	z := ReadRootKeyCapn(capMsg)
      KeyCapnToGo(z, s)
}



func KeyCapnToGo(src KeyCapn, dest *Key) *Key { 
  if dest == nil { 
    dest = &Key{} 
  }

  return dest
} 



func KeyGoToCapn(seg *capn.Segment, src *Key) KeyCapn { 
  dest := AutoNewKeyCapn(seg)

  return dest
} 
