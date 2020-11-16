<script lang="ts">
  import client from "../apollo_client";
  import { onMount } from "svelte";
  import { gql } from "apollo-boost";
import RentOffers from "./RentOffers.svelte";
import type { RentOffer } from "../models/rentoffer";

  let files: any = null;
  let errmsg: string = null;
  const MyDocuments = gql`
    query {
      mydocuments {
        label {
          text
        }
        asset {
          type
          name
          url
        }
        doctype
        createdAt
        updatedAt
      }
    }
  `;
  const MydocQuery = client.watchQuery({
    query: MyDocuments,
  });
  onMount(() => MydocQuery.startPolling(4000));
</script>

<div class="column content">
  <div class="list is-hoverable" style="max-height: 70vh;overflow-y: scroll;">
    {#await $MydocQuery}
      <progress class="progress is-small is-info" max="100">60%</progress>
    {:then results}
      {#if results}
        {#each results.data.mydocuments as file, i}
          <div class="card list-item">
            <header class="card-header">
              <p class="card-header-title">{file.label.text}</p>
            </header>
            <div class="card-content">
              <div class="content">
                {file.doctype}
                -
                {file.asset.type}
                <br />
                Last updated on
                <time datetime={file.updatedAt}>
                  {file.updatedAt.replace('T', ' ')}
                </time>
              </div>
            </div>
            <footer class="card-footer is-centered">
              <div class="field is-grouped">
                <button href="#" class="card-footer-item button is-rounded">
                  Share
                </button>
                <button href="#" class="card-footer-item button is-rounded">
                  Update
                </button>
                <button href="#" class="card-footer-item button is-rounded">
                  Delete
                </button>
              </div>
            </footer>
          </div>
        {/each}
      {/if}
    {:catch error}
      Error:
      {error}
    {/await}
  </div>
</div>
