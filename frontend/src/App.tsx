import { createSignal, onCleanup, onMount, Show } from "solid-js";
import TopBar from "./components/TopBar";
import Sidebar from "./components/Sidebar";
import RightPanel from "./components/RightPanel";

export default function App() {
  const [sidebarOpen, setSidebarOpen] = createSignal(true);
  const [panelOpen, setPanelOpen] = createSignal(true);

  function handleKeydown(e: KeyboardEvent) {
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === "b") {
      e.preventDefault();
      setSidebarOpen(!sidebarOpen());
    }
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === "j") {
      e.preventDefault();
      setPanelOpen(!panelOpen());
    }
  }

  onMount(() => {
    window.addEventListener("keydown", handleKeydown);
  });

  onCleanup(() => {
    window.removeEventListener("keydown", handleKeydown);
  });

  return (
    <div class="app-shell">
      <TopBar
        onToggleSidebar={() => setSidebarOpen(!sidebarOpen())}
        onTogglePanel={() => setPanelOpen(!panelOpen())}
      />

      <div class="app-body">
        <Show when={sidebarOpen()}>
          <Sidebar />
        </Show>

        <main class="main-content">
          <div class="main-header">
            <h1 class="main-title">Orchestrare</h1>
            <p class="main-subtitle">
              Coordonează agenți, task-uri și proiecte dintr-o singură interfață.
            </p>
          </div>

          <section class="main-cards">
            <div class="stat-card">
              <span class="stat-value">3</span>
              <span class="stat-label">Agenți activi</span>
            </div>
            <div class="stat-card">
              <span class="stat-value">2</span>
              <span class="stat-label">Task-uri în curs</span>
            </div>
            <div class="stat-card">
              <span class="stat-value">3</span>
              <span class="stat-label">Proiecte</span>
            </div>
          </section>
        </main>

        <Show when={panelOpen()}>
          <RightPanel />
        </Show>
      </div>
    </div>
  );
}
