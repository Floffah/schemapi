# SchemAPI Language Syntax Guide

Welcome to the SchemAPI language documentation! This guide covers the syntax and structure of the SchemAPI schema language, designed to describe APIs, RPC, WebSockets, streams, and custom protocols in a concise, extensible, and human-friendly way. The syntax is inspired by HashiCorp Configuration Language (HCL) and aims to be both readable and powerful.

## Table of Contents
- [Overview](#overview)
- [Callables](#callables)
- [Paths](#paths)
- [Parameters](#parameters)
- [Blocks and Nesting](#blocks-and-nesting)
- [Types](#types)
- [Strings](#strings)
- [Comments](#comments)
- [Extensibility](#extensibility)
- [Examples](#examples)

---

## Overview
SchemAPI schemas are composed of **callables** (such as HTTP endpoints, RPC methods, etc.), each with a type, identifier (often a path), and optional parameters and nested blocks. The language is whitespace-insensitive and uses braces `{}` for grouping.

## Callables
A **callable** is a top-level block that defines an API operation or endpoint. The syntax is:

```
callable <type> <identifier> {
    // ...parameters and nested blocks...
}
```

- `<type>`: The callable type, e.g., `http.get`, `rpc.method`, etc. This is extensible and not hardcoded.
- `<identifier>`: The name or path of the callable (e.g., `/healthz`, `/users/:id`).

### Example
```
callable http.get /healthz {
}

callable http.get /users/:id/dashboard {
}
```

## Paths
Paths are used as identifiers for HTTP callables. They support static and parameterized segments:
- Static: `/healthz`
- Parameterized: `/users/:id/dashboard`

Path parameters are denoted by a colon prefix (e.g., `:id`).

## Parameters
Parameters are defined inside callable blocks. The syntax for parameters is:

```
param <name> <type>
```

- `<name>`: The parameter name (e.g., `userId`)
- `<type>`: The parameter type (e.g., `string`, `int`, `url`)

Example:
```
callable http.get /users/:id {
    param id string
}
```

## Blocks and Nesting
Blocks are delimited by `{}`. You can nest blocks for more complex structures (e.g., responses, request bodies, etc.).

Example:
```
callable http.post /users {
    param name string
    param email string
    response {
        param id string
        param createdAt string
    }
}
```

## Types
Supported types include:
- `string`
- `int`
- `float`
- `bool`
- `url` (special type, not just a string)

Types are extensible and can be expanded by plugins or future language versions.

## Strings
Strings can be enclosed in either double quotes `"` or single quotes `'`:
```
param name "John Doe"
param description 'A user account'
```

## Comments
Currently, comments are not supported in the syntax. (Planned for future versions.)

## Extensibility
- **Callable types** (e.g., `http.get`, `rpc.method`) are not hardcoded and can be extended by third-party plugins.
- **Types** can also be extended.

## Examples
### Minimal HTTP GET
```
callable http.get /healthz {
}
```

### HTTP GET with Path Parameter
```
callable http.get /users/:id {
    param id string
}
```

### HTTP POST with Body and Response
```
callable http.post /users {
    param name string
    param email string
    response {
        param id string
        param createdAt string
    }
}
```

---

For more advanced features and updates, see the [project repository](https://github.com/floffah/schemapi).

