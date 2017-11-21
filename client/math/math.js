import Vue from 'vue'
import VueRouter from 'vue-router'
import Math from 'math/math.vue'
import MainMenu from 'components/mainmenu.vue'
import LinearAlgebra from 'math/linearalgebra.vue'
import ImageProcessing from 'math/imageprocessing.vue'
import ModellingAndAnalysisOfData from 'math/modellingandanalysisofdata.vue'

Vue.component('mainmenu', MainMenu);

const routes = [{ 
	path: '/linearalgebra', 
	component: LinearAlgebra, 
	children: []
    }, {
	path: '/modellingandanalysisofdata',
	component: ModellingAndAnalysisOfData,
	children: [{
	    path: '/imageprocessing',
	    component: ImageProcessing
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
