package critego

import (
	"encoding/json"
	"log"
)

func Jsonify(i interface{}) []byte {
	bs, err := json.Marshal(i)
	if err != nil {
		log.Fatalf("Error when JSONify: %v", err)
	}
	return bs
}
