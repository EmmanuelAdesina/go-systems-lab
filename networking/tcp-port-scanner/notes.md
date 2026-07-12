<div align="center">

```
┌────────────────────────────────────────────┐
│   D A Y   0 3   ::   N O T E S              │
│   R A W   E N G I N E E R I N G   L O G     │
└────────────────────────────────────────────┘
```

![Type](https://img.shields.io/badge/TYPE-RAW_NOTES-8A2BE2?style=for-the-badge&labelColor=0D1117)
![Mood](https://img.shields.io/badge/MOOD-CLICKED-39FF14?style=for-the-badge&labelColor=0D1117)

</div>

> These are the working notes I wrote *while* building, not a polished report. See [`README.md`](./README.md) in this folder for the clean writeup.

<br>

## 💡 What I Learned

Today I built my first program that communicates with another machine over the network.

The most important realization: **a TCP port scan is simply an attempt to establish a TCP connection.** There is no special "scan" operation hiding underneath. If the connection succeeds, the port accepted it. If it fails, the operating system reports why.

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 🧠 Mental Models

### `conn` — the telephone line

> I think of a connection as an open telephone line.

```
net.Dial()        →  places the call
someone answers   →  connection established
line stays open   →  both sides send / receive data
conn.Close()      →  hangs up the call
```

---

### `error` — context, not a verdict

An error does **not** necessarily mean the port is closed. It means my program could not successfully complete the operation it promised to perform.

```
Invalid hostname
Network unreachable
Timeout
DNS failure
```

The error provides *context* instead of hiding the reason for failure behind a plain "no."

---

### `defer` — an engineering solution, not a syntax trick

`defer` solves an engineering problem rather than a syntax problem.

Without it, I'd have to remember to call `conn.Close()` before **every** possible return statement — success path, error path, early return, all of them.

With `defer`, cleanup is scheduled once, and Go guarantees it runs when the function exits, however it exits.

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 🧭 Engineering Decisions

I chose:

```go
func ScanPort(host string, port int) (bool, error)
```

because the caller deserves both:

- the **result** of the scan
- the **reason** when the scan can't be completed

I also chose `net.JoinHostPort()` over `fmt.Sprintf()` because it expresses networking *intent* directly, and correctly handles IPv6 — a hand-formatted string doesn't.

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 🔁 Pattern I'm Starting to Notice

A recurring Go shape:

```go
resource, err := OpenSomething()
if err != nil {
    return err
}
defer resource.Close()
```

I expect to see this same skeleton repeatedly with:

- Files
- Network connections
- HTTP responses
- Database resources

Recognizing this pattern early feels like it'll save a lot of confusion later — it's less "new syntax to learn" and more "same shape, new resource."

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

</div>

## 🪞 Reflection

Today's project made Go feel much more like a systems programming language.

Instead of hiding networking behind abstractions, Go exposes simple building blocks that closely match how real systems actually work — a socket is a socket, a connection is a connection, and errors are just data instead of exceptions leaping out of nowhere.

I'm beginning to recognize that many Go features exist to solve **engineering problems**, not language problems — and understanding *why* a feature exists makes it far easier to learn than memorizing *how* to use it.

<div align="center">

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

**⬅ Back to** [`day-03 README`](./README.md) · [`Go Cloud Security Lab`](../README.md)

</div>
