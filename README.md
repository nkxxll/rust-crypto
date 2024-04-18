# crypto in rust study

## Diffie Hellman for Randsomware

```
 ┌─────┐                                             ┌───┐
 │Alice│                                             │Bob│
 └──┬──┘                                             └─┬─┘
    │            saves pub key in program              │
    │ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ >│
    │                                                  │
    │                                                  ────┐                             ╔═══════════════════════════════╗
    │                                                      │ generate pub priv key pair  ║priv key is deleted after use ░║
    │                                                  <───┘                             ╚═══════════════════════════════╝
    │                                                  │
    │                                                  ────┐
    │                                                      │ calc the secret with a.pub key and b.priv
    │                                                  <───┘
    │                                                  │
    │                  send pub key                    │
    │<─────────────────────────────────────────────────│
    │                                                  │
    ────┐                                              │
        │ calc secret send back after randsome is payed│
    <───┘                                              │
 ┌──┴──┐                                             ┌─┴─┐
 │Alice│                                             │Bob│
 └─────┘                                             └───┘
```
