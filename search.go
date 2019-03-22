package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var PWNEDAPIURL = "https://api.pwnedpasswords.com/range/"

func main() {
	pwdHash := hashSha1(os.Args[1])
	pwdHashHex := convertBytesToHex(pwdHash)
	resp := pwnedLookupApi(pwdHashHex)
	if resp != "" {
		searchRespForHashHex(resp, pwdHashHex)
	}
}

func hashSha1(str string) []byte {
	hasher2 := sha1.New()
	io.WriteString(hasher2, str)
	return hasher2.Sum(nil)
}

func convertBytesToHex(src []byte) string {
	encodedStr := hex.EncodeToString(src)
	return fmt.Sprintf("%s", encodedStr)
}

func pwnedLookupApi(pwdHash string) string {
	first5Chars := string(pwdHash[:5])
	resp, err := http.Get(fmt.Sprint(PWNEDAPIURL, first5Chars))

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		return fmt.Sprintf("%s", body)
	}

	log.Fatal("The request responded with something other than a 200")
	return ""
}

func searchRespForHashHex(resp string, pwdHashHex string) bool {
	minus5Chars := string(pwdHashHex[5:])
	for _, s := range strings.Split(resp, "\r\n") {
		apiHashHex := strings.Split(s, ":")
		if apiHashHex[0] == strings.ToUpper(minus5Chars) {
			fmt.Println(fmt.Sprintf("Your password has been leaked %s times", apiHashHex[1]))

			return true
		}
	}

	fmt.Println("Your password was not found in haveibeenpwned's database")
	return false
}
