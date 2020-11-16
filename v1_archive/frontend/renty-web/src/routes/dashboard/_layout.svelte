<script>
  import DashboardMenu from "../../components/DashboardMenu.svelte";
  export let segment;
</script>

<main>
  <section>
    <div class="columns is-mobile" style="height: 80vw;">
      <div class="column is-one-fifth">
        <DashboardMenu {segment} />
      </div>
      <div class="column has-text-centered">
        <slot />
      </div>
    </div>
  </section>
</main>
