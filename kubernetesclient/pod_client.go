package kubernetesclient

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/api/v1"
)

type PodOperations interface {
	ByName(namespace string, name string) (*v1.Pod, error)
	CreatePod(namespace string, resource *v1.Pod) (*v1.Pod, error)
	DeletePod(namespace string, name string) error
}

func newPodClient(client *Client) *PodClient {
	return &PodClient{
		client: client,
	}
}

type PodClient struct {
	client *Client
}

func (c *PodClient) ByName(namespace string, name string) (*v1.Pod, error) {
	return c.client.K8sClient.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
}

func (c *PodClient) CreatePod(namespace string, resource *v1.Pod) (*v1.Pod, error) {
	return c.client.K8sClient.CoreV1().Pods(namespace).Create(resource)
}

func (c *PodClient) DeletePod(namespace string, name string) error {
	return c.client.K8sClient.CoreV1().Pods(namespace).Delete(name, &metav1.DeleteOptions{})
}
