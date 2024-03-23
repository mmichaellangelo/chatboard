<script lang="ts">

    interface MessageData {
        accountName: string;
        messageBody: string;
    }

    export let data: MessageData
</script>

<div class="message_container">
    <h3>{data.accountName}</h3>
    <p>{data.messageBody}</p>
</div>