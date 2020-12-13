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
- Bridge (bridge.go)
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

#### Bridge
Bridge design patterns decouple implementation from the abstraction.  The abstact base class can be subclassed to provide different implementations.

With bridge design patterns implementation details can change at runtime.

Follows the principle of composition over inheritance.  Helps in sutiations where one should subclass multiple times orthogonal to each other.  Runtime binding of the application, mapping of orthogonal class hierarchies, and platform independence implementation are the scenarios where the bridge pattern can be applied.

##### Use-Cases and Examples
Abstraction Layer to make your application work with different databases or message brokers
Flutter Embedded Layer maybe?


##### Components:
 **Abstraction** - Interface implemented as an abstract class that clients invoke with the method on the **concrete implementer**.  Maintains a *has-a* relationship with the implementation, instead of an *is-a* relationship.  The relationship is maintained by composition.  Abstraction has a *reference* (pointer) of the implementation.  **Refined abstraction** provides more variations that abstraction.

#### Composite
A group of similar objects in a single object.  Objects are stored in a tree to persist the entire hierarchy.  This pattern can be used to change a heirarchical collection of objects.  Modeled on a *heterogeneous collection*.  A benefit is that new types of objects can be added without changing the interface and client code.

##### Use-Cases
- UI Layouts for the Web or other platforms
- Network Device Tree in a monitoring dashboard
- Managing employees across departments

##### Components:
- **Component interface** defines the default behavior of all objects and behaviors for accessing the components of the composite.
- The *composite* and *component* classes implement the component interface.
- The *client* interacts with the component interface to invoke methods in the composite.