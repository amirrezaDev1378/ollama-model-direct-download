<script lang="ts">
    import { InstallModel } from '../../../wailsjs/go/main/App.js';
    import { EventsOn, EventsOff } from '../../../wailsjs/runtime/runtime.js';
    import { onMount, onDestroy } from 'svelte';

    let modelName: string = "";
    let blobsPath: string = "";
    let loading: boolean = false;
    let error: string | null = null;
    let successMessage: string | null = null;
    let progress: number = 0;
    let progressMessage: string = "";

    onMount(() => {
        EventsOn('install_progress', (data: any) => {
            if (data) {
                progress = data.progress;
                progressMessage = data.message;
            }
        });
    });

    onDestroy(() => {
        EventsOff('install_progress');
    });

    async function handleInstall() {
        if (!modelName.trim() || !blobsPath.trim()) {
            error = "Please enter both model name and blobs path.";
            successMessage = null;
            return;
        }
        loading = true;
        error = null;
        successMessage = null;
        progress = 0;
        progressMessage = "Starting installation...";

        try {
            await InstallModel(modelName, blobsPath);
            successMessage = `Model ${modelName} installed successfully!`;
            modelName = "";
            blobsPath = "";
        } catch (e: any) {
            error = e.message || String(e);
        } finally {
            loading = false;
            progressMessage = "";
        }
    }
</script>

<div class="card">
    <h2>Install Local Model</h2>
    <div class="input-group">
        <label for="installModelName">Model Name:</label>
        <input
            type="text"
            id="installModelName"
            bind:value={modelName}
            placeholder="e.g. my-custom-model"
            disabled={loading}
        />
    </div>

    <div class="input-group">
        <label for="blobsPath">Blobs Path:</label>
        <input
            type="text"
            id="blobsPath"
            bind:value={blobsPath}
            placeholder="e.g. /path/to/downloaded/blobs"
            disabled={loading}
            on:keydown={(e) => e.key === 'Enter' && handleInstall()}
        />
    </div>

    <button class="primary-btn" on:click={handleInstall} disabled={loading}>
        {#if loading}
            <span class="spinner"></span> Installing...
        {:else}
            Install Model
        {/if}
    </button>

    {#if loading && progressMessage}
        <div class="progress-container">
            <div class="progress-info">
                <span>{progressMessage}</span>
                <span>{Math.round(progress * 100)}%</span>
            </div>
            <div class="progress-bar-bg">
                <div class="progress-bar" style="width: {progress * 100}%"></div>
            </div>
        </div>
    {/if}

    {#if error}
        <div class="error-box">
            {error}
        </div>
    {/if}

    {#if successMessage}
        <div class="success-box">
            {successMessage}
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
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 0.5rem;
    }

    .primary-btn:hover:not(:disabled) {
        background-color: var(--primary-hover, #0056b3);
    }

    .primary-btn:disabled {
        opacity: 0.6;
        cursor: not-allowed;
    }

    .error-box {
        margin-top: 1.5rem;
        padding: 1rem;
        background-color: rgba(220, 53, 69, 0.1);
        border-left: 4px solid #dc3545;
        color: #ff6b6b;
        border-radius: 4px;
    }

    .success-box {
        margin-top: 1.5rem;
        padding: 1rem;
        background-color: rgba(40, 167, 69, 0.1);
        border-left: 4px solid #28a745;
        color: #28a745;
        border-radius: 4px;
    }

    .progress-container {
        margin-top: 1.5rem;
    }

    .progress-info {
        display: flex;
        justify-content: space-between;
        margin-bottom: 0.5rem;
        font-size: 0.875rem;
        color: var(--text-muted, #a0a0b0);
    }

    .progress-bar-bg {
        width: 100%;
        height: 8px;
        background-color: var(--input-bg, #1a1a24);
        border-radius: 4px;
        overflow: hidden;
    }

    .progress-bar {
        height: 100%;
        background-color: var(--primary-color, #007bff);
        transition: width 0.3s ease;
    }

    .spinner {
        display: inline-block;
        width: 1rem;
        height: 1rem;
        border: 2px solid rgba(255, 255, 255, 0.3);
        border-radius: 50%;
        border-top-color: white;
        animation: spin 1s ease-in-out infinite;
    }

    @keyframes spin {
        to { transform: rotate(360deg); }
    }
</style>