package main
import "fmt"
import	"strconv"


func main(){

	var cha byte = 'a'
	fmt.Scanf("%c",&cha)
	switch cha {
	case 'a':
		fmt.Println("A")
	case 'z':
		fmt.Println("Z")
	default :
		fmt.Println("other")
	}



	var n2 byte = 'a'
	var n3 byte = '0'

	fmt.Println("n2,n3",n2,n3)
	fmt.Printf("%c,%c",n2,n3)
	
	var n4 int = '地'
	fmt.Printf("\n%c,%d",n4,n4)

	var int1 int = 1234
	var str3 = fmt.Sprintf("%v",int1)
	str3 = fmt.Sprintf("%d",int1)
	fmt.Printf("\n%T,%v",str3, str3)

	var str5 = strconv.Itoa(int1)
	fmt.Printf("\n%T,%v",str5, str5)
	str3 = fmt.Sprintf("%d",int1)
	fmt.Printf("\n%T,%v",str3, str3)

	fmt.Printf("\n%v",fmt.Sprintf("%f",1.00))
	fmt.Printf("\n%q",fmt.Sprintf("%t",true))
	fmt.Printf("\n%q",fmt.Sprintf("%c",'a'))

	//---------------------------
	
	var str string = "true"
	var b bool 
	b,_ = strconv.ParseBool(str)
	fmt.Printf("\n%T,%v",b, b)
   
	var str2 string = "234"
	var intstr int64
	intstr ,_ = strconv.ParseInt(str2,10,64)//将str2 转成10进制 64 位表示
	fmt.Printf("\n%T,%v",intstr, intstr)

}

