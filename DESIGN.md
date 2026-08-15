# Design System: OAgent

<!-- SEED: established with the user before implementation; re-run $impeccable document once there's code to capture the actual tokens and components. -->

## Overview

**Creative North Star: "Precision Instruments"**

OAgent arată ca un instrument de precizie pentru orchestrare: clean, dens, fără decorațiuni. Dark mode nativ, panouri definite prin contrast subtil de suprafață, un singur accent proeminent care marchează acțiunea și starea activă. Inspirat de VS Code (ierarhie de panouri, densitate, familiaritate) și Warp (accent luminos, mono pentru date, tranziții rapide).

Interfața dispare în task: developerul vede agenți, task-uri, proiecte, git și browser fără să observe design-ul. Fiecare pixel are un job. Accent-ul apare doar acolo unde trebuie să acționezi sau să observi o stare.

**Key Characteristics:**
- Dark, precis, dens
- Un singur accent proeminent (default cyan), configurabil
- Mono pentru date, sans pentru UI
- Panouri colapsabile cu tranziții rapide (150–200ms)
- Fără gradient text, fără glassmorphism decorativ, fără umbre inutile

## Colors

Dark, cu accent cyan. Contrast minim 4.5:1 pentru text, suprafețele se disting prin tonalitate, nu prin umbre.

### Primary
- **Accent Cyan** (#00d4ff): acțiuni primare, selecție activă, starea „online". Folosit cu economie — max 10% din orice ecran.

### Neutral
- **Ink** (#0d1117): fundal principal
- **Panel** (#161b22): sidebar, panouri, topbar
- **Panel Elevated** (#1c2128): elemente hover, popup-uri, carduri
- **Border** (#30363d): hairlines, separatori
- **Text** (#e6edf3): text principal
- **Muted** (#8b949e): text secundar
- **Faint** (#484f58): placeholder, disabled

### Named Rules
**The One Accent Rule.** Accent-ul cyan apare doar pe acțiune primară, selecție curentă și starea activă. Restul interfeței e neutral.

## Typography

**Body Font:** System UI stack (-apple-system, "Segoe UI", Roboto)
**Mono Font:** ui-monospace, "Cascadia Code", "SF Mono", Consolas

**Character:** Sans pentru UI, mono pentru date și măsurători. Contrast clar între cele două, fără pairing decorativ.

### Hierarchy
- **Title** (700, 20px, 1.2): titluri de panou, heading-uri principale
- **Body** (400, 13px, 1.5): text normal
- **Label** (600, 11px, 0.04em, uppercase): etichete de secțiune, tab-uri
- **Mono** (400, 12px, 1.4): date, path-uri, status, măsurători

## Layout

Grid strict de panouri: topbar 44px, sidebar stânga 240px, panou drept 340px, conținut central flexibil. Colapsarea panourilor e structurală — sidebar-ul și panoul dispar complet, conținutul central preia spațiul. Spacing pe o scară de 4px: 4, 8, 12, 16, 24, 32.

## Elevation & Depth

Fără umbre. Adâncimea e transmisă prin tonalitate: Ink → Panel → Panel Elevated. Fiecare nivel de suprafață are culoarea lui. Un singur border (1px #30363d) separă suprafețele. Hover-ul ridică tonalitatea, nu adaugă umbră.

## Shapes

Colțuri mici, precise: 6px pentru controale mici (butoane, tab-uri), 8px pentru panouri și carduri. Pills doar pentru badge-uri de status. Fără colțuri rotunde mari, fără forme organice.

## Components

### Buttons
- **Shape:** 6px radius
- **Primary:** background accent, text ink, padding 6px 14px
- **Hover / Focus:** brighten accent, focus ring 2px accent la 50% opacity
- **Secondary / Ghost:** transparent cu border 1px #30363d, text muted, hover text

### Navigation
- Sidebar: iteme cu icon + label, hover tonal, activ cu accent subtil în stânga
- Topbar: butoane icon transparente, hover tonal

### Panels
- Background Panel, border 1px #30363d, fără umbre
- Colapsare: panoul dispare, conținutul central se extinde

## Do's and Don'ts

### Do:
- **Do** folosi accent cyan doar pentru acțiune, selecție, stare activă
- **Do** folosi mono pentru path-uri, status, date
- **Do** păstra densitatea VS Code — informație cât mai multă, dar scanabilă
- **Do** folosi tonalitate (nu umbre) pentru ierarhie de suprafețe

### Don't:
- **Don't** folosi gradient text sau glassmorphism decorativ
- **Don't** adăuga umbre sau glow
- **Don't** folosi mai mult de un accent color
- **Don't** pune decorațiuni pe panouri — fiecare pixel are un job
