package main

type rectangle struct {
	width, height int
}

//This can mutate the reciever struct as it be reference
func (r *rectangle) canMutate(w, h int) {
	r.width = w
	r.height = h
}

//This can't mutate the reciever strct as it is by Value
func (r rectangle) cantMutate(w, h int) {
	r.width = w
	r.height = h
}
