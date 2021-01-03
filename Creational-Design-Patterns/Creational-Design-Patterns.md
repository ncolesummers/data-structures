# Creational Design Patterns

Creational design patterns deal with object creation mechanisms in a safe and efficient manner and decouple clients from implementation specifics.  With these patterns, the code using an object need not know details about how the object is created, or even the specific type of the object,
as long as the object adheres to the interface expected.

## Factory Method (factory.go)
A factory is an object that is used to create other objects.  In a factory method pattern, a helper method or function is defined, to enable object creation without knowing the implementation class details.  By convention, the name of a factory method will begin with "New."

## Builder (builder.go)

The Builder design pattern is useful when object creation is not straightforward.  Its useful when:

- You need to validate some parameters or derive added attributes
- Optimization is important, such as retreiving an object from a cache rather than reading from the DB
- Idempotency and thread safety are required in object creation
- The objects might have contructor arguments (typically called **telescopic constructors**), and it is difficult to remember the order of parameters for the clients.  Some of these parameters might be optional.  Telescopic constructors frequently lead to bugs in client code.

## Abstract Factory (abstract-factory.go)
An abstract factory is useful when you have related objects that need to be created together.  When client code is creating related products, how do we ensure that mistakes are not made?  This creational design pattern attempts to solve the issue with a *factory of factories* construct: a factory that groups the different related/dependent factories together without specifying their concrete classes.

## Singleton (singleton.go)
This design pattern restricts the creation of objects to a single one.  This might be useful, for example, when you want a single *coordinator* object across in multiple places of the code.

