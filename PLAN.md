# Viva Las Mesh: Master Production Release Plan and Technical Roadmap
This master production release plan establishes the technical, architectural, and asset deployment pipeline for Viva Las Mesh—a terminal-native, multi-temporal, simulturn MUD and tactical heist simulator set across the Las Vegas Valley. Built in Go utilizing the Charmbracelet ecosystem (Bubble Tea, Lip Gloss, Bubbles), the engine executes off-grid state replication over multi-transport peer-to-peer (P2P) networks, integrating Meshtastic LoRa radio hardware, Tor v3 hidden services, I2P SAM garlic tunnels, and Bluetooth Low Energy (BLE) proximity meshes.
The technical execution matrix below translates the project’s design codices into an actionable, sequential master checklist. Every task must be completed systematically to transition the engine from its pre-alpha codebase into a production-ready release candidate.

## Progress Tracking

- [x] Milestone 1: Core Architecture & Environment Setup
- [ ] Milestone 2: Simulation Foundation & Deterministic Engine
- [ ] Milestone 3: Asset Pipeline & World Building
- [ ] Milestone 4: Quality Assurance, Balancing & Optimization
- [ ] Milestone 5: Production Readiness & Deployment

## Milestone 1: Core Architecture & Environment Setup
Milestone 1 establishes the low-level Go 1.24+ runtime, the terminal sub-pixel rendering engine, the multi-transport P2P background daemon (orchestrator-server), the Conflict-Free Replicated Data Type (CRDT) state synchronization engine, and the multi-tier pulse clock hierarchy.
The rendering layer operates inside ANSI-compliant terminal emulators, utilizing upper half-block characters (▀, U+2580) to achieve double vertical resolution (120 \times 80 effective pixels in a standard 120 \times 40 terminal cell). Rendering composition is strictly structured across a 5-layer Z-buffer matrix (Z-0 to Z-4) to isolate floor grids, terrain, entities, particle effects, and HUD interfaces. The network daemon operates out-of-process, interfacing with physical radio hardware and anonymized overlay networks to sync game state using Noise IK protocols and Merkle tree verification.
### Engine Core & Network Infrastructure Specifications
| Technical Subsystem | Component Target | Performance / Operational Standard | Protocol / Library Dependency |
|---|---|---|---|
| Terminal Canvas Engine | Sub-Pixel Renderer (▀) | 120 \times 80 effective resolution at 60 FPS | charmbracelet/lipgloss, bubbletea |
| Z-Buffer Matrix | Layers Z-0 through Z-4 | Strict depth order, additive blending on Z-3 | ANSI TrueColor (\x1b[38;2;...m) |
| Off-Grid Network Daemon | orchestrator-server | Standalone background service with local IPC | Go 1.24+ runtime, go-i2p/onramp |
| P2P LoRa Transport | TR-LORA (868/915 MHz) | 237-byte max frame size, lossy mesh delivery | fxamacker/cbor/v2, Meshtastic driver |
| Darknet Anonymity | TR-TOR3 & TR-I2P | 4096B (Tor v3) / 1024B (I2P Garlic) frames | Noise IK, Signal Double Ratchet |
| State Synchronization | CRDT Delta Engine | Associative, commutative, idempotent state | ipfs/go-datastore, cbergoon/merkletree |
| Pulse Clock Engine | Micro / Macro / Meta / Eon | 250ms / 1000ms / 10.0s / 60.0s divisions | Go channels & precision tickers |
### Execution Checklist: Core Architecture & Setup
 * [x] 1.1 Repository Initialization & Environment Configuration
   * [x] Initialize Go 1.24+ module workspace structure across designated directories (/cmd/mesh, /pkg/render, /pkg/p2p, /pkg/engine).
   * [x] Configure third-party dependency locks for charmbracelet/bubbletea, charmbracelet/lipgloss, and charmbracelet/bubbles.
    * [x] Add go-i2p/onramp dependency for I2P support
    * [x] Add fxamacker/cbor/v2 dependency for CBOR serialization
    * [x] Add Meshtastic driver dependency for LoRa communication
   * [x] Establish repository license compliance headers referencing the MIT License.
   * [x] Set up automated cross-platform build scripts targeting Linux (x86_64/ARM64), macOS (Apple Silicon/Intel), and Windows (x86_64).
 * [x] 1.2 Sub-Pixel Terminal Rendering Engine (▀) Implementation
   * [x] Construct the half-block (▀, U+2580) canvas buffer establishing 120 \times 80 sub-pixel addressing within standard terminal cells.
   * [x] Implement 24-bit TrueColor ANSI escape sequence generators for upper sub-pixel foreground (\x1b[38;2;R;G;Bm) and lower sub-pixel background (\x1b[48;2;R;G;Bm) color channels.
   * [x] Construct the 5-layer Z-buffer compositing pipeline: Z-0 (Floor & Grid), Z-1 (Terrain Details), Z-2 (Entities & FOV), Z-3 (Particle & FX), Z-4 (HUD & Spectrum Analyzer).
   * [x] Build viewport resize detection and dynamic sub-pixel canvas scaling routines to prevent ASCII boundary distortion during window resizing.
 * [x] 1.3 Multi-Transport P2P Daemon (orchestrator-server)
   * [x] Construct the orchestrator-server background service architecture to handle asynchronous network transport outside the main UI thread.
   * [x] Implement TR-LORA driver interface for Meshtastic 868/915 MHz radio devices with CBOR packet serialization.
   * [x] Implement TR-TOR3 Ephemeral v3 Hidden Service wrapper for anonymized TCP routing over public networks.
   * [x] Implement TR-I2P SAM Garlic Tunnel bridge via go-i2p/onramp integration for multi-hop encrypted transport.
   * [x] Implement TR-BLE Bluetooth Low Energy driver for local proximity peer discovery and direct payload exchange.
   * [x] Integrate Noise IK Protocol handshake and Signal Double Ratchet session management using chacha20poly1305 and curve25519 cryptographic primitives.
 * [ ] 1.4 CRDT State Convergence & Persistence Infrastructure
   * [x] Integrate ipfs/go-datastore for local key-value state persistence across application sessions.
   * [x] Implement Merkle tree state verification using cbergoon/merkletree to identify missing transaction deltas between nodes.
   * [x] Build Conflict-Free Replicated Data Type (CRDT) delta convergence logic adhering to mathematical state unification rules.
   * [x] Implement deterministic state reconciliation handlers for lossy, out-of-order packet delivery over radio airwaves.
 * [ ] 1.5 Master Clock Division & Pulse Core
   * [x] Implement 250ms Micro-Tick loop for real-time movement interpolation, input polling, and sub-pixel redraws.
   * [ ] Implement 1000ms Macro-Pulse loop for room state updates, guard pathfinding execution, SCADA timers, and packet dispatching.
   * [ ] Implement 10.0s Meta-Pulse loop for metabolic decay, hydration consumption, police heat decay, and mob interest accrual.
   * [ ] Implement 60.0s Eon-Pulse loop for spectral apparition checks and cross-network casino jackpot pool drift.
## Milestone 2: Gameplay Mechanics & Systems Integration
Milestone 2 operationalizes the simulturn resolution engine, multi-temporal era logic, diegetic electromagnetic hacking, faction political state machines, survival mechanics, and mob escalation protocols.
Actions staged during micro-ticks resolve simultaneously at macro-pulse boundaries using a strict 5-tier priority matrix. In-game hacking spans the physical RF spectrum from VLF to Ku-band and optical fiber, deploying zero-days and SCADA overrides against casino infrastructure. Gameplay incorporates hardcore desert survival mechanics, a chemical dispensary catalog, a dynamic Paranoia Index ("Bat Country" visual shaders), and high-stakes mob punishment systems, including real-time desert grave escape sequences and the Morgue Wake State.
### Simulturn Priority & System Mechanics Specifications
| System Domain | Operational Rule / Mechanism | System Outcome / Priority Tier |
|---|---|---|
| Tier 1 Priority | Divine & Spectral Intervention | Transcends physical reality; overrides all physical actions |
| Tier 2 Priority | Interrupt Actions | Pre-empts T3–T5 actions; cancels target's queued staging |
| Tier 3 Priority | Fast Actions | Resolves before standard actions; Hustle attribute tie-breaker |
| Tier 4 Priority | Standard Combat & Hacking | Simultaneous damage and effect resolution |
| Tier 5 Priority | Heavy Field Operations | Sustained multi-tick staging; interrupted by damage |
| Police Heat | 0 to 5 Star Pursuit Machine | Modifies guard response rate from 10-pulse loops to 250ms ticks |
| Paranoia Index | Chemical & Heat Derivative | Triggers canvas shaders: Smooth Synergy to Existential Dread |
| Escapology | Desert Pit Real-Time Sequences | 4-phase timed survival sequence (Trunk, Dig, Execution, Standoff) |
### Execution Checklist: Gameplay Mechanics & Systems Integration
 * [ ] 2.1 Simulturn Staging & Resolution Matrix
   * [ ] Implement action staging queue processing actions across Priority Tiers 1 through 5.
   * [ ] Build attribute-based tie-breaker evaluation routines (Hustle, Hardware, Street Smarts) for equal-tier action collisions.
   * [ ] Construct simultaneous resolution solver resolving concurrent attacks, dual pickups, and spatial movement collisions.
 * [ ] 2.2 Multi-Temporal Era Engine & Anachronism Detection
   * [ ] Construct world state engines for Era 1 (1953 Atomic AEC), Era 2 (1962 Mob Syndicate), Era 3 (1993 Corporate Mega-Resort), and Era 4 (2026 Cyber-Strip).
   * [ ] Implement Anachronism Detection Engine evaluating item tech tags against current era settings.
   * [ ] Build Anachronism Flag handlers (FLAG_ALIEN_TECH, FLAG_WIRE_BUG, FLAG_ANTIQUE) triggering localized faction standing hits.
   * [ ] Construct The Transcendent King (Elvis) spectral state machine across forms: Sun Records Rockabilly (1953), Hollywood Gold-Suit (1962), White Jumpsuit Oracle (1993), and Holographic Code-God (2026).
 * [ ] 2.3 Diegetic RF Spectrum Hacking & Exploits
   * [ ] Implement full-spectrum RF receiver architecture covering VLF (3 kHz) through Ku-Band (18 GHz) and Optical/IR frequencies.
   * [ ] Build diegetic exploit payload engines: PL-IMSI (Cellular), PL-SCADA (HVAC), PL-PRNG (Slots), PL-PINCH (EMP Substation), PL-SAT (Satellite), PL-GLITCH (Vault Voltage), PL-GPS (Transport Rerouting), PL-LASER (Fiber Tap), PL-CAN (Elevator), PL-BGP (Data Center).
   * [ ] Integrate spectrum signal visualization data handlers feeding live RF waves to Z-4 HUD elements.
 * [ ] 2.4 Factions, Heat State Machine & Vault Heist FSM
   * [ ] Implement faction disposition metrics for Atomic Energy Commission (AEC), Chicago Outfit, Corporate Moguls, Darknet Hackers, and Metro PD.
   * [ ] Build real-time Police Heat state machine (0 to 5 Stars) accelerating police dispatch ticks from 10 pulses down to continuous 250ms micro-ticks.
   * [ ] Implement "The Skim" money laundering channels across high-stakes craps (LN-CRAPS), dispensary retail (LN-DISP), pawn shops (LN-PAWN), and casino count rooms (LN-CAGE).
   * [ ] Construct 4-stage simulturn vault heist state machine: Stage 1 (Surveillance Bypass), Stage 2 (Guard Neutralization), Stage 3 (Vault Unlock), Stage 4 (Cash Extraction).
 * [ ] 2.5 Physiology, Dispensary Catalog & Paranoia Shaders
   * [ ] Implement core player attribute handlers: Health (0–100), Buzz (0–10), Hydration (0–100%), Hustle (1–20), Street Smarts (1–20), Dignity (1–20), Hardware (1–20).
   * [ ] Program chemical intake mechanics for dispensaries: THC Distillate (CH-DIST), Live Resin (CH-CART), Nano-THC Soda (CH-SODA), Melatonin (CH-MEL), Psilocybin (CH-SHRM), THCA Diamond Wax (CH-WAX).
   * [ ] Construct mathematical Paranoia Index derivation engine combining THC dosage, metabolic ticks, alcohol level, heat, and street smarts.
   * [ ] Build terminal canvas color distortion shaders for Paranoia states: Smooth Synergy (1–5), Greening Out (6–12), Bat Country (13–20), Existential Dread (21+).
 * [ ] 2.6 Extreme Mob Escalation, Escapology & Wolf Protocol
   * [ ] Construct real-time 4-phase Desert Pit Escapology state machine with timed inputs: Trunk Transport (12s), Grave Digging (15s), Execution Draw (3s), Desert Standoff (5s).
   * [ ] Implement "The Wolf Protocol" body disposal mechanics across Acid Drains (DP-ACID), Hoover Dam (DP-LAKE), Incinerator (DP-BURN), and Rival Parking (DP-TRUNK).
   * [ ] Program Permanent Scarring attribute modifiers (TR-NOSE, TR-FINGER, TR-BURN, TR-LIMP).
   * [ ] Build Morgue Wake State lifecycle handler setting FLAG_TOE_TAG_ACTIVE, stripping physical inventory, and resetting police aggression.
 * [ ] 2.7 Cosmic Lore Engine, Radio Free Graceland & 777 CRDT Jackpot
   * [ ] Construct Radio Free Graceland airwave broadcast channel over LoRa (915.5 MHz), AM Band (1420 kHz), and WWV Shortwave (5.000 MHz).
   * [ ] Implement Rat Pack Pantheon luck entropy system governing Frank Sinatra, Dean Martin, Sammy Davis Jr., and Howard Hughes offerings.
   * [ ] Build "The Cosmic Cooler" roaming NPC logic that drains 50% player luck via CRDT state updates.
   * [ ] Program Global 777 CRDT Progressive Jackpot aggregating casino losses and heist failures into a cross-network pool.
## Milestone 3: Asset Pipeline & World Building
Milestone 3 focuses on populating the game world with item assets, procedural salvage generation engines, crafting workbenches, sub-pixel visual art, audio spectrum visualizers, spatial level maps, and frictionless user interface controls.
The asset pipeline utilizes 6 \times 6 dual-axis matrices to generate hundreds of procedural item profile variations from mundane desert garbage and corporate trash. Level design spans authentic historical and fictionalized Las Vegas sites across all four temporal eras. Interface controls prioritize zero execution friction, combining context-aware smart keys (Space / Enter), ANSI mouse tracking, and a fuzzy-search action palette (Ctrl+K).
### Content Pipeline & UI/UX Specifications
| Asset / UI Subsystem | Structure Target | Functional Specification |
|---|---|---|
| Procedural Salvage | 6 \times 6 Generation Matrices | 6 distinct matrices mapping material bases against wear states |
| Crafting Workbenches | Tiers 1 through 4 | Inventory Jury-Rig up to Automated Industrial Fabricators |
| Scrap Breakdown | 6 Salvage Currencies | Polymer, Conductive Metal, Solvents, Silicate, Steel, Fibers |
| Era Color Palettes | 4 Historical Schemes | 1953 Amber, 1962 Magenta, 1993 Teal, 2026 Cyan |
| World Map Sectors | 8 Key Node Locations | Test Site, Sky Room, Stardust, Jean Lake, Mirage, Tunnels, etc. |
| Frictionless UX | Smart Key (Space) | Contextually resolves highest-priority environmental action |
| Action Palette | Fuzzy Finder (Ctrl+K) | Natural language search with embedded success probability |
### Execution Checklist: Asset Pipeline & World Building
 * [ ] 3.1 Procedural Salvage & Item Asset Pipeline
   * [ ] Ingest mundane survival utilities: Terrible's Cup, Casino Water, Electrolyte Packets, Matchbooks, Ice Bucket Liners, Antacids.
   * [ ] Ingest corporate credentials: Players Club Cards, RFID Keys, VIP Wristbands, Shift Schedules, Uniform Patches, Press Passes.
   * [ ] Ingest hardware scrap: Burner Phones, SDR Dongles, Kiosk Readers, Micro-SD Shards, Multimeters, Induction Coils.
   * [ ] Ingest souvenirs & lore items: Glow Flamingos, Plastic Horseshoes, Elvis Programs, Ceramic Dice, Inflatable Palms, Cowboy Hats.
   * [ ] Ingest containers & carrying gear: Souvenir Buckets, Duty-Free Bags, Fanny Packs, Metal Cash Boxes, Laundry Hampers, Insulated Bags.
   * [ ] Ingest standardized lethal weapons: Melee (Club, Knife) and Firearms (Pistol, Revolver, Rifle, Machine Gun).
   * [ ] Construct 6 \times 6 procedural loot generation matrices:
     * [ ] Matrix 3.1: Hydration, Dumpster Finds & Slime Utility.
     * [ ] Matrix 3.2: Corporate Access, Scams & Counterfeit Credentials.
     * [ ] Matrix 3.3: Electronics, Hacking Implants & Radio Scrap.
     * [ ] Matrix 3.4: Souvenirs, Memorabilia & Psychological Warfare.
     * [ ] Matrix 3.5: Containers, Carrying Gear & Improvised Storage.
     * [ ] Matrix 3.6: Standardized Lethal Weapons Condition States.
 * [ ] 3.2 Field Crafting & Workbench Infrastructure
   * [ ] Implement Workbench Tier mechanics: Tier 1 (Field Jury-Rig), Tier 2 (Burn-Barrel Workbench), Tier 3 (Deep-Sublevel Bench), Tier 4 (Automated Industrial Fabricator).
   * [ ] Program field recipes: Styrofoam Muzzle Sleeve, Acoustic Tripwire Alarm, Glow-Decoy Flare, LoRa Signal Booster, Dual-Cell Power Bank, Cloned Mag-Stripe Key, Improvised Napalm Trap, Acidic Corrosive Drip, Waterproof Data Shroud, Improvised EMP Grenade, Spiked Caltrops, Filtered Respirator.
   * [ ] Implement scrap salvage conversion pipeline breaking items down into Polymer, Conductive Metal, Chem-Solvents, Silicate Glass, Structural Steel, and Organic Fibers.
 * [ ] 3.3 Sub-Pixel Visual Art & Era Color Palettes
   * [ ] Define sub-pixel color palettes for historical eras: 1953 (Atomic Amber), 1962 (Neon Magenta), 1993 (Corporate Teal), 2026 (Cyber Cyan).
   * [ ] Render half-block character art tiles for craps tables, slot reels, vault doors, neon marquees, and desert sand dunes.
   * [ ] Construct dynamic lighting and particle rendering handlers for RF waves, muzzle flashes, cigarette smoke, and neon glow on Z-3.
 * [ ] 3.4 Audio Spectrum Visualizer & Audio Subsystem Integration
   * [ ] Construct ▄▀█ character sub-pixel audio spectrum visualizer for the Z-4 HUD layer.
   * [ ] Integrate Radio Free Graceland streaming hooks and terminal bell (\a) audio flash triggers.
   * [ ] Program sound effect event triggers for weapons fire, lock picking, motor hums, and desert winds.
 * [ ] 3.5 Level Design & World Map Construction
   * [ ] Map and construct Era 1 spatial nodes: Nevada Test Site Bunkers, Desert Inn Sky Room.
   * [ ] Map and construct Era 2 spatial nodes: Stardust Vault & Count Room, Jean Dry Lake Bed.
   * [ ] Map and construct Era 3 spatial nodes: Mirage Mega-Vault Atrium, Underground Service Tunnels.
   * [ ] Map and construct Era 4 spatial nodes: Subterranean Flood Channels, Mega-Dispensary Complex.
   * [ ] Implement A^* pathfinding navigation grids across all interior and exterior map terrain tiles.
 * [ ] 3.6 Frictionless UI/UX Engine
   * [ ] Implement Context-Aware Smart Key (Space / Enter) resolving high-priority environmental actions.
   * [ ] Implement ANSI SGR mouse tracking (\x1b[?1006h) supporting left-click movement, right-click context radials, hover tooltips, and drag-and-drop inventory management.
   * [ ] Construct Fuzzy Action Palette (Ctrl+K / /) featuring natural language auto-complete and success probability previews.
   * [ ] Build "The Vegas Street Guide" diegetic onboarding prologue sequence set in Era 4 flood channels.
## Milestone 4: Quality Assurance, Balancing & Optimization
Milestone 4 subjects the compiled engine to rigorous performance optimization, P2P network simulation, simulturn collision testing, economy balancing, and bug triage.
Performance tuning guarantees stable sub-16ms frame render times (60 FPS) across various ANSI terminal emulators. Multi-transport P2P netcode is stress-tested under high packet drop and latency conditions to confirm CRDT convergence stability. System balancing calibrates item spawn rates, money laundering fees, paranoia escalation curves, and combat damage metrics.
QA Benchmarks & Performance Targets
| Testing Area | Metric / Target | Verification Method |
|---|---|---|
| Terminal Frame Rate | \ge 60 \text{ FPS} (sub-16ms render tick) | Automated frame timer profiling across viewports |
| Memory Allocation | Zero allocation spikes during ANSI drawing | Go benchmark allocations (b.ReportAllocs()) |
| Mesh Resilience | 90% simulated packet loss tolerance | Automated LoRa drop-rate network testing harness |
| CRDT Convergence | 100% state parity across 50 peer nodes | Automated Merkle tree comparison scripts |
| Simulturn Determinism | 0 state desyncs during 100-player shootouts | Headless simulturn collision simulation suites |
| Terminal Compatibility | Full compatibility with Kitty, Alacritty, WezTerm | Matrix validation suite across terminal drivers |
### Execution Checklist: Quality Assurance & Optimization
 * [ ] 4.1 Terminal Rendering & Sub-Pixel Performance Optimization
   * [ ] Profile Lip Gloss / Bubble Tea render loops to enforce sub-16ms frame redraw rates.
   * [ ] Benchmark Z-buffer cell compositing performance across 120 \times 80 sub-pixel viewport boundaries.
   * [ ] Optimize memory allocations during ANSI escape sequence string formatting.
   * [ ] Validate rendering output and ANSI mouse alignment across Kitty, Alacritty, WezTerm, and Windows Terminal.
 * [ ] 4.2 Multi-Transport P2P Netcode & CRDT Stress Testing
   * [ ] Simulate lossy, high-latency RF airwaves for Meshtastic LoRa transport (TR-LORA) under 90% packet drop conditions.
   * [ ] Test Tor v3 and I2P SAM garlic tunnel failover and path recovery during network partitioning.
   * [ ] Verify CRDT state convergence idempotency, associativity, and commutativity across 50 simulated peer nodes.
   * [ ] Audit Noise IK handshake efficiency and Double Ratchet key rotation under hostile packet injection scenarios.
 * [ ] 4.3 Simulturn Mechanics & Priority Tier Stress Testing
   * [ ] Conduct automated simulation tests for Tier 1 through Tier 5 simulturn priority collisions.
   * [ ] Validate attribute tie-breaker determinism during simultaneous 100-player combat pulses.
   * [ ] Test 4-stage vault heist FSM transition stability under multi-player concurrent action staging.
   * [ ] Audit Desert Pit Escapology real-time countdown timers across varying system clock loads.
 * [ ] 4.4 Economy Balancing & Parameter Tuning
   * [ ] Balance "The Skim" money laundering fee percentages and police heat generation against high-roller casino payouts.
   * [ ] Fine-tune Paranoia Index formula weights to ensure smooth transitions between Smooth Synergy and Bat Country states.
   * [ ] Calibrate procedural item spawn rates across 6 \times 6 salvage matrices to prevent economy inflation.
   * [ ] Adjust weapon damage profiles, noise radii, and wear degradation rates for standard armaments.
 * [ ] 4.5 Playtesting, Accessibility & Bug Triage
   * [ ] Execute user playtesting sessions verifying context-aware Smart Key (Space) prompt accuracy.
   * [ ] Test ANSI mouse tracking precision across varying terminal font sizes and screen geometries.
   * [ ] Conduct end-to-end playthrough audits of "The Vegas Street Guide" onboarding sequence.
   * [ ] Triage and resolve blocker, critical, and major bugs in the issue tracking system.
## Milestone 5: Production Readiness & Deployment
Milestone 5 executes the final release packaging, binary compilation, licensing compliance verification, network seed node deployment, and launch execution.
Final binaries are cross-compiled into standalone packages containing both the orchestrator-server daemon and the viva-las-mesh TUI executable. Darknet seed nodes (Tor v3 .onion and I2P SAM destinations) and Meshtastic radio channels are established to bootstrap off-grid state replication.
### Deployment Pipeline & Release Artifacts
| Release Target | Artifact Type | Distribution Channel | Compliance / Security Target |
|---|---|---|---|
| Linux (x86_64 / ARM64) | Standalone Binary Pair | GitHub Releases / Tarball | SHA-256 Checksum, MIT License |
| macOS (Universal) | Signed App / Binary | GitHub Releases / Zip | Apple Silicon / Intel Compatible |
| Windows (x86_64) | Executable Bundle | GitHub Releases / Zip | Windows Terminal TrueColor |
| Tor Seed Nodes | .onion Hidden Service | Darknet P2P Overlay | Tor v3 Ephemeral Keys |
| I2P Seed Nodes | SAM Destination | Darknet P2P Overlay | I2P Garlic Router |
| LoRa Channels | Config Payload | Physical Radio Airwaves | 868/915 MHz Public Mesh |
### Execution Checklist: Production Readiness & Deployment
 * [ ] 5.1 Release Packaging & Binary Cross-Compilation
   * [ ] Execute automated cross-compilation builds for Linux (x86_64, arm64), macOS (universal), and Windows (x86_64).
   * [ ] Package standalone orchestrator-server background daemon alongside main viva-las-mesh executable.
   * [ ] Implement automated cryptographic code signing and SHA-256 checksum generation for all build distributions.
   * [ ] Apply dead-code elimination and symbol table stripping flags (-ldflags="-s -w") to minimize binary sizes.
 * [ ] 5.2 Licensing & Compliance Verification
   * [ ] Perform repository audit ensuring proper inclusion of the MIT License text and copyright headers.
   * [ ] Conduct third-party library license compliance checks for all Go module dependencies.
   * [ ] Verify the complete elimination of temporary debug flags, unencrypted test credentials, or hardcoded secrets.
 * [ ] 5.3 Network Seed Node & Channel Deployment
   * [ ] Deploy initial bootstrap P2P seed nodes across Tor v3 .onion services and I2P SAM garlic destinations.
   * [ ] Configure and publish default Meshtastic LoRa channel definitions for regional off-grid communication.
   * [ ] Initialize and publish the starting Global 777 CRDT Progressive Jackpot state root to seed nodes.
 * [ ] 5.4 Launch Execution & Systems Activation
   * [ ] Perform final end-to-end launch verification from clean install to off-grid multi-player session.
   * [ ] Publish official viva-las-mesh v1.0 release packages to distribution channels.
   * [ ] Activate Radio Free Graceland airwave broadcast channel and initiate the global MUD heartbeat loop.
   * [ ] Transition project repository from staging to active production release state.
## Technical Summary & Critical Path Dependencies
To maintain production momentum, execution must strictly follow the critical path dependencies outlined below. No milestone phase may be bypassed:
 * Core Renderer Before Mechanics: Sub-pixel canvas rendering (▀) and Z-buffer compositing (Milestone 1) must be stabilized before rendering complex entity movement, dynamic lighting, or Paranoia canvas shaders (Milestone 2).
 * P2P Transport Daemon Before State Sync: The orchestrator-server daemon and multi-transport drivers (TR-LORA, TR-TOR3, TR-I2P) (Milestone 1) must be operational before testing CRDT state convergence or global 777 progressive jackpots (Milestone 2).
 * Simulturn Engine Before Content: Staged priority tier resolution (Milestone 2) must be fully deterministic before scripting complex vault heist state machines or procedural item matrices (Milestone 3).
 * Optimization Before Deployment: Terminal frame rate profiling and P2P packet drop simulation (Milestone 4) must meet production benchmarks before generating release candidate binaries (Milestone 5).
