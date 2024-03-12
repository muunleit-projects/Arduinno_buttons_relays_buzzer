// The Board is connected to two buttons and two relays.
// If one button ist pressed, the corresponding relays channel is switched and sound plays.

// Definitions for the buttons
const int buttonGreen = 8;  // pin for Button
const int buttonRed = 9;    // pin for Button
// places to capture the button-states
int buttonGreenPressed = 0;
int buttonRedPressed = 0;

// Definitions for the relays
const int relaysGreen = 2;  // pin for relays
const int relaysRed = 3;    // pin for relays
// the relays uses GND for ON and anything else for OFF
#define ON 0
#define OFF 1

// Definitions for the buzzer
const int buzzer = 10;  // pin for buzzer
// Defining frequency of each music note
#define NOTE_C4 262
#define NOTE_D4 294
#define NOTE_E4 330
#define NOTE_F4 349
#define NOTE_G4 392
#define NOTE_A4 440
#define NOTE_B4 494
#define NOTE_C5 523
#define NOTE_D5 587
#define NOTE_E5 659
#define NOTE_F5 698
#define NOTE_G5 784
#define NOTE_A5 880
#define NOTE_B5 988

// SETUP
void setup() {
  // initialize buttons
  pinMode(buttonGreen, INPUT);
  pinMode(buttonRed, INPUT);
  // initialize relays
  pinMode(relaysGreen, OUTPUT);
  pinMode(relaysRed, OUTPUT);
  digitalWrite(relaysGreen, OFF);
  digitalWrite(relaysRed, OFF);
  // initialize buzzer
  pinMode(buzzer, OUTPUT);
}

// LOOP
void loop() {
  // read states
  buttonGreenPressed = digitalRead(buttonGreen);
  buttonRedPressed = digitalRead(buttonRed);

  // There is no need to make a race between two buttons.
  // If they are both pressed, just do nothing.
  if (buttonGreenPressed != buttonRedPressed) {
    // Wait for a fracture of a second for new input.
    delay(200);
    return;
  }

  // If one button is pressed, switch the corresponding relays ON.
  if (buttonRedPressed) {
    diffAnsw();
  } else {
    sameAnsw();
  }
  delay(5000);
}

// sameAnsw switches on the relays for green and starts a "positiv"-feeling array of notes.
void sameAnsw() {
  int notes[] = { NOTE_A4, 0, NOTE_A4, 0, NOTE_D5 };
  // int notes[] = { NOTE_F4, 0, NOTE_B4, 0, NOTE_A5 };
  // int durations[] = { 200, 10, 200, 10, 1000 };
  int durations[] = { 200, 10, 200, 100, 400 };
  const int totalNotes = sizeof(notes) / sizeof(int);

  digitalWrite(relaysGreen, ON);
  playSound(notes, durations, totalNotes);
  digitalWrite(relaysGreen, OFF);
}

// diffAnsw switches on the relays for red and starts a "negativ"-feeling array of notes.
void diffAnsw() {
  int notes[] = { NOTE_E4, 0, NOTE_C4, 0, 0 };
  int durations[] = { 200, 10, 500, 10, 1000 };
  // int notes[] = { NOTE_E4, 0, NOTE_D4, 0, NOTE_C4 };
  // int durations[] = { 200, 10, 200, 10, 1000 };
  const int totalNotes = sizeof(notes) / sizeof(int);

  digitalWrite(relaysRed, ON);
  playSound(notes, durations, totalNotes);
  digitalWrite(relaysRed, OFF);
}

// playSounds actually plays the give array of notes
void playSound(int notes[], int durations[], int totalNotes) {
  // Loop through each note
  for (int i = 0; i < totalNotes; i++) {
    const int currentNote = notes[i];
    float dur = durations[i];

    // Play tone if currentNote is not 0 frequency, otherwise pause (noTone)
    if (currentNote != 0) {
      tone(buzzer, currentNote, dur);  // tone(pin, frequency, duration)
    } else {
      noTone(buzzer);
    };

    // delay is used to delay for tone to finish playing before moving to next loop
    delay(dur);
  }
}
