package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

const prompt = "and don't type your number in,just press ENTER when ready."

func main() {

	var firstNumber = rand.IntN(8) + 2
	var secondNumber = rand.IntN(8) + 2
	var subtraction = rand.IntN(8) + 2
	var answer = firstNumber*secondNumber - subtraction
	//fmt.Println(firstNumber, secondNumber, subtraction)
	playGame(firstNumber, secondNumber, subtraction, answer)

}

func playGame(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)

	//	 display welcome/instrucions
	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the game
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide your number by the number your originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The awnser is", answer)
}
