// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/open-telemetry/opentelemetry-operator/pkg/apis/opentelemetry/v1alpha1"
	scheme "github.com/open-telemetry/opentelemetry-operator/pkg/client/versioned/scheme"
)

// OpenTelemetryCollectorsGetter has a method to return a OpenTelemetryCollectorInterface.
// A group's client should implement this interface.
type OpenTelemetryCollectorsGetter interface {
	OpenTelemetryCollectors(namespace string) OpenTelemetryCollectorInterface
}

// OpenTelemetryCollectorInterface has methods to work with OpenTelemetryCollector resources.
type OpenTelemetryCollectorInterface interface {
	Create(*v1alpha1.OpenTelemetryCollector) (*v1alpha1.OpenTelemetryCollector, error)
	Update(*v1alpha1.OpenTelemetryCollector) (*v1alpha1.OpenTelemetryCollector, error)
	UpdateStatus(*v1alpha1.OpenTelemetryCollector) (*v1alpha1.OpenTelemetryCollector, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OpenTelemetryCollector, error)
	List(opts v1.ListOptions) (*v1alpha1.OpenTelemetryCollectorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpenTelemetryCollector, err error)
	OpenTelemetryCollectorExpansion
}

// openTelemetryCollectors implements OpenTelemetryCollectorInterface
type openTelemetryCollectors struct {
	client rest.Interface
	ns     string
}

// newOpenTelemetryCollectors returns a OpenTelemetryCollectors
func newOpenTelemetryCollectors(c *OpentelemetryV1alpha1Client, namespace string) *openTelemetryCollectors {
	return &openTelemetryCollectors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the openTelemetryCollector, and returns the corresponding openTelemetryCollector object, and an error if there is any.
func (c *openTelemetryCollectors) Get(name string, options v1.GetOptions) (result *v1alpha1.OpenTelemetryCollector, err error) {
	result = &v1alpha1.OpenTelemetryCollector{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("opentelemetrycollectors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OpenTelemetryCollectors that match those selectors.
func (c *openTelemetryCollectors) List(opts v1.ListOptions) (result *v1alpha1.OpenTelemetryCollectorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OpenTelemetryCollectorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("opentelemetrycollectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested openTelemetryCollectors.
func (c *openTelemetryCollectors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("opentelemetrycollectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a openTelemetryCollector and creates it.  Returns the server's representation of the openTelemetryCollector, and an error, if there is any.
func (c *openTelemetryCollectors) Create(openTelemetryCollector *v1alpha1.OpenTelemetryCollector) (result *v1alpha1.OpenTelemetryCollector, err error) {
	result = &v1alpha1.OpenTelemetryCollector{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("opentelemetrycollectors").
		Body(openTelemetryCollector).
		Do().
		Into(result)
	return
}

// Update takes the representation of a openTelemetryCollector and updates it. Returns the server's representation of the openTelemetryCollector, and an error, if there is any.
func (c *openTelemetryCollectors) Update(openTelemetryCollector *v1alpha1.OpenTelemetryCollector) (result *v1alpha1.OpenTelemetryCollector, err error) {
	result = &v1alpha1.OpenTelemetryCollector{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("opentelemetrycollectors").
		Name(openTelemetryCollector.Name).
		Body(openTelemetryCollector).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *openTelemetryCollectors) UpdateStatus(openTelemetryCollector *v1alpha1.OpenTelemetryCollector) (result *v1alpha1.OpenTelemetryCollector, err error) {
	result = &v1alpha1.OpenTelemetryCollector{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("opentelemetrycollectors").
		Name(openTelemetryCollector.Name).
		SubResource("status").
		Body(openTelemetryCollector).
		Do().
		Into(result)
	return
}

// Delete takes name of the openTelemetryCollector and deletes it. Returns an error if one occurs.
func (c *openTelemetryCollectors) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("opentelemetrycollectors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *openTelemetryCollectors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("opentelemetrycollectors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched openTelemetryCollector.
func (c *openTelemetryCollectors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpenTelemetryCollector, err error) {
	result = &v1alpha1.OpenTelemetryCollector{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("opentelemetrycollectors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
