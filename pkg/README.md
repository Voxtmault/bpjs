# /pkg

This folder contains reusable packages and libraries that are used within the project.

## Structure

The `/pkg` folder is organized into subdirectories based on the functionality or purpose of each package. Each subdirectory represents a separate package and contains its own `README.md` file to provide detailed information about the package.

- [interfaces](./interfaces/README.md): Interfaces define a set of method signatures that a type (in this project named service) must implement, enabling polymorphism and decoupling code from specific implementations.
- [models](./models/README.md): Models hold definition of entities (objects / tables in a DB), the entities can have different varieties. There will be one for general purpose and 1 for creation purpose at least.
- [rpc](./rpc/README.md): You can think rpc folder as controllers in MVC Model Project Structure. It's where you will call the interfaces function and chose which service to use for each implementations. In this you will need to do input validations before calling the main logic.
- [services](./services/README.md): The services directory contains the core logic and implementations of the interfaces defined in the interfaces directory, providing the actual functionality that the project relies on..

## Usage

To use a package from the `/pkg` folder, import it into your Go code using the package's import path. Make sure to follow the package's documentation and guidelines for proper usage.
