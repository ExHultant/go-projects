package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Usuario struct {
	ID     int
	Nombre string
	Titulo string
	Email  string
	Rol    string
}

func obtenerConexion() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/bdgolang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func crearUsuario(nombre, titulo, email, rol string) error {
	db, err := obtenerConexion()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO usuarios (nombre, titulo, email, rol) VALUES (?, ?, ?, ?)", nombre, titulo, email, rol)
	if err != nil {
		return err
	}

	return nil
}

func actualizarUsuario(id int, nombre, titulo, email, rol string) error {
	db, err := obtenerConexion()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE usuarios SET nombre = ?, titulo = ?, email = ?, rol = ? WHERE id = ?", nombre, titulo, email, rol, id)
	if err != nil {
		return err
	}

	return nil
}

func eliminarUsuario(id int) error {
	db, err := obtenerConexion()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM usuarios WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func mostrarFormulario(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("formulario.html"))
	tpl.Execute(w, nil)
}

func crearUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		titulo := r.FormValue("titulo")
		email := r.FormValue("email")
		rol := r.FormValue("rol")
		err := crearUsuario(nombre, titulo, email, rol)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error al crear el usuario", http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func actualizarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		titulo := r.FormValue("titulo")
		email := r.FormValue("email")
		rol := r.FormValue("rol")
		err := actualizarUsuario(id, nombre, titulo, email, rol)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error al actualizar el usuario", http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func eliminarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		err := eliminarUsuario(id)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error al eliminar el usuario", http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", mostrarFormulario)
	http.HandleFunc("/crear", crearUsuarioHandler)
	http.HandleFunc("/actualizar", actualizarUsuarioHandler)
	http.HandleFunc("/eliminar", eliminarUsuarioHandler)

	log.Println("Servidor escuchando en http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
