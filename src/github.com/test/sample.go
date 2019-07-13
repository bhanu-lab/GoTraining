package main
import (
    "fmt"
    "os"
    "bufio"
    )

func main(){
	// your code goes here
	reader := bufio.NewReader(os.Stdin)
	val,_ := reader.ReadString('\n')
	
	for val != "42" {
	    
	    val,_ = reader.ReadString('\n')
	    
	    if val == "42"{
	        break
	    }else{
	        fmt.Println(val)
	    }
	    
	}
}
