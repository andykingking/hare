package hare

import (
	"encoding/json"
)

type Document struct {
	Record `capid:"skip"`
	Body   string `capid:"0"`
}

func (doc *Document) SetBody(body string) (err error) {
	// validate JSON
	var tmp interface{}
	err = json.Unmarshal([]byte(body), &tmp)
	if err != nil {
		return err
	}

	doc.Body = body
	return nil
}

func (doc *Document) BucketName() []byte {
	return []byte("data")
}
