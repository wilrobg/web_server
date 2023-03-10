package main

import (
	"net/http"
)

// Struct router para hacer request en el servidor
type Router struct {
	//Reglas para definir de que rutas pasan a que handler, mapa que pasa de strings a handler
	//mapa que tenga como llaves strings y que mapee a HandlerFunc
	rules map[string]map[string]http.HandlerFunc
}

// forma de instanciar el router, similar al NewServer() del archivo servidor.go
func NewRouter() *Router {
	return &Router{
		//a diferencia del servidor, aqui el router debe empezar en un estado vacio, creamos un mapa vacio
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// Metodo ServeHTTP de router para poder implementar en el handler el atributo s.router en server.go
// parametros: el primero es el escritor, el segundo es el request en donde viene la informacion
// no olvidar colocar ServeHTTP con letras mayusculas
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	//impresion de mensaje respuesta que el servidor da a la ruta
	//Fprintf es un escritor, que recive w que es el escritor asignado, y el mensaje que queremos mostrar
	handler, methodExist, exist := r.FindHandler(request.Method, request.URL.Path)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	handler(w, request)
}

func (r *Router) FindHandler(method, path string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}
