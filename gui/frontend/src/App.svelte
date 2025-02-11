<script lang="ts">
    import logo from './assets/images/logo-universal.png'
    import {Greet} from '../wailsjs/go/main/App.js'
    import GetDownloadLink from './components/steps/get-download-link.svelte';

    let resultText: string = "Please enter your name below 👇"
    let name: string
    let chosenStep: "GetLink" | "Install" | undefined = undefined

    function greet(): void {
        Greet(name).then(result => resultText = result)
    }
</script>

<main>
    <img alt="Wails logo" id="logo" src="{logo}">
    <div class="result" id="result">{resultText}</div>
    <div class="input-box" id="input">
        <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>
        <button class="btn" on:click={greet}>Greet</button>
    </div>
    {#if (chosenStep === undefined) }
        <div>
            <button on:click={() => chosenStep = 'GetLink'}>Get Link</button>
            <button on:click={() => chosenStep = 'Install'}>Install</button>
        </div>
    {/if}
    {#if chosenStep === 'GetLink' }
        <GetDownloadLink/>
    {/if}
    {#if chosenStep === 'Install' }
        <p>install</p>
    {/if}
</main>

<style>

</style>
