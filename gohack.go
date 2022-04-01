package gohack

import (
	"reflect"
	"unsafe"
)

var (
	offsetGoid         uintptr
	offsetPaniconfault uintptr
	offsetLabels       uintptr
)

func init() {
	gt := getgt()
	offsetGoid = offset(gt, "goid")
	offsetPaniconfault = offset(gt, "paniconfault")
	offsetLabels = offset(gt, "labels")
}

//add pointer addition operation.
func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

func offset(t reflect.Type, f string) uintptr {
	field, found := t.FieldByName(f)
	if found {
		return field.Offset
	}
	panic("The type '" + t.Name() + "' doesn't have a field of a specified name '" + f + "'.")
}

func getg() *g {
	gp := getgp()
	if gp == nil {
		panic("Unable to get goid from runtime natively!!!")
	}
	return &g{
		goid:         *(*int64)(add(gp, offsetGoid)),
		paniconfault: (*bool)(add(gp, offsetPaniconfault)),
		labels:       (*unsafe.Pointer)(add(gp, offsetLabels)),
	}
}

type g struct {
	goid         int64
	paniconfault *bool
	labels       *unsafe.Pointer
}

func (gp *g) getPanicOnFault() bool {
	return *gp.paniconfault
}

func (gp *g) setPanicOnFault(new bool) (old bool) {
	old = *gp.paniconfault
	*gp.paniconfault = new
	return
}

func (gp *g) getLabel() unsafe.Pointer {
	return *gp.labels
}

func (gp *g) setLabel(label unsafe.Pointer) {
	*gp.labels = label
}
