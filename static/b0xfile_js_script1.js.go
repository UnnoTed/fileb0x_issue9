// generaTed by fileb0x at "2017-03-20 17:23:19.387923889 -0300 BRT" from config file "b0x.yaml"

package static

import (
	"log"
	"os"
)

// FileJsScript1Js is "js/script1.js"
var FileJsScript1Js = []byte("\x63\x6f\x6e\x73\x6f\x6c\x65\x2e\x6c\x6f\x67\x28\x27\x6f\x6b\x27\x29\x3b\x0a")

func init() {

	f, err := FS.OpenFile(CTX, "js/script1.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(FileJsScript1Js)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
