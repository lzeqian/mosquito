<template lang="">
<div class="theme-group">
  <el-row class="block-col-1">
    <el-col :span="24">
      <!-- <span class="demonstration">click 激活</span> -->
      <el-dropdown trigger="click" :hide-on-click="true" class="dropdown-toggle theme-icons menu-btn" @command="handleCommand" >
        <span class="el-dropdown-link ">
        {{current_theme}}
      <i class="el-icon-caret-bottom el-icon--right"></i>
    </span>
        <el-dropdown-menu slot="dropdown" class="theme-dropdown-list">
             <el-dropdown-item class="theme-1 dropdown-item theme-icons" :command="key" v-for="(c, key) in themeList" :key="key">{{key}}</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </el-col>
  </el-row>
</div>

</template>

<script>
import {
    mapGetters
} from 'vuex'
export default {
  name: "theme",
  data() {
    return {
      theme_index: 1,
      current_theme: "classic",
      ulActive: false
    };
  },
    watch: {
        selectTheme(newv, oldv) {
            this.updateThemeDefaultValue(newv)
        },
    },
  computed: {
      ...mapGetters({
          'minder': 'getMinder'
      }),
      themeList(){
          return kityminder.Minder.getThemeList()
      },
    class_theme_index: function () {
      return "theme-" + this.theme_index;
    },
      selectTheme(){
          let currentTheme =this.minder.queryCommandValue &&
              this.minder.queryCommandValue("Theme");
          this.updateThemeDefaultValue(currentTheme)
          return currentTheme;
      }
  },

  methods: {
      updateThemeDefaultValue(currentTheme){
          this.current_theme=currentTheme||"";
      },
    handleCommand(command) {
        this.minder.execCommand('Theme', command);
        this.current_theme=command
    },
  },
};
</script>
<style>

</style>
