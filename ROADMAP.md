# VIVA LAS MESH: UNIFIED MASTER REAL-TIME MUD & SIMULTURN SYSTEM SPECIFICATION
**Document Version:** 15.0.0-FINAL (EXHAUSTIVE NO-CODE ARCHITECTURAL SPECIFICATION)  
**Target Platform:** Terminal-Native Sub-Pixel Engine (Go / Bubble Tea / Lip Gloss Architecture)  
**Architectural Standard:** Pulse-Driven MUD Engine / Table-Driven Simulturn Resolution Matrix  

---

## ARCHITECTURAL OVERVIEW: THE VEGAS SIMULATION PARADIGM

> *"Las Vegas is the only town in the world where you can lose a fortune, get married in a pink Cadillac by a man dressed like Elvis, get shot in the desert by a guy named Frank, and wake up in a morgue in time for the 9:99 buffet."* — **Vegas Lore Codex**

**Viva Las Mesh** is a hyper-authentic, terminal-native, multi-temporal MUD engine set across the neon, dust, and darknet of the Las Vegas Valley. Built on a strict **Table-Driven Architecture**, the engine decouples world data from mechanics using real-time pulse ticks, simultaneous order staging (*Apogee/Fantasy Flight* simulturn resolution), multi-transport mesh P2P networking (Meshtastic LoRa, Tor v3, I2P), and an accessible, modern terminal user experience.

---

## SECTION 1: STRICT ARCHITECTURAL BOUNDARY & ENGINE PIPELINE

```
+---------------------------------------------------------------------------------------+
|                       ABSOLUTE ARCHITECTURAL ISOLATION WALL                           |
+---------------------------------------------------------------------------------------+
|  [ LAYER 1: REAL-WORLD MULTI-TRANSPORT ENGINE P2P INFRASTRUCTURE ]                    |
|  - Software Stack: Go Runtime / Bubble Tea TUI / Lip Gloss Sub-Pixel Engine           |
|  - Transport Layer A (Off-Grid Radio): Meshtastic LoRa Packet Mesh (868/915 MHz)      |
|  - Transport Layer B (Anonymized Darknet): Tor v3 Hidden Services (.onion)            |
|  - Transport Layer C (Garlic Routed Darknet): I2P SAM Bridge / I2P Tunnels            |
|  - Transport Layer D (Proximity Radio): Bluetooth Low Energy (BLE) Mesh               |
|  - Crypto & State Transport: Noise IK Protocol / Signal Double Ratchet + CRDT Sync    |
|  - Daemon Control: orchestrator-server background daemon managing state replication   |
|                                                                                       |
|============================ STRICT DOMAIN SEPARATION ==================================|
|                                                                                       |
|  [ LAYER 2: IN-GAME DIEGETIC UNCONSTRAINED CYBERNETIC & RF HACKING ]                  |
|  - Full-Spectrum Signals Intelligence: VLF to Millimeter Wave / Optical / Satellite   |
|  - Protocol & Software Exploits: Zero-day payloads, SCADA/PLC injects, PRNG cracking  |
|  - Hardware & Physical Implants: FPGA Micro-SAO badges, glitching, bus tapping        |
|  - Electronic Warfare: IMSI catching, rogue base stations, GPS spoofing, satellite tap|
+---------------------------------------------------------------------------------------+
```

### 1.1 Terminal Sub-Pixel Engine (`▀`) Architecture
The rendering layer operates inside ANSI-compliant terminal emulators using the upper half-block character (`▀`, `U+2580`) to achieve double vertical resolution ($120 \times 80$ effective pixels in a $120 \times 40$ character cell terminal).

* **Cell Mapping:** Each character cell renders two vertical color channels:
  * **Top Sub-Pixel:** Foreground Color ANSI Escape Sequence (`\x1b[38;2;R;G;Bm`)
  * **Bottom Sub-Pixel:** Background Color ANSI Escape Sequence (`\x1b[48;2;R;G;Bm`)

#### Table 1.1: Canvas Composition & Z-Buffer Layer Matrix
| Layer Depth | Layer Name | Cinematic Aesthetic Reference | Composite Rule | Primary Render Output |
| :---: | :--- | :--- | :--- | :--- |
| **Z-0** | **Floor & Grid** | *Tron (1982)* / Neon Grid | Base Background | CRT scanlines, asphalt textures, carpet patterns, era-bleed. |
| **Z-1** | **Terrain Details** | *Casino (1995)* / Velvet Booths | Overwrites Z-0 | Craps tables, slot reels, vault walls, desert sand dunes. |
| **Z-2** | **Entities & FOV** | *Blade Runner (1982)* / Rainy Streets | Overwrites Z-1 | Player avatar, Mob enforcers, Metro PD, King's Presence. |
| **Z-3** | **Particle & FX** | *The Matrix (1999)* / Digital Rain | Additive Blending | RF signal waves, muzzle flashes, cigarette smoke, neon glow. |
| **Z-4** | **HUD & Terminal** | *Sneakers (1992)* / Cray Mainframe | Topmost Overlay | Health, Buzz, Paranoia, Inventory, Command Log, Spectrum. |

### 1.2 Multi-Transport P2P Engine Network Architecture

```
+---------------------------------------------------------------------------------------+
|                    REAL-WORLD ENGINE MULTI-TRANSPORT ARCHITECTURE                     |
+---------------------------------------------------------------------------------------+
|                               [Terminal Client Machine]                               |
|                               +-----------------------+                               |
|                               | Bubble Tea TUI Client |                               |
|                               +-----------+-----------+                               |
|                                           | Local IPC                                 |
|                               +-----------v-----------+                               |
|                               |  orchestrator-server  |                               |
|                               |   (State Sync / CRDT) |                               |
|                               +-----/-----+-----\-----+                               |
|                                    /      |      \                                    |
|                                   /       |       \                                   |
|      +---------------------------+    +---+---+    +---------------------------+      |
|      |  Meshtastic Radio Driver  |    |  Tor  |    |     I2P SAM Bridge        |      |
|      | (LoRa 868/915MHz Off-Grid) |    | (v3)  |    | (I2P Garlic Destination)  |      |
|      +-------------+-------------+    +---+---+    +-------------+-------------+      |
|                    |                      |                      |                    |
|             Physical Airwaves         Onion Circuits        Garlic Tunnels            |
|                    |                      |                      |                    |
|      +-------------v-------------+    +---v---+    +-------------v-------------+      |
|      | Remote Meshtastic Node    |    | Remote|    | Remote I2P Peer           |      |
|      | (Physical Long-Range Mesh)|    | Onion |    | (Garlic Destination)      |      |
|      +---------------------------+    +-------+    +---------------------------+      |
+---------------------------------------------------------------------------------------+
```

#### Table 1.2: Expanded Multi-Transport Routing & Protocol Matrix
| Transport ID | Physical / Logical Bearer | Film Homage / Lore Codex | Max Frame Size | Cryptographic Protocol | Network Routing Strategy |
| :---: | :--- | :--- | :---: | :--- | :--- |
| **TR-LORA** | Meshtastic (868/915 MHz) | *The Hunt for Red October* | 237 Bytes | Noise IK + Double Ratchet | Offline, direct physical RF hop mesh; lossy CRDT state updates. |
| **TR-TOR3** | Tor v3 Ephemeral Onion | *Hackers* ("Zero Cool") | 4096 Bytes | Ephemeral v3 Hidden Keys | Encrypted TCP streams; IP hiding over public internet. |
| **TR-I2P** | I2P SAM Garlic Tunnel | *Sneakers* ("Setec Astronomy")| 1024 Bytes | ElGamal/AES + SessionTags | Multi-hop garlic routing; resistant to deep packet analysis. |
| **TR-BLE** | Local Bluetooth LE Mesh | *Enemy of the State* | 512 Bytes | AES-256-GCM Direct | Short-range proximity peer handshake for local physical traders. |

$$\text{CRDT State Convergence:} \quad S_{\text{final}} = S_A \sqcup S_B \quad \text{where } \sqcup \text{ is associative, commutative, and idempotent}$$

---

## SECTION 2: MUD-TICK PULSE HIERARCHY & SIMULTURN RESOLUTION MATRIX

The game world ticks continuously in real-time. Action execution is staged during pulse windows and resolved simultaneously according to explicit priority rules.

### 2.1 MUD Heartbeat & Pulse Hierarchy

#### Table 2.1: MUD Pulse Engine Clock Division Matrix
| Tick Level | Interval | Cinematic Homage | System Operations & Engine Scope |
| :---: | :---: | :--- | :--- |
| **Micro-Tick** | **250 ms** | *The Matrix* ("Bullet Time") | Player movement interpolation, continuous sub-pixel rendering, real-time dodge/interrupt staging. |
| **Macro-Pulse**| **1000 ms** | *WarGames* ("3, 2, 1... launch") | MUD room ticks, guard patrol pathfinding, SCADA countdowns, wireless mesh packet dispatch. |
| **Meta-Pulse** | **10.0 sec**| *Casino* ("The Skim Count") | Drug/THC metabolic decay, Hydration depletion, police backup response timers, mob debt interest. |
| **Eon-Pulse**  | **60.0 sec**| *Viva Las Vegas* ("The King Awakens") | Spectral apparition check for The King across temporal nodes; global casino jackpot drift. |

### 2.2 Simulturn Order Staging & Conflict Resolution

```
+-------------------------------------------------------------------------------+
|                       SIMULTURN CONFLICT RESOLUTION FLOW                      |
+-------------------------------------------------------------------------------+
|  STAGED ACTIONS (Tick N)  --> Sort by Priority Tier (T1 to T5)                |
|                                         |                                     |
|  EQUAL PRIORITY TIER?    ----> YES ---> Evaluate Attribute Tie-Breaker        |
|                                         | (Hustle / Hardware / Street Smarts) |
|                                         v                                     |
|                                 STILL EQUAL? (Simultaneous Impact Executed)   |
|                                 - Both shots land simultaneously              |
|                                 - Both players grab the item at once          |
+-------------------------------------------------------------------------------+
```

#### Table 2.2: Action Priority Tier & Simulturn Resolution Table
| Priority Tier | Action Type Examples | Cinematic Trope | Priority Rule / Resolution Logic |
| :---: | :--- | :--- | :--- |
| **Tier 1: Divine / Spectral**| King's Blessing, Velvet Flash | *Viva Las Vegas* | Transcends physical reality; overrides all physical interrupts and weapons. |
| **Tier 2: Interrupt** | Pocket Sand, Flashbang, EMP Pinch | *No Country for Old Men* | Pre-empts Tier 3-5 actions in the same tick; cancels target's queued action. |
| **Tier 3: Fast Action** | Sprint, Dodge, Hardware Bypass | *Quick and the Dead* | Resolves before Standard actions; uses **Hustle** stat as tie-breaker. |
| **Tier 4: Standard** | Fire Weapon, Loot Cash, Hack SCADA | *Heat* | Standard simultaneous execution; both sides deal damage if fired in same tick. |
| **Tier 5: Heavy** | Drill Vault Door, Dissolve Body | *Snatch* | Requires N continuous uninterrupted ticks; canceled if hit by Tier 2 interrupt. |

---

## SECTION 3: TEMPORAL ERA MATRIX, AUTHENTIC MAP & THE KING'S TRANSFORMATION

The game world spans four historical eras across authentic Las Vegas Valley geography. **The King (Elvis)** exists as an eternal, spectral entity who transcends all temporal boundaries.

```
+-------------------------------------------------------------------------------+
|                       CINEMATIC TEMPORAL SPATIAL MAP                          |
+-------------------------------------------------------------------------------+
| NORTH (Era 1: 1953 Atomic)          EAST (Era 2: 1962 Syndicate)            |
| - Nevada Test Site / Sky Room       - Fremont St. / Stardust / Count Room   |
| - Inspired by: *Bugsy* (1991)       - Inspired by: *Casino* (1995)          |
|-------------------------------------------------------------------------------|
| WEST (Era 4: 2026 Cyber-Strip)      SOUTH (Era 3: 1993 Corporate Resort)    |
| - Subterranean Flood Tunnels        - Mirage Mega-Vault / Keycard Locks       |
| - Inspired by: *Blade Runner 2049*  - Inspired by: *Ocean's Eleven* (2001)    |
+-------------------------------------------------------------------------------+
```

### 3.1 Temporal Eras & Cinematic Node Locations

#### Table 3.1: Exhaustive Temporal Era & Node Location Registry
| Era ID | Location Node | Film Homage & Lore Inspiration | Dominant Faction | Unique Sector Mechanic |
| :---: | :--- | :--- | :--- | :--- |
| **ERA-1** | **Nevada Test Site Bunkers** | *Atomic Cafe* / *Bugsy* | Atomic Energy Comm. | Radioactive fallout drift; requires radiation counter. |
| **ERA-1** | **Desert Inn Sky Room** | *The Aviator* (2004) | AEC / Howard Hughes | High-society infiltration; strict tuxedo dress code. |
| **ERA-2** | **Stardust Vault & Count Room**| *Casino* (1995) | The Chicago Outfit | Cash skim couriers; Mob guards shoot on sight. |
| **ERA-2** | **Jean Dry Lake Bed** | *No Country for Old Men* | Mob Enforcers | Shallow grave sites; trunk drop-offs and executions. |
| **ERA-3** | **Mirage Mega-Vault Atrium** | *Ocean's Eleven* (2001) | Corporate Casino Moguls| PTZ camera networks; silent laser alarm grids. |
| **ERA-3** | **Underground Service Tunnels**| *Hard Eight* (1996) | Casino Security / Union | Keycard access doors; employee badge masking. |
| **ERA-4** | **Subterranean Flood Channels**| *Blade Runner 2049* | Underground Hackers | Off-grid darknet relays; high-voltage water traps. |
| **ERA-4** | **Mega-Dispensary Complex** | *Fear and Loathing* | Spectrum Hackers / Cartels| THC distillate refineries; crypto payment kiosks. |

### 3.2 The Eternal Transcendent Spirit of The King (Elvis)

Elvis is not bound by mortality, history, or physical code. He shifts forms across eras, acting as an elusive spirit, mentor, or cosmic force.

```
+-------------------------------------------------------------------------------+
|                    THE TRANSCENDENT KING SPECTRAL MATRIX                      |
+-------------------------------------------------------------------------------+
|  ERA 1 (1953)  --> The Young Sun Records Prophet (Raw Rockabilly Energy)       |
|  ERA 2 (1962)  --> The Hollywood King (Gold Sequined Suit / Movie Star)       |
|  ERA 3 (1993)  --> The Vegas Jumpsuit Oracle (White Caped Vision in Neon)     |
|  ERA 4 (2026)  --> The Holographic Spectral Code-God (AI Consciousness)       |
+-------------------------------------------------------------------------------+
```

#### Table 3.2: Transcendent Manifestations of The King
| Era Active | Spectral Manifestation Form | Film / Culture Homage | Encounter Trigger | Spectral Blessing / Curse Effect |
| :---: | :--- | :--- | :--- | :--- |
| **ERA-1** | **The Sun Records Rockabilly**| *3000 Miles to Graceland* | Sing "Heartbreak Hotel" in Sky Room | **Prophetic Riff:** Grants $+4$ Hustle and permanent immunity to fear. |
| **ERA-2** | **The Hollywood Gold-Suit** | *Viva Las Vegas* (1964) | Win $10,000 at Stardust Craps Table | **Velvet Charm:** $+5$ Dignity; Mob Capos refuse to attack player. |
| **ERA-3** | **The White Jumpsuit Oracle** | *Honeymoon in Vegas* (1992)| Green-out Paranoia Index $\ge 15$ | **Neon Redemption:** Clears all Paranoia; grants 1-time Revive token. |
| **ERA-4** | **The Holographic Code-God** | *Blade Runner 2049* | Broadcast Elvis MIDI over LoRa Mesh | **Graceland Protocol:** Overrides all local casino SCADA systems for 30s. |

#### Table 3.3: Anachronism Penalty Lookup Matrix
| Item Tech Level | Era Present | Anachronism Flag | Cinematic NPC Reaction | Faction Standing Hit |
| :---: | :---: | :---: | :--- | :---: |
| **2026 Cyber** | **1953** | `FLAG_ALIEN_TECH` | AEC Agents lock down sector; suspect Soviet/UFO spy. | AEC Rep $-50$ |
| **2026 Cyber** | **1962** | `FLAG_WIRE_BUG` | Mob Capos suspect wiretap device; draw snubnose .38s. | Outfit Rep $-40$ |
| **1953 Vintage**| **2026** | `FLAG_ANTIQUE` | Pawn Brokers offer double cash; street hackers laugh. | No Hit ($+10$ Cred) |

---

## SECTION 4: UNCONSTRAINED DIEGETIC HACKING & ELECTRONIC WARFARE

In-game cybernetic hacking spans the entire electromagnetic spectrum, software zero-days, SCADA industrial systems, and hardware implants.

```
+-------------------------------------------------------------------------------+
|              UNCONSTRAINED IN-GAME ELECTROMAGNETIC ATTACK MATRIX              |
+-------------------------------------------------------------------------------+
| SPECTRUM BAND   | IN-GAME DOMAIN              | DIEGETIC ATTACK VECTOR        |
| 3 kHz - 30 kHz  | Submarine / Underground ELF | AEC Deep Bunker Intercept     |
| 3 MHz - 30 MHz  | Shortwave HF Skywave        | Cold War Numbers Station Hack |
| 150 - 450 MHz   | Analog VHF / P25 Trunking   | Metro PD Voice & Data Tap     |
| 800 - 950 MHz   | Cellular GSM / LTE / P25    | Rogue IMSI Catcher / Spoofing |
| 1.2 - 1.6 GHz   | GPS / GNSS Navigation       | Casino Cash Truck Spoofing    |
| 2.4 - 5.8 GHz   | Wi-Fi / Bluetooth / Microwave| SCADA / CCTV / ATM Intercept   |
| 12 - 18 GHz     | Ku-Band Satellite Downlink  | High-Roller Feed Hijacking    |
| Optical / IR    | Free-Space Laser & Fiber    | Vault Isolation Bypass        |
+-------------------------------------------------------------------------------+
```

### 4.1 Exhaustive In-Game Hardware & RF Attack Catalog

#### Table 4.1: Diegetic Hardware Exploit & RF Attack Catalog
| Payload ID | Target Infrastructure | Film Homage | Hardware Check | Active Duration | Real-Time MUD World Outcome |
| :---: | :--- | :--- | :---: | :---: | :--- |
| **PL-IMSI** | Cellular Towers | *Enemy of the State* | Hardware 14 | 30 Pulses | Intercept Mob Capo calls; streams coordinates to log. |
| **PL-SCADA**| Vault Environmental | *Heat* (1995) | Hardware 16 | 12 Pulses | Flushes halon gas; forces count room guards to exit. |
| **PL-PRNG** | Video Slot Machine | *Ocean's Thirteen* | Hardware 13 | 1 Pulse | Predicts reel state; guarantees $5,000 jackpot next spin. |
| **PL-PINCH**| Casino Substation | *Ocean's Eleven* | Hardware 18 | 5 Pulses | Triggers local EMP pulse; blinds cameras for 5 sec. |
| **PL-SAT**  | Ku-Band Satellite | *Sneakers* (1992) | Hardware 15 | 20 Pulses | Hijack executive suite feeds; extract blackmail data. |
| **PL-GLITCH**| Vault Controller | *Matrix Reloaded* | Hardware 17 | Instant | Voltage spike bypasses door bootloader instantly. |
| **PL-GPS**  | Cash Transport Truck | *The Italian Job* | Hardware 15 | 10 Pulses | Spoofs GNSS signals; reroutes truck into drain ambush. |
| **PL-LASER**| Fiber Optic Trunk | *Mission: Impossible*| Hardware 16 | 15 Pulses | Mirrors vault optical camera feeds to TUI display. |
| **PL-CAN**  | Casino Elevator Bus | *Speed* (1994) | Hardware 12 | Instant | Freezes elevator cars between floors; traps security. |
| **PL-BGP**  | Casino Data Center | *Live Free or Die Hard*| Hardware 19 | 60 Pulses | Reroutes online sports book bets to player crypto wallet. |

---

## SECTION 5: FACTION POLITICS, "THE SKIM" & SIMULTURN HEIST MATRIX

### 5.1 Faction Matrix & Relationships

#### Table 5.1: Faction Disposition & Standing Matrix
| Faction | Primary Representative | Cinematic Trope | Neutral State | Hostile Action | Ally Perk |
| :--- | :--- | :--- | :---: | :--- | :--- |
| **AEC** | General "Nuke" Groves | Cold War Secrecy | Ignore | Enter Restricted Bunkers | Access to bunker safehouses |
| **Outfit** | Sam "Ace" Rothstein | Suit & Tie Mobster | Tax 10% | Steal from Count Room | Unlocks Mob muscle bodyguards |
| **Corporate**| Terry Benedict | Ruthless CEO | Track CCTV | Wear Unapproved Clothing | High-Roller Suite access |
| **Hackers** | Zero Cool | Cyberpunk Renegade | Share Nodes | Broadcast Real Name | Free darknet node relay |
| **Metro PD** | Detective Vincent Hanna| Driven Lawman | Patrol | Any Felony Action | SWAT Backup (Bribed State) |

### 5.2 Real-Time Police Heat Escalation Matrix

> *"He's home... watching TV... eating a TV dinner... and you're out here in the rain waiting for him."* — **Heat (1995)**

#### Table 5.2: Real-Time Heat & Pursuit State Machine
| Heat Level | MUD Tick Response Rate | Cinematic Threat Level | Real-Time Enforcement Action |
| :---: | :---: | :--- | :--- |
| **0 Stars** | No Active Response | *The Big Lebowski* | Patrol units pass by on standard 10-pulse loop; ignoring player. |
| **1-2 Stars**| Patrol Tick every 4 Pulses| *French Connection* | Nearby uniformed officers change pathfinding to intercept player coordinates. |
| **3-4 Stars**| Pursuit Tick every 2 Pulses| *Ronin* | Unmarked vehicles attempt vehicular ramming; SWAT deploy tear gas. |
| **5 Stars** | Continuous Micro-Tick Pursuit | *Terminator 2* | Metro PD snipers lock on; deadly force authorized every 250ms tick. |

### 5.3 "The Skim" Money Laundering Pipeline

```
[Dirty Cash] --> [Casino Floor / Dispensary Concession] --> [Laundering Fee (-15% to -30%)] --> [Clean Cash]
```

#### Table 5.3: Expanded Money Laundering Pipeline Table
| Channel ID | Facility Type | Cinematic Reference | Fee % | Heat Penalty Risk | Clean Yield per $1,000 |
| :---: | :--- | :--- | :---: | :---: | :---: |
| **LN-CRAPS**| High-Stakes Craps Table | *Casino* (1995) | 20% | High (Pit Boss Eye) | $800 Clean Cash |
| **LN-DISP** | Mega-Dispensary Retail | *Jackie Brown* (1997) | 15% | Low (Legal Front) | $850 Clean Cash |
| **LN-PAWN** | Gold & Silver Pawn | *Uncut Gems* (2019) | 30% | Zero (Off-Grid) | $700 Clean Cash |
| **LN-CAGE** | Stardust Count Room | *The Godfather II* | 10% | Extreme (Mob Sit-down)| $900 Clean Cash |

### 5.4 Multi-Stage Simulturn Vault Heist Engine

```
+-------------------------------------------------------------------------------+
|                       SIMULTURN CASINO VAULT HEIST FSM                        |
+-------------------------------------------------------------------------------+
|  STAGE 1: Surveillance & Sensor Bypass (SCADA / Satellite / RF Pinch)         |
|                                     |                                         |
|  STAGE 2: Guard Neutralization (IMSI Wiretap / THC / Flashbang)               |
|                                     |                                         |
|  STAGE 3: Vault Door Unlock (Voltage Glitch / Laser Cut / Mechanical Pick)    |
|                                     |                                         |
|  STAGE 4: Cash Extraction & Escape (Weight Limit / Pursuit Resolution)        |
+-------------------------------------------------------------------------------+
```

#### Table 5.4: Simulturn Vault Heist Resolution Matrix
| Heist Phase | Player Staged Action | Guard / System Staged Action | Priority Tier | Simultaneous Outcome |
| :---: | :--- | :--- | :---: | :--- |
| **Phase 1** | Inject SCADA Payload | Sweep CCTV PTZ Camera | Tier 2 vs Tier 3 | Cameras blinded 1 tick before sweep detects player. |
| **Phase 2** | Throw Flashbang | Draw Sidearm | Tier 2 vs Tier 4 | Flashbang detonates; guard blinded; shot canceled. |
| **Phase 3** | Voltage Glitch Lock | Pull Manual Wall Alarm | Tier 3 vs Tier 4 | Vault door opens instantly before guard reaches alarm lever. |
| **Phase 4** | Loot Cash Bag ($50k) | Seal Security Vault Doors | Tier 4 vs Tier 4 | Player grabs $50k cash, but vault doors seal permanently. |

---

## SECTION 6: PHYSIOLOGY, LEGAL DISPENSARY CATALOG & "BAT COUNTRY" PARANOIA ENGINE

### 6.1 Player Attribute Definitions
* **Health (0–100):** Physical condition. Reaching 0 triggers the Morgue Wake State.
* **Buzz (0–10):** Intoxication level from alcohol or cannabinoids. Modifies Luck and Perception.
* **Hydration (0–100%):** Drains continuously during movement/time. Reaching 0% causes collapse.
* **Hustle (1–20):** Agility, movement speed, lockpicking precision, and trunk escape success.
* **Street Smarts (1–20):** Deception, lie detection, price negotiation, and paranoia mitigation.
* **Dignity (1–20):** Access to high-roller lounges, mob sit-down respect, and pit boss tolerance.
* **Hardware (1–20):** Spectrum tuning, zero-day writing, FPGA soldering, SCADA hacking.

### 6.2 Mega-Dispensary Chemical Catalog

#### Table 6.1: Expanded Dispensary Catalog & Chemical Matrix
| Item ID | Product Name | Film Reference | Buzz Δ | Hydration Δ | Special Effect State |
| :---: | :--- | :--- | :---: | :---: | :--- |
| **CH-DIST** | 1000mg THC Distillate | *Fear and Loathing* | $+5$ | $-35\%$ | Grants [Hyper-Focus]: $+4$ Hardware, $-4$ Hustle. |
| **CH-CART** | Live Resin Cartridge | *Pineapple Express* | $+2$ | $-15\%$ | Grants [Precision]: $+3$ Soldering / Hardware. |
| **CH-SODA** | Nano-THC Elixir | *The Big Lebowski* | $+3$ | $-10\%$ | Grants [Abide]: $+2$ Luck on Slot Machine Reels. |
| **CH-MEL**  | Melatonin Gummy | *Leaving Las Vegas* | $-4$ | $-5\%$ | Clears [Tripping]; restores composure for Mob sit-downs. |
| **CH-SHRM** | Psilocybin Microdose | *The Doors* (1991) | $+3$ | $-20\%$ | High-frequency aura vision; highlights hidden radio taps. |
| **CH-WAX**  | 99% THCA Diamond Wax | *Dazed and Confused* | $+6$ | $-40\%$ | Grants [Astral Walk]: $+5$ Stealth; $-5$ Movement Speed. |

### 6.3 Real-Time Paranoia Engine & "Bat Country" Shaders

$$\text{Current Paranoia} = \left( \frac{\text{THC (mg)}}{\text{Metabolic Ticks elapsed}} \right) + (\text{Alcohol Level} \times 2) + \text{Police Heat} - \text{Street Smarts}$$

#### Table 6.2: Paranoia Index & Visual Shader Lookup
| Paranoia Index | State Name | Cinematic Reference | Visual Shader Effect | System Gameplay Modifier |
| :---: | :--- | :--- | :--- | :--- |
| **1 - 5** | **Smooth Synergy** | *The Big Lebowski* | Warm sub-pixel hue shift | $+2$ Luck; $+10\%$ gambling payouts. |
| **6 - 12** | **Greening Out** | *Pineapple Express* | Viewport radial sway | 20% chance to trip/stumble between nodes. |
| **13 - 20** | **Bat Country** | *Fear and Loathing* | Undercover NPCs turn into red lizards | Every NPC appears as an undercover agent. |
| **21+** | **Existential Dread**| *Jacob's Ladder* | Extreme RGB chromatic aberration | Character locks up for 2 turns believing simulation is real. |

---

## SECTION 7: EXTREME MOB STAKES, DESERT GRAVES & ESCAPE ENGINE

When Mob Respect hits 0%, Metro PD Heat reaches 5 Stars, or a heist fails, the game escalates into gritty mob-movie sequences.

```
+-------------------------------------------------------------------------------+
|                    EXTREME ESCALATION SEQUENCE FLOW                           |
+-------------------------------------------------------------------------------+
|  BLUNDER / HEIST FAILURE  -->  MOB ENFORCER AMBUSH  -->  TRUNK TRANSPORT      |
|                                                                 |             |
|  PERMANENT SCARRING      <--  BODY DISPOSAL    <--  DESERT PIT  v             |
+-------------------------------------------------------------------------------+
```

### 7.1 Real-Time Desert Pit Escapology State Table

#### Table 7.1: Real-Time Desert Escapology Table
| Countdown Phase | Timer Window | Cinematic Trope | Required Action Input | Real-Time Success Outcome | Failure / Execution Outcome |
| :---: | :---: | :--- | :--- | :--- | :--- |
| **1: Trunk Transport** | **12 Seconds** | *Pulp Fiction* | Stage Lockpick / Glitch | Trunk unlocks while car moves; jump to highway. | Car stops at Jean Dry Lake Bed; forced out at gunpoint. |
| **2: Grave Digging** | **15 Seconds** | *Casino* | Stage Distraction / Stall | Trigger atomic flash or fake chest pain; buys +10s. | Enforcer raises .38 revolver; proceed to Phase 3. |
| **3: Execution Draw** | **3 Seconds** | *The Good, Bad, Ugly* | Stage Pocket Sand / Throw | Blinds hitman; interrupts fatal gunshot. | **SHOT IN HEAD:** Enter Morgue Wake State. |
| **4: Desert Standoff**| **5 Seconds** | *No Country for Old Men*| Stage Disarm / Shoot | Kill Enforcer or convince him you know vault codes. | Buried to neck in sand; lose ALL equipped items. |

### 7.2 Evidence Disposal Table ("The Wolf Protocol")

> *"I'm Winston Wolfe. I solve problems."* — **Pulp Fiction (1994)**

#### Table 7.2: Body Disposal & Evidence Management Table
| Disposal ID | Location Required | Film Homage | Time Cost | Faction & Heat Outcome |
| :---: | :--- | :--- | :---: | :--- |
| **DP-ACID** | Subterranean Flood Drains | *Breaking Bad* / *Snatch* | 3 Turns | Dissolves body completely; **0 Heat remaining**. |
| **DP-LAKE** | Hoover Dam Spillway | *Casino* (1995) | 5 Turns | Body sinks; $+10\%$ Mob Respect (Professionalism). |
| **DP-BURN** | Casino Maintenance Incinerator | *Goodfellas* (1990) | 2 Turns | Body burned; 15% chance smoke trips fire alarm. |
| **DP-TRUNK**| Rival Parking Garage | *Payback* (1999) | 1 Turn | Shifts crime blame to Rival Mob (-15 Rival Rep). |

### 7.3 Permanent Physical Disfigurement Table

#### Table 7.3: Permanent Scarring & Physical Trauma
| Trauma ID | Source Trigger | Film Reference | Attribute Impact | Permanent World Effect |
| :---: | :--- | :--- | :---: | :--- |
| **TR-NOSE** | Lost Brawl in Sands Lounge | *Chinatown* (1974) | $-1$ Dignity, $+2$ Smarts | Intimidates street punks; $-10\%$ on illegal gear. |
| **TR-FINGER**| Failed Interrogation Check | *The Yakuza* (1975) | $-2$ Hardware, $+10\%$ Mob | Mob sees missing pinky; grants veteran respect. |
| **TR-BURN**  | Incinerator Flashback | *Darkman* (1990) | $-3$ Dignity, $+3$ Intimidation| Casino Pit Bosses watch closely; -2 Stealth in light. |
| **TR-LIMP**  | Shot in Jean Dry Lake Bed | *Reservoir Dogs* (1992)| $-2$ Hustle, No Sprint | Footsteps make distinct noise in quiet ducts. |

### 7.4 The Morgue Wake State Table

#### Table 7.4: Morgue Wake State Table
| Parameter | State Value | Cinematic Reference | Mechanical System Result |
| :---: | :--- | :--- | :--- |
| **Location** | County Morgue / Sunrise Hospital | *Crank* (2006) / *The Hangover* | Player wakes up on stainless steel gurney. |
| **Inventory**| Stripped to $0 Cash & Clothes | *Terminator* (1984) | Lose all main/off-hand items; retain embedded micro-chips. |
| **Toe Tag**  | `FLAG_TOE_TAG_ACTIVE` | *Pulp Fiction* (1994) | Grants $+2$ Stealth; Metro PD considers player dead for 50 turns. |

---

## SECTION 8: RADIO FREE GRACELAND, RAT PACK PANTHEON & THE 777 CRDT JACKPOT

```
+-------------------------------------------------------------------------------+
|                    THE COSMIC VEGAS LORE ENGINE PIPELINE                      |
+-------------------------------------------------------------------------------+
|  [ RADIO FREE GRACELAND ]  --> LoRa / Tor Audio & Spectrum Visualizer (`▄▀█`)   |
|  [ RAT PACK PANTHEON ]     --> Luck Entropy Engine / "The Cosmic Cooler" NPC  |
|  [ 777 CRDT JACKPOT ]      --> Global Vault Pool / Solenoid Flashover Event   |
+-------------------------------------------------------------------------------+
```

### 8.1 Radio Free Graceland (The P2P Spectrum & Audio Engine)
Radio Free Graceland (RFG) is an off-grid audio and data channel broadcast across the mesh, hosted by the eternal spirit of **The King**.

* **Sub-Pixel Spectrum Visualizer:** Rendered on the Z-4 HUD using `▄▀█` ANSI characters to show audio frequency bands in real-time.
* **Tactical Airwave Commentary:** Intersperses 1950s rockabilly tunes with real-time tactical warnings based on MUD event hooks.

#### Table 8.1: Radio Free Graceland Frequency & Broadcast Schedule
| Frequency / Sub-Channel | Era Target | Broadcast Content | Tactical Player Effect |
| :---: | :--- | :--- | :--- |
| **915.5 MHz / LoRa-1** | Era 4 (2026) | Cyber-Rockabilly & Cipher Codes | Reveals hidden SCADA IP nodes in the current sector. |
| **1420 kHz / AM Band** | Era 2 (1962) | Sands Lounge Live & Outfit Wiretaps| Warns of incoming Mob Capo raids 2 pulses early. |
| **5.000 MHz / WWV** | Era 1 (1953) | Atomic Test Countdown & Sun Records| Immunity to nuclear fallout radiation for 60 seconds. |

### 8.2 The Rat Pack Pantheon & Luck Entropy Engine
Luck in *Viva Las Mesh* is an environmental pressure governed by the Rat Pack Pantheon (Sinatra, Martin, Davis Jr., Hughes).

#### Table 8.2: Rat Pack Pantheon & Luck Mechanics
| Pantheon Member | Archetype Domain | Shrine Location | Special Offering | Pantheon Blessing / State |
| :---: | :--- | :--- | :--- | :--- |
| **Frank Sinatra** | "The Chairman" | Sands Copa Room | Bottle of Jack Daniel's | **The Chairman's Aura:** Metro PD Heat drops to 0 instantly. |
| **Dean Martin** | "Saint of Spirit" | Desert Inn Bar | Vintage Bourbon Glass | **Smooth Buzz:** Max Buzz stat without triggering Paranoia. |
| **Sammy Davis Jr.**| "Lord of Stage" | Stardust Mainstage | Silk Bowtie | **Velvet Presence:** +5 Dignity; free entry to high-roller suites. |
| **Howard Hughes** | "Sky Hermit" | Desert Inn Penthouse| Wiretap Reel Spool | **Satellite Master:** Unlocks orbital CCTV override keys. |

#### Table 8.3: Luck Entropy State Matrix
| Luck State | Trigger Condition | Visual Indicator | Mechanical Outcome |
| :---: | :--- | :--- | :--- |
| **In The Zone** | Win 5 consecutive bets / actions | RGB Neon Border Flash | +50% critical strike / hacking success rate. |
| **The Cooler (NPC)**| Player Luck exceeds 18 | Black-suited NPC enters room | Drains 50% Luck of all players in same room via CRDT. |
| **Snake Eyes** | Loss streak of 3 | Cold grey terminal screen | Next failed action triggers a Tier 2 interrupt on player. |

### 8.3 The Global 777 CRDT Progressive Jackpot
A single cross-network vault pool replicated across all nodes using Conflict-Free Replicated Data Types (CRDTs).

$$\text{Global Vault Pool} = \sum \text{Casino Losses} + \sum \text{Failed Heist Loot} - \sum \text{Jackpot Payouts}$$

#### Table 8.4: 777 Progressive Jackpot Trigger Table
| Trigger Condition | Network Event | Terminal Canvas Response | Payout Outcome |
| :---: | :--- | :--- | :--- |
| **7-7-7 Slot Spin** | Broadcast via CRDT Delta | Golden `U+2580` half-block flash (3 sec) | 100% Global Vault Cash distributed to player wallet. |
| **Mirage Vault Crack**| Broadcast via LoRa/Tor | Terminal bell sequence (`\a`) x 7 | 50% Global Vault Cash + Holographic King Manifest. |

---

## SECTION 9: FRICTIONLESS ROGUELIKE UX/UI & ACCESSIBILITY ARCHITECTURE

To ensure high usability by traditional and modern gaming standards without sacrificing terminal fidelity, *Viva Las Mesh* implements a zero-friction UI/UX framework designed for low barriers to entry and intuitive operation.

```
+-------------------------------------------------------------------------------+
|                       FRICTIONLESS TERMINAL UX ENGINE                         |
+-------------------------------------------------------------------------------+
| [ Contextual Smart Key (`Space`) ]  --> Auto-resolves obvious actions          |
| [ Sub-Pixel Mouse Navigation ]     --> Full point-and-click / tap tracking    |
| [ Fuzzy Action Palette (`Ctrl+K`) ] --> Real-time auto-complete & previews     |
| [ Diegetic Onboarding Tutorial ]   --> Low-stakes mini-heist prologue          |
+-------------------------------------------------------------------------------+
```

### 9.1 Context-Aware "Smart Key" Engine
Instead of memorizing dozens of obscure ASCII keybindings, a single primary button (**`Space`** or **`Enter`**) dynamically adapts to the player's immediate context.

#### Table 9.1: Smart Action Key Resolution Table
| Player Context / Facing Object | Highlighted UI Prompt | Single-Key Action (`Space`) | Alternative Key (`Tab`) |
| :--- | :--- | :--- | :--- |
| Facing Closed Vault Door | `[SPACE] Glitch Vault Lock` | Injects optimal SCADA payload | Opens Hardware Toolkit menu |
| Standing on Cash Bag | `[SPACE] Grab Dirty Cash` | Loots cash instantly | Inspects bag weight/details |
| Adjacent to Neutral NPC | `[SPACE] Talk / Bribe` | Initiates dialogue window | Draws equipped weapon |
| In Firefight (Enemy Facing) | `[SPACE] Fire Primary Weapon` | Executes Tier 4 Attack | Stages Tier 2 Interrupt (Pocket Sand) |
| Standing at SCADA Terminal | `[SPACE] Access Terminal` | Opens GUI Hacking Overlay | Bypasses power supply |

### 9.2 Sub-Pixel Mouse & Touch Navigation
The client leverages terminal ANSI SGR mouse tracking (`\x1b[?1006h`) to provide full mouse, trackpad, and touchscreen interactivity.

#### Table 9.2: Point-and-Click Sub-Pixel Interaction Table
| Mouse Input Event | Target Canvas Zone | UI/UX System Outcome |
| :--- | :--- | :--- |
| **Left Click / Tap** | World Map Cell (Z-1 / Z-2) | Auto-paths character to cell using A* pathfinding. |
| **Right Click / Hold** | Any Entity or Tile | Opens context radial menu with explicit action choices. |
| **Hover / Pointer Over**| Weapon, NPC, or Object | Displays instant sub-pixel floating HUD tooltip card. |
| **Drag & Drop** | Inventory Slot (Z-4) | Equips, drops, or uses item without keyboard commands. |
| **Scroll Wheel** | Command History / Terminal | Smoothly scrolls combat log and RF frequency bands. |

### 9.3 Command Visualizer & Fuzzy Action Palette (`Ctrl+K` / `/`)
For players who prefer command-line interaction, a modern fuzzy-search palette powered by `Bubble Tea` allows instant discoverability.

```
+-------------------------------------------------------------------------------+
| > hack                                                                        |
|-------------------------------------------------------------------------------|
| [Action] Hack SCADA Vault Controller   (Hardware Level 14 | 85% Success)      |
| [Action] Hack CCTV PTZ Camera Feeds   (Hardware Level 10 | 99% Success)      |
| [Action] Hack Metro PD IMSI Catcher   (Hardware Level 16 | 60% Success)      |
+-------------------------------------------------------------------------------+
```

#### Table 9.3: Fuzzy Action Palette & Ergonomic Controls
| Shortcut | Command Window Name | Functional Purpose & UX Benefit |
| :---: | :--- | :--- |
| **`Ctrl+K` / `/`** | **Fuzzy Command Finder** | Type any natural word ("vault", "bribe", "radio") to see valid actions. |
| **`Tab`** | **Context Cycle** | Toggles focus between Map View, Inventory, and RF Spectrum Analyzer. |
| **`Esc`** | **Universal Back / Cancel** | Clears active prompt, closes overlays, or pauses pulse staging. |
| **`F1` / `?`** | **Diegetic Overlay Help** | Shows overlay card mapping screen elements to plain English descriptions. |

### 9.4 Diegetic Onboarding: "The Vegas Street Guide"
Instead of wall-of-text manuals, new players enter a 3-minute guided prologue set in Era 4's subterranean flood channels.

#### Table 9.4: Guided Onboarding Progression
| Step | Mission Objective | Mechanics Taught | UX Safety Rail |
| :---: | :--- | :--- | :--- |
| **1** | Walk to the glowing neon sign | Movement (`WASD`, Arrow keys, or Click) | Hazard-free starting room. |
| **2** | Grab the discarded RF Sniffer | Item interaction & Smart Key (`Space`) | Item glows in sub-pixel gold. |
| **3** | Bypass the flood gate lock | Hacking / Hardware check with UI odds preview | Success guaranteed on first try. |
| **4** | Outsmart the Metro PD Scout | Simulturn pulse staging & Priority Tiers | Time freezes during action selection. |

---

## SECTION 10: MASTER REAL-TIME EXECUTION PIPELINE & EVENT DRIVER MATRIX

The engine runtime processes every tick through a unified master pipeline combining real-world networking, game logic, lore triggers, and sub-pixel UI updates.

```
+-------------------------------------------------------------------------------+
|                       MASTER EVENT DISPATCHER                                 |
+-------------------------------------------------------------------------------+
| [ MICRO-TICK PULSE (250ms) ] --> Collect Input (Mouse / Keys / P2P Mesh)      |
|                                  |                                            |
|                                  +--> Evaluate Priority Tiers (T1-T5)         |
|                                  +--> Resolve Simulturn Conflicts             |
|                                  +--> Update Physiology, Paranoia & Luck      |
|                                  +--> Process Radio Free Graceland Audio      |
|                                  +--> Evaluate Global 777 CRDT Jackpot Pool   |
|                                  +--> Stream CRDT Deltas & Render Sub-Pixel   |
+-------------------------------------------------------------------------------+
```

#### Table 10.1: Real-Time Master Event Execution Pipeline Matrix
| Step | Phase Name | System Action | Cinematic Focus | Output Target |
| :---: | :--- | :--- | :--- | :--- |
| **Step 1** | **Pulse Signal** | Emits 250ms micro-tick signal | *The Matrix* Clock Core | Engine Core |
| **Step 2** | **Input Gathering** | Reads Smart Key, mouse clicks, and mesh frames | *Frictionless UX Stack* | Action Queue |
| **Step 3** | **Priority Sorting**| Sorts actions into Tiers 1–5 | *Fantasy Flight Order Phase*| Resolution Stack |
| **Step 4** | **Simulturn Exec** | Resolves conflicts & attribute checks | *Heat* Firefight | World State |
| **Step 5** | **Lore & Audio Check**| Updates Radio Free Graceland spectrum & audio | *Radio Free Graceland* | Audio / HUD Overlay |
| **Step 6** | **Luck & Jackpot** | Calculates Luck entropy & CRDT 777 Jackpot | *Casino* / *Ocean's Eleven* | CRDT Vault Delta |
| **Step 7** | **MUD Decay Check**| Updates Hydration, THC, and Heat | *Fear and Loathing* | Player Stats |
| **Step 8** | **Network & Render**| Streams CRDT delta via LoRa/Tor/I2P; flushes half-block sub-pixel canvas | *Blade Runner* Screen | Terminal Display |