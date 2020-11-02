<script>
  import { onMount } from "svelte";
  import client from "../apollo_client.js";
  import { gql } from "apollo-boost";
  let isloading = true;

  const RENTOFFERS = gql`
    query {
      rentoffers {
        title {
          value
        }
        location {
          geoJson {
            geometry {
              coordinates
            }
          }
        }
      }
    }
  `;

  onMount(() => {
    var mymap = L.map("offersMap", { zoomControl: false }).setView(
      [44.8333, -0.5667],
      13
    );
    var token =
      "pk.eyJ1IjoiZGF2ZS1sb3BldXIiLCJhIjoiY2s2dzVkMTl2MDhlOTNlcXZ1N2o5eHNkNiJ9.6nm9cmaZzJU85LTjo0oR6Q";
    L.tileLayer(
      "https://api.mapbox.com/styles/v1/{id}/tiles/{z}/{x}/{y}?access_token=" +
        token,
      {
        maxZoom: 18,
        attribution:
          'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a>',
        id: "mapbox/streets-v11",
        tileSize: 512,
        zoomOffset: -1,
      }
    ).addTo(mymap);
    client
      .query({
        query: RENTOFFERS,
      })
      .then((r) => {
        for (let el = 0; el < r.data.rentoffers.length; ++el) {
          const offer = r.data.rentoffers[el];
          new L.marker([
            offer.location.geoJson.geometry.coordinates[1],
            offer.location.geoJson.geometry.coordinates[0],
          ])
            .bindPopup(offer.title[0].value)
            .addTo(mymap);
        }
        isloading = false;
      })
      .catch((e) => {
        isloading = false;
        console.log(e);
      });
  });
</script>

<style>
  #offersMap {
    height: 100vh;
    width: 100%;
  }
</style>

{#if isloading}
  <progress
    style="margin-top: 10px;"
    class="progress is-small is-primary"
    max="100">
    35%
  </progress>
{/if}
<div id="offersMap" style="height: 100vh;width: 100%;z-index: -1;" />
