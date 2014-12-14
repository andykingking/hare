package hare

type Sequencer struct {
	Index uint64
	Key
}

func (seq *Sequencer) Next() uint64 {
	seq.Index++
	return seq.Index
}
