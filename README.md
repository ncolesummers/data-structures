# Data Structures in Golang
Fundamentals are important, if not always riveting.  This repository serves as a reference for common data structures.
--

## Basic Structures
- Lists (lists.go)
- Tuples (tuples.go)
- Heaps (heaps.go)
---
## Structural Design Patterns
Describe the relationships between the entities.  Used to build large structures using classes and objects.   

### Gang of Four (**GoF**) Structural design patterns
- Adapter (adapter.go)
- Bridge
- Composite
- Decorator
- Facade
- Flyweight
- Private Class Data
- Proxy 

#### Adapter

Provides a wrapper with an interface required by the API client to link incompatible types and act as a translator between the two types.

Composition includes the target, adaptee, adapter, and client:
- *Target* is the interface called by the client and invokes methods on the adapter and adaptee.
- The *client* wants the incompatible interface implemented by the adapter
- The *adapter* translates the incompatible interface of the adaptee into an interface that the client wants.