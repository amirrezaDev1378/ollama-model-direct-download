<script lang="ts">
    import {GetModelDownloadLinks} from "../../../wailsjs/go/main/App.js";
    import type {main} from "../../../wailsjs/go/models.js";
    import {ClipboardSetText} from "../../../wailsjs/runtime/runtime";

    let modelName: string = "";
    let loading: boolean = false;
    let error: string | null = null;
    let response: main.GetLinkResponse | null = null;
    let copyButtonStatus: "error" | "success" | "pending" | "default" = "default";

    async function handleGetLinks() {
        if (!modelName.trim()) {
            error = "Please enter a model name.";
            return;
        }
        loading = true;
        error = null;
        response = null;

        try {
            response = await GetModelDownloadLinks(modelName);
        } catch (e: any) {
            error = e.message || String(e);
        } finally {
            loading = false;
        }
    }

    function copyAllToClipboard() {
        if (response) {
            copyButtonStatus = "pending";
            ClipboardSetText([response.manifestLink, ...response.downloadLinks].join("\n")).then(r => {
                setTimeout(() => {
                    copyButtonStatus = "success";
                }, 800);
            }).catch(err => {
                console.error(err.message);
                copyButtonStatus = "error";

            }).finally(() => {
                setTimeout(() => {
                    copyButtonStatus = "default";
                }, 4000);
            });
        }
    }
</script>

<div class="card">
    <h2>Get Model Download Links</h2>
    <div class="input-group">
        <label for="modelName">Model Name:</label>
        <input
                type="text"
                id="modelName"
                bind:value={modelName}
                placeholder="e.g. deepseek-coder-v2:latest"
                on:keydown={(e) => e.key === 'Enter' && handleGetLinks()}
        />
    </div>

    <button class="primary-btn" on:click={handleGetLinks} disabled={loading}>
        {loading ? 'Fetching...' : 'Get Links'}
    </button>

    {#if error}
        <div class="error-box">
            {error}
        </div>
    {/if}

    {#if response}
        <div class="results">
            <h3>Manifest Link:</h3>
            <div class="link-box">
                <a href={response.manifestLink} target="_blank" rel="noopener noreferrer">{response.manifestLink}</a>
            </div>

            <h3>Layer Download Links:</h3>
            <ul class="link-list">
                {#each response.downloadLinks as link}
                    <li>
                        <a href={link} target="_blank" rel="noopener noreferrer">{link}</a>
                    </li>
                {/each}
            </ul>
            <button class="btn-copy" on:click={copyAllToClipboard}>

                {copyButtonStatus === "default" ? "Copy All Links" : ""}
                {copyButtonStatus === "error" ? "Failed to copy links!" : ""}
                {copyButtonStatus === "success" ? "Copied ..." : ""}
                {copyButtonStatus === "pending" ? " ... " : ""}
            </button>
        </div>
    {/if}
</div>

<style>
    .card {
        background-color: var(--card-bg, #2a2a35);
        padding: 2rem;
        border-radius: 8px;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        text-align: left;
    }

    h2 {
        margin-top: 0;
        margin-bottom: 1.5rem;
        color: var(--text-color, #ffffff);
    }

    h3 {
        margin-top: 1.5rem;
        margin-bottom: 0.5rem;
        font-size: 1.1rem;
        color: var(--text-color, #ffffff);
    }

    .input-group {
        margin-bottom: 1.5rem;
        display: flex;
        flex-direction: column;
    }

    label {
        margin-bottom: 0.5rem;
        font-weight: 500;
        color: var(--text-muted, #a0a0b0);
    }

    input {
        padding: 0.75rem;
        border-radius: 4px;
        border: 1px solid var(--border-color, #4a4a5a);
        background-color: var(--input-bg, #1a1a24);
        color: var(--text-color, #ffffff);
        font-size: 1rem;
    }

    input:focus {
        outline: none;
        border-color: var(--primary-color, #007bff);
        box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
    }

    .primary-btn {
        background-color: var(--primary-color, #007bff);
        color: white;
        border: none;
        padding: 0.75rem 1.5rem;
        border-radius: 4px;
        font-size: 1rem;
        cursor: pointer;
        transition: background-color 0.2s;
        width: 100%;
    }

    .primary-btn:hover:not(:disabled) {
        background-color: var(--primary-hover, #0056b3);
    }

    .primary-btn:disabled {
        opacity: 0.6;
        cursor: not-allowed;
    }

    .btn-copy {
        padding: 16px;
        border-radius: 12px;
        background: #6aefb0;
        width: 300px;
        margin-top: 24px;
        margin-inline: auto;
        border: none;
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
    }

    .btn-copy:hover {
        background: #5bd79c;
    }
    .error-box {
        margin-top: 1.5rem;
        padding: 1rem;
        background-color: rgba(220, 53, 69, 0.1);
        border-left: 4px solid #dc3545;
        color: #ff6b6b;
        border-radius: 4px;
    }

    .results {
        margin-top: 2rem;
        border-top: 1px solid var(--border-color, #4a4a5a);
        padding-top: 1rem;
        display: flex;
        flex-direction: column;
    }

    .link-box {
        background-color: var(--input-bg, #1a1a24);
        padding: 0.75rem;
        border-radius: 4px;
        word-break: break-all;
        border: 1px solid var(--border-color, #4a4a5a);
    }

    .link-list {
        list-style-type: none;
        padding: 0;
        margin: 0;
    }

    .link-list li {
        background-color: var(--input-bg, #1a1a24);
        padding: 0.75rem;
        border-radius: 4px;
        margin-bottom: 0.5rem;
        word-break: break-all;
        border: 1px solid var(--border-color, #4a4a5a);
    }

    a {
        color: var(--link-color, #4dabf7);
        text-decoration: none;
    }

    a:hover {
        text-decoration: underline;
    }
</style>
