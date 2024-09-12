package main

import "../singleton"

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			singleton.GetInstanceWithAtomic()
		}
	}()
	for i := 0; i < 100; i++ {
		singleton.GetInstanceWithAtomic()
	}
}
