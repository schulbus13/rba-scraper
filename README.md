# rba-scraper
erfasse Informationen zu allen Battles der RBA in einer JSON-Datei

## idee
In der [Battlerap-Datenbank](https://www.battlerap-datenbank.de) möchte ich gern auch die [RBA-Battles](https://www.r-b-a.de) erfassen. Diese sind allerdings nicht grundsätzlich auf YouTube zu finden, da muss ich also an der Datenbank mal noch was ändern. Wie auch immer, von den mehr als 80000 RBA-Battles brauche ich die Namen beider Rapper*innen bzw. Teams und den Link zum Battle.

## vorgehen
Die Battles haben alle eine Battle-Nummer, über die sie leicht aufgerufen werden können. Auf jeder Seite befinden sich dann zwei h2-Elemente, von denen eines die benötigten Informationen zu den Artists enthält. Da die Seiten automatisch generiert werden, ist der entsprechende String leicht zu splitten.

## resultat
Die Daten werden in eine JSON-Datei geschrieben. Jedes JSON-Objekt entspricht einem einzelnen Battle und hat die Werte "Rapper1", "Rapper2" und "Link". Denen sind dann entsprechend die Namen der am Battle beteiligten Artists und den dazugehörigen Link als String zugeordnet.

## ausführen
Der Code kann mit `go run rba-scraper.go` über die Kommandozeile ausgeführt werden. Die JSON-Datei wird im Ordner der GO-Datei unter dem Namen _rba.json_ gespeichert und überschreibt eine ggfs. bereits bestehende Datei mit diesem Namen. Da über 80000 (in Worten: achtzigtausend) Battles durchsucht werden, dauert das Ausführen wirklich lange (bei mir ca. anderthalb Stunden). Während der Ausführung werden alle Battles auch auf die Kommandozeile ausgegeben und erst am Schluss wird die JSON-Datei tatsächlich geschrieben. Bei Abbruch der Ausführung kann es also vorkommen, dass gar nichts in eine Datei geschrieben wird.
