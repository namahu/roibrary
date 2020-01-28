<template>
    <div class="signin">
        <div class="pageTitle">
            <h1>ログイン画面</h1>
        </div>
        <button type="button" class="baseButton submitButton" @click="signinWithGoogle">Googleログイン</button>
    </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';
import * as firebase from 'firebase/app';
import 'firebase/auth';

@Component({})
export default class Signin extends Vue {
    private signinWithGoogle() {
        const provider: firebase.auth.GoogleAuthProvider = new firebase.auth.GoogleAuthProvider();
        firebase
            .auth()
            .signInWithPopup(provider)
            .then(result => {
                if (result.user === null) {
                    this.$router.push('/');
                    return;
                }

                result.user.getIdToken().then(idToken => {
                    localStorage.setItem('jwt', idToken.toString());
                });

                this.$router.push('/admin');
            })
            .catch(error => {
                console.log(error);
                this.$router.push('/');
            });
    }
}
</script>
