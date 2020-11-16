<script lang="ts">
  import client from "../apollo_client.js";
  import { gql } from "apollo-boost";
  // import { onMount } from "svelte";
  import { goto } from "@sapper/app";
  import { userstore, islogged, google_clientId } from "../store.js";
  import { GoogleAuth } from "@beyonk/svelte-social-auth";

  let isloading = false;
  let googleUid = null;
  let username = "";
  let userpassword = "";
  let userconfirm = 0;
  let results = null;
  let errmsg = null;
  let token = "";
  // let fcmToken = null;

  const LOGIN = gql`
    mutation($pseudo: String!, $password: String!) {
      login(input: { pseudo: $pseudo, password: $password }) {
        token
        user {
          id
        }
      }
    }
  `;
  const SOCIALLOGIN = gql`
    query($IDToken: String!, $notificationID: String) {
      socialLogin(
        input: {
          IDToken: $IDToken
          provider: GOOGLE
          notificationID: $notificationID
        }
      ) {
        token
      }
    }
  `;
  const CONFIRM = gql`
    query($code: Int!) {
      confirm(code: $code) {
        token
      }
    }
  `;

  // onMount(async () => {
  //   try {
  //     const token = await firebase.messaging().getToken();
  //     fcmToken = token;
  //   } catch (e) {
  //     console.log(e);
  //   }
  // });

  function LoginSuccess(r) {
    console.log(r);
    if (r.data) {
      results = r.data.login ? r.data.login : r.data.socialLogin;
      token = results.token;
      isloading = false;
      localStorage.setItem("token", token);
      islogged.set(true);
      userstore.set({ id: results.user.id });
      goto("/dashboard");
    } else {
      console.log("...");
    }
  }
  function SocialLogin() {
    isloading = true;
    client
      .query({
        query: SOCIALLOGIN,
        variables: {
          IDToken: googleUid,
          notificationID: fcmToken,
        },
      })
      .then((r) => LoginSuccess(r))
      .catch((e) => {
        isloading = false;
        console.log(e.message);
        errmsg = "Invalid input please try again";
      });
  }
  function login() {
    isloading = true;
    client
      .mutate({
        mutation: LOGIN,
        variables: {
          pseudo: username,
          password: userpassword,
        },
      })
      .then((r) => LoginSuccess(r))
      .catch((e) => {
        isloading = false;
        console.log(e.message);
        errmsg = "Invalid input please try again";
      });
  }
  function googleLogin(googleUser) {
    const profile = googleUser.getBasicProfile();
    googleUid = googleUser.getAuthResponse().id_token;
    SocialLogin();
  }
  async function confirm() {
    const confirmQuery = await client.query({
      query: CONFIRM,
      variables: {
        code: userconfirm,
      },
    });
    const rep = await confirmQuery.result();
    results = rep.data.confirm;
    localStorage.setItem("token", results.token);
    await goto("/dashboard");
  }
  function isenter(event) {
    if (event.keyCode == 13) {
      login();
    }
  }
</script>

<svelte:head>
  <title>Renty - Login 🔓</title>
</svelte:head>
<section class="hero is-primary is-fullheight-with-navbar">
  <div class="hero-body">
    <div class="column is-three-fifths is-offset-one-fifth">
      <div class="box">
        {#if results}
          {#if results.otp}
            <form on:submit|preventDefault={confirm}>
              <h4>Please enter 2FA code</h4>
              <div class="field">
                <p class="control has-icons-left has-icons-right">
                  <input
                    type="number"
                    placeholder="2FA"
                    bind:value={userconfirm} />
                  <span class="icon is-small is-left">
                    <i class="fas fa-user" />
                  </span>
                  <span class="icon is-small is-right">
                    <i class="fas fa-check" />
                  </span>
                </p>
              </div>
              <div class="has-text-centered">
                {#if !isloading}
                  <button class="button" method="submit">Submit</button>
                {:else}<i class="fas fa-spinner fa-spin" />{/if}
              </div>
            </form>
          {/if}
        {:else if !isloading}
          <div class="field">
            <div class="control has-icons-left has-icons-right">
              <input
                class="input"
                placeholder="Username"
                bind:value={username} />
              <span class="icon is-small is-left">
                <i class="fas fa-user" />
              </span>
              <span class="icon is-small is-right" style="display: none;">
                <i class="fas fa-check" />
              </span>
            </div>
          </div>

          <div class="field">
            <div class="control has-icons-left">
              <input
                class="input"
                type="password"
                placeholder="Password"
                on:keydown={isenter}
                bind:value={userpassword} />
              <span class="icon is-small is-left">
                <i class="fas fa-lock" />
              </span>
            </div>
            {#if errmsg}
              <p class="help is-danger has-text-centered">{errmsg}</p>
            {/if}
          </div>
          <button class="button is-primary" on:click|preventDefault={login}>
            Login
          </button>
          <div class="column has-text-centered">
            <GoogleAuth
              text="Login with Google"
              clientId={google_clientId.client_id}
              on:auth-success={(e) => googleLogin(e.detail.user)} />
          </div>
        {:else}
          <div class="column has-text-centered">
            <i
              class="fas fa-spinner fa-spin is-primary is-large"
              style="margin-left: 10px;" />
          </div>
        {/if}
      </div>
    </div>
  </div>
</section>
