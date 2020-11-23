import Vue from 'vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';
import fr from 'vuetify/src/locale/fr';
import VueTheMask from 'vue-the-mask'

Vue.use(Vuetify);
Vue.use(VueTheMask);

export default new Vuetify({
  theme: {
    options: {
      customProperties: true,
    },
    themes: {
      light: {
        primary: '#007f96',
        secondary: '#424242',
        accent: '#82B1FF',
        error: '#D1281F',
        info: '#1282D7',
        success: '#4CAF50',
        warning: '#E25A1B',
        background: '#F4F3F0',
        header: '#E4E4E4'
      },
      dark: {
        primary: '#007f96',
        secondary: '#424242',
        accent: '#82B1FF',
        error: '#D1281F',
        info: '#1282D7',
        success: '#4CAF50',
        warning: '#E25A1B',
        background: '#424242',
        header: "#353535"
      }
    },
  },
  lang: {
    locales: { fr },
    current: 'fr',
  },
  icons: {
    iconfont: 'fa4' || 'mdi',
  },
});
