package hare

import (
// "github.com/robertkrimen/otto"
)

type View struct {
	db *DB

	Record `capid:"skip"`
	Src    string `capid:"0"`
}

func (view *View) Value() string {
	return ""
}

func (view *View) BucketName() string {
	return "views"
}

// vm := otto.New()
// vm.Run(`
//     abc = 2 + 2;
//     console.log("The value of abc is " + abc); // 4
// `)
