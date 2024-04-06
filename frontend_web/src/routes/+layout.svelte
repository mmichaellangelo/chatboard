<script>

    import Header from "$lib/elements/Header/Header.svelte";
    import { sessionStore } from "$lib/state/userStore";
    

</script>
<Header data={{username: $sessionStore.username}}/>
<slot/>