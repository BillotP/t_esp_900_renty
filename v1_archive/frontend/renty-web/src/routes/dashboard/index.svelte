<script>
  import { goto } from "@sapper/app";
  import { onMount } from "svelte";
  import { islogged } from "../../store.js";
  import UserProfile from "../../components/UserProfile.svelte";
  import UpdateProfile from "../../components/UpdateProfile.svelte";

  onMount(() => {
    if (islogged == false) {
      goto("/");
    }
  });
</script>

<svelte:head>
  <title>Dashboard | Renty</title>
</svelte:head>

<section>
  <div class="column is-vcentered is-centered">
    <UserProfile />
  </div>
  <div class="column is-vcentered is-centered">
    <UpdateProfile />
  </div>
</section>
