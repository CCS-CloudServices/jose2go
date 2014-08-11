package padding

import (
	"bytes"
)

func AddPkcs7(data []byte, blockSize int) []byte {
	
	var paddingCount int
	
	if paddingCount = blockSize - (len(data) % blockSize);paddingCount == 0 {
		paddingCount=blockSize
	}		
	
	return append(data, bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)...)
}

func RemovePkcs7(padded []byte, blockSize int) []byte {	
	
	dataLen:=len(padded)		
	paddingCount:=int(padded[dataLen-1])
	
	if(paddingCount > blockSize || paddingCount <= 0) {
		return padded //data is not padded (or not padded correctly), return as is
	}

	padding := padded[dataLen-paddingCount : dataLen-1]
			
	for _, b := range padding {
		if int(b) != paddingCount {
			return padded  //data is not padded (or not padded correcly), return as is
		}
	}		
		
	return padded[:len(padded)-paddingCount] //return data - padding
}