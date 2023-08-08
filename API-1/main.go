package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Se definește funcția de bază a serverului
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Atunci când o cerere HTTP este făcută către ruta "/" (rădăcina),
		// se va executa acest bloc de cod.

		log.Println("Hello World") // Se afișează un mesaj în consolă.

		// Se citește tot conținutul corpului cererii HTTP.
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oioioiiii", http.StatusBadRequest) // În caz de eroare, se returnează un mesaj de eroare și un cod de stare HTTP.
			return
		}
		log.Printf("data %s\n", d) // Se afișează conținutul citit în consolă.

		// Se trimite un răspuns către client cu un mesaj și conținutul citit.
		fmt.Fprintf(w, "Hello %s", d)
	})

	// Se definește o altă funcție pentru ruta "/bye".
	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		// Atunci când o cerere HTTP este făcută către ruta "/bye",
		// se va executa acest bloc de cod.
		log.Println("Bye World") // Se afișează un mesaj în consolă.
	})

	// Se pornește serverul pentru a asculta pe portul 9090.
	http.ListenAndServe(":9090", nil)
}

/*
http.HandleFunc:

Aceasta este o funcție în biblioteca standard a limbajului Go care vă permite să înregistrați funcții de tratare (handler functions) pentru anumite rute pe serverul dvs. Aceste funcții sunt apelate atunci când primiți cereri HTTP pe ruta specificată. Sintaxa sa este http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)).
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			Funcția de tratare pentru ruta "/"
			w: ResponseWriter - vă permite să trimiteți răspunsul către client.
			r: Request - conține informații despre cererea HTTP primită.

http.ResponseWriter:
Acest tip de obiect vă permite să interacționați cu răspunsul pe care serverul dvs. îl trimite către client. Prin intermediul acestui obiect, puteți scrie date în corpul răspunsului, seta antete HTTP și altele.

http.Error:
Această funcție vă permite să trimiteți un răspuns de eroare către client. Ea are sintaxa http.Error(w http.ResponseWriter, error string, code int). Prin utilizarea acestei funcții, puteți trimite răspunsuri de eroare structurate către clientul dvs. în loc să aruncați simple erori sau să întrerupeți execuția

http.ListenAndServe:
Aceasta este funcția care pornește serverul HTTP și îl pune în așteptare pentru cereri de la clienți. Sintaxa sa este http.ListenAndServe(addr string, handler http.Handler). addr reprezintă adresa și portul pe care serverul ascultă, iar handler este handler-ul global al serverului.

Când cineva accesează ruta specificată ca parametru în funcția http.HandleFunc, adică ruta specificată în argumentul pattern, atunci funcția handler asociată acelei rute va fi executată.

Mai exact, atunci când un client face o cerere HTTP către ruta specificată, serverul va căuta funcția de tratare (handler function) asociată acelei rute folosind șablonul pattern. Dacă găsește o potrivire, va apela acea funcție, pasându-i ca argumente un obiect http.ResponseWriter (pentru a trimite răspunsul) și un obiect http.Request (pentru a accesa informațiile despre cerere).

*/
