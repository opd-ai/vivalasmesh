# Viva Las Mesh

> *"Las Vegas is the only town in the world where you can lose a fortune, get married in a pink Cadillac by a man dressed like Elvis, get shot in the desert by a guy named Frank, and wake up in a morgue in time for the 9:99 buffet."*

**Viva Las Mesh** is a terminal-native, multi-temporal, simulturn MUD and tactical heist simulator set across the past, present, and darknet future of the Las Vegas Valley. Built in Go using the **Charmbracelet** ecosystem (`Bubble Tea`, `Lip Gloss`, `Bubbles`), it runs entirely inside standard ANSI-compliant terminals using a high-density sub-pixel rendering engine (`▀`, `U+2580`).

Operating off-grid via Meshtastic LoRa radios, Tor, and I2P anonymity networks, the game bridges authentic real-world P2P mesh protocols with deep, uncompromising cyberpunk and syndicate gameplay.

---

## Key Features

* **Sub-Pixel Terminal Rendering (`▀`):** Uses upper half-block characters to achieve double vertical resolution ($120 \times 80$ effective pixels in a standard terminal), managed through an exhaustive 5-layer Z-buffer matrix (`Z-0` floor to `Z-4` HUD/spectrum analyzers).
* **Simulturn Pulse Engine:** Powered by 250ms micro-ticks and 1000ms macro-pulses. Actions stage continuously and resolve simultaneously using a strict 5-tier priority matrix (from Divine/Spectral King interventions down to heavy vault drilling).
* **Multi-Transport P2P Architecture (`orchestrator-server`):** Seamlessly syncs game state across offline Meshtastic LoRa meshes, Tor, and I2P garlicked tunnels using CRDT state convergence and Noise protocol framing.
* **Four Temporal Eras:** Explore authentic Las Vegas Valley locations across 1953 (Atomic AEC Era), 1962 (Syndicate/Stardust Era), 1993 (Corporate Mega-Resort Era), and 2026 (Cyber-Strip Dystopia).
* **Radio Free Graceland & Rat Pack Pantheon:** Tune into off-grid audio streams hosted by the eternal spirit of The King, while managing your luck against the Rat Pack Pantheon and avoiding "The Cooler" NPC.
* **Hardcore Survival & Escalation:** Manage your THC/distillate levels, hydration, police heat, and **Paranoia Index** (complete with "Bat Country" visual shaders). Fail a heist, and you face real-time desert trunk breakouts, grave digging, and the dreaded Morgue Wake State.
* **Frictionless Roguelike UX:** Zero execution barrier fatigue featuring context-aware smart keys (`Space` / `Enter`), ANSI SGR point-and-click mouse navigation, and fuzzy-search action palettes (`Ctrl+K`) with embedded success odds previews.

---

## Architectural Tech Stack

* **Language:** Go 1.24+
* **TUI Framework:** `[github.com/charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea)`, `[github.com/charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss)`, `[github.com/charmbracelet/bubbles](https://github.com/charmbracelet/bubbles)`
* **Networking & Transport:** `[github.com/go-i2p/onramp](https://github.com/go-i2p/onramp)` (Unified Tor/I2P tunneling), Meshtastic LoRa packet drivers (CBOR serialization via `[github.com/fxamacker/cbor/v2](https://github.com/fxamacker/cbor/v2)`)
* **State & Sync:** `[github.com/ipfs/go-datastore](https://github.com/ipfs/go-datastore)`, `[github.com/cbergoon/merkletree](https://github.com/cbergoon/merkletree)` (CRDT state convergence)
* **Crypto:** `golang.org/x/crypto/chacha20poly1305`, `golang.org/x/crypto/curve25519` (Noise IK / Double Ratchet)

---

## Getting Started

### Prerequisites

Ensure you have Go installed on your system along with an ANSI-compliant TrueColor terminal emulator (e.g., Kitty, Alacritty, WezTerm, or Windows Terminal).

### Installation & Execution

Clone the repository and run the background daemon and TUI client:

```bash
git clone https://github.com/opd-ai/toxcore.git
cd toxcore

# Initialize dependencies
go mod tidy

# Build and launch the orchestrator daemon & TUI client
go build -o ./bin/viva-las-mesh ./cmd/mesh
./bin/viva-las-mesh
```

---

## Controls & Ergonomics

| Key / Input | Action / Context |
| :--- | :--- |
| **`WASD` / Arrow Keys** | Grid movement and menu navigation |
| **`Space` / `Enter`** | **Contextual Smart Key:** Auto-resolves highest-priority action (e.g., loot cash, glitch vault lock, talk to NPC) |
| **`Ctrl+K` / `/`** | Open Fuzzy Action Palette for natural language command lookup |
| **`Tab`** | Cycle focus between Map View, Inventory, and RF Spectrum Analyzer |
| **`Esc`** | Cancel active prompt or close overlays |
| **Left Click / Tap** | Point-and-click $A^*$ pathfinding on the world grid |
| **Right Click / Hold** | Open context radial menu on entities or tiles |

---

## Contributing

Please review `copilot-instructions.md` before submitting pull requests to ensure strict adherence to our domain isolation boundaries, sub-pixel rendering rules, and third-party Go library reuse mandate.

---

## License

Distributed under the MIT License. See `LICENSE` for more information.