<template>
  <div id="app">
    <div>
      <p>
        <router-link to="/">
          <el-button>Home</el-button>
        </router-link>
        <router-link to="/about">
          <el-button>About</el-button>
        </router-link>
      </p>
    </div>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: "app"
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
