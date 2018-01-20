<template>
    <div id="backgroundgraphics"> 
	<canvas id="glcanvas">

	</canvas>
    </div>
</template>
<script>

export default {
    mounted() {
	const canvas = document.getElementById('glcanvas');
	canvas.width = window.innerWidth;
	canvas.height = window.innerHeight;

	const gl = canvas.getContext("webgl2");

	if (!gl) {
	    alert("Unable to initialize WebGL. Your browser or machine may not support it.");
	}
	// Set clear color to black, fully opaque
	gl.clearColor(0.0, 0.0, 0.0, 1.0);
	// Clear the color buffer with specified clear color
	gl.clear(gl.COLOR_BUFFER_BIT);
    },
    components : {
    },
    data () {
	return {
	}
    }
}
</script>
<style>
</style>
