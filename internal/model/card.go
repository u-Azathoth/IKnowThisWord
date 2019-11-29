package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Card struct {
	ID              int     `json:"id"`
	Word            string  `json:"word"`
	Meaning         string  `json:"meaning"`
	RecognitionRate float32 `json:"recognition_rate"`
}

func (c *Card) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Word, validation.Required, validation.Length(3, 100)),
		validation.Field(&c.Meaning, validation.Required),
	)
}
