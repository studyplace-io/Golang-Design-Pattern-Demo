package subject_object_pattern

import "testing"

func TestSubjectObjectPattern(t *testing.T) {

	stock := &Stock{}
	issue := &Issue{}

	order := &Order{
		Id: 11,
		Observers: []Observer{stock, issue},
	}

	order.Cancel()

}
