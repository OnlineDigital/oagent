import { createSignal, For, Show } from "solid-js";

type PanelTab = "browser" | "git";

const TABS: { id: PanelTab; label: string }[] = [
  { id: "browser", label: "Browser" },
  { id: "git", label: "Git" },
];

const GIT_FILES = [
  { path: "frontend/src/App.tsx", status: "M" as const, lines: "+12 -4" },
  { path: "frontend/src/components/Sidebar.tsx", status: "A" as const, lines: "+86" },
  { path: "opencodeservice.go", status: "M" as const, lines: "+3 -1" },
  { path: "frontend/public/style.css", status: "M" as const, lines: "+210 -180" },
];

export default function RightPanel(props: { width: number }) {
  const [activeTab, setActiveTab] = createSignal<PanelTab>(readActiveTab());

  function readActiveTab(): PanelTab {
    try {
      const raw = localStorage.getItem("oagent.panel.tab");
      return raw === "git" ? "git" : "browser";
    } catch {
      return "browser";
    }
  }

  function selectTab(tab: PanelTab) {
    setActiveTab(tab);
    try {
      localStorage.setItem("oagent.panel.tab", tab);
    } catch {
      // Ignore storage failures.
    }
  }

  return (
    <aside class="right-panel" style={{ "--panel-w": `${props.width}px` }}>
      <div class="right-panel-tabs">
        <For each={TABS}>
          {(tab) => (
            <button
              class="right-panel-tab"
              classList={{ "is-active": activeTab() === tab.id }}
              onClick={() => selectTab(tab.id)}
              aria-selected={activeTab() === tab.id}
              role="tab"
            >
              {tab.label}
            </button>
          )}
        </For>
      </div>

      <div class="right-panel-body">
        <Show when={activeTab() === "browser"}>
          <div class="browser-pane">
            <div class="browser-toolbar">
              <button class="browser-btn" aria-label="Back">
                <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
                  <path
                    d="M9.5 3L5 8l4.5 5"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  />
                </svg>
              </button>
              <button class="browser-btn" aria-label="Forward">
                <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
                  <path
                    d="M6.5 3L11 8l-4.5 5"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  />
                </svg>
              </button>
              <button class="browser-btn" aria-label="Reload">
                <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
                  <path
                    d="M14 8a6 6 0 1 1-1.8-4.3M14 2.5V6h-3.5"
                    stroke="currentColor"
                    stroke-width="1.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  />
                </svg>
              </button>
              <div class="browser-url">https://localhost:3000</div>
            </div>
            <div class="browser-content">
              <div class="browser-empty">
                <svg width="32" height="32" viewBox="0 0 32 32" fill="none">
                  <rect
                    x="4"
                    y="5"
                    width="24"
                    height="19"
                    rx="2"
                    stroke="currentColor"
                    stroke-width="1.5"
                  />
                  <path
                    d="M4 9h24"
                    stroke="currentColor"
                    stroke-width="1.5"
                  />
                  <circle cx="7" cy="7" r="0.8" fill="currentColor" />
                  <circle cx="9.5" cy="7" r="0.8" fill="currentColor" />
                  <circle cx="12" cy="7" r="0.8" fill="currentColor" />
                </svg>
                <span>Browser panel</span>
              </div>
            </div>
          </div>
        </Show>

        <Show when={activeTab() === "git"}>
          <div class="git-pane">
            <div class="git-header">
              <span class="git-header-title">Changes</span>
              <span class="git-branch">main</span>
            </div>
            <ul class="git-file-list">
              <For each={GIT_FILES}>
                {(file) => (
                  <li class="git-file-item">
                    <span
                      class="git-file-status"
                      classList={{
                        "status-modified": file.status === "M",
                        "status-added": file.status === "A",
                      }}
                    >
                      {file.status}
                    </span>
                    <div class="git-file-main">
                      <span class="git-file-path">{file.path}</span>
                      <span class="git-file-lines">{file.lines}</span>
                    </div>
                  </li>
                )}
              </For>
            </ul>
            <div class="git-actions">
              <button class="btn-primary" disabled>
                Commit
              </button>
              <button class="btn-ghost" disabled>
                Push
              </button>
            </div>
          </div>
        </Show>
      </div>
    </aside>
  );
}
