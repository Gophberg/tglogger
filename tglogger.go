package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

var chUser1 = make(chan bool)
var chUser2 = make(chan bool)
var chUser3 = make(chan bool)
var chUser4 = make(chan bool)
var chUser5 = make(chan bool)

type Position struct {
	X, Y int
}

func main() {
	go getState()

	for {
		select {
		case User1 := <-chUser1:
			{
				if User1 == true {
					t := sTime()
					str := []string{fmt.Sprintf("%v", t), "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------"}
					rec(true, str)
				}
				if User1 == false {
					t := sTime()
					str := []string{"-------------", fmt.Sprintf("%v", t), "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------"}
					rec(false, str)
				}
			}
		case User2 := <-chUser2:
			{
				if User2 == true {
					t := sTime()
					str := []string{"-------------", "-------------", fmt.Sprintf("%v", t), "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------"}
					rec(true, str)
				}
				if User2 == false {
					t := sTime()
					str := []string{"-------------", "-------------", "-------------", fmt.Sprintf("%v", t), "-------------", "-------------", "-------------", "-------------", "-------------", "-------------"}
					rec(false, str)
				}
			}
		case User3 := <-chUser3:
			{
				if User3 == true {
					t := sTime()
					str := []string{"-------------", "-------------", "-------------", "-------------", fmt.Sprintf("%v", t), "-------------", "-------------", "-------------", "-------------", "-------------"}
					rec(true, str)
				}
				if User3 == false {
					t := sTime()
					str := []string{"-------------", "-------------", "-------------", "-------------", "-------------", fmt.Sprintf("%v", t), "-------------", "-------------", "-------------", "-------------"}
					rec(false, str)
				}
			}
		case User4 := <-chUser4:
			{
				if User4 == true {
					t := sTime()
					str := []string{"-------------", "-------------", "-------------", "-------------", "-------------", "-------------", fmt.Sprintf("%v", t), "-------------", "-------------", "-------------"}
					rec(true, str)
				}
				if User4 == false {
					t := sTime()
					str := []string{"-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", fmt.Sprintf("%v", t), "-------------", "-------------"}
					rec(false, str)
				}
			}
		case User5 := <-chUser5:
			{
				if User5 == true {
					t := sTime()
					str := []string{"-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", fmt.Sprintf("%v", t), "-------------"}
					rec(true, str)
				}
				if User5 == false {
					t := sTime()
					str := []string{"-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", "-------------", fmt.Sprintf("%v", t)}
					rec(false, str)
				}
			}
		}
	}
}

func getState() {

	var user1 Position
	user1.X = 53
	user1.Y = 120

	var user2 Position
	user2.X = 53
	user2.Y = 182

	var user3 Position
	user3.X = 53
	user3.Y = 243

	var user4 Position
	user4.X = 53
	user4.Y = 307

	var user5 Position
	user5.X = 53
	user5.Y = 369

	stateUser1 := captureState(user1)
	stateUser2 := captureState(user2)
	stateUser3 := captureState(user3)
	stateUser4 := captureState(user4)
	stateUser5 := captureState(user5)

	newStateUser1, newStateUser2, newStateUser3, newStateUser4, newStateUser5 := false, false, false, false, false
	var stUser1, stUser2, stUser3, stUser4, stUser5 *bool
	stUser1 = &newStateUser1
	stUser2 = &newStateUser2
	stUser3 = &newStateUser3
	stUser4 = &newStateUser4
	stUser5 = &newStateUser5
	for {
		*stUser1 = captureState(user1)
		if *stUser1 != stateUser1 {
			chUser1 <- *stUser1
			stateUser1 = *stUser1
		}
		*stUser2 = captureState(user2)
		if *stUser2 != stateUser2 {
			chUser2 <- *stUser2
			stateUser2 = *stUser2
		}
		*stUser3 = captureState(user3)
		if *stUser3 != stateUser3 {
			chUser3 <- *stUser3
			stateUser3 = *stUser3
		}
		*stUser4 = captureState(user4)
		if *stUser4 != stateUser4 {
			chUser4 <- *stUser4
			stateUser4 = *stUser4
		}
		*stUser5 = captureState(user5)
		if *stUser5 != stateUser5 {
			chUser5 <- *stUser5
			stateUser5 = *stUser5
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func captureState(p Position) bool {
	// var offLine = color.RGBA{40, 46, 51, 255}
	var onLineActived = color.RGBA{255, 255, 255, 255}
	var onLineInactived = color.RGBA{10, 231, 210, 255}

	img, err := screenshot.Capture(p.X, p.Y, 1, 1)
	if err != nil {
		panic(err)
	}
	readImg := img.At(0, 0)
	img = nil
	if readImg == onLineActived || readImg == onLineInactived {
		return true
	}
	return false
}

func rec(status bool, s []string) {
	t := time.Now().Format("02012006")
	fileName := fmt.Sprintf("%v.csv", t)
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	if err := w.Write(s); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}
	defer f.Close()
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func sTime() string {
	return time.Now().Format("0201 15:04:05")
}
