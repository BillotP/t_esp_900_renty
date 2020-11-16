<script lang="ts">
  import { goto } from "@sapper/app";
  import { gql } from "apollo-boost";
  import client from "../apollo_client.js";
  import { google_clientId } from "../store.js";
  import Snackbar from "../components/Snackbar.svelte";
  import { GoogleAuth } from "@beyonk/svelte-social-auth";

  let pseudo = null;
  let phone = null;
  let email = null;
  let password = null;
  let firstName = null;
  let lastName = null;
  let gender = null;
  let results = null;
  let response = null;
  let snackbar = false;
  let checkedTos = false;
  let google_name = null;
  let google_email = null;
  let google_idtoken = null;
  let respmessage = null;
  let errormessage = null;

  const REGISTER = gql`
    mutation(
      $pseudo: String!
      $firstName: String!
      $lastName: String!
      $gender: String!
      $password: String!
    ) {
      register(
        input: {
          pseudo: $pseudo
          firstName: $firstName
          lastName: $lastName
          gender: $gender
          password: $password
        }
      ) {
        id
        pseudo
      }
    }
  `;

  const SOCIALREGISTER = gql`
    mutation($IDToken: String!, $notificationID: String) {
      socialregister(input: { IDToken: $IDToken, provider: GOOGLE }) {
        name
      }
    }
  `;

  function validateEmail(email) {
    var checkEmail = /\S+@\S+\.\S+/;
    return checkEmail.test(String(email).toLowerCase());
  }

  function quickCheck() {
    if (email && !validateEmail(email)) {
      email = null;
      errormessage = `
      Please enter a valid email if you choose to submit one.
      `;
      return true;
    } else if (name.length < 3 || password.length < 3) {
      errormessage =
        "Your input is invalid, name and password must be longer than 3 characters.";
      errormessage += "\nPlease review them before sending again.";
      return false;
    }
    return true;
  }
  function registerSuccess(r) {
    results = r.data.register ? r.data.register : r.data.socialregister;
    response = `Sucessfully register ${results.pseudo}`;
    goto("/login");
  }

  function registerError(e) {
    console.error(e);
    errormessage = String(e.message).replace("GraphQL error:", "");
  }
  async function socialRegister() {
    let notificationID;
    if (!google_idtoken) return;
    try {
      notificationID = await firebase.messaging().getToken();
    } catch (e) {}
    try {
      let res = await client.mutate({
        mutation: SOCIALREGISTER,
        variables: {
          IDToken: google_idtoken,
          notificationID,
        },
      });
      registerSuccess(res);
    } catch (error) {
      console.log(error);
      registerError(error);
    }
  }
  function googlesignup(googleUser) {
    google_idtoken = googleUser.getAuthResponse().id_token;
    checkedTos = true;
    socialRegister();
  }

  async function register() {
    let notificationID;
    if (!quickCheck) return;
    try {
      notificationID = await firebase.messaging().getToken();
    } catch (e) {}
    client
      .mutate({
        mutation: REGISTER,
        variables: {
          pseudo,
          firstName,
          lastName,
          gender,
          password,
        },
      })
      .then((r) => registerSuccess(r))
      .catch(registerError);
  }
</script>

<svelte:head>
  <title>Renty - Register 🖋️</title>
</svelte:head>
<section class="hero is-primary is-fullheight-with-navbar">
  <div class="hero-body">
    <div class="column is-three-fifths is-offset-one-fifth">
      <div class="box">
        <div class="field">
          <div class="control has-icons-left has-icons-right">
            <input
              class={!errormessage ? 'input' : 'input is-danger'}
              placeholder="Pseudo*"
              bind:value={pseudo} />
            <span class="icon is-small is-left">
              <i class="fas fa-user" />
            </span>
            <span class="icon is-small is-right" style="display: none;">
              <i class="fas fa-check" />
            </span>
          </div>
          <p class="help is-success" style="display: none;">
            This pseudo is available
          </p>
        </div>
        <div class="field">
          <div class="control has-icons-left has-icons-right">
            <input
              class={!errormessage ? 'input' : 'input is-danger'}
              placeholder="First name*"
              bind:value={firstName} />
            <span class="icon is-small is-left">
              <i class="fas fa-user" />
            </span>
            <span class="icon is-small is-right" style="display: none;">
              <i class="fas fa-check" />
            </span>
          </div>
        </div>
        <div class="field">
          <div class="control has-icons-left has-icons-right">
            <input
              class={!errormessage ? 'input' : 'input is-danger'}
              placeholder="Last name*"
              bind:value={lastName} />
            <span class="icon is-small is-left">
              <i class="fas fa-user" />
            </span>
            <span class="icon is-small is-right" style="display: none;">
              <i class="fas fa-check" />
            </span>
          </div>
        </div>
        <div class="field">
          <div class="control has-icons-left has-icons-right">
            <input
              class={!errormessage ? 'input' : 'input is-danger'}
              placeholder="Gender*"
              bind:value={gender} />
            <span class="icon is-small is-left">
              <i class="fas fa-user" />
            </span>
            <span class="icon is-small is-right" style="display: none;">
              <i class="fas fa-check" />
            </span>
          </div>
        </div>
        <div class="field">
          <div class="control has-icons-left has-icons-right">
            <input class="input" placeholder="Email" bind:value={email} />
            <span class="icon is-small is-left">
              <i class="fas fa-envelope" />
            </span>
            <span class="icon is-small is-right" style="display: none;">
              <i class="fas fa-check" />
            </span>
          </div>
          <p class="help is-danger" style="display: none;">
            This email is invalid
          </p>
        </div>
        <div class="field">
          <p class="control has-icons-left has-icons-right">
            <input
              class="input"
              type="tel"
              placeholder="Phone Number"
              bind:value={phone} />
            <span class="icon is-small is-left">
              <i class="fas fa-phone" />
            </span>
            <span class="icon is-small is-right" style="display: none;">
              <i class="fas fa-check" />
            </span>
          </p>
        </div>

        <div class="field">
          <p class="control has-icons-left">
            <input
              class={!errormessage ? 'input' : 'input is-danger'}
              type="password"
              placeholder="Password*"
              bind:value={password} />
            <span class="icon is-small is-left">
              <i class="fas fa-lock" />
            </span>
          </p>
        </div>
        <p class="is-size-6 has-text-grey" style="margin-bottom: 5px;">
          *Required fields are marked with an asterisk
        </p>
        {#if errormessage}
          <p class="notification">{errormessage}</p>
        {/if}
        <label class="checkbox">
          <input type="checkbox" bind:checked={checkedTos} />
          I agree to the
          <a href="/blog">terms and conditions</a>
        </label>
        <div class="column has-text-centered">
          <GoogleAuth
            text="Register with Google"
            clientId={google_clientId.client_id}
            on:auth-success={(e) => googlesignup(e.detail.user)} />
          <p class="is-size-6 has-text-grey-light">
            By signing up with Google services, you automatically accept our
            terms and conditions.
          </p>
        </div>
        <div class="has-text-left" style="padding-top: 10px;">
          <button
            disabled={!checkedTos}
            on:click|preventDefault={register}
            class="button is-primary"
            method="submit">
            Register
          </button>
        </div>
      </div>
    </div>
  </div>
  {#if snackbar}
    <Snackbar text={respmessage} />
  {/if}
</section>
