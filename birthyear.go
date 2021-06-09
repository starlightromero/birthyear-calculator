package main

import (
	"bufio"
	"fmt"
	"os"
  "log"
  "errors"
  "time"
	"strconv"
)

var (
  reader = bufio.NewReader(os.Stdin)
  errCannotReadInput = errors.New("cannot read input")
)

func main() {
	age := getAge()
	birthMonth := getBirthMonth()
	birthDay := getBirthDay()
	birthYear := getBirthYear(birthMonth, birthDay, age)
  fmt.Println("Your birth year is", birthYear)
}

func getBirthYear(bMonth, bDay, age int) int {
  curYear, curMonthDate, today := time.Now().Date()
  curMonth := int(curMonthDate)
  if bMonth - curMonth == 0 && bDay - today <= 0 || bMonth - curMonth < 0 {
    return curYear - age
  }
  return curYear - age - 1
}

func getAge() int {
	fmt.Print("Enter your age: ")
	age, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(errCannotReadInput)
  }
  aStr := trimNewline(age)
  a := toInt(aStr)
  if a < 0 {
    log.Fatal(errors.New("input must be greater than 0"))
  }
	return a
}

func getBirthMonth() int {
	fmt.Print("Enter your month of birth (1-12): ")
	month, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(errCannotReadInput)
  }
  mStr := trimNewline(month)
  m := toInt(mStr)
  if m < 1 || m > 12 {
    log.Fatal(errors.New("input must be between 1 and 12"))
  }
	return m
}

func getBirthDay() int {
	fmt.Print("Enter your day of birth (1-31): ")
	day, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(errCannotReadInput)
  }
  dStr := trimNewline(day)
  d := toInt(dStr)
  if d < 1 || d > 31 {
    log.Fatal(errors.New("input must be between 1 and 31"))
  }
	return d
}

func trimNewline(s string) string {
  l := len(s)
  t := s[:l-1]
  return t
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(errors.New("Cannot convert input to int"))
	}
	return i
}
