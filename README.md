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
- Composite (composite.go)
- Decorator (decorator.go)
- Facade (facade.go)
- Flyweight (flyweight.go)
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

#### Decorator
In a scenario where class responsibilities are removed or added, the decorator pattern is applied.  The decorator pattern helps with sublassing when modifying functionality, instead of static inheritance.  

The **single responsibility principle** can be achieved using a decorator.

##### Use Cases
-  *window components* 
-  *object modeling*

##### Components
- component interface
- concrete component class - implements the component interface
- decorator class - implements the component interface and extends functionality in the same method or additional ones.  The decorator base can be a particapant representing the base class for all decorators.

#### Facade
Used to abstract subsystem interfaces with a helper.  This design pattern gets used in scenarios when the number of interfaces increases and the system gets complicated.  Facade is an *entry point* to different subsystems, and it simplifies dependencies between the systems.  It provides an interface that hides the implementation details of the hidden code.

##### Use Cases
- Loose Coupling
- Improve poorly designed APIs
- Incorporate changes to the contract and implementation using a service facade in SOA

##### Components
- **Facade Class** - Delegates the requests from the client to the module classes; hides the complexities of the subsystem logic and rules
- **Module Classes** - Implement the behaviors and functionalites of the module subsystem.
- **Client** - Invokes the `facade` method.  The `facade` class functionality can be spread across multiple packages and assemblies.

#### Flyweight
Used to manage the state of an object with high variation.  The pattern allows us to share common parts of the object state among multiple objects, instead of each object storing it.  Variable object data is referred to as *extrinsic state*, and the rest of the object state is *intrinsic*.  Extrinsic data is passed to flyweight methods and will never be stored within it.  

Flyweight patterns help reduce the overall memory usage and the object initializing overhead.  The pattern helps create interclass relationships and lower memory to a manageable level.

Flyweight objects are immutable.  

##### Components
- **FlyWeight** interface - Has a method through which flyweights can get and act on the extrinsic state
- **ConcreteFlyWeight** - implements the `FlyWeight` interface to represent flyweight objects
- **FlyweightFactory** - used to create and manage flyweight objects.  The client invokes `FlyweightFactory` to get a flyweight object.  `UnsharedFlyWeight` can have a functionality that is not shared.
- Client classes