package internal

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	step  int8     // campo privato
	steps []string // campo privato
}

func newTask() *Task {
	return &Task{
		step:  0,
		steps: []string{"stepZero", "stepUno", "stepDue", "stepTre", "stepQuattro"},
	}
}
func (t *Task) nextStepOk() int8 {
	fmt.Println("in next step ...")
	fmt.Print("From " + t.steps[t.step] + " to ")
	t.step++
	t.step = t.step % 5
	fmt.Println(t.steps[t.step] + " ... ")
	return t.step
}

func (t *Task) randomDelay() {
	rand.Seed(time.Now().UnixNano())
	randomDuration := time.Duration(rand.Intn(500)+200) * time.Millisecond
	time.Sleep(randomDuration)
}

func (t *Task) NextStep() (int8, error) {
	fmt.Println("Try to go to next step...")
	randomNumber := rand.Intn(3)
	if randomNumber == 0 {
		return -1, errors.New("something went wrong") //errore generico
	} else if randomNumber == 1 {
		//t.randomDelay() //tempo
	}
	return t.nextStepOk(), nil //va bene
}
