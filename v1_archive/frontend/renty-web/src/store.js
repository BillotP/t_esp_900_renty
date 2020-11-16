import { writable } from "svelte/store";

export const islogged = writable(false);
export const google_clientId = {
  client_id:
    "500932383497-dib4hqbo88i3vkvgvv5j0ga0v81muu0i.apps.googleusercontent.com",
};

export const userstore = writable({
  id: "",
})