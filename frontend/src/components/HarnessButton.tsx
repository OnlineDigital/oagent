import { createMemo, createSignal, For, onCleanup, onMount, Show } from "solid-js";
import {
  HarnessInfo,
  McpInfo,
  OpenCodeService,
  OpenCodeStatus,
  PluginInfo,
  SkillInfo,
} from "../../bindings/changeme";

function ToggleSwitch(props: {
  active: boolean;
  disabled?: boolean;
  onToggle: () => void;
}) {
  return (
    <button
      class="harness-toggle"
      classList={{ "is-active": props.active }}
      disabled={props.disabled}
      role="switch"
      aria-checked={props.active}
      onClick={props.onToggle}
    >
      <span class="harness-toggle-thumb" />
    </button>
  );
}

export default function HarnessButton() {
  let containerRef!: HTMLDivElement;
  let popupRef!: HTMLDivElement;
  let skillPopoverRef!: HTMLDivElement;

  const [status, setStatus] = createSignal<OpenCodeStatus | null>(null);
  const [popupOpen, setPopupOpen] = createSignal(false);
  const [checking, setChecking] = createSignal(true);
  const [settingUp, setSettingUp] = createSignal(false);
  const [setupMessage, setSetupMessage] = createSignal("");

  const [harnesses, setHarnesses] = createSignal<HarnessInfo[]>([]);
  const [selectedHarnessId, setSelectedHarnessId] = createSignal<string | null>(null);
  const [mcps, setMcps] = createSignal<McpInfo[]>([]);
  const [plugins, setPlugins] = createSignal<PluginInfo[]>([]);
  const [skills, setSkills] = createSignal<SkillInfo[]>([]);
  const [searchQuery, setSearchQuery] = createSignal("");
  const [controlsLoading, setControlsLoading] = createSignal(false);
  const [controlsError, setControlsError] = createSignal("");
  const [toggling, setToggling] = createSignal<string | null>(null);
  const [skillPopover, setSkillPopover] = createSignal<{
    skill: SkillInfo;
    x: number;
    y: number;
  } | null>(null);

  const isReady = () => status()?.ready ?? false;
  const selectedHarness = () =>
    harnesses().find((item) => item.id === selectedHarnessId()) ?? harnesses()[0] ?? null;

  const filteredMcps = createMemo(() => {
    const query = searchQuery().toLowerCase();
    if (!query) return mcps();
    return mcps().filter((item) => item.name.toLowerCase().includes(query));
  });

  const filteredPlugins = createMemo(() => {
    const query = searchQuery().toLowerCase();
    if (!query) return plugins();
    return plugins().filter((item) => item.id.toLowerCase().includes(query));
  });

  const filteredSkills = createMemo(() => {
    const query = searchQuery().toLowerCase();
    if (!query) return skills();
    return skills().filter(
      (item) =>
        item.name.toLowerCase().includes(query) ||
        item.id.toLowerCase().includes(query) ||
        (item.description ?? "").toLowerCase().includes(query),
    );
  });

  async function checkStatus() {
    setChecking(true);
    try {
      const result = await OpenCodeService.IsReady();
      setStatus(result);
    } catch (err) {
      console.error(err);
      setStatus({ ready: false, url: "", error: String(err) });
    } finally {
      setChecking(false);
    }
  }

  async function loadHarnesses() {
    try {
      const list = (await OpenCodeService.Harnesses()) ?? [];
      setHarnesses(list);
      if (!selectedHarnessId() && list.length) {
        setSelectedHarnessId(list[0].id);
      }
    } catch (err) {
      console.error(err);
      setHarnesses([]);
    }
  }

  async function loadControls() {
    setControlsLoading(true);
    setControlsError("");
    try {
      const [mcpList, pluginList, skillList] = await Promise.all([
        OpenCodeService.McpServers(),
        OpenCodeService.Plugins(),
        OpenCodeService.Skills(),
      ]);
      setMcps(mcpList ?? []);
      setPlugins(pluginList ?? []);
      setSkills(skillList ?? []);
    } catch (err) {
      console.error(err);
      setControlsError(String(err));
    } finally {
      setControlsLoading(false);
    }
  }

  async function doSetup() {
    setSettingUp(true);
    setSetupMessage("");
    try {
      const result = await OpenCodeService.Setup();
      setStatus(result);
      if (result.ready) {
        setSetupMessage(`Connected to ${result.url}`);
      } else {
        setSetupMessage(result.error || "Setup failed.");
      }
    } catch (err) {
      console.error(err);
      setSetupMessage(String(err));
    } finally {
      setSettingUp(false);
    }
  }

  async function toggleMcp(server: McpInfo) {
    const key = `mcp:${server.name}`;
    if (toggling()) return;
    setToggling(key);
    try {
      await OpenCodeService.ToggleMcp(server.name, !server.active);
      await loadControls();
    } catch (err) {
      console.error(err);
      setControlsError(String(err));
    } finally {
      setToggling(null);
    }
  }

  async function togglePlugin(plugin: PluginInfo) {
    const key = `plugin:${plugin.id}`;
    if (toggling()) return;
    setToggling(key);
    try {
      await OpenCodeService.TogglePlugin(plugin.id, !plugin.active);
      await loadControls();
    } catch (err) {
      console.error(err);
      setControlsError(String(err));
    } finally {
      setToggling(null);
    }
  }

  async function toggleSkill(skill: SkillInfo) {
    const key = `skill:${skill.id}`;
    if (toggling()) return;
    setToggling(key);
    try {
      await OpenCodeService.ToggleSkill(skill.id, !skill.active);
      await loadControls();
    } catch (err) {
      console.error(err);
      setControlsError(String(err));
    } finally {
      setToggling(null);
    }
  }

  function showSkillPopover(skill: SkillInfo, e: MouseEvent) {
    const popupRect = popupRef?.getBoundingClientRect();
    if (!popupRect) return;
    setSkillPopover({
      skill,
      x: e.clientX - popupRect.left,
      y: e.clientY - popupRect.top,
    });
  }

  function hideSkillPopover() {
    setSkillPopover(null);
  }

  function moveSkillPopover(e: MouseEvent) {
    const popover = skillPopover();
    if (!popover || !popupRef || !skillPopoverRef) return;
    const popupRect = popupRef.getBoundingClientRect();
    const popoverRect = skillPopoverRef.getBoundingClientRect();
    if (!popoverRect.width) return;

    let x = e.clientX - popupRect.left + 14;
    let y = e.clientY - popupRect.top + 14;
    const maxX = popupRect.width - popoverRect.width - 8;
    const maxY = popupRect.height - popoverRect.height - 8;
    x = Math.max(8, Math.min(x, maxX));
    y = Math.max(8, Math.min(y, maxY));
    setSkillPopover({ ...popover, x, y });
  }

  function openPopup() {
    const next = !popupOpen();
    setPopupOpen(next);
    if (next) {
      loadHarnesses();
      loadControls();
    }
  }

  function selectHarness(id: string) {
    setSelectedHarnessId(id);
    loadControls();
  }

  onMount(() => {
    checkStatus();
    loadHarnesses();
    document.addEventListener("keydown", handleEscape);
    document.addEventListener("pointerdown", handleOutsidePointer);
  });

  onCleanup(() => {
    document.removeEventListener("keydown", handleEscape);
    document.removeEventListener("pointerdown", handleOutsidePointer);
  });

  function handleEscape(e: KeyboardEvent) {
    if (e.key === "Escape" && popupOpen()) {
      setPopupOpen(false);
    }
  }

  function handleOutsidePointer(e: PointerEvent) {
    if (!popupOpen()) return;
    const target = e.target as Node;
    if (containerRef && !containerRef.contains(target)) {
      setPopupOpen(false);
    }
  }

  return (
    <div ref={containerRef}>
      <button
        class="harness-btn"
        classList={{ "is-ready": isReady() }}
        onClick={openPopup}
        aria-label="Harnesses"
        aria-expanded={popupOpen()}
      >
        <span
          class="harness-dot"
          classList={{
            "dot-online": isReady(),
            "dot-offline": !isReady() && !checking(),
            "dot-checking": checking(),
          }}
        />
        <span class="harness-label">
          {checking() ? "Checking…" : isReady() ? "OpenCode2" : "Harnesses"}
        </span>
        <svg
          width="16"
          height="16"
          viewBox="0 0 16 16"
          fill="none"
          class="harness-chevron"
          classList={{ "is-open": popupOpen() }}
        >
          <path
            d="M4 6l4 4 4-4"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </button>

      <Show when={popupOpen()}>
        <div
          class="harness-popup"
          role="dialog"
          aria-label="Harnesses"
          ref={popupRef}
        >
          <div class="harness-popup-header">
            <span class="harness-popup-title">Harnesses</span>
            <button
              class="harness-close"
              onClick={() => setPopupOpen(false)}
              aria-label="Close"
            >
              <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
                <path
                  d="M4 4l8 8M12 4l-8 8"
                  stroke="currentColor"
                  stroke-width="1.5"
                  stroke-linecap="round"
                />
              </svg>
            </button>
          </div>

          <div class="harness-popup-body">
            <div class="harness-column">
              <div class="harness-column-label">Devices</div>
              <For each={harnesses()}>
                {(harness) => (
                  <button
                    class="harness-device"
                    classList={{ "is-active": harness.id === selectedHarnessId() }}
                    onClick={() => selectHarness(harness.id)}
                  >
                    <span
                      class="harness-dot"
                      classList={{
                        "dot-online": harness.online,
                        "dot-offline": !harness.online,
                      }}
                    />
                    <span class="harness-device-info">
                      <span class="harness-device-name">{harness.name}</span>
                      <Show when={harness.description}>
                        <span class="harness-device-desc">{harness.description}</span>
                      </Show>
                    </span>
                    <Show when={harness.local}>
                      <span class="harness-local-badge">Local</span>
                    </Show>
                  </button>
                )}
              </For>

              <Show when={!harnesses().length}>
                <div class="harness-empty">No harnesses available</div>
              </Show>
            </div>

            <div class="harness-column harness-column-controls">
              <div class="harness-column-label">
                {selectedHarness()?.name ?? "Controls"}
              </div>

              <div class="harness-search">
                <svg
                  width="14"
                  height="14"
                  viewBox="0 0 16 16"
                  fill="none"
                  aria-hidden="true"
                >
                  <circle cx="7" cy="7" r="4.5" stroke="currentColor" stroke-width="1.5" />
                  <path d="M11 11l3 3" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
                </svg>
                <input
                  type="search"
                  class="harness-search-input"
                  placeholder="Search skills, MCP, plugins…"
                  value={searchQuery()}
                  onInput={(e) => setSearchQuery(e.currentTarget.value)}
                  aria-label="Search skills, MCP, and plugins"
                />
              </div>

              <Show when={controlsLoading()}>
                <div class="harness-empty">Loading…</div>
              </Show>

              <Show when={!controlsLoading() && controlsError()}>
                <div class="harness-error">{controlsError()}</div>
              </Show>

              <Show when={!controlsLoading()}>
                <div class="harness-control-section">
                  <div class="harness-control-section-title">Skills</div>
                  <For each={filteredSkills()}>
                    {(skill) => (
                      <div
                        class="harness-control-row harness-skill-row"
                        onMouseEnter={(e) => showSkillPopover(skill, e)}
                        onMouseLeave={hideSkillPopover}
                        onMouseMove={moveSkillPopover}
                      >
                        <div class="harness-skill-line">
                          <span class="harness-control-name">{skill.name}</span>
                          <ToggleSwitch
                            active={skill.active}
                            disabled={toggling() === `skill:${skill.id}`}
                            onToggle={() => toggleSkill(skill)}
                          />
                        </div>
                      </div>
                    )}
                  </For>
                  <Show when={!filteredSkills().length}>
                    <div class="harness-empty">No skills found</div>
                  </Show>
                </div>

                <div class="harness-control-section">
                  <div class="harness-control-section-title">MCP Servers</div>
                  <For each={filteredMcps()}>
                    {(mcp) => (
                      <div class="harness-control-row">
                        <span class="harness-control-name">{mcp.name}</span>
                        <ToggleSwitch
                          active={mcp.active}
                          disabled={toggling() === `mcp:${mcp.name}`}
                          onToggle={() => toggleMcp(mcp)}
                        />
                      </div>
                    )}
                  </For>
                  <Show when={!filteredMcps().length}>
                    <div class="harness-empty">No MCP servers found</div>
                  </Show>
                </div>

                <div class="harness-control-section">
                  <div class="harness-control-section-title">Plugins</div>
                  <For each={filteredPlugins()}>
                    {(plugin) => (
                      <div class="harness-control-row">
                        <span class="harness-control-name">{plugin.id}</span>
                        <ToggleSwitch
                          active={plugin.active}
                          disabled={toggling() === `plugin:${plugin.id}`}
                          onToggle={() => togglePlugin(plugin)}
                        />
                      </div>
                    )}
                  </For>
                  <Show when={!filteredPlugins().length}>
                    <div class="harness-empty">No plugins found</div>
                  </Show>
                </div>
              </Show>
            </div>
          </div>

          <Show when={skillPopover()}>
            <div
              class="harness-skill-popover"
              ref={skillPopoverRef}
              style={{
                left: `${skillPopover()!.x}px`,
                top: `${skillPopover()!.y}px`,
              }}
            >
              <div class="harness-skill-popover-title">
                {skillPopover()!.skill.name}
              </div>
              <div class="harness-skill-popover-desc">
                {skillPopover()!.skill.description || "No description"}
              </div>
            </div>
          </Show>

          <Show when={setupMessage()}>
            <p class="setup-message">{setupMessage()}</p>
          </Show>
        </div>
      </Show>
    </div>
  );
}
