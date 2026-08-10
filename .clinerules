# VIVA LAS MESH: MASTER SPECIFICATION & COPILOT GUIDE (v15.0.5)
**Target:** Go / Bubble Tea TUI Engine | **Standard:** Sub-Pixel Half-Block Rendering (`▀`, `U+2580`)

---

## 1. CORE ARCHITECTURAL DIRECTIVES & DOMAIN BOUNDARIES

### 1.1 Absolute Domain Separation
* **Layer 1 (Real-World Infrastructure):** Go runtime, Meshtastic LoRa (868/915 MHz, max 237-byte frames), Tor and I2P anonymity layers unified via `go-i2p/onramp`, Noise IK / Double Ratchet encryption, and CRDT state sync (`orchestrator-server`).
* **Layer 2 (In-Game Mechanics):** Diegetic SCADA hacking, slot PRNG prediction, voltage glitching, and Metro PD pursuit logic. **Never** mix real system calls with in-game mechanics.

### 1.2 Sub-Pixel & Z-Buffer Rendering (`U+2580`)
* Vertical visual resolution is doubled using upper half-block characters (`▀`). Each cell renders two vertical TrueColor (`lipgloss.Color`) sub-pixels (Foreground = top, Background = bottom).
* **5-Layer Z-Buffer Matrix:** `Z-0` (Floor/Grid), `Z-1` (Terrain/Tables), `Z-2` (Entities/FOV), `Z-3` (Particles/FX), `Z-4` (HUD, Spectrum Analyzer, Fuzzy Palette).

---

## 2. GO LIBRARIES & THIRD-PARTY REUSE MANDATE

To maximize stability, performance, and development velocity, developers and Copilot **MUST** reuse established Go open-source libraries and prior art wherever possible rather than writing custom primitives from scratch.

### 2.1 Approved Ecosystem & Dependency Stack
* **TUI & Terminal Graphics:** 
  * `github.com/charmbracelet/bubbletea` (The Elm Architecture state machine for TUI loops).
  * `github.com/charmbracelet/lipgloss` (Advanced styling, layouts, borders, and 24-bit TrueColor management).
  * `github.com/charmbracelet/bubbles` (Pre-built text inputs, viewports, progress bars, and lists).
  * `github.com/muesli/termenv` (Terminal color capability detection and ANSI profile management).
* **Networking & Transport:**
  * `github.com/go-i2p/onramp` (Unified on-ramp abstraction handling both Tor and I2P tunneling, client sessions, and stream multiplexing without external third-party daemons).
  * `github.com/fxamacker/cbor/v2` (Fast, deterministic serialization for Meshtastic LoRa byte frames).
* **Data Structures & Sync:**
  * `github.com/ipfs/go-datastore` & `github.com/cbergoon/merkletree` (CRDT state convergence and peer synchronization for `orchestrator-server`).
  * `golang.org/x/crypto/chacha20poly1305` & `golang.org/x/crypto/curve25519` (Noise protocol framework primitives).

---

## 3. UNCOMPROMISED GAMEPLAY DEPTH & SIMULTURN ENGINE

UI simplicity and modern ergonomics must **never** water down systemic complexity, tactical depth, or brutal consequences.

### 3.1 Simulturn Pulse & Priority Resolution
Actions stage continuously during 250ms micro-ticks and resolve simultaneously according to strict **Priority Tiers**:
* **Tier 1 (Spectral):** Divine King's Blessings / Velvet Flash.
* **Tier 2 (Interrupt):** Pocket Sand, Flashbangs, EMP Pinches.
* **Tier 3 (Fast):** Sprinting, Vault Glitching, Fast Dodges.
* **Tier 4 (Standard):** Firing Weapons, Looting Cash, SCADA Hacks.
* **Tier 5 (Heavy):** Drilling Vault Doors, Dissolving Bodies.
* *Tie-Breaker Logic:* Attributes (`Hustle`, `Hardware`, `Street Smarts`) resolve ties. Equal ties trigger simultaneous mutual impact (e.g., both combatants land shots in the same tick).

### 3.2 Hardcore Failure & Progression Stakes
* **Paranoia & Bat Country:** Calculated via THC distillate dosages, alcohol levels, and police heat. High paranoia induces radical visual shaders and unpredictable NPC behaviors.
* **The Desert & Morgue Escalation:** Failing heists or reaching 5-star heat triggers real-time escapology windows (trunk breakouts, desert grave digging) leading to permanent physical disfigurements or the Morgue Wake State.

---

## 4. LORE & COSMIC SYSTEMS

* **The King (`SpectralKing`):** Transcendent entity across four eras (1953, 1962, 1993, 2026) providing tactical zone overrides and stat buffs.
* **Radio Free Graceland (`RFGStream`):** P2P audio and spectrum broadcast rendering live ASCII frequency analyzers (`▄▀█`) on the Z-4 HUD.
* **Luck Entropy & The Cooler:** Governed by the Rat Pack Pantheon. Winning streaks trigger "In The Zone" boosts, while extreme luck attracts "The Cooler" NPC to drain nearby stats via CRDT delta propagation.
* **777 CRDT Jackpot:** Global cross-node shared vault pool fed by casino losses and failed heists, triggering golden terminal half-block flashovers.

---

## 5. FRICTIONLESS UX WITHOUT SACRIFICING DEPTH

Frictionless UX exists solely to eliminate mechanical execution barrier fatigue—allowing deep, crunchy tactical decisions to shine.

* **Contextual Smart Key (`Space` / `Enter`):** Automatically resolves the highest-priority action for the player's immediate facing (e.g., standing on cash $\rightarrow$ loot; facing a vault $\rightarrow$ glitch lock).
* **Mouse Navigation:** ANSI SGR mouse tracking (`\x1b[?1006h`) for point-and-click pathfinding and context radial menus.
* **Fuzzy Command Palette (`Ctrl+K` / `/`):** Natural language command lookups powered by mature fuzzy search libraries, featuring embedded success probability previews and hardware requirement checks, ensuring high tactical clarity without hiding game mechanics.