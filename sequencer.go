package hare

type Sequencer struct {
  index   uint64
}

func (seq *Sequencer) Next() *Key {
  seq.index++
  return NewKey(seq.index)
}
