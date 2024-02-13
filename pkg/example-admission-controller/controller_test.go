package exampleadmissioncontroller

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// var dummySvc = `{
// 	"apiVersion": "v1",
// 	"kind": "Service",
// 	"metadata": {
// 			"annotations": {
// 					"kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{},\"labels\":{\"app.kubernetes.io/managed-by\":\"tilt\"},\"name\":\"example-admission-controller-svc\",\"namespace\":\"example-admission-controller\"},\"spec\":{\"ports\":[{\"port\":443,\"targetPort\":8080}],\"selector\":{\"app\":\"example-admission-controller\"},\"type\":\"LoadBalancer\"}}\n"
// 			},
// 			"creationTimestamp": "2024-02-07T11:02:39Z",
// 			"labels": {
// 					"app.kubernetes.io/managed-by": "tilt"
// 			},
// 			"name": "example-admission-controller-svc",
// 			"namespace": "example-admission-controller",
// 			"resourceVersion": "877916",
// 			"uid": "2bc48566-1644-4bac-900d-a546cbd3e65d"
// 	},
// 	"spec": {
// 			"allocateLoadBalancerNodePorts": true,
// 			"clusterIP": "10.96.71.8",
// 			"clusterIPs": [
// 					"10.96.71.8"
// 			],
// 			"externalTrafficPolicy": "Cluster",
// 			"internalTrafficPolicy": "Cluster",
// 			"ipFamilies": [
// 					"IPv4"
// 			],
// 			"ipFamilyPolicy": "SingleStack",
// 			"ports": [
// 					{
// 							"nodePort": 31131,
// 							"port": 443,
// 							"protocol": "TCP",
// 							"targetPort": 8080
// 					}
// 			],
// 			"selector": {
// 					"app": "example-admission-controller"
// 			},
// 			"sessionAffinity": "None",
// 			"type": "LoadBalancer"
// 	},
// 	"status": {
// 			"loadBalancer": {
// 					"ingress": [
// 							{
// 									"ip": "100.127.255.2"
// 							}
// 					]
// 			}
// 	}
// }`

func TestDeserialize(t *testing.T) {
	writer := httptest.NewRecorder()

	_, err := unmarshalService([]byte(`{"apiVersion": "v1"}`), writer)
	assert.Error(t, err, errors.New("invalid input. not an admission review object"))
	
	_, err = unmarshalService([]byte("not_json"), writer)
	assert.Error(t, err, errors.New("ibody if request is not a json object"))
	
}