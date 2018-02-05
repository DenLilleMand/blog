<template>
    <div id="contactbox"> 
	<div class="d-flex justify-content-center align-self-center">
	    <div class="border border-dark">
		<h1>Matti Andreas Nielsen</h1>
		<ul style="list-style: none;">
		    <li v-for="item in infos">
			<i :class="item.icon"></i>
			<span>{{ item.value }}</span>
		    </li>
		</ul>
		<ul style="list-style: none">
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
	    infos: [{
		name: "Email",
		icon: "fas fa-envelope-square",
		value: "matti@denlillemand.com"
	    }],
	    socialMedias: [{
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
