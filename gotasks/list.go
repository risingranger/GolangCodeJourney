[12:25 PM] Dr. Tarkeshwar Barua

package main

 

import (

    "container/list"

    "fmt"

)

 

func main() {

    var obj list.List

    obj.PushBack(2)

    obj.PushBack(6)

    obj.PushBack(3)

    for index := obj.Front(); index != nil; index = index.Next() {

        fmt.Println(index.Value.(int))

    }

}

 