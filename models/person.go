package models

import (
	"../utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

type Person struct {
	gorm.Model
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type PersonAggr struct {
	Cnt int32
}

func GetFirstPerson() *Person {
	temp := &Person{};
	err := GetDB().Debug().Table("person").First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		println("db eror")
	}

	return temp
}

func Filter(vs []Person, f func(Person) bool) []Person {
	vsf := make([]Person, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func SumAgeBefore(age int32) PersonAggr {

	var people []Person

	start:=time.Now()
	err := GetDB().Debug().Table("person").Find(&people).Error

	elapsed:=time.Since(start)

	fmt.Printf("query time %v\n",elapsed)

	if err != nil && err != gorm.ErrRecordNotFound {
		println("db eror")
	}

	return calcAge()(people, age)
}

func calcAge() func([]Person, int32) PersonAggr {
	return func(people []Person, age int32) PersonAggr {
		cont := utils.Container{}
		for _, value := range people {
			cont.Put(value)
		}
		filteredPeople := cont.Filer(func(i interface{}) bool {
			tmp, _ := i.(Person)
			if (tmp.Age < age) {
				return true
			}
			return false
		})
		var sumAge int32
		for _, element := range filteredPeople {
			tmp, _ := element.(Person)
			sumAge += tmp.Age
		}
		return PersonAggr{Cnt: sumAge}
	}
}

func putIntoChan(wg *sync.WaitGroup, ch chan PersonAggr, item func(people []Person, age int32) PersonAggr, p []Person, age int32, i int) {
	defer wg.Done()
	fmt.Println("process ", i)
	ch <- item(p, age)
}

func ManyPersonQueries() [] PersonAggr {
	var aggrChannel=make(chan PersonAggr,10)

	var wg sync.WaitGroup

	wg.Add(10)

	for i:=0; i<10;i++ {
		go func() {
			defer wg.Done()
			aggrChannel <- SumAgeBefore(50)
		}()
	}

	wg.Wait()
	close(aggrChannel)

	var aggrArray=make([] PersonAggr,10)
	for result:= range aggrChannel {
		aggrArray=append(aggrArray,result )
	}
	return aggrArray
}

func SumAgeBeforeParallel(age int32) PersonAggr {
	var people []Person

	start := time.Now()

	err := GetDB()/*.Debug()*/.Table("person").Find(&people).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		println("db eror")
	}
	elapsed := time.Since(start)
	fmt.Println("queryTime %s", elapsed)

	peopleSize := len(people)
	fmt.Println("people size", len(people))
	batchSize := peopleSize / 10.0


	capacity:=10

	var aggrs = make([]PersonAggr, capacity)

	personChan := make(chan PersonAggr, capacity)

	var wg sync.WaitGroup
	wg.Add(capacity)
	for i := 0; i < capacity; i++ {
		sl := people[i*capacity : batchSize+i*capacity]
		//fmt.Println("i %s batchSize %s", i, batchSize)
		go putIntoChan(&wg, personChan, calcAge(), sl, 50, i)
	}
	wg.Wait()

	close(personChan)

	i := 0
	for resp := range personChan {
		//fmt.Println("aggr finished", i)
		aggrs[i] = resp
		i++
	}

	//for i := 0; i < 10; i++ {
	//	fmt.Println("aggr finished",i)
	//	aggrs[i] = <-personChan
	//}

	var sumAge int32

	for _, value := range aggrs {
		sumAge += value.Cnt
	}

	return PersonAggr{Cnt: sumAge}

}
