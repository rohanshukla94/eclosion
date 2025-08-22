<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a>
</p>

<h3 align="center">Eclosion web framework written in Go</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center">This is my attempt to create a web framework in Golang.
    <br> 
    What I built: “Eclosion is a lightweight starter for Go web apps that standardizes sessions, rendering, routing, and env‑based configuration. It lets me switch between Jet and Go templates, and swap session backends (cookie/Redis/Postgres) based on environment.”

    Key design choices:

    Chi for composable middleware and lean router.

    SCS for secure session handling with pluggable stores.

    Jet for fast, expressive templating + optional fallback to stdlib.

    Reliability: Centralized logging, Recoverer middleware, configurable timeouts on the server.
</p>

### Todo

- [ ] finish session stores,
- [ ] add CSRF and Static Serving
- [ ] fixing config pitfalls
- [ ] small simple auth feature to demo sessions and templates


## 📝 Table of Contents

- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [TODO](../TODO.md)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)


## 🏁 Getting Started <a name = "getting_started"></a>

```
go get github.com/rohanshukla94/eclosion
```

### Prerequisites

What things you need to install the software and how to install them.

```
Give examples
```

### Installing

A step by step series of examples that tell you how to get a development env running.

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo.

## 🔧 Running the tests <a name = "tests"></a>

Explain how to run the automated tests for this system.

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## 🎈 Usage <a name="usage"></a>

Add notes about how to use the system.

## 🚀 Deployment <a name = "deployment"></a>

Add additional notes about how to deploy this on a live system.