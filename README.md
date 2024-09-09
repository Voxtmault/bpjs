# BPJS Health Insurance Module

A Sub-Project (or Modules if you'd like) to connect with BPJS Health Insurance of Indonesia.

## Folder Structure

- [config](./config/README.md): Where you will find many configurations regarding how the service will behave.
- [docs](./docs/README.md): Where you will find OpenAPI Standards Documentation regarding the service.
- [pkg](./pkg/README.md): Where you will find the inner parts of this service.

## Getting Started

To get started with this project, follow these steps:

Clone this GitHub repository to your local machine.

```bash
git clone github.com/voxtmault/bpjs-rs-module
```

Or you can make a fork from this repo if you'd like.

Run

```bash
go get
```

To download all of the necessary packages.

Once you have done that, you can create your own .env files, just by renaming the .env.example into .env (or you can copy-paste it, do what you like)

## Security Guidelines

This is kinda obvious but, **NEVER EVER**, i repeat **NEVER** push the configured .env files into github.
The main reason is because you will need to put confidential API Keys and Secrets into the .env file, you can refer to [.env.example](.env.example) for examples.

I know who you are :D, please don't make this even harder than it is already
