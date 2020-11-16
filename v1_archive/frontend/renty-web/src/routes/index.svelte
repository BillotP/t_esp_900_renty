<script lang="ts">
  import Map from "../components/Map.svelte";
</script>

<style>
</style>

<svelte:head>
  <title>Renty - Your personal real estate agency 🏙️</title>
</svelte:head>
<section class="hero is-fullheight-with-navbar">
  <Map />
  <div class="is-overlay" style="padding-top: 35vw;padding-bottom: 35vw">
    <div class="card" style="max-width: 50vw;">
      <div class="card-content">
        <p class="title">“Be ready, you will move in your home in no time !”</p>
        <p class="subtitle">Renty</p>
      </div>
      <footer class="card-footer">
        <p class="card-footer-item">
          <span> 🚀 Quick <a href="/register">register</a> </span>
        </p>
        <p class="card-footer-item">
          <span>
            🏘️ See available
            <a href="/alloffers">offers</a>
            near me
          </span>
        </p>
      </footer>
    </div>
  </div>
</section>
