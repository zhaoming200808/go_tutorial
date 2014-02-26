package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)	//get reflect_obj s
	fn := st.Elem()
	field := st.Field(0)	//get filst_type F
	color_value := field.Tag.Get("color") //get s.F species 
	species_value := field.Tag.Get("species")
	nill_value := field.Tag.Get("NULL")
	fmt.Printf(" =|%v|= , =|%v|= , =|%v|= \n",color_value, species_value,nill_value)

}

