import { initializeApp } from 'firebase/app';
import { firebaseConfig } from "./firebase-config";
import {
    getAuth,
    getRedirectResult,
    GoogleAuthProvider,
    signInWithRedirect
} from "firebase/auth";

const app = initializeApp(firebaseConfig);
const auth = getAuth();
const googleProvider = new GoogleAuthProvider();

window.App = {
    async init() {
        await this.handleRedirect()
    },
    async handleRedirect() {
        try {
            const result = await getRedirectResult(auth)
            if (result == null) return

            const credential = GoogleAuthProvider.credentialFromResult(result);
            const token = credential.accessToken;
            const user = result.user;

            const idToken = await user.getIdToken(true)

            console.log("USER: ", user, "\nACCESS TOKEN: ", token, "\nID TOKEN: ", idToken)

        } catch (error) {
            console.log(error)
        }
    },
    googleSignInRedirect() {
        signInWithRedirect(auth, googleProvider);
    }
}

document.addEventListener("DOMContentLoaded", async () => {
    await App.init()
})