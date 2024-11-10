package main

// go get -u github.com/rs/zerolog  اضافه کردن پکیج زیرولاگ به پروژه

import (
	"LogToFile/core"
	"errors"
)

var logger = core.NewFileLogger()

func main() {

	logger.Print("Test") // این لاکک ساختار دارد

	logger.Info().
		Str("Category: ", "Search").
		Msg("search for a thing")

	err := errors.New("this is an error")
	logger.Error().Err(err).Msg("Error occurred")

	err = func3()
	if err != nil {
		logger.Error().Stack().Err(err).Msg("Error occurred func3 ")

	}
}

func func1() error {
	return errors.New("this is an error in func1")
}

func func2() error {
	err := func1()
	if err != nil {
		return err
	}
	return nil
}

func func3() error {
	err := func2()
	if err != nil {
		return err
	}
	return nil
}

// trace  debug info warn  error fatal(critical) panic   سطوح لاگ
