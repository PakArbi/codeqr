package main

import "time"

// Person adalah struktur data untuk data diri
type Person struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string    `json:"name,omitempty" bson:"name,omitempty"`
	Age           int       `json:"age,omitempty" bson:"age,omitempty"`
	MaritalStatus string    `json:"maritalStatus,omitempty" bson:"maritalStatus,omitempty"`
	QRCode        string    `json:"qrCode,omitempty" bson:"qrCode,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
