package state

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type FState int

const (
	OffHook FState = iota
	Connecting
	Connected
	OnHold
	OnHook
)

func (s FState) String() string {
	switch s {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	default:
		return "Unknown"
	}
}

type Trigger int

const (
	CallDialed Trigger = iota
	HungUp
	CallConnected
	PlacedOnHold
	TakenOffHold
	LeftMessage
)

func (t Trigger) String() string {
	switch t {
	case CallDialed:
		return "CallDialed"
	case HungUp:
		return "HungUp"
	case CallConnected:
		return "CallConnected"
	case PlacedOnHold:
		return "PlacedOnHold"
	case TakenOffHold:
		return "TakenOffHold"
	case LeftMessage:
		return "LeftMessage"
	default:
		return "Unknown"
	}
}

type TriggerResult struct {
	Trigger Trigger
	State   FState
}

var Rules = map[FState][]TriggerResult{
	OffHook: {
		{CallDialed, Connecting},
	},
	Connecting: {
		{HungUp, OnHook},
		{CallConnected, Connected},
	},
	Connected: {
		{LeftMessage, OnHook},
		{HungUp, OnHook},
		{PlacedOnHold, OnHold},
	},
	OnHold: {
		{TakenOffHold, Connected},
		{HungUp, OnHook},
	},
}

func ProcessPhoneCall() {
	callState, exitState := OffHook, OnHook
	for ok := true; ok; ok = callState != exitState {
		fmt.Println("The phone is currently", callState)
		fmt.Println("Select a trigger:")

		for i := 0; i < len(Rules[callState]); i++ {
			tr := Rules[callState][i]
			fmt.Println(strconv.Itoa(i), ".", tr.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		tr := Rules[callState][i]
		callState = tr.State
	}
	fmt.Println("We are done using the phone")
}