package hare

type Fact struct {
	Record `capid:"skip"`

	KView   uint64 `capid:"0"`
	KDomain uint64 `capid:"1"`
	KTarget uint64 `capid:"2"`
}

func NewFact(view *View, domain *Document, target *Document) *Fact {
	if !(view.Saved() && domain.Saved() && target.Saved()) {
		return nil
	}

	return &Fact{KView: view.K, KDomain: domain.K, KTarget: target.K}
}

func (fact *Fact) BucketName() []byte {
	return append([]byte("facts:"), KeyToBytes(fact.KView)...)
}

func (fact *Fact) Key() []byte {
	return append(KeyToBytes(fact.KDomain), KeyToBytes(fact.KTarget)...)
}
