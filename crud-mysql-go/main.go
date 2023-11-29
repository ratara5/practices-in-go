/*
CRUD EN Go y BD Mysql + TEMPLATES
https://www.youtube.com/watch?v=G58gN0lIbyI
*/
package main

import (
	"net/http"
	"log"
	"fmt"
	"text/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connectionBD()(connection *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := ""
	Name := "system"

	connection, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+Name)
	if err != nil {
		panic(err.Error())
	}
	return connection
}

var tmps = template.Must(template.ParseGlob("templates/*")) //templates es una carpeta en este directorio

func main() {
	http.HandleFunc("/", Home) //O Inicio en vez de Home
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)

	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)

}
type Employee struct {
	Id int
	Name string
	Email string
}
func Home(w http.ResponseWriter, r *http.Request){ //función q recibe w y r
	//fmt.Fprint(w, "Hi Developer")
	
	connectionOn := connectionBD()

	//insertReg, err := connectionOn.Prepare("INSERT INTO employees(name, email) VALUES('Alv', 'alv@mail.com')")
	//insertReg.Exec()

	regs, err := connectionOn.Query("SELECT * FROM employees")
		
	if err != nil {
		panic(err.Error())
	}

	employee := Employee{}
	arrayEmployees := []Employee{}

	for regs.Next(){
		var id int
		var name, email string
		err = regs.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.Email = email

		arrayEmployees = append(arrayEmployees, employee)
	}
	//fmt.Println(arrayEmployees)

	tmps.ExecuteTemplate(w, "home", arrayEmployees) //nil significa que no se le envían parámetros
}

func Create(w http.ResponseWriter, r *http.Request){ 
	tmps.ExecuteTemplate(w, "create", nil) 
}

func Insert(w http.ResponseWriter, r *http.Request){ 
	if r.Method == "POST"{
		name := r.FormValue("name") //El name del input para name
		email := r.FormValue("email") 
	
		connectionOn := connectionBD()

		insertReg, err := connectionOn.Prepare("INSERT INTO employees(name, email) VALUES(?, ?)")

		if err != nil {
			panic(err.Error())
		}

		insertReg.Exec(name, email)

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request){
	idEmployee := r.URL.Query().Get("id")
	fmt.Println(idEmployee)

	connectionOn := connectionBD()

	deleteReg, err := connectionOn.Prepare("DELETE FROM employees WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	deleteReg.Exec(idEmployee)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")

	employee := Employee{}
	connectionOn := connectionBD()

	//insertReg, err := connectionOn.Prepare("INSERT INTO employees(name, email) VALUES('Alv', 'alv@mail.com')")
	//insertReg.Exec()

	reg, err := connectionOn.Query("SELECT * FROM employees WHERE id=?", idEmployee)

	for reg.Next(){
		var id int
		var name, email string
		err = reg.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.Email = email
	}

	fmt.Println(employee)
	tmps.ExecuteTemplate(w, "edit", employee)
}

func Update(w http.ResponseWriter, r *http.Request){ 
	if r.Method == "POST"{
		id := r.FormValue("id") 
		name := r.FormValue("name") 
		email := r.FormValue("email") 
	
		connectionOn := connectionBD()

		updateReg, err := connectionOn.Prepare("UPDATE employees SET name=?, email=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		updateReg.Exec(name, email, id)

		http.Redirect(w, r, "/", 301)
	}
}
//Detener y Reiniciar servicio cada vez q se haga cambio en el código