package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T)()  {
  tests := []struct{
    name string
    tm time.Time
    want string 
  }{
    {
      name:"UTC",
      tm:time.Date(2022, 10,18,11,38,0,0, time.UTC),
      want: "18 Oct 2022 at 11:38",
    },
    {
      name: "Empty",
      tm :time.Time{},
      want: "",
    },
    {
      name: "EAT", 
      tm: time.Date(2022, 10, 18, 11 ,38, 0, 0,time.FixedZone("EAT", -3*60*60)),
      want: "18 Oct 2022 at 14:38",
    },
  }
  for _,tt := range tests{
    t.Run(tt.name, func(t *testing.T){ 
      hd := humanDate(tt.tm)
      if hd != tt.want{
        t.Errorf("want %q; got %q", tt.want, hd)
      }
    })
  }
}
