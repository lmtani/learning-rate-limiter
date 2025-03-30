package entity

import "encoding/json"

type TokenMap map[string]int

// Decode permite que envconfig decodifique a variável de ambiente (string JSON)
// para o tipo TokenMap.
func (tm *TokenMap) Decode(value string) error {
	return json.Unmarshal([]byte(value), tm)
}
