package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Mahasiswa
type Mahasiswa struct {
	ID   int    `json:"id"`
	NIM  int    `json:"nim"`
	Name string `json:"name"`
}

// NewMahasiswa
func NewMahasiswa() []Mahasiswa {
	mhs := []Mahasiswa{
		Mahasiswa{
			ID:   1,
			NIM:  123756675454,
			Name: "Didik Prabowo",
		},
		Mahasiswa{
			ID:   2,
			NIM:  9245753454,
			Name: "Joni Gunawan",
		},
		Mahasiswa{
			ID:   3,
			NIM:  92342343454,
			Name: "Gugus Irwan",
		},
		Mahasiswa{
			ID:   4,
			NIM:  9232423423454,
			Name: "Muhammad Irwan",
		},
		Mahasiswa{
			ID:   5,
			NIM:  9232324454,
			Name: "Sira Irwan",
		},
		Mahasiswa{
			ID:   6,
			NIM:  9234547978,
			Name: "Muhammad dd",
		},
		Mahasiswa{
			ID:   7,
			NIM:  92345435434,
			Name: "Alen Irwan",
		},
	}
	return mhs
}

// GetMahasiswa
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mhs := NewMahasiswa()
		datamahasiswa, err := json.Marshal(mhs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datamahasiswa)
		return
	}

	http.Error(w, "Maaf salah server!", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/data", GetMahasiswa)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":7", nil); err != nil {
		log.Fatal(err)
	}
}
