<div align="center">

# 🛡️ Go Cloud Security Lab

### Learning Go in public — by building real, working systems.

*No isolated exercises. No throwaway tutorials. Just projects that mirror how Go is actually used in production.*

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/status-actively%20building-brightgreen?style=for-the-badge)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)
![Last Commit](https://img.shields.io/badge/dynamic-learning%20log-orange?style=for-the-badge)

</div>

---

## 🎯 What This Is

This repository is a **public engineering log** — a running record of how I'm learning Go by building things that look and behave like real infrastructure, not classroom exercises.

The goal isn't to "finish learning Go." Languages don't get finished — systems do, and then they get maintained, broken, and rebuilt. So instead of chasing syntax for its own sake, every concept here gets tied to something that has to actually *work*: parse input it didn't design for, fail predictably, and be readable by someone who isn't me.

I'm focused on the areas where Go is doing real load-bearing work in the industry:

| Domain | Why It Matters |
|---|---|
| ☁️ **Cloud Security** | Go powers the tooling that secures the infrastructure everything else runs on |
| 🔐 **DevSecOps** | Fast, static-binary tooling that fits cleanly into CI/CD pipelines |
| 🤖 **AI Agents & Automation** | Concurrency and simplicity make Go a strong fit for agentic tooling |
| 📦 **Backend Services** | The language of choice for scalable, low-latency APIs |
| 🐳 **Docker & Kubernetes** | Both are written in Go — understanding the language demystifies the ecosystem |
| 🖥️ **CLI Tools** | Go's standard library makes production-grade CLIs approachable |
| 🌐 **Distributed Systems** | Where Go's concurrency model actually earns its keep |

---

## 🧭 Philosophy

Every project in this repo follows the same loop:

```
1. Pick one Go concept.
2. Build one meaningful project around it.
3. Document what broke, what I learned, and why it matters.
4. Push it publicly — imperfect, but real.
```

I'm optimizing for **honest progress over polished appearances**. A working project with rough edges and clear notes teaches more — to me and to anyone reading this — than a clean tutorial copy-paste ever could.

---

## 🗺️ Learning Path

> **This roadmap is intentionally a living document, not a fixed curriculum.**
> I add, reorder, and revisit topics as real projects demand them — because that's how learning actually works outside of a classroom. Treat the list below as "where things stand today," not a locked syllabus.

**Currently exploring / done:**
- [x] Structs

**Up next (order subject to change):**
- [ ] Methods
- [ ] Pointers
- [ ] Interfaces
- [ ] Error Handling
- [ ] Packages
- [ ] Context
- [ ] Concurrency
- [ ] Testing
- [ ] Building Production CLI Tools

New topics get added the moment a project needs them — this list will keep growing.

---

## 📁 Repository Structure

Each day/topic lives in its own self-contained folder:

```
day-01-structs/
├── main.go          # the actual project code
├── notes.md         # what I learned, what confused me, what I'd do differently
└── README.md        # project-specific context

day-02-methods/
day-03-pointers/
...
```

No folder depends on another — you can drop into any day and understand it on its own.

---

## 🚀 Why Follow Along

If you're learning Go yourself, evaluating my work, or just enjoy watching engineering skill compound in public — this repo gives you:

- **Unfiltered progress** — including the mistakes, not just the wins
- **Real project code**, not snippets copied from docs
- **Security and infrastructure framing** for every concept, not abstract theory
- A roadmap that evolves the way real skill-building does

---

<div align="center">

**⭐ Star this repo if you want to watch it grow.**
*New commits, new concepts, new projects — added consistently.*

</div>
