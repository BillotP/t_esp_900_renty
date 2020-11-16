<script lang="ts">
  import { fly } from "svelte/transition";
  export let text;
</script>

<style>
  .snackbar {
    min-width: 250px; /* Set a default minimum width */
    margin-left: -125px; /* Divide value of min-width by 2 */
    background-color: #333; /* Black background color */
    color: #fff; /* White text color */
    text-align: center; /* Centered text */
    border-radius: 2px; /* Rounded borders */
    padding: 16px; /* Padding */
    position: fixed; /* Sit on top of the screen */
    z-index: 1; /* Add a z-index if needed */
    left: 50%; /* Center the snackbar */
    bottom: 30px; /* 30px from the bottom */
  }
</style>

<div class="snackbar" transition:fly={{ y: 40, duration: 1000 }}>{text}</div>
