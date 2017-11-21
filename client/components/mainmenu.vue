<template>
    <div class="container-fluid"> 
	<div class="row">
	    <div class="col-md-1 offset-md-8">
		<a href="/">Home</a>
	    </div>
	    <div class="col-md-1">
		<a href="/math">Math</a>
	    </div>
	</div>
    </div>
</template>

<script>
export default {
  data () {
    return {
      msg: 'hello world!'
    }
  }
}
</script>

<style>
@import '~bootstrap/dist/css/bootstrap.min.css';
@import '~font-awesome/css/font-awesome.min.css';
.message {
  color: blue;
}
</style>
