# Design System: OAgent

<!-- SEED: established with the user before implementation; re-run $impeccable document once there's code to capture the actual tokens and components. -->

## Overview

**Creative North Star: "Precision Instruments"**

OAgent looks like a precision instrument for orchestration: clean, dense, no decoration. Native dark mode, panels defined through subtle surface contrast, a single prominent accent that marks action and active state. Inspired by VS Code (panel hierarchy, density, familiarity) and Warp (bright accent, mono for data, fast transitions).

The interface disappears into the task: the developer sees agents, tasks, projects, git, and browser without noticing the design. Every pixel has a job. The accent appears only where you need to act or observe a state.

**Key Characteristics:**
- Dark, precise, dense
- A single prominent accent (default cyan), configurable
- Mono for data, sans for UI
- Collapsible panels with fast transitions (150–200ms)
- No text gradients, no decorative glassmorphism, no unnecessary shadows

## Colors

Dark, with cyan accent. Minimum 4.5:1 contrast for text; surfaces are distinguished by tonality, not shadows.

### Primary
- **Accent Cyan** (#00d4ff): primary actions, active selection, "online" state. Used sparingly — max 10% of any screen.

### Neutral
- **Ink** (#0d1117): main background
- **Panel** (#161b22): sidebar, panels, topbar
- **Panel Elevated** (#1c2128): hover elements, popups, cards
- **Border** (#30363d): hairlines, separators
- **Text** (#e6edf3): main text
- **Muted** (#8b949e): secondary text
- **Faint** (#484f58): placeholder, disabled

### Named Rules
**The One Accent Rule.** The cyan accent appears only on primary action, current selection, and active state. The rest of the interface is neutral.

## Typography

**Body Font:** System UI stack (-apple-system, "Segoe UI", Roboto)
**Mono Font:** ui-monospace, "Cascadia Code", "SF Mono", Consolas

**Character:** Sans for UI, mono for data and measurements. Clear contrast between the two, no decorative pairing.

### Hierarchy
- **Title** (700, 20px, 1.2): panel titles, main headings
- **Body** (400, 13px, 1.5): normal text
- **Label** (600, 11px, 0.04em, uppercase): section labels, tabs
- **Mono** (400, 12px, 1.4): data, paths, status, measurements

## Layout

Strict panel grid: topbar 44px, left sidebar 240px, right panel 340px, flexible central content. Panel collapsing is structural — the sidebar and panel disappear completely, and the central content takes the space. Spacing on a 4px scale: 4, 8, 12, 16, 24, 32.

## Elevation & Depth

No shadows. Depth is conveyed through tonality: Ink → Panel → Panel Elevated. Each surface level has its own color. A single border (1px #30363d) separates surfaces. Hover raises tonality, not shadow.

## Shapes

Small, precise corners: 6px for small controls (buttons, tabs), 8px for panels and cards. Pills only for status badges. No large rounded corners, no organic shapes.

## Components

### Buttons
- **Shape:** 6px radius
- **Primary:** accent background, ink text, 6px 14px padding
- **Hover / Focus:** brighten accent, 2px focus ring accent at 50% opacity
- **Secondary / Ghost:** transparent with 1px #30363d border, muted text, hover text

### Navigation
- Sidebar: items with icon + label, tonal hover, active with subtle accent on the left
- Topbar: transparent icon buttons, tonal hover

### Panels
- Panel background, 1px #30363d border, no shadows
- Collapse: panel disappears, central content expands

## Do's and Don'ts

### Do:
- **Do** use cyan accent only for action, selection, active state
- **Do** use mono for paths, status, data
- **Do** keep VS Code density — as much information as possible, but scannable
- **Do** use tonality (not shadows) for surface hierarchy

### Don't:
- **Don't** use text gradients or decorative glassmorphism
- **Don't** add shadows or glow
- **Don't** use more than one accent color
- **Don't** put decorations on panels — every pixel has a job
