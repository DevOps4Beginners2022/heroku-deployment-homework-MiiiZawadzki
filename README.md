# Heroku deployment homework

## Zadanie 1

Przygotuj Dockerfile, który skopiuje pliki naszej aplikacji, zbuduję ją i uruchomi.

Wymagania dot. dockefile:

1. Uzyj obrazu `golang:1.17-alpine`
2. Zbuduj aplikacje 
  - skopiuj pliki
  - uruchom polecenie: go mod download
  - zbuduj aplikacje za pomocą polecenia: `go build -o app-server` 
3. Zezwól na ustawienie wartości dla instrukcji `EXPOSE` za pomocą zmiennej środowiskowej `PORT`.
4. Ustaw instrukcje CMD, tak zeby uruchamiala naszą aplikacje podczas startu kontenera.


## Zadanie 2

Przygotuj manifest heroku.yml, który będzie zawierał:
1. Addon `heroku-postgresql`
2. Sekcje build wskazującą na nasz Dockerfile

## Zadanie 3

Na podstawie manifestu heroku.yml utwórz aplikacje na heroku. 
Adres twojej aplikacji (pełny adres z protokołem https:// ) zapisz w pliku hostname.txt

## Zadanie 4

Ustaw zmienną środowiskową `SECRET` na poziomie panelu heroku z wartością: `d4f74c4d3my`

## Zadanie 5

Wypchnij swoje zmiany na repozytorium heroku, zeby wyzwolić proces wdrozenia aplikacji.