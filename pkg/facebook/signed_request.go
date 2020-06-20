package facebook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

func ParseSignedRequest(signedRequest string, secret string) (map[string]interface{}, error) {
	encoded := strings.SplitN(signedRequest, ".", 2)

	sig := encoded[0]
	jsonReader := base64.NewDecoder(base64.URLEncoding, strings.NewReader(encoded[1]))
	rawJson, err := ioutil.ReadAll(jsonReader)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(rawJson, data)
	if err != nil {
		return nil, err
	}

	algorithm, ok := data["algorithm"]
	if !ok {
		return nil, errors.New("JSON data is missing algorithm")
	}
	if algorithm != "HMAC-SHA256" {
		return nil, errors.New("algorithm is unexpected")
	}

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(encoded[1]))
	expectedSig := base64.StdEncoding.EncodeToString(h.Sum(nil))
	expectedSig = strings.ReplaceAll(expectedSig, "+", "-")
	expectedSig = strings.ReplaceAll(expectedSig, "/", "_")
	expectedSig = strings.ReplaceAll(expectedSig, "=", "")

	if sig != expectedSig {
		return nil, errors.New("signatures did not match")
	}

	return data, nil
}
