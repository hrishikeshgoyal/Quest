package main

func main() {
	software := NewSoftwareStateContext()

	software.print()
	software.Develop()
	software.print()
	software.Develop()
	software.print()
	software.Launch()
	software.print()
	software.Deprecate()
	software.print()
	software.Develop()
	software.print()
}
