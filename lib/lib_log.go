package lib

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/drewstinnett/gout/v2"
	"github.com/gookit/color"
)

type LogStreamer struct {
	b bytes.Buffer
}

func (l *LogStreamer) String() string {
	return l.b.String()
}

func (l *LogStreamer) Write(p []byte) (n int, err error) {
	a := strings.TrimSpace(string(p))
	l.b.WriteString(a + "\n")
	log.Println(a)
	return len(p), nil
}

func ActionLog(Header string, char rune) {
	//conmpare rune
	if char == 0 {
		char = '*'
	}
	count := strings.Count(Header, "\x1b")
	str_len := len(Header) + 6
	if count > 0 {
		str_len = str_len - (count * 4) - 1
	}
	b := make([]rune, str_len)
	for i := range b {
		b[i] = char
	}
	colour := color.FgBlue.Render

	fmt.Println(colour(string(b)))
	fmt.Println(colour("** "+Header) + colour(" **"))
	fmt.Println(colour(string(b)))

}
func ActionLogOK(Header string, char rune) {
	color := color.FgGreen.Render
	ActionLog(fmt.Sprintf("%s: %s", Header, color("OK")), char)
}
func ActionLogFail(Header string, char rune) {
	color := color.FgRed.Render
	ActionLog(fmt.Sprintf("%s: %s", Header, color("Failed")), char)
}
func PrintlnOK(msg string) {
	color := color.FgGreen.Render
	log.Printf("%s: %s\n", msg, color("OK"))
}
func PrintOK(msg string) {
	color := color.FgGreen.Render
	log.Printf("%s: %s", msg, color("OK"))
}
func PrintlnFail(msg string) {
	color := color.FgRed.Render
	log.Printf("%s: %s\n", msg, color("Fail"))
}
func PrintFail(msg string) {
	color := color.FgRed.Render
	log.Printf("%s: %s", msg, color("Fail"))
}
func LogVerbose(msg string) {
	color := color.FgMagenta.Render
	fmt.Print(color(msg))

}
func FormatResults(title string, data interface{}) error {
	color := color.FgMagenta.Render
	fmt.Printf("%s\n", color(title))
	w, err := gout.New()
	if err != nil {
		panic(err)
	}
	// By Default, print to stdout. Let's change it to stderr though
	w.SetWriter(os.Stderr)

	// Print it on out!
	w.Print(data)
	return nil
}
