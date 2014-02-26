package main

//import "os"
import "fmt"

var (
	ports []int
	domains []int
	ports_status map[int]bool
	all_domaim []Domain
)

type Domain struct {
	domain int
	ports_status map[int]bool
}

func NewDomain(d int,ps map[int]bool) *Domain{
	return &Domain{
		domain : d,
		ports_status  : ps,
	}
}

func init(){
	ports = []int{80,8000,8001,8002,8003,8004}
	domains = []int{4,5,6}
	ports_status = make(map[int]bool)
	for _,port := range ports{
		ports_status[port] = false
	}
	for _ , domain_num := range domains{
		domain := *NewDomain(domain_num,ports_status)
		all_domaim = append(all_domaim,domain)
	}
}

func setTrue(d,p int){
	for _, domain := range all_domaim{
		if domain.domain == d {
			domain.ports_status[p] = true
			return
		}
	}
}


func main() {
	fmt.Printf("%v",all_domaim)
	return
}

func test_fun() {
	fmt.Printf("i am test func \n")
	return
}


