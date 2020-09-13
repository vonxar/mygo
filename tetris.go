package main

//include <conio.

import "C"

import (
	"fmt"
	"time"
)

func mainLoop() {
	//ループミリ秒
	var loopms time.Duration = 33
	//タイマー
	timer := time.NewTimer(loopms * time.Millisecond)

	//メインループ
	for {
		select {
		case <-timer.C:
			timer = time.NewTimer(loopms * time.Millisecond)
			str := ""
			if C.kbhit() != 0 {
				dtr = dtring(C.getche())
				fmt.Println(str)
			} else {
				fmt.Println(".")
			}
			break
		}
	}
}
