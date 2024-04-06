<script lang="ts">
    interface HeaderData {
        username: string;
    }
    export let data: HeaderData;
</script>

<div id="header_container">
    <h1>ChatBoard</h1>
    {#if data.username != ""}
    <p>Logged in as {data.username}</p>
    <a href="/logout">Log Out</a>
    {/if}
</div>