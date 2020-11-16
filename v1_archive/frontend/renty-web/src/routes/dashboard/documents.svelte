<script>
  import DocUpload from "../../components/DocUpload.svelte";
  import DocList from "../../components/DocList.svelte";
</script>

<svelte:head>
  <title>Renty - My Documents</title>
</svelte:head>

<section>
  <h2 class="title">My Documents</h2>
  <div class="column">
    <DocUpload />
  </div>
  <div class="column">
    <DocList />
  </div>
</section>
