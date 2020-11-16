<script lang="ts">
  import client from "../apollo_client.js";
  import { userstore } from "../store.js";
  import { onMount } from "svelte";
  import { gql } from "apollo-boost";
  let username = null;
  let errormsg = null;

  const GetGender = (gender) => {
    const values = {
      XX: {
        value: null,
        description: "...",
      },
      female: {
        value: "MS",
        description: "👩 Ms",
      },
      lady: {
        value: "MSS",
        description: "👧 Miss",
      },
      male: {
        value: "MR",
        description: "👨 Mr",
      },
      other: {
        value: "MX",
        description: "🦸 Mx",
      },
    };
    return values[gender] || values['XX'];
  };
  const USER = gql`
    query($id: String!) {
      user(id: $id) {
        pseudo
        gender
        firstName
        lastName
        email {
          value
          valid
        }
        phone {
          countryCode
          value
        }
        createdAt
        updatedAt
      }
    }
  `;
  const UserQuery = client.watchQuery({
    query: USER,
    variables: {
      id: $userstore.id,
    },
  });
  onMount(() => UserQuery.startPolling(6000));
</script>

<section>
  {#await $UserQuery}
    <span class="icon"> <i class="fas fa-spinner fa-spin is-primary" /> </span>
  {:then results}
    {#if results}
      <div class="card">
        <div class="card-content">
          <div class="media">
            <div class="media-left">
              <figure class="image is-48x48">
                <img src="https://i.imgur.com/47PuEVC.png" alt="Placeholder" />
              </figure>
            </div>
            <div class="media-content">
              <p class="title is-4">
                {`${GetGender(results.data.user.gender).value} ${results.data.user.pseudo}`}
              </p>
              {#if results.data.user.firstName && results.data.user.lastName}
                <p class="subtitle is-6">
                  {`${results.data.user.firstName} ${results.data.user.lastName}`}
                </p>
              {/if}
            </div>
          </div>
          <div class="content has-text-left">
            {#if results.data.user.email}
              <p class="subtitle is-6">
                Email :
                {results.data.user.email.value}
                {#if results.data.user.email.verified}
                  <span class="icon has-text-success">
                    <i class="fas fa-check-square" />
                  </span>
                {:else}
                  <span class="icon has-text-warning">
                    <i class="fas fa-exclamation-triangle" />
                  </span>
                {/if}
              </p>
            {/if}
            {#if results.data.user.phone}
              <p class="subtitle is-6">
                Phone :
                {`${results.data.user.phone.value}`}
                {#if results.data.user.phone.verified}
                  <span class="icon has-text-success">
                    <i class="fas fa-check-square" />
                  </span>
                {:else}
                  <span class="icon has-text-warning">
                    <i class="fas fa-exclamation-triangle" />
                  </span>
                {/if}
              </p>
            {/if}
            <p>
              Member since :
              <time datetime={results.data.user.createdAt}>
                {new Date(results.data.user.createdAt).toLocaleString()}
              </time>
            </p>
            <p>
              Last Update :
              <time datetime={results.data.user.updatedAt}>
                {new Date(results.data.user.updatedAt).toLocaleString()}
              </time>
            </p>
            <p class="help">
              *Is confirmed
              <span class="icon has-text-success">
                <i class="fas fa-check-square" />
              </span>
              / need confirmation
              <span class="icon has-text-warning">
                <i class="fas fa-exclamation-triangle" />
              </span>
            </p>
          </div>
        </div>
      </div>
    {/if}
  {/await}
</section>
