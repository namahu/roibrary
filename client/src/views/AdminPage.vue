<template>
    <div class="adminPage">
        <div class="pageTitle">
            <h1>管理者メニュー</h1>
        </div>
        <router-link to="/admin/newbook" class="baseButton linkbutton">本を追加する</router-link>
    </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';

@Component({})
export default class AdminPage extends Vue {}
</script>
