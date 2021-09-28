RPS - CBG 
==========

DEV
---
```commandline
# execute make:dev to run
make dev

# execute make:run to execute
make run
```

RUN 
------
```commandline
notifyd --config=data/config.ini
# WINDOWS
notifyd.exe --config=data/config.ini
```

CONFIG (config.ini)
--------------------------
```editorconfig
[runtime]
# host - kontekst w jakim należy zapisać dane, wartości ustawiane w DB cbg.radiopoznan.fm (np. rp -> Radio poznań)
ctx     = contextValue
# token - zabezpieczenie zapisywania danych w kontekscie
token   = tokenValue
# host - adres http gdzie przesyłać dane do zapisu (lokalizacja aplikacji zapisującej CBG)
host    = http://localhost
# file - ścieszka absolutna pliku tekstowego do obserwacji zmiany treści
# (plik w którym zapisywane jest co obecnie jest grane)
file    = /usr/XXX/tmp/inputFile.txt
# timeout - co ile sekund [s] należy sprawdzić zawartość pliku wskazanego przez 'file'
timeout = 3
```