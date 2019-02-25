package jedilevel8

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSONData ... unmarshalling data from json
func UnmarshalJSONData() {
	jsonout := []byte(`[{"First":"Peter","Last":"Parker","Age":32},{"First":"Sharon","Last":"Tuchey","Age":30}]`)

	var users []user

	err := json.Unmarshal(jsonout, &users)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", users)
}
