package main

import (
	"fmt"
	"log"
	"os"
)


func main () {

	var xor int = 1;
	var fpath string
	var fpathOutput string


	
	for xor > 0 && xor < 27 {

		fmt.Print("To quit enter 0\nenter the xor shift [+]:")
		fmt.Scan(&xor)
		fmt.Println(xor)
		if xor == 0 {
			break
		}
		fmt.Print("Enter the path for the file you want to cipher\n[+]:")
		fmt.Scan(&fpath)
		fmt.Print("Enter the output path for your encrypted file\n[+]:")
		fmt.Scan(&fpathOutput)
		
		
		// Reading the input file
		fpathHistory, err := os.ReadFile(fpath)
		if err != nil {
			log.Fatal(err)
		}
		
		var charA int

		charA = charA >> charA
		fmt.Println(string(charA))

		result := ""
		for _, char := range fpathHistory {
			if char >= 'a' && char <= 'z' {
				result +=string(((int(char)-'a'+xor)%26)+'a')
			} else if char >= 'A' && char <= 'Z' {
				result +=string(((int(char)-'A'+xor)%26)+'A')
			} else {
				//Anything other than lowercase and uppercase
				result+=string(char)
			}	
			

		}

		
		fpathOutput, err := os.Create(fpathOutput)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer fpathOutput.Close()

		_, err = fpathOutput.WriteString(string(result))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}