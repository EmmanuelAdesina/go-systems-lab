Day 02 Notes

What I Learned

Today I learned that pointer receivers are not simply about memory addresses—they are about ownership and mutation.

A value receiver receives a copy of a struct, so any changes remain local to that copy.

A pointer receiver receives access to the original object, allowing changes to be visible everywhere the object is used.

That distinction made the purpose of pointers much clearer.

---

Mental Model

I now think about methods in two categories:

Read-only methods

These answer questions about an object without changing it.

Examples:

- "IsAdmin()"
- "HasHomeDirectory()"

These are good candidates for value receivers when copying the object is inexpensive.

Modifying methods

These change the object's internal state.

Examples:

- "ChangePort()"
- "EnableTLS()"

These require pointer receivers because the original object should be updated.

---

Mistakes I Made

- I initially kept a "DebugEnabled" field even though "LogLevel" already represented similar behavior.
- After removing the field, I forgot to update the struct literal and received a compiler error:
  - "unknown field DebugEnabled in struct literal of type Config"

The compiler immediately pointed out the inconsistency, which reinforced how helpful Go's type system is.

---

Engineering Lessons

Instead of asking:

«"Should I use a pointer?"»

I should ask:

«"Should this method modify the original object?"»

The answer naturally determines whether a pointer receiver is appropriate.

---

Reflection

Today's project felt closer to real engineering than a language exercise.

Configuration exists in almost every production system—from web servers to cloud platforms—and understanding how mutable configuration is modeled in Go makes pointer receivers feel practical instead of abstract.

I'm beginning to see that learning Go isn't about memorizing syntax; it's about understanding why the language is designed the way it is and how those design choices support building reliable systems.
