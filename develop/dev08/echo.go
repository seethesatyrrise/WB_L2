package main

func echo(text interface{}) interface{} {
	print(text, "\n")

	return text
}
