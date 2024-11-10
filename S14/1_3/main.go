package main

// go get -u github.com/rs/zerolog  اضافه کردن پکیج زیرولاگ به پروژه

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel) //لول لاگ را مشخص کردیم

	log.Print("Test") // این لاکک ساختار دارد

	log.Info().Str("Category: ", "Search").Msg("search for a thing")

	err := errors.New("this is an error")
	log.Error().Err(err).Msg("Error occurred")

	err = func3()
	if err != nil {
		log.Error().Stack().Err(err).Msg("Error occurred func3 ")

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
