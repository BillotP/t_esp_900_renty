// import "isomorphic-fetch";
import { createUploadLink } from "apollo-upload-client";
import { InMemoryCache } from "apollo-cache-inmemory";
import { setContext } from "apollo-link-context";
import ApolloClient from "apollo-client";
import { ApolloLink } from "apollo-link";
import fetch from "node-fetch";

const authLink = setContext((_, { headers }) => {
  // get the authentication token from local storage if it exists
  const token = localStorage.getItem("token");
  // return the headers to the context so httpLink can read them
  return {
    headers: {
      ...headers,
      Authorization: token ? `Bearer ${token}` : "",
    },
  };
});

const link = ApolloLink.from([
  authLink,
  createUploadLink({ fetch, uri: "http://api.192-168-1-34.sslip.io/api/query" }),
]);

const client = new ApolloClient({
  // onError: ({ networkError, graphQLErrors }) => {
  //   console.log("graphQLErrors :", graphQLErrors);
  //   console.log("networkError :", networkError);
  //   return { networkError, graphQLErrors };
  // },
  link,
  cache: new InMemoryCache(),
});

export default client;
