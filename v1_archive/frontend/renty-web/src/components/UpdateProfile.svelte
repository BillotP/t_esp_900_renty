<script lang="ts">
  import { gql } from "apollo-boost";
  import client from "../apollo_client";
  import { GraphQLEnumType } from "graphql";
  import Snackbar from "./Snackbar.svelte";

  let firstName = null;
  let lastName = null;
  let civility = null;
  let email = null;
  let phone = null;
  let phoneprefix = null;
  let phonenumber = null;
  let snackbar = false;
  let response = null;
  let isloading = false;

  const PhonePefixes = [
    {
      value: null,
      countrycode: null,
      description: "..."
    },
    {
      value: "+33",
      countrycode: "FR",
      description: "🇫🇷 +33"
    },
    {
      value: "+32",
      countrycode: "BE",
      description: "🇧🇪 +32"
    },
    {
      value: "+34",
      countrycode: "ES",
      description: "🇪🇸 +34"
    },
    {
      value: "+48",
      countrycode: "PL",
      description: "🇵🇱 +48"
    },
    {
      value: "+44",
      countrycode: "EN",
      description: "🏴󠁧󠁢󠁥󠁮󠁧󠁿 +44"
    }
  ];

  const Civility = new GraphQLEnumType({
    name: "Civility",
    values: {
      XX: {
        value: null,
        description: "..."
      },
      MS: {
        value: "MS",
        description: "👩 Ms"
      },
      MSS: {
        value: "MSS",
        description: "👧 Miss"
      },
      MR: {
        value: "MR",
        description: "👨 Mr"
      },
      MX: {
        value: "MX",
        description: "🦸 Mx"
      }
    }
  });

  const UPDATECONTACT = gql`
    mutation(
      $firstName: String
      $lastName: String
      $gender: Civility
      $email: String
      $phone: PhoneInput
    ) {
      updatecontact(
        input: {
          firstName: $firstName
          lastName: $lastName
          gender: $civility
          email: $email
          phone: $phone
        }
      ) {
        name
      }
    }
  `;
  function updateContact() {
    isloading = true;
    console.log(phoneprefix);
    if (phoneprefix && phonenumber) {
      phone = {
        value: phonenumber,
        countryCode: phoneprefix
      };
      console.log(phone);
    }
    client
      .mutate({
        mutation: UPDATECONTACT,
        variables: {
          firstName,
          lastName,
          civility,
          phone,
          email
        }
      })
      .then(r => {
        console.log(r);
        const name = r.data.updatecontact.name;
        response = `Succesfully updated ${name}`;
        isloading = false;
      })
      .catch(e => {
        console.log(e);
        response = e;
      });
    isloading = false;
    snackbar = true;
    setTimeout(() => (snackbar = false), 4000);
  }
</script>

<section>
  <div class="card">
    <div class="card-content has-text-left">
      <div class="control content">
        <div class="field">
          <!-- svelte-ignore a11y-label-has-associated-control -->
          <label class="label">Gender</label>
          <div class="select is-info">
            <select bind:value={civility}>
              {#each Civility.getValues() as vals}
                <option value={vals.value}>{vals.description}</option>
              {/each}
            </select>
          </div>
          <p class="help">How should we call you ?</p>
        </div>
        <div class="field">
          <!-- svelte-ignore a11y-label-has-associated-control -->
          <label class="label">First name</label>
          <div class="control">
            <input
              bind:value={firstName}
              class="input"
              type="text"
              placeholder="e.g John" />
          </div>
        </div>

        <div class="field">
          <!-- svelte-ignore a11y-label-has-associated-control -->
          <label class="label">Last name</label>
          <div class="control">
            <input
              bind:value={lastName}
              class="input"
              type="text"
              placeholder="e.g Doe" />
          </div>
        </div>

        <div class="field">
          <!-- svelte-ignore a11y-label-has-associated-control -->
          <label class="label">Email</label>
          <div class="control">
            <input
              class="input"
              type="email"
              bind:value={email}
              placeholder="e.g. alexsmith@gmail.com" />
          </div>
        </div>
        <div class="field">
          <!-- svelte-ignore a11y-label-has-associated-control -->
          <label class="label">Phone</label>
          <div class="field-body">
            <div class="field is-expanded">
              <div class="field has-addons">
                <p class="control">
                  <span class="select">
                    <select bind:value={phoneprefix}>
                      {#each PhonePefixes as vals}
                        <option value={vals.countrycode}>
                          {vals.description}
                        </option>
                      {/each}
                    </select>
                  </span>
                </p>
                <p class="control is-expanded">
                  <input
                    class="input"
                    type="tel"
                    bind:value={phonenumber}
                    placeholder="*Your phone number" />
                </p>
              </div>
              <p class="help has-text-grey">*Do not enter the first zero</p>
            </div>
          </div>
        </div>
      </div>
      <div class="control">
        <button
          on:click|preventDefault={updateContact}
          class:is-loading={isloading}
          class="button is-primary">
          Update
        </button>
        <button class="button is-success">
          <span class="icon is-small is-left">
            <i class="fas fa-check" />
          </span>
          &nbsp; Submit for validation
        </button>

      </div>
    </div>
  </div>
  {#if snackbar}
    <Snackbar text={response} />
  {/if}
</section>
