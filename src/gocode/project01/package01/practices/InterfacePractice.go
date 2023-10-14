package practices
import "fmt"
import "sort"
import "math/rand"


type Hero struct{
	Name string
	Age int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int{
	return len(hs)
}

func (hs HeroSlice) Less(i,j int) bool{
	// return hs[i].Age < hs[j].Age
	return hs[i].Name < hs[j].Name
}

func (hs HeroSlice)Swap(i,j int){
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	hs[i],hs[j] = hs[j],hs[i]
} 


func Test(){

	var heros HeroSlice
	for i :=0; i < 10; i++ {
		hero := Hero{
			Name : fmt.Sprintf("cat%d",rand.Intn(100)),
			Age : rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	fmt.Println("--before sort-",heros)

	sort.Sort(heros)

	fmt.Println("-after sort--",heros)

	for _, v := range heros{
		fmt.Println(v)
	}
	var new int = 0
	fmt.Println(new)


}
