package utility

import "encoding/json"

//WebFToJSON returns a json string with a WebF type entry
func WebFToJSON(input WebF) (output string) {
	var tmp JSONWebF

	tmp.Title = input.Title
	tmp.Content = MixPosWithWords(input.Words, input.Content)

	//Meta
	for n := 0; n < len(input.Meta); n++ {
		tmp.Meta = append(tmp.Meta, JSONMeta{
			Key:   input.Meta[n].Key,
			Value: input.Meta[n].Value,
		})
	}

	a, _ := json.Marshal(tmp)
	output = string(a)

	return
}
