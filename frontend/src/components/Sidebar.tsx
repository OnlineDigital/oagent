import { For, Show, createMemo, createSignal, onCleanup, onMount } from "solid-js";
import { ProjectInfo, SessionInfo } from "../../bindings/changeme";

type Conversation = {
  id: string;
  parentId: string;
  projectId: string;
  project: string;
  initials: string;
  title: string;
  detail: string;
  status: "active" | "idle";
  activityAt: number;
  active: boolean;
};

function basename(path: string): string {
  if (!path) return "";
  const normalized = path.replace(/[\\/]+$/, "");
  const parts = normalized.split(/[\\/]/);
  return parts[parts.length - 1] || normalized;
}

function initialsFor(name: string): string {
  const clean = name.trim();
  if (!clean) return "??";
  const words = clean.split(/[\s_-]+/).filter(Boolean);
  if (words.length >= 2) {
    return (words[0][0] + words[1][0]).toUpperCase();
  }
  return clean.slice(0, 2).toUpperCase();
}

function relativeTime(timestamp: number): string {
  if (!timestamp) return "—";
  const diff = Math.max(0, Date.now() - timestamp);
  const minutes = Math.floor(diff / 60000);
  if (minutes < 1) return "now";
  if (minutes < 60) return `${minutes}m`;
  const hours = Math.floor(minutes / 60);
  if (hours < 24) return `${hours}h`;
  const days = Math.floor(hours / 24);
  return `${days}d`;
}

const FILTER_STORAGE_KEY = "oagent.sidebar.filter";

function readFilter(): string {
  try {
    const raw = localStorage.getItem(FILTER_STORAGE_KEY);
    return raw || "all";
  } catch {
    return "all";
  }
}

function persistFilter(value: string) {
  try {
    localStorage.setItem(FILTER_STORAGE_KEY, value);
  } catch {
    // Ignore storage failures; the filter just won't persist.
  }
}

function sortConversations(items: Conversation[]): Conversation[] {
  return [...items].sort((a, b) => {
    if (a.active !== b.active) return a.active ? -1 : 1;
    return b.activityAt - a.activityAt;
  });
}

function FolderIcon() {
  return (
    <svg width="13" height="13" viewBox="0 0 16 16" fill="none" aria-hidden="true">
      <path
        d="M1.5 3.5A1.5 1.5 0 0 1 3 2h3l1.5 2H13a1.5 1.5 0 0 1 1.5 1.5v6A1.5 1.5 0 0 1 13 13H3a1.5 1.5 0 0 1-1.5-1.5v-8z"
        fill="currentColor"
        fill-rule="evenodd"
        clip-rule="evenodd"
      />
    </svg>
  );
}

function ChevronIcon() {
  return (
    <svg width="12" height="12" viewBox="0 0 16 16" fill="none" aria-hidden="true">
      <path
        d="M4 6l4 4 4-4"
        stroke="currentColor"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  );
}

function CheckIcon() {
  return (
    <svg width="12" height="12" viewBox="0 0 16 16" fill="none" aria-hidden="true">
      <path
        d="M3 8l3 3 7-7"
        stroke="currentColor"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  );
}

function BranchIcon() {
  return (
    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" aria-hidden="true">
      <line x1="6" y1="3" x2="6" y2="15" stroke="currentColor" stroke-width="2" />
      <circle cx="18" cy="6" r="3" stroke="currentColor" stroke-width="2" />
      <circle cx="6" cy="18" r="3" stroke="currentColor" stroke-width="2" />
      <path d="M18 9a9 9 0 0 1-9 9" stroke="currentColor" stroke-width="2" />
    </svg>
  );
}

function AgentIcon() {
  return (
    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" aria-hidden="true">
      <path
        d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  );
}

export default function Sidebar(props: {
  width: number;
  projects: ProjectInfo[];
  sessions: SessionInfo[];
  selectedId: string | null;
  loading: boolean;
  error: string;
  onSelect: (id: string) => void;
  onRetry: () => void;
}) {
  const [filter, setFilter] = createSignal(readFilter());
  const [filterOpen, setFilterOpen] = createSignal(false);
  const [filterPos, setFilterPos] = createSignal({ top: 0, left: 0 });

  const projectName = createMemo(() => {
    const map = new Map<string, string>();
    for (const project of props.projects) {
      const name = project.name || basename(project.canonical) || "Untitled";
      map.set(project.id, name);
    }
    return map;
  });

  const conversations = createMemo<Conversation[]>(() =>
    props.sessions.map((session) => {
      const project = projectName().get(session.projectId) || basename(session.directory || "") || "Untitled";
      return {
        id: session.id,
        parentId: session.parentId || "",
        projectId: session.projectId,
        project,
        initials: initialsFor(project),
        title: session.title || session.agent || "Untitled session",
        detail: basename(session.directory || ""),
        status: session.active ? "active" : "idle",
        activityAt: session.updatedAt || session.createdAt,
        active: session.active,
      };
    }),
  );

  const subagentsFor = createMemo(() => {
    const map = new Map<string, Conversation[]>();
    for (const conversation of conversations()) {
      if (!conversation.parentId) continue;
      const list = map.get(conversation.parentId) || [];
      list.push(conversation);
      map.set(conversation.parentId, list);
    }
    for (const list of map.values()) {
      list.sort((a, b) => a.activityAt - b.activityAt);
    }
    return map;
  });

  const filterOptions = createMemo(() => {
    const seen = new Map<string, string>();
    for (const conversation of conversations()) {
      if (conversation.parentId) continue;
      if (!seen.has(conversation.projectId)) {
        seen.set(conversation.projectId, conversation.project);
      }
    }
    return [...seen.entries()].map(([id, name]) => ({ id, name }));
  });

  const visible = createMemo(() => {
    const selected = filter();
    const matches = conversations().filter((item) => {
      if (item.parentId) return false;
      if (selected === "all") return true;
      if (selected === "active") return item.active;
      if (selected === "done") return !item.active;
      return item.projectId === selected;
    });
    return sortConversations(matches);
  });

  const filterLabel = createMemo(() => {
    const selected = filter();
    if (selected === "all") return "All projects";
    if (selected === "active") return "Active";
    if (selected === "done") return "Idle";
    const option = filterOptions().find((item) => item.id === selected);
    return option ? option.name : selected;
  });

  function handleEscape(e: KeyboardEvent) {
    if (e.key === "Escape" && filterOpen()) {
      setFilterOpen(false);
    }
  }

  onMount(() => {
    document.addEventListener("keydown", handleEscape);
  });

  onCleanup(() => {
    document.removeEventListener("keydown", handleEscape);
  });

  function toggleFilter(e: MouseEvent) {
    const rect = (e.currentTarget as HTMLElement).getBoundingClientRect();
    setFilterPos({ top: rect.bottom + 4, left: rect.left });
    setFilterOpen((open) => !open);
  }

  function chooseFilter(next: string) {
    setFilter(next);
    persistFilter(next);
    setFilterOpen(false);
  }

  return (
    <aside class="sidebar" style={{ "--sidebar-w": `${props.width}px` }}>
      <div class="sidebar-header">
        <span class="sidebar-header-title">Conversations</span>
        <button
          class="conv-filter-btn"
          onClick={toggleFilter}
          aria-label="Filter conversations"
          aria-expanded={filterOpen()}
        >
          <FolderIcon />
          <span class="conv-filter-label">{filterLabel()}</span>
          <span class="conv-filter-chevron" classList={{ "is-open": filterOpen() }}>
            <ChevronIcon />
          </span>
        </button>
      </div>

      <Show when={filterOpen()}>
        <div class="conv-filter-backdrop" onClick={() => setFilterOpen(false)} />
        <div
          class="conv-filter-popup"
          style={{ top: `${filterPos().top}px`, left: `${filterPos().left}px` }}
        >
          <button
            class="conv-filter-option"
            classList={{ "is-selected": filter() === "all" }}
            onClick={() => chooseFilter("all")}
          >
            <span>All projects</span>
            <Show when={filter() === "all"}>
              <CheckIcon />
            </Show>
          </button>
          <button
            class="conv-filter-option"
            classList={{ "is-selected": filter() === "active" }}
            onClick={() => chooseFilter("active")}
          >
            <span>Active</span>
            <Show when={filter() === "active"}>
              <CheckIcon />
            </Show>
          </button>
          <button
            class="conv-filter-option"
            classList={{ "is-selected": filter() === "done" }}
            onClick={() => chooseFilter("done")}
          >
            <span>Idle</span>
            <Show when={filter() === "done"}>
              <CheckIcon />
            </Show>
          </button>
          <Show when={filterOptions().length > 0}>
            <div class="conv-filter-divider" />
          </Show>
          <For each={filterOptions()}>
            {(option) => (
              <button
                class="conv-filter-option"
                classList={{ "is-selected": filter() === option.id }}
                onClick={() => chooseFilter(option.id)}
              >
                <span>{option.name}</span>
                <Show when={filter() === option.id}>
                  <CheckIcon />
                </Show>
              </button>
            )}
          </For>
        </div>
      </Show>

      <div class="conv-list">
        <Show when={props.loading}>
          <div class="conv-empty">Loading conversations…</div>
        </Show>

        <Show when={!props.loading && props.error}>
          <div class="conv-empty">
            <div class="conv-error-msg">{props.error}</div>
            <button class="btn-ghost conv-retry" onClick={props.onRetry}>
              Retry
            </button>
          </div>
        </Show>

        <Show when={!props.loading && !props.error}>
          <For each={visible()}>
            {(conversation) => (
              <>
                <button
                  class="conv-item"
                  classList={{ "is-active": props.selectedId === conversation.id }}
                  onClick={() => props.onSelect(conversation.id)}
                  aria-label={`${conversation.title}, ${conversation.project}`}
                  aria-pressed={props.selectedId === conversation.id}
                >
                  <div class="conv-row">
                    <span class="conv-project">
                      <span class="conv-project-badge">{conversation.initials}</span>
                      <span class="conv-project-name">{conversation.project}</span>
                    </span>
                    <span
                      class="conv-status"
                      classList={{
                        "is-working": conversation.status === "active",
                        "is-done": conversation.status === "idle",
                      }}
                    >
                      <span class="conv-status-dot" />
                      {conversation.status === "active" ? "Active" : "Idle"} ·{" "}
                      {relativeTime(conversation.activityAt)}
                    </span>
                  </div>

                  <div class="conv-row">
                    <span class="conv-title">{conversation.title}</span>
                  </div>

                  <div class="conv-row">
                    <span class="conv-branch">
                      <BranchIcon />
                      <span class="conv-branch-name">
                        {conversation.detail || conversation.project}
                      </span>
                    </span>
                    <span class="conv-meta">
                      <span class="conv-thread">
                        <AgentIcon />
                        <span>{conversation.status === "active" ? "Live" : "Recent"}</span>
                      </span>
                    </span>
                  </div>
                </button>

                <Show when={subagentsFor().get(conversation.id)?.length}>
                  <div class="conv-subagents">
                    <For each={subagentsFor().get(conversation.id)!}>
                      {(subagent) => (
                        <button
                          class="conv-subagent"
                          classList={{ "is-active": props.selectedId === subagent.id }}
                          onClick={() => props.onSelect(subagent.id)}
                          aria-label={`${subagent.title}, subagent`}
                          aria-pressed={props.selectedId === subagent.id}
                        >
                          <span class="conv-subagent-rail" />
                          <AgentIcon />
                          <span class="conv-subagent-title">{subagent.title || "Subagent"}</span>
                          <span
                            class="conv-subagent-dot"
                            classList={{
                              "is-working": subagent.status === "active",
                              "is-done": subagent.status === "idle",
                            }}
                          />
                        </button>
                      )}
                    </For>
                  </div>
                </Show>
              </>
            )}
          </For>

          <Show when={visible().length === 0}>
            <div class="conv-empty">No conversations match this filter.</div>
          </Show>
        </Show>
      </div>
    </aside>
  );
}
