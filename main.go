package main

 import(
	"eirc/structure"
	"fmt"
	"encoding/json"
 )


func main(){
	student1 := structure.Students{
		Name:   "amy",
		ID: 1,
	}
	student2 := structure.Students{}
		student2.Name ="Bin"
		student2.ID= 2
	
 class:=[]structure.Students{student1,student2}
 	fmt.Println(class)

 class2:=[]structure.Students{}

		for index, value := range class{
		class2=append(class2,value)
		fmt.Println(index,class2)
}
marshal,err := json.Marshal(class)
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(marshal))
	
class3:=[]structure.Students{}
err=json.Unmarshal(marshal,&class3)
if err != nil {
	fmt.Println(err)
}
fmt.Println(class3)

a := 2
fmt.Println("a =", a)

//&獲取變數的記憶體位置
b := &a
fmt.Println("b is address of a = ", b)

//*獲取指標變數中的值
c := *b
fmt.Printf("c is value of a = %d\n", c)
d := *&a
fmt.Printf("d is the same as c = %d", d)

}

	
