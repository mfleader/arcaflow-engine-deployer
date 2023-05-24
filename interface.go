// Package deployer provides interfaces for all deployers of plugins to follow.
package deployer

import (
	"context"
	"io"

	log "go.arcalot.io/log/v2"
	"go.flow.arcalot.io/pluginsdk/schema"
)

// ConnectorFactory is an abstraction that hides away the complexity of instantiating a Connector. Its main purpose is
// to provide the configuration schema for the connector and then create an instance of said connector.
type ConnectorFactory[ConfigType any] interface {
	ID() string
	ConfigurationSchema() *schema.TypedScopeSchema[ConfigType]
	Create(config ConfigType, logger log.Logger) (Connector, error)
}

// AnyConnectorFactory is the untyped version of ConnectorFactory.
type AnyConnectorFactory interface {
	ID() string
	ConfigurationSchema() schema.Object
	Create(config any, logger log.Logger) (Connector, error)
}

// Connector is responsible for deploying a container image on the specified target. Once deployed and ready, the
// connector returns an I/O to communicate with the plugin.
type Connector interface {
	// Deploy instructs the connector to aquire the plugin and run it,
	// resulting in the plugin starting its ATP server.
	// The ATP server will be accessible through the `Plugin` interface.
	//
	// Parameters:
	// ctx: The context that lasts the length of the deployment.
	// 	     Cancelling the context will send a SIGTERM request to terminate the deployment,
	//       which can be used to cancel the running step.
	// image: The tag of the image to run.
	Deploy(ctx context.Context, image string) (Plugin, error)
}

// Plugin is single, possibly containerized instance of a plugin. When read from, this interface provides the stdout of
// the plugin, supposedly talking the Arcaflow Transport Protocol, whereas the writer will write to the standard input
// of the plugin. The Close() method will shut the plugin down and return when the shutdown was successful.
type Plugin interface {
	io.Reader
	io.Writer
	io.Closer
}
