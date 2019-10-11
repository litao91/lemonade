package server

import (
	"log"

	"github.com/atotto/clipboard"
	"github.com/lemonade-command/lemonade/lemon"
)

type Clipboard struct{}

var fallback string

func (_ *Clipboard) Copy(text string, _ *struct{}) error {
	<-connCh
	fallback = lemon.ConvertLineEnding(text, LineEndingOpt)
	err := clipboard.WriteAll(fallback)
	if err != nil {
		log.Printf("%v\n", err)
	}
	return nil
}

func (_ *Clipboard) Paste(_ struct{}, resp *string) error {
	<-connCh
	t, err := clipboard.ReadAll()
	if err != nil {
		*resp = fallback
	} else {
		*resp = t
	}
	return nil
}
