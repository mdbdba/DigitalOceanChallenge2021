package ranking

import (
	"fmt"
	"math/rand"
	"time"
)

var battles = []string {
	"John Lennon vs Bill O'Reilly",
	"Darth Vader vs Adolf Hitler",
	"Abe Lincoln vs Chuck Norris",
	"Justin Bieber vs Beethoven",
	"Albert Einstein vs Stephen Hawking",
	"Genghis Khan vs The Easter Bunny",
	"Billy Mays vs Ben Franklin",
	"Gandalf vs Dumbledore",
	"Dr Seuss vs William Shakespeare",
	"Mr. T vs Mr. Rogers",
	"Captain Kirk vs Christopher Columbus",
	"Mario Bros. vs Wright Bros.",
	"Michael Jackson vs Elvis Presley",
	"Cleopatra vs Marilyn Monroe",
	"Steve Jobs vs Bill Gates",
	"Frank Sinatra vs Freddie Mercury",
	"Doc Brown vs Doctor Who",
	"Bruce Lee vs Clint Eastwood",
	"Batman vs Sherlock Holmes",
	"Moses vs Santa Claus",
	"Adam vs Eve",
	"Gandhi vs Martin Luther King Jr.",
	"Nikola Tesla vs Thomas Edison",
	"Babe Ruth vs Lance Armstrong",
	"Mozart vs Skrillex",
	"Rasputin vs Stalin",
	"Blackbeard vs Al Capone",
	"Miley Cyrus vs Joan Of Arc",
	"Bob Ross vs Pablo Picasso",
	"Michael Jordan vs Muhammad Ali",
	"Rick Grimes vs Walter White",
	"Goku vs. Superman",
	"Stephen King vs Edgar Allan Poe",
	"Sir Isaac Newton vs Bill Nye",
	"George Washington vs William Wallace",
	"Ghostbusters vs Mythbusters",
	"Zeus vs Thor",
	"Jack The Ripper vs Hannibal Lecter",
	"Steven Spielberg vs Alfred Hitchcock",
	"David Copperfield vs Harry Houdini",
	"Terminator vs Robocop",
	"Shaka Zulu vs Julius Caesar",
	"Jim Henson vs Stan Lee",
	"J. R. R. Tolkien vs George R. R. Martin",
	"Gordon Ramsay vs Julia Child",
	"James Bond vs Austin Powers",
	"Bruce Banner vs Bruce Jenner",
	"Alexander the Great vs Ivan the Terrible",
	"Tony Hawk vs Wayne Gretzky",
	"Theodore Roosevelt vs Winston Churchill",
	"Batman vs Elon Musk",
	"Deadpool vs Boba Fett",
	"George Carlin vs Richard Pryor",
	"Guy Fawkes vs Che Guevara",
	"Harry Potter vs Luke Skywalker",
	"Jeff Bezos vs Mansa Musa",
	"Larry Bird vs Big Bird",
	"Mother Teresa vs Sigmund Freud",
	"Ragnar Lodbrok vs Richard The Lionheart",
	"Ronald McDonald vs The Burger King",
	"Thanos vs J. Robert Oppenheimer",
	"The Joker vs Pennywise",
	"Wolverine vs Freddy Krueger"}

// Voices export voices
var Voices = []string{
	"Darth Vader",
	"Macho Man",
	"Stephen Hawking",
	"Christopher Walken",
	"Marilyn Monroe",
	"Santa Claus",
	"Mother Teresa",
	"Harry Potter",
	"Big Bird",
	"Joan of Arc"}

var ratingValues = []int{
	0,  // Hated
	1,  // Meh
	2,  // Liked
}

// GenRatings - export this function
func GenRatings() ([]string, []string) {
	var ratings []string
	var ratingsHtml []string
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	for n := 0; n < 20 ; n++ {
		// randomly select a title, two voices, and a rating
		message := fmt.Sprintf("<tr><td>%s</td>", battles[rand.Intn(len(battles))])
		hmsg := fmt.Sprintf("%s|", battles[rand.Intn(len(battles))])
		tVoice1 := Voices[rand.Intn(len(Voices))]
		message = fmt.Sprintf("%s<td>%s</td>", message, tVoice1)
		hmsg = fmt.Sprintf("%s%s|", hmsg, tVoice1)
		tVoice2 := Voices[rand.Intn(len(Voices))]
		if tVoice1 == tVoice2 { //try again
			choice := rand.Intn(len(Voices))
			tVoice2 = Voices[choice]
			if tVoice1 == tVoice2 {
				if rand.Intn(100) < 50 {
					if choice == 0 {
						choice = len(Voices) -1
					} else {
						choice -= 1
					}
				} else {
					if choice == len(Voices) -1 {
						choice = 0
					} else {
						choice +=1
					}
				}
				tVoice2 = Voices[choice]
			}
		}
		rVal := ratingValues[rand.Intn(len(ratingValues))]
		message = fmt.Sprintf("%s<td>%s</td><td>%d</td><tr>",
			message, tVoice2, rVal)
		hmsg = fmt.Sprintf("%s%s|%d", hmsg, tVoice2, rVal)
		//fmt.Println(message)
		ratings = append(ratings, message)
		ratingsHtml = append(ratingsHtml, hmsg)
	}
	return ratings, ratingsHtml
}
