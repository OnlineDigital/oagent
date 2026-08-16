import { For, Show, createSignal, onCleanup, onMount } from "solid-js";

type ToolStatus = "running" | "done" | "error";
type ToolKind = "list_dir" | "grep" | "patch" | "bash" | "read_file";

type ToolTemplate = {
  id: string;
  kind: ToolKind;
  label: string;
  args: string;
  duration: string;
  diff?: boolean;
  output: string[];
};

const TOOL_TEMPLATES: ToolTemplate[] = [
  {
    id: "tc-list",
    kind: "list_dir",
    label: "read_directory",
    args: "frontend/src",
    duration: "38ms",
    output: [
      "components/",
      "  TopBar.tsx",
      "  Sidebar.tsx",
      "  RightPanel.tsx",
      "  HarnessButton.tsx",
      "App.tsx",
      "main.tsx",
      "auth.ts",
      "vite-env.d.ts",
    ],
  },
  {
    id: "tc-grep",
    kind: "grep",
    label: "grep",
    args: "auth|session|token",
    duration: "52ms",
    output: [
      "frontend/src/auth.ts:3    export function saveSession(",
      "frontend/src/auth.ts:12   export function loadSession(",
      "frontend/src/App.tsx:31   const session = loadSession();",
    ],
  },
  {
    id: "tc-patch",
    kind: "patch",
    label: "edit_file",
    args: "frontend/src/auth.ts",
    duration: "46ms",
    diff: true,
    output: [
      "-  localStorage.setItem(TOKEN_KEY, token);",
      "+  localStorage.setItem(TOKEN_KEY, token);",
      "+  window.dispatchEvent(new Event(\"oagent.session.changed\"));",
    ],
  },
  {
    id: "tc-bash",
    kind: "bash",
    label: "bash",
    args: "npx tsc --noEmit",
    duration: "1.8s",
    output: [
      "$ npx tsc --noEmit",
      "frontend/src/auth.ts:12:5 - error TS2304: Cannot find name 'token'.",
      "Found 1 error.",
    ],
  },
];

const SUBAGENT_TOOLS: ToolTemplate[] = [
  {
    id: "sub-read-1",
    kind: "read_file",
    label: "read_file",
    args: "frontend/src/auth.ts",
    duration: "24ms",
    output: [
      "const TOKEN_KEY = \"oagent.session\";",
      "export function saveSession(token: string) {",
      "  localStorage.setItem(TOKEN_KEY, token);",
      "}",
      "export function loadSession(): string | null {",
      "  return localStorage.getItem(TOKEN_KEY);",
      "}",
    ],
  },
  {
    id: "sub-grep-1",
    kind: "grep",
    label: "grep",
    args: "auth|session|token  *.go",
    duration: "41ms",
    output: [
      "opencodeservice.go:32  func (s *OpenCodeService) IsReady() OpenCodeStatus {",
      "opencodeservice.go:42  if out == \"\" || out == \"stopped\" {",
      "opencodeservice.go:50  func (s *OpenCodeService) Setup() OpenCodeStatus {",
    ],
  },
];

const ALL_TOOLS = new Map(
  [...TOOL_TEMPLATES, ...SUBAGENT_TOOLS].map((tool) => [tool.id, tool]),
);

const USER_TEXT =
  "Find where authentication is implemented and explain the flow end to end.";

const MAIN_THINKING =
  "I should not answer from memory. I'll list the frontend structure, grep for auth terms, and spawn an explore subagent to map the full frontend-to-backend handshake before drawing a conclusion.";

const SUBAGENT_THINKING =
  "I need to trace the flow from the login form to token storage, then out to the Go service. I'll read the auth store first and confirm the backend handshake.";

const ASSISTANT_TEXT =
  "Authentication is client-side and local: the login form stores a session token in localStorage, the app reads it on startup, and the backend OpenCodeService verifies the local opencode2 service with IsReady() / Setup().";

const ASSISTANT_POINTS = [
  { before: "Token lives at ", code: "oagent.session", after: "." },
  { before: "No refresh-token path exists yet.", code: "", after: "" },
  {
    before: "Status is checked once on mount via the Wails binding.",
    code: "",
    after: "",
  },
];

function ToolGlyph(props: { kind: ToolKind }) {
  const paths: Record<ToolKind, string> = {
    list_dir:
      "M1.5 3.5A1.5 1.5 0 0 1 3 2h3l1.5 2H13a1.5 1.5 0 0 1 1.5 1.5v6A1.5 1.5 0 0 1 13 13H3a1.5 1.5 0 0 1-1.5-1.5v-8z",
    grep: "M7 1.5a5.5 5.5 0 1 0 0 11c1.32 0 2.54-.46 3.48-1.24l3.13 3.13 1.28-1.28-3.13-3.13A5.5 5.5 0 0 0 7 1.5zm0 1.5a4 4 0 1 1 0 8 4 4 0 0 1 0-8z",
    patch:
      "M3 2h7l3 3v9H3V2zm7 1.5V5h1.5L10 3.5zM4 3v10h8V6h-3V3H4zm1 2h5v1H5V5zm0 2h5v1H5V7zm0 2h3v1H5V9z",
    bash:
      "M2 4h12v8H2V4zm1 1v6h10V5H3zm1 1h2v1H4V6zm4 1l2.5 1.5L8 10V7z",
    read_file:
      "M4 2h5l3 3v9H4V2zm5 1.5V5h1.5L9 3.5zM5 3v10h6V6H9V3H5zm1 2h5v1H6V5zm0 2h5v1H6V7zm0 2h3v1H6V9z",
  };
  return (
    <svg width="14" height="14" viewBox="0 0 16 16" fill="none" aria-hidden="true">
      <path d={paths[props.kind]} fill="currentColor" fill-rule="evenodd" clip-rule="evenodd" />
    </svg>
  );
}

function Chevron(props: { open: boolean }) {
  return (
    <svg
      width="12"
      height="12"
      viewBox="0 0 16 16"
      fill="none"
      class="trace-chevron"
      classList={{ "is-open": props.open }}
      aria-hidden="true"
    >
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

function ThinkingBlock(props: {
  text: string;
  open: boolean;
  onToggle: () => void;
  active?: boolean;
  compact?: boolean;
}) {
  return (
    <div
      class="trace-thinking"
      classList={{ "is-active": props.active, "is-compact": props.compact }}
    >
      <button
        class="trace-thinking-toggle"
        onClick={props.onToggle}
        aria-expanded={props.open}
      >
        <span class="trace-thinking-dot" />
        <span class="trace-thinking-label">Thinking</span>
        <span class="trace-thinking-spacer" />
        <Chevron open={props.open} />
      </button>
      <Show when={props.open}>
        <p class="trace-thinking-text">{props.text}</p>
      </Show>
    </div>
  );
}

function ToolCallCard(props: {
  call: ToolTemplate;
  open: boolean;
  onToggle: () => void;
  status: ToolStatus;
  output: string[];
  nested?: boolean;
}) {
  return (
    <div
      class="trace-tool"
      classList={{ "is-nested": props.nested, "is-error": props.status === "error" }}
    >
      <button
        class="trace-tool-toggle"
        onClick={props.onToggle}
        aria-expanded={props.open}
      >
        <span
          class="trace-tool-icon"
          classList={{
            "is-running": props.status === "running",
            "is-done": props.status === "done",
            "is-error": props.status === "error",
          }}
        >
          <ToolGlyph kind={props.call.kind} />
        </span>
        <span class="trace-tool-name">{props.call.label}</span>
        <span class="trace-tool-args">{props.call.args}</span>
        <span class="trace-tool-spacer" />
        <span class="trace-tool-duration">
          {props.status === "running" ? "…" : props.call.duration}
        </span>
        <span class="trace-tool-status">
          {props.status === "running"
            ? "Running"
            : props.status === "done"
              ? "Done"
              : "Error"}
        </span>
        <Chevron open={props.open} />
      </button>
      <Show when={props.open}>
        <div class="trace-tool-output">
          <For each={props.output}>
            {(line) => (
              <span
                class="trace-tool-line"
                classList={{
                  "is-add": props.call.diff && line.startsWith("+"),
                  "is-remove": props.call.diff && line.startsWith("-"),
                  "is-command": props.call.kind === "bash" && line.startsWith("$"),
                }}
              >
                {line}
              </span>
            )}
          </For>
        </div>
      </Show>
    </div>
  );
}

function SubagentPanel(props: {
  open: boolean;
  onToggle: () => void;
  status: "running" | "done";
  thinkingOpen: boolean;
  onThinkingToggle: () => void;
  thinkingActive: boolean;
  thinkingText: string;
  toolIds: string[];
  resultVisible: boolean;
  toolStatus: (id: string) => ToolStatus;
  toolOutput: (id: string) => string[];
  toolOpen: (id: string) => boolean;
  onToolToggle: (id: string) => void;
}) {
  return (
    <div class="trace-subagent">
      <div class="trace-subagent-header">
        <span class="trace-subagent-glyph" aria-hidden="true">
          <svg width="14" height="14" viewBox="0 0 16 16" fill="none">
            <rect x="2" y="2" width="5" height="5" rx="1" stroke="currentColor" stroke-width="1.5" />
            <rect x="9" y="2" width="5" height="5" rx="1" stroke="currentColor" stroke-width="1.5" />
            <rect x="2" y="9" width="5" height="5" rx="1" stroke="currentColor" stroke-width="1.5" />
            <rect x="9" y="9" width="5" height="5" rx="1" stroke="currentColor" stroke-width="1.5" />
          </svg>
        </span>
        <span class="trace-subagent-name">explore</span>
        <span class="trace-subagent-kind">subagent</span>
        <span class="trace-subagent-spacer" />
        <span class="trace-subagent-meta">4 steps · 1m 12s</span>
        <span
          class="trace-subagent-status"
          classList={{ "is-running": props.status === "running" }}
        >
          <span class="trace-subagent-status-dot" />
          {props.status === "running" ? "Working" : "Done"}
        </span>
        <button
          class="trace-subagent-toggle"
          onClick={props.onToggle}
          aria-expanded={props.open}
          aria-label={props.open ? "Collapse subagent" : "Expand subagent"}
        >
          <Chevron open={props.open} />
        </button>
      </div>

      <Show when={props.open}>
        <div class="trace-subagent-purpose">
          Map the authentication flow across frontend and backend.
        </div>
        <div class="trace-subagent-steps">
          <Show when={props.thinkingText.length > 0}>
            <ThinkingBlock
              text={props.thinkingText}
              open={props.thinkingOpen}
              onToggle={props.onThinkingToggle}
              active={props.thinkingActive}
              compact
            />
          </Show>
          <For each={props.toolIds}>
            {(id) => {
              const call = ALL_TOOLS.get(id)!;
              return (
                <ToolCallCard
                  call={call}
                  open={props.toolOpen(id)}
                  onToggle={() => props.onToolToggle(id)}
                  status={props.toolStatus(id)}
                  output={props.toolOutput(id)}
                  nested
                />
              );
            }}
          </For>
        </div>
        <Show when={props.resultVisible}>
          <div class="trace-subagent-result">
            <div class="trace-subagent-result-label">
              <svg width="12" height="12" viewBox="0 0 16 16" fill="none" aria-hidden="true">
                <path
                  d="M13 3H5.5L2 6.5 5.5 10H13a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zM2 6.5V12"
                  stroke="currentColor"
                  stroke-width="1.5"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </svg>
              Result returned to OAgent
            </div>
            <p class="trace-subagent-result-text">
              Found a client-side token flow:{" "}
              <span class="trace-mono">Login → auth store → OpenCodeService.IsReady/Setup</span>.
              The backend validates through the local <span class="trace-mono">opencode2</span> CLI.
              No refresh-token path exists yet.
            </p>
          </div>
        </Show>
      </Show>
    </div>
  );
}

export default function ConversationDemo() {
  const [userVisible, setUserVisible] = createSignal(false);

  const [mainThinkingOpen, setMainThinkingOpen] = createSignal(false);
  const [mainThinkingActive, setMainThinkingActive] = createSignal(false);
  const [mainThinkingText, setMainThinkingText] = createSignal("");

  const [mainToolIds, setMainToolIds] = createSignal<string[]>([]);

  const [subagentVisible, setSubagentVisible] = createSignal(false);
  const [subagentOpen, setSubagentOpen] = createSignal(false);
  const [subagentStatus, setSubagentStatus] = createSignal<"running" | "done">("running");
  const [subagentThinkingOpen, setSubagentThinkingOpen] = createSignal(false);
  const [subagentThinkingActive, setSubagentThinkingActive] = createSignal(false);
  const [subagentThinkingText, setSubagentThinkingText] = createSignal("");
  const [subagentToolIds, setSubagentToolIds] = createSignal<string[]>([]);
  const [subagentResultVisible, setSubagentResultVisible] = createSignal(false);

  const [assistantVisible, setAssistantVisible] = createSignal(false);
  const [assistantStreaming, setAssistantStreaming] = createSignal("");
  const [assistantDone, setAssistantDone] = createSignal(false);
  const [assistantPointsVisible, setAssistantPointsVisible] = createSignal(false);

  const [toolStatus, setToolStatus] = createSignal<Record<string, ToolStatus>>({});
  const [toolOutput, setToolOutput] = createSignal<Record<string, string[]>>({});
  const [toolOpen, setToolOpen] = createSignal<Record<string, boolean>>({});

  let cancelled = false;

  function sleep(ms: number) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }

  async function streamText(text: string, setter: (value: string) => void, duration: number) {
    setter("");
    const stepMs = Math.max(1, Math.floor(duration / Math.max(1, text.length)));
    let index = 0;
    while (index < text.length) {
      if (cancelled) return;
      index += 1;
      setter(text.slice(0, index));
      await sleep(stepMs);
    }
    setter(text);
  }

  function setToolStatusById(id: string, status: ToolStatus) {
    setToolStatus((map) => ({ ...map, [id]: status }));
  }

  function setToolOutputById(id: string, lines: string[]) {
    setToolOutput((map) => ({ ...map, [id]: lines }));
  }

  function setToolOpenById(id: string, open: boolean) {
    setToolOpen((map) => ({ ...map, [id]: open }));
  }

  function collapseOtherTools(exceptId: string) {
    setToolOpen((map) => {
      const next: Record<string, boolean> = {};
      for (const key of Object.keys(map)) {
        next[key] = key === exceptId;
      }
      return next;
    });
  }

  function collapseAllTools() {
    setToolOpen((map) => {
      const next: Record<string, boolean> = {};
      for (const key of Object.keys(map)) {
        next[key] = false;
      }
      return next;
    });
  }

  async function runTool(template: ToolTemplate) {
    collapseOtherTools(template.id);
    setToolOpenById(template.id, true);
    setToolStatusById(template.id, "running");
    setToolOutputById(template.id, []);
    await sleep(650);
    if (cancelled) return;

    setToolStatusById(template.id, "done");
    for (let index = 1; index <= template.output.length; index += 1) {
      if (cancelled) return;
      setToolOutputById(template.id, template.output.slice(0, index));
      await sleep(90);
    }
    await sleep(320);
  }

  async function playScenario() {
    await sleep(420);
    if (cancelled) return;
    setUserVisible(true);

    await sleep(720);
    if (cancelled) return;
    setMainThinkingOpen(true);
    setMainThinkingActive(true);
    await streamText(MAIN_THINKING, setMainThinkingText, 1500);
    if (cancelled) return;
    setMainThinkingActive(false);
    setMainThinkingOpen(false);

    setMainToolIds([TOOL_TEMPLATES[0].id]);
    await runTool(TOOL_TEMPLATES[0]);

    setMainToolIds((ids) => [...ids, TOOL_TEMPLATES[1].id]);
    await runTool(TOOL_TEMPLATES[1]);

    collapseAllTools();
    setSubagentVisible(true);
    setSubagentOpen(true);
    setSubagentStatus("running");
    setSubagentThinkingOpen(true);
    setSubagentThinkingActive(true);
    await streamText(SUBAGENT_THINKING, setSubagentThinkingText, 1000);
    if (cancelled) return;
    setSubagentThinkingActive(false);
    setSubagentThinkingOpen(false);

    setSubagentToolIds([SUBAGENT_TOOLS[0].id]);
    await runTool(SUBAGENT_TOOLS[0]);

    setSubagentToolIds((ids) => [...ids, SUBAGENT_TOOLS[1].id]);
    await runTool(SUBAGENT_TOOLS[1]);

    setSubagentResultVisible(true);
    setSubagentStatus("done");
    await sleep(480);
    if (cancelled) return;

    setMainToolIds((ids) => [...ids, TOOL_TEMPLATES[2].id]);
    await runTool(TOOL_TEMPLATES[2]);

    setMainToolIds((ids) => [...ids, TOOL_TEMPLATES[3].id]);
    await runTool(TOOL_TEMPLATES[3]);

    setAssistantVisible(true);
    await streamText(ASSISTANT_TEXT, setAssistantStreaming, 2200);
    if (cancelled) return;
    setAssistantDone(true);
    setAssistantPointsVisible(true);
  }

  onMount(() => {
    playScenario();
  });

  onCleanup(() => {
    cancelled = true;
  });

  function toggleTool(id: string) {
    setToolOpenById(id, !(toolOpen()[id] ?? false));
  }

  return (
    <div class="conversation">
      <div class="conversation-header">
        <div class="conversation-title-row">
          <span class="conversation-badge">OA</span>
          <div class="conversation-heading">
            <h2 class="conversation-title">Map authentication flow</h2>
            <span class="conversation-subtitle">
              OAgent · <span class="trace-mono">oagent/auth-audit</span> · #142
            </span>
          </div>
          <span class="conversation-active">
            <span class="conversation-active-dot" />
            Active
          </span>
        </div>
      </div>

      <div class="conversation-body">
        <div class="trace-list">
          <Show when={userVisible()}>
            <div class="trace-entry trace-entry-user">
              <div class="trace-msg">
                <div class="trace-msg-label">You</div>
                <p class="trace-msg-text">{USER_TEXT}</p>
              </div>
            </div>
          </Show>

          <Show when={mainThinkingText().length > 0}>
            <div class="trace-entry">
              <ThinkingBlock
                text={mainThinkingText()}
                open={mainThinkingOpen()}
                onToggle={() => setMainThinkingOpen(!mainThinkingOpen())}
                active={mainThinkingActive()}
              />
            </div>
          </Show>

          <For each={mainToolIds()}>
            {(id) => {
              const call = ALL_TOOLS.get(id)!;
              return (
                <div class="trace-entry">
                  <ToolCallCard
                    call={call}
                    open={toolOpen()[id] ?? false}
                    onToggle={() => toggleTool(id)}
                    status={toolStatus()[id] ?? "running"}
                    output={toolOutput()[id] ?? []}
                  />
                </div>
              );
            }}
          </For>

          <Show when={subagentVisible()}>
            <div class="trace-entry trace-entry-subagent">
              <SubagentPanel
                open={subagentOpen()}
                onToggle={() => setSubagentOpen(!subagentOpen())}
                status={subagentStatus()}
                thinkingOpen={subagentThinkingOpen()}
                onThinkingToggle={() => setSubagentThinkingOpen(!subagentThinkingOpen())}
                thinkingActive={subagentThinkingActive()}
                thinkingText={subagentThinkingText()}
                toolIds={subagentToolIds()}
                resultVisible={subagentResultVisible()}
                toolStatus={(toolId) => toolStatus()[toolId] ?? "done"}
                toolOutput={(toolId) => toolOutput()[toolId] ?? []}
                toolOpen={(toolId) => toolOpen()[toolId] ?? false}
                onToolToggle={toggleTool}
              />
            </div>
          </Show>

          <Show when={assistantVisible()}>
            <div class="trace-entry trace-entry-assistant">
              <div class="trace-msg">
                <div class="trace-msg-label">OAgent</div>
                <Show
                  when={assistantDone()}
                  fallback={<p class="trace-msg-text">{assistantStreaming()}</p>}
                >
                  <p class="trace-msg-text">
                    Authentication is client-side and local: the login form stores a session
                    token in <span class="trace-mono">localStorage</span>, the app reads it on
                    startup, and the backend <span class="trace-mono">OpenCodeService</span>{" "}
                    verifies the local <span class="trace-mono">opencode2</span> service with{" "}
                    <span class="trace-mono">IsReady()</span> /{" "}
                    <span class="trace-mono">Setup()</span>.
                  </p>
                </Show>
                <Show when={assistantPointsVisible()}>
                  <ul class="trace-msg-list">
                    <For each={ASSISTANT_POINTS}>
                      {(point) => (
                        <li>
                          {point.before}
                          <Show when={point.code}>
                            <span class="trace-mono">{point.code}</span>
                          </Show>
                          {point.after}
                        </li>
                      )}
                    </For>
                  </ul>
                </Show>
              </div>
            </div>
          </Show>
        </div>
      </div>
    </div>
  );
}
