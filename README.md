# rba-scraper
erfasse Informationen zu allen Battles der RBA in einer CSV-Datei

## idee
In der [Battlerap-Datenbank](https://www.battlerap-datenbank.de) möchte ich gern auch die [RBA-Battles](https://www.r-b-a.de) erfassen. Diese sind allerdings nicht grundsätzlich auf YouTube zu finden, da muss ich also an der Datenbank mal noch was ändern. Wie auch immer, von den mehr als 80000 RBA-Battles brauche ich die Namen beider Rapper*innen bzw. Teams und den Link zum Battle.

## vorgehen
Die Battles haben alle eine Battle-ID, über die sie leicht aufgerufen werden können. Auf jeder Seite befinden sich dann zwei h2-Elemente, von denen eines die benötigten Informationen zu den Artists enthält. Da die Seiten automatisch generiert werden, ist der entsprechende String leicht zu splitten.

## resultat
Die Daten werden zunächst in eine CSV-Datei geschrieben. Perspektivisch wäre ein lokal startbares Skript (also ohne erst einen PHP-Server starten zu müssen) wünschenswert, damit das Daten Scrapen möglichst einfach wiederholt werden kann. Für den Fall, dass mal noch jemand die Daten braucht. Auch wäre eine JSON-Datei als Output praktischer, da manche Namen Kommas enthalten, das macht ein bisschen was kaputt für die CSV-Variante.
