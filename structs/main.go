package structs

import (
	"github.com/paidTrainee27/gocodingpractice/model"
	"fmt"
	"reflect"
	"sync"
	"time"
)

/*How to access private fields of public struct*/
func PrivateField() {
	l := new(model.List)
	l.AddFront(4)

	// Get a reflect.Value fv for the unexported field len.
	fv := reflect.ValueOf(l).Elem().FieldByName("len")
	fmt.Println(fv.Int()) // 4
}

/*Advantage of empty struct*/
func EmptyStruct() {
	finish := make(chan struct{})
	t0 := time.Now()
	go func() {
		time.Sleep(time.Second * 2)
		finish <- struct{}{} //OR
		// close(finish)
	}()
	<-finish
	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))

}

func EmptyStruct1() {
	finish := make(chan struct{})
	var done sync.WaitGroup
	done.Add(1)
	go func() {
		//OR
		select {
		case <-time.After(1 * time.Hour):
		case <-finish: //called when recieved a value or channel is closed
		}
		done.Done()
	}()
	t0 := time.Now()
	close(finish)
	done.Wait()
	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))
}
