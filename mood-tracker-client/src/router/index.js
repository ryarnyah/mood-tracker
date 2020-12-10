import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';
import Mood from '../views/GetMood.vue';
import Entry from '../views/AddEntry.vue';

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/mood/:moodId/:moodSignature',
        name: 'Mood',
        component: Mood
    },
    {
        path: '/entry/:moodId/:entryId/:entrySignature',
        name: 'Entry',
        component: Entry
    },
    {
        path: '/about',
        name: 'About',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    }
];

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
});

export default router;
