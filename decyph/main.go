package main

import (
	"crypto/rsa"
	"io/ioutil"
	"math/big"
)

func main() {
	// pubKey, _ := ioutil.ReadFile("pubKey.blob")
	file, _ := ioutil.ReadFile("11月信公.xls.prs")
	// fmt.Printf("%x", string(pubKey))
	nStr := "91b442a841a0278488e937c59c21070e6cd9633c594ec9ce9b816bcbde6d1394308c38a3ad277fd94eeac88b8baadb8cfecd9b6cb6bc0360e32922769408acdafbc1f073c54b54943ad63ce9fa8489bb162a1eefd6bc43e015b1fd556ae403855a459b6454fb375da2f85f1ba7f539fec22ec629241b0b88742b55f661d9b8ab"
	n := big.NewInt(0)
	n.SetString(nStr, 16)
	pubKey := rsa.PublicKey{N: n, E: 65537}
}
