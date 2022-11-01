package students

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/student-management/models"
	"github.com/student-management/service"

	"github.com/gorilla/mux"
)

type studentApi struct {
	studentService service.Service
}

func New(studentService service.Service) *studentApi {
	return &studentApi{studentService: studentService}
}

func (ua *studentApi) GetstudentByIdHandler(wr http.ResponseWriter, req *http.Request) {
	// set content-type to json
	wr.Header().Set("content-type", "application/json")

	// get id from url
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := models.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
		return
	}

	// get data from service layer
	usrData, err := ua.studentService.GetById(id)
	if err != nil {
		resErr := models.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: err.Error()}
		wr.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
		return
	}

	// convert student data to json object and write on the response
	outData := models.HttpResponse{
		Data:       usrData,
		Message:    "Retrieved",
		StatusCode: http.StatusOK,
	}
	resp, _ := json.Marshal(outData)
	_, _ = wr.Write(resp)
	wr.WriteHeader(http.StatusOK)
}

func (ua *studentApi) GetAllstudentHandler(wr http.ResponseWriter, req *http.Request) {
	// set content-type to json
	wr.Header().Set("content-type", "application/json")

	// get data from service layer
	data, err := ua.studentService.GetAll()
	if err != nil {
		resErr := models.HttpErrs{ErrCode: http.StatusNotFound, ErrMsg: err.Error()}
		wr.WriteHeader(http.StatusNotFound)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
		return
	}

	// convert student data to json object and write on the response
	outData := models.HttpResponse{
		Data:       data,
		Message:    "Retrieved",
		StatusCode: http.StatusOK,
	}
	resp, _ := json.Marshal(outData)
	_, _ = wr.Write(resp)
}

func (ua *studentApi) CreatestudentHandler(wr http.ResponseWriter, req *http.Request) {
	// set content-type to json
	wr.Header().Set("content-type", "application/json")
	body := req.Body
	var studentData models.Student
	err := json.NewDecoder(body).Decode(&studentData)
	if err != nil {
		respErr := models.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "bad request"}
		res, _ := json.Marshal(respErr)
		wr.WriteHeader(http.StatusInternalServerError)
		wr.Write(res)
		return
	}

	err = ua.studentService.Create(studentData)
	if err != nil {
		respErr := models.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, _ := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		wr.Write(res)
		return
	}

	wr.WriteHeader(http.StatusCreated)
	response := models.HttpResponse{
		Data:       studentData,
		Message:    "student created",
		StatusCode: http.StatusOK,
	}
	resp, _ := json.Marshal(response)
	wr.Write(resp)
}

func (ua *studentApi) UpdatestudentHandler(wr http.ResponseWriter, req *http.Request) {
	// set content-type to json
	wr.Header().Set("content-type", "application/json")
	body := req.Body
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := models.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
		return
	}

	var studentData models.Student

	err = json.NewDecoder(body).Decode(&studentData)
	if err != nil {
		resErr := models.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
		return
	}

	studentData.Id = id
	// call to service layer
	err = ua.studentService.Update(studentData)
	if err != nil {
		respErr := models.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, _ := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		_, _ = wr.Write(res)
		return
	}

	// give status OK (200) if everything goes OK
	wr.WriteHeader(http.StatusOK)
	response := models.HttpResponse{
		Data:       studentData,
		Message:    "Data updated",
		StatusCode: http.StatusOK,
	}
	resp, _ := json.Marshal(response)
	wr.Write(resp)
}

func (ua *studentApi) DeletestudentHandler(wr http.ResponseWriter, req *http.Request) {
	// set content-type to json
	wr.Header().Set("content-type", "application/json")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := models.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
		return
	}

	// call to service layer
	err = ua.studentService.Delete(id)
	if err != nil {
		respErr := models.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, _ := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		_, _ = wr.Write(res)
		return
	}

	// give status OK (200) if everything goes OK
	wr.WriteHeader(http.StatusOK)
	response := models.HttpResponse{
		Data:       id,
		Message:    "student deleted",
		StatusCode: http.StatusOK,
	}
	resp, _ := json.Marshal(response)
	wr.Write(resp)
}
