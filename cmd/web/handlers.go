package main

import "net/http"

func (app *application) virtualTerminal (w http.ResponseWriter, r *http.Request){
	
	err:=app.renderTemplate(w,r,"terminal",nil)
	if err!=nil{
		app.errorLog.Println(err)
	}
}