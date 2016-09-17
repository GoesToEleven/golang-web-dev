package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	pathTeacherProspectingUdemy = "/crm/teacher/prospecting/udemy"
)

func initCRM(r *httprouter.Router) {
	r.GET(pathTeacherProspectingUdemy, teacherProspectingUdemy)
}

func teacherProspectingUdemy(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	d := something()
	serveTemplate(res, "crmTeacherProspectingUdemy.gohtml", d)
}