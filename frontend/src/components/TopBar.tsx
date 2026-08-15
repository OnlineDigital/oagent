import HarnessButton from "./HarnessButton";

export default function TopBar(props: {
  onToggleSidebar: () => void;
  onTogglePanel: () => void;
}) {
  return (
    <header class="topbar">
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
    </header>
  );
}
