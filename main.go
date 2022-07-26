package main

//Proyecto de Api Rest de crear catalogo de juegos

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//definicion Struct/Objeto/Clase games
type games struct {
	//definicion de atributos de Struct/Objeto/Clase games
	//Adicional van a ser tipo Json ya que responderemos en Json
	Id       int64  `json:id`
	Name     string `json:name`
	Category string `json:category`
}

//Definimos lista, para lograr consultar, editar etc uno o varios juegos
type allGames []games

//definimos lista de elmentos dummy para pruebas
var vgames = allGames{
	{
		Id:       0,
		Name:     "GTA VI",
		Category: "Open World",
	},
}

//definimos funcion que se ejecutara en el endpoint /
//primer parametro w es la respuesta que le daremos al cliente
//el segundo el r peticion del usuario
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my Api")
}

//Definimos funciona que devuelve todas las tareas
func getGames(w http.ResponseWriter, r *http.Request) {
	//Responder formato JSON
	json.NewEncoder(w).Encode(vgames)
}

func main() {

	//El metodo .StrictSlash(true) es para decirle que siempre debe redigiri por ejemplo a /games y no a /games/, se vuelve ma estricto
	//instanciamos
	//creamos enrutador con libreria mux para servidor
	routerLocal := mux.NewRouter().StrictSlash(true)
	//definimos endpoint principal con HandlerFunc
	routerLocal.HandleFunc("/", indexRoute)
	//definimos endpoint para retornar todos los juegos
	routerLocal.HandleFunc("/games", getGames)
	//creamos servidor http, se pasa el metodo ListenAndServe
	//recibe dos parametros, el primero el puerto, el segundo el objeto del enrutador en este caso es router
	http.ListenAndServe(":3000", routerLocal)
	//para rastrear que sucede usamos libreria log
	log.Fatal(http.ListenAndServe(":3000", routerLocal))

}
