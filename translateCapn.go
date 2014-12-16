package hare

import (
  capn "github.com/glycerine/go-capnproto"
  "io"
  "fmt"
)




func (s *Document) Save(w io.Writer) {
  	seg := capn.NewBuffer(nil)
  	DocumentGoToCapn(seg, s)
  	seg.WriteTo(w)
}
 

 
func (s *Document) Load(r io.Reader) {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
  	}
  	z := ReadRootDocumentCapn(capMsg)
      DocumentCapnToGo(z, s)
}



func DocumentCapnToGo(src DocumentCapn, dest *Document) *Document { 
  if dest == nil { 
    dest = &Document{} 
  }
  dest.Body = src.Body()

  return dest
} 



func DocumentGoToCapn(seg *capn.Segment, src *Document) DocumentCapn { 
  dest := AutoNewDocumentCapn(seg)
  dest.SetBody(src.Body)

  return dest
} 



func (s *Fact) Save(w io.Writer) {
  	seg := capn.NewBuffer(nil)
  	FactGoToCapn(seg, s)
  	seg.WriteTo(w)
}
 

 
func (s *Fact) Load(r io.Reader) {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
  	}
  	z := ReadRootFactCapn(capMsg)
      FactCapnToGo(z, s)
}



func FactCapnToGo(src FactCapn, dest *Fact) *Fact { 
  if dest == nil { 
    dest = &Fact{} 
  }
  dest.KView = uint64(src.KView())
  dest.KDomain = uint64(src.KDomain())
  dest.KTarget = uint64(src.KTarget())

  return dest
} 



func FactGoToCapn(seg *capn.Segment, src *Fact) FactCapn { 
  dest := AutoNewFactCapn(seg)
  dest.SetKView(uint64(src.KView))
  dest.SetKDomain(uint64(src.KDomain))
  dest.SetKTarget(uint64(src.KTarget))

  return dest
} 



func (s *Sequencer) Save(w io.Writer) {
  	seg := capn.NewBuffer(nil)
  	SequencerGoToCapn(seg, s)
  	seg.WriteTo(w)
}
 

 
func (s *Sequencer) Load(r io.Reader) {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
  	}
  	z := ReadRootSequencerCapn(capMsg)
      SequencerCapnToGo(z, s)
}



func SequencerCapnToGo(src SequencerCapn, dest *Sequencer) *Sequencer { 
  if dest == nil { 
    dest = &Sequencer{} 
  }
  dest.Mark = uint64(src.Mark())

  return dest
} 



func SequencerGoToCapn(seg *capn.Segment, src *Sequencer) SequencerCapn { 
  dest := AutoNewSequencerCapn(seg)
  dest.SetMark(uint64(src.Mark))

  return dest
} 



func (s *View) Save(w io.Writer) {
  	seg := capn.NewBuffer(nil)
  	ViewGoToCapn(seg, s)
  	seg.WriteTo(w)
}
 

 
func (s *View) Load(r io.Reader) {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
  	}
  	z := ReadRootViewCapn(capMsg)
      ViewCapnToGo(z, s)
}



func ViewCapnToGo(src ViewCapn, dest *View) *View { 
  if dest == nil { 
    dest = &View{} 
  }
  dest.Src = src.Src()

  return dest
} 



func ViewGoToCapn(seg *capn.Segment, src *View) ViewCapn { 
  dest := AutoNewViewCapn(seg)
  dest.SetSrc(src.Src)

  return dest
} 
