package utility

import (
	"encoding/json"
)

//ReadJSON reads a json content and retransmits the data in a JsonWebF variable
func ReadJSON(data []byte) (jsonWebF JSONWebF) {

	errr := json.Unmarshal(data, &jsonWebF)
	Check(errr)

	return
}
