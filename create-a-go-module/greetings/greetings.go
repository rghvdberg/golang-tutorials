package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the given name.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Create a message using a random formate.
	message := fmt.Sprintf(randomFormat(), name)
	// break it
	//message := fmt.Sprintf(randomFormat())
	return message, nil

}

// Hellos returens a map that asscociates each of the named people
// with a greeting
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the
		// the name.
		messages[name] = message
	}
	return messages, nil

} 

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

//randomFormat returns one of a set of greetings messages. The returned
// message is selected at random
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}