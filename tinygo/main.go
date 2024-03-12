// this is just a prove of concept to try to use tinygo instead of the
// Arduino-IDE

// The Board is connected to two buttons, two LEDs via two relays and a buzzer.
// If one button is pressed, the corresponding LED is switched on and some sound
// will be played.

package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
)

// A note is a combination of tone and duration
type note struct {
	tone      float64
	durartion float64
}

// A sequence is a slice of notes
type sequence []note

// Define pins, relays-states and time to wait
const (
	// For the buttons pin 8 and 9
	buttonGreen = machine.D8
	buttonRed   = machine.D9
	// For the speaker pin 10
	speakerPin = machine.D10
	// For the relays pin 2 and 3
	relaysGreen = machine.D2
	relaysRed   = machine.D3
	// we need to switch true and false, because the relays use GND for ON and
	// anything else for OFF
	rON  = false
	rOFF = true
	// time to wait before a new button-input is accepted
	wait = 5
)

func main() {
	/*
		SETUP
	*/
	// initialize buttons
	buttonGreen.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonRed.Configure(machine.PinConfig{Mode: machine.PinInput})

	// initialize relays
	relaysGreen.Configure(machine.PinConfig{Mode: machine.PinOutput})
	relaysRed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	relaysGreen.Set(rOFF)
	relaysRed.Set(rOFF)

	// initialize Speaker
	speakerPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker := buzzer.New(speakerPin)

	// set sound-sequences
	goodSequence := sequence{
		{buzzer.F4, buzzer.Eighth / 2},
		{buzzer.B4, buzzer.Eighth / 2},
		{buzzer.A5, buzzer.Eighth},
	}
	badSequnce := sequence{
		{buzzer.E4, buzzer.Eighth / 2},
		{buzzer.E4, buzzer.Eighth / 2},
		{buzzer.E4, buzzer.Eighth},
	}

	/*
		MAIN LOOP
	*/
	for {
		// save button-states to variables because they are very volatile
		buttonGreenPressed := buttonGreen.Get()
		buttonRedPressed := buttonRed.Get()

		// There is no need to make a race between two buttons. If they are both
		// pressed, just do nothing.
		if buttonGreenPressed != buttonRedPressed { // Sadly I bought one opener and one closer. So I had to replace "==" with "!="
			// Wait for a fracture of a second for new input.
			time.Sleep(time.Millisecond * 100)
			continue
		}

		// If one button is pressed, start the show.
		if buttonRedPressed {
			show(speaker, relaysRed, badSequnce)
		} else {
			show(speaker, relaysGreen, goodSequence)
		}

		// wait before a new input is accepted
		time.Sleep(time.Second * wait)
	}
}

// show switches on the corresponding relay and plays the sounds.
func show(speaker buzzer.Device, relays machine.Pin, seq sequence) {
	// light ON
	relays.Set(rON)
	// play sounds
	for _, s := range seq {
		speaker.Tone(s.tone, s.durartion)
	}
	// light OFF
	relays.Set(rOFF)
}
