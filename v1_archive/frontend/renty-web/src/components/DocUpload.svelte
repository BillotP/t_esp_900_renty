<script lang="ts">
  import { gql } from "apollo-boost";
  import client from "../apollo_client.js";
  import Snackbar from "./Snackbar.svelte";
  import { GraphQLEnumType } from "graphql";

  let files;
  let snackbar = false;
  let isloading = false;
  let respmessage = null;
  let documenttype;

  const DocumentTypes = new GraphQLEnumType({
    name: "DocumentType",
    values: {
      IDDOC: {
        value: "IDDOC"
      },
      PROOFOFINCOME: {
        value: "PROOFOFINCOME"
      },
      EMPLOYMENTPROOF: {
        value: "EMPLOYMENTPROOF"
      },
      PROOFOFADDRESS: {
        value: "PROOFOFADDRESS"
      }
    }
  });
  const UploadAsset = gql`
    mutation($documentType: DocumentType!, $label: String!, $file: Upload!) {
      uploadasset(
        input: { documentType: $documentType, label: $label, file: $file }
      )
    }
  `;
  async function uploadAsset() {
    isloading = true;
    try {
      console.log(documenttype);
      const uploadMutation = await client.mutate({
        mutation: UploadAsset,
        variables: {
          documentType: documenttype,
          label: files[0].name,
          file: files[0]
        }
      });
      const resp = (await uploadMutation.data).uploadasset;
      console.log(resp);
      isloading = false;
      respmessage = `Sucessfully uploaded ${files[0].name}`;
    } catch (e) {
      console.log(e);
      isloading = false;
      respmessage = `Ohoh it's their is a problem with your file, please try again later`;
    }
    snackbar = true;
    setTimeout(() => (snackbar = false), 4000);
  }
</script>

<div class="column">
  <div class="file has-name is-fullwidth">
    <label class="file-label">
      <input class="file-input" type="file" bind:files name="resume" />

      <span class="file-cta">
        <span class="file-icon">
          <i class="fas fa-upload" />
        </span>
        <span class="file-label">Choose a file…</span>
      </span>
      <span class="file-name">
        {#if files && files[0]}{files[0].name}{:else}...{/if}
      </span>
    </label>
  </div>
</div>
<div class="select is-success">
  <select bind:value={documenttype}>
    {#each DocumentTypes.getValues() as vals}
      <option value={vals.value}>{vals.value}</option>
    {/each}
  </select>
</div>
{#if !isloading}
  <button
    disabled={files == null}
    class="button is-success"
    on:click|preventDefault={uploadAsset}>
    Upload
  </button>
{:else}
  <button class="button is-primary is-loading" />
{/if}
{#if snackbar}
  <Snackbar text={respmessage} />
{/if}
