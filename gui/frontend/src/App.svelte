<script lang="ts">
    import logo from "./assets/images/ollama.png";
    import GetDownloadLink from "./components/steps/get-download-link.svelte";
    import Install from "./components/steps/install.svelte";

    let chosenStep: "GetLink" | "Install" = "GetLink";
</script>

<main>
    <header>
        <img alt="Ollama Direct Download logo" id="logo" src="{logo}">
        <h1>Ollama Model Manager</h1>
    </header>

    <div class="nav">
        <button
            class:active={chosenStep === 'GetLink'}
            on:click={() => chosenStep = 'GetLink'}>
            Get Link
        </button>
        <button
            class:active={chosenStep === 'Install'}
            on:click={() => chosenStep = 'Install'}>
            Install
        </button>
    </div>

    <div class="content">
        {#if chosenStep === 'GetLink'}
            <GetDownloadLink/>
        {:else if chosenStep === 'Install'}
            <Install/>
        {/if}
    </div>
</main>

<style>
    main {
        max-width: 800px;
        margin: 0 auto;
        padding: 2rem;
        display: flex;
        flex-direction: column;
        align-items: stretch;
    }

    header {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-bottom: 2rem;
    }

    #logo {
        height: 100px;
        margin-bottom: 1rem;
    }

    h1 {
        font-size: 2rem;
        color: var(--text-color, #ffffff);
        margin: 0;
    }

    .nav {
        display: flex;
        justify-content: center;
        gap: 1rem;
        margin-bottom: 2rem;
    }

    .nav button {
        background-color: var(--input-bg, #1a1a24);
        color: var(--text-muted, #a0a0b0);
        border: 1px solid var(--border-color, #4a4a5a);
        padding: 0.75rem 1.5rem;
        border-radius: 4px;
        font-size: 1rem;
        cursor: pointer;
        transition: all 0.2s;
    }

    .nav button:hover {
        background-color: var(--hover-bg, #2a2a35);
        color: var(--text-color, #ffffff);
    }

    .nav button.active {
        background-color: var(--primary-color, #007bff);
        color: white;
        border-color: var(--primary-color, #007bff);
    }

    .content {
        width: 100%;
    }
</style>
