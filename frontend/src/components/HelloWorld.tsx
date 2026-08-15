import { createSignal, onMount, Show } from "solid-js";
import { OpenCodeService, OpenCodeStatus } from "../../bindings/changeme";

export default function HelloWorld() {
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
        setSetupMessage(`Conectat la ${result.url}`);
      } else {
        setSetupMessage(result.error || "Setup eșuat.");
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
  });

  return (
    <>
      <main class="container">
        <header class="brand">
          <span class="brand-mark" aria-label="OAgent">
            <img src="/wails.png" class="brand-logo" alt="OAgent logo" />
          </span>
        </header>

        <h1 class="title">
          <span class="title-accent">OAgent</span>
        </h1>
        <p class="subtitle">Agent Orchestrator — OpenCode V2 Harness prin API.</p>

        <button
          class="harness-btn"
          classList={{ "is-ready": isReady() }}
          onClick={() => setPopupOpen(!popupOpen())}
          aria-label="Harnesses"
        >
          <span
            class="harness-dot"
            classList={{
              "dot-green": isReady(),
              "dot-red": !isReady() && !checking(),
            }}
          />
          <span class="harness-label">
            {checking() ? "Verific…" : isReady() ? "OpenCode" : "Harnesses"}
          </span>
        </button>

        <Show when={popupOpen()}>
          <div class="harness-popup" role="dialog" aria-label="Harnesses">
            <div class="harness-popup-header">
              <span>Harnesses</span>
              <button
                class="harness-close"
                onClick={() => setPopupOpen(false)}
                aria-label="Close"
              >
                ×
              </button>
            </div>

            <div class="harness-item">
              <span class="harness-item-name">OpenCode2</span>
              <Show
                when={isReady()}
                fallback={
                  <button
                    class="btn setup-btn"
                    onClick={doSetup}
                    disabled={settingUp()}
                  >
                    {settingUp() ? "Instalez…" : "Setup"}
                  </button>
                }
              >
                <span class="harness-ready-badge">
                  <span class="harness-dot dot-green" />
                  Online
                </span>
              </Show>
            </div>

            <Show when={setupMessage()}>
              <p class="setup-message">{setupMessage()}</p>
            </Show>
          </div>
        </Show>
      </main>
    </>
  );
}
