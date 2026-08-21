import { createSignal, createMemo, onCleanup, onMount, Show } from "solid-js";
import TopBar from "./components/TopBar";
import Sidebar from "./components/Sidebar";
import RightPanel from "./components/RightPanel";
import Conversation from "./components/Conversation";
import {
  ConversationMessage,
  OpenCodeService,
  ProjectInfo,
  SessionInfo,
  SubagentInfo,
} from "../bindings/changeme";

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

function basename(path: string): string {
  if (!path) return "";
  const normalized = path.replace(/[\\/]+$/, "");
  const parts = normalized.split(/[\\/]/);
  return parts[parts.length - 1] || normalized;
}

export default function App() {
  const [sidebarOpen, setSidebarOpen] = createSignal(readBoolean(SIDEBAR_OPEN_KEY, true));
  const [panelOpen, setPanelOpen] = createSignal(readBoolean(PANEL_OPEN_KEY, true));
  const [sidebarWidth, setSidebarWidth] = createSignal(readNumber(SIDEBAR_WIDTH_KEY, 240));
  const [panelWidth, setPanelWidth] = createSignal(readNumber(PANEL_WIDTH_KEY, 340));

  const [projects, setProjects] = createSignal<ProjectInfo[]>([]);
  const [sessions, setSessions] = createSignal<SessionInfo[]>([]);
  const [selectedId, setSelectedId] = createSignal<string | null>(null);
  const [messages, setMessages] = createSignal<ConversationMessage[]>([]);
  const [subagents, setSubagents] = createSignal<SubagentInfo[]>([]);
  const [nextCursor, setNextCursor] = createSignal("");
  const [hasMoreHistory, setHasMoreHistory] = createSignal(false);
  const [loadingOlder, setLoadingOlder] = createSignal(false);
  const [listLoading, setListLoading] = createSignal(true);
  const [listError, setListError] = createSignal("");
  const [detailLoading, setDetailLoading] = createSignal(false);
  const [detailError, setDetailError] = createSignal("");

  let historyToken = 0;

  const projectName = createMemo(() => {
    const map = new Map<string, string>();
    for (const project of projects()) {
      map.set(project.id, project.name || basename(project.canonical) || "Untitled");
    }
    return map;
  });

  const selectedSession = createMemo(() => {
    const session = sessions().find((item) => item.id === selectedId());
    if (!session) return null;
    return {
      id: session.id,
      title: session.title || session.agent || "Untitled session",
      project: projectName().get(session.projectId) || basename(session.directory || "") || "Untitled",
      agent: session.agent,
      directory: session.directory,
      active: session.active,
    };
  });

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

  async function loadList() {
    setListLoading(true);
    setListError("");
    try {
      const [projectList, sessionList] = await Promise.all([
        OpenCodeService.Projects(),
        OpenCodeService.Sessions(100),
      ]);
      setProjects(projectList ?? []);
      setSessions(sessionList ?? []);
      if (!selectedId() && sessionList?.length) {
        setSelectedId(sessionList[0].id);
      }
    } catch (err) {
      console.error(err);
      setListError(String(err));
    } finally {
      setListLoading(false);
    }
  }

  async function loadDetail(id: string) {
    if (!id) {
      setMessages([]);
      setSubagents([]);
      setNextCursor("");
      setHasMoreHistory(false);
      return;
    }

    const token = ++historyToken;
    setDetailLoading(true);
    setDetailError("");
    setLoadingOlder(false);
    setMessages([]);
    setNextCursor("");
    setHasMoreHistory(false);
    try {
      const [page, subagentList] = await Promise.all([
        OpenCodeService.ConversationPage(id, ""),
        OpenCodeService.Subagents(id),
      ]);
      if (token !== historyToken) return;
      setMessages(page.messages ?? []);
      setNextCursor(page.nextCursor ?? "");
      setHasMoreHistory(page.hasMore);
      setSubagents(subagentList ?? []);
    } catch (err) {
      if (token !== historyToken) return;
      console.error(err);
      setDetailError(String(err));
    } finally {
      if (token === historyToken) {
        setDetailLoading(false);
      }
    }
  }

  async function loadOlder() {
    const id = selectedId();
    if (!id || loadingOlder() || !hasMoreHistory()) return;

    const token = historyToken;
    setLoadingOlder(true);
    try {
      const page = await OpenCodeService.ConversationPage(id, nextCursor());
      if (token !== historyToken) return;
      const older = page.messages ?? [];
      setMessages((current) => {
        const known = new Set(current.map((m) => m.id));
        const added = older.filter((m) => !known.has(m.id));
        return [...added, ...current];
      });
      setNextCursor(page.nextCursor ?? "");
      setHasMoreHistory(page.hasMore);
    } catch (err) {
      if (token !== historyToken) return;
      console.error(err);
      setDetailError(String(err));
    } finally {
      if (token === historyToken) {
        setLoadingOlder(false);
      }
    }
  }

  async function loadAllHistory() {
    const token = historyToken;
    while (token === historyToken && hasMoreHistory()) {
      const before = nextCursor();
      await loadOlder();
      if (nextCursor() === before) break;
    }
  }

  function selectSession(id: string) {
    if (id === selectedId()) return;
    setSelectedId(id);
    loadDetail(id).then(() => {
      loadAllHistory();
    });
  }

  onMount(() => {
    window.addEventListener("keydown", handleKeydown);
    loadList().then(() => {
      const id = selectedId();
      if (id) {
        loadDetail(id).then(() => {
          loadAllHistory();
        });
      }
    });
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
          <Sidebar
            width={sidebarWidth()}
            projects={projects()}
            sessions={sessions()}
            selectedId={selectedId()}
            loading={listLoading()}
            error={listError()}
            onSelect={selectSession}
            onRetry={loadList}
          />
          <div
            class="resize-handle resize-handle-left"
            role="separator"
            aria-orientation="vertical"
            aria-label="Resize sidebar"
            onPointerDown={(e) => startSidebarResize(e)}
          />
        </Show>

        <main class="main-content">
          <Conversation
            session={selectedSession()}
            messages={messages()}
            subagents={subagents()}
            loading={detailLoading()}
            loadingOlder={loadingOlder()}
            hasMoreHistory={hasMoreHistory()}
            onLoadOlder={loadOlder}
            error={detailError()}
          />
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
