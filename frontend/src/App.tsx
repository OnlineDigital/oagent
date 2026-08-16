import { createSignal, onCleanup, onMount, Show } from "solid-js";
import TopBar from "./components/TopBar";
import Sidebar from "./components/Sidebar";
import RightPanel from "./components/RightPanel";
import ConversationDemo from "./components/ConversationDemo";

const SIDEBAR_OPEN_KEY = "oagent.sidebar.open";
const SIDEBAR_WIDTH_KEY = "oagent.sidebar.width";
const PANEL_OPEN_KEY = "oagent.panel.open";
const PANEL_WIDTH_KEY = "oagent.panel.width";

function readBoolean(key: string, fallback: boolean): boolean {
  try {
    const raw = localStorage.getItem(key);
    return raw === null ? fallback : raw === "true";
  } catch {
    return fallback;
  }
}

function readNumber(key: string, fallback: number): number {
  try {
    const raw = localStorage.getItem(key);
    return raw === null ? fallback : Number(raw);
  } catch {
    return fallback;
  }
}

export default function App() {
  const [sidebarOpen, setSidebarOpen] = createSignal(readBoolean(SIDEBAR_OPEN_KEY, true));
  const [panelOpen, setPanelOpen] = createSignal(readBoolean(PANEL_OPEN_KEY, true));
  const [sidebarWidth, setSidebarWidth] = createSignal(readNumber(SIDEBAR_WIDTH_KEY, 240));
  const [panelWidth, setPanelWidth] = createSignal(readNumber(PANEL_WIDTH_KEY, 340));

  function persistLayout() {
    try {
      localStorage.setItem(SIDEBAR_OPEN_KEY, String(sidebarOpen()));
      localStorage.setItem(PANEL_OPEN_KEY, String(panelOpen()));
      localStorage.setItem(SIDEBAR_WIDTH_KEY, String(sidebarWidth()));
      localStorage.setItem(PANEL_WIDTH_KEY, String(panelWidth()));
    } catch {
      // Ignore storage failures; layout just won't persist.
    }
  }

  function toggleSidebar() {
    setSidebarOpen(!sidebarOpen());
    persistLayout();
  }

  function togglePanel() {
    setPanelOpen(!panelOpen());
    persistLayout();
  }

  function handleKeydown(e: KeyboardEvent) {
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === "b") {
      e.preventDefault();
      toggleSidebar();
    }
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === "j") {
      e.preventDefault();
      togglePanel();
    }
  }

  onMount(() => {
    window.addEventListener("keydown", handleKeydown);
  });

  function startSidebarResize(e: PointerEvent) {
    e.preventDefault();
    const startX = e.clientX;
    const startWidth = sidebarWidth();
    const minWidth = 160;
    const maxWidth = Math.min(520, window.innerWidth * 0.6);

    function onMove(ev: PointerEvent) {
      const next = Math.min(maxWidth, Math.max(minWidth, startWidth + (ev.clientX - startX)));
      setSidebarWidth(next);
    }

    function onUp() {
      window.removeEventListener("pointermove", onMove);
      window.removeEventListener("pointerup", onUp);
      persistLayout();
      document.body.classList.remove("is-resizing");
    }

    window.addEventListener("pointermove", onMove);
    window.addEventListener("pointerup", onUp);
    document.body.classList.add("is-resizing");
  }

  function startPanelResize(e: PointerEvent) {
    e.preventDefault();
    const startX = e.clientX;
    const startWidth = panelWidth();
    const minWidth = 200;
    const maxWidth = Math.min(560, window.innerWidth * 0.6);

    function onMove(ev: PointerEvent) {
      const next = Math.min(maxWidth, Math.max(minWidth, startWidth - (ev.clientX - startX)));
      setPanelWidth(next);
    }

    function onUp() {
      window.removeEventListener("pointermove", onMove);
      window.removeEventListener("pointerup", onUp);
      persistLayout();
      document.body.classList.remove("is-resizing");
    }

    window.addEventListener("pointermove", onMove);
    window.addEventListener("pointerup", onUp);
    document.body.classList.add("is-resizing");
  }

  onCleanup(() => {
    window.removeEventListener("keydown", handleKeydown);
  });

  return (
    <div class="app-shell">
      <TopBar onToggleSidebar={toggleSidebar} onTogglePanel={togglePanel} />

      <div class="app-body">
        <Show when={sidebarOpen()}>
          <Sidebar width={sidebarWidth()} />
          <div
            class="resize-handle resize-handle-left"
            role="separator"
            aria-orientation="vertical"
            aria-label="Resize sidebar"
            onPointerDown={(e) => startSidebarResize(e)}
          />
        </Show>

        <main class="main-content">
          <ConversationDemo />
        </main>

        <Show when={panelOpen()}>
          <div
            class="resize-handle resize-handle-right"
            role="separator"
            aria-orientation="vertical"
            aria-label="Resize right panel"
            onPointerDown={(e) => startPanelResize(e)}
          />
          <RightPanel width={panelWidth()} />
        </Show>
      </div>
    </div>
  );
}
