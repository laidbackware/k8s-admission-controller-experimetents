package exampleadmissioncontroller

import (
	"encoding/json"
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
		klog.Error("empty body")
		http.Error(writer, "empty body", http.StatusBadRequest)
		return
	}
	klog.Info("Received request")

	if r.URL.Path != "/validate" {
		klog.Error("no validate")
		http.Error(writer, "no validate", http.StatusBadRequest)
		return
	}
	
	svc, err := unmarshalService(body, writer)
	if err != nil {
		return
	}

	klog.Info(svc)
	
}

func unmarshalService(body []byte, writer http.ResponseWriter) (svc corev1.Service, err error) {
	arRequest := admissionv1.AdmissionReview{}
	if err = json.Unmarshal(body, &arRequest); err != nil {
		klog.Error("incorrect body")
		http.Error(writer, "incorrect body", http.StatusBadRequest)
		klog.Error(body)
		return
	}

	raw := arRequest.Request.Object.Raw
	svc = corev1.Service{}

	if err = json.Unmarshal(raw, &svc); err != nil {
		klog.Error("error deserializing service")
	}
	return
}