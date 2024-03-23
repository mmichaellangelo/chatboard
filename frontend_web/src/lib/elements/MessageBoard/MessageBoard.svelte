<script lang="ts">
    import Message from "../Message/Message.svelte";

    const messages = [
        {
            accountName: "anon",
            messageBody: "a message",
        },
        {
            accountName: "anon",
            messageBody: "a second message",
        },
        {
            accountName: "anon",
            messageBody: "a third message",
        },
        {
            accountName: "anon",
            messageBody: "a fourth",
        },
    ]
</script>

{#each messages as message}
    <Message data={message} />
{/each}