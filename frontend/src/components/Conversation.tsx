import { For, Show, createMemo, createSignal } from "solid-js";
import {
  ConversationMessage,
  SubagentInfo,
  ToolCall,
} from "../../bindings/changeme";

type SessionContext = {
  id: string;
  title: string;
  project: string;
  agent?: string;
  directory?: string;
  active: boolean;
};

function ToolGlyph(props: { name: string }) {
  const name = props.name.toLowerCase();
  let path = "M4 2h5l3 3v9H4V2zm5 1.5V5h1.5L9 3.5zM5 3v10h6V6H9V3H5z";
  if (name.includes("bash") || name.includes("shell") || name.includes("exec")) {
    path = "M2 4h12v8H2V4zm1 1v6h10V5H3zm1 1h2v1H4V6zm4 1l2.5 1.5L8 10V7z";
  } else if (name.includes("grep") || name.includes("search") || name.includes("find")) {
    path = "M7 1.5a5.5 5.5 0 1 0 0 11c1.32 0 2.54-.46 3.48-1.24l3.13 3.13 1.28-1.28-3.13-3.13A5.5 5.5 0 0 0 7 1.5zm0 1.5a4 4 0 1 1 0 8 4 4 0 0 1 0-8z";
  } else if (name.includes("list") || name.includes("directory") || name.includes("glob")) {
    path = "M1.5 3.5A1.5 1.5 0 0 1 3 2h3l1.5 2H13a1.5 1.5 0 0 1 1.5 1.5v6A1.5 1.5 0 0 1 13 13H3a1.5 1.5 0 0 1-1.5-1.5v-8z";
  } else if (name.includes("patch") || name.includes("edit") || name.includes("write")) {
    path = "M3 2h7l3 3v9H3V2zm7 1.5V5h1.5L10 3.5zM4 3v10h8V6h-3V3H4zm1 2h5v1H5V5zm0 2h5v1H5V7zm0 2h3v1H5V9z";
  }

  return (
    <svg width="14" height="14" viewBox="0 0 16 16" fill="none" aria-hidden="true">
      <path d={path} fill="currentColor" fill-rule="evenodd" clip-rule="evenodd" />
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

function ThinkingBlock(props: { text: string; agent?: string }) {
  const [open, setOpen] = createSignal(false);
  return (
    <div class="trace-thinking">
      <button class="trace-thinking-toggle" onClick={() => setOpen(!open())} aria-expanded={open()}>
        <span class="trace-thinking-dot" />
        <span class="trace-thinking-label">Thinking</span>
        <Show when={props.agent}>
          <span class="trace-thinking-agent">{props.agent}</span>
        </Show>
        <span class="trace-thinking-spacer" />
        <Chevron open={open()} />
      </button>
      <Show when={open()}>
        <p class="trace-thinking-text">{props.text}</p>
      </Show>
    </div>
  );
}

function lineMeta(tool: ToolCall): { kind: "bash" | "diff" | "plain"; lines: string[] } {
  const raw = tool.output || "";
  if (!raw) return { kind: "plain", lines: [] };

  if (tool.diff) {
    return { kind: "diff", lines: raw.split(/\r?\n/) };
  }

  const isBash = /bash|shell|exec|command/i.test(tool.name);
  return { kind: isBash ? "bash" : "plain", lines: raw.split(/\r?\n/) };
}

function toolDuration(tool: ToolCall): string {
  if (tool.completedAt && tool.createdAt) {
    const ms = tool.completedAt - tool.createdAt;
    if (ms > 0) {
      if (ms < 1000) return `${ms}ms`;
      return `${(ms / 1000).toFixed(1)}s`;
    }
  }
  return "";
}

function toolStatusLabel(tool: ToolCall): "running" | "done" | "error" {
  const status = (tool.status || "").toLowerCase();
  if (status.includes("error") || status.includes("fail")) return "error";
  if (status.includes("run") || status.includes("stream")) return "running";
  return "done";
}

function ToolCallCard(props: { tool: ToolCall; nested?: boolean }) {
  const [open, setOpen] = createSignal(false);
  const status = toolStatusLabel(props.tool);
  const output = lineMeta(props.tool);
  const inputPreview = (toolInputPreview(props.tool) || "").slice(0, 120);

  return (
    <div class="trace-tool" classList={{ "is-nested": props.nested, "is-error": status === "error" }}>
      <button class="trace-tool-toggle" onClick={() => setOpen(!open())} aria-expanded={open()}>
        <span
          class="trace-tool-icon"
          classList={{
            "is-running": status === "running",
            "is-done": status === "done",
            "is-error": status === "error",
          }}
        >
          <ToolGlyph name={props.tool.name} />
        </span>
        <span class="trace-tool-name">{props.tool.name}</span>
        <span class="trace-tool-args">{inputPreview}</span>
        <span class="trace-tool-spacer" />
        <span class="trace-tool-duration">{toolDuration(props.tool)}</span>
        <span class="trace-tool-status">
          {status === "running" ? "Running" : status === "done" ? "Done" : "Error"}
        </span>
        <Chevron open={open()} />
      </button>
      <Show when={open()}>
        <div class="trace-tool-output">
          <For each={output.lines}>
            {(line) => (
              <span
                class="trace-tool-line"
                classList={{
                  "is-add": output.kind === "diff" && line.startsWith("+"),
                  "is-remove": output.kind === "diff" && line.startsWith("-"),
                  "is-command": output.kind === "bash" && line.startsWith("$"),
                }}
              >
                {line}
              </span>
            )}
          </For>
          <Show when={output.lines.length === 0 && props.tool.input}>
            <span class="trace-tool-line">{props.tool.input}</span>
          </Show>
        </div>
      </Show>
    </div>
  );
}

function toolInputPreview(tool: ToolCall): string {
  if (!tool.input) return "";
  try {
    const parsed = JSON.parse(tool.input);
    if (parsed && typeof parsed === "object") {
      if (typeof parsed.command === "string") return parsed.command;
      if (typeof parsed.filePath === "string") return parsed.filePath;
      if (typeof parsed.path === "string") return parsed.path;
      if (typeof parsed.pattern === "string") return parsed.pattern;
      return Object.keys(parsed).slice(0, 3).join(", ");
    }
    return tool.input;
  } catch {
    return tool.input;
  }
}

function SubagentPanel(props: { subagent: SubagentInfo }) {
  const [open, setOpen] = createSignal(false);
  const messages = () => props.subagent.messages || [];
  const purpose = createMemo(() => {
    const first = messages().find((m) => m.role === "user");
    return (first?.text || props.subagent.title || "").slice(0, 160);
  });
  const result = createMemo(() => {
    const last = messages().filter((m) => m.role === "assistant" && m.text).pop();
    return last?.text || "";
  });
  const toolCount = createMemo(() =>
    messages().reduce((sum, m) => sum + (m.tools?.length || 0), 0),
  );

  return (
    <div class="trace-subagent">
      <button
        class="trace-subagent-header"
        type="button"
        onClick={() => setOpen(!open())}
        aria-expanded={open()}
        aria-label={open() ? "Collapse subagent" : "Expand subagent"}
      >
        <span class="trace-subagent-glyph" aria-hidden="true">
          <svg width="14" height="14" viewBox="0 0 16 16" fill="none">
            <rect x="2" y="2" width="5" height="5" rx="1" stroke="currentColor" stroke-width="1.5" />
            <rect x="9" y="2" width="5" height="5" rx="1" stroke="currentColor" stroke-width="1.5" />
            <rect x="2" y="9" width="5" height="5" rx="1" stroke="currentColor" stroke-width="1.5" />
            <rect x="9" y="9" width="5" height="5" rx="1" stroke="currentColor" stroke-width="1.5" />
          </svg>
        </span>
        <span class="trace-subagent-name">{props.subagent.agent || "subagent"}</span>
        <span class="trace-subagent-kind">subagent</span>
        <span class="trace-subagent-spacer" />
        <span class="trace-subagent-meta">{toolCount()} tools</span>
        <span
          class="trace-subagent-status"
          classList={{ "is-running": props.subagent.status === "active" }}
        >
          <span class="trace-subagent-status-dot" />
          {props.subagent.status === "active" ? "Working" : "Done"}
        </span>
        <span class="trace-subagent-chevron">
          <Chevron open={open()} />
        </span>
      </button>

      <Show when={open()}>
        <Show when={purpose()}>
          <div class="trace-subagent-purpose">{purpose()}</div>
        </Show>
        <div class="trace-subagent-steps">
          <For each={messages()}>
            {(message) => (
              <>
                <Show when={message.reasoning}>
                  <ThinkingBlock text={message.reasoning!} agent={message.agent} />
                </Show>
                <For each={message.tools || []}>
                  {(tool) => <ToolCallCard tool={tool} nested />}
                </For>
              </>
            )}
          </For>
        </div>
        <Show when={result()}>
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
            <p class="trace-subagent-result-text">{result()}</p>
          </div>
        </Show>
      </Show>
    </div>
  );
}

function UserMessage(props: { message: ConversationMessage }) {
  return (
    <div class="trace-entry trace-entry-user">
      <div class="trace-msg">
        <div class="trace-msg-label">You</div>
        <p class="trace-msg-text">{props.message.text}</p>
      </div>
    </div>
  );
}

function AssistantMessage(props: { message: ConversationMessage }) {
  const label = () => props.message.agent || "OAgent";
  return (
    <div class="trace-entry trace-entry-assistant">
      <div class="trace-msg">
        <div class="trace-msg-label">{label()}</div>
        <Show when={props.message.reasoning}>
          <ThinkingBlock text={props.message.reasoning!} agent={props.message.agent} />
        </Show>
        <For each={props.message.tools || []}>
          {(tool) => <ToolCallCard tool={tool} />}
        </For>
        <Show when={props.message.text}>
          <p class="trace-msg-text">{props.message.text}</p>
        </Show>
      </div>
    </div>
  );
}

function SystemMessage(props: { message: ConversationMessage }) {
  const label = props.message.role === "synthetic" ? "Tool" : props.message.role;
  return (
    <div class="trace-entry trace-entry-system">
      <div class="trace-msg">
        <div class="trace-msg-label">{label}</div>
        <p class="trace-msg-text">{props.message.text}</p>
      </div>
    </div>
  );
}

export default function Conversation(props: {
  session: SessionContext | null;
  messages: ConversationMessage[];
  subagents: SubagentInfo[];
  loading: boolean;
  error: string;
}) {
  return (
    <div class="conversation">
      <div class="conversation-header">
        <div class="conversation-title-row">
          <span class="conversation-badge">OA</span>
          <div class="conversation-heading">
            <h2 class="conversation-title">
              {props.session?.title || "No conversation selected"}
            </h2>
            <span class="conversation-subtitle">
              {props.session?.project || "OAgent"} ·{" "}
              <span class="trace-mono">{props.session?.directory || ""}</span>
            </span>
          </div>
          <Show when={props.session?.active}>
            <span class="conversation-active">
              <span class="conversation-active-dot" />
              Active
            </span>
          </Show>
        </div>
      </div>

      <div class="conversation-body">
        <Show when={props.loading}>
          <div class="conversation-empty">Loading conversation…</div>
        </Show>

        <Show when={!props.loading && props.error}>
          <div class="conversation-empty conversation-error">{props.error}</div>
        </Show>

        <Show when={!props.loading && !props.error}>
          <div class="trace-list">
            <For each={props.messages}>
              {(message) => {
                if (message.role === "user") return <UserMessage message={message} />;
                if (message.role === "assistant") return <AssistantMessage message={message} />;
                return <SystemMessage message={message} />;
              }}
            </For>

            <For each={props.subagents}>
              {(subagent) => (
                <div class="trace-entry trace-entry-subagent">
                  <SubagentPanel subagent={subagent} />
                </div>
              )}
            </For>

            <Show when={props.messages.length === 0 && props.subagents.length === 0}>
              <div class="conversation-empty">No messages in this session yet.</div>
            </Show>
          </div>
        </Show>
      </div>
    </div>
  );
}
