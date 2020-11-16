<script lang="ts">
  import client from "../apollo_client.js";
  import RentOffer from "./RentOffer.svelte";
  import { onMount } from "svelte";
  import { gql } from "apollo-boost";

  const RentOffers = gql`
    query {
      rentoffers {
        surface {
          value
          unit
        }
        price {
          value
          currency
        }
        title {
          value
        }
        description {
          value
        }
        assets {
          url
        }
      }
    }
  `;
  const RentOffersQuery = client.watchQuery({
    query: RentOffers,
  });
  // const RentOffersQuery = GET_RENTOFFERS();
  onMount(() => {
    RentOffersQuery.refetch().catch((e) => console.log(e));
    RentOffersQuery.startPolling(10000);
  });
</script>

<div class="column">
  {#await $RentOffersQuery}
    <span class="icon"> <i class="fas fa-spinner fa-spin is-primary" /> </span>
  {:then results}
    {#if results}
      <!-- <h5>{JSON.stringify(results)}</h5> -->
      <div class="list" style="max-height: 90vh;overflow-y: scroll">
        {#each results.data.rentoffers as offer}
          <div class="box">
            <RentOffer {offer} />
          </div>
        {/each}
      </div>
    {/if}
  {/await}
</div>
