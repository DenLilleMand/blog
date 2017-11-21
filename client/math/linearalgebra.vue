<template>
    <div> 
	Linear Algebra
	<ul> 
	    <li v-for="post in posts">
		<router-link v-bind:to="post.name"> {{post.name}} </router-link>
	    </li>
	</ul>
	<router-view> </router-view>
    </div>
</template>
<script>
export default {
  data () {
    return {
	posts: [ ]
    }
  }
}
</script>

<style>

</style>
