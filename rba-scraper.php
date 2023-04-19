<?php

/** 
 * 
 * PHP-Skript zum erfassen aller RBA-Battles (r-b-a.de) für die Battlerap-Datenbank (battlerap-datenbank.de)
 * 
 * es wird eine CSV-Datei mit den Namen beider Battle-Parteien und dem Link zum Battle für alle Battles geschrieben
 *
 * @author Felix Milke <battlerap-datenbank@gmail.com>
 * @version 1.0
 * 
 * @todo Ausgabedatei zu einer JSON-Datei umwandeln
 * @todo Skript zu einer lokal ausführbaren Datei machen (dann wohl kein PHP mehr)
 */

//  unbegrenzt Zeit zur Verfügung stellen, Programm bricht sonst zu früh ab
set_time_limit(0);

//  Schleife über alle BATTLE-IDs: 1 bis 81644
for($i = 1; $i < 81645; $i++) {
    //  h2-Elemente finden, benötigt wird das zweite
    $link = "https://www.r-b-a.de/index.php?ID=4101&BATTLE=" . $i;
    $h2element = getH2Element($link);

    //  wenn der Inhalt vom h2-Element " vs. (? : ?)" ist, dann die andere ID (4020 statt 4101) versucht
    if(strpos($h2element, " vs.") == 0) {
        $link = "https://www.r-b-a.de/index.php?ID=4020&BATTLE=" . $i;
        $h2element = getH2Element($link);
    }

    //  aus dem h2-Element die beiden Rapper extrahieren
    $vsPosition = strpos($h2element, " vs. ");
    $pointsPosition = strpos($h2element, " (");
    $rapper1 = substr($h2element, 0, $vsPosition);
    $rapper2 = substr($h2element, $vsPosition + 5, $pointsPosition - $vsPosition - 5);

    //  beide Rapper + den Link speichern
    write($rapper1, $rapper2, $link);
}


/**
 *  Daten in eine bestimmte Datei schreiben
 * 
 * @todo Dateipfad sollte übergeben werden können
 * @todo Ausgabe der Zahl der geschriebenen Einträge am Ende, an Stelle jedes einzelnen Links
 * @todo Datei am Anfang einmal öffnen, am Ende schließen, statt das jedes Mal zu tun
 */
function write($rapper1, $rapper2, $link) {
    //  bei gelöschten Accounts wird der Rappername durch "R-I-P" und eine Nummer ersetzt, die Battles sollen nicht gespeichert werden
    if(strpos($rapper1, "R-I-P") !== false || strpos($rapper2, "R-I-P") !== false) return;

    //  ist einer der Rappernamen leer, soll das Battle auch nicht gespeichert werden
    if(strlen($rapper1) == 0 || strlen($rapper2) == 0) return;
    
    //  Datei öffnen, "a" fügt neue Inhalte ans Ende der Datei
    $file = fopen("rba-data.csv", "a");

    //  einen String in die Datei schreiben, der beide Rapper und den Link zum Battle, durch Kommas getrennt, enthält
    fwrite($file, $rapper1 . "," . $rapper2 . "," . $link . "\n");

    //  Ausgabe auf die Seite, aber das sollte mal weg
    echo nl2br($link . "\n");

    //  Datei schließen
    fclose($file);
}

//  finde das zweite h2-Element auf der angegebenen Website
function getH2Element($link) {
    $dom = new DOMDocument;
    @$dom->loadHTMLFile($link);     // das @ verhindert irgendeinen Fehler, welchen genau hab ich schon wieder vergessen
    $elements = $dom->getElementsByTagName('h2');

    //  werden weniger oder mehr als zwei h2-Elemente gefunden, wird ein Platzhalter zurückgegeben, der später abgefangen wird
    if($elements->length != 2) return " vs. (? : ?)";

    //  werden genau zwei h2-Elemente gefunden, wird das zweite zurückgegeben
    return $elements->item(1)->nodeValue;
}

?>