<script lang="js">
  export let segment;
  import { onMount } from "svelte";
  import { goto } from "@sapper/app";
  import { islogged, google_clientId } from "../store.js";
  import client from "../apollo_client.js";
  // let islogged_value;

  // const unsubscribe = islogged.subscribe(value => {
  //   islogged_value = value;
  // });
  const parseJwt = (token) => {
    try {
      return JSON.parse(atob(token.split(".")[1]));
    } catch (e) {
      return null;
    }
  };
  // onDestroy(unsubscribe);
  onMount(() => {
    // let google = require("googleapis");
    gapi.load("auth2", () => gapi.auth2.init(google_clientId));
    islogged.update((v) => {
      let token = localStorage.getItem("token");
      if (!token) return (v = false);
      let values = parseJwt(token);
      if (values == null) return (v = false);
      if (Date.now() >= values.exp * 1000) return (v = false);
      return (v = true);
    });
    // Get all "navbar-burger" elements
    const $navbarBurgers = Array.prototype.slice.call(
      document.querySelectorAll(".navbar-burger"),
      0
    );
    // Check if there are any navbar burgers
    if ($navbarBurgers.length > 0) {
      // Add a click event on each of them
      $navbarBurgers.forEach((el) => {
        el.addEventListener("click", () => {
          // Get the target from the "data-target" attribute
          const target = el.dataset.target;
          const targetEl = document.getElementById(target);

          // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
          el.classList.toggle("is-active");
          targetEl.classList.toggle("is-active");
        });
      });
    }
  });

  async function logout() {
    // let google = require("googleapis");
    await client.clearStore();
    var auth2 = gapi.auth2.getAuthInstance();
    if (auth2) {
      await auth2.signOut();
    }
    localStorage.clear();
    islogged.set(false);
    await goto("/");
  }
</script>

<style>
  nav {
    border-bottom: 1px solid rgba(255, 62, 0, 0.1);
    font-weight: 300;
    padding: 0 1em;
  }

  .selected {
    position: relative;
    display: inline-block;
  }

  .selected::after {
    position: absolute;
    content: "";
    width: calc(100% - 1em);
    height: 2px;
    background-color: #3498db;
    display: block;
    bottom: -1px;
  }

  a {
    text-decoration: none;
    padding: 1em 0.5em;
    display: block;
  }
</style>

<nav class="navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/">
      <img alt="Renty logo" src="/logo.png" width="112" height="28" />
    </a>

    <button
      role="button"
      class="navbar-burger button"
      aria-label="menu"
      aria-expanded="false"
      data-target="navBar">
      <span aria-hidden="true" />
      <span aria-hidden="true" />
      <span aria-hidden="true" />
    </button>
  </div>
  <div id="navBar" class="navbar-menu">
    <div class="navbar-start">
      <a class="navbar-item" class:selected={segment === undefined} href=".">
        home
      </a>

      <a class="navbar-item" class:selected={segment === 'about'} href="about">
        about
      </a>

      <a
        class="navbar-item"
        rel="prefetch"
        class:selected={segment === 'alloffers'}
        href="alloffers">
        All Offers
      </a>
    </div>

    <div class="navbar-end">
      <div class="navbar-item">
        <div class="buttons">
          {#if !$islogged}
            <a
              class:selected={segment === 'login'}
              href="login"
              class="button is-light">
              Sign in
            </a>
            <a
              class:selected={segment === 'register'}
              href="register"
              class="button is-primary">
              <strong>Sign up</strong>
            </a>
          {:else}
            <a
              class:selected={segment === 'dashboard'}
              href="dashboard"
              class="button is-primary">
              <strong>My Renty</strong>
            </a>
            <button
              style="margin-top: 5px;"
              on:click|preventDefault={logout}
              class="button has-text-info">
              <i class="fas fa-sign-out-alt" />
            </button>
          {/if}
        </div>
      </div>
    </div>
  </div>
</nav>
