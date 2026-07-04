<div align="center">

```
┌────────────────────────────────────────────┐
│   D A Y   0 1   ::   S T R U C T S          │
│   S Y S T E M   U S E R   I N S P E C T O R │
└────────────────────────────────────────────┘
```

![Concept](https://img.shields.io/badge/CONCEPT-SYSTEM-ENGINEERING-00F0FF?style=for-the-badge&labelColor=0D1117)
![Status](https://img.shields.io/badge/STATUS-COMPLETE-39FF14?style=for-the-badge&labelColor=0D1117)
![Lang](https://img.shields.io/badge/GO-00ADD8?style=for-the-badge&logo=go&logoColor=000000&labelColor=0D1117)

</div>

<br>

## 🎯 Goal

Learn how to model real-world entities using Go **structs**, and attach behavior to them using **methods**.

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 🧩 Concepts Covered

```
Structs
Methods
Functions
Slices
Range loops
```

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 🛠️ What I Built

A simple Linux user model — `SystemUser` — that can answer questions about itself without any external logic reaching in to inspect its fields directly:

```go
type SystemUser struct {
    Username string
    UID      int
    IsAdmin  bool
    HomeDir  string
}

func (u SystemUser) IsAdministrator() bool {
    return u.IsAdmin
}

func (u SystemUser) HasHomeDirectory() bool {
    return u.HomeDir != ""
}
```

**Behavior it exposes:**
- ✅ Is the user an administrator?
- ✅ Does the user have a home directory?

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 🧠 Reflection

> Today I learned that methods aren't "inside" a struct — they're functions with a **receiver**, which lets behavior stay close to the data it belongs to without collapsing the two into one thing.

That distinction matters more than it looks: a struct is just a shape for data. A method is a function that happens to be scoped to that shape. Nothing is hidden or magic — it's still just a function call, `IsAdministrator(u)` in disguise. Understanding that early makes interfaces (later on the roadmap) click much faster, since interfaces are really just "does this type have the right methods," not "does this type belong to the right hierarchy."

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 📂 Files in This Folder

| File | Purpose |
|---|---|
| `main.go` | The working `SystemUser` model and its methods |
| `notes.md` | Raw notes taken while building — messier than this README, on purpose |
| `README.md` | This file — the polished writeup |

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

**⬅ Back to** [`Go Cloud Security Lab`](../README.md)

</div>

