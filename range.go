package main

import "sync"

func Ints(s []int) {
	var c int
	for _, v := range s {
		c = v
	}
	_ = c
}

func IntsRef(s []int) {
	var c *int
	for _, v := range s {
		c = &v
	}
	_ = c
}

func IntsClosures(s []int) {
	var c int
	var wg sync.WaitGroup
	wg.Add(len(s))
	for _, v := range s {
		go func() {
			c = v
			wg.Done()
		}()
	}
	wg.Wait()
	_ = c
}

func Strings(s []string) {
	var c string
	for _, v := range s {
		c = v
	}
	_ = c
}

func StringsRef(s []string) {
	var c *string
	for _, v := range s {
		c = &v
	}
	_ = c
}

func StringsClosures(s []string) {
	var c string
	var wg sync.WaitGroup
	wg.Add(len(s))
	for _, v := range s {
		go func() {
			c = v
			wg.Done()
		}()
	}
	wg.Wait()
	_ = c
}

func Structs(s []LargeStruct) {
	var c LargeStruct
	for _, v := range s {
		c = v
	}
	_ = c
}

func StructsRef(s []LargeStruct) {
	var c *LargeStruct
	for _, v := range s {
		c = &v
	}
	_ = c
}

func StructsClosures(s []LargeStruct) {
	var c LargeStruct
	var wg sync.WaitGroup
	wg.Add(len(s))
	for _, v := range s {
		go func() {
			c = v
			wg.Done()
		}()
	}
	wg.Wait()
	_ = c
}
