<template>
    <div> 
	<mainmenu> </mainmenu>
	<ul> 
	    <li v-for="subject in subjects">
		<router-link v-bind:to="subject.name"> {{subject.name}} </router-link>
	    </li>
	</ul>
	<router-view> </router-view>
    </div>
</template>

<script>
export default {
  data () {
    return {
      subjects: [{ 
	      name: "Linear algebra", 
	      posts: []
	  }, { 
	      name: "Modelling and analysis of data", 
	      posts: [{
		  name: "Image processing",
		  date: "2017-11-21 22:17"
	      }]
	  }
      ]
    }
  }
}
</script>

<style>
.message {
  color: blue;
}
</style>
