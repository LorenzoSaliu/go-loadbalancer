package main

import (
	"fmt"
	
	"loadbalancer/models"
)

type Loadbalancer struct {
	port uint8
	servers []*models.Server
	roundRobinCount int
}

func newLoadBalancer(){}
func (lb *Loadbalancer) addServer(s *models.Server){
	lb.servers = append(lb.servers, s)	
}

func main(){
	fmt.Println("start")

}
