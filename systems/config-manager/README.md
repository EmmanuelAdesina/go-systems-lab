<div align="center">

```
┌────────────────────────────────────────────┐
│   D A Y   0 2   ::   M E T H O D S          │
│   C O N F I G   M A N A G E R               │
└────────────────────────────────────────────┘
```

![Concept](https://img.shields.io/badge/CONCEPT-POINTER_RECEIVERS-00F0FF?style=for-the-badge&labelColor=0D1117)
![Status](https://img.shields.io/badge/STATUS-COMPLETE-39FF14?style=for-the-badge&labelColor=0D1117)
![Lang](https://img.shields.io/badge/GO-00ADD8?style=for-the-badge&logo=go&logoColor=000000&labelColor=0D1117)

</div>

<br>

## 🎯 Objective

Learn how Go uses **pointer receivers** to modify the original object instead of operating on a copy.

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## ⚙️ Engineering Problem

Configuration sits at the core of almost every real system — web servers, APIs, Kubernetes controllers, cloud services, and automation tools all depend on configuration objects that can be read, updated, and trusted at runtime.

This project explores how configuration can be modeled in Go, and why **pointer receivers** are required the moment an object needs to update its own state instead of just describing it.

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 🧩 Concepts Covered

```
Structs
Methods
Pointer Receivers
Value Receivers
time.Duration
Composite Literals
```

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 🛠️ The Project

A simple configuration manager for a web service. The model captures the values almost every service needs at boot:

```go
type Config struct {
    Host       string
    Port       int
    TLSEnabled bool
    CertFile   string
    KeyFile    string
    Timeout    time.Duration
    LogLevel   string
}
```

**Values modeled:**
- Host
- Port
- TLS configuration
- Certificate paths
- Request timeout
- Log level

Configuration changes persist across the program because the mutating methods use **pointer receivers** — they operate on the real `Config`, not a throwaway copy of it.

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 🔧 Methods Implemented

```go
func (c *Config) ChangePort(port int)
func (c *Config) EnableTLS(certFile, keyFile string)
```

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 📤 Sample Output

```diff
=== Before changes ===
{Host:0.0.0.0 Port:8080 TLSEnabled:false CertFile: KeyFile: Timeout:30s LogLevel:info}

=== After changes ===
{Host:0.0.0.0 Port:8443 TLSEnabled:true CertFile:/etc/certs/server.crt KeyFile:/etc/certs/server.key Timeout:30s LogLevel:info}
```

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 🧠 Key Takeaways

- Methods with **value receivers** operate on a copy — changes vanish the moment the method returns.
- Methods with **pointer receivers** modify the original object — changes persist for every caller holding that pointer.
- Configuration objects are mutable by nature, which makes them a textbook case for pointer receivers — you're not describing config, you're *managing* it.

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## ⏭️ Next Step

The next project introduces interaction with the outside world — file I/O or networking — which naturally leads into Go's **error handling** model. Once a program touches the outside world, "it compiled" stops meaning "it works."

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

**⬅ Back to** [`Go Cloud Security Lab`](../README.md)

</div>
