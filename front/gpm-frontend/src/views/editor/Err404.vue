<style scoped>
    .pdf-container{
        width: 100%;
        height: 100%;
        position: absolute;
        top:0;
        left:0;
        bottom:0;
        right:0;
    }
</style>
<template>

    <div style="height: 100%;background-image: url('/img/404.png')">
    </div>

</template>
<script>
    export default {
        data() {
            return {
                pdfSrc: null,
            }
        },
        computed: {
        },
        watch: {
        },
        methods: {

        },
        mounted() {
            this.initData()
        }
    }
</script>
