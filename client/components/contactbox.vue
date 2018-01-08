<template>
    <div id="contactbox"> 
	<div class="d-flex justify-content-center align-self-center">
	    <div class="border border-dark">
		<h1>Matti Andreas Nielsen</h1>
		<ul>
		    <li v-for="item in socialMedias"> 
			<i :class="item.icon"></i>
			<a target="_blank" :href="item.url">{{ item.name }}</a>
		    </li>
		</ul>
	    </div>
	</div>
    </div>
</template>
<script>

import MainMenu from 'components/mainmenu.vue'
export default {
    components : {
	MainMenu
    },
    data () {
	return {
	    socialMedias:[{
		name: "Github",
		icon: "fab fa-github",
		url: "https://github.com/denlillemand",
	    },{
		name: "Linkedin",
		icon: "fab fa-linkedin-in",
		url: "https://www.linkedin.com/in/matti-andreas-nielsen",
	    }]
	}
    }
}
</script>
<style>
</style>
