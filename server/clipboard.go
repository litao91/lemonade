package server

import (
	"github.com/atotto/clipboard"
	"github.com/pocke/lemonade/lemon"
)

type Clipboard struct{}

var fallback string

func (_ *Clipboard) Copy(text string, _ *struct{}) error {
	<-connCh
	fallback = lemon.ConvertLineEnding(text, LineEndingOpt)
	return clipboard.WriteAll(fallback)
}

func (_ *Clipboard) Paste(_ struct{}, resp *string) error {
	<-connCh
	t, err := clipboard.ReadAll()
	if err != nil {
		*resp = fallback
	} else {
		*resp = t
	}
	return err
}
