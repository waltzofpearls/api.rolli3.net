package main

import "log"

func main() {
	c := newConfig()
	if err := c.load(); err != nil {
		log.Fatalln("[ERROR] Failed to load config.", err)
	}
	l := newLogger(c)
	l.Fatal(newAPI(c, l).run())
}
