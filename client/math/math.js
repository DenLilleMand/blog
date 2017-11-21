import Vue from 'vue'
import VueRouter from 'vue-router'
import Math from 'math/math.vue'
import MainMenu from 'components/mainmenu.vue'
import Subject from 'math/subject.vue'
import Post from 'math/post.vue'

Vue.component('mainmenu', MainMenu);

const routes = [{ 
	path: '/:subject', 
	component: Subject, 
	children: [{
	    path: '/:post',
	    component: Post,
	    props: true
	}]
}];

const router = new VueRouter({
    routes,
});

Vue.use(VueRouter);

new Vue({
    router: router,
    el: '#math',
    template: '<Math />',
    components: { Math }
});
