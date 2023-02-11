package app

import (
	"crypto/rand"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"os"
)

const tokenSize = 15

func wget(link string) []byte {
	u, err := url.Parse(link)
	if err != nil {
		return []byte{}
	} else if u.Hostname() != "127.0.0.1" && u.Hostname() != "localhost" {
		return []byte{}
	}

	response, err := http.Get(u.String())
	if err != nil {
		return []byte{}
	} else if response.StatusCode != http.StatusOK {
		return []byte{}
	}

	if body, err := ioutil.ReadAll(response.Body); err != nil {
		return []byte{}
	} else {
		return body
	}
}

func save(body []byte) (string, error) {
	tok, err := genToken(tokenSize)
	if err != nil {
		return "", err
	}

	filename := "./uploads/" + tok + ".bin"
	if err = os.WriteFile(filename, body, 0644); err != nil {
		return "", err
	}
	return tok, nil
}

func load(token string) []byte {
	if body, err := os.ReadFile("./uploads/" + token + ".bin"); err != nil {
		return []byte{}
	} else {
		return body
	}
}

func genToken(n int) (string, error) {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	size := big.NewInt(int64(len(alphabet)))
	tok := make([]rune, n)

	for i := 0; i < n; i++ {
		j, err := rand.Int(rand.Reader, size)
		if err != nil {
			return "", err
		}
		tok[i] = alphabet[j.Int64()]
	}
	return string(tok), nil
}
