<template>
  <div class="about">
    <h1>This is an about page</h1>
    <ul v-if="posts && posts.length">
      <li v-for="post of posts" :key="post.postID">
        <p><strong>{{post.createdBy}}</strong></p>
        <p>{{post.messageBody}}</p>
      </li>
    </ul>

    <ul v-if="errors && errors.length">
      <li v-for="error of errors" :key="error">
        {{error.message}}
      </li>
    </ul>
  </div>
</template>
<script>
import axios from "axios";

export default {
  name: "about",
  data() {
    return {
      posts: [],
      errors: []
    };
  },
  async created() {
    // promise version
    // axios
    //   .get(`http://127.0.0.1:8090/api/post`)
    //   .then(response => {
    //     this.posts = response.data;
    //   })
    //   .catch(e => {
    //     this.errors.push(e);
    //   });

    try {
      const response = await axios.get(`http://localhost:8090/api/post`);
      this.posts = response.data;
    } catch (e) {
      this.errors.push(e);
    }
  }
};
</script>
