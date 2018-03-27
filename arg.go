package arg

import "reflect"

type test struct {

}


type Arg struct {
	Name string
	ShortName string
	Default interface{}
	Type reflect.Type
	Desc string
	Required bool
}


func (o *Arg) Int() {

}

func (o *Arg) String() {

}

func (o *Arg) Float() {

}

