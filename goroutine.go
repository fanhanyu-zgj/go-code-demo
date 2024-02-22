package main

import (
	"errors"
	"fmt"
	"sync"
)

func test1() (string, error) {
	return "test1Res", nil
}

func test2() (string, error) {
	return "test2Res", errors.New("occured:message")
}

func main() {
	var wg sync.WaitGroup
	var errCh = make(chan error, 2)
	wg.Add(1)
	var test1Res string
	go func() {
		defer wg.Done()
		errCh <- func() error {
			res, err := test1()
			if err != nil {
				return err
			}
			test1Res = res
			return nil
		}()
	}()

	wg.Add(1)
	var test2Res string
	go func() {
		defer wg.Done()
		errCh <- func() error {
			res, err := test2()
			if err != nil {
				return err
			}
			test2Res = res
			return nil
		}()
	}()

	wg.Wait()
	for len(errCh) > 0 {
		err := <-errCh
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(test1Res, test2Res)
}
