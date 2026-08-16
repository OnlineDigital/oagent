import { For, Show, createMemo, createSignal, onCleanup, onMount } from "solid-js";

type Conversation = {
  id: string;
  project: string;
  initials: string;
  title: string;
  branch: string;
  status: "working" | "done";
  durationMin: number;
  lastActivityMin: number;
  thread: string;
  pinned: boolean;
};

const CONVERSATIONS: Conversation[] = [
  {
    id: "conv-142",
    project: "OAgent",
    initials: "OA",
    title: "Refactor sidebar into a project-aware inbox",
    branch: "oagent/sidebar-inbox",
    status: "working",
    durationMin: 14,
    lastActivityMin: 4,
    thread: "#142",
    pinned: true,
  },
  {
    id: "conv-138",
    project: "OAgent",
    initials: "OA",
    title: "Fix agent task scheduling race",
    branch: "oagent/fix-task-scheduling",
    status: "working",
    durationMin: 7,
    lastActivityMin: 9,
    thread: "#138",
    pinned: false,
  },
  {
    id: "conv-131",
    project: "OAgent",
    initials: "OA",
    title: "Add browser preview to the right panel",
    branch: "oagent/browser-preview",
    status: "done",
    durationMin: 22,
    lastActivityMin: 25,
    thread: "#131",
    pinned: true,
  },
  {
    id: "conv-119",
    project: "OAgent",
    initials: "OA",
    title: "Wire OpenCode2 setup status",
    branch: "oagent/opencode-setup",
    status: "done",
    durationMin: 31,
    lastActivityMin: 47,
    thread: "#119",
    pinned: false,
  },
  {
    id: "conv-104",
    project: "CLI Tools",
    initials: "CL",
    title: "Build the auth command for the CLI",
    branch: "cli-tools/feature/auth",
    status: "working",
    durationMin: 53,
    lastActivityMin: 12,
    thread: "#104",
    pinned: false,
  },
  {
    id: "conv-097",
    project: "CLI Tools",
    initials: "CL",
    title: "Migrate config to TOML",
    branch: "cli-tools/feature/toml",
    status: "done",
    durationMin: 41,
    lastActivityMin: 68,
    thread: "#097",
    pinned: false,
  },
  {
    id: "conv-088",
    project: "Whales",
    initials: "WH",
    title: "Refactor the Whales API client",
    branch: "whales/feature/api-client",
    status: "working",
    durationMin: 12,
    lastActivityMin: 18,
    thread: "#088",
    pinned: false,
  },
  {
    id: "conv-076",
    project: "Whales",
    initials: "WH",
    title: "Add parser tests",
    branch: "whales/feature/parser-tests",
    status: "done",
    durationMin: 19,
    lastActivityMin: 82,
    thread: "#076",
    pinned: false,
  },
];

function sortConversations(items: Conversation[]): Conversation[] {
  return [...items].sort((a, b) => {
    if (a.status !== b.status) return a.status === "working" ? -1 : 1;
    return a.lastActivityMin - b.lastActivityMin;
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

function ThreadIcon() {
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

function PinIcon() {
  return (
    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" aria-hidden="true">
      <path
        d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
      <circle cx="12" cy="10" r="3" stroke="currentColor" stroke-width="2" />
    </svg>
  );
}

export default function Sidebar(props: { width: number }) {
  const [filter, setFilter] = createSignal("all");
  const [filterOpen, setFilterOpen] = createSignal(false);
  const [filterPos, setFilterPos] = createSignal({ top: 0, left: 0 });
  const [selectedId, setSelectedId] = createSignal("conv-142");

  const projects = createMemo(() => [...new Set(CONVERSATIONS.map((item) => item.project))]);

  const visible = createMemo(() => {
    const selected = filter();
    const matches = CONVERSATIONS.filter((item) => {
      if (selected === "all") return true;
      if (selected === "active") return item.status === "working";
      if (selected === "done") return item.status === "done";
      return item.project === selected;
    });
    return sortConversations(matches);
  });

  const filterLabel = createMemo(() => {
    const selected = filter();
    if (selected === "all") return "All projects";
    if (selected === "active") return "Active";
    if (selected === "done") return "Done";
    return selected;
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
            <span>Done</span>
            <Show when={filter() === "done"}>
              <CheckIcon />
            </Show>
          </button>
          <div class="conv-filter-divider" />
          <For each={projects()}>
            {(project) => (
              <button
                class="conv-filter-option"
                classList={{ "is-selected": filter() === project }}
                onClick={() => chooseFilter(project)}
              >
                <span>{project}</span>
                <Show when={filter() === project}>
                  <CheckIcon />
                </Show>
              </button>
            )}
          </For>
        </div>
      </Show>

      <div class="conv-list">
        <For each={visible()}>
          {(conversation) => (
            <button
              class="conv-item"
              classList={{ "is-active": selectedId() === conversation.id }}
              onClick={() => setSelectedId(conversation.id)}
              aria-label={`${conversation.title}, ${conversation.project}`}
              aria-pressed={selectedId() === conversation.id}
            >
              <div class="conv-row">
                <span class="conv-project">
                  <span class="conv-project-badge">{conversation.initials}</span>
                  <span class="conv-project-name">{conversation.project}</span>
                </span>
                <span
                  class="conv-status"
                  classList={{
                    "is-working": conversation.status === "working",
                    "is-done": conversation.status === "done",
                  }}
                >
                  <span class="conv-status-dot" />
                  {conversation.status === "working" ? "Working" : "Done"} ·{" "}
                  {conversation.durationMin}m
                </span>
              </div>

              <div class="conv-row">
                <span class="conv-title">{conversation.title}</span>
                <span class="conv-time">{conversation.lastActivityMin}m</span>
              </div>

              <div class="conv-row">
                <span class="conv-branch">
                  <BranchIcon />
                  <span class="conv-branch-name">{conversation.branch}</span>
                </span>
                <span class="conv-meta">
                  <span class="conv-thread">
                    <ThreadIcon />
                    <span>{conversation.thread}</span>
                  </span>
                  <Show when={conversation.pinned}>
                    <span class="conv-pin" aria-label="Pinned">
                      <PinIcon />
                    </span>
                  </Show>
                </span>
              </div>
            </button>
          )}
        </For>

        <Show when={visible().length === 0}>
          <div class="conv-empty">
            No conversations match this filter.
          </div>
        </Show>
      </div>
    </aside>
  );
}
