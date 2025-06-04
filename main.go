//RLE (RUN-LENGTH ENCODING)

package main

import (
	"fmt"
	"os"
	"strings"
)

const helpMessage = "Usage: [rle-coder decode filename] or [rle-coder decode filename]"

func encode(file []byte) (encodedData []byte) {
	var sameCount byte = 1

	for i := 0; i < len(file)-1; i++ {
		if file[i] == file[i+1] {
			sameCount++
		} else {
			encodedData = append(encodedData, sameCount)
			encodedData = append(encodedData, file[i])
			sameCount = 1
		}
	}
	encodedData = append(encodedData, sameCount)
	encodedData = append(encodedData, file[len(file)-1])
	return
}

func decode(file []byte) (decodedData []byte) {
	var recordData byte
	var count int = 0
	for i := 0; i < len(file); i++ {
		if i%2 == 0 {
			count = int(file[i])
		} else {
			recordData = file[i]
			for k := 0; k < count; k++ {
				decodedData = append(decodedData, recordData)
			}
		}
	}
	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(helpMessage)
		return
	}

	file, err := os.ReadFile(os.Args[2])

	if err != nil {
		fmt.Println("File open error: " + os.Args[2])
		return
	}

	switch os.Args[1] {
	case "encode":
		os.WriteFile("encoded_"+os.Args[2], encode(file), os.ModePerm)
		fmt.Println("Encoded")
	case "decode":
		fileName := strings.Replace(os.Args[2], "encoded_", "decoded_", 1)
		os.WriteFile(fileName, decode(file), os.ModePerm)
		fmt.Println("Decoded")
	default:
		fmt.Println(helpMessage)
	}
}
