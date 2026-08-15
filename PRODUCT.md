# Product

<!-- impeccable:product-schema 1 -->

## Platform

web

## Users

Dezvoltatori solo care orchestrează agenți AI din desktop pentru task-uri de codare pe proiecte locale.

## Product Purpose

OAgent este un orchestrator desktop care permite unui singur developer să lanseze, monitorizeze și coordoneze agenți AI (prin harnesse precum OpenCode2) care execută task-uri de codare. Succesul înseamnă: agenții rulează, task-urile progresează, iar developerul are control și vizibilitate totală din aceeași interfață.

## Positioning

OAgent orchestrează agenți AI locali prin API-uri harness (OpenCode2 V2), nu prin servicii cloud. Developerul își păstrează controlul asupra mediului local, git-ului și browser-ului din aceeași aplicație desktop.

## Operating Context

- Desktop app (Wails v3 + SolidJS), rulând local pe Windows
- Developerul lucrează cu terminal, git, editor de cod și browser în paralel
- Agenții AI rulează ca servicii locale (OpenCode2) și comunică prin API
- Task-urile sunt legate de proiecte git locale
- Interfața trebuie să fie scanabilă și rapidă — developerul comută des între ea și editor

## Capabilities and Constraints

- Sidebar stânga: navigare între agenți, task-uri, proiecte (colapsabil)
- Top bar: toggle sidebar, toggle panou drept, buton harnesse cu status (OpenCode2)
- Panou drept: tab-uri Browser și Git (colapsabil), extensibil cu funcții viitoare
- Harness OpenCode2: verificare status (bulina verde/roșie) și setup când e offline
- Backend: Wails service OpenCodeService (IsReady, Setup)
- Datele din sidebar și panoul drept sunt momentan hardcodate

## Brand Commitments

- Numele produsului este OAgent
- Direcție vizuală: clean, minimalist, modern, high-tech
- Accent color proeminent, configurabil (default cyan)
- Nivel de craft aliniat cu VS Code / Warp — precis, dens, fără decorațiuni
- Tema dark, potrivită pentru utilizare la birou, seara

## Evidence on Hand

- Frontend SolidJS funcțional cu layout shell (TopBar, Sidebar, RightPanel)
- Binding-uri Wails generate: OpenCodeService.IsReady / Setup
- Backend Go cu serviciul OpenCodeService
- Nicio imagistică de brand, logo sau asset vizual autorizat

## Product Principles

1. Vizibilitate totală: starea agenților și task-urilor e vizibilă dintr-o privire
2. Control local: developerul deține mediul, git-ul și browser-ul din aceeași interfață
3. Rapid și dens: informația e scanabilă, nu decorativă
4. Extensibil: panoul drept și sidebar-ul pot găzdui funcții noi fără refactor de layout

## Accessibility & Inclusion

Interfața trebuie să fie utilizabilă cu tastatura (focus visible, tab order logic), contrast minim 4.5:1 pentru text normal, și să nu depindă exclusiv de culoare pentru transmiterea stării (bulina de status are și text).
