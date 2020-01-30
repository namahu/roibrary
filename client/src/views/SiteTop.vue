<template>
    <div class="siteTop">
        <div class="pageTitle">
            <h1>蔵書一覧</h1>
        </div>
        <div class="bookList">
            <div class="bookInfo">
                <img src="#," alt="bookImage" class="bookImage" />
                <span class="bookTitle">title</span>
            </div>
            <div class="bookInfo">
                <img src="#," alt="bookImage" class="bookImage" />
                <span class="bookTitle">title</span>
            </div>
            <div class="bookInfo">
                <img src="#," alt="bookImage" class="bookImage" />
                <span class="bookTitle">title</span>
            </div>
            <div class="bookInfo">
                <img src="#," alt="bookImage" class="bookImage" />
                <span class="bookTitle">title</span>
            </div>
        </div>
        <router-link to="/admin" class="baseButton linkButton">管理者メニューへ</router-link>
    </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'vue-property-decorator';
import axios from 'axios';

@Component
export default class SiteTop extends Vue {
    private mounted() {
        axios
            .get('http://localhost:8000/getbook')
            .then(res => console.log(res))
            .catch(err => console.log(err));
    }
}
</script>

<style lang="scss" scoped></style>
