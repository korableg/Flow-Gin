package errs

import (
	"encoding/json"
	"testing"
)

func TestErrs(t *testing.T) {

	e := New(ERR_NODE_NOT_FOUND)

	if ERR_NODE_NOT_FOUND.Error() != e.Error() {
		t.Errorf("error must be %s", ERR_NODE_NOT_FOUND.Error())
	}

	_, err := json.Marshal(e)
	if err != nil {
		t.Error(err)
	}

}
