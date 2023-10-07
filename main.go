package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new Fyne application
	myApp := app.New()

	// Create 3 arrays for the table
	affirmativeAnswers := []string{"It is certain", "It is decidedly so", "Without a doubt", "Yes definitely", "You may rely on it", "As I see it, yes", "Most likely", "Outlook good", "Yes", "Signs point to yes"}
	nonCommittalAnswers := []string{"Reply hazy, try again", "Ask again later", "Better not tell you now", "Cannot predict now", "Concentrate and ask again", "Very doubtful"}
	negativeAnswers := []string{"Don’t count on it", "My reply is no", "My sources say no", "Outlook not so good"}

	// Create a map to store the arrays
	answers := map[string][]string{
		"Affirmative Answers":   affirmativeAnswers,
		"Non-Committal Answers": nonCommittalAnswers,
		"Negative Answers":      negativeAnswers,
	}

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Create a label to display the Magic 8 Ball's answer
	answerLabel := widget.NewLabel("")

	// Create a text input field for the user to enter a question
	questionEntry := widget.NewEntry()
	questionEntry.SetPlaceHolder("Please ask your question?")

	// Create a button to ask the Magic 8 Ball a question
	askButton := widget.NewButton("Ask", func() {
		// Get the user's question
		question := questionEntry.Text

		// Generate a random number between 0 and the length of the map - 1
		randomIndex := rand.Intn(len(answers))

		// Get the key (category) at the random index
		i := 0
		var category string
		for k := range answers {
			if i == randomIndex {
				category = k
				break
			}
			i++
		}

		// Get the answer from the array at the random index
		answer := answers[category][rand.Intn(len(answers[category]))]

		// Set the answer label text to the Magic 8 Ball's answer
		answerLabel.SetText(fmt.Sprintf("Question: %s\nAnswer: %s", question, answer))
	})

	// Create a button to quit the application
	quitButton := widget.NewButton("Quit", func() {
		myApp.Quit()
	})

	// Create a horizontal box to hold the buttons
	buttonBox := container.NewHBox(askButton, quitButton)

	// Create a vertical box to hold the question entry, buttons, and answer label
	content := container.NewVBox(
		layout.NewSpacer(),
		questionEntry,
		layout.NewSpacer(),
		buttonBox,
		layout.NewSpacer(),
		answerLabel,
		layout.NewSpacer(),
	)

	// Create a new window and set its content
	myWindow := myApp.NewWindow("Magic 8 Ball")
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 200))
	// Show the window and run the application
	myWindow.ShowAndRun()
}
