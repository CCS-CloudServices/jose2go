package jose

import (
	"errors"
)

func init() {
	RegisterJwa(new(Direct))
}

type Direct struct{
}

func (alg *Direct) Name() string {
	return "dir"
}

func (alg *Direct) WrapNewKey(cekSizeBits int, key interface{}, header map[string]interface{}) (cek []byte, encryptedCek []byte, err error) {
	
	if cek,ok:=key.([]byte); ok {
		return cek,[]byte{},nil
	}
	
	return nil,nil,errors.New("Direct.WrapNewKey(): expected key to be '[]byte' array")
}

func (alg *Direct) Unwrap(encryptedCek []byte, key interface{}, cekSizeBits int, header map[string]interface{}) (cek []byte, err error) {

	if(len(encryptedCek)!=0) {
		return nil, errors.New("Direct key management expects empty encrypted CEK")
	}
	
	if cek,ok:=key.([]byte); ok {
		return cek,nil
	}
		
	return nil,errors.New("Direct.Unwrap(): expected key to be '[]byte' array")	
}
