package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ein einzelnes RBA-Battle mit zwei unterschiedlichen Rappern bzw. Teams und dem dazugehörigen Link
type Battle struct {
	Rapper1 string
	Rapper2 string
	Link    string
}

func main() {
	//	die Menge aller erfassten Battles
	battleSlice := []Battle{}

	//	Schleife über alle Battle-Nummern von 1 bis 81644
	for i := 1; i < 81645; i++ {
		//	erster Versuch pro Battle-Nummer ist die ID 4101
		link := "https://www.r-b-a.de/index.php?ID=4101&BATTLE=" + strconv.Itoa(i)
		//	erfasse das zweite h2-Element auf der Seite
		title := getTitle(link, i)

		//	falls es kein Battle mit der ID 4101 und der entsprechenden Battle-Nummer gibt, beginnt das h2-Element mit " vs."
		if strings.HasPrefix(title, " vs.") {
			//	zweiter Versuch: die ID 4020
			link = "https://www.r-b-a.de/index.php?ID=4020&BATTLE=" + strconv.Itoa(i)
			title = getTitle(link, i)
		}

		//	damit der String aus dem h2-Element gesplittet werden kann, werden die Positionen von zwei Teilstrings bestimmt
		vsPosition := strings.Index(title, " vs. ")
		pointsPosition := strings.Index(title, " (")

		//	in manchen Fällen haben Artists " (" in ihrem Namen, das führt zu Problemen
		//	TODO besser abfangen wäre notwendig, hier gehen ein paar Battles verloren
		if vsPosition+5 >= pointsPosition {
			continue
		}

		//	die Informationen werden in einem neuen Battle gespeichert, das im Anschluss zum Slice aller Battles hinzugefügt wird
		battle := Battle{Rapper1: title[:vsPosition], Rapper2: title[vsPosition+5 : pointsPosition], Link: link}
		battleSlice = addToSlice(battle, battleSlice, i)
	}

	//	alle gespeicherten Battles werden in eine JSON-Datei geschrieben
	printBattles(battleSlice)
}

// ein Battle wird zum Slice aller Battles hinzugefügt, wenn beide Rapper auch tatsächlich aufgeführt sind
func addToSlice(battle Battle, slice []Battle, i int) []Battle {
	//	ist einer der beiden Rapper ein leerer String, wird das Battle übersprungen
	if len(strings.TrimSpace(battle.Rapper1)) == 0 || len(strings.TrimSpace(battle.Rapper2)) == 0 {
		return slice
	}

	//	ist einer der beiden Rapper als "R-I-P No. X" aufgeführt (vermutlich wegen Löschung des RBA-Kontos), wird das Battle übersprungen
	if strings.Contains(battle.Rapper1, "R-I-P No.") || strings.Contains(battle.Rapper2, "R-I-P No.") {
		return slice
	}

	//	alternative Schreibweise für gelöschte Kontos ist "RIPX" mit X als Zahl, auch diese Battles werden übersprungen
	//	möglicherweise werden hier auch tatsächlich existierende Artists übersprungen, aber das wird einfach in Kauf genommen
	if strings.HasPrefix(battle.Rapper1, "RIP") || strings.HasPrefix(battle.Rapper2, "RIP") {
		return slice
	}

	//	die Battle-Nummer und die beiden Artists werden auf die Konsole ausgegeben
	fmt.Printf("%v: %v vs. %v\n", i, battle.Rapper1, battle.Rapper2)

	// zurückgegeben wird der Slice mit angehängtem neuen Battle
	return append(slice, battle)
}

// alle gespeicherten Battles werden in eine JSON-Datei geschrieben
func printBattles(slice []Battle) {
	//	umwandeln des Slices in ein JSON-Objekt
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	//	die Links enthalten das Zeichen "&", das soll nicht escaped werden
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(slice)

	if err != nil {
		panic(err)
	}

	//	speichere das JSON-Objekt in der Datei "rba.json"
	err = os.WriteFile("rba.json", buffer.Bytes(), 0664)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("-----------------\nsuccessfully written to rba.json")
	}
}

// das zweite h2-Element von der mittels link übergebenen Seite als String zurückgeben
// Großteil dieses Codes übernommen von zetcode.com/golang/goquery
func getTitle(link string, i int) string {
	resp, err := http.Get(link)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		panic(err)
	}

	//	aus der Liste aller gefundenen h2-Elemente wird ein Slice mit nur dem zweiten Element gemacht,
	//	das dann in einen String konvertiert wird
	return doc.Find("h2").Slice(1, 2).Text()
}
