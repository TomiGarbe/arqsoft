package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

type Hotel struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

var hotels []Hotel

func main() {
	// Inicializar la lista de hoteles (para fines de prueba)
	hotels = []Hotel{
		{
			ID:          1,
			Title:       "Hotel A",
			Description: "Descripción del Hotel A",
			Image:       "hotelA.jpg",
		},
		{
			ID:          2,
			Title:       "Hotel B",
			Description: "Descripción del Hotel B",
			Image:       "hotelB.jpg",
		},
	}

	// Configurar las rutas del servidor
	router := mux.NewRouter()
	router.HandleFunc("/api/hotels", GetHotels).Methods("GET")
	router.HandleFunc("/api/hotels", CreateHotel).Methods("POST")

	// Iniciar el servidor
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetHotels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotels)
}

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	// Parsear los datos del formulario, incluyendo la imagen
	err := r.ParseMultipartForm(10 << 20) // tamaño máximo de la imagen: 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Leer los datos del formulario
	title := r.FormValue("title")
	description := r.FormValue("description")
	imageFile, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer imageFile.Close()

	// Generar un nombre único para la imagen
	imageName := uuid.New().String() + ".jpg" // usar extensión .jpg por simplicidad

	// Guardar la imagen en el sistema de archivos del servidor
	imagePath := filepath.Join("uploads", imageName)
	imageData, err := ioutil.ReadAll(imageFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = ioutil.WriteFile(imagePath, imageData, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Crear un nuevo hotel con el nombre del archivo de la imagen
	hotel := Hotel{
		ID:          len(hotels) + 1,
		Title:       title,
		Description: description,
		Image:       imageName,
	}
	hotels = append(hotels, hotel)

	// Devolver la respuesta con el hotel creado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotel)
}
