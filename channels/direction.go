package main

import (
	"path/filepath"
	"io/ioutil"
	"regexp"
	"fmt"
	"time"
	"strconv"
	"math"
	"strings"
)

var (
	windRegex = regexp.MustCompile(`\d METAR. *EGLL \d*Z [A-Z]*(\d{5}KT|VRB\d{2}KT).*=`)
	validation = regexp.MustCompile(`.*TAF.*`)
	comment = regexp.MustCompile(`\w*#.*`)
	metaclose = regexp.MustCompile(`.*=`)
	variablewind = regexp.MustCompile(`.*VRB\d{2}KT`)
	winddironly = regexp.MustCompile(`\d{5}KT`)
	winddistance = [8]int{}
)

func parseToArray(txtchannel chan string, metachannel chan []string) {
	for text := range txtchannel {
		lines := strings.Split(text, "\n")
		metaslice := make([]string, 0, len(lines))
		string := ""

		for _, line := range lines {
			if validation.MatchString(line) {
				break
			}
			if !comment.MatchString(line) {
				string += strings.Trim(line, " ")
			}
			if metaclose.MatchString(line) {
				metaslice = append(metaslice, string)
				string = ""
			}
		}
		metachannel <- metaslice
	}
	close(metachannel)
}

func extractDirection(metachannel, windschannel chan []string) {
	for metas := range metachannel {
		winds := make([]string, 0, len(metas))
		for _, meta := range metas {
			if windRegex.MatchString(meta) {
				winds = append(winds, windRegex.FindAllStringSubmatch(meta, -1)[0][1])
			}
		}
		windschannel <- winds
	}
	close(windschannel)
}

func distribution(windschannel chan []string, distchannel chan[8]int) {
	for winds := range windschannel {
		for _, wind := range winds {
			if variablewind.MatchString(wind) {
				for i := 0 ; i < 8 ; i++ {
					winddistance[i]++
				}
			} else if validation.MatchString(wind) {
				windstring := winddironly.FindAllStringSubmatch(wind, -1)[0][1]
				if d, err := strconv.ParseFloat(windstring, 64); err == nil {
					index := int(math.Round(d/45.0)) % 8
					winddistance[index]++
				}
			}
		}
	}
	distchannel <- winddistance
}

func main() {

	textChannel := make(chan string)
	metaChannel := make(chan []string)
	windsChannel := make(chan []string)
	resultChannel := make(chan [8]int)

	go parseToArray(textChannel, metaChannel)
	go extractDirection(metaChannel, windsChannel)
	go distribution(windsChannel, resultChannel)

	abspath, _ := filepath.Abs("./metafiles")
	files, _ := ioutil.ReadDir(abspath)
	start := time.Now()

	for _, file := range files {
		data, err := ioutil.ReadFile(filepath.Join(abspath, file.Name()))
		if err != nil {
			panic(err)
		}

		text := string(data)
		textChannel <- text
	}

	close(textChannel)

	results := <-resultChannel
	elapsed := time.Since(start)

	fmt.Printf("%v\n", winddistance)
	fmt.Printf("Processing took %s\n", elapsed)
	fmt.Println(results)
}