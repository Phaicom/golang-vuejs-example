<template>
  <div class="about">
    <h1>This is an about page</h1>
    <el-button @click="addPost(10)">Add 10 posts</el-button>
    <el-button @click="reset()">Reset</el-button>
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
import { goAPI } from "../http-constants";

export default {
  name: "about",
  data() {
    return {
      posts: [],
      errors: []
    };
  },
  methods: {
    addPost(size) {
      goAPI
        .post(`api/post`, {
          size: size
        })
        .then(response => {
          this.posts = response.data;
        })
        .catch(e => {
          this.errors.push(e);
        });
    },
    reset() {
      goAPI
        .post(`api/reset`)
        .then(response => {
          this.posts = response.data;
        })
        .catch(e => {
          this.errors.push(e);
        });
    }
  },
  created() {
    // promise version
    goAPI
      .get(`api/post`)
      .then(response => {
        this.posts = response.data;
      })
      .catch(e => {
        this.errors.push(e);
      });

    // async await  version
    // try {
    //   const response = await axios.get(`http://localhost:8090/api/post`);
    //   this.posts = response.data;
    // } catch (e) {
    //   this.errors.push(e);
    // }
  }
};
</script>
