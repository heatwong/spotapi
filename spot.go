package main

// import "time"

type Spot struct {
   Latitude    float32  `json:"latitude"`
   Longitude   float32  `json:"longitude"`
   Altitude    float32  `json:"altitude"`
}

type Spots []Spot