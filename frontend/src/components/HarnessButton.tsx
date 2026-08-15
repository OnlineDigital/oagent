import { createSignal, onCleanup, onMount, Show } from "solid-js";
import { OpenCodeService, OpenCodeStatus } from "../../bindings/changeme";

export default function HarnessButton() {
  const [status, setStatus] = createSignal<OpenCodeStatus | null>(null);
  const [popupOpen, setPopupOpen] = createSignal(false);
  const [checking, setChecking] = createSignal(true);
  const [settingUp, setSettingUp] = createSignal(false);
  const [setupMessage, setSetupMessage] = createSignal("");

  const isReady = () => status()?.ready ?? false;

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

  onMount(() => {
    checkStatus();
    document.addEventListener("keydown", handleEscape);
  });

  onCleanup(() => {
    document.removeEventListener("keydown", handleEscape);
  });

  function handleEscape(e: KeyboardEvent) {
    if (e.key === "Escape" && popupOpen()) {
      setPopupOpen(false);
    }
  }

  return (
    <>
      <button
        class="harness-btn"
        classList={{ "is-ready": isReady() }}
        onClick={() => setPopupOpen(!popupOpen())}
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
        <div class="harness-popup" role="dialog" aria-label="Harnesses">
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

          <div class="harness-item">
            <div class="harness-item-info">
              <span class="harness-item-name">OpenCode2</span>
              <span class="harness-item-desc">
                {isReady() ? "Service active" : "Service stopped"}
              </span>
            </div>
            <Show
              when={isReady()}
              fallback={
                <button
                  class="btn-primary"
                  onClick={doSetup}
                  disabled={settingUp()}
                >
                  {settingUp() ? "Installing…" : "Setup"}
                </button>
              }
            >
              <span class="harness-status-badge">
                <span class="harness-dot dot-online" />
                Online
              </span>
            </Show>
          </div>

          <Show when={setupMessage()}>
            <p class="setup-message">{setupMessage()}</p>
          </Show>
        </div>
      </Show>
    </>
  );
}
