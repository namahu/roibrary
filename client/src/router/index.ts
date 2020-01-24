import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/SiteTop.vue';

import * as firebase from 'firebase/app';

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        name: 'home',
        component: Home,
        meta: { isPublic: true },
    },
    {
        path: '/about',
        name: 'about',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
    },
    {
        path: '/addnewbook',
        name: 'addnewbook',
        component: () => import(/* webpackChunkName: "addnewbook" */ '../views/BookAdd.vue'),
    },
    {
        path: '/signin',
        name: 'signin',
        meta: { isPublic: true },
        component: () => import(/* webpackChunkName: "signin" */ '../views/Signin.vue'),
    },
];

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes,
});

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.isPublic)) {
        next();
        return;
    }

    const currentUser = firebase.auth().currentUser;
    if (!currentUser) {
        next('/signin');
    } else {
        next();
    }
});

export default router;
