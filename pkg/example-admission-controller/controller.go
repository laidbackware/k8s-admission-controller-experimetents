package exampleadmissioncontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"k8s.io/klog/v2"

	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
)

//ExampleServerHandler listen to admission requests and serve responses
type ExampleServerHandler struct {

}

func (handler *ExampleServerHandler) Validate(writer http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		if data, err := io.ReadAll(r.Body); err == nil {
			body = data
		}
	}
	if len(body) == 0 {
		klog.Error("empty body!")
		http.Error(writer, "empty body!", http.StatusBadRequest)
		return
	}
	klog.Info("Received request")

	if r.URL.Path != "/validate" {
		klog.Error("Invalid endpoint")
		http.Error(writer, "Invalid endpoint", http.StatusBadRequest)
		return
	}
	
	svc, err := unmarshalService(body, writer)
	if err != nil {
		klog.Error(err)
		http.Error(writer, fmt.Sprint(err), http.StatusBadRequest)
		return
	}

	klog.Info(svc)
	klog.Error("it works!")

	
	
}

func unmarshalService(body []byte, writer http.ResponseWriter) (svc corev1.Service, err error) {
	arRequest := admissionv1.AdmissionReview{}
	if err = json.Unmarshal(body, &arRequest); err != nil {
		klog.Error("body if request is not a json object")
		http.Error(writer, "body if request is not a json object", http.StatusBadRequest)
		klog.Error(body)
		return
	}

	if arRequest.Request != nil{
		raw := arRequest.Request.Object.Raw
		svc = corev1.Service{}
	
		if err = json.Unmarshal(raw, &svc); err != nil {
			klog.Error("error deserializing service")
		}
		return
	}
	err = errors.New("invalid input. not an admission review object")
	klog.Error("error deserializing service")
	return
}