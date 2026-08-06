# VIVA LAS MESH — ITEMS, SURVIVAL FRAMEWORK & PROCEDURAL SALVAGE (v5.0)
**Document Version:** 5.0.0  
**Domain:** Inventory System, Mundane Utility, Atmospheric Lore Design, Matrix-Based Procedural Generation, and Field Crafting

---

## 1. DESIGN PHILOSOPHY: MUNDANE UTILITY VS. LETHAL SIMPLICITY

In the Vegas Valley—spanning atomic test sites, corporate neon strips, and encrypted darknet radio zones—survival is rarely about raw firepower. **Lethal weapons are intentionally kept generic, functional, and brutal** (pistol, revolver, knife, rifle, machine gun, club), shifting player ingenuity toward clever, environmentally contextual utility items.

A rusted souvenir cup, a laminated players card, a discarded convention lanyard, or a soggy receipt from an all-night diner often provides a higher tactical advantage than a firearm when navigating police heat, corporate security scans, or the dehydrating expanse of the Mojave Desert. The overarching design mandate prioritizes emergent tactical depth through systemic friction, ensuring that every piece of garbage discovered in the ashes of the Strip serves a distinct mechanical purpose during simulturn resolution or exploration phases. Players must master the delicate balance between inventory weight limits, social stealth camouflage profiles, and the relentless degradation of improvised gear under high-heat desert conditions.

---

## 2. INVENTORY CLASSIFICATIONS & FRAMEWORK

### 2.1 Hydration & Survival Utilities (The Desert Edge)
The Valley environment kills faster than bullets. Water management, heat stroke prevention, and local ecosystem tricks dictate early-game survival, requiring players to constantly weigh hydration sources against severe toxic contamination vectors and heat exhaustion meters.

* **Cracked 32oz Terrible’s Styrofoam Cup (Dumpster-Dived)**
  * *Description:* A sun-bleached, cracked, moisture-rotted foam cup fished straight from a grease-slicked commercial dumpster behind an abandoned off-strip convenience store. It smells faintly of stale diet cola and chemical regret.
  * *Vegas Benefit:* Completely worthless for actual hydration (it leaks through micro-fractures), but acts as an exceptional acoustic dampening sleeve when shoved over the muzzle of a pistol or used as a crude funnel for siphoning automotive gasoline.
* **Casino Bottled Water (Sealed vs. Refilled)**
  * *Description:* Premium artesian water bottled exclusively for high-roller suites, complete with gold foil labeling and alkaline mineral additives. 
  * *Vegas Benefit:* Restores hydration without triggering radiation/toxin penalties. Empty bottles can be refilled with graywater for chemical experiments or used as crude acoustic tripwire noisemakers in narrow hotel corridors.
* **Electrolyte Packets & Casino Salt Licks**
  * *Description:* Pocket-sized recovery salts lifted directly from hotel hospitality amenity baskets and VIP lounges.
  * *Vegas Benefit:* Mitigates the stamina depletion effects of high heat waves and offsets the dangerous dehydration penalties caused by alcohol consumption, stim use, or synthetic THC distillates.
* **Sun-Baked Cardboard Matchbook**
  * *Description:* A damp matchbook from a defunct strip club featuring faded foil typography.
  * *Vegas Benefit:* Provides a reliable chemical spark source for igniting burn-barrel campfires, though individual matches degrade by 20% for each hour spent wading through high-humidity sewer zones.
* **Melted Hotel Ice Bucket Liner**
  * *Description:* A thin, translucent plastic bag salvaged from a luxury guest room ice bucket.
  * *Vegas Benefit:* Can be tied around plant roots to harvest condensation dew overnight, or worn over boots as a makeshift vapor barrier to protect feet from toxic industrial marsh sludge.
* **Expired Antacid Tablets (Alka-Seltzer Strip)**
  * *Description:* A blister pack of chalky mint tablets found in a pharmacy medicine cabinet.
  * *Vegas Benefit:* Neutralizes minor food poisoning and acid reflux penalties caused by consuming contaminated rations, restoring baseline stamina regeneration rates.

### 2.2 Corporate Access & Social Engineering
Bypassing physical security requires blending into the tourist herd or exploiting corporate credential systems to trick automated biometric checkpoints and AI surveillance drones.

* **Laminated Players Club Card (Gold/Platinum tier)**
  * *Description:* Scratched magnetic-stripe card from a defunct or active Strip mega-resort, carrying residual high-roller status markers.
  * *Vegas Benefit:* Grants low-level biometric clearance in casino back-of-house service corridors. Lowers suspicion meters when wandering restricted employee areas, provided you maintain a confident, unbothered "lost tourist" stride.
* **Stolen Keycard / RFID Room Key**
  * *Description:* Unprogrammed or active room key lifted from a slot machine change tray or left behind on a sticky cocktail table.
  * *Vegas Benefit:* Bypasses electronic turnstiles and elevator locks in specific resort zones. Can be physically de-magnetized against CRT terminals to improvise temporary magnetic triggers for trap doors.
* **VIP Cabana Wristband**
  * *Description:* A cloth festival-style wristband soaked in spilled gin, sunscreen, and UV-reactive dye.
  * *Vegas Benefit:* Grants immediate passage through VIP security checkpoints and allows temporary sanctuary from Metro PD patrols inside crowded pool areas and luxury cabana complexes.
* **Forged Security Shift Schedule**
  * *Description:* A grease-stained clipboard sheet detailing patrol paths for night-shift guards in sector 4.
  * *Vegas Benefit:* Reveals blind spots in surveillance camera arcs and grants a temporary bonus to stealth movement rolls when navigating corporate administration wings.
* **Unregistered Maintenance Uniform Patch**
  * *Description:* A frayed blue embroidered patch reading "HVAC SYSTEMS - FACILITIES."
  * *Vegas Benefit:* Can be safety-pinned over civilian jackets to gain immediate immunity from lower-tier security challenges when lingering near utility access panels and cooling towers.
* **Convention Press Pass Badge**
  * *Description:* A plastic laminate badge featuring a generic smiling headshot and a holographic barcode.
  * *Vegas Benefit:* Bypasses barricades set up around media zones and press-only conference halls, allowing unimpeded movement through corporate event complexes.

### 2.3 Electronics, Hacking & Communications
Bridging the physical grid with Tor, I2P, and Meshtastic LoRa infrastructures requires specialized hardware scavenging, custom kernel flashing, and tactical field modifications.

* **Burner Smartphone (Pre-paid / Rooted)**
  * *Description:* A cracked plastic handset running a custom Linux kernel with integrated LoRa transceiver drivers and stripped telemetry modules.
  * *Vegas Benefit:* Acts as the primary client interface for local mesh pinging, running `go-i2p/onramp` node connections, and scanning local Wi-Fi/Bluetooth beacons for open SCADA endpoints.
* **SDR (Software Defined Radio) USB Dongle**
  * *Description:* A pocket-sized aluminum receiver tuned to intercept local emergency dispatch, hotel maintenance frequencies, and wireless slot machine telemetry.
  * *Vegas Benefit:* Visualizes police heat vectors, corporate drone telemetry, and active patrol cones directly on the Z-4 HUD spectrum analyzer.
* **Modded Casino Loyalty Kiosk Reader**
  * *Description:* A pocket-sized skimmer wired directly into a custom battery pack and an Arduino microcontroller.
  * *Vegas Benefit:* Allows passive harvesting of tourist credit telemetry and room-key hashes when brushed against pocket terminals in crowded casino elevators or monorail cars.
* **Encrypted Micro-SD Data Shard**
  * *Description:* A tiny storage medium recovered from a burned server rack in an abandoned tech incubator.
  * *Vegas Benefit:* Contains cached darknet routing tables and cryptographic keys that can be traded with information brokers for high-tier weapon blueprints or safehouse coordinates.
* **Handheld Multimeter with Frayed Probes**
  * *Description:* A yellow plastic electrical tester with electrical tape holding the display bezel together.
  * *Vegas Benefit:* Measures live electrical currents in hacked door panels and prevents accidental electrocution minigame failures when interacting with exposed wiring.
* **Desoldered Copper Induction Coil**
  * *Description:* A tight spool of enameled wire harvested from the chassis of an old television monitor.
  * *Vegas Benefit:* Used to broadcast targeted electromagnetic pulses that temporarily disable magnetic door locks and perimeter security alarms.

### 2.4 Souvenirs & Psychological Warfare (Lore Items)
Mundane tourist trash transformed into tactical assets through sheer audacity, local superstition, and psychological disruption.

* **Glow-in-the-Dark Plastic Flamingo**
  * *Description:* A cheap neon lawn ornament bought at an open-air gift shop off Fremont Street.
  * *Vegas Benefit:* Emits a faint, non-lethal UV luminescence. Can be left in dark service stairwells to trick automated motion turret targeting or used as a crude distraction marker for distracted human guards.
* **Wedding Chapel Plastic Horseshoe**
  * *Description:* A mold-injected plastic good-luck charm from an all-night drive-thru chapel near the old courthouse.
  * *Vegas Benefit:* Can be tossed to create a sharp plastic clatter, drawing NPC attention away from your actual position. Tied to local luck mechanics; carrying it slightly buffers minor RNG failure rolls against slot and table games.
* **Faded Elvis Show Program (1976)**
  * *Description:* Water-stained glossy paper featuring theatrical imagery of The King in sequined attire.
  * *Vegas Benefit:* Instantly pacifies or triggers nostalgic dialogue loops from older NPC security guards and casino pit bosses who remember the golden era, granting a safe conversational window.
* **Neon Ceramic Dice Set**
  * *Description:* A pair of weighted casino dice salvaged from an imploded gaming floor.
  * *Vegas Benefit:* Can be rolled on hard concrete surfaces to generate distinct acoustic echoes, masking the sound of lockpicking or reloading mechanisms.
* **Inflatable Casino Neon Palm Tree**
  * *Description:* A deflated vinyl novelty toy smelling of old pool chlorine and warm plasticizers.
  * *Vegas Benefit:* Can be inflated with an air pump and stuffed into air ventilation shafts to block tear gas dispersion or create a visual barricade against security cameras.
* **Gaudy Rhinestone Cowboy Hat**
  * *Description:* A heavily decorated white felt hat covered in fake diamonds and dust.
  * *Vegas Benefit:* Increases charisma checks when dealing with wasteland merchants and traveling traders, but inflicts a severe penalty to stealth camouflage ratings in shadow zones.

### 2.5 Containers, Carrying Gear & Concealment
Inventory volume is strictly bound to physical weight, spatial dimension limits, and social profile visibility constraints.

* **Souvenir Plastic Bucket (Sealed with Lid)**
  * *Description:* A 2-gallon neon-pink bucket originally filled with 190-proof neon slushie alcohol, now washed out and fitted with a gasket lid.
  * *Vegas Benefit:* Waterproof container used for carrying stolen casino chips, chemical components, or dry rations across the desert. Can be inverted to serve as a makeshift stool or signal reflector.
* **Overstuffed Duty-Free Shopping Bag**
  * *Description:* Heavy paper bag from a high-end designer boutique filled with crumpled tissue paper and discarded packaging.
  * *Vegas Benefit:* Completely conceals generic firearms, hacking tools, and illicit hardware from casual visual inspection by security guards, allowing you to walk right through metal detectors disguised as a wealthy shopper.
* **Tactical Fanny Pack (Nylon)**
  * *Description:* Weathered waist pack worn aggressively across the chest, featuring frayed zippers and faded military webbing.
  * *Vegas Benefit:* Provides instant-access hot slots for items (lockpicks, stims, burner phone) without requiring inventory menu navigation during active turn pulses.
* **Rusted Metal Cash Box**
  * *Description:* A heavy lockbox salvaged from a ticket counter, dented by blunt force trauma.
  * *Vegas Benefit:* Offers secure, fire-resistant storage for sensitive ciphers and currency, but adds a heavy encumbrance penalty to inventory movement speed.
* **Heavy Canvas Laundry Hamper Bag**
  * *Description:* A massive industrial canvas sack marked with resort linen inventory codes.
  * *Vegas Benefit:* Holds an immense volume of bulky salvage and heavy scrap metal, but requires two hands to drag, completely disabling sprinting mechanics.
* **Insulated Cooler Bag**
  * *Description:* A soft-sided nylon beverage carrier lined with deteriorated silver foil insulation.
  * *Vegas Benefit:* Protects temperature-sensitive medical stims, chemical reagents, and biological samples from extreme ambient desert heat degradation.

### 2.6 Lethal Weapons (Standardized & Nonspecific)
To maintain tactical crunch and focus gameplay on environmental interaction, lethal armaments remain standardized, preventing gear score inflation:

* **Melee / Impact:** `Club` (tire iron, heavy flashlight, pipe), `Knife` (switchblade, ceramic blade).
* **Firearms:** `Pistol`, `Revolver`, `Rifle`, `Machine Gun`.
* *Design Note:* Weapons provide flat damage values and noise signatures during simulturn resolution. They lack complex customization trees, forcing players to rely on tactical positioning, surprise, and environmental manipulation rather than gear score escalation. Every weapon model features distinct wear patterns—from pitted rust reducing durability to blood-stained grips that attract tracking hounds—ensuring that even standard armaments demand careful field maintenance and tactical appraisal.

---

## 3. MATRIX-BASED PROCEDURAL SALVAGE & LOOT ENGINES

To maximize item variety and procedural emergent storytelling across the Vegas Valley, random salvage is determined via **Dual-Axis $6 \times 6$ Generation Matrices**. Rolling a coordinate (Row $\times$ Column) combines a foundational material state with a specific urban context or degradation vector, yielding hundreds of unique item profiles with distinct tactical advantages.

### 3.1 Hydration, Dumpster Finds & Slime Utility Matrix
*Instructions: Roll 1d6 for the Material Base (Rows) and 1d6 for the Urban Degradation State (Columns).*

| Material \ State | **1. Sun-Bleached & Rotted** | **2. Grease-Slicked & Sludgy** | **3. Rat-Chewed & Feral** | **4. Chemically Contaminated** | **5. Compressed & Crushed** | **6. Waterlogged & Moldy** |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **A. Styrofoam Cup** | Leaking foam cup; useful as a crude muzzle acoustic silencer sleeve. | Greasy foam cup containing stagnant fast-fry oil; flammable trap component. | Rat-gnawed cup base; makes a papery crunch noise when used as a tripwire. | Acid-etched foam cup; dissolves slowly to release caustic chemical drips. | Flat wafer of compressed foam; burns with dense, choking black smoke. | Soggy foam wedge; dampens footsteps or seals small air-draft cracks. |
| **B. Aluminum Can** | Sun-warped energy drink can; can be cut into sharp lock shims. | Half-full sugary stimulant syrup; gives temporary Hustle but high crash. | Punctured can infested with desert ants; useful as a distraction bait. | Corroded battery-acid can; stains hands and poisons close-combat strikes. | Flattened aluminum strip; perfect for emergency radio antenna wiring. | Sludge-filled can; heavy blunt-impact weight when swung in a sock. |
| **C. Plastic Bottle** | Crinkled 20oz flask; makes loud crunch warnings if stepped on. | Oily liquor flask; excellent accelerant for burn-barrel traps. | Chewed-through nipple cap; holds dry matches or micro-SD data cards. | Fluoride-clouded water; restores minimal hydration with minor toxicity. | Crushed flat plastic ribbon; high-tensile binding cord for trap rigging. | Algae-choked squeeze bottle; foul-tasting water that induces mild nausea. |
| **D. Cardboard Box** | Brittle pizza carton; strips tear easily for dry tinder fuel. | Grease-soaked slice box; smells strongly enough to attract tracker hounds. | Nesting box fragment lined with rodent hair; toxic allergen risk. | Solvent-soaked box paper; emits chemical fumes that bypass gas masks. | Compressed pulp brick; acts as makeshift body armor padding against blunt force. | Moldering cardboard mush; can be plastered over windows to block light. |
| **E. Metal Ice Bucket** | Brushed aluminum bucket; un-grounded Faraday shield for burner phones. | Sludge-lined bucket; heavy blunt-force `Club` with tetanus hazard. | Gnawed handle rim; sharp jagged edges useful for cutting wire traps. | Rust-scaling bucket; dissolves rapidly when exposed to battery electrolyte. | Stomp-flattened metal disc; makeshift door barricade wedge or shrapnel plate. | Water-holding basin; useful for catching overnight condensation dew. |
| **F. Plastic Pail** | Sun-cracked 2-gallon pail; acts as a dry container for electronic spares. | Chemically stained bucket; carries caustic waste for trap deployment. | Chewed rim bucket; holds loose scrap metal or heavy mechanical components. | Industrial runoff bucket; emits low-level Geiger counter clicks. | Folded plastic shell; lightweight waterproof shield against desert dust. | Silt-filled bucket; heavy dead-weight anchor for rigging subterranean traps. |

---

### 3.2 Corporate Access, Scams & Counterfeit Credentials Matrix
*Instructions: Roll 1d6 for the Document Type (Rows) and 1d6 for the Forgery/Wear State (Columns).*

| Document \ State | **1. Laminated & Faded** | **2. Hand-Marked & Altered** | **3. Photocopied Cheaply** | **4. Scratched & Demagnetized** | **5. Stolen & Blood-Stained** | **6. Expired & Obsolete** |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **A. Players Club Card** | Gold-tier loyalty card; fools low-level casino floor security. | Sharpie-altered name badge; mimics high-roller VIP status poorly. | Blurry paper card printout; works only on distracted human attendants. | Scratched mag-stripe card; triggers error buzzers on smart door locks. | Blood-spattered VIP card; induces high suspicion from corporate guards. | 1990s casino card; evokes nostalgic laughter or dismissal from old pit bosses. |
| **B. Employee Name Tag** | Plastic "CHAD" badge; grants back-of-house service hallway access. | Sharpie-written "MANAGER" overlay; temporary social immunity in breakrooms. | Paper cutout name tag taped to cardboard; fails close visual inspection. | Melted corner badge; implies you barely survived a kitchen fire accident. | Stolen dead-staff badge; links to an active security investigation alert. | Vintage 2005 resort badge; recognized only by legacy maintenance staff. |
| **C. Parking Pass** | Laminated garage pass; opens manual boom gates via sheer confidence. | Hand-drawn validation stamp; fools tired night-shift booth attendants. | Blurry photocopy stub; forces a frustrated manual hand-wave exit. | Bent barcode ticket; jams optical scanners and flashes red alarms. | Blood-smeared valet tag; linked to a stolen vehicle alert profile. | Expired 2012 festival pass; allows entry to abandoned parking structures. |
| **D. Convention Lanyard** | Tech conference fabric strap; grants entry to corporate tech lounges. | Marker-corrected "PRESS" pass; bypasses media barricades. | Paper badge insert drawn with crayons; highly suspicious up close. | Frayed nylon strap; easily snaps if yanked during physical altercations. | Stained tech-bro badge; smells of spilled craft beer and desperation. | Defunct cyber-expo pass; unlocks doors in abandoned server basements. |
| **E. Security Epaulet** | Gold-stitched uniform stripe; upgrades disguise value vs. guards. | Sharpie rank upgrade; fools junior guards but inspected by captains. | Printed paper rank insignia taped on cloth; tears in heavy wind. | Moth-eaten fabric patch; looks unprofessional and invites scrutiny. | Torn guard-sleeve piece; smells of gun oil and sweat; high realism. | Obsolete 1980s casino guard pin; grants entry to subterranean vaults. |
| **F. Transit Ticket** | Monorail RFID card; slides mechanical turnstiles open smoothly. | Punched-hole ticket stub; grants one-time passage on service lifts. | Paper bus transfer slip; worthless for trains, good for street panhandlers. | Cracked plastic transit pass; shorts out optical turnstile gates. | Crushed ticket from a crime scene; links player to transit police alerts. | Century-old rail pass; valued only by local artifact collectors and fences. |

---

### 3.3 Electronics, Hacking Implants & Radio Scrap Matrix
*Instructions: Roll 1d6 for the Hardware Component (Rows) and 1d6 for the Damage/Mod State (Columns).*

| Hardware \ State | **1. Rusted & Corroded** | **2. Water-Logged & Shorted** | **3. Stripped & Salvaged** | **4. Burned & Melted** | **5. Duct-Taped & Rigged** | **6. Cracked & Fractured** |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **A. Car Radio Faceplate** | Rusted AM/FM chassis; yields copper wire and variable capacitors. | Soaked faceplate; components fried; good only for scrap metal weight. | Stripped circuit board; harvested for micro-resistors and display LCDs. | Melted plastic housing; emits toxic smoke when heated with a torch. | Duct-taped stereo deck; wired into a crude radio signal booster. | Cracked glass display; sharp glass shards can be used as cutting tools. |
| **B. Vape Battery (18650)** | Leaking lithium cell; highly volatile and corrosive chemical risk. | Submerged battery; dead and completely inert, toxic heavy metal core. | Stripped cell casing; yields high-density lithium foil for incendiary traps. | Bloated thermal-runaway cell; risks exploding if overcharged by a rig. | Wire-wrapped dual cell pack; high-capacity portable power bank for phones. | Dented outer sleeve; sparks intermittently when jostled in a toolkit. |
| **C. Slot Bill Acceptor** | Jammed currency module; yields optical sensors and tiny solenoids. | Water-logged sensor housing; optical lenses clouded by mineral scale. | Mechanically gutted frame; heavy steel housing functions as a mini-anvil. | Singed bill validator; smells of burned paper currency and ozone. | Rigged cash stepper; tricks mechanical token chutes into false payout drops. | Cracked optical glass; creates random false-positive optical triggers. |
| **D. Smartphone Motherboard** | Mineral-crusted board; gold contact pads can be scraped off with a knife. | Fully drowned board; completely dead silicon; silver solder harvest only. | Component-plucked board; zero chips left, useful only as a rigid shim. | Charred processor traces; completely useless short-circuit hazard. | Jumper-wire Frankenboard; runs a rudimentary rogue LoRa relay script. | Shattered glass sandwich; severe laceration hazard when handled bare-handed. |
| **E. Magnet Wire Spool** | Rusted copper coil; brittle wire snaps under high-tension winding. | Mold-encrusted wire spool; insulation rotting away, prone to shorting. | Unwound copper hair strands; perfect for fine antenna array tuning. | Heat-fused wire lump; melted into a solid, useless block of copper slag. | Improvised induction loop; traps magnetic pulses from door locks. | Tangled wire bird's nest; takes 1 turn in inventory to untangle safely. |
| **F. Garage Keyfob** | Corroded circuit clicker; battery terminals covered in green crust. | Drowned plastic shell; water sloshing inside micro-switches. | Gutted button board; housing used to conceal micro-SD data cards. | Melted button pad; buttons permanently depressed, emitting continuous RF. | Rewired brute-force clicker; cycles rolling codes against automated gates. | Cracked plastic case held by rubber bands; fragile trigger mechanism. |

---

### 3.4 Souvenirs, Weird Memorabilia & Psychological Warfare Matrix
*Instructions: Roll 1d6 for the Novelty Item (Rows) and 1d6 for the Weirdness/Wear State (Columns).*

| Novelty \ State | **1. Glows in Dark** | **2. Water-Damaged** | **3. Moth-Eaten & Shedding** | **4. Melted in Sun** | **5. Glitter-Covered** | **6. Brittle & Sun-Bleached** |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **A. Plastic Skull Ice Cube** | Neon radioactive glow; excellent distraction decoy in dark tunnels. | Slime-coated skull; smells of stale bar syrup and stagnant condensation. | Fur-lined cheap tourist skull; shedding synthetic fibers everywhere. | Warped melting skull; sticks to inventory pockets like hot glue. | Glitter-encrusted skull; catches stray flashlight beams, ruining stealth. | Brittle chalky skull; crumbles instantly when thrown, creating chalk dust. |
| **B. Elvis Wedding Program** | UV-reactive ink booklet; reveals hidden UV secret safehouse codes. | Water-soaked glossy paper; pages stuck together into a solid brick. | Moth-eaten paper pages; can be torn into fine confetti security trackers. | Sun-baked curling booklet; smells of old cardboard and desert dust. | Glitter laminate cover; blinds pursuers when flashed in bright neon lights. | Brittle paper brochure; snaps in half instantly; high sentimental value. |
| **C. Inflatable Palm Toy** | Neon green-brown vinyl; inflated via straw to block air-duct drafts. | Moldy deflated palm tree; smells of stagnant indoor pool chlorine. | Shredded vinyl strips; can be tied into crude ghillie suit camouflage. | Melted plastic lump; chemical adhesive property useful for sealing gaps. | Glittery inflatable fronds; catches light to signal allied mesh operators. | Dry-rotted vinyl; shatters into sharp plastic shards when kicked. |
| **D. Golden Nugget Brick** | Gold-painted plastic weight; heavy bludgeon `Club` or paperweight. | Sand-logged plastic brick; water sloshes inside, altering balance. | Velvet-lined display box scrap; useful for cradling delicate microchips. | Sun-warped gold brick; paint flakes off to leave sticky yellow residue. | Glitter-coated nugget; distracting reflection catches guard sniper lines. | Brittle plastic block; cracks on hard impact, revealing dry desert sand. |
| **E. Aviator Sunglasses** | UV-blocking lenses; shields eyes from tactical flash-bang grenades. | Mineral-crusted lenses; severely distorts vision, increasing missed shots. | Rusted metal frames; missing one earpiece, held on by dirty string. | Heat-warped frames; lenses pop out spontaneously under extreme heat. | Rhinestone-studded frames; catches light, making stealth movement harder. | Brittle plastic frames; snap instantly if sat upon or dropped. |
| **F. Feather Boa** | Fluorescent pink boa; catches UV light for underground rave parties. | Soggy feather mass; heavy, water-logged biological weight. | Heavily shedding feathers; leaves a clear tracking trail for enemy hounds. | Heat-fused feather clump; stiff and sharp like wire brush bristles. | Glitter-bombed boa; explodes into a blinding cloud of micro-plastics when torn. | Dry brittle feathers; excellent flash-fire tinder for emergency barrels. |

---

### 3.5 Containers, Carrying Gear & Improvised Storage Matrix
*Instructions: Roll 1d6 for the Container Type (Rows) and 1d6 for the Structural Integrity State (Columns).*

| Container \ State | **1. Torn & Seam-Split** | **2. Oily & Foul-Smelling** | **3. Waterproof-Sealed** | **4. Heavy & Reinforced** | **5. Overstuffed & Bulging** | **6. Perforated & Leaking** |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **A. Mesh Laundry Bag** | Split-seam nylon sack; gear slowly leaks out during parkour sprints. | Bleach-stained bag; chemical odor masks player scent from patrol dogs. | Double-lined waterproof sack; protects electronics from flash floods. | Double-stitched linen sack; holds massive bulky loot without tearing. | Bulging laundry sack; high profile, zeroes all stealth and disguise bonuses. | Hole-ridden mesh sack; drops small items randomly while moving. |
| **B. Cardboard Liquor Box** | Crushed corrugated box; bottle dividers collapse under heavy weight. | Bourbon-soaked box; highly flammable hazard if exposed to open sparks. | Paraffin-wax coated box; completely waterproof storage for dry rations. | Double-wall liquor carton; excellent impact protection for fragile vials. | Overpacked bottle box; flaps won't close, spilling contents when tilted. | Punctured carton base; leaks loose hardware and spare ammunition rounds. |
| **C. Tyvek Courier Envelope** | Rip-tear postal envelope; waterproof paper protects vital cipher notes. | Grease-stained envelope; smells of stale corporate coffee and ink. | Heat-sealed Tyvek sleeve; submersible data storage for flash drives. | Cardboard-backed envelope; rigid enough to act as a knife-stab block. | Stuffed shipping envelope; seams balloon outward, threatening to burst. | Shredded-edge envelope; lets fine desert dust coat sensitive documents. |
| **D. Room Service Tray** | Bent stainless steel tray; serves as a crude anti-shrapnel chest plate. | Grease-caked food tray; smells of old hollandaise sauce and dead flies. | Edge-sealed tray rim; can be used as a shallow pan for boiling water. | Heavy-gauge hotel tray; durable melee bludgeon (`Club`) in close combat. | Stacked high tray tower; rattles loudly, alerting guards across the floor. | Rusted perforated tray; swiss-cheese metal sheet useless for liquid holding. |
| **E. Nylon Fanny Pack** | Jammed plastic buckle; requires safety pin to stay secured to waist. | Daiquiri-stained pouch; sticky zippers jam during urgent item draws. | Sealed waterproof pouch; keeps matches and micro-electronics bone dry. | Tactical reinforced pouch; extra MOLLE webbing for external tool clips. | Overstuffed zipper pouch; items pop out unpredictably during rolls. | Ripped pouch lining; small items slip directly into the inner jacket lining. |
| **F. Contractor Garbage Bag** | Punctured plastic film; requires heavy duct tape wrapping to hold loot. | Chemical waste bag; emits toxic fumes that inflict mild status debuffs. | Heat-sealed industrial bag; airtight vacuum shroud for body disposal. | 6mm thick contractor bag; highly puncture-resistant heavy loot sack. | Bloated trash bag; slow movement speed and high profile to guards. | Shredded plastic sheet; useful only as improvised ground-sheet insulation. |

---

### 3.6 Standardized Lethal Weapons Matrix
*Instructions: Roll 1d6 for the Weapon Class (Rows) and 1d6 for Sanitized/Condition State (Columns).*

| Weapon Class \ State | **1. Rusted & Pitted** | **2. Blood-Stained & Dirty** | **3. Field-Stripped & Clean** | **4. Improvised & Crude** | **5. Silenced & Modified** | **6. Factory-Standard** |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **A. Melee Club** | Rusted iron pipe `Club`; high tetanus infection risk on hit. | Blood-caked tire iron `Club`; terrifying appearance lowers guard morale. | Polished tactical baton `Club`; balanced weight for quick strike recovery. | Concrete-weighted rebar `Club`; heavy swing, stamina draining. | Foam-wrapped pipe `Club`; silent impact dampening for stealth takedowns. | Standard heavy flashlight `Club`; dual-purpose illumination and bludgeon. |
| **B. Melee Knife** | Rusted kitchen `Knife`; fragile blade snaps on heavy armor plates. | Stained combat `Knife`; biological residue leaves forensic tracking traces. | Razor-sharp hunting `Knife`; maximum cutting precision and stealth lethality. | Ceramic shiv `Knife`; completely undetectable by basic metal detectors. | Wrapped handle `Knife`; non-reflective matte blade coating for night ops. | Standard military combat `Knife`; durable, reliable utility blade edge. |
| **C. Revolver** | Pitted cylinder `Revolver`; high chance of casing jam in desert sand. | Blood-spattered frame `Revolver`; heavy stopping power, deafening report. | Oiled clean-action `Revolver`; smooth trigger pull and reliable lockup. | Field-filed trigger `Revolver`; hair-trigger snap with increased misfire risk. | Ported barrel `Revolver`; reduced muzzle flash at the cost of higher noise. | Standard civilian `Revolver`; balanced performance across all engagement ranges. |
| **D. Semi-Auto Pistol** | Jam-prone slide `Pistol`; requires constant oiling in dusty environments. | Grime-filled grip `Pistol`; slick handling during high-stress simulturn ticks. | Clean-feed magazine `Pistol`; flawless cycling speed and fast reload times. | Tape-wrapped frame `Pistol`; makeshift grip panels fashioned from plastic scrap. | Threaded-barrel `Pistol`; fitted with an improvised oil-filter suppressor. | Standard-issue service `Pistol`; reliable sidearm with standard magazine size. |
| **E. Hunting Rifle** | Warped barrel `Rifle`; long-range accuracy degraded by environmental rust. | Mud-caked stock `Rifle`; heavy weight slows down movement across grid tiles. | Zeroed-scope `Rifle`; pin-point accuracy for elevated sniper vantage points. | Sawed-off barrel `Rifle`; wide spread, reduced range, massive noise profile. | Suppressed long `Rifle`; zero muzzle flash, ultra-quiet sub-sonic rounds. | Standard bolt-action `Rifle`; high-velocity stopping power at long distances. |
| **F. Machine Gun** | Jammed feed-tray gun; prone to catastrophic ammunition stovepipes. | Battle-worn receiver gun; smokes heavily during sustained suppressive fire. | Clean gas-piston gun; consistent high-rate cyclic fire without overheating. | Drum-mag welded gun; excessive weight penalty but massive ammo reserve. | Suppressed tactical gun; reduced muzzle signature during full auto bursts. | Standard military squad gun; ultimate suppressive firepower for open combat. |

---

## 4. FIELD CRAFTING & SCRAP FABRICATION SYSTEM

### 4.1 Crafting Stations & Workbench Tiers
* **Tier 1: Field Jury-Rig (Inventory-Only)**
  * *Requirements:* Hands only, 1 turn of exposed time.
  * *Capability:* Combines basic consumables, wraps melee weapons in tape, or creates loud acoustic noisemakers and simple distraction baits. High failure rate if interrupted by environmental hazards or patrol sightings.
* **Tier 2: Burn-Barrel Workbench (Safehouse / Alleyway)**
  * *Requirements:* Empty oil drum or casino trash bin, makeshift tools (pliers, screwdriver, wire cutters).
  * *Capability:* Fabricates electrical mods, antenna arrays, chemical traps, and rudimentary suppressor sleeves. Moderate failure rate depending on player skill checks and workshop ambient lighting conditions.
* **Tier 3: Deep-Sublevel Electronics Bench (Encrypted Hideout)**
  * *Requirements:* Soldering iron, stolen digital multimeter, stable low-voltage power supply generator.
  * *Capability:* Assembles complex LoRa mesh nodes, reprograms casino access keyfobs, builds custom SDR surveillance gear, and services advanced encrypted comms hardware. Zero random failure rate under stable power.
* **Tier 4: Automated Industrial Fabricator (Corporate Vault)**
  * *Requirements:* Reprogrammed CNC milling terminal, high-purity chemical baths, heavy electrical grid connection.
  * *Capability:* Manufactures high-precision replacement parts, military-grade armor plates, and complex cryptographic processing hardware. Requires rare clearance tokens to unlock operational privileges.

### 4.2 Core Crafting Recipes & Blueprints
| Recipe Name | Required Components | Workbench Tier | Functional Result & Tactical Effect |
| :--- | :--- | :--- | :--- |
| **Styrofoam Muzzle Sleeve** | Cracked Styrofoam Cup (Base A) + Electrical Tape | Tier 1 (Field) | Slips over a `Pistol` or `Revolver`. Reduces the acoustic report signature by 50% for 3 shots before the foam completely disintegrates. |
| **Acoustic Tripwire Alarm** | Crinkled Plastic Bottle (Base C) + Metal Wire + Tin Cans | Tier 1 (Field) | Rigs a noisy perimeter warning line across hotel doorways or hallway thresholds; alerts players to enemy movement within 2 grid tiles. |
| **Glow-Decoy Flare** | Glow-in-the-Dark Skull (Base A) + Flare Residue | Tier 1 (Field) | Creates a false thermal/optical distraction marker that tricks automated security turrets and guard line-of-sight checks for 1 turn. |
| **LoRa Signal Booster** | Car Radio Faceplate (Base A) + Aluminum Can Strip (Base B) | Tier 2 (Burn Barrel) | Extends burner phone mesh networking range by +3 grid sectors, allowing encrypted packet routing through concrete casino walls. |
| **Dual-Cell Power Bank** | 2x Vape Batteries (18650, Base B) + Wire Strippings | Tier 2 (Burn Barrel) | Restores full power to dead tactical flashlights, encrypted keypads, or SDR dongles during extended subterranean operations. |
| **Cloned Mag-Stripe Key** | Blank Plastic Keycard + Rigged Keyfob Board (Base F) | Tier 3 (Sublevel) | Mimics active casino room key data, bypassing electronic turnstiles and back-of-house service doors without triggering lockdown alarms. |
| **Improvised Napalm Trap** | Styrofoam Cup (Base A) + Oily Liquor Flask (Base C) | Tier 2 (Burn Barrel) | Dissolves foam into sticky, highly volatile incendiary sludge; creates a high-damage area-denervation fire hazard when ignited in hallways. |
| **Acidic Corrosive Drip** | Corroded Ice Bucket (Base E) + Battery Acid Scrap | Tier 2 (Burn Barrel) | Deploys a slow-release chemical drip that weakens mechanical door hinges or dissolves lock mechanisms over 2 turn cycles. |
| **Waterproof Data Shroud** | Tyvek Courier Envelope (Base C) + Heat-Sealer / Lighter | Tier 1 (Field) | Encapsulates micro-SD memory cards, cryptographic keys, and sensitive paper ciphers in an airtight, submersible protective barrier. |
| **Improvised EMP Grenade** | Vape Battery (18650) + Copper Coil + Flash Bulb | Tier 3 (Sublevel) | Emits a localized electromagnetic shockwave that temporarily disables security cameras and electronic locks within a 3-tile radius. |
| **Spiked Caltrop Cluster** | Cut Aluminum Cans (Base B) + Heavy Wire Binding | Tier 1 (Field) | Scatters jagged metal triangles across entry corridors, inflicting crippling movement speed debuffs on pursuing guard units or tracker hounds. |
| **Filtered Charcoal Respirator** | Plastic Bottle (Base C) + Crushed Carbon Filters + Cloth | Tier 2 (Burn Barrel) | Provides temporary immunity to tear gas, chemical runoff vapors, and toxic sewer miasma during subterranean exploration phases. |

### 4.3 Scrap Salvage Conversion Values & Advanced Processing
When items cannot be used directly, they can be broken down at any Tier 2 workbench into raw salvage currencies used in advanced fabrication:
* **Scrap Polymer (Plastics / Foam / Vinyl):** Yielded from cups, bottles, keyfob housings, and plastic pails. Used for casings, insulation, and lightweight structural shims.
* **Scrap Conductive Metal (Copper / Aluminum):** Yielded from soda cans, car radios, and magnet wire. Used for antenna coils, electrical wiring, and jumper leads.
* **Hazardous Chem-Solvents (Acid / Lithium / Oily Residue):** Yielded from vape batteries, corroded buckets, and chemical run-offs. Used for traps, incendiaries, and power generation.
* **Refined Silicate Glass (Lenses / Displays / Shards):** Yielded from broken optical scanners, smartphone screens, and neon tubing. Used for optical triggers, solar charging nodes, and precision cutting tools.
* **Heavy Structural Steel (Frames / Trays / Bars):** Yielded from ice buckets, cash boxes, and shelving units. Used for heavy melee reinforcements, door barricade wedges, and structural framing.
* **Organic Fibers & Textiles (Linen / Canvas / Cordage):** Yielded from laundry bags, clothing scraps, and convention lanyards. Used for binding ropes, acoustic dampening wraps, and primitive camouflage netting.