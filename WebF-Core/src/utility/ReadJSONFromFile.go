package utility

import "io/ioutil"

//ReadJSONFromFile reads a .json file and retransmits the data in a JsonWebF variable
func ReadJSONFromFile(path string) (jsonWebF JSONWebF) {

	data, err := ioutil.ReadFile(path)
	Check(err)

	jsonWebF = ReadJSON(data)

	return
}
