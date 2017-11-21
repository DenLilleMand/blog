<template>
    <div> 
	<mainmenu> </mainmenu>
	<ul> 
	    <li v-for="subject in subjects">
		<router-link v-bind:to="'/math/' + subject.name"> {{subject.name}} </router-link>
	    </li>
	</ul>
	<router-view> </router-view>
    </div>
</template>

<script>
export default {
  data () {
    return {
      subjects: [
	  { name: "Linear algebra" },
	  { name: "Modelling and analysis of data" }
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
