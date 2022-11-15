# Arcaflow deployer interface

This library holds the deployer interface that allows for writing deployment systems for Arcaflow plugins.

New deployment systems have to implement the [`ConnectorFactory`](interface.go) as a primary entry point. As a type parameter, you will have to define your configuration structure, and your factory will have to return a schema that can be used to validate configuration options.