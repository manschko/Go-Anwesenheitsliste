# Go-Anwesenheitsliste

Anwesenheitsliste über QR code

## Run
### comandline
Build binary
```
go build
```
Testing
```
go test
```
Run Program
```
go run main.go
```
### goland
Um Flags in Goland hinzuzufügen  

Run > Edit Configuration... 

Unter dem Feld **"Program arguments"** könne Flags einfach definiert werden
## Flags
| Flags      | Description                        | Expacted type  |
|------------|------------------------------------|----------------|
| -portQR    | port für die QR code Seite         | int (Ganzzahl) |
| -portLogin | port für die An- und Abmelderseite | int (Ganzzahl) |
| -valid     | Gültigkeits dauer der Token in Sekunden       | int (Ganzzahl) |
| -url    | URL für das erreichen der Webseite      | string (text) |