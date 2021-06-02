package main

import (
	"fmt"
	"math/rand"
	"time"
)

func AbsI(x int) int{
	if x < 0 {
		return -x
	}
	return x
}


func main() {
	var gene[9]int
	var add = 0
	var goal = 0
	var mut = 0
	for j:= 0; j < 9; j++{
		rand.Seed(time.Now().UTC().UnixNano())
		gene[j] = rand.Intn(9)
	}
	fmt.Println(gene)
	for {
		add = 0
		mut = 0
		for x:= 0; x < 9; x++{
			add = add + gene[x]
		}
		if add == goal {
			fmt.Println("correct sequence found")
			break
		} else {
			time.Sleep(3 * time.Second)
			rand.Seed(time.Now().UTC().UnixNano())
			var mutatedgene[9] int = gene
			mutatedgene[rand.Intn(9)] = rand.Intn(9)
			mut = 0
			for y:= 0; y < 9; y++{
				mut = mut + mutatedgene[y]
			}
			if AbsI(mut - goal)<AbsI(add - goal){
				fmt.Println(" V")
				fmt.Println(mutatedgene, mut," success")
				gene = mutatedgene
			}else{
				fmt.Println(" |  ",mutatedgene, " failed")
			}
		}
	}
}

