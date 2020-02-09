import Vue from 'vue';
import App from './App.vue';
import router from './router';

import * as firebase from 'firebase/app';
import 'firebase/auth';

const firebaseConfig = {
    apiKey: process.env.VUE_APP_FIREBASE_APIKEY,
    authDomain: process.env.VUE_APP_FIREBASE_AUTH_DOMAIN,
    databaseURL: process.env.VUE_APP_FIREBASE_DATABASE_URL,
    projectId: process.env.VUE_APP_FIREBASE_PROJECT_ID,
    storageBucket: process.env.VUE_APP_FIREBASE_STORAGE_BUCKET,
    messagingSenderId: process.env.VUE_APP_FIREBASE_MESSAGINGSENDER_ID,
    appId: process.env.VUE_APP_FIREBASE_APP_ID,
};

firebase.initializeApp(firebaseConfig);

Vue.config.productionTip = false;

let app: any;

firebase.auth().onAuthStateChanged(user => {
    if (!app) {
        new Vue({
            router,
            render: h => h(App),
        }).$mount('#app');
    }
});
