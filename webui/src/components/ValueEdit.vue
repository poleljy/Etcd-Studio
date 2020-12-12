<template>
    <div id="valueEdit">
        <editor ref="myEditor" @init="editorInit" :lang="lang" theme="chrome" width="100%" height="100%"></editor>
    </div>
</template>

<script>
export default {
    name: "ValueEdit",
    data(){
        return {
            lang: "json",
        }
    },
    methods:{
        editorInit: function () {
            require("brace/ext/language_tools"); //language extension prerequsite...
            require("brace/mode/html");
            require("brace/mode/javascript"); //language
            require("brace/mode/json"); //language
            require("brace/mode/less");
            require("brace/theme/chrome");
            require("brace/snippets/javascript"); //snippet

            this.editor = this.$refs.myEditor.editor;
            this.editor.getSession().setTabSize(4);
            this.editor.setReadOnly(true);
            this.editor.resize();
        },
        change() {
            this.lang = "text";
        },
        format(type) {
            if (type === "json") {
                let val = JSON.parse(this.editor.getValue());
                this.editor.setValue(JSON.stringify(val, null, 4));
                // this.editor.getSession().setMode('ace/mode/' + 'json');
            }
            this.editor.clearSelection();
            this.editor.navigateFileStart();
        }
    },
    components: {
        editor: require("vue2-ace-editor")
    }
}
</script>

<style lang="stylus" scoped>
    #valueEdit


</style>