// this is just a prove of concept to try to use tinygo instead of the
// ardiuno-IDE

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

func main() {
	/*
	 SETUP
	*/
	// Define pins for the buttons
	buttonGreen := machine.D8
	buttonRed := machine.D9
	// initialize buttons
	buttonGreen.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonRed.Configure(machine.PinConfig{Mode: machine.PinInput})

	// Define pins for the relays
	relaysGreen := machine.D2
	relaysRed := machine.D3
	// we need to switch true and false, because the relays use GND for ON and
	// anything else for OFF
	rON := false
	rOFF := true
	// initialize relays
	relaysGreen.Configure(machine.PinConfig{Mode: machine.PinOutput})
	relaysRed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	relaysGreen.Set(rOFF)
	relaysRed.Set(rOFF)

	// Define pin for the speaker
	speakerPin := machine.D10
	// initialize Speaker
	speakerPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker := buzzer.New(speakerPin)

	// Define sound-sequences
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

	// show switches on the corresponding relay and plays the sounds.
	show := func(pin machine.Pin, seq sequence) {
		// light ON
		pin.Set(rON)
		// play sounds
		for _, s := range seq {
			speaker.Tone(s.tone, s.durartion)
		}
		// light OFF
		pin.Set(rOFF)
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
			show(relaysRed, badSequnce)
		} else {
			show(relaysGreen, goodSequence)
		}
	}
}
