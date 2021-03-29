import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import store from './store';
import "./plugins/components";
import "./plugins/forms";
import vuetify from './plugins/vuetify';
import "./plugins/forms";
import { ApolloClient } from 'apollo-client'
import {ApolloLink, concat} from 'apollo-link'
import { createUploadLink } from 'apollo-upload-client'
import { InMemoryCache } from 'apollo-cache-inmemory'
import VueApollo from 'vue-apollo'

Vue.use(VueApollo)

// HTTP connection to the API
const httpLink = createUploadLink({
  // You should use an absolute URL here
  uri: 'http://localhost:8080/query',
})

const authMiddleware = new ApolloLink((operation, forward) => {
  // add the authorization to the headers
  operation.setContext({
    headers: {
      authorization: localStorage.getItem('token') || null,
    }
  });
  return forward(operation);
})

// Cache implementation
const cache = new InMemoryCache()

// Create the apollo client
export const apolloClient = new ApolloClient({
  link: concat(authMiddleware, httpLink),
  cache,
})

const apolloProvider = new VueApollo({
  defaultClient: apolloClient,
})

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  apolloProvider,
  vuetify,
  render: h => h(App)
}).$mount('#app');
