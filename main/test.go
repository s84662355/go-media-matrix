package main
import (
 
   "fmt"  
    "time"
)

type A struct{

}

func (l *A) aaa(a int , b int){
	fmt.Println(a,b)
}

type Cb func(a int , b int)

func main() {

	  
   // var eventMap = make(map[string]map[string]Cb)

   aaaa := new(A)

  var mmm =  make(map[string]Cb)


  mmm["b"] = Cb(func(a int , b int){
    		fmt.Println(a,b)
    })
 mmm["c"] = Cb(aaaa.aaa) 
     mmm["b"](312,34543)

  //  eventMap["a"] =  mmm["b"]

mmm["c"](23423,4656)
///eventMap["a"]["b"](123,154645)
time.Sleep(10* time.Second)

  
}