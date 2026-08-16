import { createSignal, onCleanup, onMount, Show } from "solid-js";
import { Events, System, Window } from "@wailsio/runtime";
import HarnessButton from "./HarnessButton";

export default function TopBar(props: {
  onToggleSidebar: () => void;
  onTogglePanel: () => void;
}) {
  const [maximised, setMaximised] = createSignal(false);
  const [isMac, setIsMac] = createSignal(false);

  onMount(() => {
    try {
      setIsMac(System.IsMac());
    } catch {
      // Not running inside Wails (e.g. browser dev server).
    }
    Window.IsMaximised()
      .then(setMaximised)
      .catch(() => {});
    const offMax = Events.On("common:WindowMaximise", () => setMaximised(true));
    const offUnmax = Events.On("common:WindowUnMaximise", () => setMaximised(false));
    onCleanup(() => {
      offMax();
      offUnmax();
    });
  });

  function minimise() {
    Window.Minimise().catch(() => {});
  }

  async function toggleMaximise() {
    try {
      await Window.ToggleMaximise();
      setMaximised(await Window.IsMaximised());
    } catch {
      // Not running inside Wails (e.g. browser dev server).
    }
  }

  function closeWindow() {
    Window.Close().catch(() => {});
  }

  function handleDblClick(e: MouseEvent) {
    if ((e.target as HTMLElement).closest("button")) return;
    toggleMaximise();
  }

  return (
    <header class="topbar" onDblClick={handleDblClick}>
      <button
        class="topbar-icon-btn"
        onClick={props.onToggleSidebar}
        aria-label="Toggle sidebar"
        title="Toggle sidebar (Ctrl+B)"
      >
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <path
            d="M2 3.5h12M2 8h12M2 12.5h12"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
          />
        </svg>
      </button>

      <div class="topbar-brand">
        <span class="topbar-brand-mark" aria-hidden="true">
          <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
            <circle cx="8" cy="8" r="6.5" stroke="currentColor" stroke-width="1.5" />
            <circle cx="8" cy="8" r="2" fill="currentColor" />
          </svg>
        </span>
        <span class="topbar-brand-name">OAgent</span>
      </div>

      <div class="topbar-divider" />

      <div class="topbar-spacer" />

      <HarnessButton />

      <div class="topbar-divider" />

      <button
        class="topbar-icon-btn"
        onClick={props.onTogglePanel}
        aria-label="Toggle right panel"
        title="Toggle right panel (Ctrl+J)"
      >
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <rect
            x="8.5"
            y="1.5"
            width="6"
            height="13"
            rx="1"
            stroke="currentColor"
            stroke-width="1.5"
          />
          <rect
            x="1.5"
            y="1.5"
            width="6"
            height="13"
            rx="1"
            stroke="currentColor"
            stroke-width="1.5"
          />
        </svg>
      </button>

      <Show when={!isMac()}>
        <div class="window-controls">
          <button
            class="window-control-btn"
            onClick={minimise}
            aria-label="Minimize"
            title="Minimize"
          >
            <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
              <path d="M3 8.5h10" stroke="currentColor" />
            </svg>
          </button>
          <button
            class="window-control-btn"
            onClick={toggleMaximise}
            aria-label={maximised() ? "Restore" : "Maximize"}
            title={maximised() ? "Restore" : "Maximize"}
          >
            <Show
              when={!maximised()}
              fallback={
                <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
                  <path d="M9.5 9.5h2v-7h-7v2" stroke="currentColor" />
                  <rect x="2.5" y="4.5" width="7" height="7" stroke="currentColor" />
                </svg>
              }
            >
              <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
                <rect x="3.5" y="3.5" width="9" height="9" stroke="currentColor" />
              </svg>
            </Show>
          </button>
          <button
            class="window-control-btn window-control-close"
            onClick={closeWindow}
            aria-label="Close"
            title="Close"
          >
            <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
              <path d="M3 3l10 10M13 3L3 13" stroke="currentColor" />
            </svg>
          </button>
        </div>
      </Show>
    </header>
  );
}
