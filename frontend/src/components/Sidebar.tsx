import { For, createSignal } from "solid-js";

type NavSection = "agents" | "tasks" | "projects";

const NAV_ITEMS: { section: NavSection; label: string; icon: string }[] = [
  {
    section: "agents",
    label: "Agenți",
    icon: "M8 8a2.5 2.5 0 1 0 0-5 2.5 2.5 0 0 0 0 5zm0 1.5c-2.67 0-8 1.34-8 4v1.5h16V13.5c0-2.66-5.33-4-8-4z",
  },
  {
    section: "tasks",
    label: "Task-uri",
    icon: "M5 2h6a1 1 0 0 1 1 1v1h2a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h1V3a1 1 0 0 1 1-1zm2 2h4V3H7v1zm-3 2v7h8V6H4z",
  },
  {
    section: "projects",
    label: "Proiecte",
    icon: "M1.5 3.5A1.5 1.5 0 0 1 3 2h3l1.5 2H13a1.5 1.5 0 0 1 1.5 1.5v6A1.5 1.5 0 0 1 13 13H3a1.5 1.5 0 0 1-1.5-1.5v-8z",
  },
];

const AGENTS = [
  { id: "agent-1", name: "Agent Alpha", status: "running" as const, task: "Refactor API" },
  { id: "agent-2", name: "Agent Beta", status: "idle" as const, task: null },
  { id: "agent-3", name: "Agent Gamma", status: "running" as const, task: "Fix bug #42" },
];

const TASKS = [
  { id: "task-1", name: "Refactor API", status: "in-progress" as const, agent: "Agent Alpha" },
  { id: "task-2", name: "Fix bug #42", status: "in-progress" as const, agent: "Agent Gamma" },
  { id: "task-3", name: "Deploy staging", status: "queued" as const, agent: null },
];

const PROJECTS = [
  { id: "project-1", name: "OAgent", branch: "main", changes: 3 },
  { id: "project-2", name: "Whales", branch: "feature/setup", changes: 7 },
  { id: "project-3", name: "CLI Tools", branch: "main", changes: 0 },
];

export default function Sidebar() {
  const [activeSection, setActiveSection] = createSignal<NavSection>("agents");

  return (
    <aside class="sidebar">
      <div class="sidebar-section-label">Workspace</div>

      <nav class="sidebar-nav">
        <For each={NAV_ITEMS}>
          {(item) => (
            <button
              class="sidebar-nav-btn"
              classList={{ "is-active": activeSection() === item.section }}
              onClick={() => setActiveSection(item.section)}
              aria-label={item.label}
              aria-current={activeSection() === item.section ? "page" : undefined}
            >
              <svg
                width="16"
                height="16"
                viewBox="0 0 16 16"
                fill="none"
                class="sidebar-nav-icon"
              >
                <path
                  d={item.icon}
                  fill="currentColor"
                  fill-rule="evenodd"
                  clip-rule="evenodd"
                />
              </svg>
              <span class="sidebar-nav-label">{item.label}</span>
            </button>
          )}
        </For>
      </nav>

      <div class="sidebar-divider" />

      <div class="sidebar-section-label">
        {activeSection() === "agents"
          ? "Agenți"
          : activeSection() === "tasks"
            ? "Task-uri"
            : "Proiecte"}
      </div>

      <div class="sidebar-items">
        <ShowSection section={activeSection()} />
      </div>
    </aside>
  );
}

function ShowSection(props: { section: NavSection }) {
  switch (props.section) {
    case "agents":
      return (
        <For each={AGENTS}>
          {(agent) => (
            <button class="sidebar-item" aria-label={agent.name}>
              <span
                class="agent-status-dot"
                classList={{
                  "dot-online": agent.status === "running",
                  "dot-idle": agent.status === "idle",
                }}
              />
              <span class="sidebar-item-main">
                <span class="sidebar-item-name">{agent.name}</span>
                {agent.task && (
                  <span class="sidebar-item-sub">{agent.task}</span>
                )}
              </span>
            </button>
          )}
        </For>
      );
    case "tasks":
      return (
        <For each={TASKS}>
          {(task) => (
            <button class="sidebar-item" aria-label={task.name}>
              <span
                class="task-status-dot"
                classList={{
                  "dot-online": task.status === "in-progress",
                  "dot-idle": task.status === "queued",
                }}
              />
              <span class="sidebar-item-main">
                <span class="sidebar-item-name">{task.name}</span>
                {task.agent && (
                  <span class="sidebar-item-sub">{task.agent}</span>
                )}
              </span>
            </button>
          )}
        </For>
      );
    case "projects":
      return (
        <For each={PROJECTS}>
          {(project) => (
            <button class="sidebar-item" aria-label={project.name}>
              <span class="sidebar-item-main">
                <span class="sidebar-item-name">{project.name}</span>
                <span class="sidebar-item-sub">
                  {project.branch}
                  {project.changes > 0 && ` · ${project.changes} changes`}
                </span>
              </span>
            </button>
          )}
        </For>
      );
  }
}
