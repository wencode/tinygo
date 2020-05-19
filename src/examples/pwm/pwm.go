package main

// This example demonstrates some features of the PWM support.

import (
	"machine"
	"time"
)

var (
	// Configuration on an Adafruit Circuit Playground Express.
	pwm  = machine.PWM1
	pinA = machine.A2
	pinB = machine.A3

	// Configuration on an Arduino Uno.
	//pwm  = machine.Timer2
	//pinA = machine.PB3 // pin 11 on the Uno
	//pinB = machine.PD3 // pin 3 on the Uno
)

const delayBetweenPeriods = time.Second * 5

func main() {
	// Delay a bit on startup to easily catch the first messages.
	time.Sleep(time.Second * 2)

	// Configure the PWM with the given period.
	err := pwm.Configure(machine.PWMConfig{
		Period: 16384e3, // 16.384ms
	})
	if err != nil {
		println("failed to configure PWM")
		return
	}

	// The top value is the highest value that can be passed to PWMChannel.Set.
	// It is usually an even number.
	println("top:", pwm.Top())

	// Configure the two channels we'll use as outputs.
	channelA, err := pwm.Channel(pinA)
	if err != nil {
		println("failed to configure channel A")
		return
	}
	channelB, err := pwm.Channel(pinB)
	if err != nil {
		println("failed to configure channel B")
		return
	}

	// Invert one of the channels to demonstrate output polarity.
	channelB.SetInverting(true)

	// Test out various frequencies below, including some edge cases.

	println("running at 0% duty cycle")
	channelA.Set(0)
	channelB.Set(0)
	time.Sleep(delayBetweenPeriods)

	println("running at 1")
	channelA.Set(1)
	channelB.Set(1)
	time.Sleep(delayBetweenPeriods)

	println("running at 25% duty cycle")
	channelA.Set(pwm.Top() / 4)
	channelB.Set(pwm.Top() / 4)
	time.Sleep(delayBetweenPeriods)

	println("running at top-1")
	channelA.Set(pwm.Top() - 1)
	channelB.Set(pwm.Top() - 1)
	time.Sleep(delayBetweenPeriods)

	println("running at 100% duty cycle")
	channelA.Set(pwm.Top())
	channelB.Set(pwm.Top())
	time.Sleep(delayBetweenPeriods)

	for {
		time.Sleep(time.Second)
	}
}
