<template>
    <div id="index"> 
	<MainMenu> </MainMenu>
	<ContactBox> </ContactBox>
    </div>
</template>
<script>
import MainMenu from 'components/mainmenu.vue'
import ContactBox from 'components/contactbox.vue'
import BackgroundGraphics from 'components/backgroundgraphics.vue'
import fontawesome from 'fontawesome/fontawesome-all.min.js'


export default {
    components : {
	MainMenu,
	ContactBox,
	BackgroundGraphics
    },
    data () {
	return {
	}
    }
}
</script>
<style>
</style>
