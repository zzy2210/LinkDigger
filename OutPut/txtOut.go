package OutPut

import (
	"bufio"
	"log"
	"os"
)

func TxtOut(path string,conSlice []string){
	txtPath := path + ".txt"

	file,err := os.OpenFile(txtPath,os.O_CREATE|os.O_RDWR,0640)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _,v := range conSlice {
		writer.WriteString(v+"\n")
		writer.Flush()
	}


}